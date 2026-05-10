package asset

import "time"

type AssetType string

const (
	AssetTypeContainer  AssetType = "container"
	AssetTypePallet     AssetType = "pallet"
	AssetTypeForklift   AssetType = "forklift"
	AssetTypeRack       AssetType = "rack"
	AssetTypeIBC        AssetType = "ibc"
	AssetTypeDrum       AssetType = "drum"
	AssetTypeTrailer    AssetType = "trailer"
)

type AssetStatus string

const (
	AssetStatusAvailable  AssetStatus = "available"
	AssetStatusInUse      AssetStatus = "in_use"
	AssetStatusTransit    AssetStatus = "in_transit"
	AssetStatusMaintenance AssetStatus = "maintenance"
	AssetStatusRetired    AssetStatus = "retired"
	AssetStatusLost       AssetStatus = "lost"
)

// Asset is any trackable physical unit within the system.
type Asset struct {
	ID           string      `json:"id" db:"id"`
	TenantID     string      `json:"tenant_id" db:"tenant_id"`
	PlantID      string      `json:"plant_id" db:"plant_id"`
	ZoneID       string      `json:"zone_id,omitempty" db:"zone_id"`
	Name         string      `json:"name" db:"name"`
	Code         string      `json:"code" db:"code"`         // internal code
	AirtableID   string      `json:"airtable_id,omitempty" db:"airtable_id"`
	QRCode       string      `json:"qr_code" db:"qr_code"`   // QR payload
	Barcode      string      `json:"barcode,omitempty" db:"barcode"`
	Type         AssetType   `json:"type" db:"type"`
	Status       AssetStatus `json:"status" db:"status"`
	Description  string      `json:"description,omitempty" db:"description"`
	Manufacturer string      `json:"manufacturer,omitempty" db:"manufacturer"`
	Model        string      `json:"model,omitempty" db:"model"`
	SerialNumber string      `json:"serial_number,omitempty" db:"serial_number"`
	WeightKg     float64     `json:"weight_kg" db:"weight_kg"`
	Dimensions   Dimensions  `json:"dimensions" db:"dimensions"`
	PurchaseDate *time.Time  `json:"purchase_date,omitempty" db:"purchase_date"`
	PurchaseCost float64     `json:"purchase_cost,omitempty" db:"purchase_cost"`
	ClientID     string      `json:"client_id,omitempty" db:"client_id"`
	Metadata     map[string]interface{} `json:"metadata,omitempty" db:"metadata"`
	LastScanAt   *time.Time  `json:"last_scan_at,omitempty" db:"last_scan_at"`
	LastScanBy   string      `json:"last_scan_by,omitempty" db:"last_scan_by"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" db:"updated_at"`
}

type Dimensions struct {
	LengthCm float64 `json:"length_cm"`
	WidthCm  float64 `json:"width_cm"`
	HeightCm float64 `json:"height_cm"`
}

func (d Dimensions) VolumeCm3() float64 {
	return d.LengthCm * d.WidthCm * d.HeightCm
}

// AssetEvent records every state transition or scan event.
type AssetEvent struct {
	ID          string      `json:"id" db:"id"`
	TenantID    string      `json:"tenant_id" db:"tenant_id"`
	AssetID     string      `json:"asset_id" db:"asset_id"`
	EventType   EventType   `json:"event_type" db:"event_type"`
	FromStatus  AssetStatus `json:"from_status" db:"from_status"`
	ToStatus    AssetStatus `json:"to_status" db:"to_status"`
	FromPlantID string      `json:"from_plant_id,omitempty" db:"from_plant_id"`
	ToPlantID   string      `json:"to_plant_id,omitempty" db:"to_plant_id"`
	UserID      string      `json:"user_id" db:"user_id"`
	ShipmentID  string      `json:"shipment_id,omitempty" db:"shipment_id"`
	Notes       string      `json:"notes,omitempty" db:"notes"`
	Geo         *GeoPoint   `json:"geo,omitempty" db:"geo"`
	OccurredAt  time.Time   `json:"occurred_at" db:"occurred_at"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type EventType string

const (
	EventTypeScan        EventType = "scan"
	EventTypeCheckIn     EventType = "check_in"
	EventTypeCheckOut    EventType = "check_out"
	EventTypeTransfer    EventType = "transfer"
	EventTypeMaintenance EventType = "maintenance"
	EventTypeInspection  EventType = "inspection"
	EventTypeShipment    EventType = "shipment"
	EventTypeReturn      EventType = "return"
)

// MaintenanceRecord tracks scheduled and completed maintenance.
type MaintenanceRecord struct {
	ID          string    `json:"id" db:"id"`
	TenantID    string    `json:"tenant_id" db:"tenant_id"`
	AssetID     string    `json:"asset_id" db:"asset_id"`
	Type        string    `json:"type" db:"type"`
	ScheduledAt time.Time `json:"scheduled_at" db:"scheduled_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	TechnicianID string   `json:"technician_id,omitempty" db:"technician_id"`
	Cost        float64   `json:"cost" db:"cost"`
	Notes       string    `json:"notes,omitempty" db:"notes"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
