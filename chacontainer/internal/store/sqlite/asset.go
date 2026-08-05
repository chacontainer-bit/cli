package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/asset"
)

type AssetStore struct {
	db *sql.DB
}

func NewAssetStore(db *sql.DB) *AssetStore {
	return &AssetStore{db: db}
}

func (s *AssetStore) Create(tenantID string, a *asset.Asset) error {
	a.ID = newID()
	a.TenantID = tenantID
	now := parseTime(nowStr())
	a.CreatedAt = now
	a.UpdatedAt = now
	if a.Status == "" {
		a.Status = asset.AssetStatusAvailable
	}
	if a.QRCode == "" {
		a.QRCode = fmt.Sprintf("%s|%s|%s|1|%s", tenantID, a.ID, a.Code, a.ID[:8])
	}
	_, err := s.db.Exec(
		`INSERT INTO assets (id, tenant_id, plant_id, zone_id, name, code, airtable_id, qr_code, barcode, type, status,
			description, manufacturer, model, serial_number, weight_kg, dimensions, purchase_date, purchase_cost,
			client_id, metadata, last_scan_at, last_scan_by, created_at, updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		a.ID, a.TenantID, a.PlantID, nullStr(a.ZoneID), a.Name, a.Code, nullStr(a.AirtableID), a.QRCode, nullStr(a.Barcode),
		string(a.Type), string(a.Status), nullStr(a.Description), nullStr(a.Manufacturer), nullStr(a.Model), nullStr(a.SerialNumber),
		a.WeightKg, toJSON(a.Dimensions), nullTimePtr(a.PurchaseDate), a.PurchaseCost, nullStr(a.ClientID), toJSON(a.Metadata),
		nullTimePtr(a.LastScanAt), nullStr(a.LastScanBy), fmtTime(a.CreatedAt), fmtTime(a.UpdatedAt),
	)
	return err
}

func scanAsset(row interface {
	Scan(dest ...interface{}) error
}) (*asset.Asset, error) {
	var a asset.Asset
	var zoneID, airtableID, barcode, description, manufacturer, model, serial, clientID, lastScanBy sql.NullString
	var purchaseDate sql.NullString
	var lastScanAt sql.NullString
	var dimensions, metadata string
	var typeStr, statusStr string
	var createdAt, updatedAt string
	var purchaseCostF sql.NullFloat64

	err := row.Scan(
		&a.ID, &a.TenantID, &a.PlantID, &zoneID, &a.Name, &a.Code, &airtableID, &a.QRCode, &barcode,
		&typeStr, &statusStr, &description, &manufacturer, &model, &serial,
		&a.WeightKg, &dimensions, &purchaseDate, &purchaseCostF, &clientID, &metadata,
		&lastScanAt, &lastScanBy, &createdAt, &updatedAt,
	)
	if err != nil {
		return nil, err
	}

	a.ZoneID = zoneID.String
	a.AirtableID = airtableID.String
	a.Barcode = barcode.String
	a.Type = asset.AssetType(typeStr)
	a.Status = asset.AssetStatus(statusStr)
	a.Description = description.String
	a.Manufacturer = manufacturer.String
	a.Model = model.String
	a.SerialNumber = serial.String
	a.ClientID = clientID.String
	a.LastScanBy = lastScanBy.String
	a.LastScanAt = parseTimePtr(lastScanAt)
	if purchaseDate.Valid {
		t := parseTime(purchaseDate.String)
		a.PurchaseDate = &t
	}
	if purchaseCostF.Valid {
		a.PurchaseCost = purchaseCostF.Float64
	}
	fromJSON(dimensions, &a.Dimensions)
	fromJSON(metadata, &a.Metadata)
	a.CreatedAt = parseTime(createdAt)
	a.UpdatedAt = parseTime(updatedAt)
	return &a, nil
}

const assetSelectCols = `id, tenant_id, plant_id, zone_id, name, code, airtable_id, qr_code, barcode, type, status,
	description, manufacturer, model, serial_number, weight_kg, dimensions, purchase_date, purchase_cost,
	client_id, metadata, last_scan_at, last_scan_by, created_at, updated_at`

func (s *AssetStore) GetByID(tenantID, id string) (*asset.Asset, error) {
	row := s.db.QueryRow(`SELECT `+assetSelectCols+` FROM assets WHERE tenant_id = ? AND id = ?`, tenantID, id)
	return scanAsset(row)
}

func (s *AssetStore) GetByQR(tenantID, qr string) (*asset.Asset, error) {
	row := s.db.QueryRow(`SELECT `+assetSelectCols+` FROM assets WHERE tenant_id = ? AND qr_code = ?`, tenantID, qr)
	return scanAsset(row)
}

func (s *AssetStore) List(tenantID string, filter handlers.AssetFilter) ([]*asset.Asset, int, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{tenantID}

	if filter.PlantID != "" {
		where = append(where, "plant_id = ?")
		args = append(args, filter.PlantID)
	}
	if filter.ZoneID != "" {
		where = append(where, "zone_id = ?")
		args = append(args, filter.ZoneID)
	}
	if filter.Type != "" {
		where = append(where, "type = ?")
		args = append(args, filter.Type)
	}
	if filter.Status != "" {
		where = append(where, "status = ?")
		args = append(args, filter.Status)
	}
	if filter.ClientID != "" {
		where = append(where, "client_id = ?")
		args = append(args, filter.ClientID)
	}
	if filter.Search != "" {
		where = append(where, "(name LIKE ? OR code LIKE ?)")
		like := "%" + filter.Search + "%"
		args = append(args, like, like)
	}
	whereClause := strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM assets WHERE `+whereClause, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	page, perPage := filter.Page, filter.PerPage
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 50
	}
	offset := (page - 1) * perPage

	queryArgs := append(append([]interface{}{}, args...), perPage, offset)
	rows, err := s.db.Query(`SELECT `+assetSelectCols+` FROM assets WHERE `+whereClause+` ORDER BY created_at DESC LIMIT ? OFFSET ?`, queryArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []*asset.Asset
	for rows.Next() {
		a, err := scanAsset(rows)
		if err != nil {
			return nil, 0, err
		}
		out = append(out, a)
	}
	return out, total, rows.Err()
}

func (s *AssetStore) Update(tenantID, id string, patch map[string]interface{}) (*asset.Asset, error) {
	allowed := map[string]bool{
		"name": true, "plant_id": true, "zone_id": true, "status": true, "description": true,
		"manufacturer": true, "model": true, "serial_number": true, "weight_kg": true,
		"client_id": true, "last_scan_at": true, "last_scan_by": true, "barcode": true,
	}
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
	args = append(args, nowStr())
	args = append(args, tenantID, id)

	_, err := s.db.Exec(`UPDATE assets SET `+strings.Join(sets, ", ")+` WHERE tenant_id = ? AND id = ?`, args...)
	if err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

func (s *AssetStore) Delete(tenantID, id string) error {
	_, err := s.db.Exec(`DELETE FROM assets WHERE tenant_id = ? AND id = ?`, tenantID, id)
	return err
}

func (s *AssetStore) RecordEvent(e *asset.AssetEvent) error {
	e.ID = newID()
	now := nowStr()
	if e.OccurredAt.IsZero() {
		e.OccurredAt = parseTime(now)
	}
	e.CreatedAt = parseTime(now)
	_, err := s.db.Exec(
		`INSERT INTO asset_events (id, tenant_id, asset_id, event_type, from_status, to_status, from_plant_id, to_plant_id,
			user_id, shipment_id, notes, geo, occurred_at, created_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		e.ID, e.TenantID, e.AssetID, string(e.EventType), string(e.FromStatus), string(e.ToStatus),
		nullStr(e.FromPlantID), nullStr(e.ToPlantID), e.UserID, nullStr(e.ShipmentID), nullStr(e.Notes),
		nullGeoJSON(e.Geo), fmtTime(e.OccurredAt), fmtTime(e.CreatedAt),
	)
	return err
}

func (s *AssetStore) ListEvents(tenantID, assetID string) ([]*asset.AssetEvent, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, asset_id, event_type, from_status, to_status, from_plant_id, to_plant_id,
			user_id, shipment_id, notes, geo, occurred_at, created_at
		 FROM asset_events WHERE tenant_id = ? AND asset_id = ? ORDER BY occurred_at DESC`, tenantID, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*asset.AssetEvent
	for rows.Next() {
		var e asset.AssetEvent
		var fromStatus, toStatus, fromPlant, toPlant, shipmentID, notes, geo sql.NullString
		var eventType, occurredAt, createdAt string
		if err := rows.Scan(&e.ID, &e.TenantID, &e.AssetID, &eventType, &fromStatus, &toStatus, &fromPlant, &toPlant,
			&e.UserID, &shipmentID, &notes, &geo, &occurredAt, &createdAt); err != nil {
			return nil, err
		}
		e.EventType = asset.EventType(eventType)
		e.FromStatus = asset.AssetStatus(fromStatus.String)
		e.ToStatus = asset.AssetStatus(toStatus.String)
		e.FromPlantID = fromPlant.String
		e.ToPlantID = toPlant.String
		e.ShipmentID = shipmentID.String
		e.Notes = notes.String
		if geo.Valid && geo.String != "" {
			var g asset.GeoPoint
			fromJSON(geo.String, &g)
			e.Geo = &g
		}
		e.OccurredAt = parseTime(occurredAt)
		e.CreatedAt = parseTime(createdAt)
		out = append(out, &e)
	}
	return out, rows.Err()
}

func nullStr(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

func nullGeoJSON(g *asset.GeoPoint) interface{} {
	if g == nil {
		return nil
	}
	return toJSON(g)
}
