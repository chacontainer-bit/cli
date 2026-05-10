package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
)

// StatsStore implements handlers.StatsStore with real PostgreSQL aggregations.
type StatsStore struct{ db *sql.DB }

func NewStatsStore(db *sql.DB) *StatsStore { return &StatsStore{db: db} }

// ── GetDashboardStats ─────────────────────────────────────────────────────────

func (s *StatsStore) GetDashboardStats(tenantID string) (*handlers.DashboardStats, error) {
	stats := &handlers.DashboardStats{
		AssetsByStatus:    map[string]int{},
		AssetsByType:      map[string]int{},
		ShipmentsByStatus: map[string]int{},
		PlantOccupancy:    []handlers.PlantOccupancy{},
		TopClients:        []handlers.ClientActivity{},
		GeneratedAt:       time.Now(),
	}

	// Assets by status
	rows, err := s.db.Query(`
		SELECT status, COUNT(*) FROM assets
		WHERE tenant_id = $1 AND status NOT IN ('retired','lost')
		GROUP BY status`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("assets by status: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var status string
		var count int
		rows.Scan(&status, &count)
		stats.AssetsByStatus[status] = count
		stats.TotalAssets += count
		if status == "in_transit" {
			stats.AssetsInTransit = count
		}
	}

	// Assets by type
	rows2, err := s.db.Query(`
		SELECT type, COUNT(*) FROM assets
		WHERE tenant_id = $1 AND status NOT IN ('retired','lost')
		GROUP BY type`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("assets by type: %w", err)
	}
	defer rows2.Close()
	for rows2.Next() {
		var typ string
		var count int
		rows2.Scan(&typ, &count)
		stats.AssetsByType[typ] = count
	}

	// Shipments by status
	rows3, err := s.db.Query(`
		SELECT status, COUNT(*) FROM shipments
		WHERE tenant_id = $1
		GROUP BY status`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("shipments by status: %w", err)
	}
	defer rows3.Close()
	for rows3.Next() {
		var status string
		var count int
		rows3.Scan(&status, &count)
		stats.ShipmentsByStatus[status] = count
	}

	// Late shipments
	s.db.QueryRow(`
		SELECT COUNT(*) FROM shipments
		WHERE tenant_id = $1
		  AND scheduled_delivery < NOW()
		  AND status NOT IN ('delivered','cancelled')`, tenantID).
		Scan(&stats.LateShipments)

	// Shipments today
	s.db.QueryRow(`
		SELECT COUNT(*) FROM shipments
		WHERE tenant_id = $1
		  AND DATE_TRUNC('day', created_at) = CURRENT_DATE`, tenantID).
		Scan(&stats.ShipmentsToday)

	// QR scans last 24h
	s.db.QueryRow(`
		SELECT COUNT(*) FROM qr_scans
		WHERE tenant_id = $1 AND scanned_at >= NOW() - INTERVAL '24 hours'`, tenantID).
		Scan(&stats.ScanEvents24h)

	// Plant occupancy
	occupancy, err := s.GetPlantOccupancy(tenantID)
	if err == nil {
		stats.PlantOccupancy = occupancy
	}

	// Top 5 clients by shipment volume (30d)
	rows4, err := s.db.Query(`
		SELECT c.id, c.name,
		       COUNT(s.id) AS shipments_30d,
		       COUNT(DISTINCT a.id) AS assets
		FROM clients c
		LEFT JOIN shipments s ON s.client_id = c.id
		    AND s.tenant_id = c.tenant_id
		    AND s.created_at >= NOW() - INTERVAL '30 days'
		LEFT JOIN assets a ON a.client_id = c.id AND a.tenant_id = c.tenant_id
		WHERE c.tenant_id = $1 AND c.status = 'active'
		GROUP BY c.id, c.name
		ORDER BY shipments_30d DESC
		LIMIT 5`, tenantID)
	if err == nil {
		defer rows4.Close()
		for rows4.Next() {
			ca := handlers.ClientActivity{}
			rows4.Scan(&ca.ClientID, &ca.ClientName, &ca.Shipments, &ca.Assets)
			stats.TopClients = append(stats.TopClients, ca)
		}
	}

	return stats, nil
}

// ── GetAssetTrend ─────────────────────────────────────────────────────────────

func (s *StatsStore) GetAssetTrend(tenantID string, days int) ([]handlers.TrendPoint, error) {
	if days > 365 {
		days = 365
	}
	rows, err := s.db.Query(`
		SELECT TO_CHAR(d::date, 'YYYY-MM-DD') AS day,
		       COUNT(a.id) AS total
		FROM generate_series(
		         NOW() - ($2 || ' days')::interval,
		         NOW(),
		         '1 day'
		     ) d
		LEFT JOIN assets a
		       ON a.tenant_id = $1
		      AND DATE_TRUNC('day', a.created_at) = DATE_TRUNC('day', d)
		      AND a.status NOT IN ('retired','lost')
		GROUP BY d
		ORDER BY d`,
		tenantID, days,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTrendPoints(rows)
}

// ── GetShipmentTrend ──────────────────────────────────────────────────────────

func (s *StatsStore) GetShipmentTrend(tenantID string, days int) ([]handlers.TrendPoint, error) {
	if days > 365 {
		days = 365
	}
	rows, err := s.db.Query(`
		SELECT TO_CHAR(d::date, 'YYYY-MM-DD') AS day,
		       COUNT(sh.id) AS total
		FROM generate_series(
		         NOW() - ($2 || ' days')::interval,
		         NOW(),
		         '1 day'
		     ) d
		LEFT JOIN shipments sh
		       ON sh.tenant_id = $1
		      AND DATE_TRUNC('day', sh.created_at) = DATE_TRUNC('day', d)
		GROUP BY d
		ORDER BY d`,
		tenantID, days,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanTrendPoints(rows)
}

// ── GetPlantOccupancy ─────────────────────────────────────────────────────────

func (s *StatsStore) GetPlantOccupancy(tenantID string) ([]handlers.PlantOccupancy, error) {
	rows, err := s.db.Query(`
		SELECT
		    p.id::text,
		    p.name,
		    (p.capacity->>'total_slots')::int  AS total_slots,
		    COUNT(a.id)                         AS used_slots
		FROM plants p
		LEFT JOIN assets a
		       ON a.plant_id = p.id
		      AND a.tenant_id = p.tenant_id
		      AND a.status NOT IN ('retired','lost','in_transit')
		WHERE p.tenant_id = $1 AND p.status = 'active'
		GROUP BY p.id, p.name, p.capacity
		ORDER BY p.name`, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []handlers.PlantOccupancy
	for rows.Next() {
		po := handlers.PlantOccupancy{}
		rows.Scan(&po.PlantID, &po.PlantName, &po.TotalSlots, &po.UsedSlots)
		if po.TotalSlots > 0 {
			po.OccupancyPct = float64(po.UsedSlots) / float64(po.TotalSlots) * 100
		}
		out = append(out, po)
	}
	return out, rows.Err()
}

// ── helpers ───────────────────────────────────────────────────────────────────

func scanTrendPoints(rows *sql.Rows) ([]handlers.TrendPoint, error) {
	var out []handlers.TrendPoint
	for rows.Next() {
		var tp handlers.TrendPoint
		if err := rows.Scan(&tp.Date, &tp.Value); err != nil {
			return nil, err
		}
		out = append(out, tp)
	}
	return out, rows.Err()
}

var _ handlers.StatsStore = (*StatsStore)(nil)
