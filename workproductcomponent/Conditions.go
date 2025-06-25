package workproductcomponent

// The conditions under which this analysis has been carried out
type Conditions struct {
	// The average pressure in which the gas permeability has been measured.                        
	MeanPressure                                                                           *float64 `json:"MeanPressure,omitempty"`
	// The pressure of the fluid in the pore space.                                                 
	PorePressure                                                                           *float64 `json:"PorePressure,omitempty"`
	// The pressure condition for the analysis.                                                     
	Pressure                                                                               *float64 `json:"Pressure,omitempty"`
	// Identifies the environment where the pressure was measured. E.g. Ambient, Overburden         
	PressureMeasurementTypeID                                                              *string  `json:"PressureMeasurementTypeID,omitempty"`
	// The temperature condition for the analysis.                                                  
	Temperature                                                                            *float64 `json:"Temperature,omitempty"`
}
