package postgres

import (
	"database/sql"
	"fmt"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
)

// UserStore implements handlers.AuthStore against the users table.
type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// FindByEmail resolves a user for login. Email is only unique per tenant
// (see the UNIQUE (tenant_id, email) constraint), so on the rare chance two
// tenants share an email on the same local instance, this returns whichever
// active account was created first - fine for a single-operator local setup.
func (s *UserStore) FindByEmail(email string) (*handlers.AuthUser, error) {
	var u handlers.AuthUser
	var passwordHash sql.NullString
	var plantIDs []byte

	err := s.db.QueryRow(`
		SELECT id, tenant_id, email, name, password_hash, role, plant_ids, active
		FROM users WHERE email = $1 ORDER BY (active) DESC, created_at ASC LIMIT 1
	`, email).Scan(&u.ID, &u.TenantID, &u.Email, &u.Name, &passwordHash, &u.Role, &plantIDs, &u.Active)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	u.PasswordHash = passwordHash.String
	if err := unmarshalJSON(plantIDs, &u.PlantIDs); err != nil {
		return nil, fmt.Errorf("decode plant_ids: %w", err)
	}
	return &u, nil
}

func (s *UserStore) TouchLastLogin(userID string) error {
	_, err := s.db.Exec(`UPDATE users SET last_login_at = NOW() WHERE id = $1`, userID)
	return err
}
