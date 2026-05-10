package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/asset"
	"github.com/lib/pq"
)

// AssetStore implements handlers.AssetStore against PostgreSQL.
type AssetStore struct{ db *sql.DB }

func NewAssetStore(db *sql.DB) *AssetStore { return &AssetStore{db: db} }

// ── Create ────────────────────────────────────────────────────────────────────

func (s *AssetStore) Create(tenantID string, a *asset.Asset) error {
	a.ID = newUUID()
	a.TenantID = tenantID
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	if a.Status == "" {
		a.Status = asset.AssetStatusAvailable
	}

	dimJSON := marshalJSON(a.Dimensions)
	metaJSON := marshalJSON(a.Metadata)

	_, err := s.db.Exec(`
		INSERT INTO assets (
			id, tenant_id, plant_id, zone_id, name, code, qr_code, barcode,
			type, status, description, manufacturer, model, serial_number,
			weight_kg, dimensions, purchase_date, purchase_cost,
			client_id, metadata, airtable_id, created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,
			$9,$10,$11,$12,$13,$14,
			$15,$16,$17,$18,
			$19,$20,$21,$22,$23
		)`,
		a.ID, tenantID,
		nullString(a.PlantID), nullString(a.ZoneID),
		a.Name, a.Code, a.QRCode, nullString(a.Barcode),
		string(a.Type), string(a.Status),
		nullString(a.Description), nullString(a.Manufacturer),
		nullString(a.Model), nullString(a.SerialNumber),
		a.WeightKg, dimJSON, nullTime(a.PurchaseDate), nullFloat(a.PurchaseCost),
		nullString(a.ClientID), metaJSON, nullString(a.AirtableID),
		a.CreatedAt, a.UpdatedAt,
	)
	return err
}

// ── GetByID ───────────────────────────────────────────────────────────────────

