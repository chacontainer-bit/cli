package sqlite

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/chacontainer/backend/internal/domain/tenant"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(db *sql.DB) *AuthStore {
	return &AuthStore{db: db}
}

var slugNonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts a company name into a URL/identifier-safe slug,
// disambiguating collisions with a numeric suffix.
func (s *AuthStore) Slugify(name string) (string, error) {
	base := strings.Trim(slugNonAlnum.ReplaceAllString(strings.ToLower(name), "-"), "-")
	if base == "" {
		base = "empresa"
	}
	slug := base
	for i := 2; ; i++ {
		var exists int
		if err := s.db.QueryRow(`SELECT COUNT(*) FROM tenants WHERE slug = ?`, slug).Scan(&exists); err != nil {
			return "", err
		}
		if exists == 0 {
			return slug, nil
		}
		slug = fmt.Sprintf("%s-%d", base, i)
	}
}

func (s *AuthStore) CreateTenant(t *tenant.Tenant) error {
	t.ID = newID()
	now := nowStr()
	if t.Plan == "" {
		t.Plan = tenant.PlanStarter
	}
	if t.Status == "" {
		t.Status = tenant.StatusTrial
	}
	if t.Country == "" {
		t.Country = "MX"
	}
	if t.Timezone == "" {
		t.Timezone = "America/Mexico_City"
	}
	_, err := s.db.Exec(
		`INSERT INTO tenants (id, name, slug, plan, status, contact_email, country, timezone, logo_url, settings, created_at, updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
		t.ID, t.Name, t.Slug, string(t.Plan), string(t.Status), t.ContactEmail, t.Country, t.Timezone, t.LogoURL, toJSON(t.Settings), now, now,
	)
	return err
}

// EmailTaken checks email uniqueness across ALL tenants (self-serve signup
// uses email-only login, so we enforce global uniqueness at the app layer
// even though the schema's constraint is scoped per-tenant).
func (s *AuthStore) EmailTaken(email string) (bool, error) {
	var n int
	err := s.db.QueryRow(`SELECT COUNT(*) FROM users WHERE email = ?`, strings.ToLower(email)).Scan(&n)
	return n > 0, err
}

func (s *AuthStore) CreateUser(u *tenant.User, passwordHash string) error {
	u.ID = newID()
	u.Email = strings.ToLower(u.Email)
	now := nowStr()
	u.CreatedAt = parseTime(now)
	if u.Role == "" {
		u.Role = tenant.RoleAdmin
	}
	_, err := s.db.Exec(
		`INSERT INTO users (id, tenant_id, email, name, password_hash, role, plant_ids, active, created_at)
		 VALUES (?,?,?,?,?,?,?,?,?)`,
		u.ID, u.TenantID, u.Email, u.Name, passwordHash, string(u.Role), toJSON(u.PlantIDs), boolToInt(true), now,
	)
	return err
}

// GetUserByEmail returns the user, its password hash, and its tenant.
func (s *AuthStore) GetUserByEmail(email string) (*tenant.User, string, *tenant.Tenant, error) {
	row := s.db.QueryRow(
		`SELECT u.id, u.tenant_id, u.email, u.name, u.password_hash, u.role, u.plant_ids, u.active, u.last_login_at, u.created_at,
		        t.id, t.name, t.slug, t.plan, t.status, t.contact_email, t.country, t.timezone, t.logo_url, t.settings, t.created_at, t.updated_at
		 FROM users u JOIN tenants t ON t.id = u.tenant_id
		 WHERE u.email = ?`, strings.ToLower(email))

	var u tenant.User
	var t tenant.Tenant
	var hash string
	var plantIDs, settings string
	var active int
	var lastLogin sql.NullString
	var uCreated, tCreated, tUpdated string
	var role, plan, status string

	err := row.Scan(
		&u.ID, &u.TenantID, &u.Email, &u.Name, &hash, &role, &plantIDs, &active, &lastLogin, &uCreated,
		&t.ID, &t.Name, &t.Slug, &plan, &status, &t.ContactEmail, &t.Country, &t.Timezone, &t.LogoURL, &settings, &tCreated, &tUpdated,
	)
	if err != nil {
		return nil, "", nil, err
	}

	u.Role = tenant.Role(role)
	u.Active = active == 1
	u.LastLoginAt = parseTimePtr(lastLogin)
	u.CreatedAt = parseTime(uCreated)
	fromJSON(plantIDs, &u.PlantIDs)

	t.Plan = tenant.Plan(plan)
	t.Status = tenant.Status(status)
	t.CreatedAt = parseTime(tCreated)
	t.UpdatedAt = parseTime(tUpdated)
	fromJSON(settings, &t.Settings)

	return &u, hash, &t, nil
}

func (s *AuthStore) TouchLastLogin(userID string) error {
	_, err := s.db.Exec(`UPDATE users SET last_login_at = ? WHERE id = ?`, nowStr(), userID)
	return err
}
