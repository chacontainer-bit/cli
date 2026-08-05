package postgres

import (
	"database/sql"
	"fmt"

	"github.com/cli/cli/v2/chacontainer/internal/domain/plant"
)

// PlantStore persists plants and their zones in PostgreSQL.
type PlantStore struct {
	db *sql.DB
}

func NewPlantStore(db *sql.DB) *PlantStore {
	return &PlantStore{db: db}
}

func (s *PlantStore) Create(tenantID string, p *plant.Plant) error {
	address, err := marshalJSON(p.Address)
	if err != nil {
		return err
	}
	capacity, err := marshalJSON(p.Capacity)
	if err != nil {
		return err
	}
	geo, err := marshalJSON(p.Geo)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(`
		INSERT INTO plants (tenant_id, code, name, address, capacity, status, manager_id, geo)
		VALUES ($1, $2, $3, $4::jsonb, $5::jsonb, COALESCE(NULLIF($6, ''), 'active'), $7, $8::jsonb)
		RETURNING id, created_at, updated_at
	`, tenantID, p.Code, p.Name, address, capacity, string(p.Status), nullStr(p.ManagerID), geo)

	if err := row.Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return fmt.Errorf("insert plant: %w", err)
	}
	p.TenantID = tenantID
	if p.Status == "" {
		p.Status = plant.StatusActive
	}
	return nil
}

func (s *PlantStore) GetByID(tenantID, id string) (*plant.Plant, error) {
	row := s.db.QueryRow(`
		SELECT id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at
		FROM plants WHERE tenant_id = $1 AND id = $2
	`, tenantID, id)
	return scanPlant(row)
}

func (s *PlantStore) List(tenantID string) ([]*plant.Plant, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at
		FROM plants WHERE tenant_id = $1 ORDER BY name
	`, tenantID)
	if err != nil {
		return nil, fmt.Errorf("list plants: %w", err)
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

var plantUpdatable = map[string]columnKind{
	"code":       kindText,
	"name":       kindText,
	"address":    kindJSON,
	"capacity":   kindJSON,
	"status":     kindText,
	"manager_id": kindText,
	"geo":        kindJSON,
}

func (s *PlantStore) Update(tenantID, id string, patch map[string]interface{}) (*plant.Plant, error) {
	setClause, args, err := buildUpdate(plantUpdatable, patch)
	if err != nil {
		return nil, err
	}
	args = append(args, tenantID, id)
	query := fmt.Sprintf(`
		UPDATE plants SET %s, updated_at = NOW()
		WHERE tenant_id = $%d AND id = $%d
		RETURNING id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at
	`, setClause, len(args)-1, len(args))

	row := s.db.QueryRow(query, args...)
	return scanPlant(row)
}

func (s *PlantStore) CreateZone(tenantID string, z *plant.Zone) error {
	row := s.db.QueryRow(`
		INSERT INTO plant_zones (tenant_id, plant_id, code, name, type, capacity)
		VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'warehouse'), $6)
		RETURNING id, created_at
	`, tenantID, z.PlantID, z.Code, z.Name, string(z.Type), z.Capacity)

	if err := row.Scan(&z.ID, &z.CreatedAt); err != nil {
		return fmt.Errorf("insert zone: %w", err)
	}
	z.TenantID = tenantID
	if z.Type == "" {
		z.Type = plant.ZoneTypeWarehouse
	}
	return nil
}

func (s *PlantStore) ListZones(tenantID, plantID string) ([]*plant.Zone, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, plant_id, code, name, type, capacity, created_at
		FROM plant_zones WHERE tenant_id = $1 AND plant_id = $2 ORDER BY code
	`, tenantID, plantID)
	if err != nil {
		return nil, fmt.Errorf("list zones: %w", err)
	}
	defer rows.Close()

	var out []*plant.Zone
	for rows.Next() {
		var z plant.Zone
		var typ string
		if err := rows.Scan(&z.ID, &z.TenantID, &z.PlantID, &z.Code, &z.Name, &typ, &z.Capacity, &z.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan zone: %w", err)
		}
		z.Type = plant.ZoneType(typ)
		out = append(out, &z)
	}
	return out, rows.Err()
}

// rowScanner abstracts *sql.Row and *sql.Rows so scanPlant works for both a
// single Get and a List loop.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

func scanPlant(row rowScanner) (*plant.Plant, error) {
	var p plant.Plant
	var status string
	var managerID sql.NullString
	var address, capacity, geo []byte
	if err := row.Scan(&p.ID, &p.TenantID, &p.Code, &p.Name, &address, &capacity, &status, &managerID, &geo, &p.CreatedAt, &p.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("plant not found")
		}
		return nil, fmt.Errorf("scan plant: %w", err)
	}
	p.Status = plant.Status(status)
	p.ManagerID = strOrEmpty(managerID)
	if err := unmarshalJSON(address, &p.Address); err != nil {
		return nil, fmt.Errorf("decode plant address: %w", err)
	}
	if err := unmarshalJSON(capacity, &p.Capacity); err != nil {
		return nil, fmt.Errorf("decode plant capacity: %w", err)
	}
	if err := unmarshalJSON(geo, &p.Geo); err != nil {
		return nil, fmt.Errorf("decode plant geo: %w", err)
	}
	return &p, nil
}
