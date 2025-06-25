package workproductcomponent

// Describes the workflows and/or personas that the technical assurance value is not valid
// for (e.g., This data has a technical assurance property of "trusted", but it is not
// suitable for Seismic Interpretation).
type FluffyUnacceptableUsage struct {
	// The QualityDataRuleSet, which did not pass successfully to achieve this level of                 
	// technical assurance.                                                                             
	QualityDataRuleSetID                                                                        *string `json:"QualityDataRuleSetID,omitempty"`
	// The stage of business where the record is not acceptable for workflow usage.                     
	ValueChainStatusTypeID                                                                      *string `json:"ValueChainStatusTypeID,omitempty"`
	// Name of the role or personas that the record is technical assurance value is not valid           
	// for.                                                                                             
	WorkflowPersona                                                                             *string `json:"WorkflowPersona,omitempty"`
	// Name of the business activities, processes, and/or workflows that the record is technical        
	// assurance value is not valid for.                                                                
	WorkflowUsage                                                                               *string `json:"WorkflowUsage,omitempty"`
}
