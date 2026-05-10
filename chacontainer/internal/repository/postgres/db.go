package postgres

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Open returns a tuned *sql.DB for PostgreSQL.
func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	return db, nil
}

// newUUID generates a random UUID v4 without external dependencies.
func newUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// nullString returns sql.NullString.
func nullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// nullTime returns sql.NullTime.
func nullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// marshalJSON marshals v into JSON bytes, returning nil on error.
func marshalJSON(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}

// scanJSON scans a JSON []byte column into dst.
func scanJSON(data []byte, dst interface{}) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}
	return json.Unmarshal(data, dst)
}
