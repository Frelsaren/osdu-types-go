package masterdata

// Set of attributes capturing the Life Cycle Statuses of the Reservoir, a concept which is
// typically chronological.
type LifeCycleStatuses struct {
	// The date and time at which the reservoir status type becomes effective.                    
	EffectiveDateTime                                                                     *string `json:"EffectiveDateTime,omitempty"`
	// The Life Cycle Status of the Reservoir, a concept which is typically chronological.        
	LifeCycleStatusID                                                                     *string `json:"LifeCycleStatusID,omitempty"`
	// Date of a status change                                                                    
	StatusDate                                                                            *string `json:"StatusDate,omitempty"`
	// The date and time at which the reservoir status type is no longer in effect.               
	TerminationDateTime                                                                   *string `json:"TerminationDateTime,omitempty"`
}
