package masterdata

import "time"

// The history of expenditure approvals.
type FundsAuthorizations struct {
	// Internal Company control number which identifies the allocation of funds to the Project.           
	AuthorizationID                                                                            *string    `json:"AuthorizationID,omitempty"`
	// Type of currency for the authorized expenditure.                                                   
	CurrencyID                                                                                 *string    `json:"CurrencyID,omitempty"`
	// The date and time when the funds were approved.                                                    
	EffectiveDateTime                                                                          *time.Time `json:"EffectiveDateTime,omitempty"`
	// The level of expenditure approved.                                                                 
	FundsAmount                                                                                *float64   `json:"FundsAmount,omitempty"`
}
