package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/domain/cycle"
)

// CycleStore is the persistence interface for reverse-logistics cycles.
type CycleStore interface {
	Create(tenantID string, c *cycle.Cycle) error
	GetByID(tenantID, id string) (*cycle.Cycle, error)
	List(tenantID string, filter CycleFilter) ([]*cycle.Cycle, int, error)
	RecordMilestone(tenantID, id string, milestone cycle.Milestone, at time.Time) (*cycle.Cycle, error)
}

type CycleFilter struct {
	PilotConfigID string
	AssetID       string
	Status        string
	Page          int
	PerPage       int
}

type CyclesHandler struct {
	store CycleStore
}

func NewCyclesHandler(store CycleStore) *CyclesHandler {
	return &CyclesHandler{store: store}
}

func (h *CyclesHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()
	filter := CycleFilter{
		PilotConfigID: q.Get("pilot_config_id"),
		AssetID:       q.Get("asset_id"),
		Status:        q.Get("status"),
		Page:          intParam(q.Get("page"), 1),
		PerPage:       intParam(q.Get("per_page"), 50),
	}

	cycles, total, err := h.store.List(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     cycles,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *CyclesHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())

	var c cycle.Cycle
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	c.TenantID = tenantID
	if c.Status == "" {
		c.Status = cycle.StatusOpen
	}

	if err := h.store.Create(tenantID, &c); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

func (h *CyclesHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	c, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "cycle not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

// RecordMilestone timestamps one step of the cycle (despacho, recepción,
// vacío, listo para retorno, recolección, recepción de retorno, inicio de
// inspección o liberación) and triggers recomputation of every day-metric
// that milestone unlocks.
func (h *CyclesHandler) RecordMilestone(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	var req struct {
		Milestone  string     `json:"milestone"`
		OccurredAt *time.Time `json:"occurred_at"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	at := time.Now()
	if req.OccurredAt != nil {
		at = *req.OccurredAt
	}

	updated, err := h.store.RecordMilestone(tenantID, id, cycle.Milestone(req.Milestone), at)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}
