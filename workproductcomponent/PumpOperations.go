package workproductcomponent

import "time"

// Information related to pump operations on a drilling/operations report
type PumpOperations struct {
	// Comments and remarks                                                        
	Comments                                                             *string   `json:"Comments,omitempty"`
	// Date and time when the pump operation occurred                              
	DateTime                                                             time.Time `json:"DateTime"`
	// Liner inside diameter.                                                      
	LinerInnerDiameter                                                   *float64  `json:"LinerInnerDiameter,omitempty"`
	// Along-hole measured depth of the measurement from the drill datum.          
	MeasuredDepthBit                                                     *float64  `json:"MeasuredDepthBit,omitempty"`
	// Type of pump operation.                                                     
	OperationType                                                        *string   `json:"OperationType,omitempty"`
	// how much fluid is moved by the pump in one cycle                            
	PumpDisplacement                                                     *float64  `json:"PumpDisplacement,omitempty"`
	// Pump efficiency.                                                            
	PumpEfficiency                                                       *float64  `json:"PumpEfficiency,omitempty"`
	// The model name for the pump                                                 
	PumpModelName                                                        *string   `json:"PumpModelName,omitempty"`
	// The sequence number of the pump (e.g., pump 1)                              
	PumpNumber                                                           *string   `json:"PumpNumber,omitempty"`
	// Pump output (included for efficiency).                                      
	PumpOutput                                                           *float64  `json:"PumpOutput,omitempty"`
	// Pump pressure recorded.                                                     
	PumpPressure                                                         *float64  `json:"PumpPressure,omitempty"`
	// Pump stroke length.                                                         
	StrokeLength                                                         *float64  `json:"StrokeLength,omitempty"`
	// Pump rate (strokes per minute).                                             
	StrokeRate                                                           *float64  `json:"StrokeRate,omitempty"`
}
