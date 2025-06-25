package workproductcomponent

// Hesitation Squeeze History
type HesitationSqueezeHistory struct {
	// Final Pressure                          
	FinalPressure                     *float64 `json:"FinalPressure,omitempty"`
	// Duration between pumping stages         
	HesitationTime                    *float64 `json:"HesitationTime,omitempty"`
	// Initial Pressure                        
	InitialPressure                   *float64 `json:"InitialPressure,omitempty"`
	// Remarks                                 
	Remarks                           *string  `json:"Remarks,omitempty"`
	// Sequence Number                         
	SequenceNo                        *int64   `json:"SequenceNo,omitempty"`
	// Volume Pumped                           
	Volume                            *float64 `json:"Volume,omitempty"`
}
