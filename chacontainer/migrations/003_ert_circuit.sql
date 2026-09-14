-- CHACONTAINER - ERT Circuit: Reverse Logistics & Value Recovery
-- PostgreSQL 15+
-- Multi-tenant: every table carries tenant_id with RLS enforced at the app layer.
-- Adds the transactional model behind the ERT pilot ledger (CHC-ERT-WBK-LI-001):
-- pilot config, closed-loop cycles, per-asset inspections and the value recovery
-- ledger that feeds the ROI gate. Design references: AIAG RC-5 (returnable
-- containers management), GS1 GRAI/EPCIS (asset identity and event visibility).
-- Run with: psql $DATABASE_URL -f 003_ert_circuit.sql

BEGIN;

-- ─── PILOT CONFIG ────────────────────────────────────────────────────────────
-- One row per reverse-logistics circuit under governance (01_CONFIG sheet).
CREATE TABLE pilot_configs (
    id                            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id                     UUID          NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    code                          VARCHAR(50)   NOT NULL,
    client_id                     UUID          NOT NULL REFERENCES clients(id),
    circuit                       VARCHAR(300)  NOT NULL,
    currency                      CHAR(3)       NOT NULL DEFAULT 'MXN',
    start_date                    DATE,
    end_date                      DATE,
    sample_assets_target          INTEGER       NOT NULL DEFAULT 0,
    circuit_pool_total            INTEGER       NOT NULL DEFAULT 0,
    daily_demand                  INTEGER,
    replacement_cost_avg_mxn      NUMERIC(14,2),
    cycle_time_baseline_days      NUMERIC(8,2),
    sla_cycle_days                NUMERIC(8,2),
    dwell_critical_hours          NUMERIC(8,2),
    loss_rate_baseline_pct        NUMERIC(6,3),
    reverse_freight_baseline_mxn  NUMERIC(14,2),
    fixed_cost_mxn                NUMERIC(14,2) NOT NULL DEFAULT 0,
    traceability_cost_mxn         NUMERIC(14,2) NOT NULL DEFAULT 0,
    transport_cost_mxn            NUMERIC(14,2) NOT NULL DEFAULT 0,
    other_cost_mxn                NUMERIC(14,2) NOT NULL DEFAULT 0,
    roi_policy                    TEXT          NOT NULL DEFAULT
        'Solo ahorro realizado + costo evitado verificado. Excluye exposicion patrimonial recuperada y oportunidad potencial.',
    status                        VARCHAR(20)   NOT NULL DEFAULT 'draft'
                                    CHECK (status IN ('draft','active','closed')),
    created_at                    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    updated_at                    TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    UNIQUE (tenant_id, code)
);

CREATE INDEX idx_pilot_configs_client ON pilot_configs (tenant_id, client_id);

