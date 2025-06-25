package masterdata

import "time"

// A significant occurrence in the life of a facility, which often changes its state, or the
// state of one of its components. It can describe a point-in-time (event) or a time
// interval of a specific type (FacilityEventType).
type AbstractFacilityEvent struct {
	// The date and time at which the event took place or takes effect.                                    
	EffectiveDateTime                                                                           *time.Time `json:"EffectiveDateTime,omitempty"`
	// The facility event type is a picklist. Examples: 'Permit', 'Spud', 'Abandon', etc.                  
	FacilityEventTypeID                                                                         *string    `json:"FacilityEventTypeID,omitempty"`
	// A comment or remark about the facility event.                                                       
	Remark                                                                                      *string    `json:"Remark,omitempty"`
	// The date and time at which the event is no longer in effect. For point-in-time events the           
	// 'TerminationDateTime' must be set equal to 'EffectiveDateTime'. Open time intervals have            
	// an absent 'TerminationDateTime'.                                                                    
	TerminationDateTime                                                                         *time.Time `json:"TerminationDateTime,omitempty"`
}
