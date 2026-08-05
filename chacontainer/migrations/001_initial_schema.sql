-- CHACONTAINER - Initial Schema
-- PostgreSQL 15+
-- Multi-tenant: every table carries tenant_id with RLS enforced at the app layer.
-- Run with: psql $DATABASE_URL -f 001_initial_schema.sql
-- (no BEGIN/COMMIT here: the server's migration runner wraps each file in
-- its own transaction; psql --single-transaction does the same manually)

-- ─── Extensions ──────────────────────────────────────────────────────────────
-- No PostGIS: no column in this schema uses a native geometry type (geo/lat/lng
-- are plain JSONB/NUMERIC), so a plain postgres:15 image is enough to run this
-- fully offline on a local machine.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";    -- fuzzy search on names

-- ─── TENANTS ─────────────────────────────────────────────────────────────────
CREATE TABLE tenants (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR(200)  NOT NULL,
    slug          VARCHAR(100)  NOT NULL UNIQUE,
    plan          VARCHAR(30)   NOT NULL DEFAULT 'starter'
                    CHECK (plan IN ('starter','professional','enterprise')),
    status        VARCHAR(20)   NOT NULL DEFAULT 'trial'
                    CHECK (status IN ('active','suspended','trial')),
    contact_email VARCHAR(255)  NOT NULL,
    country       CHAR(2)       NOT NULL DEFAULT 'MX',
    timezone      VARCHAR(50)   NOT NULL DEFAULT 'America/Mexico_City',
    logo_url      TEXT,
    settings      JSONB         NOT NULL DEFAULT '{}',
    created_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

-- ─── USERS ───────────────────────────────────────────────────────────────────
CREATE TABLE users (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id     UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    email         VARCHAR(255)  NOT NULL,
    name          VARCHAR(200)  NOT NULL,
    password_hash VARCHAR(255),
    role          VARCHAR(20)   NOT NULL DEFAULT 'operator'
                    CHECK (role IN ('admin','manager','operator','viewer','api')),
    plant_ids     JSONB         NOT NULL DEFAULT '[]', -- array of plant UUIDs (driver-agnostic)
    active        BOOLEAN       NOT NULL DEFAULT TRUE,
    last_login_at TIMESTAMPTZ,
    created_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, email)
);

-- ─── PLANTS ──────────────────────────────────────────────────────────────────
CREATE TABLE plants (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id     UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    code          VARCHAR(20)   NOT NULL,
    name          VARCHAR(200)  NOT NULL,
    address       JSONB         NOT NULL DEFAULT '{}',
    capacity      JSONB         NOT NULL DEFAULT '{"total_slots":0,"used_slots":0,"max_weight_kg":0}',
    status        VARCHAR(20)   NOT NULL DEFAULT 'active'
                    CHECK (status IN ('active','inactive')),
    manager_id    UUID          REFERENCES users(id),
    geo           JSONB,
    created_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, code)
);

CREATE TABLE plant_zones (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id  UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    plant_id   UUID        NOT NULL REFERENCES plants(id) ON DELETE CASCADE,
    code       VARCHAR(20) NOT NULL,
    name       VARCHAR(200) NOT NULL,
    type       VARCHAR(20) NOT NULL DEFAULT 'warehouse'
                CHECK (type IN ('warehouse','dock','yard','production','quarantine')),
    capacity   INTEGER     NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (plant_id, code)
);

-- ─── CLIENTS (CRM) ───────────────────────────────────────────────────────────
CREATE TABLE clients (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id         UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    code              VARCHAR(30)   NOT NULL,
    name              VARCHAR(200)  NOT NULL,
    tax_id            VARCHAR(50),
    type              VARCHAR(30)   NOT NULL DEFAULT 'shipper'
                        CHECK (type IN ('shipper','receiver','forwarder','manufacturer','distributor')),
    status            VARCHAR(20)   NOT NULL DEFAULT 'active'
                        CHECK (status IN ('active','inactive','prospect')),
    industry          VARCHAR(100),
    website           VARCHAR(255),
    credit_limit      NUMERIC(14,2) NOT NULL DEFAULT 0,
    payment_terms_days INTEGER      NOT NULL DEFAULT 30,
    currency          CHAR(3)       NOT NULL DEFAULT 'MXN',
    address           JSONB         NOT NULL DEFAULT '{}',
    airtable_id       VARCHAR(50),
    erp_id            VARCHAR(50),
    tags              JSONB         NOT NULL DEFAULT '[]', -- array of strings (driver-agnostic)
    created_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, code)
);

