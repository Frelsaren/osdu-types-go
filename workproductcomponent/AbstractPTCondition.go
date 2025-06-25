package workproductcomponent

// Used to capture the conditions at which measurement have been made/computed
//
// The pair of absolute pressure and temperature values describing the condition for a
// particular volume measurement or estimation. The unit of measure context is defined via
// the meta[] block in the record. Search responses will return pressure in Pa (Pascal) and
// temperature in K (Kelvin).
type AbstractPTCondition struct {
	// To capture when Measurement have been made at Standard Conditions (25°C / 100 kPa)                
	// Mutually Exclusive with Pressure/Temperature.                                                     
	// Capture                                                                                           
	IsStandardConditions                                                                         *bool   `json:"IsStandardConditions,omitempty"`
	// Open Text Box to capture the P & T Reference when measurements are made at non standard           
	// conditions (such as "Reservoir", "Tank",…)                                                        
	NonStandardConditionsReference                                                               *string `json:"NonStandardConditionsReference,omitempty"`
	// The recorded absolute pressure condition. The unit of measure context is defined via              
	// meta[] in the Storage record while the Search responses return the value in base SI unit          
	// Pa (Pascal).                                                                                      
	Pressure                                                                                     float64 `json:"Pressure"`
	// The recorded temperature condition. The unit of measure context is defined via meta[] in          
	// the Storage record while the Search responses return the value in base SI unit K (Kelvin).        
	Temperature                                                                                  float64 `json:"Temperature"`
}
