package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
)

// StatsStore computes dashboard KPIs directly from the operational tables.
// A local single-tenant install has no need for the materialized views in
// migrations/002 (those exist for multi-tenant SaaS scale) - live aggregate
// queries are simpler to keep correct and are always up to date.
type StatsStore struct {
	db *sql.DB
}

func NewStatsStore(db *sql.DB) *StatsStore {
	return &StatsStore{db: db}
}

func (s *StatsStore) GetDashboardStats(tenantID string) (*handlers.DashboardStats, error) {
	stats := &handlers.DashboardStats{
		AssetsByStatus:    map[string]int{},
		AssetsByType:      map[string]int{},
		ShipmentsByStatus: map[string]int{},
		GeneratedAt:       time.Now(),
	}

	rows, err := s.db.Query(`SELECT status, COUNT(*) FROM assets WHERE tenant_id = $1 GROUP BY status`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("assets by status: %w", err)
	}
	for rows.Next() {
		var status string
		var n int
		if err := rows.Scan(&status, &n); err != nil {
			rows.Close()
			return nil, err
		}
		stats.AssetsByStatus[status] = n
		stats.TotalAssets += n
		if status == "in_transit" {
			stats.AssetsInTransit = n
		}
	}
	rows.Close()

	rows, err = s.db.Query(`SELECT type, COUNT(*) FROM assets WHERE tenant_id = $1 GROUP BY type`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("assets by type: %w", err)
	}
	for rows.Next() {
		var typ string
		var n int
		if err := rows.Scan(&typ, &n); err != nil {
			rows.Close()
			return nil, err
		}
		stats.AssetsByType[typ] = n
	}
	rows.Close()

	rows, err = s.db.Query(`SELECT status, COUNT(*) FROM shipments WHERE tenant_id = $1 GROUP BY status`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("shipments by status: %w", err)
	}
	for rows.Next() {
		var status string
		var n int
		if err := rows.Scan(&status, &n); err != nil {
			rows.Close()
			return nil, err
		}
		stats.ShipmentsByStatus[status] = n
	}
	rows.Close()

	if err := s.db.QueryRow(`
		SELECT COUNT(*) FROM shipments
		WHERE tenant_id = $1 AND status NOT IN ('delivered','cancelled') AND scheduled_delivery < NOW()
	`, tenantID).Scan(&stats.LateShipments); err != nil {
		return nil, fmt.Errorf("late shipments: %w", err)
	}

	if err := s.db.QueryRow(`
		SELECT COUNT(*) FROM shipments WHERE tenant_id = $1 AND scheduled_pickup::date = CURRENT_DATE
	`, tenantID).Scan(&stats.ShipmentsToday); err != nil {
		return nil, fmt.Errorf("shipments today: %w", err)
	}

	if err := s.db.QueryRow(`
		SELECT COUNT(*) FROM qr_scans WHERE tenant_id = $1 AND scanned_at >= NOW() - INTERVAL '24 hours'
	`, tenantID).Scan(&stats.ScanEvents24h); err != nil {
		return nil, fmt.Errorf("scan events 24h: %w", err)
	}

	occupancy, err := s.GetPlantOccupancy(tenantID)
	if err != nil {
		return nil, err
	}
	stats.PlantOccupancy = occupancy

	topClients, err := s.topClients(tenantID)
	if err != nil {
		return nil, err
	}
	stats.TopClients = topClients

	return stats, nil
}

func (s *StatsStore) topClients(tenantID string) ([]handlers.ClientActivity, error) {
	rows, err := s.db.Query(`
		SELECT c.id, c.name,
			COUNT(DISTINCT sh.id) FILTER (WHERE sh.created_at >= NOW() - INTERVAL '30 days') AS shipments_30d,
			COUNT(DISTINCT a.id) AS assets
		FROM clients c
		LEFT JOIN shipments sh ON sh.client_id = c.id AND sh.tenant_id = c.tenant_id
		LEFT JOIN assets a ON a.client_id = c.id AND a.tenant_id = c.tenant_id
		WHERE c.tenant_id = $1
		GROUP BY c.id, c.name
		ORDER BY shipments_30d DESC, assets DESC
		LIMIT 5
	`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("top clients: %w", err)
	}
	defer rows.Close()

	out := []handlers.ClientActivity{}
	for rows.Next() {
		var c handlers.ClientActivity
		if err := rows.Scan(&c.ClientID, &c.ClientName, &c.Shipments, &c.Assets); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func (s *StatsStore) GetAssetTrend(tenantID string, days int) ([]handlers.TrendPoint, error) {
	return s.dailyCountTrend(tenantID, days, `
		SELECT d::date AS day, COUNT(ae.id)
		FROM generate_series(CURRENT_DATE - ($2 - 1) * INTERVAL '1 day', CURRENT_DATE, INTERVAL '1 day') d
		LEFT JOIN asset_events ae ON ae.tenant_id = $1 AND ae.occurred_at::date = d::date
		GROUP BY d ORDER BY d
	`)
}

func (s *StatsStore) GetShipmentTrend(tenantID string, days int) ([]handlers.TrendPoint, error) {
	return s.dailyCountTrend(tenantID, days, `
		SELECT d::date AS day, COUNT(sh.id)
		FROM generate_series(CURRENT_DATE - ($2 - 1) * INTERVAL '1 day', CURRENT_DATE, INTERVAL '1 day') d
		LEFT JOIN shipments sh ON sh.tenant_id = $1 AND sh.created_at::date = d::date
		GROUP BY d ORDER BY d
	`)
}

func (s *StatsStore) dailyCountTrend(tenantID string, days int, query string) ([]handlers.TrendPoint, error) {
	if days <= 0 {
		days = 30
	}
	rows, err := s.db.Query(query, tenantID, days)
	if err != nil {
		return nil, fmt.Errorf("trend query: %w", err)
	}
	defer rows.Close()

	out := []handlers.TrendPoint{}
	for rows.Next() {
		var day time.Time
		var n int
		if err := rows.Scan(&day, &n); err != nil {
			return nil, err
		}
		out = append(out, handlers.TrendPoint{Date: day.Format("2006-01-02"), Value: n})
	}
	return out, rows.Err()
}

func (s *StatsStore) GetPlantOccupancy(tenantID string) ([]handlers.PlantOccupancy, error) {
	rows, err := s.db.Query(`
		SELECT p.id, p.name,
			COALESCE((p.capacity->>'total_slots')::int, 0) AS total_slots,
			COUNT(a.id) FILTER (WHERE a.status NOT IN ('retired','lost')) AS used_slots
		FROM plants p
		LEFT JOIN assets a ON a.plant_id = p.id AND a.tenant_id = p.tenant_id
		WHERE p.tenant_id = $1
		GROUP BY p.id, p.name, p.capacity
		ORDER BY p.name
	`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("plant occupancy: %w", err)
	}
	defer rows.Close()

	out := []handlers.PlantOccupancy{}
	for rows.Next() {
		var o handlers.PlantOccupancy
		var totalSlots, usedSlots int
		if err := rows.Scan(&o.PlantID, &o.PlantName, &totalSlots, &usedSlots); err != nil {
			return nil, err
		}
		o.TotalSlots, o.UsedSlots = totalSlots, usedSlots
		if totalSlots > 0 {
			o.OccupancyPct = float64(usedSlots) / float64(totalSlots) * 100
		}
		out = append(out, o)
	}
	return out, rows.Err()
}
