package sqlite

import (
	"database/sql"

	"github.com/chacontainer/backend/internal/api/handlers"
)

// WebhookProcessor persists inbound webhook events for audit. Outbound
// dispatch to Make.com/ERP/Airtable (internal/integrations/*) is documented
// as backlog in PRD.md — those clients exist but aren't wired to writes yet.
type WebhookProcessor struct {
	db *sql.DB
}

func NewWebhookProcessor(db *sql.DB) *WebhookProcessor {
	return &WebhookProcessor{db: db}
}

func (p *WebhookProcessor) record(event handlers.WebhookEvent) error {
	_, err := p.db.Exec(
		`INSERT INTO webhook_log (id, source, event, payload, received_at) VALUES (?,?,?,?,?)`,
		newID(), event.Source, event.Event, toJSON(event.Payload), nowStr(),
	)
	return err
}

func (p *WebhookProcessor) ProcessMakeEvent(event handlers.WebhookEvent) error     { return p.record(event) }
func (p *WebhookProcessor) ProcessERPEvent(event handlers.WebhookEvent) error      { return p.record(event) }
func (p *WebhookProcessor) ProcessAirtableEvent(event handlers.WebhookEvent) error { return p.record(event) }
