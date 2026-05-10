package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/client"
	"github.com/lib/pq"
)

// ClientStore implements handlers.ClientStore against PostgreSQL.
type ClientStore struct{ db *sql.DB }

func NewClientStore(db *sql.DB) *ClientStore { return &ClientStore{db: db} }

// ── Create ────────────────────────────────────────────────────────────────────

func (s *ClientStore) Create(tenantID string, c *client.Client) error {
	c.ID = newUUID()
	c.TenantID = tenantID
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	if c.Status == "" {
		c.Status = client.ClientStatusActive
	}
	if c.Currency == "" {
		c.Currency = "MXN"
	}

	addrJSON := marshalJSON(c.Address)

	_, err := s.db.Exec(`
		INSERT INTO clients (
			id, tenant_id, code, name, tax_id,
			type, status, industry, website,
			credit_limit, payment_terms_days, currency, address,
			airtable_id, erp_id, tags,
			created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,
			$6,$7,$8,$9,
			$10,$11,$12,$13,
			$14,$15,$16,
			$17,$18
		)`,
		c.ID, tenantID, c.Code, c.Name, nullString(c.TaxID),
		string(c.Type), string(c.Status),
		nullString(c.Industry), nullString(c.Website),
		c.CreditLimit, c.PaymentTerms, c.Currency, addrJSON,
		nullString(c.AirtableID), nullString(c.ERPID),
		pq.Array(tagsOrEmpty(c.Tags)),
		c.CreatedAt, c.UpdatedAt,
	)
	return err
}

// ── GetByID ───────────────────────────────────────────────────────────────────

