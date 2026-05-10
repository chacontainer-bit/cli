package tenant

import (
	"time"
)

type Plan string

const (
	PlanStarter    Plan = "starter"
	PlanProfessional Plan = "professional"
	PlanEnterprise Plan = "enterprise"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusTrial     Status = "trial"
)

// Tenant is the top-level SaaS customer (company using CHACONTAINER).
type Tenant struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Slug        string    `json:"slug" db:"slug"`
	Plan        Plan      `json:"plan" db:"plan"`
	Status      Status    `json:"status" db:"status"`
	ContactEmail string   `json:"contact_email" db:"contact_email"`
	Country     string    `json:"country" db:"country"`
	Timezone    string    `json:"timezone" db:"timezone"`
	LogoURL     string    `json:"logo_url,omitempty" db:"logo_url"`
	Settings    Settings  `json:"settings" db:"settings"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Settings struct {
	Currency          string `json:"currency"`
	DateFormat        string `json:"date_format"`
	DefaultPlantID    string `json:"default_plant_id,omitempty"`
	ERPEnabled        bool   `json:"erp_enabled"`
	AirtableSyncEnabled bool `json:"airtable_sync_enabled"`
	QRTrackingEnabled bool   `json:"qr_tracking_enabled"`
	MaxContainers     int    `json:"max_containers"`
	MaxPlants         int    `json:"max_plants"`
	MaxUsers          int    `json:"max_users"`
}

// User belongs to a Tenant and operates within it.
type User struct {
	ID         string    `json:"id" db:"id"`
	TenantID   string    `json:"tenant_id" db:"tenant_id"`
	Email      string    `json:"email" db:"email"`
	Name       string    `json:"name" db:"name"`
	Role       Role      `json:"role" db:"role"`
	PlantIDs   []string  `json:"plant_ids" db:"plant_ids"`
	Active     bool      `json:"active" db:"active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleManager   Role = "manager"
	RoleOperator  Role = "operator"
	RoleViewer    Role = "viewer"
	RoleAPI       Role = "api"
)

func (r Role) CanWrite() bool {
	return r == RoleAdmin || r == RoleManager || r == RoleOperator
}

func (r Role) CanAdmin() bool {
	return r == RoleAdmin
}
