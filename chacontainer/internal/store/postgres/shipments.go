package postgres

import (
	"database/sql"
	"fmt"

	"github.com/cli/cli/v2/chacontainer/internal/api/handlers"
	"github.com/cli/cli/v2/chacontainer/internal/domain/shipment"
)

// ShipmentStore persists shipments, their lines and status timeline in PostgreSQL.
type ShipmentStore struct {
	db *sql.DB
}

func NewShipmentStore(db *sql.DB) *ShipmentStore {
	return &ShipmentStore{db: db}
}

const shipmentSelect = `
	SELECT id, tenant_id, reference, client_id, origin_plant_id, dest_plant_id, dest_address, status, mode,
		carrier_name, carrier_ref, tracking_url, scheduled_pickup, actual_pickup, scheduled_delivery,
		actual_delivery, total_weight_kg, total_items, notes, customs_ref, erp_id, created_by, created_at, updated_at
	FROM shipments`

func (s *ShipmentStore) Create(tenantID string, sh *shipment.Shipment) error {
	destAddress, err := marshalJSON(sh.DestAddress)
	if err != nil {
		return err
	}

	row := s.db.QueryRow(`
		INSERT INTO shipments (tenant_id, reference, client_id, origin_plant_id, dest_plant_id, dest_address,
			status, mode, carrier_name, carrier_ref, tracking_url, scheduled_pickup, scheduled_delivery,
			total_weight_kg, total_items, notes, customs_ref, erp_id, created_by)
		VALUES ($1, $2, $3, $4, $5, $6::jsonb, COALESCE(NULLIF($7, ''), 'draft'), COALESCE(NULLIF($8, ''), 'road'),
			$9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id, created_at, updated_at
	`, tenantID, sh.Reference, sh.ClientID, sh.OriginPlantID, nullStr(sh.DestPlantID), destAddress,
		string(sh.Status), string(sh.Mode), nullStr(sh.CarrierName), nullStr(sh.CarrierRef), nullStr(sh.TrackingURL),
		sh.ScheduledPickup, sh.ScheduledDelivery, sh.TotalWeightKg, sh.TotalItems, nullStr(sh.Notes),
		nullStr(sh.CustomsRef), nullStr(sh.ERPID), sh.CreatedBy)

	if err := row.Scan(&sh.ID, &sh.CreatedAt, &sh.UpdatedAt); err != nil {
		return fmt.Errorf("insert shipment: %w", err)
	}
	sh.TenantID = tenantID
	if sh.Status == "" {
		sh.Status = shipment.StatusDraft
	}
	if sh.Mode == "" {
		sh.Mode = shipment.ModeRoad
	}
	return nil
}

func (s *ShipmentStore) GetByID(tenantID, id string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(shipmentSelect+` WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanShipment(row)
}

func (s *ShipmentStore) GetByReference(tenantID, ref string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(shipmentSelect+` WHERE tenant_id = $1 AND reference = $2`, tenantID, ref)
	return scanShipment(row)
}

func (s *ShipmentStore) List(tenantID string, filter handlers.ShipmentFilter) ([]*shipment.Shipment, int, error) {
	where := "WHERE tenant_id = $1"
	args := []interface{}{tenantID}

	if filter.ClientID != "" {
		args = append(args, filter.ClientID)
		where += fmt.Sprintf(" AND client_id = $%d", len(args))
	}
	if filter.PlantID != "" {
		args = append(args, filter.PlantID)
		where += fmt.Sprintf(" AND (origin_plant_id = $%d OR dest_plant_id = $%d)", len(args), len(args))
	}
	if filter.Status != "" {
		args = append(args, filter.Status)
		where += fmt.Sprintf(" AND status = $%d", len(args))
	}
	if filter.From != nil {
		args = append(args, *filter.From)
		where += fmt.Sprintf(" AND scheduled_pickup >= $%d", len(args))
	}
	if filter.To != nil {
		args = append(args, *filter.To)
		where += fmt.Sprintf(" AND scheduled_delivery <= $%d", len(args))
	}

	var total int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM shipments `+where, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count shipments: %w", err)
	}

	page, perPage := pagination(filter.Page, filter.PerPage)
	args = append(args, perPage, (page-1)*perPage)
	query := shipmentSelect + " " + where + fmt.Sprintf(" ORDER BY scheduled_delivery DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("list shipments: %w", err)
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
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}
	defer tx.Rollback()

	// Lock and read the current status first so the event we log below
	// records the real "from" state, not the one we're about to write.
	var fromStatus string
	err = tx.QueryRow(`SELECT status FROM shipments WHERE tenant_id = $1 AND id = $2 FOR UPDATE`, tenantID, id).Scan(&fromStatus)
	if err == sql.ErrNoRows {
		return fmt.Errorf("shipment not found: %s", id)
	} else if err != nil {
		return fmt.Errorf("lock shipment: %w", err)
	}

	extraSet := ""
	switch status {
	case shipment.StatusInTransit:
		extraSet = ", actual_pickup = COALESCE(actual_pickup, NOW())"
	case shipment.StatusDelivered:
		extraSet = ", actual_delivery = COALESCE(actual_delivery, NOW())"
	}
	if _, err := tx.Exec(fmt.Sprintf(`UPDATE shipments SET status = $1, updated_at = NOW()%s WHERE tenant_id = $2 AND id = $3`, extraSet),
		string(status), tenantID, id); err != nil {
		return fmt.Errorf("update shipment status: %w", err)
	}

	if description == "" {
		description = fmt.Sprintf("Status changed to %s", status)
	}
	if _, err := tx.Exec(`
		INSERT INTO shipment_events (tenant_id, shipment_id, from_status, to_status, description, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, tenantID, id, nullStr(fromStatus), string(status), description, userID); err != nil {
		return fmt.Errorf("insert shipment event: %w", err)
	}

	return tx.Commit()
}

func (s *ShipmentStore) AddLine(tenantID string, line *shipment.ShipmentLine) error {
	row := s.db.QueryRow(`
		INSERT INTO shipment_lines (tenant_id, shipment_id, asset_id, quantity, weight_kg, notes, scanned_at)
		VALUES ($1, $2, $3, COALESCE(NULLIF($4, 0), 1), $5, $6, $7)
		ON CONFLICT (shipment_id, asset_id) DO UPDATE SET quantity = shipment_lines.quantity + EXCLUDED.quantity
		RETURNING id, created_at
	`, tenantID, line.ShipmentID, line.AssetID, line.Quantity, line.WeightKg, nullStr(line.Notes), nullTime(line.ScannedAt))

	if err := row.Scan(&line.ID, &line.CreatedAt); err != nil {
		return fmt.Errorf("insert shipment line: %w", err)
	}
	line.TenantID = tenantID
	return nil
}

func (s *ShipmentStore) ListLines(tenantID, shipmentID string) ([]*shipment.ShipmentLine, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, shipment_id, asset_id, quantity, weight_kg, notes, scanned_at, created_at
		FROM shipment_lines WHERE tenant_id = $1 AND shipment_id = $2 ORDER BY created_at
	`, tenantID, shipmentID)
	if err != nil {
		return nil, fmt.Errorf("list shipment lines: %w", err)
	}
	defer rows.Close()

	var out []*shipment.ShipmentLine
	for rows.Next() {
		var l shipment.ShipmentLine
		var notes sql.NullString
		var scannedAt sql.NullTime
		if err := rows.Scan(&l.ID, &l.TenantID, &l.ShipmentID, &l.AssetID, &l.Quantity, &l.WeightKg, &notes, &scannedAt, &l.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan shipment line: %w", err)
		}
		l.Notes, l.ScannedAt = strOrEmpty(notes), timeOrNil(scannedAt)
		out = append(out, &l)
	}
	return out, rows.Err()
}

