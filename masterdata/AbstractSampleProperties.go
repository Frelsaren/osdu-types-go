package masterdata

// This captures information pertaining to the observed physical properties of the sample.
//
// A nested object definition for ordinary sample properties.
type AbstractSampleProperties struct {
	// This captures the diameter of the sample. This is mostly applicable in core samples.              
	SampleDiameter                                                                              *float64 `json:"SampleDiameter,omitempty"`
	// This refers to the length of the sample. Applicable to rock / core samples.                       
	SampleLength                                                                                *float64 `json:"SampleLength,omitempty"`
	// The kind of orientation of the rock sample with respect to the bedding or drilling                
	// direction. Typical values are Horizontal, Vertical, Axial.                                        
	SampleOrientationTypeID                                                                     *string  `json:"SampleOrientationTypeID,omitempty"`
	// This refers to the volume of the sample acquired. It is applicable to both rock and fluid         
	// samples.                                                                                          
	SampleVolume                                                                                *float64 `json:"SampleVolume,omitempty"`
	// This captures the weight or mass of the sample.                                                   
	SampleWeight                                                                                *float64 `json:"SampleWeight,omitempty"`
}
