package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/asset"
	"github.com/chacontainer/backend/internal/domain/shipment"
)

// WebhookProcessor implements handlers.WebhookProcessor.
// It routes incoming events from Make.com, ERP, and Airtable to DB state changes.
type WebhookProcessor struct{ db *sql.DB }

func NewWebhookProcessor(db *sql.DB) *WebhookProcessor { return &WebhookProcessor{db: db} }

// ── ProcessMakeEvent ──────────────────────────────────────────────────────────

func (p *WebhookProcessor) ProcessMakeEvent(event handlers.WebhookEvent) error {
	switch event.Event {
	case "shipment_status_update":
		return p.handleMakeShipmentUpdate(event.Payload)
	case "asset_maintenance_complete":
		return p.handleMakeMaintenanceComplete(event.Payload)
	default:
		log.Printf("make: unhandled event %q", event.Event)
		return nil
	}
}

func (p *WebhookProcessor) handleMakeShipmentUpdate(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	ref, _ := payload["shipment_ref"].(string)
	newStatus, _ := payload["status"].(string)
	description, _ := payload["description"].(string)

	if tenantID == "" || ref == "" || newStatus == "" {
		return fmt.Errorf("make shipment_update: missing required fields")
	}

	var id string
	err := p.db.QueryRow(
		`SELECT id FROM shipments WHERE tenant_id = $1 AND reference = $2`,
		tenantID, ref,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("shipment not found ref=%s: %w", ref, err)
	}

	store := &ShipmentStore{db: p.db}
	return store.UpdateStatus(tenantID, id,
		shipment.ShipmentStatus(newStatus),
		"make-automation", description)
}

func (p *WebhookProcessor) handleMakeMaintenanceComplete(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	assetID, _ := payload["asset_id"].(string)
	notes, _ := payload["notes"].(string)

	if tenantID == "" || assetID == "" {
		return fmt.Errorf("make maintenance_complete: missing fields")
	}

	_, err := p.db.Exec(`
		UPDATE maintenance_records
		SET completed_at = NOW(), notes = COALESCE($1, notes)
		WHERE tenant_id = $2 AND asset_id = $3
		  AND completed_at IS NULL
		  AND scheduled_at <= NOW()`,
		nullString(notes), tenantID, assetID,
	)
	if err != nil {
		return err
	}

	// Move asset back to available
	_, err = p.db.Exec(`
		UPDATE assets SET status = 'available', updated_at = NOW()
		WHERE tenant_id = $1 AND id = $2 AND status = 'maintenance'`,
		tenantID, assetID,
	)
	return err
}

// ── ProcessERPEvent ───────────────────────────────────────────────────────────

func (p *WebhookProcessor) ProcessERPEvent(event handlers.WebhookEvent) error {
	switch event.Event {
	case "goods_receipt":
		return p.handleERPGoodsReceipt(event.Payload)
	case "purchase_order_confirmed":
		return p.handleERPPOConfirmed(event.Payload)
	default:
		log.Printf("erp: unhandled event %q", event.Event)
		return nil
	}
}

func (p *WebhookProcessor) handleERPGoodsReceipt(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	erpID, _ := payload["erp_id"].(string)
	if tenantID == "" || erpID == "" {
		return fmt.Errorf("erp goods_receipt: missing fields")
	}

	// Mark shipment as delivered if ERP confirms goods receipt
	_, err := p.db.Exec(`
		UPDATE shipments
		SET status = 'delivered', actual_delivery = NOW(), erp_id = $1, updated_at = NOW()
		WHERE tenant_id = $2 AND erp_id = $1 AND status = 'in_transit'`,
		erpID, tenantID,
	)
	return err
}

func (p *WebhookProcessor) handleERPPOConfirmed(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	erpRef, _ := payload["reference"].(string)
	if tenantID == "" || erpRef == "" {
		return fmt.Errorf("erp po_confirmed: missing fields")
	}

	// Confirm the draft shipment that matches the ERP reference
	_, err := p.db.Exec(`
		UPDATE shipments
		SET status = 'confirmed', updated_at = NOW()
		WHERE tenant_id = $1 AND reference = $2 AND status = 'draft'`,
		tenantID, erpRef,
	)
	return err
}

// ── ProcessAirtableEvent ──────────────────────────────────────────────────────

func (p *WebhookProcessor) ProcessAirtableEvent(event handlers.WebhookEvent) error {
	switch event.Event {
	case "client_updated":
		return p.handleAirtableClientUpdate(event.Payload)
	case "asset_status_override":
		return p.handleAirtableAssetOverride(event.Payload)
	default:
		log.Printf("airtable: unhandled event %q", event.Event)
		return nil
	}
}

func (p *WebhookProcessor) handleAirtableClientUpdate(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	ccID, _ := payload["chacontainer_id"].(string)
	if tenantID == "" || ccID == "" {
		return fmt.Errorf("airtable client_updated: missing fields")
	}

	fields := []string{}
	args := []interface{}{}
	n := 1

	if v, ok := payload["credit_limit"]; ok {
		fields = append(fields, fmt.Sprintf("credit_limit = $%d", n))
		args = append(args, v)
		n++
	}
	if v, ok := payload["status"]; ok {
		fields = append(fields, fmt.Sprintf("status = $%d", n))
		args = append(args, v)
		n++
	}
	if len(fields) == 0 {
		return nil
	}

	fields = append(fields, fmt.Sprintf("updated_at = $%d", n))
	args = append(args, time.Now())
	n++

	args = append(args, tenantID, ccID)
	query := "UPDATE clients SET " + joinStrings(fields, ", ") +
		fmt.Sprintf(" WHERE tenant_id = $%d AND id = $%d", n, n+1)
	_, err := p.db.Exec(query, args...)
	return err
}

func (p *WebhookProcessor) handleAirtableAssetOverride(payload map[string]interface{}) error {
	tenantID, _ := payload["tenant_id"].(string)
	ccID, _ := payload["chacontainer_id"].(string)
	newStatus, _ := payload["status"].(string)
	if tenantID == "" || ccID == "" || newStatus == "" {
		return fmt.Errorf("airtable asset_status_override: missing fields")
	}

	aStore := &AssetStore{db: p.db}

	// Fetch current status for the event record
	a, err := aStore.GetByID(tenantID, ccID)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(`
		UPDATE assets SET status = $1, updated_at = NOW()
		WHERE tenant_id = $2 AND id = $3`,
		newStatus, tenantID, ccID,
	)
	if err != nil {
		return err
	}

	return aStore.RecordEvent(&asset.AssetEvent{
		TenantID:  tenantID,
		AssetID:   ccID,
		EventType: asset.EventTypeInspection,
		FromStatus: a.Status,
		ToStatus:  asset.AssetStatus(newStatus),
		UserID:    "airtable-sync",
		Notes:     "Status override from Airtable",
		OccurredAt: time.Now(),
		CreatedAt:  time.Now(),
	})
}

// ── helpers ───────────────────────────────────────────────────────────────────

func joinStrings(ss []string, sep string) string {
	out := ""
	for i, s := range ss {
		if i > 0 {
			out += sep
		}
		out += s
	}
	return out
}

var _ handlers.WebhookProcessor = (*WebhookProcessor)(nil)
