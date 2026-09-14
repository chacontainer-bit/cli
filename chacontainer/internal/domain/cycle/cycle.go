package cycle

import "time"

type Status string

const (
	StatusOpen         Status = "open"
	StatusClosed       Status = "closed"
	StatusException    Status = "exception"
	StatusNoConciliado Status = "no_conciliado"
)

// Cycle is one closed reverse-logistics loop for an asset: dispatch,
// client use, return transit, inspection and release. Day-metric fields
// are computed and stored by the application as each milestone lands and
// when the cycle closes.
type Cycle struct {
	ID                  string     `json:"id" db:"id"`
	TenantID            string     `json:"tenant_id" db:"tenant_id"`
	PilotConfigID       string     `json:"pilot_config_id" db:"pilot_config_id"`
	AssetID             string     `json:"asset_id" db:"asset_id"`
	DispatchedAt        *time.Time `json:"dispatched_at,omitempty" db:"dispatched_at"`
	ReceivedByClientAt  *time.Time `json:"received_by_client_at,omitempty" db:"received_by_client_at"`
	EmptiedAt           *time.Time `json:"emptied_at,omitempty" db:"emptied_at"`
	ReadyForReturnAt    *time.Time `json:"ready_for_return_at,omitempty" db:"ready_for_return_at"`
	CollectedAt         *time.Time `json:"collected_at,omitempty" db:"collected_at"`
	ReceivedReturnAt    *time.Time `json:"received_return_at,omitempty" db:"received_return_at"`
	InspectionStartedAt *time.Time `json:"inspection_started_at,omitempty" db:"inspection_started_at"`
	ReleasedAt          *time.Time `json:"released_at,omitempty" db:"released_at"`
	OutboundTransitDays *float64   `json:"outbound_transit_days,omitempty" db:"outbound_transit_days"`
	ClientUseDays       *float64   `json:"client_use_days,omitempty" db:"client_use_days"`
	DwellReturnDays     *float64   `json:"dwell_return_days,omitempty" db:"dwell_return_days"`
	ReturnTransitDays   *float64   `json:"return_transit_days,omitempty" db:"return_transit_days"`
	ReconditioningDays  *float64   `json:"reconditioning_days,omitempty" db:"reconditioning_days"`
	CycleTimeDays       *float64   `json:"cycle_time_days,omitempty" db:"cycle_time_days"`
	SLACycleDays        *float64   `json:"sla_cycle_days,omitempty" db:"sla_cycle_days"`
	MeetsSLA            *bool      `json:"meets_sla,omitempty" db:"meets_sla"`
	DwellCriticalHours  *float64   `json:"dwell_critical_hours,omitempty" db:"dwell_critical_hours"`
	ExceedsDwell        *bool      `json:"exceeds_dwell,omitempty" db:"exceeds_dwell"`
	Status              Status     `json:"status" db:"status"`
	Notes               string     `json:"notes,omitempty" db:"notes"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}

// Milestone identifies which timestamp of the cycle is being recorded.
type Milestone string

const (
	MilestoneDispatched       Milestone = "dispatched_at"
	MilestoneReceivedByClient Milestone = "received_by_client_at"
	MilestoneEmptied          Milestone = "emptied_at"
	MilestoneReadyForReturn   Milestone = "ready_for_return_at"
	MilestoneCollected        Milestone = "collected_at"
	MilestoneReceivedReturn   Milestone = "received_return_at"
	MilestoneInspectionStart  Milestone = "inspection_started_at"
	MilestoneReleased         Milestone = "released_at"
)

// ApplyMilestone sets the timestamp for a milestone and recomputes every
// day-metric that milestone unlocks. Mirrors the CICLOS sheet's derived
// columns (Outbound_Transit_Dias, Uso_Cliente_Dias, ...).
func (c *Cycle) ApplyMilestone(m Milestone, at time.Time) {
	switch m {
	case MilestoneDispatched:
		c.DispatchedAt = &at
	case MilestoneReceivedByClient:
		c.ReceivedByClientAt = &at
	case MilestoneEmptied:
		c.EmptiedAt = &at
	case MilestoneReadyForReturn:
		c.ReadyForReturnAt = &at
	case MilestoneCollected:
		c.CollectedAt = &at
	case MilestoneReceivedReturn:
		c.ReceivedReturnAt = &at
	case MilestoneInspectionStart:
		c.InspectionStartedAt = &at
	case MilestoneReleased:
		c.ReleasedAt = &at
	}
	c.recomputeMetrics()
}

func (c *Cycle) recomputeMetrics() {
	c.OutboundTransitDays = daysBetween(c.DispatchedAt, c.ReceivedByClientAt)
	c.ClientUseDays = daysBetween(c.ReceivedByClientAt, c.EmptiedAt)
	c.DwellReturnDays = daysBetween(c.EmptiedAt, c.CollectedAt)
	c.ReturnTransitDays = daysBetween(c.CollectedAt, c.ReceivedReturnAt)
	c.ReconditioningDays = daysBetween(c.ReceivedReturnAt, c.ReleasedAt)
	c.CycleTimeDays = daysBetween(c.DispatchedAt, c.ReleasedAt)

	if c.CycleTimeDays != nil && c.SLACycleDays != nil {
		meets := *c.CycleTimeDays <= *c.SLACycleDays
		c.MeetsSLA = &meets
	}
	if c.DwellReturnDays != nil && c.DwellCriticalHours != nil {
		exceeds := (*c.DwellReturnDays * 24) > *c.DwellCriticalHours
		c.ExceedsDwell = &exceeds
	}
	if c.ReleasedAt != nil && c.Status == StatusOpen {
		c.Status = StatusClosed
	}
}

func daysBetween(from, to *time.Time) *float64 {
	if from == nil || to == nil {
		return nil
	}
	d := to.Sub(*from).Hours() / 24
	return &d
}
