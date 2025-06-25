package masterdata

import "time"

// The organisation that was responsible for a facility at some point in time.
type AbstractFacilityOperator struct {
	// The date and time at which the facility operator becomes effective.                                
	EffectiveDateTime                                                                          *time.Time `json:"EffectiveDateTime,omitempty"`
	// Internal, unique identifier for an item 'AbstractFacilityOperator'. This identifier is             
	// used by 'AbstractFacility.CurrentOperatorID' and 'AbstractFacility.InitialOperatorID'.             
	FacilityOperatorID                                                                         *string    `json:"FacilityOperatorID,omitempty"`
	// The company that currently operates, or previously operated the facility                           
	FacilityOperatorOrganisationID                                                             *string    `json:"FacilityOperatorOrganisationID,omitempty"`
	// A comment or remark about the facility operator.                                                   
	Remark                                                                                     *string    `json:"Remark,omitempty"`
	// The date and time at which the facility operator is no longer in effect. If the operator           
	// is still effective, the 'TerminationDateTime' is left absent.                                      
	TerminationDateTime                                                                        *time.Time `json:"TerminationDateTime,omitempty"`
}
