package api

// Stub implementations satisfy handler interfaces during development.
// Replace each stub with a concrete postgres store before production.

import (
	"fmt"
	"time"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/domain/asset"
	"github.com/cli/cli/v2/chacontainer/internal/domain/client"
	"github.com/cli/cli/v2/chacontainer/internal/domain/cycle"
	"github.com/cli/cli/v2/chacontainer/internal/domain/inspection"
	"github.com/cli/cli/v2/chacontainer/internal/domain/plant"
	"github.com/cli/cli/v2/chacontainer/internal/domain/recovery"
	"github.com/cli/cli/v2/chacontainer/internal/domain/shipment"
)

// ── Asset stub ────────────────────────────────────────────────────────────────

type stubAssetStore struct{}

func (s *stubAssetStore) Create(_ string, a *asset.Asset) error {
	a.ID = newID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}
func (s *stubAssetStore) GetByID(_, id string) (*asset.Asset, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubAssetStore) GetByQR(_, qr string) (*asset.Asset, error) {
	return nil, fmt.Errorf("not found: %s", qr)
}
func (s *stubAssetStore) List(_ string, _ handlers.AssetFilter) ([]*asset.Asset, int, error) {
	return []*asset.Asset{}, 0, nil
}
func (s *stubAssetStore) Update(_, id string, _ map[string]interface{}) (*asset.Asset, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubAssetStore) Delete(_, _ string) error                            { return nil }
func (s *stubAssetStore) RecordEvent(_ *asset.AssetEvent) error               { return nil }
func (s *stubAssetStore) ListEvents(_, _ string) ([]*asset.AssetEvent, error) { return nil, nil }

// ── Shipment stub ─────────────────────────────────────────────────────────────

type stubShipmentStore struct{}

func (s *stubShipmentStore) Create(_ string, sh *shipment.Shipment) error {
	sh.ID = newID()
	sh.CreatedAt = time.Now()
	sh.UpdatedAt = time.Now()
	return nil
}
func (s *stubShipmentStore) GetByID(_, id string) (*shipment.Shipment, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubShipmentStore) GetByReference(_, _ string) (*shipment.Shipment, error) {
	return nil, fmt.Errorf("not found")
}
func (s *stubShipmentStore) List(_ string, _ handlers.ShipmentFilter) ([]*shipment.Shipment, int, error) {
	return []*shipment.Shipment{}, 0, nil
}
func (s *stubShipmentStore) UpdateStatus(_, _ string, _ shipment.ShipmentStatus, _, _ string) error {
	return nil
}
func (s *stubShipmentStore) AddLine(_ string, _ *shipment.ShipmentLine) error { return nil }
func (s *stubShipmentStore) ListLines(_, _ string) ([]*shipment.ShipmentLine, error) {
	return nil, nil
}
func (s *stubShipmentStore) ListEvents(_, _ string) ([]*shipment.ShipmentEvent, error) {
	return nil, nil
}

// ── Client stub ───────────────────────────────────────────────────────────────

type stubClientStore struct{}

func (s *stubClientStore) Create(_ string, c *client.Client) error {
	c.ID = newID()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return nil
}
func (s *stubClientStore) GetByID(_, id string) (*client.Client, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubClientStore) List(_ string, _ handlers.ClientFilter) ([]*client.Client, int, error) {
	return []*client.Client{}, 0, nil
}
func (s *stubClientStore) Update(_, id string, _ map[string]interface{}) (*client.Client, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubClientStore) ListContacts(_, _ string) ([]*client.Contact, error)     { return nil, nil }
func (s *stubClientStore) CreateContact(_ string, _ *client.Contact) error         { return nil }
func (s *stubClientStore) CreateInteraction(_ string, _ *client.Interaction) error { return nil }
func (s *stubClientStore) ListInteractions(_, _ string) ([]*client.Interaction, error) {
	return nil, nil
}

// ── Plant stub ────────────────────────────────────────────────────────────────

type stubPlantStore struct{}

func (s *stubPlantStore) Create(_ string, p *plant.Plant) error {
	p.ID = newID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}
func (s *stubPlantStore) GetByID(_, id string) (*plant.Plant, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubPlantStore) List(_ string) ([]*plant.Plant, error) { return nil, nil }
func (s *stubPlantStore) Update(_, id string, _ map[string]interface{}) (*plant.Plant, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubPlantStore) CreateZone(_ string, _ *plant.Zone) error     { return nil }
func (s *stubPlantStore) ListZones(_, _ string) ([]*plant.Zone, error) { return nil, nil }

