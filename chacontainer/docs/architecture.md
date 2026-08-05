# CHACONTAINER — Technical Architecture

> **Running this on your own machine only?** See
> [`local-deployment.md`](./local-deployment.md) for the single-command
> Docker Compose setup (local Postgres, no Redis, no cloud integrations
> wired in by default). The multi-tenant SaaS design below is the target
> architecture for a hosted deployment; the local stack is a deliberately
> smaller subset of it.

## System Overview

```
┌─────────────────────────────────────────────────────────────────────────┐
│                        CHACONTAINER SaaS Platform                       │
│                                                                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌────────────┐ │
│  │  Web App     │  │  Mobile App  │  │ Make.com     │  │ ERP System │ │
│  │  (React)     │  │  QR Scanner  │  │ Automations  │  │ (SAP/Odoo) │ │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘  └─────┬──────┘ │
│         │                 │                 │                 │        │
│         └─────────────────┴────────┬────────┴─────────────────┘        │
│                                    │ HTTPS / JWT                        │
│                          ┌─────────▼──────────┐                        │
│                          │   API Gateway /     │                        │
│                          │   Load Balancer     │                        │
│                          └─────────┬───────────┘                       │
│                                    │                                    │
│              ┌─────────────────────┼──────────────────────┐            │
│              │                     │                      │            │
│    ┌─────────▼───────┐  ┌──────────▼──────┐  ┌───────────▼──────┐    │
│    │  REST API        │  │  Webhook        │  │  Background       │   │
│    │  (Go / HTTP)     │  │  Receivers      │  │  Workers (cron)   │   │
│    │                  │  │  /webhooks/*    │  │  Python scripts   │   │
│    └─────────┬────────┘  └──────────┬──────┘  └───────────┬──────┘   │
│              │                      │                      │           │
│              └──────────────────────┼──────────────────────┘           │
│                                     │                                   │
│                          ┌──────────▼──────────┐                       │
│                          │   PostgreSQL 15      │                       │
│                          │   Multi-tenant DB    │                       │
│                          │   + Materialized     │                       │
│                          │     Views            │                       │
│                          └──────────────────────┘                      │
│                                                                         │
│                          ┌──────────────────────┐                      │
│                          │   Redis              │                       │
│                          │   Rate limiting /    │                       │
│                          │   Session cache      │                       │
│                          └──────────────────────┘                      │
└─────────────────────────────────────────────────────────────────────────┘
```

## Multi-Tenant Model

```
Tenant (company)
  ├── Users (roles: admin, manager, operator, viewer, api)
  ├── Plants (physical facilities)
  │     └── Zones (warehouse, dock, yard, production, quarantine)
  ├── Assets (containers, pallets, forklifts, racks, IBC, drums, trailers)
  │     ├── Asset Events (immutable audit trail)
  │     └── QR Scans (location + action log)
  ├── Clients (CRM)
  │     ├── Contacts
  │     └── Interactions (call, email, meeting, visit, note)
  ├── Shipments
  │     ├── Shipment Lines (which assets)
  │     └── Shipment Events (status timeline)
  └── Maintenance Records
```

Tenant isolation: every row carries `tenant_id`. All queries are scoped.

## API Endpoints

### Authentication
All endpoints (except webhooks) require `Authorization: Bearer <JWT>`.

JWT payload:
```json
{
  "tenant_id": "uuid",
  "user_id":   "uuid",
  "role":      "admin|manager|operator|viewer|api",
  "plant_ids": ["uuid"],
  "exp":       1234567890
}
```

### REST Routes

| Method | Path | Description |
|--------|------|-------------|
| GET | /health | Health check |
| GET | /api/v1/dashboard/summary | KPI summary |
| GET | /api/v1/dashboard/assets/trend | Asset count over time |
| GET | /api/v1/dashboard/shipments/trend | Shipments over time |
| GET | /api/v1/dashboard/plants/occupancy | Plant utilization |
| GET | /api/v1/assets | List assets (filter: plant_id, type, status, q) |
| POST | /api/v1/assets | Create asset |
| GET | /api/v1/assets/{id} | Get asset |
| PATCH | /api/v1/assets/{id} | Update asset |
| DELETE | /api/v1/assets/{id} | Delete asset |
| GET | /api/v1/assets/{id}/events | Asset event history |
| POST | /api/v1/assets/scan | Process QR scan |
| GET | /api/v1/shipments | List shipments |
| POST | /api/v1/shipments | Create shipment |
| GET | /api/v1/shipments/{id} | Get shipment |
| POST | /api/v1/shipments/{id}/transition | Change status |
| GET | /api/v1/shipments/{id}/lines | Get lines |
| POST | /api/v1/shipments/{id}/lines | Add line |
| GET | /api/v1/shipments/{id}/timeline | Status history |
| GET | /api/v1/clients | List clients |
| POST | /api/v1/clients | Create client |
| GET | /api/v1/clients/{id} | Get client |
| PATCH | /api/v1/clients/{id} | Update client |
| GET | /api/v1/clients/{id}/contacts | List contacts |
| POST | /api/v1/clients/{id}/contacts | Add contact |
| GET | /api/v1/clients/{id}/interactions | CRM timeline |
| POST | /api/v1/clients/{id}/interactions | Log interaction |
| GET | /api/v1/plants | List plants |
| POST | /api/v1/plants | Create plant |
| GET | /api/v1/plants/{id} | Get plant |
| PATCH | /api/v1/plants/{id} | Update plant |
| GET | /api/v1/plants/{id}/zones | List zones |
| POST | /api/v1/plants/{id}/zones | Create zone |
| POST | /webhooks/make | Make.com event receiver |
| POST | /webhooks/erp | ERP event receiver |
| POST | /webhooks/airtable | Airtable sync event |

