package masterdata

// The gas-oil ratio recorded for this sample recombination process as well as the
// volumetric reference conditions for both the oil and gas phases. This is typically
// required for fluid sample types.
type RecombinationGasOilRatio struct {
	// The Gas Oil Ratio calculated at the reference conditions specified for each stream (Oil                    
	// or Gas)                                                                                                    
	GasOilRatio                                                                               float64             `json:"GasOilRatio"`
	// The pressure and temperature reference values for the gas stream.                                          
	VolumeReferenceConditionGas                                                               AbstractPTCondition `json:"VolumeReferenceConditionGas"`
	// The pressure and temperature reference values for the oil stream.                                          
	VolumeReferenceConditionOil                                                               AbstractPTCondition `json:"VolumeReferenceConditionOil"`
}
