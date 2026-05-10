package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/domain/shipment"
)

type ShipmentStore interface {
	Create(tenantID string, s *shipment.Shipment) error
	GetByID(tenantID, id string) (*shipment.Shipment, error)
	GetByReference(tenantID, ref string) (*shipment.Shipment, error)
	List(tenantID string, filter ShipmentFilter) ([]*shipment.Shipment, int, error)
	UpdateStatus(tenantID, id string, status shipment.ShipmentStatus, userID, description string) error
	AddLine(tenantID string, line *shipment.ShipmentLine) error
	ListLines(tenantID, shipmentID string) ([]*shipment.ShipmentLine, error)
	ListEvents(tenantID, shipmentID string) ([]*shipment.ShipmentEvent, error)
}

type ShipmentFilter struct {
	ClientID string
	PlantID  string
	Status   string
	From     *time.Time
	To       *time.Time
	Page     int
	PerPage  int
}

type ShipmentsHandler struct {
	store ShipmentStore
}

func NewShipmentsHandler(store ShipmentStore) *ShipmentsHandler {
	return &ShipmentsHandler{store: store}
}

func (h *ShipmentsHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()

	filter := ShipmentFilter{
		ClientID: q.Get("client_id"),
		PlantID:  q.Get("plant_id"),
		Status:   q.Get("status"),
		Page:     intParam(q.Get("page"), 1),
		PerPage:  intParam(q.Get("per_page"), 50),
	}

	shipments, total, err := h.store.List(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     shipments,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *ShipmentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())

	var s shipment.Shipment
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	s.TenantID = tenantID
	s.CreatedBy = userID
	s.Status = shipment.StatusDraft

	if err := h.store.Create(tenantID, &s); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, s)
}

func (h *ShipmentsHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	s, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "shipment not found")
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *ShipmentsHandler) Transition(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())
	id := pathParam(r, "id")

	var req struct {
		Status      string `json:"status"`
		Description string `json:"description"`
		Location    string `json:"location"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	status := shipment.ShipmentStatus(req.Status)
	if err := h.store.UpdateStatus(tenantID, id, status, userID, req.Description); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s, _ := h.store.GetByID(tenantID, id)
	writeJSON(w, http.StatusOK, s)
}

func (h *ShipmentsHandler) AddLine(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	var line shipment.ShipmentLine
	if err := json.NewDecoder(r.Body).Decode(&line); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	line.TenantID = tenantID
	line.ShipmentID = id

	if err := h.store.AddLine(tenantID, &line); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, line)
}

func (h *ShipmentsHandler) GetLines(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	lines, err := h.store.ListLines(tenantID, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": lines})
}

func (h *ShipmentsHandler) GetTimeline(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	events, err := h.store.ListEvents(tenantID, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": events})
}
