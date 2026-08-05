package postgres

import (
	"database/sql"
	"fmt"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/domain/client"
)

// ClientStore persists CRM clients, contacts and interactions in PostgreSQL.
type ClientStore struct {
	db *sql.DB
}

func NewClientStore(db *sql.DB) *ClientStore {
	return &ClientStore{db: db}
}

func (s *ClientStore) Create(tenantID string, c *client.Client) error {
	address, err := marshalJSON(c.Address)
	if err != nil {
		return err
	}
	tags, err := marshalJSON(c.Tags)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(`
		INSERT INTO clients (tenant_id, code, name, tax_id, type, status, industry, website,
			credit_limit, payment_terms_days, currency, address, airtable_id, erp_id, tags)
		VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'shipper'), COALESCE(NULLIF($6, ''), 'active'),
			$7, $8, $9, COALESCE(NULLIF($10, 0), 30), COALESCE(NULLIF($11, ''), 'MXN'), $12::jsonb, $13, $14, $15::jsonb)
		RETURNING id, created_at, updated_at
	`, tenantID, c.Code, c.Name, nullStr(c.TaxID), string(c.Type), string(c.Status), nullStr(c.Industry),
		nullStr(c.Website), c.CreditLimit, c.PaymentTerms, c.Currency, address, nullStr(c.AirtableID), nullStr(c.ERPID), tags)

	if err := row.Scan(&c.ID, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return fmt.Errorf("insert client: %w", err)
	}
	c.TenantID = tenantID
	return nil
}

func (s *ClientStore) GetByID(tenantID, id string) (*client.Client, error) {
	row := s.db.QueryRow(clientSelect+` WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanClient(row)
}

const clientSelect = `
	SELECT id, tenant_id, code, name, tax_id, type, status, industry, website,
		credit_limit, payment_terms_days, currency, address, airtable_id, erp_id, tags, created_at, updated_at
	FROM clients`

func (s *ClientStore) List(tenantID string, filter handlers.ClientFilter) ([]*client.Client, int, error) {
	where := "WHERE tenant_id = $1"
	args := []interface{}{tenantID}

	if filter.Type != "" {
		args = append(args, filter.Type)
		where += fmt.Sprintf(" AND type = $%d", len(args))
	}
	if filter.Status != "" {
		args = append(args, filter.Status)
		where += fmt.Sprintf(" AND status = $%d", len(args))
	}
	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		where += fmt.Sprintf(" AND (name ILIKE $%d OR code ILIKE $%d)", len(args), len(args))
	}

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM clients `+where, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count clients: %w", err)
	}

	page, perPage := pagination(filter.Page, filter.PerPage)
	args = append(args, perPage, (page-1)*perPage)
	query := clientSelect + " " + where + fmt.Sprintf(" ORDER BY name LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list clients: %w", err)
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

var clientUpdatable = map[string]columnKind{
	"code":               kindText,
	"name":               kindText,
	"tax_id":             kindText,
	"type":               kindText,
	"status":             kindText,
	"industry":           kindText,
	"website":            kindText,
	"credit_limit":       kindNumber,
	"payment_terms_days": kindNumber,
	"currency":           kindText,
	"address":            kindJSON,
	"airtable_id":        kindText,
	"erp_id":             kindText,
	"tags":               kindJSON,
}

func (s *ClientStore) Update(tenantID, id string, patch map[string]interface{}) (*client.Client, error) {
	setClause, args, err := buildUpdate(clientUpdatable, patch)
	if err != nil {
		return nil, err
	}
	args = append(args, tenantID, id)
	query := fmt.Sprintf(`UPDATE clients SET %s, updated_at = NOW() WHERE tenant_id = $%d AND id = $%d RETURNING id`,
		setClause, len(args)-1, len(args))

	var updatedID string
	if err := s.db.QueryRow(query, args...).Scan(&updatedID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client not found: %s", id)
		}
		return nil, fmt.Errorf("update client: %w", err)
	}
	return s.GetByID(tenantID, updatedID)
}

