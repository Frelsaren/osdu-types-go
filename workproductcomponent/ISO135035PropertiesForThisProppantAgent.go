package workproductcomponent

// ISO13503_5 point properties for this proppant agent.
type ISO135035PropertiesForThisProppantAgent struct {
	// Proppant conductivity under stress.                          
	Conductivity                                           *float64 `json:"Conductivity,omitempty"`
	// Proppant permeability under stress.                          
	Permeability                                           *float64 `json:"Permeability,omitempty"`
	// The amount of stress applied.                                
	StressApplied                                          *float64 `json:"StressApplied,omitempty"`
	// The temperature at the time measurements were taken.         
	Temperature                                            *float64 `json:"Temperature,omitempty"`
}
