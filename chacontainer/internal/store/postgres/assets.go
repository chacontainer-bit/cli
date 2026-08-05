package postgres

import (
	"database/sql"
	"fmt"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/domain/asset"
)

// AssetStore persists assets and their event trail in PostgreSQL.
type AssetStore struct {
	db *sql.DB
}

func NewAssetStore(db *sql.DB) *AssetStore {
	return &AssetStore{db: db}
}

const assetSelect = `
	SELECT id, tenant_id, plant_id, zone_id, name, code, airtable_id, qr_code, barcode, type, status,
		description, manufacturer, model, serial_number, weight_kg, dimensions, purchase_date,
		purchase_cost, client_id, metadata, last_scan_at, last_scan_by, created_at, updated_at
	FROM assets`

func (s *AssetStore) Create(tenantID string, a *asset.Asset) error {
	dimensions, err := marshalJSON(a.Dimensions)
	if err != nil {
		return err
	}
	metadata, err := marshalJSON(a.Metadata)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(`
		INSERT INTO assets (tenant_id, plant_id, zone_id, name, code, airtable_id, qr_code, barcode, type, status,
			description, manufacturer, model, serial_number, weight_kg, dimensions, purchase_date,
			purchase_cost, client_id, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, COALESCE(NULLIF($10, ''), 'available'),
			$11, $12, $13, $14, $15, $16::jsonb, $17,
			$18, $19, $20::jsonb)
		RETURNING id, created_at, updated_at
	`, tenantID, a.PlantID, nullStr(a.ZoneID), nullStr(a.Name), a.Code, nullStr(a.AirtableID), a.QRCode, nullStr(a.Barcode),
		string(a.Type), string(a.Status), nullStr(a.Description), nullStr(a.Manufacturer), nullStr(a.Model),
		nullStr(a.SerialNumber), a.WeightKg, dimensions, nullDate(a.PurchaseDate),
		nullFloat(a.PurchaseCost), nullStr(a.ClientID), metadata)

	if err := row.Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt); err != nil {
		return fmt.Errorf("insert asset: %w", err)
	}
	a.TenantID = tenantID
	if a.Status == "" {
		a.Status = asset.AssetStatusAvailable
	}
	return nil
}

func (s *AssetStore) GetByID(tenantID, id string) (*asset.Asset, error) {
	row := s.db.QueryRow(assetSelect+` WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanAsset(row)
}

func (s *AssetStore) GetByQR(tenantID, qrCode string) (*asset.Asset, error) {
	row := s.db.QueryRow(assetSelect+` WHERE tenant_id = $1 AND qr_code = $2`, tenantID, qrCode)
	return scanAsset(row)
}

func (s *AssetStore) List(tenantID string, filter handlers.AssetFilter) ([]*asset.Asset, int, error) {
	where := "WHERE tenant_id = $1"
	args := []interface{}{tenantID}

	addFilter := func(col, val string) {
		if val == "" {
			return
		}
		args = append(args, val)
		where += fmt.Sprintf(" AND %s = $%d", col, len(args))
	}
	addFilter("plant_id", filter.PlantID)
	addFilter("zone_id", filter.ZoneID)
	addFilter("type", filter.Type)
	addFilter("status", filter.Status)
	addFilter("client_id", filter.ClientID)
	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		where += fmt.Sprintf(" AND (name ILIKE $%d OR code ILIKE $%d)", len(args), len(args))
	}

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM assets `+where, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count assets: %w", err)
	}

	page, perPage := pagination(filter.Page, filter.PerPage)
	args = append(args, perPage, (page-1)*perPage)
	query := assetSelect + " " + where + fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list assets: %w", err)
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

var assetUpdatable = map[string]columnKind{
	"plant_id":      kindText,
	"zone_id":       kindText,
	"name":          kindText,
	"status":        kindText,
	"description":   kindText,
	"manufacturer":  kindText,
	"model":         kindText,
	"serial_number": kindText,
	"weight_kg":     kindNumber,
	"dimensions":    kindJSON,
	"purchase_cost": kindNumber,
	"client_id":     kindText,
	"metadata":      kindJSON,
	"last_scan_at":  kindText,
	"last_scan_by":  kindText,
}

func (s *AssetStore) Update(tenantID, id string, patch map[string]interface{}) (*asset.Asset, error) {
	setClause, args, err := buildUpdate(assetUpdatable, patch)
	if err != nil {
		return nil, err
	}
	args = append(args, tenantID, id)
	query := fmt.Sprintf(`UPDATE assets SET %s, updated_at = NOW() WHERE tenant_id = $%d AND id = $%d RETURNING id`,
		setClause, len(args)-1, len(args))

	var updatedID string
	if err := s.db.QueryRow(query, args...).Scan(&updatedID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("asset not found: %s", id)
		}
		return nil, fmt.Errorf("update asset: %w", err)
	}
	return s.GetByID(tenantID, updatedID)
}