func (s *ShipmentStore) ListEvents(tenantID, shipmentID string) ([]*shipment.ShipmentEvent, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, shipment_id, from_status, to_status, location, description, user_id, occurred_at
		FROM shipment_events WHERE tenant_id = $1 AND shipment_id = $2 ORDER BY occurred_at DESC
	`, tenantID, shipmentID)
	if err != nil {
		return nil, fmt.Errorf("list shipment events: %w", err)
	}
	defer rows.Close()

	var out []*shipment.ShipmentEvent
	for rows.Next() {
		var e shipment.ShipmentEvent
		var fromStatus, location sql.NullString
		var toStatus string
		if err := rows.Scan(&e.ID, &e.TenantID, &e.ShipmentID, &fromStatus, &toStatus, &location, &e.Description, &e.UserID, &e.OccurredAt); err != nil {
			return nil, fmt.Errorf("scan shipment event: %w", err)
		}
		e.FromStatus, e.ToStatus, e.Location = shipment.ShipmentStatus(fromStatus.String), shipment.ShipmentStatus(toStatus), strOrEmpty(location)
		out = append(out, &e)
	}
	return out, rows.Err()
}

func scanShipment(row rowScanner) (*shipment.Shipment, error) {
	var sh shipment.Shipment
	var status, mode string
	var destPlantID, carrierName, carrierRef, trackingURL, notes, customsRef, erpID sql.NullString
	var actualPickup, actualDelivery sql.NullTime
	var destAddress []byte

	if err := row.Scan(&sh.ID, &sh.TenantID, &sh.Reference, &sh.ClientID, &sh.OriginPlantID, &destPlantID, &destAddress,
		&status, &mode, &carrierName, &carrierRef, &trackingURL, &sh.ScheduledPickup, &actualPickup,
		&sh.ScheduledDelivery, &actualDelivery, &sh.TotalWeightKg, &sh.TotalItems, &notes, &customsRef, &erpID,
		&sh.CreatedBy, &sh.CreatedAt, &sh.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("shipment not found")
		}
		return nil, fmt.Errorf("scan shipment: %w", err)
	}

	sh.Status, sh.Mode = shipment.ShipmentStatus(status), shipment.TransportMode(mode)
	sh.DestPlantID, sh.CarrierName, sh.CarrierRef = strOrEmpty(destPlantID), strOrEmpty(carrierName), strOrEmpty(carrierRef)
	sh.TrackingURL, sh.Notes, sh.CustomsRef, sh.ERPID = strOrEmpty(trackingURL), strOrEmpty(notes), strOrEmpty(customsRef), strOrEmpty(erpID)
	sh.ActualPickup, sh.ActualDelivery = timeOrNil(actualPickup), timeOrNil(actualDelivery)
	if len(destAddress) > 0 {
		var a shipment.Address
		if err := unmarshalJSON(destAddress, &a); err != nil {
			return nil, fmt.Errorf("decode dest address: %w", err)
		}
		sh.DestAddress = &a
	}
	return &sh, nil
}
