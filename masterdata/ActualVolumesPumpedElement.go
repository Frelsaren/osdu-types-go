package masterdata

// Individual description, for a given pressure, of the actual volume pumped for a single
// step of the test.
type ActualVolumesPumpedElement struct {
	// Maximum actual volume pumped at specified pressure.                            
	MaximumVolume                                                            *float64 `json:"MaximumVolume,omitempty"`
	// Minimum actual volume pumped at specified pressure.                            
	MinimumVolume                                                            *float64 `json:"MinimumVolume,omitempty"`
	// Value for actual pressure.                                                     
	Pressure                                                                 *float64 `json:"Pressure,omitempty"`
	// The pumping rate for the actual volume pumped at a specified pressure.         
	PumpingRate                                                              *float64 `json:"PumpingRate,omitempty"`
	// Actual volume pumped at specified pressure                                     
	Volume                                                                   *float64 `json:"Volume,omitempty"`
}
