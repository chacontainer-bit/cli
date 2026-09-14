// Package recovery implements the ERT pilot's financial ledger and ROI
// gate: the value recovery entries behind 06_VALUE_RECOVERY and the
// aggregated dashboard behind 07_ROI in the CHC-ERT-WBK-LI-001 workbook.
package recovery

import "time"

// Classification is the ROI-eligibility bucket for a benefit
// (08_CATALOGOS: CLASIFICACION_FINANCIERA). Only AhorroRealizado and
// CostoEvitadoVerificado, once validated, may count toward ROI — see
// PilotConfig.ROIPolicy and Entry.IncludeInROI.
type Classification string

const (
	AhorroRealizado                 Classification = "AHORRO_REALIZADO"
	CostoEvitadoVerificado          Classification = "COSTO_EVITADO_VERIFICADO"
	ExposicionPatrimonialRecuperada Classification = "EXPOSICION_PATRIMONIAL_RECUPERADA"
	OportunidadPotencial            Classification = "OPORTUNIDAD_POTENCIAL"
)

// ROIEligible reports whether this classification may ever count toward
// ROI, independent of validation state.
func (c Classification) ROIEligible() bool {
	return c == AhorroRealizado || c == CostoEvitadoVerificado
}

type ValidationStatus string

const (
	ValidationPending  ValidationStatus = "pendiente"
	ValidationApproved ValidationStatus = "validado"
	ValidationRejected ValidationStatus = "rechazado"
)

// Entry is one line of the value recovery ledger.
type Entry struct {
	ID                      string           `json:"id" db:"id"`
	TenantID                string           `json:"tenant_id" db:"tenant_id"`
	PilotConfigID           string           `json:"pilot_config_id" db:"pilot_config_id"`
	CycleID                 string           `json:"cycle_id,omitempty" db:"cycle_id"`
	AssetID                 string           `json:"asset_id,omitempty" db:"asset_id"`
	EntryDate               time.Time        `json:"entry_date" db:"entry_date"`
	Category                string           `json:"category" db:"category"`
	FinancialClassification Classification   `json:"financial_classification" db:"financial_classification"`
	Quantity                int              `json:"quantity" db:"quantity"`
	BaseCostUnitMXN         float64          `json:"base_cost_unit_mxn" db:"base_cost_unit_mxn"`
	ActualCostUnitMXN       float64          `json:"actual_cost_unit_mxn" db:"actual_cost_unit_mxn"`
	RecoveredValueUnitMXN   float64          `json:"recovered_value_unit_mxn" db:"recovered_value_unit_mxn"`
	GrossBenefitMXN         float64          `json:"gross_benefit_mxn" db:"gross_benefit_mxn"`
	ImplementationCostMXN   float64          `json:"implementation_cost_mxn" db:"implementation_cost_mxn"`
	NetBenefitMXN           float64          `json:"net_benefit_mxn" db:"net_benefit_mxn"`
	IncludeInROI            bool             `json:"include_in_roi" db:"include_in_roi"`
	EvidenceURL             string           `json:"evidence_url,omitempty" db:"evidence_url"`
	ValidatorID             string           `json:"validator_id,omitempty" db:"validator_id"`
	ValidationStatus        ValidationStatus `json:"validation_status" db:"validation_status"`
	Comments                string           `json:"comments,omitempty" db:"comments"`
	CreatedAt               time.Time        `json:"created_at" db:"created_at"`
}

// EligibleForROI applies the workbook's gate rule directly: an entry only
// counts toward ROI once validated AND classified as realized savings or
// verified avoided cost. Exposure recovered and potential opportunity are
// always excluded to avoid double counting.
func (e *Entry) EligibleForROI() bool {
	return e.ValidationStatus == ValidationApproved && e.FinancialClassification.ROIEligible()
}

