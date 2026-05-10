package plant

import "time"

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

// Plant represents a physical facility/warehouse/plant.
type Plant struct {
	ID        string    `json:"id" db:"id"`
	TenantID  string    `json:"tenant_id" db:"tenant_id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Address   Address   `json:"address" db:"address"`
	Capacity  Capacity  `json:"capacity" db:"capacity"`
	Status    Status    `json:"status" db:"status"`
	ManagerID string    `json:"manager_id" db:"manager_id"`
	Geo       GeoPoint  `json:"geo" db:"geo"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

type Capacity struct {
	TotalSlots int `json:"total_slots"`
	UsedSlots  int `json:"used_slots"`
	MaxWeight  float64 `json:"max_weight_kg"`
}

func (c Capacity) AvailableSlots() int {
	return c.TotalSlots - c.UsedSlots
}

func (c Capacity) OccupancyPct() float64 {
	if c.TotalSlots == 0 {
		return 0
	}
	return float64(c.UsedSlots) / float64(c.TotalSlots) * 100
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Zone represents a section within a plant (rack, bay, dock, etc.)
type Zone struct {
	ID        string    `json:"id" db:"id"`
	PlantID   string    `json:"plant_id" db:"plant_id"`
	TenantID  string    `json:"tenant_id" db:"tenant_id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Type      ZoneType  `json:"type" db:"type"`
	Capacity  int       `json:"capacity" db:"capacity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ZoneType string

const (
	ZoneTypeWarehouse ZoneType = "warehouse"
	ZoneTypeDock      ZoneType = "dock"
	ZoneTypeYard      ZoneType = "yard"
	ZoneTypeProduction ZoneType = "production"
	ZoneTypeQuarantine ZoneType = "quarantine"
)
