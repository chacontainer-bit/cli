package sqlite

import (
	"database/sql"
	"strings"

	"github.com/chacontainer/backend/internal/domain/plant"
)

type PlantStore struct {
	db *sql.DB
}

func NewPlantStore(db *sql.DB) *PlantStore {
	return &PlantStore{db: db}
}

func (s *PlantStore) Create(tenantID string, p *plant.Plant) error {
	p.ID = newID()
	p.TenantID = tenantID
	now := parseTime(nowStr())
	p.CreatedAt = now
	p.UpdatedAt = now
	if p.Status == "" {
		p.Status = plant.StatusActive
	}
	_, err := s.db.Exec(
		`INSERT INTO plants (id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?)`,
		p.ID, p.TenantID, p.Code, p.Name, toJSON(p.Address), toJSON(p.Capacity), string(p.Status), nullStr(p.ManagerID),
		toJSON(p.Geo), fmtTime(p.CreatedAt), fmtTime(p.UpdatedAt),
	)
	return err
}

func scanPlant(row interface{ Scan(dest ...interface{}) error }) (*plant.Plant, error) {
	var p plant.Plant
	var address, capacity, geo string
	var managerID sql.NullString
	var status, createdAt, updatedAt string

	if err := row.Scan(&p.ID, &p.TenantID, &p.Code, &p.Name, &address, &capacity, &status, &managerID, &geo, &createdAt, &updatedAt); err != nil {
		return nil, err
	}
	p.Status = plant.Status(status)
	p.ManagerID = managerID.String
	fromJSON(address, &p.Address)
	fromJSON(capacity, &p.Capacity)
	fromJSON(geo, &p.Geo)
	p.CreatedAt = parseTime(createdAt)
	p.UpdatedAt = parseTime(updatedAt)
	return &p, nil
}

const plantSelectCols = `id, tenant_id, code, name, address, capacity, status, manager_id, geo, created_at, updated_at`

func (s *PlantStore) GetByID(tenantID, id string) (*plant.Plant, error) {
	row := s.db.QueryRow(`SELECT `+plantSelectCols+` FROM plants WHERE tenant_id = ? AND id = ?`, tenantID, id)
	return scanPlant(row)
}

func (s *PlantStore) List(tenantID string) ([]*plant.Plant, error) {
	rows, err := s.db.Query(`SELECT `+plantSelectCols+` FROM plants WHERE tenant_id = ? ORDER BY created_at DESC`, tenantID)
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

func (s *PlantStore) Update(tenantID, id string, patch map[string]interface{}) (*plant.Plant, error) {
	allowed := map[string]bool{"name": true, "status": true, "manager_id": true}
	var sets []string
	var args []interface{}
	for k, v := range patch {
		if !allowed[k] {
			continue
		}
		sets = append(sets, k+" = ?")
		args = append(args, v)
	}
	if len(sets) == 0 {
		return s.GetByID(tenantID, id)
	}
	sets = append(sets, "updated_at = ?")
	args = append(args, nowStr(), tenantID, id)
	if _, err := s.db.Exec(`UPDATE plants SET `+strings.Join(sets, ", ")+` WHERE tenant_id = ? AND id = ?`, args...); err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

func (s *PlantStore) CreateZone(tenantID string, z *plant.Zone) error {
	z.ID = newID()
	z.TenantID = tenantID
	z.CreatedAt = parseTime(nowStr())
	if z.Type == "" {
		z.Type = plant.ZoneTypeWarehouse
	}
	_, err := s.db.Exec(
		`INSERT INTO plant_zones (id, tenant_id, plant_id, code, name, type, capacity, created_at) VALUES (?,?,?,?,?,?,?,?)`,
		z.ID, z.TenantID, z.PlantID, z.Code, z.Name, string(z.Type), z.Capacity, fmtTime(z.CreatedAt),
	)
	return err
}

func (s *PlantStore) ListZones(tenantID, plantID string) ([]*plant.Zone, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, plant_id, code, name, type, capacity, created_at FROM plant_zones WHERE tenant_id = ? AND plant_id = ? ORDER BY created_at`,
		tenantID, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*plant.Zone
	for rows.Next() {
		var z plant.Zone
		var typeStr, createdAt string
		if err := rows.Scan(&z.ID, &z.TenantID, &z.PlantID, &z.Code, &z.Name, &typeStr, &z.Capacity, &createdAt); err != nil {
			return nil, err
		}
		z.Type = plant.ZoneType(typeStr)
		z.CreatedAt = parseTime(createdAt)
		out = append(out, &z)
	}
	return out, rows.Err()
}
