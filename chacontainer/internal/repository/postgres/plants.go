package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/plant"
)

// PlantStore implements handlers.PlantStore against PostgreSQL.
type PlantStore struct{ db *sql.DB }

func NewPlantStore(db *sql.DB) *PlantStore { return &PlantStore{db: db} }

// ── Create ────────────────────────────────────────────────────────────────────

func (s *PlantStore) Create(tenantID string, p *plant.Plant) error {
	p.ID = newUUID()
	p.TenantID = tenantID
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	if p.Status == "" {
		p.Status = plant.StatusActive
	}

	addrJSON := marshalJSON(p.Address)
	capJSON := marshalJSON(p.Capacity)
	geoJSON := marshalJSON(p.Geo)

	_, err := s.db.Exec(`
		INSERT INTO plants (id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		p.ID, tenantID, p.Code, p.Name,
		addrJSON, capJSON, string(p.Status),
		nullString(p.ManagerID), geoJSON,
		p.CreatedAt, p.UpdatedAt,
	)
	return err
}

// ── GetByID ───────────────────────────────────────────────────────────────────

func (s *PlantStore) GetByID(tenantID, id string) (*plant.Plant, error) {
	row := s.db.QueryRow(`SELECT `+plantCols+` FROM plants
		WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanPlant(row)
}

// ── List ──────────────────────────────────────────────────────────────────────

func (s *PlantStore) List(tenantID string) ([]*plant.Plant, error) {
	rows, err := s.db.Query(`SELECT `+plantCols+` FROM plants
		WHERE tenant_id = $1 ORDER BY name`, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*plant.Plant
	for rows.Next() {
		p, err := scanPlant(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// ── Update ────────────────────────────────────────────────────────────────────

var allowedPlantPatch = map[string]bool{
	"name": true, "status": true, "manager_id": true,
}

func (s *PlantStore) Update(tenantID, id string, patch map[string]interface{}) (*plant.Plant, error) {
	sets := []string{}
	args := []interface{}{}
	n := 1

	// Handle JSONB sub-objects specially
	if addr, ok := patch["address"]; ok {
		sets = append(sets, fmt.Sprintf("address = $%d", n))
		args = append(args, marshalJSON(addr))
		n++
	}
	if cap, ok := patch["capacity"]; ok {
		sets = append(sets, fmt.Sprintf("capacity = $%d", n))
		args = append(args, marshalJSON(cap))
		n++
	}
	if geo, ok := patch["geo"]; ok {
		sets = append(sets, fmt.Sprintf("geo = $%d", n))
		args = append(args, marshalJSON(geo))
		n++
	}

	for k, v := range patch {
		if !allowedPlantPatch[k] {
			continue
		}
		sets = append(sets, fmt.Sprintf("%s = $%d", k, n))
		args = append(args, v)
		n++
	}
	if len(sets) == 0 {
		return s.GetByID(tenantID, id)
	}

	sets = append(sets, fmt.Sprintf("updated_at = $%d", n))
	args = append(args, time.Now())
	n++

	args = append(args, tenantID, id)
	_, err := s.db.Exec(
		`UPDATE plants SET `+strings.Join(sets, ", ")+
			fmt.Sprintf(` WHERE tenant_id = $%d AND id = $%d`, n, n+1),
		args...,
	)
	if err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

// ── Zones ─────────────────────────────────────────────────────────────────────

func (s *PlantStore) CreateZone(tenantID string, z *plant.Zone) error {
	z.ID = newUUID()
	z.TenantID = tenantID
	z.CreatedAt = time.Now()

	_, err := s.db.Exec(`
		INSERT INTO plant_zones (id, tenant_id, plant_id, code, name, type, capacity, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		z.ID, tenantID, z.PlantID, z.Code, z.Name, string(z.Type), z.Capacity, z.CreatedAt,
	)
	return err
}

func (s *PlantStore) ListZones(tenantID, plantID string) ([]*plant.Zone, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, plant_id::text, code, name, type, capacity, created_at
		FROM plant_zones
		WHERE tenant_id = $1 AND plant_id = $2
		ORDER BY code`, tenantID, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*plant.Zone
	for rows.Next() {
		z := &plant.Zone{}
		var zType string
		if err := rows.Scan(&z.ID, &z.TenantID, &z.PlantID,
			&z.Code, &z.Name, &zType, &z.Capacity, &z.CreatedAt); err != nil {
			return nil, err
		}
		z.Type = plant.ZoneType(zType)
		out = append(out, z)
	}
	return out, rows.Err()
}

// ── SQL helpers ───────────────────────────────────────────────────────────────

const plantCols = `
	id, tenant_id, code, name,
	address, capacity, status,
	COALESCE(manager_id::text,''), geo,
	created_at, updated_at`

type plantScanner interface {
	Scan(dest ...interface{}) error
}

func scanPlant(row plantScanner) (*plant.Plant, error) {
	p := &plant.Plant{}
	var status string
	var addrJSON, capJSON, geoJSON []byte

	err := row.Scan(
		&p.ID, &p.TenantID, &p.Code, &p.Name,
		&addrJSON, &capJSON, &status,
		&p.ManagerID, &geoJSON,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("plant not found")
	}
	if err != nil {
		return nil, err
	}

	p.Status = plant.Status(status)
	scanJSON(addrJSON, &p.Address)
	scanJSON(capJSON, &p.Capacity)
	scanJSON(geoJSON, &p.Geo)
	return p, nil
}

var _ handlers.PlantStore = (*PlantStore)(nil)
