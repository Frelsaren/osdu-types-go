package masterdata

// The depth range over which the the activity takes place
type DepthRange struct {
	// The end depth of the activity. Depth relative to Planned wellbore ZDP. Navigate via               
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in               
	// data.VerticalMeasurement.                                                                         
	ActivityDepthEnd                                                                            *float64 `json:"ActivityDepthEnd,omitempty"`
	// The start depth of the activity. Depth relative to Planned wellbore ZDP. Navigate via             
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in               
	// data.VerticalMeasurement.                                                                         
	ActivityDepthStart                                                                          *float64 `json:"ActivityDepthStart,omitempty"`
	// The depth of the hole at the end of the activity. Depth relative to Planned wellbore ZDP.         
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth               
	// reference in data.VerticalMeasurement.                                                            
	HoleDepthEnd                                                                                *float64 `json:"HoleDepthEnd,omitempty"`
	// The depth of the hole at the start of the activity. Depth relative to Planned wellbore            
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth          
	// reference in data.VerticalMeasurement.                                                            
	HoleDepthStart                                                                              *float64 `json:"HoleDepthStart,omitempty"`
}
