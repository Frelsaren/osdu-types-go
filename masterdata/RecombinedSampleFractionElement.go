package masterdata

// This object holds information about a sample or component of a sample and its
// contribution to a target sample created through the recombination process.
type RecombinedSampleFractionElement struct {
	// The mass fraction contribution of this component to the sample composition.                     
	MassFraction                                                                              *float64 `json:"MassFraction,omitempty"`
	// The mole fraction contribution of this component to the sample composition.                     
	MoleFraction                                                                              *float64 `json:"MoleFraction,omitempty"`
	// This is the OSDU Record ID of an item in the list of  sample components. This attribute         
	// references the components that make up a sample.                                                
	SampleCompositionComponentID                                                              string   `json:"SampleCompositionComponentID"`
	// The volume fraction contribution of this component to the sample composition.                   
	VolumeFraction                                                                            *float64 `json:"VolumeFraction,omitempty"`
}
