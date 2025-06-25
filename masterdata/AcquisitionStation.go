package masterdata

// A Station is defined as a static point within the pass - where one or many tests can be
// tried out.
type AcquisitionStation struct {
	// / Isolated Interval) present at this test station                                                 
	CompletionID                                                                                *string  `json:"CompletionID,omitempty"`
	// Measured Depth of the station                                                                     
	PressurePointMeasuredDepth                                                                  *float64 `json:"PressurePointMeasuredDepth,omitempty"`
	// List of Names - in the reference InterpretationSet (Marker or Interval) array - of the            
	// prognosed interpretations to be tested at this station                                            
	PrognosedInterpretationNames                                                                []string `json:"PrognosedInterpretationNames,omitempty"`
	// Array of Identifier of the reservoir units (Reservoir, Segment, Sectors,…) expected to be         
	// tested at this station                                                                            
	PrognosedReservoirUnitsIDs                                                                  []string `json:"PrognosedReservoirUnitsIDs,omitempty"`
	StationIdentifier                                                                           *int64   `json:"StationIdentifier,omitempty"`
}
