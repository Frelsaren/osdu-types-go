package workproductcomponent

// An injection test involving multiple steps of injection rate and pressure, where a curve
// deflection and change of slope indicates the fracture breakdown pressure.
type StimStepDownTest struct {
	// The density of the fluid at the bottom of the hole adjusting for bottomhole temperature                            
	// and pressure during the step-down test.                                                                            
	BHFluidDensity                                                                              *float64                  `json:"BHFluidDensity,omitempty"`
	// A coefficient used in the equation for calculation of the pressure drop across a                                   
	// perforation set.                                                                                                   
	DischargeCoefficient                                                                        *float64                  `json:"DischargeCoefficient,omitempty"`
	// Diameter of the injection point or perforation.                                                                    
	EntryHoleDiameter                                                                           *float64                  `json:"EntryHoleDiameter,omitempty"`
	// The data related to a particular step in the step-down test.                                                       
	FlowBackTestStep                                                                            []FlowBackTestStepElement `json:"FlowBackTestStep,omitempty"`
	// The initial shut in pressure.                                                                                      
	InitialShutinPressure                                                                       *float64                  `json:"InitialShutinPressure,omitempty"`
	// The number of perforations in the interval being tested that are calculated to be open to                          
	// injection, which is determined during the step-down test.                                                          
	PerforationsEffective                                                                       *int64                    `json:"PerforationsEffective,omitempty"`
	// The total number of perforations in the interval being tested.                                                     
	PerforationsTotal                                                                           *int64                    `json:"PerforationsTotal,omitempty"`
	// General remarks about this Stim Step Down Test.                                                                    
	Remarks                                                                                     *string                   `json:"Remarks,omitempty"`
}
