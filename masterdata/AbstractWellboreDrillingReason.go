package masterdata

import "time"

// Purpose for drilling a wellbore, which often is an indication of the level of risk.
type AbstractWellboreDrillingReason struct {
	// Identifier of the drilling reason type for the corresponding time period.                       
	DrillingReasonTypeID                                                                    *string    `json:"DrillingReasonTypeID,omitempty"`
	// The date and time at which the event becomes effective.                                         
	EffectiveDateTime                                                                       *time.Time `json:"EffectiveDateTime,omitempty"`
	// The Lahee classification, based on the traditional, commonly accepted, scheme to                
	// categorize wells by the general degree of risk assumed by the operator at the time of           
	// drilling.                                                                                       
	LaheeClassID                                                                            *string    `json:"LaheeClassID,omitempty"`
	// A remark or comment explaining the drilling reason or LaheeClass assignment.                    
	Remark                                                                                  *string    `json:"Remark,omitempty"`
	// The date and time at which the event is no longer in effect.                                    
	TerminationDateTime                                                                     *time.Time `json:"TerminationDateTime,omitempty"`
}
