package workproductcomponent

import "time"

// Pump flow back test for this stim job diagnostic session.
type StimPumpFlowBackTest struct {
	// Casing pressure.                                                      
	CasingPressure                                             *float64      `json:"CasingPressure,omitempty"`
	// End time for the test.                                                
	FlowBackEndDateTime                                        *time.Time    `json:"FlowBackEndDateTime,omitempty"`
	// Start time for the test.                                              
	FlowBackStartDateTime                                      *time.Time    `json:"FlowBackStartDateTime,omitempty"`
	// Total fluid volume recovered during a flow back test.                 
	FlowBackVolume                                             *float64      `json:"FlowBackVolume,omitempty"`
	// The time required for the fracture width to become zero.              
	FractureCloseDuration                                      *float64      `json:"FractureCloseDuration,omitempty"`
	// The pressure when the fracture width becomes zero.                    
	FractureClosePressure                                      *float64      `json:"FractureClosePressure,omitempty"`
	// General remarks for this stim pump flow back test.                    
	Remarks                                                    *string       `json:"Remarks,omitempty"`
	// Step                                                                  
	Step                                                       []StepElement `json:"Step,omitempty"`
	// Tubing pressure.                                                      
	TubingPressure                                             *float64      `json:"TubingPressure,omitempty"`
}
