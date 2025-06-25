package workproductcomponent

import "time"

// Record that reflects the status of the Assembly - as 'installed', 'pulled', 'planned',...
// - Applicable to tubing/completions as opposed to drillstrings
type TubularAssemblyStateElement struct {
	// Date the status has been established                                           
	Date                                                                   *time.Time `json:"Date,omitempty"`
	// Used to describe the reason of Activity - such as cut/pull, pulling,           
	Description                                                            *string    `json:"Description,omitempty"`
	// SRN of a reference value status type                                           
	StatusTypeID                                                           *string    `json:"StatusTypeID,omitempty"`
}