func (s *AssetStore) GetByID(tenantID, id string) (*asset.Asset, error) {
	row := s.db.QueryRow(`
		SELECT `+assetColumns+`
		FROM assets
		WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanAsset(row)
}

// ── GetByQR ───────────────────────────────────────────────────────────────────

func (s *AssetStore) GetByQR(tenantID, qrCode string) (*asset.Asset, error) {
	row := s.db.QueryRow(`
		SELECT `+assetColumns+`
		FROM assets
		WHERE tenant_id = $1 AND qr_code = $2`, tenantID, qrCode)
	return scanAsset(row)
}

// ── List ──────────────────────────────────────────────────────────────────────

func (s *AssetStore) List(tenantID string, f handlers.AssetFilter) ([]*asset.Asset, int, error) {
	where := []string{"tenant_id = $1"}
	args := []interface{}{tenantID}
	n := 2

	if f.PlantID != "" {
		where = append(where, fmt.Sprintf("plant_id = $%d", n))
		args = append(args, f.PlantID)
		n++
	}
	if f.ZoneID != "" {
		where = append(where, fmt.Sprintf("zone_id = $%d", n))
		args = append(args, f.ZoneID)
		n++
	}
	if f.Type != "" {
		where = append(where, fmt.Sprintf("type = $%d", n))
		args = append(args, f.Type)
		n++
	}
	if f.Status != "" {
		where = append(where, fmt.Sprintf("status = $%d", n))
		args = append(args, f.Status)
		n++
	}
	if f.ClientID != "" {
		where = append(where, fmt.Sprintf("client_id = $%d", n))
		args = append(args, f.ClientID)
		n++
	}
	if f.Search != "" {
		where = append(where, fmt.Sprintf("(name ILIKE $%d OR code ILIKE $%d)", n, n))
		args = append(args, "%"+f.Search+"%")
		n++
	}

	clause := "WHERE " + strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM assets "+clause, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	if f.Page < 1 {
		f.Page = 1
	}
	if f.PerPage < 1 || f.PerPage > 200 {
		f.PerPage = 50
	}
	offset := (f.Page - 1) * f.PerPage

	args = append(args, f.PerPage, offset)
	rows, err := s.db.Query(
		`SELECT `+assetColumns+` FROM assets `+clause+
			fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, n, n+1),
		args...,
	)
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

// ── Update ────────────────────────────────────────────────────────────────────

// allowedAssetPatch lists every column that external callers may update.
var allowedAssetPatch = map[string]bool{
	"name": true, "status": true, "plant_id": true, "zone_id": true,
	"description": true, "client_id": true, "last_scan_at": true,
	"last_scan_by": true, "metadata": true, "airtable_id": true,
	"barcode": true, "weight_kg": true,
}

func (s *AssetStore) Update(tenantID, id string, patch map[string]interface{}) (*asset.Asset, error) {
	sets := []string{}
	args := []interface{}{}
	n := 1

	for k, v := range patch {
		if !allowedAssetPatch[k] {
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
		`UPDATE assets SET `+strings.Join(sets, ", ")+
			fmt.Sprintf(` WHERE tenant_id = $%d AND id = $%d`, n, n+1),
		args...,
	)
	if err != nil {
		return nil, err
	}
	return s.GetByID(tenantID, id)
}

// ── Delete ────────────────────────────────────────────────────────────────────

func (s *AssetStore) Delete(tenantID, id string) error {
	_, err := s.db.Exec(
		`UPDATE assets SET status = 'retired', updated_at = NOW()
		 WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return err
}

// ── RecordEvent ───────────────────────────────────────────────────────────────

func (s *AssetStore) RecordEvent(e *asset.AssetEvent) error {
	if e.ID == "" {
		e.ID = newUUID()
	}
	e.CreatedAt = time.Now()

	geoJSON := marshalJSON(e.Geo)
	_, err := s.db.Exec(`
		INSERT INTO asset_events (
			id, tenant_id, asset_id, event_type,
			from_status, to_status, from_plant_id, to_plant_id,
			user_id, shipment_id, notes, geo, occurred_at, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`,
		e.ID, e.TenantID, e.AssetID, string(e.EventType),
		nullString(string(e.FromStatus)), nullString(string(e.ToStatus)),
		nullString(e.FromPlantID), nullString(e.ToPlantID),
		e.UserID, nullString(e.ShipmentID), nullString(e.Notes),
		geoJSON, e.OccurredAt, e.CreatedAt,
	)
	return err
}

// ── ListEvents ────────────────────────────────────────────────────────────────

func (s *AssetStore) ListEvents(tenantID, assetID string) ([]*asset.AssetEvent, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, asset_id, event_type,
		       COALESCE(from_status,''), COALESCE(to_status,''),
		       COALESCE(from_plant_id::text,''), COALESCE(to_plant_id::text,''),
		       user_id::text, COALESCE(shipment_id::text,''),
		       COALESCE(notes,''), geo, occurred_at, created_at
		FROM asset_events
		WHERE tenant_id = $1 AND asset_id = $2
		ORDER BY occurred_at DESC
		LIMIT 200`, tenantID, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*asset.AssetEvent
	for rows.Next() {
		e := &asset.AssetEvent{}
		var evType, fromStatus, toStatus string
		var geoJSON []byte
		if err := rows.Scan(
			&e.ID, &e.TenantID, &e.AssetID, &evType,
			&fromStatus, &toStatus,
			&e.FromPlantID, &e.ToPlantID,
			&e.UserID, &e.ShipmentID,
			&e.Notes, &geoJSON, &e.OccurredAt, &e.CreatedAt,
		); err != nil {
			return nil, err
		}
		e.EventType = asset.EventType(evType)
		e.FromStatus = asset.AssetStatus(fromStatus)
		e.ToStatus = asset.AssetStatus(toStatus)
		if len(geoJSON) > 0 {
			e.Geo = &asset.GeoPoint{}
			scanJSON(geoJSON, e.Geo)
		}
		out = append(out, e)
	}
	return out, rows.Err()
}

// ── SQL helpers ───────────────────────────────────────────────────────────────

const assetColumns = `
	id, tenant_id,
	COALESCE(plant_id::text,''), COALESCE(zone_id::text,''),
	COALESCE(name,''), code, qr_code, COALESCE(barcode,''),
	type, status,
	COALESCE(description,''), COALESCE(manufacturer,''),
	COALESCE(model,''), COALESCE(serial_number,''),
	weight_kg, dimensions,
	purchase_date, COALESCE(purchase_cost,0),
	COALESCE(client_id::text,''), metadata, COALESCE(airtable_id,''),
	last_scan_at, COALESCE(last_scan_by::text,''),
	created_at, updated_at`

type assetScanner interface {
	Scan(dest ...interface{}) error
}

func scanAsset(row assetScanner) (*asset.Asset, error) {
	a := &asset.Asset{}
	var typ, status string
	var dimJSON, metaJSON []byte
	var purchaseDate sql.NullTime
	var lastScanAt sql.NullTime

	err := row.Scan(
		&a.ID, &a.TenantID,
		&a.PlantID, &a.ZoneID,
		&a.Name, &a.Code, &a.QRCode, &a.Barcode,
		&typ, &status,
		&a.Description, &a.Manufacturer,
		&a.Model, &a.SerialNumber,
		&a.WeightKg, &dimJSON,
		&purchaseDate, &a.PurchaseCost,
		&a.ClientID, &metaJSON, &a.AirtableID,
		&lastScanAt, &a.LastScanBy,
		&a.CreatedAt, &a.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("asset not found")
	}
	if err != nil {
		return nil, err
	}

	a.Type = asset.AssetType(typ)
	a.Status = asset.AssetStatus(status)

	scanJSON(dimJSON, &a.Dimensions)
	scanJSON(metaJSON, &a.Metadata)

	if purchaseDate.Valid {
		t := purchaseDate.Time
		a.PurchaseDate = &t
	}
	if lastScanAt.Valid {
		t := lastScanAt.Time
		a.LastScanAt = &t
	}
	return a, nil
}

func nullFloat(f float64) interface{} {
	if f == 0 {
		return nil
	}
	return f
}

// compile-time interface check
var _ handlers.AssetStore = (*AssetStore)(nil)

// suppress unused import warning
var _ = pq.Array
