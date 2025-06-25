package masterdata

import "time"

// A list of obligations or allowed activities specified by the agreement that apply to
// stored resources.  These are translated into rules, which the Entitlement Rulebase
// enforces.  Each rule should reference the agreement it codifies.
type Terms struct {
	// The Date when the obligation no longer needs to be fulfilled.                                      
	EndDate                                                                                    *time.Time `json:"EndDate,omitempty"`
	// Lengthy description of a legal restriction imposed on data governed by the agreement.              
	ObligationDescription                                                                      *string    `json:"ObligationDescription,omitempty"`
	// Reference to the general class of obligation, such as nondisclosure, termination of use,           
	// non-assignment, export restriction, limitation on derivatives.                                     
	ObligationTypeID                                                                           *string    `json:"ObligationTypeID,omitempty"`
	// The Date when the obligation becomes enforceable.                                                  
	StartDate                                                                                  *time.Time `json:"StartDate,omitempty"`
}
