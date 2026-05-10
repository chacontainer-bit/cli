package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/chacontainer/backend/internal/api/handlers"
	"github.com/chacontainer/backend/internal/domain/shipment"
)

// ShipmentStore implements handlers.ShipmentStore against PostgreSQL.
type ShipmentStore struct{ db *sql.DB }

func NewShipmentStore(db *sql.DB) *ShipmentStore { return &ShipmentStore{db: db} }

// ── Create ────────────────────────────────────────────────────────────────────

func (s *ShipmentStore) Create(tenantID string, sh *shipment.Shipment) error {
	sh.ID = newUUID()
	sh.TenantID = tenantID
	sh.CreatedAt = time.Now()
	sh.UpdatedAt = time.Now()
	if sh.Status == "" {
		sh.Status = shipment.StatusDraft
	}

	destAddrJSON := marshalJSON(sh.DestAddress)

	_, err := s.db.Exec(`
		INSERT INTO shipments (
			id, tenant_id, reference, client_id,
			origin_plant_id, dest_plant_id, dest_address,
			status, mode, carrier_name, carrier_ref, tracking_url,
			scheduled_pickup, actual_pickup,
			scheduled_delivery, actual_delivery,
			total_weight_kg, total_items, notes, customs_ref, erp_id,
			created_by, created_at, updated_at
		) VALUES (
			$1,$2,$3,$4,
			$5,$6,$7,
			$8,$9,$10,$11,$12,
			$13,$14,
			$15,$16,
			$17,$18,$19,$20,$21,
			$22,$23,$24
		)`,
		sh.ID, tenantID, sh.Reference, sh.ClientID,
		sh.OriginPlantID, nullString(sh.DestPlantID), destAddrJSON,
		string(sh.Status), string(sh.Mode),
		nullString(sh.CarrierName), nullString(sh.CarrierRef), nullString(sh.TrackingURL),
		sh.ScheduledPickup, nullTime(sh.ActualPickup),
		sh.ScheduledDelivery, nullTime(sh.ActualDelivery),
		sh.TotalWeightKg, sh.TotalItems,
		nullString(sh.Notes), nullString(sh.CustomsRef), nullString(sh.ERPID),
		sh.CreatedBy, sh.CreatedAt, sh.UpdatedAt,
	)
	return err
}

// ── GetByID ───────────────────────────────────────────────────────────────────

func (s *ShipmentStore) GetByID(tenantID, id string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(`SELECT `+shipmentCols+` FROM shipments
		WHERE tenant_id = $1 AND id = $2`, tenantID, id)
	return scanShipment(row)
}

// ── GetByReference ────────────────────────────────────────────────────────────

func (s *ShipmentStore) GetByReference(tenantID, ref string) (*shipment.Shipment, error) {
	row := s.db.QueryRow(`SELECT `+shipmentCols+` FROM shipments
		WHERE tenant_id = $1 AND reference = $2`, tenantID, ref)
	return scanShipment(row)
}

// ── List ──────────────────────────────────────────────────────────────────────

