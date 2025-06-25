package masterdata

// Describes a BusinessAssociate's status in time.
type BAStatus struct {
	// The date and time at which a given business associate status becomes effective.             
	EffectiveDate                                                                          *string `json:"EffectiveDate,omitempty"`
	// A remark about the business associate status in time.                                       
	Remark                                                                                 *string `json:"Remark,omitempty"`
	// The current status of the Business Associate, such as Active, In Receivership, Sold,        
	// Merged. Property #1                                                                         
	StatusTypeID                                                                           *string `json:"StatusTypeID,omitempty"`
	// The date and time at which a given business associate status is no longer in effect.        
	TerminationDate                                                                        *string `json:"TerminationDate,omitempty"`
}
