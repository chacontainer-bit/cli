package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/cli/cli/v2/chacontainer/internal/api/middleware"
)

// AuthUser is the subset of a user record needed to authenticate a login
// request and issue a JWT. It intentionally lives outside the tenant domain
// package so the password hash never has to flow through code paths that
// serialize user objects back to clients.
type AuthUser struct {
	ID           string
	TenantID     string
	Email        string
	Name         string
	Role         string
	PlantIDs     []string
	PasswordHash string
	Active       bool
}

// AuthStore looks up credentials for login.
type AuthStore interface {
	FindByEmail(email string) (*AuthUser, error)
	TouchLastLogin(userID string) error
}

type AuthHandler struct {
	store     AuthStore
	jwtSecret string
	tokenTTL  time.Duration
}

func NewAuthHandler(store AuthStore, jwtSecret string) *AuthHandler {
	return &AuthHandler{store: store, jwtSecret: jwtSecret, tokenTTL: 24 * time.Hour}
}

// Login authenticates against the local database and issues the same JWT
// format every other endpoint already expects via middleware.Auth. There is
// no external identity provider involved - this is the whole auth story for
// a private, single-machine install.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if req.Email == "" || req.Password == "" {
		writeError(w, http.StatusBadRequest, "email and password are required")
		return
	}

	u, err := h.store.FindByEmail(req.Email)
	if err != nil || u == nil || !u.Active {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}
	if u.PasswordHash == "" || bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) != nil {
		writeError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := middleware.GenerateToken(u.TenantID, u.ID, u.Role, u.PlantIDs, h.jwtSecret, h.tokenTTL)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not issue token")
		return
	}
	_ = h.store.TouchLastLogin(u.ID)

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_in": int(h.tokenTTL.Seconds()),
		"user": map[string]interface{}{
			"id":        u.ID,
			"tenant_id": u.TenantID,
			"email":     u.Email,
			"name":      u.Name,
			"role":      u.Role,
		},
	})
}