func (s *AssetStore) Delete(tenantID, id string) error {
	_, err := s.db.Exec(`DELETE FROM assets WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	if err != nil {
		return fmt.Errorf("delete asset: %w", err)
	}
	return nil
}

func (s *AssetStore) RecordEvent(e *asset.AssetEvent) error {
	geo, err := marshalJSON(e.Geo)
	if err != nil {
		return err
	}
	row := s.db.QueryRow(`
		INSERT INTO asset_events (tenant_id, asset_id, event_type, from_status, to_status,
			from_plant_id, to_plant_id, user_id, shipment_id, notes, geo, occurred_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11::jsonb, COALESCE(NULLIF($12, '')::timestamptz, NOW()))
		RETURNING id, created_at, occurred_at
	`, e.TenantID, e.AssetID, string(e.EventType), nullStr(string(e.FromStatus)), nullStr(string(e.ToStatus)),
		nullStr(e.FromPlantID), nullStr(e.ToPlantID), e.UserID, nullStr(e.ShipmentID), nullStr(e.Notes), geo, nullTimeStr(e.OccurredAt))

	return row.Scan(&e.ID, &e.CreatedAt, &e.OccurredAt)
}

func (s *AssetStore) ListEvents(tenantID, assetID string) ([]*asset.AssetEvent, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, asset_id, event_type, from_status, to_status, from_plant_id, to_plant_id,
			user_id, shipment_id, notes, geo, occurred_at, created_at
		FROM asset_events WHERE tenant_id = $1 AND asset_id = $2 ORDER BY occurred_at DESC
	`, tenantID, assetID)
	if err != nil {
		return nil, fmt.Errorf("list asset events: %w", err)
	}
	defer rows.Close()

	var out []*asset.AssetEvent
	for rows.Next() {
		var e asset.AssetEvent
		var eventType, fromStatus, toStatus, fromPlantID, toPlantID, shipmentID, notes sql.NullString
		var geo []byte
		if err := rows.Scan(&e.ID, &e.TenantID, &e.AssetID, &eventType, &fromStatus, &toStatus, &fromPlantID, &toPlantID,
			&e.UserID, &shipmentID, &notes, &geo, &e.OccurredAt, &e.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan asset event: %w", err)
		}
		e.EventType = asset.EventType(eventType.String)
		e.FromStatus, e.ToStatus = asset.AssetStatus(fromStatus.String), asset.AssetStatus(toStatus.String)
		e.FromPlantID, e.ToPlantID, e.ShipmentID, e.Notes = fromPlantID.String, toPlantID.String, shipmentID.String, notes.String
		if len(geo) > 0 {
			var g asset.GeoPoint
			if err := unmarshalJSON(geo, &g); err != nil {
				return nil, fmt.Errorf("decode event geo: %w", err)
			}
			e.Geo = &g
		}
		out = append(out, &e)
	}
	return out, rows.Err()
}

func scanAsset(row rowScanner) (*asset.Asset, error) {
	var a asset.Asset
	var typ, status string
	var zoneID, name, airtableID, barcode, description, manufacturer, model, serialNumber, clientID, lastScanBy sql.NullString
	var purchaseDate sql.NullTime
	var purchaseCost sql.NullFloat64
	var lastScanAt sql.NullTime
	var dimensions, metadata []byte

	if err := row.Scan(&a.ID, &a.TenantID, &a.PlantID, &zoneID, &name, &a.Code, &airtableID, &a.QRCode, &barcode, &typ, &status,
		&description, &manufacturer, &model, &serialNumber, &a.WeightKg, &dimensions, &purchaseDate,
		&purchaseCost, &clientID, &metadata, &lastScanAt, &lastScanBy, &a.CreatedAt, &a.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("asset not found")
		}
		return nil, fmt.Errorf("scan asset: %w", err)
	}

	a.Type, a.Status = asset.AssetType(typ), asset.AssetStatus(status)
	a.ZoneID, a.Name, a.AirtableID, a.Barcode = strOrEmpty(zoneID), strOrEmpty(name), strOrEmpty(airtableID), strOrEmpty(barcode)
	a.Description, a.Manufacturer, a.Model = strOrEmpty(description), strOrEmpty(manufacturer), strOrEmpty(model)
	a.SerialNumber, a.ClientID, a.LastScanBy = strOrEmpty(serialNumber), strOrEmpty(clientID), strOrEmpty(lastScanBy)
	a.LastScanAt = timeOrNil(lastScanAt)
	if purchaseDate.Valid {
		a.PurchaseDate = timeOrNil(purchaseDate)
	}
	if purchaseCost.Valid {
		a.PurchaseCost = purchaseCost.Float64
	}
	if err := unmarshalJSON(dimensions, &a.Dimensions); err != nil {
		return nil, fmt.Errorf("decode asset dimensions: %w", err)
	}
	if len(metadata) > 0 {
		if err := unmarshalJSON(metadata, &a.Metadata); err != nil {
			return nil, fmt.Errorf("decode asset metadata: %w", err)
		}
	}
	return &a, nil
}
