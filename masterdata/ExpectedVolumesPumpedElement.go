package masterdata

// Individual description, for a given pressure, of the expected volume pumped for a single
// step of the test.
type ExpectedVolumesPumpedElement struct {
	// Maximum expected volume pumped at specified pressure.                            
	MaximumVolume                                                              *float64 `json:"MaximumVolume,omitempty"`
	// Minimum expected  volume pumped at specified pressure.                           
	MinimumVolume                                                              *float64 `json:"MinimumVolume,omitempty"`
	// Value for planned pressure.                                                      
	Pressure                                                                   float64  `json:"Pressure"`
	// The pumping rate for the expected volume pumped at a specified pressure.         
	PumpingRate                                                                *float64 `json:"PumpingRate,omitempty"`
	// Expected volume pumped at specified pressure                                     
	Volume                                                                     *float64 `json:"Volume,omitempty"`
}
