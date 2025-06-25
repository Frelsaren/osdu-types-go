package workproductcomponent

// Describes the workflows and/or personas that the technical assurance value is valid for
// (e.g., This data has a technical assurance property of "trusted" and it is suitable for
// Seismic Interpretation).
type FluffyAcceptableUsage struct {
	// The QualityDataRuleSet, which had to pass successfully to achieve this level of technical        
	// assurance.                                                                                       
	QualityDataRuleSetID                                                                        *string `json:"QualityDataRuleSetID,omitempty"`
	// The stage of business where the record is acceptable for workflow usage.                         
	ValueChainStatusTypeID                                                                      *string `json:"ValueChainStatusTypeID,omitempty"`
	// Name of the role or personas that the record is technical assurance value is valid for.          
	WorkflowPersona                                                                             *string `json:"WorkflowPersona,omitempty"`
	// Name of the business activities, processes, and/or workflows that the record is technical        
	// assurance value is valid for.                                                                    
	WorkflowUsage                                                                               *string `json:"WorkflowUsage,omitempty"`
}
