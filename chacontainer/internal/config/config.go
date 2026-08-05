package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DatabaseURL string
	JWTSecret   string
	Environment string

	AirtableAPIKey string
	AirtableBaseID string

	ERPBaseURL string
	ERPAPIKey  string

	MakeWebhookSecret string

	QRBaseURL string

	// SeedDemoData bootstraps a demo tenant/admin user/plant/assets on first
	// run so a fresh local install has something to look at immediately.
	// It only ever writes when the tenants table is empty (see
	// postgres.SeedDemoData), so leaving this on is safe after the first run.
	SeedDemoData bool
}

func Load() (*Config, error) {
	port := 8080
	if v := os.Getenv("PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT: %w", err)
		}
		port = p
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://chacontainer:chacontainer@localhost:5432/chacontainer?sslmode=disable"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	seedDemoData := true
	if v := os.Getenv("SEED_DEMO_DATA"); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf("invalid SEED_DEMO_DATA: %w", err)
		}
		seedDemoData = b
	}

	return &Config{
		Port:              port,
		DatabaseURL:       dbURL,
		JWTSecret:         jwtSecret,
		Environment:       getEnv("ENVIRONMENT", "development"),
		AirtableAPIKey:    os.Getenv("AIRTABLE_API_KEY"),
		AirtableBaseID:    os.Getenv("AIRTABLE_BASE_ID"),
		ERPBaseURL:        os.Getenv("ERP_BASE_URL"),
		ERPAPIKey:         os.Getenv("ERP_API_KEY"),
		MakeWebhookSecret: os.Getenv("MAKE_WEBHOOK_SECRET"),
		QRBaseURL:         getEnv("QR_BASE_URL", "https://api.chacontainer.com/qr"),
		SeedDemoData:      seedDemoData,
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
