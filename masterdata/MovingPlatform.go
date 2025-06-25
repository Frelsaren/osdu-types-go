package masterdata

// An array that describes the survey acquisition properties relevant to a moving platform
// survey, for example, airborne, marine vessel or other moving vehicle
type MovingPlatform struct {
	// The total length of the line distance calculated from the line data.                                              
	CalculatedLineLength                                                                         *float64                `json:"CalculatedLineLength,omitempty"`
	// Captures the average flying height and the flying height parameters such as the                                   
	// instrumentation, datum and the clearance method.                                                                  
	FlyingHeightParameters                                                                       *FlyingHeightParameters `json:"FlyingHeightParameters,omitempty"`
	// A boolean flag indicating if the survey was acquired from a moving platform, for example,                         
	// airborne, marine vessel or other moving vehicle.                                                                  
	// This boolean flag also permits a moving platform survey to be identified prior to the                             
	// loading of the moving platform related properties.                                                                
	// This should be true if any of the below moving platform properties could be populated now                         
	// or in the future.                                                                                                 
	IsMovingPlatformSurvey                                                                       *bool                   `json:"IsMovingPlatformSurvey,omitempty"`
	// The nominal line length covered by the survey. This value is usually entered by the end                           
	// user. Equivalent to the two boxes for an area.                                                                    
	NominalLineLength                                                                            *float64                `json:"NominalLineLength,omitempty"`
	// The free text name of the platform or the vessel that acquired the survey                                         
	PlatformName                                                                                 *string                 `json:"PlatformName,omitempty"`
	// The distance between primary lines. If acquired with seismic data this will be the inline.                        
	PrimaryLineSpacing                                                                           *float64                `json:"PrimaryLineSpacing,omitempty"`
	// The sampling station intervals recorded in time. Sampling intervals are often received in                         
	// 1 second intervals or similar and/or p190 for marine data                                                         
	SamplingIntervalTime                                                                         *float64                `json:"SamplingIntervalTime,omitempty"`
	// An association to a seismic acquisition survey when seismic and gravity and magnetic data                         
	// are acquired simultaneously, typically in the marine environment.                                                 
	// The seismic acquisition survey object contains data relevant to the acquisition such as                           
	// the survey geometry.                                                                                              
	SeismicAcquisitionSurveyID                                                                   *string                 `json:"SeismicAcquisitionSurveyID,omitempty"`
	// The distance between lines that intersect the primary lines. These lines are at a                                 
	// different azimuth to the primary lines. Typically perpendicular to the primary line.                              
	TieLineSpacing                                                                               *float64                `json:"TieLineSpacing,omitempty"`
}