-- ─── CYCLES ──────────────────────────────────────────────────────────────────
-- One closed reverse-logistics loop per asset movement (05_CICLOS sheet).
-- Day-metric columns are populated by the application when each milestone is
-- recorded and when the cycle closes, mirroring shipment.IsLate()/DaysLate()
-- style computation rather than DB-generated columns.
CREATE TABLE cycles (
    id                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id                 UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    pilot_config_id           UUID        NOT NULL REFERENCES pilot_configs(id),
    asset_id                  UUID        NOT NULL REFERENCES assets(id),
    dispatched_at             TIMESTAMPTZ,
    received_by_client_at     TIMESTAMPTZ,
    emptied_at                TIMESTAMPTZ,
    ready_for_return_at       TIMESTAMPTZ,
    collected_at              TIMESTAMPTZ,
    received_return_at        TIMESTAMPTZ,
    inspection_started_at     TIMESTAMPTZ,
    released_at               TIMESTAMPTZ,
    outbound_transit_days     NUMERIC(8,2),
    client_use_days           NUMERIC(8,2),
    dwell_return_days         NUMERIC(8,2),
    return_transit_days       NUMERIC(8,2),
    reconditioning_days       NUMERIC(8,2),
    cycle_time_days           NUMERIC(8,2),
    sla_cycle_days            NUMERIC(8,2),
    meets_sla                 BOOLEAN,
    dwell_critical_hours      NUMERIC(8,2),
    exceeds_dwell             BOOLEAN,
    status                    VARCHAR(20) NOT NULL DEFAULT 'open'
                                CHECK (status IN ('open','closed','exception','no_conciliado')),
    notes                     TEXT,
    created_at                TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at                TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_cycles_asset ON cycles (tenant_id, asset_id, created_at DESC);
CREATE INDEX idx_cycles_pilot_status ON cycles (tenant_id, pilot_config_id, status);

-- ─── ASSET EVENTS: circuit linkage ────────────────────────────────────────────
-- Parity with 03_EVENTOS: ties every asset_event to a cycle and adds the
-- exception-handling fields the pilot's event capture protocol requires.
ALTER TABLE asset_events
    ADD COLUMN cycle_id          UUID REFERENCES cycles(id) ON DELETE SET NULL,
    ADD COLUMN exception         BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN exception_reason  TEXT,
    ADD COLUMN immediate_action  TEXT,
    ADD COLUMN validated         BOOLEAN NOT NULL DEFAULT FALSE;

CREATE INDEX idx_asset_events_cycle ON asset_events (cycle_id) WHERE cycle_id IS NOT NULL;

-- ─── INSPECTIONS ─────────────────────────────────────────────────────────────
-- Condition, damage and cost breakdown on return receipt (04_INSPECCION sheet).
CREATE TABLE inspections (
    id                           UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id                    UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    cycle_id                     UUID        REFERENCES cycles(id) ON DELETE SET NULL,
    asset_id                     UUID        NOT NULL REFERENCES assets(id),
    received_at                  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    inspector_id                 UUID        REFERENCES users(id),
    condition_in                 CHAR(1)     CHECK (condition_in IN ('A','B','C','D','E')),
    dirt_level                   VARCHAR(20),
    damage_type                  VARCHAR(50),
    severity                     VARCHAR(20) CHECK (severity IN ('ninguna','leve','moderada','severa')),
    action                       VARCHAR(30) NOT NULL
                                    CHECK (action IN ('LIBERAR','LIMPIEZA','REPARACION','LIMPIEZA_REPARACION','CUARENTENA','SCRAP')),
    labor_minutes                INTEGER       NOT NULL DEFAULT 0,
    cleaning_cost_mxn            NUMERIC(12,2) NOT NULL DEFAULT 0,
    repair_cost_mxn              NUMERIC(12,2) NOT NULL DEFAULT 0,
    parts_cost_mxn               NUMERIC(12,2) NOT NULL DEFAULT 0,
    scrap_value_mxn              NUMERIC(12,2) NOT NULL DEFAULT 0,
    total_intervention_cost_mxn  NUMERIC(12,2) NOT NULL DEFAULT 0,
    released_at                  TIMESTAMPTZ,
    condition_out                CHAR(1)     CHECK (condition_out IN ('A','B','C','D','E')),
    released                     BOOLEAN     NOT NULL DEFAULT FALSE,
    evidence_before_url          TEXT,
    evidence_after_url           TEXT,
    probable_cause                TEXT,
    probable_damage_node         VARCHAR(100),
    approved_by                  UUID        REFERENCES users(id),
    notes                        TEXT,
    created_at                   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_inspections_cycle ON inspections (cycle_id) WHERE cycle_id IS NOT NULL;
CREATE INDEX idx_inspections_asset ON inspections (tenant_id, asset_id, received_at DESC);

-- ─── VALUE RECOVERY LEDGER ───────────────────────────────────────────────────
-- Financial benefit ledger behind the ROI gate (06_VALUE_RECOVERY sheet).
-- Enforces the workbook's own rule: only a VALIDADO entry classified as
-- AHORRO_REALIZADO or COSTO_EVITADO_VERIFICADO may count toward ROI.
CREATE TABLE value_recovery_entries (
    id                         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id                  UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    pilot_config_id            UUID        NOT NULL REFERENCES pilot_configs(id),
    cycle_id                   UUID        REFERENCES cycles(id) ON DELETE SET NULL,
    asset_id                   UUID        REFERENCES assets(id),
    entry_date                 DATE        NOT NULL DEFAULT CURRENT_DATE,
    category                   VARCHAR(50) NOT NULL,
    financial_classification   VARCHAR(40) NOT NULL
                                  CHECK (financial_classification IN
                                    ('AHORRO_REALIZADO','COSTO_EVITADO_VERIFICADO',
                                     'EXPOSICION_PATRIMONIAL_RECUPERADA','OPORTUNIDAD_POTENCIAL')),
    quantity                   INTEGER       NOT NULL DEFAULT 1,
    base_cost_unit_mxn         NUMERIC(12,2) NOT NULL DEFAULT 0,
    actual_cost_unit_mxn       NUMERIC(12,2) NOT NULL DEFAULT 0,
    recovered_value_unit_mxn   NUMERIC(12,2) NOT NULL DEFAULT 0,
    gross_benefit_mxn          NUMERIC(14,2) NOT NULL DEFAULT 0,
    implementation_cost_mxn    NUMERIC(14,2) NOT NULL DEFAULT 0,
    net_benefit_mxn            NUMERIC(14,2) NOT NULL DEFAULT 0,
    include_in_roi             BOOLEAN       NOT NULL DEFAULT FALSE,
    evidence_url               TEXT,
    validator_id                UUID         REFERENCES users(id),
    validation_status          VARCHAR(20)   NOT NULL DEFAULT 'pendiente'
                                  CHECK (validation_status IN ('pendiente','validado','rechazado')),
    comments                   TEXT,
    created_at                 TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    CHECK (
        NOT include_in_roi
        OR (validation_status = 'validado'
            AND financial_classification IN ('AHORRO_REALIZADO','COSTO_EVITADO_VERIFICADO'))
    )
);

CREATE INDEX idx_value_recovery_pilot ON value_recovery_entries (tenant_id, pilot_config_id);
CREATE INDEX idx_value_recovery_roi ON value_recovery_entries (tenant_id, pilot_config_id)
    WHERE include_in_roi;

-- ─── UPDATED_AT triggers ─────────────────────────────────────────────────────
CREATE TRIGGER trg_pilot_configs_updated BEFORE UPDATE ON pilot_configs FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER trg_cycles_updated       BEFORE UPDATE ON cycles       FOR EACH ROW EXECUTE FUNCTION set_updated_at();

COMMIT;
