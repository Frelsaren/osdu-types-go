package workproductcomponent

import "time"

// Investigation Status history of the NPT event
type NPTInvestigationStatus struct {
	// Status Date/time                           
	InvestigationStatusDateTime        *time.Time `json:"InvestigationStatusDateTime,omitempty"`
	// Investigation Status Remarks               
	InvestigationStatusRemarks         *string    `json:"InvestigationStatusRemarks,omitempty"`
	// NPT Investigation Status Type ID           
	InvestigationStatusTypeID          *string    `json:"InvestigationStatusTypeID,omitempty"`
	// Investigation Team Reviewer                
	InvestigationTeamReviewer          *string    `json:"InvestigationTeamReviewer,omitempty"`
}
