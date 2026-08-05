package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// DemoAdminEmail and DemoAdminPassword are the credentials created by
// SeedDemoData, printed to the server log so a first-time local run is
// usable immediately without a separate provisioning step.
const (
	DemoAdminEmail    = "admin@chacontainer.local"
	DemoAdminPassword = "chacontainer123"
)

// SeedDemoData creates one demo tenant, an admin user, a plant and a
// handful of assets/clients/shipments - but only the very first time the
// database is empty. It is safe to call on every server start.
func SeedDemoData(db *sql.DB) error {
	var tenantCount int
	if err := db.QueryRow(`SELECT COUNT(*) FROM tenants`).Scan(&tenantCount); err != nil {
		return fmt.Errorf("check existing tenants: %w", err)
	}
	if tenantCount > 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}
	defer tx.Rollback()

	var tenantID string
	if err := tx.QueryRow(`
		INSERT INTO tenants (name, slug, plan, status, contact_email, country, timezone)
		VALUES ('Demo Local', 'local', 'professional', 'active', $1, 'MX', 'America/Mexico_City')
		RETURNING id
	`, DemoAdminEmail).Scan(&tenantID); err != nil {
		return fmt.Errorf("insert demo tenant: %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(DemoAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash demo password: %w", err)
	}

	var userID string
	if err := tx.QueryRow(`
		INSERT INTO users (tenant_id, email, name, password_hash, role, active)
		VALUES ($1, $2, 'Administrador Local', $3, 'admin', TRUE)
		RETURNING id
	`, tenantID, DemoAdminEmail, string(passwordHash)).Scan(&userID); err != nil {
		return fmt.Errorf("insert demo user: %w", err)
	}

	var plantID string
	if err := tx.QueryRow(`
		INSERT INTO plants (tenant_id, code, name, address, capacity, status, manager_id)
		VALUES ($1, 'PL-01', 'Planta Principal',
			'{"street":"Av. Industrial 100","city":"Santiago","state":"RM","country":"CL","zip_code":"8320000"}'::jsonb,
			'{"total_slots":200,"used_slots":0,"max_weight_kg":50000}'::jsonb, 'active', $2)
		RETURNING id
	`, tenantID, userID).Scan(&plantID); err != nil {
		return fmt.Errorf("insert demo plant: %w", err)
	}

	var zoneID string
	if err := tx.QueryRow(`
		INSERT INTO plant_zones (tenant_id, plant_id, code, name, type, capacity)
		VALUES ($1, $2, 'BOD-N', 'Bodega Norte', 'warehouse', 100)
		RETURNING id
	`, tenantID, plantID).Scan(&zoneID); err != nil {
		return fmt.Errorf("insert demo zone: %w", err)
	}

	var clientID string
	if err := tx.QueryRow(`
		INSERT INTO clients (tenant_id, code, name, type, status, industry, credit_limit, currency, address)
		VALUES ($1, 'CLI-001', 'BASF RM', 'shipper', 'active', 'Química', 500000, 'MXN',
			'{"street":"Camino a Melipilla 8000","city":"Santiago","state":"RM","country":"CL","zip_code":"9020000"}'::jsonb)
		RETURNING id
	`, tenantID).Scan(&clientID); err != nil {
		return fmt.Errorf("insert demo client: %w", err)
	}

	assetSpecs := []struct {
		code, qr, assetType, status string
		weight                      float64
	}{
		{"CHC-0421", "demo-chc-0421", "ibc", "available", 65},
		{"CHC-0422", "demo-chc-0422", "drum", "in_use", 22},
		{"CHC-0430", "demo-chc-0430", "container", "in_transit", 1200},
		{"CHC-0433", "demo-chc-0433", "pallet", "available", 18},
		{"CHC-0439", "demo-chc-0439", "ibc", "maintenance", 65},
	}
	for _, a := range assetSpecs {
		if _, err := tx.Exec(`
			INSERT INTO assets (tenant_id, plant_id, zone_id, name, code, qr_code, type, status, weight_kg, client_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		`, tenantID, plantID, zoneID, a.code, a.code, a.qr, a.assetType, a.status, a.weight, clientID); err != nil {
			return fmt.Errorf("insert demo asset %s: %w", a.code, err)
		}
	}

	if _, err := tx.Exec(`
		INSERT INTO shipments (tenant_id, reference, client_id, origin_plant_id, status, mode,
			scheduled_pickup, scheduled_delivery, total_weight_kg, total_items, created_by)
		VALUES
		($1, 'SHP-2026-001', $2, $3, 'in_transit', 'road', NOW() - INTERVAL '2 days', NOW() + INTERVAL '1 day', 1200, 1, $4),
		($1, 'SHP-2026-002', $2, $3, 'delivered', 'road', NOW() - INTERVAL '5 days', NOW() - INTERVAL '3 days', 130, 3, $4)
	`, tenantID, clientID, plantID, userID); err != nil {
		return fmt.Errorf("insert demo shipments: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit demo seed: %w", err)
	}

	log.Printf("==================================================================")
	log.Printf("Seeded local demo data. Log in with:")
	log.Printf("  email:    %s", DemoAdminEmail)
	log.Printf("  password: %s", DemoAdminPassword)
	log.Printf("==================================================================")
	return nil
}
