package workproductcomponent

// Log Evaluation Depth
type LogEvaluationDepth struct {
	// Evaluated Interval Base Measured Depth         
	EvaluatedIntervalMeasuredDepthBase       *float64 `json:"EvaluatedIntervalMeasuredDepthBase,omitempty"`
	// Evaluated Interval Top Measured Depth          
	EvaluatedIntervalMeasuredDepthTop        *float64 `json:"EvaluatedIntervalMeasuredDepthTop,omitempty"`
	// Is Isolated                                    
	IsIsolated                               *bool    `json:"IsIsolated,omitempty"`
	// Remarks                                        
	Remarks                                  *string  `json:"Remarks,omitempty"`
}
