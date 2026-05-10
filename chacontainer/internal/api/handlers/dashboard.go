package handlers

import (
	"net/http"
	"time"

	"github.com/chacontainer/backend/internal/api/middleware"
)

// DashboardStats aggregates KPIs for the operational dashboard.
type DashboardStats struct {
	AssetsByStatus    map[string]int     `json:"assets_by_status"`
	AssetsByType      map[string]int     `json:"assets_by_type"`
	ShipmentsByStatus map[string]int     `json:"shipments_by_status"`
	PlantOccupancy    []PlantOccupancy   `json:"plant_occupancy"`
	LateShipments     int                `json:"late_shipments"`
	TotalAssets       int                `json:"total_assets"`
	AssetsInTransit   int                `json:"assets_in_transit"`
	ShipmentsToday    int                `json:"shipments_today"`
	ScanEvents24h     int                `json:"scan_events_24h"`
	TopClients        []ClientActivity   `json:"top_clients"`
	GeneratedAt       time.Time          `json:"generated_at"`
}

type PlantOccupancy struct {
	PlantID       string  `json:"plant_id"`
	PlantName     string  `json:"plant_name"`
	TotalSlots    int     `json:"total_slots"`
	UsedSlots     int     `json:"used_slots"`
	OccupancyPct  float64 `json:"occupancy_pct"`
}

type ClientActivity struct {
	ClientID   string `json:"client_id"`
	ClientName string `json:"client_name"`
	Shipments  int    `json:"shipments_30d"`
	Assets     int    `json:"assets"`
}

// StatsStore is the analytics read model.
type StatsStore interface {
	GetDashboardStats(tenantID string) (*DashboardStats, error)
	GetAssetTrend(tenantID string, days int) ([]TrendPoint, error)
	GetShipmentTrend(tenantID string, days int) ([]TrendPoint, error)
	GetPlantOccupancy(tenantID string) ([]PlantOccupancy, error)
}

type TrendPoint struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

type DashboardHandler struct {
	stats StatsStore
}

func NewDashboardHandler(stats StatsStore) *DashboardHandler {
	return &DashboardHandler{stats: stats}
}

func (h *DashboardHandler) Summary(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	data, err := h.stats.GetDashboardStats(tenantID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, data)
}

func (h *DashboardHandler) AssetTrend(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	days := intParam(r.URL.Query().Get("days"), 30)
	points, err := h.stats.GetAssetTrend(tenantID, days)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": points})
}

func (h *DashboardHandler) ShipmentTrend(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	days := intParam(r.URL.Query().Get("days"), 30)
	points, err := h.stats.GetShipmentTrend(tenantID, days)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": points})
}

func (h *DashboardHandler) PlantOccupancy(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	occupancy, err := h.stats.GetPlantOccupancy(tenantID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": occupancy})
}
