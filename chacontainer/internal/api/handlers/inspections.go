package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/domain/inspection"
)

// InspectionStore is the persistence interface for return inspections.
type InspectionStore interface {
	Create(tenantID string, insp *inspection.Inspection) error
	GetByID(tenantID, id string) (*inspection.Inspection, error)
	List(tenantID string, filter InspectionFilter) ([]*inspection.Inspection, int, error)
	Release(tenantID, id string, conditionOut inspection.Condition, releasedAt time.Time) (*inspection.Inspection, error)
}

type InspectionFilter struct {
	AssetID string
	CycleID string
	Action  string
	Page    int
	PerPage int
}

type InspectionsHandler struct {
	store InspectionStore
}

func NewInspectionsHandler(store InspectionStore) *InspectionsHandler {
	return &InspectionsHandler{store: store}
}

func (h *InspectionsHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()
	filter := InspectionFilter{
		AssetID: q.Get("asset_id"),
		CycleID: q.Get("cycle_id"),
		Action:  q.Get("action"),
		Page:    intParam(q.Get("page"), 1),
		PerPage: intParam(q.Get("per_page"), 50),
	}

	items, total, err := h.store.List(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     items,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *InspectionsHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())

	var insp inspection.Inspection
	if err := json.NewDecoder(r.Body).Decode(&insp); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	insp.TenantID = tenantID
	if insp.InspectorID == "" {
		insp.InspectorID = userID
	}
	if insp.ReceivedAt.IsZero() {
		insp.ReceivedAt = time.Now()
	}
	insp.ComputeTotalCost()

	if err := h.store.Create(tenantID, &insp); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, insp)
}

func (h *InspectionsHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	insp, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "inspection not found")
		return
	}
	writeJSON(w, http.StatusOK, insp)
}

// Release closes the inspection: exit condition, timestamp, and the
// LIBERAR disposition required before the asset re-enters circulation.
func (h *InspectionsHandler) Release(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	var req struct {
		ConditionOut string `json:"condition_out"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	updated, err := h.store.Release(tenantID, id, inspection.Condition(req.ConditionOut), time.Now())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}
