package workproductcomponent

// Liner Top Test
type LinerTopTest struct {
	// Cement Found On Tool?                                        
	IsCementFoundOnTool                                    *bool    `json:"IsCementFoundOnTool,omitempty"`
	// Liner Negative Test Tool                                     
	LinerNegativeTestTool                                  *string  `json:"LinerNegativeTestTool,omitempty"`
	// Liner Overlap                                                
	LinerOverlap                                           *float64 `json:"LinerOverlap,omitempty"`
	// Liner Positive Test Tool                                     
	LinerPositiveTestTool                                  *string  `json:"LinerPositiveTestTool,omitempty"`
	// Elapsed Time Before Test following end of cement job         
	LinerTopTestElapsedTime                                *float64 `json:"LinerTopTestElapsedTime,omitempty"`
	// Comments or Remarks                                          
	LinerTopTestRemarks                                    *string  `json:"LinerTopTestRemarks,omitempty"`
	// Liner Negative Test Equivalent Mud Weight                    
	NegativeTestEquivalentMudWeight                        *float64 `json:"NegativeTestEquivalentMudWeight,omitempty"`
	// Liner Positive Test Equivalent Mud Weight                    
	PositiveTestEquivalentMudWeight                        *float64 `json:"PositiveTestEquivalentMudWeight,omitempty"`
	// Test Base Measured Depth                                     
	TestBaseMeasuredDepth                                  *float64 `json:"TestBaseMeasuredDepth,omitempty"`
	// Test Base True Vertical Depth                                
	TestBaseTrueVerticalDepth                              *float64 `json:"TestBaseTrueVerticalDepth,omitempty"`
	// Test Top Measured Depth                                      
	TestTopMeasuredDepth                                   *float64 `json:"TestTopMeasuredDepth,omitempty"`
	// Test Top True Vertical Depth                                 
	TestTopTrueVerticalDepth                               *float64 `json:"TestTopTrueVerticalDepth,omitempty"`
}
