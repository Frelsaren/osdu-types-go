package workproductcomponent

// If populated RoutineCoreAnalysis contains conditions and results of a routine core
// analysis. The array AnalysisTypeIDs must then contain a value of RoutineCoreAnalysis.
type RoutineCoreAnalysis struct {
	// Th measured grain density                                                                    
	GrainDensity                                                                   *float64         `json:"GrainDensity,omitempty"`
	// The measurement type used to obtain the GrainDensity value.                                  
	GrainDensityMeasurementTypeID                                                  *string          `json:"GrainDensityMeasurementTypeID,omitempty"`
	// The conditions under which this analysis has been carried out                                
	RCAMeasurements                                                                []RCAMeasurement `json:"RCAMeasurements,omitempty"`
	// The fraction of gas volume compared to the total pore volume.                                
	SaturationGas                                                                  *float64         `json:"SaturationGas,omitempty"`
	// The saturation method type, like Dean Stark, Retort, Karl Fischer.                           
	SaturationMethodTypeID                                                         *string          `json:"SaturationMethodTypeID,omitempty"`
	// The fraction of oil volume compared to the total pore volume.                                
	SaturationOil                                                                  *float64         `json:"SaturationOil,omitempty"`
	// The fraction of water volume compared to the total pore volume.                              
	SaturationWater                                                                *float64         `json:"SaturationWater,omitempty"`
	// The weight of salt in the brine, needed in conjunction with SaturationWater.                 
	WeightOfSalt                                                                   *float64         `json:"WeightOfSalt,omitempty"`
}
