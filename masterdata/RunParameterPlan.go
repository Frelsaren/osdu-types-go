package masterdata

// Operating parameters of a drill string run
type RunParameterPlan struct {
	// A series of operating parameters observed during the run                                                         
	OperationParameterPlan                                                                      *OperationParameterPlan `json:"OperationParameterPlan,omitempty"`
	// The end measured depth of the hole section at the end of the run. Depth relative to                              
	// Planned wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which                        
	// holds the depth reference in data.VerticalMeasurement.                                                           
	RunEndHoleMeasuredDepth                                                                     *float64                `json:"RunEndHoleMeasuredDepth,omitempty"`
	// The start measured depth of the hole section at the start of the run. Depth relative to                          
	// Planned wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which                        
	// holds the depth reference in data.VerticalMeasurement.                                                           
	RunStartHoleMeasuredDepth                                                                   *float64                `json:"RunStartHoleMeasuredDepth,omitempty"`
}
