package masterdata

// A list of Resources that are governed by the agreement.  Note that different terms may
// apply to different resources, but that granularity is handled by the Entitlements
// Rulebase.
type RestrictedResources struct {
	// Reference to an information Resource which is governed by the agreement.        
	ResourceID                                                                 *string `json:"ResourceID,omitempty"`
}
