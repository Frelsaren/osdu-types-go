package workproductcomponent

// Viscosity values observed
type RheometerViscosity struct {
	// Rotational speed of the rheometer, typically in RPM.                                         
	Speed                                                                                   float64 `json:"Speed"`
	// The raw reading from a rheometer. This could be, but is not necessarily, a viscosity.        
	Viscosity                                                                               float64 `json:"Viscosity"`
}
