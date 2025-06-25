package masterdata

import "time"

// The history of life cycle states that the Project has been through..
type ProjectStates struct {
	// The date and time at which the state becomes effective.                
	EffectiveDateTime                                              *time.Time `json:"EffectiveDateTime,omitempty"`
	// The Project life cycle state from planning to completion.              
	ProjectStateTypeID                                             *string    `json:"ProjectStateTypeID,omitempty"`
	// The date and time at which the state is no longer in effect.           
	TerminationDateTime                                            *time.Time `json:"TerminationDateTime,omitempty"`
}
