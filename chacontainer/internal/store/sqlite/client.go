package sqlite

import (
	"database/sql"
	"strings"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/client"
)

type ClientStore struct {
	db *sql.DB
}

func NewClientStore(db *sql.DB) *ClientStore {
	return &ClientStore{db: db}
}

func (s *ClientStore) Create(tenantID string, c *client.Client) error {
	c.ID = newID()
	c.TenantID = tenantID
	now := parseTime(nowStr())
	c.CreatedAt = now
	c.UpdatedAt = now
	if c.Status == "" {
		c.Status = client.ClientStatusActive
	}
	if c.Currency == "" {
		c.Currency = "MXN"
	}
	_, err := s.db.Exec(
		`INSERT INTO clients (id, tenant_id, code, name, tax_id, type, status, industry, website, credit_limit,
			payment_terms_days, currency, address, airtable_id, erp_id, tags, created_at, updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		c.ID, c.TenantID, c.Code, c.Name, nullStr(c.TaxID), string(c.Type), string(c.Status), nullStr(c.Industry),
		nullStr(c.Website), c.CreditLimit, c.PaymentTerms, c.Currency, toJSON(c.Address), nullStr(c.AirtableID),
		nullStr(c.ERPID), toJSON(c.Tags), fmtTime(c.CreatedAt), fmtTime(c.UpdatedAt),
	)
	return err
}

const clientSelectCols = `id, tenant_id, code, name, tax_id, type, status, industry, website, credit_limit,
	payment_terms_days, currency, address, airtable_id, erp_id, tags, created_at, updated_at`

func scanClient(row interface{ Scan(dest ...interface{}) error }) (*client.Client, error) {
	var c client.Client
	var taxID, industry, website, airtableID, erpID sql.NullString
	var typeStr, statusStr, address, tags, createdAt, updatedAt string

	err := row.Scan(&c.ID, &c.TenantID, &c.Code, &c.Name, &taxID, &typeStr, &statusStr, &industry, &website,
		&c.CreditLimit, &c.PaymentTerms, &c.Currency, &address, &airtableID, &erpID, &tags, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	c.TaxID = taxID.String
	c.Type = client.ClientType(typeStr)
	c.Status = client.ClientStatus(statusStr)
	c.Industry = industry.String
	c.Website = website.String
	c.AirtableID = airtableID.String
	c.ERPID = erpID.String
	fromJSON(address, &c.Address)
	fromJSON(tags, &c.Tags)
	c.CreatedAt = parseTime(createdAt)
	c.UpdatedAt = parseTime(updatedAt)
	return &c, nil
}

func (s *ClientStore) GetByID(tenantID, id string) (*client.Client, error) {
	row := s.db.QueryRow(`SELECT `+clientSelectCols+` FROM clients WHERE tenant_id = ? AND id = ?`, tenantID, id)
	return scanClient(row)
}

func (s *ClientStore) List(tenantID string, filter handlers.ClientFilter) ([]*client.Client, int, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{tenantID}
	if filter.Type != "" {
		where = append(where, "type = ?")
		args = append(args, filter.Type)
	}
	if filter.Status != "" {
		where = append(where, "status = ?")
		args = append(args, filter.Status)
	}
	if filter.Search != "" {
		where = append(where, "(name LIKE ? OR code LIKE ?)")
		like := "%" + filter.Search + "%"
		args = append(args, like, like)
	}
	whereClause := strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM clients WHERE `+whereClause, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	page, perPage := filter.Page, filter.PerPage
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 50
	}
	offset := (page - 1) * perPage
	queryArgs := append(append([]interface{}{}, args...), perPage, offset)

	rows, err := s.db.Query(`SELECT `+clientSelectCols+` FROM clients WHERE `+whereClause+` ORDER BY created_at DESC LIMIT ? OFFSET ?`, queryArgs...)
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

func (s *ClientStore) Update(tenantID, id string, patch map[string]interface{}) (*client.Client, error) {
	allowed := map[string]bool{
		"name": true, "status": true, "industry": true, "website": true,
		"credit_limit": true, "payment_terms_days": true,
	}
	var sets []string
	var args []interface{}
	for k, v := range patch {
		if !allowed[k] {
			continue
		}
		sets = append(sets, k+" = ?")
		args = append(args, v)
	}
	if len(sets) == 0 {
		return s.GetByID(tenantID, id)
	}
	sets = append(sets, "updated_at = ?")
	args = append(args, nowStr(), tenantID, id)
	if _, err := s.db.Exec(`UPDATE clients SET `+strings.Join(sets, ", ")+` WHERE tenant_id = ? AND id = ?`, args...); err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

func (s *ClientStore) ListContacts(tenantID, clientID string) ([]*client.Contact, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, client_id, first_name, last_name, email, phone, position, is_primary, created_at
		 FROM client_contacts WHERE tenant_id = ? AND client_id = ? ORDER BY created_at`, tenantID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*client.Contact
	for rows.Next() {
		var c client.Contact
		var email, phone, position sql.NullString
		var isPrimary int
		var createdAt string
		if err := rows.Scan(&c.ID, &c.TenantID, &c.ClientID, &c.FirstName, &c.LastName, &email, &phone, &position, &isPrimary, &createdAt); err != nil {
			return nil, err
		}
		c.Email = email.String
		c.Phone = phone.String
		c.Position = position.String
		c.IsPrimary = isPrimary == 1
		c.CreatedAt = parseTime(createdAt)
		out = append(out, &c)
	}
	return out, rows.Err()
}

func (s *ClientStore) CreateContact(tenantID string, c *client.Contact) error {
	c.ID = newID()
	c.TenantID = tenantID
	c.CreatedAt = parseTime(nowStr())
	_, err := s.db.Exec(
		`INSERT INTO client_contacts (id, tenant_id, client_id, first_name, last_name, email, phone, position, is_primary, created_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?)`,
		c.ID, c.TenantID, c.ClientID, c.FirstName, c.LastName, nullStr(c.Email), nullStr(c.Phone), nullStr(c.Position),
		boolToInt(c.IsPrimary), fmtTime(c.CreatedAt),
	)
	return err
}

func (s *ClientStore) CreateInteraction(tenantID string, i *client.Interaction) error {
	i.ID = newID()
	i.TenantID = tenantID
	now := parseTime(nowStr())
	if i.OccurredAt.IsZero() {
		i.OccurredAt = now
	}
	i.CreatedAt = now
	_, err := s.db.Exec(
		`INSERT INTO client_interactions (id, tenant_id, client_id, contact_id, user_id, type, subject, notes, occurred_at, created_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?)`,
		i.ID, i.TenantID, i.ClientID, nullStr(i.ContactID), i.UserID, string(i.Type), i.Subject, nullStr(i.Notes),
		fmtTime(i.OccurredAt), fmtTime(i.CreatedAt),
	)
	return err
}

func (s *ClientStore) ListInteractions(tenantID, clientID string) ([]*client.Interaction, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, client_id, contact_id, user_id, type, subject, notes, occurred_at, created_at
		 FROM client_interactions WHERE tenant_id = ? AND client_id = ? ORDER BY occurred_at DESC`, tenantID, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*client.Interaction
	for rows.Next() {
		var i client.Interaction
		var contactID, notes sql.NullString
		var typeStr, occurredAt, createdAt string
		if err := rows.Scan(&i.ID, &i.TenantID, &i.ClientID, &contactID, &i.UserID, &typeStr, &i.Subject, &notes, &occurredAt, &createdAt); err != nil {
			return nil, err
		}
		i.ContactID = contactID.String
		i.Notes = notes.String
		i.Type = client.InteractionType(typeStr)
		i.OccurredAt = parseTime(occurredAt)
		i.CreatedAt = parseTime(createdAt)
		out = append(out, &i)
	}
	return out, rows.Err()
}
