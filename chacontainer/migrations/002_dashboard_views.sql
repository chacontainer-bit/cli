-- CHACONTAINER - Dashboard Materialized Views & Analytics Queries
-- Refresh: CALL refresh_dashboard_views(); (scheduled every 5 min via pg_cron or cron job)

BEGIN;

-- ─── Asset inventory snapshot per plant ──────────────────────────────────────
CREATE MATERIALIZED VIEW mv_plant_inventory AS
SELECT
    a.tenant_id,
    a.plant_id,
    p.name               AS plant_name,
    a.type               AS asset_type,
    a.status,
    COUNT(*)             AS total,
    SUM(a.weight_kg)     AS total_weight_kg
FROM assets a
JOIN plants p ON p.id = a.plant_id
WHERE a.status NOT IN ('retired','lost')
GROUP BY a.tenant_id, a.plant_id, p.name, a.type, a.status;

CREATE UNIQUE INDEX ON mv_plant_inventory (tenant_id, plant_id, asset_type, status);

-- ─── Shipment KPIs last 30 days ───────────────────────────────────────────────
CREATE MATERIALIZED VIEW mv_shipment_kpis AS
SELECT
    tenant_id,
    COUNT(*)                                                             AS total,
    COUNT(*) FILTER (WHERE status = 'delivered')                        AS delivered,
    COUNT(*) FILTER (WHERE status IN ('in_transit','confirmed','draft')) AS active,
    COUNT(*) FILTER (WHERE status = 'exception')                        AS exceptions,
    COUNT(*) FILTER (WHERE scheduled_delivery < NOW()
                     AND status NOT IN ('delivered','cancelled'))        AS overdue,
    AVG(EXTRACT(EPOCH FROM (actual_delivery - actual_pickup))/3600)
        FILTER (WHERE actual_delivery IS NOT NULL
                AND actual_pickup IS NOT NULL)                           AS avg_transit_hours,
    SUM(total_weight_kg)                                                 AS total_weight_kg
FROM shipments
WHERE created_at >= NOW() - INTERVAL '30 days'
GROUP BY tenant_id;

CREATE UNIQUE INDEX ON mv_shipment_kpis (tenant_id);

-- ─── Daily scan activity ─────────────────────────────────────────────────────
CREATE MATERIALIZED VIEW mv_daily_scans AS
SELECT
    tenant_id,
    DATE_TRUNC('day', scanned_at)::DATE AS scan_date,
    COUNT(*)                             AS scan_count,
    COUNT(DISTINCT asset_id)             AS unique_assets,
    COUNT(DISTINCT scanned_by)           AS unique_users
FROM qr_scans
WHERE scanned_at >= NOW() - INTERVAL '90 days'
GROUP BY tenant_id, DATE_TRUNC('day', scanned_at)::DATE;

CREATE UNIQUE INDEX ON mv_daily_scans (tenant_id, scan_date);

-- ─── Refresh procedure ───────────────────────────────────────────────────────
CREATE OR REPLACE PROCEDURE refresh_dashboard_views()
LANGUAGE plpgsql AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_plant_inventory;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_shipment_kpis;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_daily_scans;
END;
$$;

COMMIT;

-- ─── Useful ad-hoc queries ───────────────────────────────────────────────────

-- Plant occupancy for a tenant
-- SELECT plant_id, plant_name,
--        SUM(total) AS used_slots,
--        (SELECT capacity->>'total_slots' FROM plants WHERE id = plant_id)::int AS total_slots
-- FROM mv_plant_inventory
-- WHERE tenant_id = $1
-- GROUP BY plant_id, plant_name;

-- Late shipments with client name
-- SELECT s.reference, s.scheduled_delivery, c.name AS client,
--        EXTRACT(DAY FROM NOW() - s.scheduled_delivery) AS days_late
-- FROM shipments s
-- JOIN clients c ON c.id = s.client_id
-- WHERE s.tenant_id = $1
--   AND s.scheduled_delivery < NOW()
--   AND s.status NOT IN ('delivered','cancelled')
-- ORDER BY days_late DESC;

-- Asset utilisation per type over last 7 days
-- SELECT a.type,
--        COUNT(*) FILTER (WHERE e.event_type = 'check_out') AS checkouts,
--        COUNT(*) FILTER (WHERE e.event_type = 'check_in')  AS checkins
-- FROM asset_events e
-- JOIN assets a ON a.id = e.asset_id
-- WHERE e.tenant_id = $1
--   AND e.occurred_at >= NOW() - INTERVAL '7 days'
-- GROUP BY a.type;
