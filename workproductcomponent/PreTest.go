package workproductcomponent

// Describes the pieces of information required to identity individual pretests within the
// complete record.
type PreTest struct {
	// DEPRECATED: Time index where the individual pretest ends into the sequence           
	PreTestEndTime                                                                 *float64 `json:"PreTestEndTime,omitempty"`
	// DEPRECATED: Sequential number identifying the pretest within the record              
	PreTestNumber                                                                  *int64   `json:"PreTestNumber,omitempty"`
	// DEPRECATED: Time index where the individual pretest starts into the sequence         
	PreTestStartTime                                                               *float64 `json:"PreTestStartTime,omitempty"`
}
