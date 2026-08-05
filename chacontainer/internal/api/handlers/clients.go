package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chacontainer/backend/internal/api/middleware"
	"github.com/chacontainer/backend/internal/domain/client"
)

type ClientStore interface {
	Create(tenantID string, c *client.Client) error
	GetByID(tenantID, id string) (*client.Client, error)
	List(tenantID string, filter ClientFilter) ([]*client.Client, int, error)
	Update(tenantID, id string, patch map[string]interface{}) (*client.Client, error)
	ListContacts(tenantID, clientID string) ([]*client.Contact, error)
	CreateContact(tenantID string, c *client.Contact) error
	CreateInteraction(tenantID string, i *client.Interaction) error
	ListInteractions(tenantID, clientID string) ([]*client.Interaction, error)
}

type ClientFilter struct {
	Type    string
	Status  string
	Search  string
	Page    int
	PerPage int
}

type ClientsHandler struct {
	store ClientStore
}

func NewClientsHandler(store ClientStore) *ClientsHandler {
	return &ClientsHandler{store: store}
}

func (h *ClientsHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	q := r.URL.Query()
	filter := ClientFilter{
		Type:    q.Get("type"),
		Status:  q.Get("status"),
		Search:  q.Get("q"),
		Page:    intParam(q.Get("page"), 1),
		PerPage: intParam(q.Get("per_page"), 50),
	}
	clients, total, err := h.store.List(tenantID, filter)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"data":     clients,
		"total":    total,
		"page":     filter.Page,
		"per_page": filter.PerPage,
	})
}

func (h *ClientsHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	var c client.Client
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	c.TenantID = tenantID
	if err := h.store.Create(tenantID, &c); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

func (h *ClientsHandler) Get(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	id := pathParam(r, "id")
	c, err := h.store.GetByID(tenantID, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "client not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (h *ClientsHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *ClientsHandler) ListContacts(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	clientID := pathParam(r, "id")
	contacts, err := h.store.ListContacts(tenantID, clientID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": contacts})
}

func (h *ClientsHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	clientID := pathParam(r, "id")
	var c client.Contact
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	c.TenantID = tenantID
	c.ClientID = clientID
	if err := h.store.CreateContact(tenantID, &c); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, c)
}

func (h *ClientsHandler) CreateInteraction(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	userID := middleware.UserFromContext(r.Context())
	clientID := pathParam(r, "id")
	var i client.Interaction
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	i.TenantID = tenantID
	i.ClientID = clientID
	i.UserID = userID
	if err := h.store.CreateInteraction(tenantID, &i); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, i)
}

func (h *ClientsHandler) ListInteractions(w http.ResponseWriter, r *http.Request) {
	tenantID := middleware.TenantFromContext(r.Context())
	clientID := pathParam(r, "id")
	interactions, err := h.store.ListInteractions(tenantID, clientID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{"data": interactions})
}
