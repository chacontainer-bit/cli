// Package postgres implements every store interface declared in
// internal/api/handlers against a local PostgreSQL database, so the whole
// CHACONTAINER API can run with real persistence entirely on one machine -
// no managed cloud database required.
package postgres

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"sort"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/cli/cli/v2/chacontainer/migrations"
)

// Connect opens the database and retries for a while, since in a
// docker-compose stack the api container can start before postgres is ready
// to accept connections.
func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	var lastErr error
	for attempt := 1; attempt <= 20; attempt++ {
		if lastErr = db.Ping(); lastErr == nil {
			return db, nil
		}
		log.Printf("database not ready (attempt %d/20): %v", attempt, lastErr)
		time.Sleep(2 * time.Second)
	}
	db.Close()
	return nil, fmt.Errorf("database unreachable: %w", lastErr)
}

// Migrate applies every embedded *.sql file that hasn't run yet, in
// filename order, tracking progress in a schema_migrations table. It is
// idempotent and safe to run on every server start.
func Migrate(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			filename    TEXT PRIMARY KEY,
			applied_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`); err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	entries, err := fs.ReadDir(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)

	for _, name := range names {
		var already bool
		if err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE filename = $1)`, name).Scan(&already); err != nil {
			return fmt.Errorf("check migration %s: %w", name, err)
		}
		if already {
			continue
		}

		content, err := migrations.FS.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("begin tx for %s: %w", name, err)
		}
		if _, err := tx.Exec(string(content)); err != nil {
			tx.Rollback()
			return fmt.Errorf("apply migration %s: %w", name, err)
		}
		if _, err := tx.Exec(`INSERT INTO schema_migrations (filename) VALUES ($1)`, name); err != nil {
			tx.Rollback()
			return fmt.Errorf("record migration %s: %w", name, err)
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit migration %s: %w", name, err)
		}
		log.Printf("applied migration %s", name)
	}
	return nil
}
