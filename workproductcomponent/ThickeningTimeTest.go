package workproductcomponent

// Thickening Time Test data
type ThickeningTimeTest struct {
	// Compressive Strength Index                                                               
	TestIndex                                                                          int64    `json:"TestIndex"`
	// Thickening text consistency/slurry viscosity: Bearden Consistency (Bc) 0 to 100.         
	ThickeningTestConsistency                                                          *int64   `json:"ThickeningTestConsistency,omitempty"`
	// Thickening Test Elapsed Time                                                             
	ThickeningTestElapsedTime                                                          *float64 `json:"ThickeningTestElapsedTime,omitempty"`
	// Thickening Test Pressure                                                                 
	ThickeningTestPressure                                                             *float64 `json:"ThickeningTestPressure,omitempty"`
	// Thickening Test Temperature                                                              
	ThickeningTestTemperature                                                          *float64 `json:"ThickeningTestTemperature,omitempty"`
}
