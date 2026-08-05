package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chacontainer/backend/internal/domain/asset"
	"github.com/chacontainer/backend/internal/domain/client"
	"github.com/chacontainer/backend/internal/domain/plant"
	"github.com/chacontainer/backend/internal/domain/shipment"
	"github.com/chacontainer/backend/internal/domain/tenant"
	"golang.org/x/crypto/bcrypt"
)

// DemoCredentials are printed to the log after a fresh seed so the app is
// immediately usable without going through the registration form.
const (
	DemoEmail    = "demo@chacontainer.mx"
	DemoPassword = "Demo12345!"
)

// SeedIfEmpty populates a demo tenant with sample data the first time the
// database is created, so the app is fully explorable right after `go run`.
// It reports whether seeding actually happened.
func SeedIfEmpty(db *sql.DB) (bool, error) {
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM tenants`).Scan(&count); err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}

	auth := NewAuthStore(db)
	t := &tenant.Tenant{
		Name:         "CHACONTAINER Demo",
		Slug:         "chacontainer-demo",
		ContactEmail: DemoEmail,
		Settings: tenant.Settings{
			Currency: "MXN", DateFormat: "DD/MM/YYYY", QRTrackingEnabled: true,
			MaxContainers: 500, MaxPlants: 10, MaxUsers: 20,
		},
	}
	if err := auth.CreateTenant(t); err != nil {
		return false, fmt.Errorf("seed tenant: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(DemoPassword), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	admin := &tenant.User{TenantID: t.ID, Email: DemoEmail, Name: "Admin Demo", Role: tenant.RoleAdmin, PlantIDs: []string{}}
	if err := auth.CreateUser(admin, string(hash)); err != nil {
		return false, fmt.Errorf("seed user: %w", err)
	}

	plants := NewPlantStore(db)
	p1 := &plant.Plant{
		Code: "PLT-QRO", Name: "Planta Querétaro",
		Address:  plant.Address{City: "Querétaro", State: "Querétaro", Country: "MX"},
		Capacity: plant.Capacity{TotalSlots: 200, UsedSlots: 0, MaxWeight: 50000},
		Geo:      plant.GeoPoint{Lat: 20.5888, Lng: -100.3899},
	}
	if err := plants.Create(t.ID, p1); err != nil {
		return false, fmt.Errorf("seed plant: %w", err)
	}
	p2 := &plant.Plant{
		Code: "PLT-SLP", Name: "Planta San Luis Potosí",
		Address:  plant.Address{City: "San Luis Potosí", State: "SLP", Country: "MX"},
		Capacity: plant.Capacity{TotalSlots: 120, UsedSlots: 0, MaxWeight: 30000},
		Geo:      plant.GeoPoint{Lat: 22.1565, Lng: -100.9855},
	}
	if err := plants.Create(t.ID, p2); err != nil {
		return false, fmt.Errorf("seed plant 2: %w", err)
	}

	clients := NewClientStore(db)
	c1 := &client.Client{
		Code: "CLI-001", Name: "Autopartes del Bajío", Type: client.ClientTypeManufacturer,
		Status: client.ClientStatusActive, Industry: "Automotriz", CreditLimit: 500000, PaymentTerms: 30,
		Currency: "MXN", Address: client.Address{City: "Querétaro", Country: "MX"},
	}
	if err := clients.Create(t.ID, c1); err != nil {
		return false, fmt.Errorf("seed client: %w", err)
	}

	assets := NewAssetStore(db)
	sampleAssets := []*asset.Asset{
		{PlantID: p1.ID, Name: "Contenedor plegable 600L", Code: "AST-0001", Type: asset.AssetTypeContainer, Status: asset.AssetStatusAvailable, WeightKg: 18, Dimensions: asset.Dimensions{LengthCm: 120, WidthCm: 100, HeightCm: 78}},
		{PlantID: p1.ID, Name: "Rack metálico apilable", Code: "AST-0002", Type: asset.AssetTypeRack, Status: asset.AssetStatusInUse, ClientID: c1.ID, WeightKg: 45, Dimensions: asset.Dimensions{LengthCm: 140, WidthCm: 110, HeightCm: 160}},
		{PlantID: p2.ID, Name: "IBC 1000L", Code: "AST-0003", Type: asset.AssetTypeIBC, Status: asset.AssetStatusTransit, WeightKg: 60, Dimensions: asset.Dimensions{LengthCm: 120, WidthCm: 100, HeightCm: 116}},
	}
	for _, a := range sampleAssets {
		if err := assets.Create(t.ID, a); err != nil {
			return false, fmt.Errorf("seed asset %s: %w", a.Code, err)
		}
		_ = assets.RecordEvent(&asset.AssetEvent{
			TenantID: t.ID, AssetID: a.ID, EventType: asset.EventTypeCheckIn,
			ToStatus: a.Status, UserID: admin.ID, OccurredAt: time.Now(),
		})
	}

	shipments := NewShipmentStore(db)
	sh := &shipment.Shipment{
		ClientID: c1.ID, OriginPlantID: p1.ID, DestPlantID: p2.ID,
		Mode: shipment.ModeRoad, CarrierName: "Transportes Bajío",
		ScheduledPickup: time.Now().Add(-48 * time.Hour), ScheduledDelivery: time.Now().Add(24 * time.Hour),
		CreatedBy: admin.ID,
	}
	if err := shipments.Create(t.ID, sh); err != nil {
		return false, fmt.Errorf("seed shipment: %w", err)
	}
	if err := shipments.UpdateStatus(t.ID, sh.ID, shipment.StatusInTransit, admin.ID, "salió de planta Querétaro"); err != nil {
		return false, fmt.Errorf("seed shipment transition: %w", err)
	}

	return true, nil
}
