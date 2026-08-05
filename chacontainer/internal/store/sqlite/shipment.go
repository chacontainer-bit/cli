package sqlite

import (
	"database/sql"
	"strings"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/shipment"
)

type ShipmentStore struct {
	db *sql.DB
}

func NewShipmentStore(db *sql.DB) *ShipmentStore {
	return &ShipmentStore{db: db}
}

func (s *ShipmentStore) Create(tenantID string, sh *shipment.Shipment) error {
	sh.ID = newID()
	sh.TenantID = tenantID
	now := parseTime(nowStr())
	sh.CreatedAt = now
	sh.UpdatedAt = now
	if sh.Status == "" {
		sh.Status = shipment.StatusDraft
	}
	if sh.Mode == "" {
		sh.Mode = shipment.ModeRoad
	}
	if sh.Reference == "" {
		sh.Reference = "SHP-" + sh.ID[:8]
	}
	_, err := s.db.Exec(
		`INSERT INTO shipments (id, tenant_id, reference, client_id, origin_plant_id, dest_plant_id, dest_address,
			status, mode, carrier_name, carrier_ref, tracking_url, scheduled_pickup, actual_pickup,
			scheduled_delivery, actual_delivery, total_weight_kg, total_items, notes, customs_ref, erp_id,
			created_by, created_at, updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		sh.ID, sh.TenantID, sh.Reference, sh.ClientID, sh.OriginPlantID, nullStr(sh.DestPlantID), nullAddrJSON(sh.DestAddress),
		string(sh.Status), string(sh.Mode), nullStr(sh.CarrierName), nullStr(sh.CarrierRef), nullStr(sh.TrackingURL),
		fmtTime(sh.ScheduledPickup), nullTimePtr(sh.ActualPickup), fmtTime(sh.ScheduledDelivery), nullTimePtr(sh.ActualDelivery),
		sh.TotalWeightKg, sh.TotalItems, nullStr(sh.Notes), nullStr(sh.CustomsRef), nullStr(sh.ERPID),
		sh.CreatedBy, fmtTime(sh.CreatedAt), fmtTime(sh.UpdatedAt),
	)
	if err != nil {
		return err
	}
	return s.addEvent(tenantID, sh.ID, "", sh.Status, "", "shipment created", sh.CreatedBy)
}

const shipmentSelectCols = `id, tenant_id, reference, client_id, origin_plant_id, dest_plant_id, dest_address,
	status, mode, carrier_name, carrier_ref, tracking_url, scheduled_pickup, actual_pickup,
	scheduled_delivery, actual_delivery, total_weight_kg, total_items, notes, customs_ref, erp_id,
	created_by, created_at, updated_at`

func scanShipment(row interface{ Scan(dest ...interface{}) error }) (*shipment.Shipment, error) {
	var sh shipment.Shipment
	var destPlant, destAddress, carrierName, carrierRef, trackingURL, notes, customsRef, erpID sql.NullString
	var actualPickup, actualDelivery sql.NullString
	var statusStr, modeStr, scheduledPickup, scheduledDelivery, createdAt, updatedAt string

	err := row.Scan(&sh.ID, &sh.TenantID, &sh.Reference, &sh.ClientID, &sh.OriginPlantID, &destPlant, &destAddress,
		&statusStr, &modeStr, &carrierName, &carrierRef, &trackingURL, &scheduledPickup, &actualPickup,
		&scheduledDelivery, &actualDelivery, &sh.TotalWeightKg, &sh.TotalItems, &notes, &customsRef, &erpID,
		&sh.CreatedBy, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	sh.DestPlantID = destPlant.String
	sh.Status = shipment.ShipmentStatus(statusStr)
	sh.Mode = shipment.TransportMode(modeStr)
	sh.CarrierName = carrierName.String
	sh.CarrierRef = carrierRef.String
	sh.TrackingURL = trackingURL.String
	sh.Notes = notes.String
	sh.CustomsRef = customsRef.String
	sh.ERPID = erpID.String
	if destAddress.Valid && destAddress.String != "" {
		var addr shipment.Address
		fromJSON(destAddress.String, &addr)
		sh.DestAddress = &addr
	}
	sh.ScheduledPickup = parseTime(scheduledPickup)
	sh.ScheduledDelivery = parseTime(scheduledDelivery)
	sh.ActualPickup = parseTimePtr(actualPickup)
	sh.ActualDelivery = parseTimePtr(actualDelivery)
	sh.CreatedAt = parseTime(createdAt)
	sh.UpdatedAt = parseTime(updatedAt)
	return &sh, nil
}

func (s *ShipmentStore) GetByID(tenantID, id string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(`SELECT `+shipmentSelectCols+` FROM shipments WHERE tenant_id = ? AND id = ?`, tenantID, id)
	return scanShipment(row)
}

func (s *ShipmentStore) GetByReference(tenantID, ref string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(`SELECT `+shipmentSelectCols+` FROM shipments WHERE tenant_id = ? AND reference = ?`, tenantID, ref)
	return scanShipment(row)
}

func (s *ShipmentStore) List(tenantID string, filter handlers.ShipmentFilter) ([]*shipment.Shipment, int, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{tenantID}
	if filter.ClientID != "" {
		where = append(where, "client_id = ?")
		args = append(args, filter.ClientID)
	}
	if filter.PlantID != "" {
		where = append(where, "(origin_plant_id = ? OR dest_plant_id = ?)")
		args = append(args, filter.PlantID, filter.PlantID)
	}
	if filter.Status != "" {
		where = append(where, "status = ?")
		args = append(args, filter.Status)
	}
	whereClause := strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM shipments WHERE `+whereClause, args...).Scan(&total); err != nil {
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

	rows, err := s.db.Query(`SELECT `+shipmentSelectCols+` FROM shipments WHERE `+whereClause+` ORDER BY created_at DESC LIMIT ? OFFSET ?`, queryArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []*shipment.Shipment
	for rows.Next() {
		sh, err := scanShipment(rows)
		if err != nil {
			return nil, 0, err
		}
		out = append(out, sh)
	}
	return out, total, rows.Err()
}

func (s *ShipmentStore) UpdateStatus(tenantID, id string, status shipment.ShipmentStatus, userID, description string) error {
	current, err := s.GetByID(tenantID, id)
	if err != nil {
		return err
	}
	sets := []string{"status = ?", "updated_at = ?"}
	args := []interface{}{string(status), nowStr()}
	if status == shipment.StatusDelivered {
		sets = append(sets, "actual_delivery = ?")
		args = append(args, nowStr())
	}
	if status == shipment.StatusInTransit && current.ActualPickup == nil {
		sets = append(sets, "actual_pickup = ?")
		args = append(args, nowStr())
	}
	args = append(args, tenantID, id)
	if _, err := s.db.Exec(`UPDATE shipments SET `+strings.Join(sets, ", ")+` WHERE tenant_id = ? AND id = ?`, args...); err != nil {
		return err
	}
	if description == "" {
		description = "status changed to " + string(status)
	}
	return s.addEvent(tenantID, id, current.Status, status, "", description, userID)
}

func (s *ShipmentStore) addEvent(tenantID, shipmentID string, from, to shipment.ShipmentStatus, location, description, userID string) error {
	_, err := s.db.Exec(
		`INSERT INTO shipment_events (id, tenant_id, shipment_id, from_status, to_status, location, description, user_id, occurred_at)
		 VALUES (?,?,?,?,?,?,?,?,?)`,
		newID(), tenantID, shipmentID, nullStr(string(from)), string(to), nullStr(location), description, userID, nowStr(),
	)
	return err
}

func (s *ShipmentStore) AddLine(tenantID string, line *shipment.ShipmentLine) error {
	line.ID = newID()
	line.TenantID = tenantID
	line.CreatedAt = parseTime(nowStr())
	if line.Quantity == 0 {
		line.Quantity = 1
	}
	_, err := s.db.Exec(
		`INSERT INTO shipment_lines (id, tenant_id, shipment_id, asset_id, quantity, weight_kg, notes, scanned_at, created_at)
		 VALUES (?,?,?,?,?,?,?,?,?)`,
		line.ID, line.TenantID, line.ShipmentID, line.AssetID, line.Quantity, line.WeightKg, nullStr(line.Notes),
		nullTimePtr(line.ScannedAt), fmtTime(line.CreatedAt),
	)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(
		`UPDATE shipments SET total_items = total_items + ?, total_weight_kg = total_weight_kg + ? WHERE tenant_id = ? AND id = ?`,
		line.Quantity, line.WeightKg, tenantID, line.ShipmentID,
	)
	return err
}

func (s *ShipmentStore) ListLines(tenantID, shipmentID string) ([]*shipment.ShipmentLine, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, shipment_id, asset_id, quantity, weight_kg, notes, scanned_at, created_at
		 FROM shipment_lines WHERE tenant_id = ? AND shipment_id = ? ORDER BY created_at`, tenantID, shipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*shipment.ShipmentLine
	for rows.Next() {
		var l shipment.ShipmentLine
		var notes, scannedAt sql.NullString
		var createdAt string
		if err := rows.Scan(&l.ID, &l.TenantID, &l.ShipmentID, &l.AssetID, &l.Quantity, &l.WeightKg, &notes, &scannedAt, &createdAt); err != nil {
			return nil, err
		}
		l.Notes = notes.String
		l.ScannedAt = parseTimePtr(scannedAt)
		l.CreatedAt = parseTime(createdAt)
		out = append(out, &l)
	}
	return out, rows.Err()
}

func (s *ShipmentStore) ListEvents(tenantID, shipmentID string) ([]*shipment.ShipmentEvent, error) {
	rows, err := s.db.Query(
		`SELECT id, tenant_id, shipment_id, from_status, to_status, location, description, user_id, occurred_at
		 FROM shipment_events WHERE tenant_id = ? AND shipment_id = ? ORDER BY occurred_at DESC`, tenantID, shipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*shipment.ShipmentEvent
	for rows.Next() {
		var e shipment.ShipmentEvent
		var fromStatus, location sql.NullString
		var toStatus, occurredAt string
		if err := rows.Scan(&e.ID, &e.TenantID, &e.ShipmentID, &fromStatus, &toStatus, &location, &e.Description, &e.UserID, &occurredAt); err != nil {
			return nil, err
		}
		e.FromStatus = shipment.ShipmentStatus(fromStatus.String)
		e.ToStatus = shipment.ShipmentStatus(toStatus)
		e.Location = location.String
		e.OccurredAt = parseTime(occurredAt)
		out = append(out, &e)
	}
	return out, rows.Err()
}

func nullAddrJSON(a *shipment.Address) interface{} {
	if a == nil {
		return nil
	}
	return toJSON(a)
}