func (s *ClientStore) ListContacts(tenantID, clientID string) ([]*client.Contact, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, client_id, first_name, last_name, email, phone, position, is_primary, created_at
		FROM client_contacts WHERE tenant_id = $1 AND client_id = $2 ORDER BY is_primary DESC, first_name
	`, tenantID, clientID)
	if err != nil {
		return nil, fmt.Errorf("list contacts: %w", err)
	}
	defer rows.Close()

	var out []*client.Contact
	for rows.Next() {
		var c client.Contact
		var email, phone, position sql.NullString
		if err := rows.Scan(&c.ID, &c.TenantID, &c.ClientID, &c.FirstName, &c.LastName, &email, &phone, &position, &c.IsPrimary, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan contact: %w", err)
		}
		c.Email, c.Phone, c.Position = strOrEmpty(email), strOrEmpty(phone), strOrEmpty(position)
		out = append(out, &c)
	}
	return out, rows.Err()
}

func (s *ClientStore) CreateContact(tenantID string, c *client.Contact) error {
	row := s.db.QueryRow(`
		INSERT INTO client_contacts (tenant_id, client_id, first_name, last_name, email, phone, position, is_primary)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`, tenantID, c.ClientID, c.FirstName, c.LastName, nullStr(c.Email), nullStr(c.Phone), nullStr(c.Position), c.IsPrimary)

	if err := row.Scan(&c.ID, &c.CreatedAt); err != nil {
		return fmt.Errorf("insert contact: %w", err)
	}
	c.TenantID = tenantID
	return nil
}

func (s *ClientStore) CreateInteraction(tenantID string, i *client.Interaction) error {
	row := s.db.QueryRow(`
		INSERT INTO client_interactions (tenant_id, client_id, contact_id, user_id, type, subject, notes, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, COALESCE(NULLIF($8, '')::timestamptz, NOW()))
		RETURNING id, created_at, occurred_at
	`, tenantID, i.ClientID, nullStr(i.ContactID), i.UserID, string(i.Type), i.Subject, nullStr(i.Notes), nullTimeStr(i.OccurredAt))

	if err := row.Scan(&i.ID, &i.CreatedAt, &i.OccurredAt); err != nil {
		return fmt.Errorf("insert interaction: %w", err)
	}
	i.TenantID = tenantID
	return nil
}

func (s *ClientStore) ListInteractions(tenantID, clientID string) ([]*client.Interaction, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, client_id, contact_id, user_id, type, subject, notes, occurred_at, created_at
		FROM client_interactions WHERE tenant_id = $1 AND client_id = $2 ORDER BY occurred_at DESC
	`, tenantID, clientID)
	if err != nil {
		return nil, fmt.Errorf("list interactions: %w", err)
	}
	defer rows.Close()

	var out []*client.Interaction
	for rows.Next() {
		var i client.Interaction
		var contactID, notes sql.NullString
		var typ string
		if err := rows.Scan(&i.ID, &i.TenantID, &i.ClientID, &contactID, &i.UserID, &typ, &i.Subject, &notes, &i.OccurredAt, &i.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan interaction: %w", err)
		}
		i.ContactID, i.Notes, i.Type = strOrEmpty(contactID), strOrEmpty(notes), client.InteractionType(typ)
		out = append(out, &i)
	}
	return out, rows.Err()
}

func scanClient(row rowScanner) (*client.Client, error) {
	var c client.Client
	var typ, status string
	var taxID, industry, website, airtableID, erpID sql.NullString
	var address, tags []byte
	if err := row.Scan(&c.ID, &c.TenantID, &c.Code, &c.Name, &taxID, &typ, &status, &industry, &website,
		&c.CreditLimit, &c.PaymentTerms, &c.Currency, &address, &airtableID, &erpID, &tags, &c.CreatedAt, &c.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client not found")
		}
		return nil, fmt.Errorf("scan client: %w", err)
	}
	c.Type, c.Status = client.ClientType(typ), client.ClientStatus(status)
	c.TaxID, c.Industry, c.Website = strOrEmpty(taxID), strOrEmpty(industry), strOrEmpty(website)
	c.AirtableID, c.ERPID = strOrEmpty(airtableID), strOrEmpty(erpID)
	if err := unmarshalJSON(address, &c.Address); err != nil {
		return nil, fmt.Errorf("decode client address: %w", err)
	}
	if err := unmarshalJSON(tags, &c.Tags); err != nil {
		return nil, fmt.Errorf("decode client tags: %w", err)
	}
	return &c, nil
}
