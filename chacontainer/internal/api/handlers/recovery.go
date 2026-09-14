package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
	"github.com/cli/cli/v2/chacontainer/internal/domain/recovery"
)

// RecoveryStore is the persistence interface for pilot config, the value
// recovery ledger and the computed ROI dashboard.
type RecoveryStore interface {
	CreatePilotConfig(tenantID string, pc *recovery.PilotConfig) error
	GetPilotConfig(tenantID, id string) (*recovery.PilotConfig, error)
	ListPilotConfigs(tenantID string) ([]*recovery.PilotConfig, error)

	CreateEntry(tenantID string, e *recovery.Entry) error
	ListEntries(tenantID string, filter EntryFilter) ([]*recovery.Entry, int, error)
	ValidateEntry(tenantID, id, validatorID string, status recovery.ValidationStatus) (*recovery.Entry, error)

	GetROISummary(tenantID, pilotConfigID string) (*recovery.Summary, error)
}

type EntryFilter struct {
	PilotConfigID  string
	Classification string
	IncludeInROI   *bool
	Page           int
	PerPage        int
}

type RecoveryHandler struct {
	store RecoveryStore
}

func NewRecoveryHandler(store RecoveryStore) *RecoveryHandler {
	return &RecoveryHandler{store: store}
}

// ── Pilot config ────────────────────────────────────────────────────────────

func (h *RecoveryHandler) ListPilotConfigs(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	configs, err := h.store.ListPilotConfigs(tenantID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": configs})
}

func (h *RecoveryHandler) CreatePilotConfig(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())

	var pc recovery.PilotConfig
	if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	pc.TenantID = tenantID
	if pc.Status == "" {
		pc.Status = "draft"
	}
	if pc.Currency == "" {
		pc.Currency = "MXN"
	}
	if pc.ROIPolicy == "" {
		pc.ROIPolicy = "Solo ahorro realizado + costo evitado verificado. Excluye exposicion patrimonial recuperada y oportunidad potencial."
	}

	if err := h.store.CreatePilotConfig(tenantID, &pc); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, pc)
}

func (h *RecoveryHandler) GetPilotConfig(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	pc, err := h.store.GetPilotConfig(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "pilot config not found")
		return
	}
	writeJSON(w, http.StatusOK, pc)
}

// ROISummary computes the 07_ROI dashboard for one circuit: KPIs, the
// benefit actually eligible under the ROI gate, and the four-criterion
// scale-up decision gate.
func (h *RecoveryHandler) ROISummary(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")

	summary, err := h.store.GetROISummary(tenantID, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, summary)
}

// ── Value recovery ledger ───────────────────────────────────────────────────

func (h *RecoveryHandler) ListEntries(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()
	filter := EntryFilter{
		PilotConfigID:  q.Get("pilot_config_id"),
		Classification: q.Get("classification"),
		Page:           intParam(q.Get("page"), 1),
		PerPage:        intParam(q.Get("per_page"), 50),
	}
	if v := q.Get("include_in_roi"); v != "" {
		b := v == "true"
		filter.IncludeInROI = &b
	}

	entries, total, err := h.store.ListEntries(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     entries,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *RecoveryHandler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())

	var e recovery.Entry
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	e.TenantID = tenantID
	if e.EntryDate.IsZero() {
		e.EntryDate = time.Now()
	}
	if e.ValidationStatus == "" {
		e.ValidationStatus = recovery.ValidationPending
	}
	// Never trust the caller's include_in_roi flag: only the gate rule
	// (validado + clasificación elegible) may turn it on.
	e.IncludeInROI = e.EligibleForROI()

	if err := h.store.CreateEntry(tenantID, &e); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, e)
}

// ValidateEntry approves or rejects a ledger entry. Approving an entry
// already classified as AHORRO_REALIZADO or COSTO_EVITADO_VERIFICADO is
// what actually unlocks it for the ROI calculation — never the create call.
func (h *RecoveryHandler) ValidateEntry(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())
	id := pathParam(r, "id")

	var req struct {
		Status string `json:"status"` // "validado" | "rechazado"
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	updated, err := h.store.ValidateEntry(tenantID, id, userID, recovery.ValidationStatus(req.Status))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, updated)
}
