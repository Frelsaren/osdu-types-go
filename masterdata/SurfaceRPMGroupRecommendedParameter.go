package masterdata

// The definition of the point
type SurfaceRPMGroupRecommendedParameter struct {
	// The measured depth at which the measurement was observed. Depth relative to Planned                              
	// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds                          
	// the depth reference in data.VerticalMeasurement.                                                                 
	ObservationMeasuredDepth                                                                  *float64                  `json:"ObservationMeasuredDepth,omitempty"`
	// The source indicator associated with this point.                                                                 
	PointsSources                                                                             []MischievousPointsSource `json:"PointsSources,omitempty"`
	// The value observed at the measured depth                                                                         
	Value                                                                                     *float64                  `json:"Value,omitempty"`
	// Unit of measure of the Value value                                                                               
	ValueUnitID                                                                               *string                   `json:"ValueUnitID,omitempty"`
}
