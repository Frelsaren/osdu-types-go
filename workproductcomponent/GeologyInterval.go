package workproductcomponent

// Geology Interval Component Schema
type GeologyInterval struct {
	// Chronostratigraphic classification.                                                
	ChronostratigraphicIDs                                                       []string `json:"ChronostratigraphicIDs,omitempty"`
	// Description of item and details.                                                   
	Description                                                                  *string  `json:"Description,omitempty"`
	// Measured depth at base of interval.                                                
	IntervalMeasuredDepthBase                                                    float64  `json:"IntervalMeasuredDepthBase"`
	// Measured depth at top of interval.                                                 
	IntervalMeasuredDepthTop                                                     float64  `json:"IntervalMeasuredDepthTop"`
	// True vertical depth at base of interval.                                           
	IntervalTVDBase                                                              *float64 `json:"IntervalTVDBase,omitempty"`
	// True vertical depth at top of the section.                                         
	IntervalTVDTop                                                               *float64 `json:"IntervalTVDTop,omitempty"`
	// The geological name for the type of lithology from the enum table listing          
	// a                                                                                  
	// subset of the OneGeology / CGI defined formation types.                            
	LithologyTypeID                                                              string   `json:"LithologyTypeID"`
	// Array of lithological types - allowing for one or more lithological types.         
	LithologyTypeIDs                                                             []string `json:"LithologyTypeIDs,omitempty"`
	// Name of lithostratigraphy, regionally dependent.                                   
	LithostratigraphicIDs                                                        []string `json:"LithostratigraphicIDs,omitempty"`
	// Description of the Reference Datum used for Depth Measurement                      
	ReferenceDatum                                                               *string  `json:"ReferenceDatum,omitempty"`
	// Identifier of the reference trajectory for TVD measurements                        
	ReferenceWellboreTrajectoryID                                                *string  `json:"ReferenceWellboreTrajectoryID,omitempty"`
}
