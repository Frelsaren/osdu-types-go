package masterdata

// The base of the IsolatedInterval.
type IsolatedIntervalBase struct {
	// Bottom Measured depth of the interval                                                             
	MeasuredDepthBase                                                                           *float64 `json:"MeasuredDepthBase,omitempty"`
	// Bottom true vertical depth of the interval                                                        
	TrueVerticalDepthBase                                                                       *float64 `json:"TrueVerticalDepthBase,omitempty"`
	// The record id  of the wellbore object, to which this IsolatedIntervalBase element belongs         
	WellboreID                                                                                  *string  `json:"WellboreID,omitempty"`
}
