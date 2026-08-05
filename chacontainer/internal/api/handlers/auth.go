package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/chacontainer/backend/internal/domain/tenant"
	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 24 * time.Hour

// AuthStore is the persistence interface for tenant/user registration and login.
type AuthStore interface {
	Slugify(name string) (string, error)
	CreateTenant(t *tenant.Tenant) error
	EmailTaken(email string) (bool, error)
	CreateUser(u *tenant.User, passwordHash string) error
	GetUserByEmail(email string) (*tenant.User, string, *tenant.Tenant, error)
	TouchLastLogin(userID string) error
}

// TokenIssuer signs JWTs for authenticated sessions.
type TokenIssuer func(tenantID, userID, role string, plantIDs []string, ttl time.Duration) (string, error)

type AuthHandler struct {
	store  AuthStore
	issue  TokenIssuer
}

func NewAuthHandler(store AuthStore, issue TokenIssuer) *AuthHandler {
	return &AuthHandler{store: store, issue: issue}
}

type authResponse struct {
	Token  string       `json:"token"`
	User   *tenant.User `json:"user"`
	Tenant *tenant.Tenant `json:"tenant"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CompanyName string `json:"company_name"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	req.CompanyName = strings.TrimSpace(req.CompanyName)
	req.Name = strings.TrimSpace(req.Name)

	if req.CompanyName == "" || req.Name == "" || req.Email == "" {
		writeError(w, http.StatusBadRequest, "company_name, name and email are required")
		return
	}
	if len(req.Password) < 8 {
		writeError(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}
	if !strings.Contains(req.Email, "@") {
		writeError(w, http.StatusBadRequest, "invalid email")
		return
	}

	taken, err := h.store.EmailTaken(req.Email)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if taken {
		writeError(w, http.StatusConflict, "email already registered")
		return
	}

	slug, err := h.store.Slugify(req.CompanyName)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t := &tenant.Tenant{
		Name:         req.CompanyName,
		Slug:         slug,
		ContactEmail: req.Email,
		Settings: tenant.Settings{
			Currency:          "MXN",
			DateFormat:        "DD/MM/YYYY",
			QRTrackingEnabled: true,
			MaxContainers:     500,
			MaxPlants:         10,
			MaxUsers:          20,
		},
	}
	if err := h.store.CreateTenant(t); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not hash password")
		return
	}

	u := &tenant.User{
		TenantID: t.ID,
		Email:    req.Email,
		Name:     req.Name,
		Role:     tenant.RoleAdmin,
		PlantIDs: []string{},
	}
	if err := h.store.CreateUser(u, string(hash)); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.issue(t.ID, u.ID, string(u.Role), u.PlantIDs, tokenTTL)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not issue token")
		return
	}

	writeJSON(w, http.StatusCreated, authResponse{Token: token, User: u, Tenant: t})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	u, hash, t, err := h.store.GetUserByEmail(strings.TrimSpace(req.Email))
	if err != nil {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}
	if !u.Active {
		writeError(w, http.StatusForbidden, "user is deactivated")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	_ = h.store.TouchLastLogin(u.ID)

	token, err := h.issue(u.TenantID, u.ID, string(u.Role), u.PlantIDs, tokenTTL)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not issue token")
		return
	}

	writeJSON(w, http.StatusOK, authResponse{Token: token, User: u, Tenant: t})
}
