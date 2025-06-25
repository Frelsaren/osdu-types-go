package masterdata

// A cost value associated to a WellActivityPhaseType value.
type WellboreCost struct {
	// The activity type, to which the Cost property is attributed to. The intended target         
	// reference-data type is WellActivityType; Wellbore:1.3.0 and Wellbore:1.4.0 used             
	// WellActivityPhaseType incorrectly.                                                          
	ActivityTypeID                                                                        *string  `json:"ActivityTypeID,omitempty"`
	// The cost value associated with the WellActivityPhaseType.                                   
	Cost                                                                                  *float64 `json:"Cost,omitempty"`
}
