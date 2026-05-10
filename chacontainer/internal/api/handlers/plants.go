package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/domain/plant"
)

type PlantStore interface {
	Create(tenantID string, p *plant.Plant) error
	GetByID(tenantID, id string) (*plant.Plant, error)
	List(tenantID string) ([]*plant.Plant, error)
	Update(tenantID, id string, patch map[string]interface{}) (*plant.Plant, error)
	CreateZone(tenantID string, z *plant.Zone) error
	ListZones(tenantID, plantID string) ([]*plant.Zone, error)
}

type PlantsHandler struct {
	store PlantStore
}

func NewPlantsHandler(store PlantStore) *PlantsHandler {
	return &PlantsHandler{store: store}
}

func (h *PlantsHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	plants, err := h.store.List(tenantID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": plants})
}

func (h *PlantsHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	var p plant.Plant
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	p.TenantID = tenantID
	if err := h.store.Create(tenantID, &p); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, p)
}

func (h *PlantsHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")
	p, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "plant not found")
		return
	}
	writeJSON(w, http.StatusOK, p)
}

func (h *PlantsHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *PlantsHandler) ListZones(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	plantID := pathParam(r, "id")
	zones, err := h.store.ListZones(tenantID, plantID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": zones})
}

func (h *PlantsHandler) CreateZone(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	plantID := pathParam(r, "id")
	var z plant.Zone
	if err := json.NewDecoder(r.Body).Decode(&z); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	z.TenantID = tenantID
	z.PlantID = plantID
	if err := h.store.CreateZone(tenantID, &z); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, z)
}
