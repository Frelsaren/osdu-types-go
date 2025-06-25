package masterdata

// Generic description of an individual pressure criterion for a single step of the test
type HighPressureCriterionElement struct {
	// Boolean that states if the criteria are for the high pressure stage (low pressure stage           
	// otherwise)                                                                                        
	IsHigh                                                                                      *bool    `json:"IsHigh,omitempty"`
	// the maximum allowable pressure loss for a pressure decline test (expressed as an absolute         
	// value)                                                                                            
	MaxAllowedAbsoluteDecline                                                                   *float64 `json:"MaxAllowedAbsoluteDecline,omitempty"`
	// the maximum allowable pressure loss for a pressure decline test (expressed as a                   
	// percentage of test pressure)                                                                      
	MaxAllowedPercentageDecline                                                                 *float64 `json:"MaxAllowedPercentageDecline,omitempty"`
	// Criteria pressure rate of change for a flat-line test. Averaged over the duration of  the         
	// Validation Duration                                                                               
	MaxAllowedRateOfChange                                                                      *float64 `json:"MaxAllowedRateOfChange,omitempty"`
	// A successful negative pressure test or inflow test using the shut-in pressure method              
	// requires the pressure to remain below the maximum allowable target test pressure.                 
	TargetPressureMax                                                                           *float64 `json:"TargetPressureMax,omitempty"`
	// A successful positive pressure test requires the pressure to remain above the minimum             
	// allowable target test pressure.                                                                   
	TargetPressureMin                                                                           *float64 `json:"TargetPressureMin,omitempty"`
	// The time duration for which a pressure test must “hold” (to a prescribed testing                  
	// criteria) in order for the test to be validated (typically minutes e.g. 5mins (BOP,               
	// 30mins Casing)                                                                                    
	ValidationDuration                                                                          *float64 `json:"ValidationDuration,omitempty"`
}
