package make_com

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Sender dispatches events to Make.com scenario webhooks.
type Sender struct {
	http *http.Client
}

func NewSender() *Sender {
	return &Sender{http: &http.Client{Timeout: 10 * time.Second}}
}

// ShipmentEvent is sent to Make when a shipment changes status.
type ShipmentEvent struct {
	TenantID    string    `json:"tenant_id"`
	ShipmentRef string    `json:"shipment_ref"`
	ClientID    string    `json:"client_id"`
	FromStatus  string    `json:"from_status"`
	ToStatus    string    `json:"to_status"`
	CarrierName string    `json:"carrier_name"`
	OccurredAt  time.Time `json:"occurred_at"`
}

// AssetAlertEvent is sent when an asset enters an alert state.
type AssetAlertEvent struct {
	TenantID   string    `json:"tenant_id"`
	AssetID    string    `json:"asset_id"`
	AssetCode  string    `json:"asset_code"`
	AlertType  string    `json:"alert_type"` // overdue_maintenance, lost, in_wrong_zone
	PlantID    string    `json:"plant_id"`
	Message    string    `json:"message"`
	OccurredAt time.Time `json:"occurred_at"`
}

// ScanEvent triggers downstream automation on QR scans.
type ScanEvent struct {
	TenantID  string    `json:"tenant_id"`
	AssetCode string    `json:"asset_code"`
	Action    string    `json:"action"`
	UserID    string    `json:"user_id"`
	PlantID   string    `json:"plant_id"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	ScannedAt time.Time `json:"scanned_at"`
}

// Send fires a JSON payload to a Make.com webhook URL.
func (s *Sender) Send(webhookURL string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.http.Do(req)
	if err != nil {
		return fmt.Errorf("make http: %w", err)
	}
	defer resp.Body.Close()
	io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return fmt.Errorf("make returned %d", resp.StatusCode)
	}
	return nil
}

// MakeScenarios holds webhook URLs keyed by event type.
// These are configured per-tenant in the tenants.settings JSONB column.
type MakeScenarios struct {
	ShipmentStatusChanged string `json:"shipment_status_changed"`
	AssetAlert            string `json:"asset_alert"`
	QRScan                string `json:"qr_scan"`
	NewClient             string `json:"new_client"`
	LateShipment          string `json:"late_shipment"`
}