func (s *ShipmentStore) List(tenantID string, f handlers.ShipmentFilter) ([]*shipment.Shipment, int, error) {
	where := []string{"tenant_id = $1"}
	args := []interface{}{tenantID}
	n := 2

	if f.ClientID != "" {
		where = append(where, fmt.Sprintf("client_id = $%d", n))
		args = append(args, f.ClientID)
		n++
	}
	if f.PlantID != "" {
		where = append(where, fmt.Sprintf("origin_plant_id = $%d", n))
		args = append(args, f.PlantID)
		n++
	}
	if f.Status != "" {
		where = append(where, fmt.Sprintf("status = $%d", n))
		args = append(args, f.Status)
		n++
	}
	if f.From != nil {
		where = append(where, fmt.Sprintf("scheduled_delivery >= $%d", n))
		args = append(args, *f.From)
		n++
	}
	if f.To != nil {
		where = append(where, fmt.Sprintf("scheduled_delivery <= $%d", n))
		args = append(args, *f.To)
		n++
	}

	clause := "WHERE " + strings.Join(where, " AND ")

	var total int
	if err := s.db.QueryRow("SELECT COUNT(*) FROM shipments "+clause, args...).Scan(&total); err != nil {
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
		`SELECT `+shipmentCols+` FROM shipments `+clause+
			fmt.Sprintf(` ORDER BY created_at DESC LIMIT $%d OFFSET $%d`, n, n+1),
		args...,
	)
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

// ── UpdateStatus ──────────────────────────────────────────────────────────────

func (s *ShipmentStore) UpdateStatus(tenantID, id string, status shipment.ShipmentStatus, userID, description string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Fetch current status for event record
	var fromStatus string
	tx.QueryRow(`SELECT status FROM shipments WHERE tenant_id = $1 AND id = $2`,
		tenantID, id).Scan(&fromStatus)

	// Set actual timestamps on terminal transitions
	var extraSet string
	if status == shipment.StatusDelivered {
		extraSet = ", actual_delivery = NOW()"
	} else if status == shipment.StatusInTransit {
		extraSet = ", actual_pickup = NOW()"
	}

	_, err = tx.Exec(
		`UPDATE shipments SET status = $1, updated_at = NOW()`+extraSet+
			` WHERE tenant_id = $2 AND id = $3`,
		string(status), tenantID, id,
	)
	if err != nil {
		return err
	}

	// Append shipment event
	_, err = tx.Exec(`
		INSERT INTO shipment_events (id, tenant_id, shipment_id, from_status, to_status, description, user_id, occurred_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())`,
		newUUID(), tenantID, id, fromStatus, string(status), description, userID,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// ── AddLine ───────────────────────────────────────────────────────────────────

func (s *ShipmentStore) AddLine(tenantID string, line *shipment.ShipmentLine) error {
	line.ID = newUUID()
	line.TenantID = tenantID
	line.CreatedAt = time.Now()

	_, err := s.db.Exec(`
		INSERT INTO shipment_lines (id, tenant_id, shipment_id, asset_id, quantity, weight_kg, notes, created_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (shipment_id, asset_id) DO UPDATE
		  SET quantity = EXCLUDED.quantity, weight_kg = EXCLUDED.weight_kg, notes = EXCLUDED.notes`,
		line.ID, tenantID, line.ShipmentID, line.AssetID,
		line.Quantity, line.WeightKg, nullString(line.Notes), line.CreatedAt,
	)
	return err
}

// ── ListLines ─────────────────────────────────────────────────────────────────

func (s *ShipmentStore) ListLines(tenantID, shipmentID string) ([]*shipment.ShipmentLine, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, shipment_id, asset_id::text,
		       quantity, weight_kg, COALESCE(notes,''), scanned_at, created_at
		FROM shipment_lines
		WHERE tenant_id = $1 AND shipment_id = $2
		ORDER BY created_at`, tenantID, shipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*shipment.ShipmentLine
	for rows.Next() {
		l := &shipment.ShipmentLine{}
		var scannedAt sql.NullTime
		if err := rows.Scan(&l.ID, &l.TenantID, &l.ShipmentID, &l.AssetID,
			&l.Quantity, &l.WeightKg, &l.Notes, &scannedAt, &l.CreatedAt); err != nil {
			return nil, err
		}
		if scannedAt.Valid {
			t := scannedAt.Time
			l.ScannedAt = &t
		}
		out = append(out, l)
	}
	return out, rows.Err()
}

// ── ListEvents ────────────────────────────────────────────────────────────────

func (s *ShipmentStore) ListEvents(tenantID, shipmentID string) ([]*shipment.ShipmentEvent, error) {
	rows, err := s.db.Query(`
		SELECT id, tenant_id, shipment_id,
		       COALESCE(from_status,''), to_status,
		       COALESCE(location,''), description, user_id::text, occurred_at
		FROM shipment_events
		WHERE tenant_id = $1 AND shipment_id = $2
		ORDER BY occurred_at DESC`, tenantID, shipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []*shipment.ShipmentEvent
	for rows.Next() {
		e := &shipment.ShipmentEvent{}
		var fromStatus, toStatus string
		if err := rows.Scan(&e.ID, &e.TenantID, &e.ShipmentID,
			&fromStatus, &toStatus,
			&e.Location, &e.Description, &e.UserID, &e.OccurredAt); err != nil {
			return nil, err
		}
		e.FromStatus = shipment.ShipmentStatus(fromStatus)
		e.ToStatus = shipment.ShipmentStatus(toStatus)
		out = append(out, e)
	}
	return out, rows.Err()
}

// ── SQL helpers ───────────────────────────────────────────────────────────────

const shipmentCols = `
	id, tenant_id, reference, client_id::text,
	origin_plant_id::text, COALESCE(dest_plant_id::text,''), dest_address,
	status, mode,
	COALESCE(carrier_name,''), COALESCE(carrier_ref,''), COALESCE(tracking_url,''),
	scheduled_pickup, actual_pickup,
	scheduled_delivery, actual_delivery,
	total_weight_kg, total_items,
	COALESCE(notes,''), COALESCE(customs_ref,''), COALESCE(erp_id,''),
	created_by::text, created_at, updated_at`

type shipmentScanner interface {
	Scan(dest ...interface{}) error
}

func scanShipment(row shipmentScanner) (*shipment.Shipment, error) {
	sh := &shipment.Shipment{}
	var status, mode string
	var destAddrJSON []byte
	var actualPickup, actualDelivery sql.NullTime

	err := row.Scan(
		&sh.ID, &sh.TenantID, &sh.Reference, &sh.ClientID,
		&sh.OriginPlantID, &sh.DestPlantID, &destAddrJSON,
		&status, &mode,
		&sh.CarrierName, &sh.CarrierRef, &sh.TrackingURL,
		&sh.ScheduledPickup, &actualPickup,
		&sh.ScheduledDelivery, &actualDelivery,
		&sh.TotalWeightKg, &sh.TotalItems,
		&sh.Notes, &sh.CustomsRef, &sh.ERPID,
		&sh.CreatedBy, &sh.CreatedAt, &sh.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("shipment not found")
	}
	if err != nil {
		return nil, err
	}

	sh.Status = shipment.ShipmentStatus(status)
	sh.Mode = shipment.TransportMode(mode)

	if len(destAddrJSON) > 0 {
		sh.DestAddress = &shipment.Address{}
		scanJSON(destAddrJSON, sh.DestAddress)
	}
	if actualPickup.Valid {
		t := actualPickup.Time
		sh.ActualPickup = &t
	}
	if actualDelivery.Valid {
		t := actualDelivery.Time
		sh.ActualDelivery = &t
	}
	return sh, nil
}

var _ handlers.ShipmentStore = (*ShipmentStore)(nil)
