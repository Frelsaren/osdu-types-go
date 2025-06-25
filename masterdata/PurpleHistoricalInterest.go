package masterdata

// A 'well interest' at some time period as defined by effective and termination date.
type PurpleHistoricalInterest struct {
	// The date and time at which the well interest type becomes effective.                            
	EffectiveDateTime                                                                          *string `json:"EffectiveDateTime,omitempty"`
	// Business Interest [Well Interest Type] describes whether a company currently considers a        
	// well or its data to be a real or planned asset, and if so, the nature of and motivation         
	// for that company's interest.                                                                    
	InterestTypeID                                                                             *string `json:"InterestTypeID,omitempty"`
	// The date and time at which the well interest type is no longer in effect.                       
	TerminationDateTime                                                                        *string `json:"TerminationDateTime,omitempty"`
}
