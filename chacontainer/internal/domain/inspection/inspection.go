package inspection

import "time"

// Condition is the AIAG-style grade scale used on receipt and release
// (08_CATALOGOS: CONDICION A-E).
type Condition string

const (
	ConditionA Condition = "A" // like new
	ConditionB Condition = "B" // light wear, fully functional
	ConditionC Condition = "C" // visible wear, needs cleaning
	ConditionD Condition = "D" // damaged, needs repair
	ConditionE Condition = "E" // unusable
)

type Severity string

const (
	SeverityNone     Severity = "ninguna"
	SeverityLight    Severity = "leve"
	SeverityModerate Severity = "moderada"
	SeveritySevere   Severity = "severa"
)

// Action is the disposition decided by the inspector
// (08_CATALOGOS: ACCION_INSPECCION).
type Action string

const (
	ActionRelease        Action = "LIBERAR"
	ActionClean          Action = "LIMPIEZA"
	ActionRepair         Action = "REPARACION"
	ActionCleanAndRepair Action = "LIMPIEZA_REPARACION"
	ActionQuarantine     Action = "CUARENTENA"
	ActionScrap          Action = "SCRAP"
)

// Inspection records condition, damage and cost breakdown on return
// receipt through release (04_INSPECCION sheet).
type Inspection struct {
	ID                       string     `json:"id" db:"id"`
	TenantID                 string     `json:"tenant_id" db:"tenant_id"`
	CycleID                  string     `json:"cycle_id,omitempty" db:"cycle_id"`
	AssetID                  string     `json:"asset_id" db:"asset_id"`
	ReceivedAt               time.Time  `json:"received_at" db:"received_at"`
	InspectorID              string     `json:"inspector_id,omitempty" db:"inspector_id"`
	ConditionIn              Condition  `json:"condition_in,omitempty" db:"condition_in"`
	DirtLevel                string     `json:"dirt_level,omitempty" db:"dirt_level"`
	DamageType               string     `json:"damage_type,omitempty" db:"damage_type"`
	Severity                 Severity   `json:"severity,omitempty" db:"severity"`
	Action                   Action     `json:"action" db:"action"`
	LaborMinutes             int        `json:"labor_minutes" db:"labor_minutes"`
	CleaningCostMXN          float64    `json:"cleaning_cost_mxn" db:"cleaning_cost_mxn"`
	RepairCostMXN            float64    `json:"repair_cost_mxn" db:"repair_cost_mxn"`
	PartsCostMXN             float64    `json:"parts_cost_mxn" db:"parts_cost_mxn"`
	ScrapValueMXN            float64    `json:"scrap_value_mxn" db:"scrap_value_mxn"`
	TotalInterventionCostMXN float64    `json:"total_intervention_cost_mxn" db:"total_intervention_cost_mxn"`
	ReleasedAt               *time.Time `json:"released_at,omitempty" db:"released_at"`
	ConditionOut             Condition  `json:"condition_out,omitempty" db:"condition_out"`
	Released                 bool       `json:"released" db:"released"`
	EvidenceBeforeURL        string     `json:"evidence_before_url,omitempty" db:"evidence_before_url"`
	EvidenceAfterURL         string     `json:"evidence_after_url,omitempty" db:"evidence_after_url"`
	ProbableCause            string     `json:"probable_cause,omitempty" db:"probable_cause"`
	ProbableDamageNode       string     `json:"probable_damage_node,omitempty" db:"probable_damage_node"`
	ApprovedBy               string     `json:"approved_by,omitempty" db:"approved_by"`
	Notes                    string     `json:"notes,omitempty" db:"notes"`
	CreatedAt                time.Time  `json:"created_at" db:"created_at"`
}

// ComputeTotalCost sums the cost breakdown into TotalInterventionCostMXN,
// mirroring the workbook's Costo_Total_Intervencion_MXN formula.
func (i *Inspection) ComputeTotalCost() {
	i.TotalInterventionCostMXN = i.CleaningCostMXN + i.RepairCostMXN + i.PartsCostMXN
}