func (s *ClientStore) GetByID(tenantID, id string) (*client.Client, error) {
	row := s.db.QueryRow(`SELECT `+clientCols+` FROM clients
		WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanClient(row)
}

// ── List ──────────────────────────────────────────────────────────────────────

func (s *ClientStore) List(tenantID string, f handlers.ClientFilter) ([]*client.Client, int, error) {
	where := []string{"tenant_id = $1"}
	args := []interface{}{tenantID}
	n := 2

	if f.Type != "" {
		where = append(where, fmt.Sprintf("type = $%d", n))
		args = append(args, f.Type)
		n++
	}
	if f.Status != "" {
		where = append(where, fmt.Sprintf("status = $%d", n))
		args = append(args, f.Status)
		n++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("(name ILIKE $%d OR code ILIKE $%d OR tax_id ILIKE $%d)", n, n, n))
		args = append(args, "%"+f.Search+"%")
		n++
	}

	clause := "WHERE " + strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM clients "+clause, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	if f.Page < 1 {
		f.Page = 1
	}
	if f.PerPage < 1 || f.PerPage > 200 {
		f.PerPage = 50
	}
	offset := (f.Page - 1) * f.PerPage

	args = append(args, f.PerPage, offset)
	rows, err := s.db.Query(
		`SELECT `+clientCols+` FROM clients `+clause+
			fmt.Sprintf(` ORDER BY name ASC LIMIT $%d OFFSET $%d`, n, n+1),
		args...,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []*client.Client
	for rows.Next() {
		c, err := scanClient(rows)
		if err != nil {
			return nil, 0, err
		}
		out = append(out, c)
	}
	return out, total, rows.Err()
}

// ── Update ────────────────────────────────────────────────────────────────────

var allowedClientPatch = map[string]bool{
	"name": true, "status": true, "type": true, "industry": true,
	"website": true, "credit_limit": true, "payment_terms_days": true,
	"currency": true, "airtable_id": true, "erp_id": true, "tax_id": true,
}

func (s *ClientStore) Update(tenantID, id string, patch map[string]interface{}) (*client.Client, error) {
	sets := []string{}
	args := []interface{}{}
	n := 1

	for k, v := range patch {
		if !allowedClientPatch[k] {
			continue
		}
		sets = append(sets, fmt.Sprintf("%s = $%d", k, n))
		args = append(args, v)
		n++
	}
	if len(sets) == 0 {
		return s.GetByID(tenantID, id)
	}

	sets = append(sets, fmt.Sprintf("updated_at = $%d", n))
	args = append(args, time.Now())
	n++

	args = append(args, tenantID, id)
	_, err := s.db.Exec(
		`UPDATE clients SET `+strings.Join(sets, ", ")+
			fmt.Sprintf(` WHERE tenant_id = $%d AND id = $%d`, n, n+1),
		args...,
	)
	if err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

// ── Contacts ──────────────────────────────────────────────────────────────────

func (s *ClientStore) CreateContact(tenantID string, c *client.Contact) error {
	c.ID = newUUID()
	c.TenantID = tenantID
	c.CreatedAt = time.Now()

	// Only one primary per client
	if c.IsPrimary {
		s.db.Exec(`UPDATE client_contacts SET is_primary = false
			WHERE tenant_id = $1 AND client_id = $2`, tenantID, c.ClientID)
	}

	_, err := s.db.Exec(`
		INSERT INTO client_contacts (id, tenant_id, client_id, first_name, last_name, email, phone, position, is_primary, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		c.ID, tenantID, c.ClientID,
		c.FirstName, c.LastName,
		nullString(c.Email), nullString(c.Phone), nullString(c.Position),
		c.IsPrimary, c.CreatedAt,
	)
	return err
}

func (s *ClientStore) ListContacts(tenantID, clientID string) ([]*client.Contact, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, client_id::text, first_name, last_name,
		       COALESCE(email,''), COALESCE(phone,''), COALESCE(position,''),
		       is_primary, created_at
		FROM client_contacts
		WHERE tenant_id = $1 AND client_id = $2
		ORDER BY is_primary DESC, created_at`, tenantID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*client.Contact
	for rows.Next() {
		c := &client.Contact{}
		if err := rows.Scan(&c.ID, &c.TenantID, &c.ClientID,
			&c.FirstName, &c.LastName,
			&c.Email, &c.Phone, &c.Position,
			&c.IsPrimary, &c.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// ── Interactions ──────────────────────────────────────────────────────────────

func (s *ClientStore) CreateInteraction(tenantID string, i *client.Interaction) error {
	i.ID = newUUID()
	i.TenantID = tenantID
	i.CreatedAt = time.Now()
	if i.OccurredAt.IsZero() {
		i.OccurredAt = time.Now()
	}

	_, err := s.db.Exec(`
		INSERT INTO client_interactions (id, tenant_id, client_id, contact_id, user_id, type, subject, notes, occurred_at, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		i.ID, tenantID, i.ClientID, nullString(i.ContactID), i.UserID,
		string(i.Type), i.Subject, nullString(i.Notes),
		i.OccurredAt, i.CreatedAt,
	)
	return err
}

func (s *ClientStore) ListInteractions(tenantID, clientID string) ([]*client.Interaction, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, client_id::text, COALESCE(contact_id::text,''),
		       user_id::text, type, subject, COALESCE(notes,''),
		       occurred_at, created_at
		FROM client_interactions
		WHERE tenant_id = $1 AND client_id = $2
		ORDER BY occurred_at DESC
		LIMIT 500`, tenantID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*client.Interaction
	for rows.Next() {
		i := &client.Interaction{}
		var iType string
		if err := rows.Scan(&i.ID, &i.TenantID, &i.ClientID, &i.ContactID,
			&i.UserID, &iType, &i.Subject, &i.Notes,
			&i.OccurredAt, &i.CreatedAt); err != nil {
			return nil, err
		}
		i.Type = client.InteractionType(iType)
		out = append(out, i)
	}
	return out, rows.Err()
}

// ── SQL helpers ───────────────────────────────────────────────────────────────

const clientCols = `
	id, tenant_id, code, name, COALESCE(tax_id,''),
	type, status,
	COALESCE(industry,''), COALESCE(website,''),
	credit_limit, payment_terms_days, currency, address,
	COALESCE(airtable_id,''), COALESCE(erp_id,''), tags,
	created_at, updated_at`

type clientScanner interface {
	Scan(dest ...interface{}) error
}

func scanClient(row clientScanner) (*client.Client, error) {
	c := &client.Client{}
	var typ, status string
	var addrJSON []byte
	var tags pq.StringArray

	err := row.Scan(
		&c.ID, &c.TenantID, &c.Code, &c.Name, &c.TaxID,
		&typ, &status,
		&c.Industry, &c.Website,
		&c.CreditLimit, &c.PaymentTerms, &c.Currency, &addrJSON,
		&c.AirtableID, &c.ERPID, &tags,
		&c.CreatedAt, &c.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("client not found")
	}
	if err != nil {
		return nil, err
	}

	c.Type = client.ClientType(typ)
	c.Status = client.ClientStatus(status)
	c.Tags = []string(tags)
	scanJSON(addrJSON, &c.Address)
	return c, nil
}

func tagsOrEmpty(tags []string) []string {
	if tags == nil {
		return []string{}
	}
	return tags
}

var _ handlers.ClientStore = (*ClientStore)(nil)
