package api

// The webhook processor stays a no-op stub even in local/private mode: this
// server accepts inbound Make.com/ERP/Airtable webhooks (so those platforms
// can still push events in if you wire them up yourself), but never
// initiates outbound calls to any of them. Wire a real processor here only
// if you explicitly want that integration back.

import "github.com/cli/cli/v2/chacontainer/internal/api/handlers"

type stubWebhookProcessor struct{}

func (s *stubWebhookProcessor) ProcessMakeEvent(_ handlers.WebhookEvent) error     { return nil }
func (s *stubWebhookProcessor) ProcessERPEvent(_ handlers.WebhookEvent) error      { return nil }
func (s *stubWebhookProcessor) ProcessAirtableEvent(_ handlers.WebhookEvent) error { return nil }
