package masterdata

import "time"

// The life cycle status of a facility at some point in time.
type AbstractFacilityState struct {
	// The date and time at which the facility state becomes effective.                                    
	EffectiveDateTime                                                                           *time.Time `json:"EffectiveDateTime,omitempty"`
	// Life Cycle [Facility State Type] is a set of major phases that are significant to                   
	// regulators and/or business stakeholders. Life Cycle may apply to a well or its components           
	// [or other facility].                                                                                
	FacilityStateTypeID                                                                         *string    `json:"FacilityStateTypeID,omitempty"`
	// A comment or remark about the facility state.                                                       
	Remark                                                                                      *string    `json:"Remark,omitempty"`
	// The date and time at which the facility state is no longer in effect.                               
	TerminationDateTime                                                                         *time.Time `json:"TerminationDateTime,omitempty"`
}
