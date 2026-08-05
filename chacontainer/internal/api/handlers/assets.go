package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chacontainer/backend/internal/api/middleware"
	"github.com/chacontainer/backend/internal/domain/asset"
)

// AssetStore is the persistence interface for assets.
type AssetStore interface {
	Create(tenantID string, a *asset.Asset) error
	GetByID(tenantID, id string) (*asset.Asset, error)
	GetByQR(tenantID, qrCode string) (*asset.Asset, error)
	List(tenantID string, filter AssetFilter) ([]*asset.Asset, int, error)
	Update(tenantID, id string, patch map[string]interface{}) (*asset.Asset, error)
	Delete(tenantID, id string) error
	RecordEvent(e *asset.AssetEvent) error
	ListEvents(tenantID, assetID string) ([]*asset.AssetEvent, error)
}

type AssetFilter struct {
	PlantID  string
	ZoneID   string
	Type     string
	Status   string
	ClientID string
	Search   string
	Page     int
	PerPage  int
}

type AssetsHandler struct {
	store AssetStore
}

func NewAssetsHandler(store AssetStore) *AssetsHandler {
	return &AssetsHandler{store: store}
}

func (h *AssetsHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()
	filter := AssetFilter{
		PlantID:  q.Get("plant_id"),
		ZoneID:   q.Get("zone_id"),
		Type:     q.Get("type"),
		Status:   q.Get("status"),
		ClientID: q.Get("client_id"),
		Search:   q.Get("q"),
		Page:     intParam(q.Get("page"), 1),
		PerPage:  intParam(q.Get("per_page"), 50),
	}

	assets, total, err := h.store.List(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     assets,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *AssetsHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())

	var a asset.Asset
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	a.TenantID = tenantID

	if err := h.store.Create(tenantID, &a); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.store.RecordEvent(&asset.AssetEvent{
		TenantID:   tenantID,
		AssetID:    a.ID,
		EventType:  asset.EventTypeCheckIn,
		ToStatus:   a.Status,
		UserID:     userID,
		OccurredAt: time.Now(),
		CreatedAt:  time.Now(),
	})

	writeJSON(w, http.StatusCreated, a)
}

func (h *AssetsHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	a, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "asset not found")
		return
	}
	writeJSON(w, http.StatusOK, a)
}

func (h *AssetsHandler) Update(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	updated, err := h.store.Update(tenantID, id, patch)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *AssetsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	if err := h.store.Delete(tenantID, id); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *AssetsHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	events, err := h.store.ListEvents(tenantID, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": events})
}

// ScanQR handles a mobile QR scan event.
func (h *AssetsHandler) ScanQR(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())

	var req struct {
		QRCode  string  `json:"qr_code"`
		Action  string  `json:"action"`
		PlantID string  `json:"plant_id"`
		ZoneID  string  `json:"zone_id"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
		Notes   string  `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	a, err := h.store.GetByQR(tenantID, req.QRCode)
	if err != nil {
		writeError(w, http.StatusNotFound, "asset not found for QR code")
		return
	}

	geo := &asset.GeoPoint{Lat: req.Lat, Lng: req.Lng}
	event := &asset.AssetEvent{
		TenantID:   tenantID,
		AssetID:    a.ID,
		EventType:  asset.EventType(req.Action),
		FromStatus: a.Status,
		ToStatus:   a.Status,
		UserID:     userID,
		Notes:      req.Notes,
		Geo:        geo,
		OccurredAt: time.Now(),
		CreatedAt:  time.Now(),
	}
	h.store.RecordEvent(event)

	// Update last scan
	h.store.Update(tenantID, a.ID, map[string]interface{}{
		"last_scan_at": time.Now(),
		"last_scan_by": userID,
		"plant_id":     req.PlantID,
		"zone_id":      req.ZoneID,
	})

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"asset": a,
		"event": event,
	})
}
