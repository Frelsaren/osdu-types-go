package workproductcomponent

import "time"

// Individual Non Productive Time Action Description
type NonProductiveTimeAction struct {
	// Action Completed Date           
	ActionCompletedDate     *string    `json:"ActionCompletedDate,omitempty"`
	// Action Description              
	ActionDescription       *string    `json:"ActionDescription,omitempty"`
	// Action Owner name               
	ActionOwner             *string    `json:"ActionOwner,omitempty"`
	// Action Remarks                  
	ActionRemarks           *string    `json:"ActionRemarks,omitempty"`
	// Action Title                    
	ActionTitle             *time.Time `json:"ActionTitle,omitempty"`
}
