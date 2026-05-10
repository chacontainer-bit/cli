package airtable

import (
	"fmt"
	"log"
	"time"

	"github.com/chacontainer/backend/internal/domain/asset"
	"github.com/chacontainer/backend/internal/domain/client"
	"github.com/chacontainer/backend/internal/domain/shipment"
)

// TableNames maps domain entities to Airtable table names.
var TableNames = struct {
	Assets    string
	Clients   string
	Shipments string
}{
	Assets:    "Assets",
	Clients:   "Clients",
	Shipments: "Shipments",
}

// SyncService bidirectionally syncs CHACONTAINER ↔ Airtable.
type SyncService struct {
	client *Client
}

func NewSyncService(apiKey, baseID string) *SyncService {
	return &SyncService{client: NewClient(apiKey, baseID)}
}

// PushAsset upserts one asset to Airtable.
func (s *SyncService) PushAsset(a *asset.Asset) (string, error) {
	fields := map[string]interface{}{
		"Code":          a.Code,
		"QR Code":       a.QRCode,
		"Type":          string(a.Type),
		"Status":        string(a.Status),
		"Description":   a.Description,
		"Plant ID":      a.PlantID,
		"Weight (kg)":   a.WeightKg,
		"Last Scan":     formatTime(a.LastScanAt),
		"CHACONTAINER ID": a.ID,
	}

	records := []*Record{{Fields: fields}}
	if a.AirtableID != "" {
		records[0].ID = a.AirtableID
	}

	resp, err := s.client.Upsert(TableNames.Assets, records)
	if err != nil {
		return "", fmt.Errorf("push asset: %w", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("no record returned")
	}
	return resp[0].ID, nil
}

// PushClient upserts a client record to Airtable.
func (s *SyncService) PushClient(c *client.Client) (string, error) {
	fields := map[string]interface{}{
		"Code":            c.Code,
		"Name":            c.Name,
		"Type":            string(c.Type),
		"Status":          string(c.Status),
		"Tax ID":          c.TaxID,
		"Industry":        c.Industry,
		"Credit Limit":    c.CreditLimit,
		"Currency":        c.Currency,
		"CHACONTAINER ID": c.ID,
	}

	records := []*Record{{Fields: fields}}
	if c.AirtableID != "" {
		records[0].ID = c.AirtableID
	}

	resp, err := s.client.Upsert(TableNames.Clients, records)
	if err != nil {
		return "", fmt.Errorf("push client: %w", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("no record returned")
	}
	return resp[0].ID, nil
}

// PushShipment upserts a shipment record to Airtable.
func (s *SyncService) PushShipment(sh *shipment.Shipment) (string, error) {
	fields := map[string]interface{}{
		"Reference":         sh.Reference,
		"Status":            string(sh.Status),
		"Mode":              string(sh.Mode),
		"Carrier":           sh.CarrierName,
		"Carrier Ref":       sh.CarrierRef,
		"Scheduled Pickup":  sh.ScheduledPickup.Format(time.RFC3339),
		"Scheduled Delivery": sh.ScheduledDelivery.Format(time.RFC3339),
		"Total Weight (kg)": sh.TotalWeightKg,
		"Total Items":       sh.TotalItems,
		"CHACONTAINER ID":   sh.ID,
	}

	records := []*Record{{Fields: fields}}
	resp, err := s.client.Upsert(TableNames.Shipments, records)
	if err != nil {
		return "", fmt.Errorf("push shipment: %w", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("no record returned")
	}
	return resp[0].ID, nil
}

// PullClients fetches all client records from Airtable and returns them
// as a map keyed by CHACONTAINER ID for merge logic.
func (s *SyncService) PullClients() (map[string]*Record, error) {
	records, err := s.client.List(TableNames.Clients, "")
	if err != nil {
		return nil, err
	}
	out := make(map[string]*Record, len(records))
	for _, r := range records {
		if id, ok := r.Fields["CHACONTAINER ID"].(string); ok && id != "" {
			out[id] = r
		}
	}
	return out, nil
}

// BulkPushAssets pushes up to 10 assets per API call.
func (s *SyncService) BulkPushAssets(assets []*asset.Asset) error {
	batch := make([]*Record, 0, 10)
	flush := func() error {
		if len(batch) == 0 {
			return nil
		}
		if _, err := s.client.Upsert(TableNames.Assets, batch); err != nil {
			return err
		}
		batch = batch[:0]
		return nil
	}

	for _, a := range assets {
		batch = append(batch, &Record{
			ID: a.AirtableID,
			Fields: map[string]interface{}{
				"Code":            a.Code,
				"Status":          string(a.Status),
				"CHACONTAINER ID": a.ID,
			},
		})
		if len(batch) == 10 {
			if err := flush(); err != nil {
				log.Printf("bulk push error: %v", err)
			}
			time.Sleep(250 * time.Millisecond) // respect rate limit (5 req/s)
		}
	}
	return flush()
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}

// AirtableID is a helper field injected into domain structs during sync.
// The postgres column is airtable_id VARCHAR(50).
type AirtableID = string
