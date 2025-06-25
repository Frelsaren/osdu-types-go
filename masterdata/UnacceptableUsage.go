package masterdata

// Describes the workflows and/or personas that the technical assurance value is not valid
// for (e.g., This data has a technical assurance property of "trusted", but it is not
// suitable for Seismic Interpretation).
type UnacceptableUsage struct {
	// The relationship to a work-product-component--DataQuality assessment record, which was            
	// used as the basis for this decision.                                                              
	DataQualityID                                                                                *string `json:"DataQualityID,omitempty"`
	// The DataQualityRuleSet, which did not pass successfully to achieve this level of                  
	// technical assurance.                                                                              
	DataQualityRuleSetID                                                                         *string `json:"DataQualityRuleSetID,omitempty"`
	// DEPRECATED: superseded by DataQualityRuleSetID referring to DataQualityRuleSet. The               
	// QualityDataRuleSet, which did not pass successfully to achieve this level of technical            
	// assurance.                                                                                        
	QualityDataRuleSetID                                                                         *string `json:"QualityDataRuleSetID,omitempty"`
	// The stage of business where the record is not acceptable for workflow usage.                      
	ValueChainStatusTypeID                                                                       *string `json:"ValueChainStatusTypeID,omitempty"`
	// DEPRECATED: superseded by WorkflowPersonaTypeID. Name of the role or personas that the            
	// record is technical assurance value is not valid for.                                             
	WorkflowPersona                                                                              *string `json:"WorkflowPersona,omitempty"`
	// Name of the role or personas that the record is technical assurance value is not valid            
	// for.                                                                                              
	WorkflowPersonaTypeID                                                                        *string `json:"WorkflowPersonaTypeID,omitempty"`
	// DEPRECATED: superseded by WorkflowUsageTypeID. Name of the business activities,                   
	// processes, and/or workflows that the record is technical assurance value is not valid for.        
	WorkflowUsage                                                                                *string `json:"WorkflowUsage,omitempty"`
	// Name of the business activities, processes, and/or workflows that the record's technical          
	// assurance value is not valid for.                                                                 
	WorkflowUsageTypeID                                                                          *string `json:"WorkflowUsageTypeID,omitempty"`
}
