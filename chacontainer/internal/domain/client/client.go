package client

import "time"

type ClientType string

const (
	ClientTypeShipper      ClientType = "shipper"
	ClientTypeReceiver     ClientType = "receiver"
	ClientTypeForwarder    ClientType = "forwarder"
	ClientTypeManufacturer ClientType = "manufacturer"
	ClientTypeDistributor  ClientType = "distributor"
)

type ClientStatus string

const (
	ClientStatusActive   ClientStatus = "active"
	ClientStatusInactive ClientStatus = "inactive"
	ClientStatusProspect ClientStatus = "prospect"
)

// Client is an end-customer in the CRM.
type Client struct {
	ID           string       `json:"id" db:"id"`
	TenantID     string       `json:"tenant_id" db:"tenant_id"`
	Code         string       `json:"code" db:"code"`
	Name         string       `json:"name" db:"name"`
	TaxID        string       `json:"tax_id,omitempty" db:"tax_id"`
	Type         ClientType   `json:"type" db:"type"`
	Status       ClientStatus `json:"status" db:"status"`
	Industry     string       `json:"industry,omitempty" db:"industry"`
	Website      string       `json:"website,omitempty" db:"website"`
	CreditLimit  float64      `json:"credit_limit" db:"credit_limit"`
	PaymentTerms int          `json:"payment_terms_days" db:"payment_terms_days"`
	Currency     string       `json:"currency" db:"currency"`
	Address      Address      `json:"address" db:"address"`
	AirtableID   string       `json:"airtable_id,omitempty" db:"airtable_id"`
	ERPID        string       `json:"erp_id,omitempty" db:"erp_id"`
	Tags         []string     `json:"tags,omitempty" db:"tags"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

// Contact is a person associated with a Client.
type Contact struct {
	ID        string    `json:"id" db:"id"`
	TenantID  string    `json:"tenant_id" db:"tenant_id"`
	ClientID  string    `json:"client_id" db:"client_id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone,omitempty" db:"phone"`
	Position  string    `json:"position,omitempty" db:"position"`
	IsPrimary bool      `json:"is_primary" db:"is_primary"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (c Contact) FullName() string {
	return c.FirstName + " " + c.LastName
}

// Interaction logs every CRM touchpoint.
type Interaction struct {
	ID         string          `json:"id" db:"id"`
	TenantID   string          `json:"tenant_id" db:"tenant_id"`
	ClientID   string          `json:"client_id" db:"client_id"`
	ContactID  string          `json:"contact_id,omitempty" db:"contact_id"`
	UserID     string          `json:"user_id" db:"user_id"`
	Type       InteractionType `json:"type" db:"type"`
	Subject    string          `json:"subject" db:"subject"`
	Notes      string          `json:"notes,omitempty" db:"notes"`
	OccurredAt time.Time       `json:"occurred_at" db:"occurred_at"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
}

type InteractionType string

const (
	InteractionTypeCall    InteractionType = "call"
	InteractionTypeEmail   InteractionType = "email"
	InteractionTypeMeeting InteractionType = "meeting"
	InteractionTypeVisit   InteractionType = "visit"
	InteractionTypeNote    InteractionType = "note"
)