## QR Code System

### Payload Format
```
{tenant_id}|{asset_id}|{asset_code}|1|{hmac_sha256_12chars}
```

### Scan Flow
```
Mobile App → scan QR
  → POST /api/v1/assets/scan
    {qr_code, action, plant_id, zone_id, lat, lng}
  → Verify signature
  → Resolve asset
  → Record QR scan event
  → Update asset.last_scan_at / plant_id / zone_id
  → Optionally fire Make.com webhook
  → Return asset details
```

### QR Generation (Python)
```bash
python scripts/qr_generator.py \
  --tenant <tenant_uuid> \
  --secret $QR_SECRET \
  --input assets.csv \
  --output ./qr_codes/
```

## Airtable Integration

### Tables Required in Airtable Base
| Table | Key Fields |
|-------|------------|
| Assets | Code, Type, Status, Weight (kg), Plant ID, CHACONTAINER ID |
| Clients | Code, Name, Type, Status, Credit Limit, CHACONTAINER ID |
| Shipments | Reference, Status, Mode, Carrier, Scheduled Delivery, CHACONTAINER ID |

### Sync Strategy
- **Push** (CHACONTAINER → Airtable): triggered after every write operation via async goroutine.
- **Pull** (Airtable → CHACONTAINER): run via cron every 15 min for field updates (status, credit limit).
- Conflict resolution: CHACONTAINER is source of truth for operational data; Airtable overrides for commercial fields (credit_limit, tags).

## Make.com Automation Scenarios

### Recommended Scenarios

| Scenario | Trigger | Actions |
|----------|---------|---------|
| Shipment Delivered | POST /webhooks/make (status=delivered) | Email client, update Airtable, post Slack |
| Late Shipment Alert | Cron every 30min via Python script | Email ops team, create Airtable task |
| QR Scan Check-Out | POST /webhooks/make (action=check_out) | Update ERP stock, log to Google Sheet |
| New Client | POST /webhooks/make (event=new_client) | Create Airtable CRM record, send welcome |
| Asset Maintenance Due | Cron weekly | Create maintenance work order in ERP |

### Webhook Payload (Make receives)
```json
{
  "source":       "chacontainer",
  "event":        "shipment_status_changed",
  "tenant_id":    "uuid",
  "shipment_ref": "SHP-2026-001",
  "from_status":  "in_transit",
  "to_status":    "delivered",
  "client_id":    "uuid",
  "occurred_at":  "2026-05-10T14:00:00Z"
}
```

## Database Schema Summary

```sql
tenants          -- SaaS customers
users            -- tenant users with roles
plants           -- physical facilities
plant_zones      -- sections within plants
clients          -- CRM accounts
client_contacts  -- contacts per client
client_interactions -- CRM touchpoints
assets           -- trackable physical units
asset_events     -- state change audit trail
qr_scans         -- every QR scan event
shipments        -- logistics movements
shipment_lines   -- assets per shipment
shipment_events  -- shipment status timeline
maintenance_records -- scheduled maintenance
airtable_sync_log   -- sync audit

-- Materialized views (refresh every 5 min):
mv_plant_inventory   -- asset counts per plant/type/status
mv_shipment_kpis     -- 30-day shipment KPIs per tenant
mv_daily_scans       -- scan activity over 90 days
```

## Security

| Layer | Mechanism |
|-------|-----------|
| API Auth | JWT HS256, 24h TTL |
| Tenant isolation | tenant_id on every query |
| Webhook validation | HMAC-SHA256 signature per source |
| QR integrity | HMAC-SHA256 payload signing |
| Transport | TLS 1.3 enforced |
| Passwords | bcrypt (cost 12) |
| Secrets | Environment variables, never in code |
| Rate limiting | Per tenant_id, Redis token bucket (target design; today's middleware is a passthrough - add this before exposing the API beyond your own machine) |

## Deployment

```
Production stack:
  - Go binary: Docker container, 2 replicas minimum
  - PostgreSQL 15: managed (e.g. RDS, Supabase, Neon)
  - Redis: ElastiCache or Upstash
  - Cron workers: Kubernetes CronJob or GitHub Actions scheduled
  - Reverse proxy: nginx or Caddy (TLS termination)
```

### Environment Variables
```bash
PORT=8080
DATABASE_URL=postgres://...
JWT_SECRET=<256-bit random>
AIRTABLE_API_KEY=pat...       # optional, unused unless you wire up the integration
AIRTABLE_BASE_ID=app...       # optional, unused unless you wire up the integration
ERP_BASE_URL=https://erp.client.com   # optional, unused unless you wire up the integration
ERP_API_KEY=...               # optional, unused unless you wire up the integration
MAKE_WEBHOOK_SECRET=...       # optional, only verifies inbound webhook signatures
QR_BASE_URL=https://app.chacontainer.com/qr
SEED_DEMO_DATA=true           # inserts demo tenant/user/data on first empty run
ENVIRONMENT=production
```

## Python Automation Scripts

| Script | Purpose | Schedule |
|--------|---------|----------|
| `qr_generator.py` | Bulk QR PNG generation | On-demand |
| `airtable_sync.py` | Bidirectional Airtable sync | */15 * * * * |
| `late_shipment_alert.py` | Overdue shipment alerts | */30 * * * * |
| `inventory_report.py` | Excel/CSV snapshot + email | 0 6 * * 1-5 |
