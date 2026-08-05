package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chacontainer/backend/internal/api"
	"github.com/chacontainer/backend/internal/config"
	"github.com/chacontainer/backend/internal/store/sqlite"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	db, err := sqlite.Open(cfg.DBPath)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()

	seeded, err := sqlite.SeedIfEmpty(db)
	if err != nil {
		log.Fatalf("seed: %v", err)
	}
	if seeded {
		log.Printf("seeded demo tenant — login with %s / %s", sqlite.DemoEmail, sqlite.DemoPassword)
	}

	router := api.NewRouter(cfg, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("CHACONTAINER API listening on :%d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		cancel()
		log.Fatalf("shutdown: %v", err)
	}
	cancel()
	log.Println("server stopped")
}
