package workproductcomponent

// Fluid Compressive Strength Test data
type CompressiveStrength struct {
	// Compressive Strength                                           
	CompressiveStrength                                      *float64 `json:"CompressiveStrength,omitempty"`
	// Compressive Strength Elapsed Time                              
	CompressiveStrengthElapsedTime                           *float64 `json:"CompressiveStrengthElapsedTime,omitempty"`
	// Compressive Strength Index                                     
	CompressiveStrengthIndex                                 int64    `json:"CompressiveStrengthIndex"`
	// Compressive strength temperature                               
	CompressiveStrengthTemperature                           *float64 `json:"CompressiveStrengthTemperature,omitempty"`
	// Pressure held during Compressive Strength test                 
	CompressiveStrengthTestPressure                          *float64 `json:"CompressiveStrengthTestPressure,omitempty"`
	// Transit Time measured over a set length (microseconds)         
	CompressiveStrengthTransitTime                           *float64 `json:"CompressiveStrengthTransitTime,omitempty"`
}
