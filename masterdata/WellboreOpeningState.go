package masterdata

import "time"

// The life cycle status of a WellboreOpening at some point in time.
type WellboreOpeningState struct {
	// The date and time at which the WellboreOpening state becomes effective.                          
	EffectiveDateTime                                                                        *time.Time `json:"EffectiveDateTime,omitempty"`
	// A comment or remark attributed to the state.                                                     
	Remark                                                                                   *string    `json:"Remark,omitempty"`
	// The date and time at which the WellboreOpening state is no longer in effect.                     
	TerminationDateTime                                                                      *time.Time `json:"TerminationDateTime,omitempty"`
	// WellboreOpening State Type is a set of major phases that are significant to regulators           
	// and/or business stakeholders.                                                                    
	WellboreOpeningStateTypeID                                                               *string    `json:"WellboreOpeningStateTypeID,omitempty"`
}
