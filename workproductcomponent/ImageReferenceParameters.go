package workproductcomponent

// Similar to ImageParameters, but reference-data relationships describing reference-value
// list controlled image properties.
type ImageReferenceParameters struct {
	// The type of beam in the electron microscope for this sample image, for example, electron         
	// or ion beam.                                                                                     
	BeamTypeID                                                                                  *string `json:"BeamTypeID,omitempty"`
	// The colour space for this sample image, for example, could be RGB, Monochrome, or CMYK.          
	ColourSpaceTypeID                                                                           *string `json:"ColourSpaceTypeID,omitempty"`
	// The mode of lighting for this sample image, for example, Natural, White and Ultraviolet.         
	LightingConditionTypeID                                                                     *string `json:"LightingConditionTypeID,omitempty"`
	// The type of filter applied to the light source of this sample image, for example, Cross          
	// Polarised, Brightfield and Darkfield.                                                            
	OpticalLightFilterTypeID                                                                    *string `json:"OpticalLightFilterTypeID,omitempty"`
	// The light path for this sample image, for example, incident or transmitted.                      
	OpticalLightPathTypeID                                                                      *string `json:"OpticalLightPathTypeID,omitempty"`
	// The source of the pixel resolution for this sample image, for example, it may be inferred        
	// from a scale bar or a known reference.                                                           
	ResolutionSourceTypeID                                                                      *string `json:"ResolutionSourceTypeID,omitempty"`
	// The type of X Ray detector used for this sample image, for example, Wavelength Dispersive        
	// Detector WDS/WDX) and Energy Dispersive Detector (EDS/EDX).                                      
	XRayDetectorTypeID                                                                          *string `json:"XRayDetectorTypeID,omitempty"`
}
