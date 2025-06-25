package masterdata

// The geometric sections that combine to define the shape of a complex target
type GeometrySection struct {
	// Direction of straight line section or radius of arc for continuous curve section.                                
	AngleArc                                                                                   *float64                 `json:"AngleArc,omitempty"`
	// Length of straight line section or radius of arc for continuous curve section.                                   
	LengthRadius                                                                               *float64                 `json:"LengthRadius,omitempty"`
	// 2D coordinates that defines the start of the section.                                                            
	Location                                                                                   *AbstractSpatialLocation `json:"Location,omitempty"`
	// Name of this Geometry Section                                                                                    
	SectionName                                                                                *string                  `json:"SectionName,omitempty"`
	// Sequence number of the section (1,2,3…)                                                                          
	SectionNumber                                                                              *float64                 `json:"SectionNumber,omitempty"`
	// True Vertical Depth of this section                                                                              
	SectionTVD                                                                                 *float64                 `json:"SectionTVD,omitempty"`
	// Height of target above center point at the start of the section.  In the case of an arc,                         
	// the thickness above should vary linearly with the arc length.                                                    
	ThicknessAbove                                                                             *float64                 `json:"ThicknessAbove,omitempty"`
	// Depth of target below center point at the start of the section. In the case of an arc,                           
	// the thickness below should vary linearly with the arc length.                                                    
	ThicknessBelow                                                                             *float64                 `json:"ThicknessBelow,omitempty"`
	// Section scope : Line or Arc                                                                                      
	TypeTargetSectionScope                                                                     *string                  `json:"TypeTargetSectionScope,omitempty"`
}
