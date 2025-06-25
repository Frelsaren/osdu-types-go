package workproductcomponent

import "time"

// One step/interval in an Activity's or ProjectActivity's state.
//
// The current or last state this activity transitioned to. It is a copy of the last element
// in ActivityStates[]. If there is only one state recorded, the ActivityStates[] can stay
// empty.
type AbstractActivityState struct {
	// The ActivityStatus is a set of major activity phases that are significant to business             
	// stakeholders.                                                                                     
	ActivityStatusID                                                                          *string    `json:"ActivityStatusID,omitempty"`
	// The date and time at which the activity status becomes effective.                                 
	EffectiveDateTime                                                                         *time.Time `json:"EffectiveDateTime,omitempty"`
	// An optional remark associated with the ActivityStatus and time interval.                          
	Remark                                                                                    *string    `json:"Remark,omitempty"`
	// The date and time at which the activity status is no longer in effect. For still                  
	// effective activity states, the TerminationDateTime is left absent. For zero-duration              
	// intervals (events), the TerminationDateTime set to the same value as EffectiveDateTime.           
	TerminationDateTime                                                                       *time.Time `json:"TerminationDateTime,omitempty"`
}
