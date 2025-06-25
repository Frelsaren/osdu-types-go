package masterdata

// Tubular Sensor Component Schema
type Sensor struct {
	// Additional comments and remarks                                       
	Comments                                                        *string  `json:"Comments,omitempty"`
	// Effective offset of measurement of sensor                             
	MeasurementOffset                                               *string  `json:"MeasurementOffset,omitempty"`
	// Sensor offset from the bottom of the component                        
	OffsetBottom                                                    *float64 `json:"OffsetBottom,omitempty"`
	// Specifies the type of sensor in a tubular string.                     
	SensorMeasurementType                                           *string  `json:"SensorMeasurementType,omitempty"`
	// Unique identifier for this Instance of Sensor log                     
	SensorTypeID                                                    string   `json:"SensorTypeID"`
	// An array of well log generic tool types used in this Log Run.         
	ToolClassIDs                                                    []string `json:"ToolClassIDs,omitempty"`
	// An array of PWLS tool mnemonics used in this Log Run.                 
	ToolCodeIDs                                                     []string `json:"ToolCodeIDs,omitempty"`
}