// ── Cycle stub ────────────────────────────────────────────────────────────────

type stubCycleStore struct{}

func (s *stubCycleStore) Create(_ string, c *cycle.Cycle) error {
	c.ID = newID()
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return nil
}
func (s *stubCycleStore) GetByID(_, id string) (*cycle.Cycle, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubCycleStore) List(_ string, _ handlers.CycleFilter) ([]*cycle.Cycle, int, error) {
	return []*cycle.Cycle{}, 0, nil
}
func (s *stubCycleStore) RecordMilestone(_, id string, _ cycle.Milestone, _ time.Time) (*cycle.Cycle, error) {
	return nil, fmt.Errorf("not found: %s", id)
}

// ── Inspection stub ──────────────────────────────────────────────────────────

type stubInspectionStore struct{}

func (s *stubInspectionStore) Create(_ string, insp *inspection.Inspection) error {
	insp.ID = newID()
	insp.CreatedAt = time.Now()
	return nil
}
func (s *stubInspectionStore) GetByID(_, id string) (*inspection.Inspection, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubInspectionStore) List(_ string, _ handlers.InspectionFilter) ([]*inspection.Inspection, int, error) {
	return []*inspection.Inspection{}, 0, nil
}
func (s *stubInspectionStore) Release(_, id string, _ inspection.Condition, _ time.Time) (*inspection.Inspection, error) {
	return nil, fmt.Errorf("not found: %s", id)
}

// ── Recovery stub (pilot config, value recovery ledger, ROI) ───────────────────

type stubRecoveryStore struct{}

func (s *stubRecoveryStore) CreatePilotConfig(_ string, pc *recovery.PilotConfig) error {
	pc.ID = newID()
	pc.CreatedAt = time.Now()
	pc.UpdatedAt = time.Now()
	return nil
}
func (s *stubRecoveryStore) GetPilotConfig(_, id string) (*recovery.PilotConfig, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubRecoveryStore) ListPilotConfigs(_ string) ([]*recovery.PilotConfig, error) {
	return []*recovery.PilotConfig{}, nil
}
func (s *stubRecoveryStore) CreateEntry(_ string, e *recovery.Entry) error {
	e.ID = newID()
	e.CreatedAt = time.Now()
	return nil
}
func (s *stubRecoveryStore) ListEntries(_ string, _ handlers.EntryFilter) ([]*recovery.Entry, int, error) {
	return []*recovery.Entry{}, 0, nil
}
func (s *stubRecoveryStore) ValidateEntry(_, id, _ string, _ recovery.ValidationStatus) (*recovery.Entry, error) {
	return nil, fmt.Errorf("not found: %s", id)
}
func (s *stubRecoveryStore) GetROISummary(_, pilotConfigID string) (*recovery.Summary, error) {
	return &recovery.Summary{
		PilotConfigID: pilotConfigID,
		Gate:          recovery.DecisionGate{},
		GeneratedAt:   time.Now(),
	}, nil
}

// ── Stats stub ────────────────────────────────────────────────────────────────

type stubStatsStore struct{}

func (s *stubStatsStore) GetDashboardStats(_ string) (*handlers.DashboardStats, error) {
	return &handlers.DashboardStats{
		AssetsByStatus:    map[string]int{"available": 0},
		AssetsByType:      map[string]int{},
		ShipmentsByStatus: map[string]int{},
		PlantOccupancy:    []handlers.PlantOccupancy{},
		TopClients:        []handlers.ClientActivity{},
		GeneratedAt:       time.Now(),
	}, nil
}
func (s *stubStatsStore) GetAssetTrend(_ string, _ int) ([]handlers.TrendPoint, error) {
	return nil, nil
}
func (s *stubStatsStore) GetShipmentTrend(_ string, _ int) ([]handlers.TrendPoint, error) {
	return nil, nil
}
func (s *stubStatsStore) GetPlantOccupancy(_ string) ([]handlers.PlantOccupancy, error) {
	return nil, nil
}

// ── Webhook stub ──────────────────────────────────────────────────────────────

type stubWebhookProcessor struct{}

func (s *stubWebhookProcessor) ProcessMakeEvent(_ handlers.WebhookEvent) error     { return nil }
func (s *stubWebhookProcessor) ProcessERPEvent(_ handlers.WebhookEvent) error      { return nil }
func (s *stubWebhookProcessor) ProcessAirtableEvent(_ handlers.WebhookEvent) error { return nil }

// ── helpers ───────────────────────────────────────────────────────────────────

func newID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
