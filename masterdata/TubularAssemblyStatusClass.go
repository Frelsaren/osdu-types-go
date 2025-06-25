package masterdata

import "time"

// Reflects the current status of the Assembly - as 'installed', 'pulled', 'planned',... -
// Applicable to tubing/completions as opposed to drillstrings. Historical states are
// recorded in TubularAssemblyStates.
type TubularAssemblyStatusClass struct {
	// Date the status has been established.                                                     
	Date                                                                              *time.Time `json:"Date,omitempty"`
	// Used to describe the reason of Activity - such as installed, cut/pull, pulling.           
	Description                                                                       *string    `json:"Description,omitempty"`
	// The relationship to the  reference value status type.                                     
	StatusTypeID                                                                      *string    `json:"StatusTypeID,omitempty"`
}
