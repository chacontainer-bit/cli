// Package sqlite implements the persistence layer for CHACONTAINER using an
// embedded, pure-Go SQLite database. It replaces the in-memory stubs that
// previously backed internal/api/handlers.
package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Open creates (if needed) and migrates the SQLite database at path.
// Use ":memory:" for tests.
func Open(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=foreign_keys(1)")
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}
	// modernc.org/sqlite is not safe for concurrent writers; serialize access.
	db.SetMaxOpenConns(1)

	if err := migrate(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return db, nil
}

var schema = []string{
	`CREATE TABLE IF NOT EXISTS tenants (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		slug TEXT NOT NULL UNIQUE,
		plan TEXT NOT NULL DEFAULT 'starter',
		status TEXT NOT NULL DEFAULT 'trial',
		contact_email TEXT NOT NULL,
		country TEXT NOT NULL DEFAULT 'MX',
		timezone TEXT NOT NULL DEFAULT 'America/Mexico_City',
		logo_url TEXT,
		settings TEXT NOT NULL DEFAULT '{}',
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		email TEXT NOT NULL,
		name TEXT NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'operator',
		plant_ids TEXT NOT NULL DEFAULT '[]',
		active INTEGER NOT NULL DEFAULT 1,
		last_login_at TEXT,
		created_at TEXT NOT NULL,
		UNIQUE (tenant_id, email)
	)`,
	`CREATE TABLE IF NOT EXISTS plants (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		code TEXT NOT NULL,
		name TEXT NOT NULL,
		address TEXT NOT NULL DEFAULT '{}',
		capacity TEXT NOT NULL DEFAULT '{}',
		status TEXT NOT NULL DEFAULT 'active',
		manager_id TEXT,
		geo TEXT NOT NULL DEFAULT '{}',
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		UNIQUE (tenant_id, code)
	)`,
	`CREATE TABLE IF NOT EXISTS plant_zones (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		plant_id TEXT NOT NULL REFERENCES plants(id) ON DELETE CASCADE,
		code TEXT NOT NULL,
		name TEXT NOT NULL,
		type TEXT NOT NULL DEFAULT 'warehouse',
		capacity INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL,
		UNIQUE (plant_id, code)
	)`,
	`CREATE TABLE IF NOT EXISTS clients (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		code TEXT NOT NULL,
		name TEXT NOT NULL,
		tax_id TEXT,
		type TEXT NOT NULL DEFAULT 'shipper',
		status TEXT NOT NULL DEFAULT 'active',
		industry TEXT,
		website TEXT,
		credit_limit REAL NOT NULL DEFAULT 0,
		payment_terms_days INTEGER NOT NULL DEFAULT 30,
		currency TEXT NOT NULL DEFAULT 'MXN',
		address TEXT NOT NULL DEFAULT '{}',
		airtable_id TEXT,
		erp_id TEXT,
		tags TEXT NOT NULL DEFAULT '[]',
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		UNIQUE (tenant_id, code)
	)`,
	`CREATE TABLE IF NOT EXISTS client_contacts (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		client_id TEXT NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT,
		phone TEXT,
		position TEXT,
		is_primary INTEGER NOT NULL DEFAULT 0,
		created_at TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS client_interactions (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		client_id TEXT NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
		contact_id TEXT,
		user_id TEXT NOT NULL,
		type TEXT NOT NULL,
		subject TEXT NOT NULL,
		notes TEXT,
		occurred_at TEXT NOT NULL,
		created_at TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS assets (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		plant_id TEXT NOT NULL,
		zone_id TEXT,
		name TEXT NOT NULL,
		code TEXT NOT NULL,
		airtable_id TEXT,
		qr_code TEXT NOT NULL,
		barcode TEXT,
		type TEXT NOT NULL,
		status TEXT NOT NULL DEFAULT 'available',
		description TEXT,
		manufacturer TEXT,
		model TEXT,
		serial_number TEXT,
		weight_kg REAL NOT NULL DEFAULT 0,
		dimensions TEXT NOT NULL DEFAULT '{}',
		purchase_date TEXT,
		purchase_cost REAL,
		client_id TEXT,
		metadata TEXT NOT NULL DEFAULT '{}',
		last_scan_at TEXT,
		last_scan_by TEXT,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		UNIQUE (tenant_id, code),
		UNIQUE (tenant_id, qr_code)
	)`,
	`CREATE TABLE IF NOT EXISTS asset_events (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		asset_id TEXT NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
		event_type TEXT NOT NULL,
		from_status TEXT,
		to_status TEXT,
		from_plant_id TEXT,
		to_plant_id TEXT,
		user_id TEXT NOT NULL,
		shipment_id TEXT,
		notes TEXT,
		geo TEXT,
		occurred_at TEXT NOT NULL,
		created_at TEXT NOT NULL
	)`,
	`CREATE INDEX IF NOT EXISTS idx_asset_events_asset ON asset_events (asset_id, occurred_at DESC)`,
	`CREATE TABLE IF NOT EXISTS shipments (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		reference TEXT NOT NULL,
		client_id TEXT NOT NULL,
		origin_plant_id TEXT NOT NULL,
		dest_plant_id TEXT,
		dest_address TEXT,
		status TEXT NOT NULL DEFAULT 'draft',
		mode TEXT NOT NULL DEFAULT 'road',
		carrier_name TEXT,
		carrier_ref TEXT,
		tracking_url TEXT,
		scheduled_pickup TEXT,
		actual_pickup TEXT,
		scheduled_delivery TEXT,
		actual_delivery TEXT,
		total_weight_kg REAL NOT NULL DEFAULT 0,
		total_items INTEGER NOT NULL DEFAULT 0,
		notes TEXT,
		customs_ref TEXT,
		erp_id TEXT,
		created_by TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		UNIQUE (tenant_id, reference)
	)`,
	`CREATE TABLE IF NOT EXISTS shipment_lines (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		shipment_id TEXT NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
		asset_id TEXT NOT NULL,
		quantity INTEGER NOT NULL DEFAULT 1,
		weight_kg REAL NOT NULL DEFAULT 0,
		notes TEXT,
		scanned_at TEXT,
		created_at TEXT NOT NULL,
		UNIQUE (shipment_id, asset_id)
	)`,
	`CREATE TABLE IF NOT EXISTS shipment_events (
		id TEXT PRIMARY KEY,
		tenant_id TEXT NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
		shipment_id TEXT NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
		from_status TEXT,
		to_status TEXT NOT NULL,
		location TEXT,
		description TEXT NOT NULL,
		user_id TEXT NOT NULL,
		occurred_at TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS webhook_log (
		id TEXT PRIMARY KEY,
		source TEXT NOT NULL,
		event TEXT NOT NULL,
		payload TEXT NOT NULL,
		received_at TEXT NOT NULL
	)`,
}

func migrate(db *sql.DB) error {
	for _, stmt := range schema {
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("exec %q: %w", stmt[:min(40, len(stmt))], err)
		}
	}
	return nil
}