CREATE TABLE client_contacts (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id  UUID         NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    client_id  UUID         NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    email      VARCHAR(255),
    phone      VARCHAR(30),
    position   VARCHAR(100),
    is_primary BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE client_interactions (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id   UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    client_id   UUID        NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    contact_id  UUID        REFERENCES client_contacts(id),
    user_id     UUID        NOT NULL REFERENCES users(id),
    type        VARCHAR(20) NOT NULL
                  CHECK (type IN ('call','email','meeting','visit','note')),
    subject     VARCHAR(300) NOT NULL,
    notes       TEXT,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── ASSETS ──────────────────────────────────────────────────────────────────
CREATE TABLE assets (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id     UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    plant_id      UUID          NOT NULL REFERENCES plants(id),
    zone_id       UUID          REFERENCES plant_zones(id),
    name          VARCHAR(200),
    code          VARCHAR(50)   NOT NULL,
    airtable_id   VARCHAR(50),
    qr_code       VARCHAR(200)  NOT NULL,
    barcode       VARCHAR(100),
    type          VARCHAR(30)   NOT NULL
                    CHECK (type IN ('container','pallet','forklift','rack','ibc','drum','trailer')),
    status        VARCHAR(20)   NOT NULL DEFAULT 'available'
                    CHECK (status IN ('available','in_use','in_transit','maintenance','retired','lost')),
    description   TEXT,
    manufacturer  VARCHAR(100),
    model         VARCHAR(100),
    serial_number VARCHAR(100),
    weight_kg     NUMERIC(10,2) NOT NULL DEFAULT 0,
    dimensions    JSONB         NOT NULL DEFAULT '{}',
    purchase_date DATE,
    purchase_cost NUMERIC(14,2),
    client_id     UUID          REFERENCES clients(id),
    metadata      JSONB         NOT NULL DEFAULT '{}',
    last_scan_at  TIMESTAMPTZ,
    last_scan_by  UUID          REFERENCES users(id),
    created_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, code),
    UNIQUE (tenant_id, qr_code)
);

-- Partial index for active assets per plant — fast occupancy queries
CREATE INDEX idx_assets_plant_status ON assets (tenant_id, plant_id, status)
    WHERE status NOT IN ('retired', 'lost');

CREATE INDEX idx_assets_qr ON assets (qr_code);
CREATE INDEX idx_assets_search ON assets USING gin(to_tsvector('spanish', name || ' ' || code))
    WHERE name IS NOT NULL;

-- ─── ASSET EVENTS ────────────────────────────────────────────────────────────
CREATE TABLE asset_events (
    id           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id    UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    asset_id     UUID        NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    event_type   VARCHAR(20) NOT NULL,
    from_status  VARCHAR(20),
    to_status    VARCHAR(20),
    from_plant_id UUID       REFERENCES plants(id),
    to_plant_id  UUID        REFERENCES plants(id),
    user_id      UUID        NOT NULL REFERENCES users(id),
    shipment_id  UUID,
    notes        TEXT,
    geo          JSONB,
    occurred_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_asset_events_asset ON asset_events (asset_id, occurred_at DESC);
CREATE INDEX idx_asset_events_tenant_day ON asset_events (tenant_id, occurred_at DESC);

-- ─── QR SCAN LOG ─────────────────────────────────────────────────────────────
CREATE TABLE qr_scans (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id   UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    asset_id    UUID        NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    qr_code     VARCHAR(200) NOT NULL,
    scanned_by  UUID        NOT NULL REFERENCES users(id),
    plant_id    UUID        REFERENCES plants(id),
    zone_id     UUID        REFERENCES plant_zones(id),
    lat         NUMERIC(10,7),
    lng         NUMERIC(10,7),
    device_id   VARCHAR(100),
    app_version VARCHAR(20),
    action      VARCHAR(20) NOT NULL DEFAULT 'inventory',
    notes       TEXT,
    scanned_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_qr_scans_asset ON qr_scans (asset_id, scanned_at DESC);
CREATE INDEX idx_qr_scans_24h ON qr_scans (tenant_id, scanned_at DESC);

-- ─── SHIPMENTS ───────────────────────────────────────────────────────────────
CREATE TABLE shipments (
    id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id           UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    reference           VARCHAR(50)   NOT NULL,
    client_id           UUID          NOT NULL REFERENCES clients(id),
    origin_plant_id     UUID          NOT NULL REFERENCES plants(id),
    dest_plant_id       UUID          REFERENCES plants(id),
    dest_address        JSONB,
    status              VARCHAR(20)   NOT NULL DEFAULT 'draft'
                          CHECK (status IN ('draft','confirmed','in_transit','at_customs','delivered','cancelled','exception')),
    mode                VARCHAR(20)   NOT NULL DEFAULT 'road'
                          CHECK (mode IN ('road','sea','air','rail','multimodal')),
    carrier_name        VARCHAR(200),
    carrier_ref         VARCHAR(100),
    tracking_url        TEXT,
    scheduled_pickup    TIMESTAMPTZ   NOT NULL,
    actual_pickup       TIMESTAMPTZ,
    scheduled_delivery  TIMESTAMPTZ   NOT NULL,
    actual_delivery     TIMESTAMPTZ,
    total_weight_kg     NUMERIC(10,2) NOT NULL DEFAULT 0,
    total_items         INTEGER       NOT NULL DEFAULT 0,
    notes               TEXT,
    customs_ref         VARCHAR(100),
    erp_id              VARCHAR(50),
    created_by          UUID          NOT NULL REFERENCES users(id),
    created_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, reference)
);

CREATE INDEX idx_shipments_status ON shipments (tenant_id, status);
CREATE INDEX idx_shipments_client ON shipments (tenant_id, client_id);
CREATE INDEX idx_shipments_late ON shipments (tenant_id, scheduled_delivery)
    WHERE status NOT IN ('delivered','cancelled');

CREATE TABLE shipment_lines (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id   UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    shipment_id UUID          NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    asset_id    UUID          NOT NULL REFERENCES assets(id),
    quantity    INTEGER       NOT NULL DEFAULT 1,
    weight_kg   NUMERIC(10,2) NOT NULL DEFAULT 0,
    notes       TEXT,
    scanned_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (shipment_id, asset_id)
);

CREATE TABLE shipment_events (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id   UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    shipment_id UUID        NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
    from_status VARCHAR(20),
    to_status   VARCHAR(20) NOT NULL,
    location    VARCHAR(200),
    description TEXT        NOT NULL,
    user_id     UUID        NOT NULL REFERENCES users(id),
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── MAINTENANCE ─────────────────────────────────────────────────────────────
CREATE TABLE maintenance_records (
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id      UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    asset_id       UUID          NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    type           VARCHAR(50)   NOT NULL,
    scheduled_at   TIMESTAMPTZ   NOT NULL,
    completed_at   TIMESTAMPTZ,
    technician_id  UUID          REFERENCES users(id),
    cost           NUMERIC(14,2) NOT NULL DEFAULT 0,
    notes          TEXT,
    created_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW()
);

-- ─── AIRTABLE SYNC LOG ───────────────────────────────────────────────────────
CREATE TABLE airtable_sync_log (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id     UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    entity_type   VARCHAR(30) NOT NULL,
    entity_id     UUID        NOT NULL,
    airtable_id   VARCHAR(50),
    direction     VARCHAR(10) NOT NULL CHECK (direction IN ('push','pull')),
    status        VARCHAR(10) NOT NULL CHECK (status IN ('ok','error')),
    error_message TEXT,
    synced_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- ─── UPDATED_AT trigger ──────────────────────────────────────────────────────
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;

CREATE TRIGGER trg_tenants_updated   BEFORE UPDATE ON tenants   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_plants_updated    BEFORE UPDATE ON plants    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_clients_updated   BEFORE UPDATE ON clients   FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_assets_updated    BEFORE UPDATE ON assets    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_shipments_updated BEFORE UPDATE ON shipments FOR EACH ROW EXECUTE FUNCTION set_updated_at();
