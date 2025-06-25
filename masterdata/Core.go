package masterdata

// A rock sample obtained by drilling into the earth with a pipe conveyed hollow bit and
// core barrel
type Core struct {
	// A description of the core                                                                         
	Comments                                                                                    *string  `json:"Comments,omitempty"`
	// A measurement that represents the diameter of the recovered core.                                 
	CoreDiameter                                                                                *float64 `json:"CoreDiameter,omitempty"`
	// An identifier, assigned by the well operator, that uniquely identifies the sample.                
	CoreNumber                                                                                  *string  `json:"CoreNumber,omitempty"`
	// The name of the formation from which the core was extracted                                       
	FormationName                                                                               *string  `json:"FormationName,omitempty"`
	// A measurement that represents the length of the core that was recovered from the core             
	// acquisition activity.                                                                             
	Length                                                                                      *float64 `json:"Length,omitempty"`
	// The base depth of the interval to which the coring activity refers. Depth relative to             
	// Planned wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which         
	// holds the depth reference in data.VerticalMeasurement.                                            
	MeasuredDepthBase                                                                           *float64 `json:"MeasuredDepthBase,omitempty"`
	// The start depth of the interval to which the coring activity refers. Depth relative to            
	// Planned wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which         
	// holds the depth reference in data.VerticalMeasurement.                                            
	MeasuredDepthTop                                                                            *float64 `json:"MeasuredDepthTop,omitempty"`
}
