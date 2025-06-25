package workproductcomponent

import "time"

// General information about a gas reading taken during the drill report period
type GasReading struct {
	// Date and time of the gas reading.                                                               
	DateTime                                                                                 time.Time `json:"DateTime"`
	// Ethane (C2) concentration.                                                                      
	Ethane                                                                                   *float64  `json:"Ethane,omitempty"`
	// Top measured depth for which the gas reading was conducted                                      
	GasReadingMeasureDepthTop                                                                *float64  `json:"GasReadingMeasureDepthTop,omitempty"`
	// Bottom true vertical depth interval over which the gas reading was conducted.                   
	GasReadingTvdBase                                                                        *float64  `json:"GasReadingTvdBase,omitempty"`
	// Iso-butane (iC4) concentration.                                                                 
	Isobutane                                                                                *float64  `json:"Isobutane,omitempty"`
	// Iso-pentane (iC5) concentration                                                                 
	Isopentane                                                                               *float64  `json:"Isopentane,omitempty"`
	// Methane (C1) concentration                                                                      
	Methane                                                                                  *float64  `json:"Methane,omitempty"`
	// Nor-butane (nC4) concentration.                                                                 
	Norbutane                                                                                *float64  `json:"Norbutane,omitempty"`
	// Propane (C3) concentration.                                                                     
	Propane                                                                                  *float64  `json:"Propane,omitempty"`
	// Type of gas reading, e.g.circulating, background gas, connection gas, drilling                  
	// background, gas, drilling gas peak, flow check gas, no readings, other, shut down gas,          
	// trip gas, unknown.                                                                              
	ReadingType                                                                              *string   `json:"ReadingType,omitempty"`
	// The average gas reading                                                                         
	TotalGasAverage                                                                          *float64  `json:"TotalGasAverage,omitempty"`
	// The highest gas reading.                                                                        
	TotalGasMax                                                                              *float64  `json:"TotalGasMax,omitempty"`
	// The lowest gas reading.                                                                         
	TotalGasMin                                                                              *float64  `json:"TotalGasMin,omitempty"`
}
