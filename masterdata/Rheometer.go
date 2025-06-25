package masterdata

// Rheometer values observed
type Rheometer struct {
	// The pressure at which the rheometer values were measured                        
	PressureRheometer                                             float64              `json:"PressureRheometer"`
	// The viscosities recorder during the Rheometer test                              
	RheometerViscosities                                          []RheometerViscosity `json:"RheometerViscosities,omitempty"`
	// The temperature at which the rheometer values were measured                     
	TemperatureRheometer                                          float64              `json:"TemperatureRheometer"`
}
