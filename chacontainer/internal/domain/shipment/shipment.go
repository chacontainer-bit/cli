package shipment

import "time"

type ShipmentStatus string

const (
	StatusDraft      ShipmentStatus = "draft"
	StatusConfirmed  ShipmentStatus = "confirmed"
	StatusInTransit  ShipmentStatus = "in_transit"
	StatusAtCustoms  ShipmentStatus = "at_customs"
	StatusDelivered  ShipmentStatus = "delivered"
	StatusCancelled  ShipmentStatus = "cancelled"
	StatusException  ShipmentStatus = "exception"
)

type TransportMode string

const (
	ModeRoad     TransportMode = "road"
	ModeSea      TransportMode = "sea"
	ModeAir      TransportMode = "air"
	ModeRail     TransportMode = "rail"
	ModeMultimodal TransportMode = "multimodal"
)

// Shipment represents a logistics movement of assets between locations.
type Shipment struct {
	ID              string         `json:"id" db:"id"`
	TenantID        string         `json:"tenant_id" db:"tenant_id"`
	Reference       string         `json:"reference" db:"reference"`
	ClientID        string         `json:"client_id" db:"client_id"`
	OriginPlantID   string         `json:"origin_plant_id" db:"origin_plant_id"`
	DestPlantID     string         `json:"dest_plant_id,omitempty" db:"dest_plant_id"`
	DestAddress     *Address       `json:"dest_address,omitempty" db:"dest_address"`
	Status          ShipmentStatus `json:"status" db:"status"`
	Mode            TransportMode  `json:"mode" db:"mode"`
	CarrierName     string         `json:"carrier_name,omitempty" db:"carrier_name"`
	CarrierRef      string         `json:"carrier_ref,omitempty" db:"carrier_ref"`
	TrackingURL     string         `json:"tracking_url,omitempty" db:"tracking_url"`
	ScheduledPickup time.Time      `json:"scheduled_pickup" db:"scheduled_pickup"`
	ActualPickup    *time.Time     `json:"actual_pickup,omitempty" db:"actual_pickup"`
	ScheduledDelivery time.Time    `json:"scheduled_delivery" db:"scheduled_delivery"`
	ActualDelivery  *time.Time     `json:"actual_delivery,omitempty" db:"actual_delivery"`
	TotalWeightKg   float64        `json:"total_weight_kg" db:"total_weight_kg"`
	TotalItems      int            `json:"total_items" db:"total_items"`
	Notes           string         `json:"notes,omitempty" db:"notes"`
	CustomsRef      string         `json:"customs_ref,omitempty" db:"customs_ref"`
	ERPID           string         `json:"erp_id,omitempty" db:"erp_id"`
	CreatedBy       string         `json:"created_by" db:"created_by"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

type Address struct {
	Street  string  `json:"street"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	ZipCode string  `json:"zip_code"`
	Lat     float64 `json:"lat,omitempty"`
	Lng     float64 `json:"lng,omitempty"`
}

// ShipmentLine links individual assets to a shipment.
type ShipmentLine struct {
	ID         string    `json:"id" db:"id"`
	TenantID   string    `json:"tenant_id" db:"tenant_id"`
	ShipmentID string    `json:"shipment_id" db:"shipment_id"`
	AssetID    string    `json:"asset_id" db:"asset_id"`
	Quantity   int       `json:"quantity" db:"quantity"`
	WeightKg   float64   `json:"weight_kg" db:"weight_kg"`
	Notes      string    `json:"notes,omitempty" db:"notes"`
	ScannedAt  *time.Time `json:"scanned_at,omitempty" db:"scanned_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// ShipmentEvent is an immutable audit trail entry for a shipment.
type ShipmentEvent struct {
	ID           string         `json:"id" db:"id"`
	TenantID     string         `json:"tenant_id" db:"tenant_id"`
	ShipmentID   string         `json:"shipment_id" db:"shipment_id"`
	FromStatus   ShipmentStatus `json:"from_status" db:"from_status"`
	ToStatus     ShipmentStatus `json:"to_status" db:"to_status"`
	Location     string         `json:"location,omitempty" db:"location"`
	Description  string         `json:"description" db:"description"`
	UserID       string         `json:"user_id" db:"user_id"`
	OccurredAt   time.Time      `json:"occurred_at" db:"occurred_at"`
}

func (s Shipment) IsLate() bool {
	if s.Status == StatusDelivered || s.ActualDelivery != nil {
		return false
	}
	return time.Now().After(s.ScheduledDelivery)
}

func (s Shipment) DaysLate() int {
	if !s.IsLate() {
		return 0
	}
	return int(time.Since(s.ScheduledDelivery).Hours() / 24)
}
