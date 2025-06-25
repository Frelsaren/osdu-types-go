package masterdata

// Description, for a given step of the test, of the Fluid System used for this step.
type TestFluid struct {
	// Ambient temperate value                                    
	AmbientTemperature                                   *float64 `json:"AmbientTemperature,omitempty"`
	// Downhole temperature value                                 
	DownholeTemperature                                  *float64 `json:"DownholeTemperature,omitempty"`
	// Density value for the system fluid                         
	FluidDensitySystem                                   *float64 `json:"FluidDensitySystem,omitempty"`
	// Density value for the test fluid                           
	FluidDensityTest                                     *float64 `json:"FluidDensityTest,omitempty"`
	// Same type reference list as used in Fluids Program         
	TestFluidTypeID                                      *string  `json:"TestFluidTypeID,omitempty"`
}