// PilotConfig is the governance and baseline data for one circuit under
// pilot (01_CONFIG sheet).
type PilotConfig struct {
	ID                        string     `json:"id" db:"id"`
	TenantID                  string     `json:"tenant_id" db:"tenant_id"`
	Code                      string     `json:"code" db:"code"`
	ClientID                  string     `json:"client_id" db:"client_id"`
	Circuit                   string     `json:"circuit" db:"circuit"`
	Currency                  string     `json:"currency" db:"currency"`
	StartDate                 *time.Time `json:"start_date,omitempty" db:"start_date"`
	EndDate                   *time.Time `json:"end_date,omitempty" db:"end_date"`
	SampleAssetsTarget        int        `json:"sample_assets_target" db:"sample_assets_target"`
	CircuitPoolTotal          int        `json:"circuit_pool_total" db:"circuit_pool_total"`
	DailyDemand               *int       `json:"daily_demand,omitempty" db:"daily_demand"`
	ReplacementCostAvgMXN     *float64   `json:"replacement_cost_avg_mxn,omitempty" db:"replacement_cost_avg_mxn"`
	CycleTimeBaselineDays     *float64   `json:"cycle_time_baseline_days,omitempty" db:"cycle_time_baseline_days"`
	SLACycleDays              *float64   `json:"sla_cycle_days,omitempty" db:"sla_cycle_days"`
	DwellCriticalHours        *float64   `json:"dwell_critical_hours,omitempty" db:"dwell_critical_hours"`
	LossRateBaselinePct       *float64   `json:"loss_rate_baseline_pct,omitempty" db:"loss_rate_baseline_pct"`
	ReverseFreightBaselineMXN *float64   `json:"reverse_freight_baseline_mxn,omitempty" db:"reverse_freight_baseline_mxn"`
	FixedCostMXN              float64    `json:"fixed_cost_mxn" db:"fixed_cost_mxn"`
	TraceabilityCostMXN       float64    `json:"traceability_cost_mxn" db:"traceability_cost_mxn"`
	TransportCostMXN          float64    `json:"transport_cost_mxn" db:"transport_cost_mxn"`
	OtherCostMXN              float64    `json:"other_cost_mxn" db:"other_cost_mxn"`
	ROIPolicy                 string     `json:"roi_policy" db:"roi_policy"`
	Status                    string     `json:"status" db:"status"`
	CreatedAt                 time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at" db:"updated_at"`
}

// TotalPilotCostMXN sums the pilot's own cost buckets — never client
// replacement cost or asset value, matching 01_CONFIG's "Costo total
// piloto" formula.
func (p PilotConfig) TotalPilotCostMXN() float64 {
	return p.FixedCostMXN + p.TraceabilityCostMXN + p.TransportCostMXN + p.OtherCostMXN
}

// Summary is the computed ROI dashboard (07_ROI sheet): KPIs, the
// benefit actually eligible for ROI, and the decision gate criteria.
type Summary struct {
	PilotConfigID            string       `json:"pilot_config_id"`
	AssetsTarget             int          `json:"assets_target"`
	ClosedCycles             int          `json:"closed_cycles"`
	AvgCycleTimeDays         float64      `json:"avg_cycle_time_days"`
	AvgDwellReturnDays       float64      `json:"avg_dwell_return_days"`
	PctCyclesWithinSLA       float64      `json:"pct_cycles_within_sla"`
	UnreconciledAssets       int          `json:"unreconciled_assets"`
	AvailabilityGap          int          `json:"availability_gap"`
	TotalPilotCostMXN        float64      `json:"total_pilot_cost_mxn"`
	BenefitIncludedInROIMXN  float64      `json:"benefit_included_in_roi_mxn"`
	TotalEconomicObservedMXN float64      `json:"total_economic_observed_mxn"`
	ROI                      *float64     `json:"roi,omitempty"`
	PaybackMultiple          *float64     `json:"payback_multiple,omitempty"`
	Gate                     DecisionGate `json:"gate"`
	GeneratedAt              time.Time    `json:"generated_at"`
}

// DecisionGate mirrors the ROI sheet's "GATE DE DECISIÓN" block: the pilot
// may not be scaled until every criterion is met.
type DecisionGate struct {
	TraceabilityComplete bool `json:"traceability_complete"` // >=95% mandatory events captured
	CycleTimeImproved    bool `json:"cycle_time_improved"`   // demonstrable vs. baseline
	ValueProven          bool `json:"value_proven"`          // verifiable benefit > pilot cost
	ExceptionsControlled bool `json:"exceptions_controlled"` // every exception has an owner and is closed
}

// ReadyToScale is true only once all four gate criteria pass — no single
// strong metric may substitute for the others.
func (g DecisionGate) ReadyToScale() bool {
	return g.TraceabilityComplete && g.CycleTimeImproved && g.ValueProven && g.ExceptionsControlled
}
