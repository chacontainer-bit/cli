package sqlite

import (
	"database/sql"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
)

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

	rows, err := s.db.Query(`SELECT status, COUNT(*) FROM assets WHERE tenant_id = ? GROUP BY status`, tenantID)
	if err != nil {
		return nil, err
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

	rows, err = s.db.Query(`SELECT type, COUNT(*) FROM assets WHERE tenant_id = ? GROUP BY type`, tenantID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t string
		var n int
		if err := rows.Scan(&t, &n); err != nil {
			rows.Close()
			return nil, err
		}
		stats.AssetsByType[t] = n
	}
	rows.Close()

	rows, err = s.db.Query(`SELECT status, COUNT(*) FROM shipments WHERE tenant_id = ? GROUP BY status`, tenantID)
	if err != nil {
		return nil, err
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

	today := time.Now().UTC().Format("2006-01-02")
	if err := s.db.QueryRow(
		`SELECT COUNT(*) FROM shipments WHERE tenant_id = ? AND substr(scheduled_pickup, 1, 10) = ?`,
		tenantID, today).Scan(&stats.ShipmentsToday); err != nil {
		return nil, err
	}

	if err := s.db.QueryRow(
		`SELECT COUNT(*) FROM shipments WHERE tenant_id = ? AND status NOT IN ('delivered','cancelled') AND scheduled_delivery < ?`,
		tenantID, nowStr()).Scan(&stats.LateShipments); err != nil {
		return nil, err
	}

	since := time.Now().Add(-24 * time.Hour).UTC().Format(time.RFC3339Nano)
	if err := s.db.QueryRow(
		`SELECT COUNT(*) FROM asset_events WHERE tenant_id = ? AND event_type IN ('scan','check_in','check_out') AND occurred_at >= ?`,
		tenantID, since).Scan(&stats.ScanEvents24h); err != nil {
		return nil, err
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
			(SELECT COUNT(*) FROM shipments sh WHERE sh.client_id = c.id AND sh.created_at >= ?) AS shipments_30d,
			(SELECT COUNT(*) FROM assets a WHERE a.client_id = c.id) AS assets
		FROM clients c WHERE c.tenant_id = ?
		ORDER BY shipments_30d DESC LIMIT 5`,
		time.Now().Add(-30*24*time.Hour).UTC().Format(time.RFC3339Nano), tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []handlers.ClientActivity
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
	return s.dailyCountTrend(tenantID, days, `SELECT substr(created_at,1,10) d, COUNT(*) FROM assets WHERE tenant_id = ? AND created_at >= ? GROUP BY d ORDER BY d`)
}

func (s *StatsStore) GetShipmentTrend(tenantID string, days int) ([]handlers.TrendPoint, error) {
	return s.dailyCountTrend(tenantID, days, `SELECT substr(created_at,1,10) d, COUNT(*) FROM shipments WHERE tenant_id = ? AND created_at >= ? GROUP BY d ORDER BY d`)
}

func (s *StatsStore) dailyCountTrend(tenantID string, days int, query string) ([]handlers.TrendPoint, error) {
	if days <= 0 {
		days = 30
	}
	since := time.Now().AddDate(0, 0, -days).UTC().Format(time.RFC3339Nano)
	rows, err := s.db.Query(query, tenantID, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []handlers.TrendPoint
	for rows.Next() {
		var p handlers.TrendPoint
		if err := rows.Scan(&p.Date, &p.Value); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func (s *StatsStore) GetPlantOccupancy(tenantID string) ([]handlers.PlantOccupancy, error) {
	rows, err := s.db.Query(`SELECT id, name, capacity FROM plants WHERE tenant_id = ?`, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	type row struct {
		id, name, capacity string
	}
	var plants []row
	for rows.Next() {
		var r row
		if err := rows.Scan(&r.id, &r.name, &r.capacity); err != nil {
			return nil, err
		}
		plants = append(plants, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	out := make([]handlers.PlantOccupancy, 0, len(plants))
	for _, p := range plants {
		var capacity struct {
			TotalSlots int `json:"total_slots"`
		}
		fromJSON(p.capacity, &capacity)

		var used int
		if err := s.db.QueryRow(
			`SELECT COUNT(*) FROM assets WHERE tenant_id = ? AND plant_id = ? AND status NOT IN ('retired','lost')`,
			tenantID, p.id).Scan(&used); err != nil {
			return nil, err
		}

		pct := 0.0
		if capacity.TotalSlots > 0 {
			pct = float64(used) / float64(capacity.TotalSlots) * 100
		}
		out = append(out, handlers.PlantOccupancy{
			PlantID: p.id, PlantName: p.name, TotalSlots: capacity.TotalSlots, UsedSlots: used, OccupancyPct: pct,
		})
	}
	return out, nil
}
