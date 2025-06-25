package masterdata

// The Sidewall Coring specific sub-structure.
type SidewallCoring struct {
	// The name identifying the type of coring procedure used to extract the sidewall core,                       
	// e.g., percussion or rotary (mechanical)  sidewall core.                                                    
	CoreTypeID                                                                               *string              `json:"CoreTypeID,omitempty"`
	// The ratio between planned sidewall cores versus the actually retrieved sidewall cores.                     
	PlannedVersusActual                                                                      *float64             `json:"PlannedVersusActual,omitempty"`
	// The array of sidewall core samples, their depth and description.                                           
	SidewallCores                                                                            []SidewallCoreSample `json:"SidewallCores,omitempty"`
}
