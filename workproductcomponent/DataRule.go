package workproductcomponent

// The DataQualityRule and the DataQualityRule properties captured at the time of assessment.
type DataRule struct {
	// The assessment results for this this rule.                                                                
	AssessmentResults                                                                         []AssessmentResult `json:"AssessmentResults,omitempty"`
	// The reference to the individual DataQualityRule record ID. This is a mandatory property                   
	// value.                                                                                                    
	DataQualityRuleID                                                                         string             `json:"DataQualityRuleID"`
	// The rule's evaluation status, e.g., Active, Inactive, Development.                                        
	DataQualityRuleStatusID                                                                   *string            `json:"DataQualityRuleStatusID,omitempty"`
	// The rule's dimension type reference as captured during the assessment.                                    
	DataRuleDimensionTypeID                                                                   *string            `json:"DataRuleDimensionTypeID,omitempty"`
	// The rule's purpose in the context of this DataQualityRule as captured during the                          
	// assessment.                                                                                               
	DataRulePurposeTypeID                                                                     *string            `json:"DataRulePurposeTypeID,omitempty"`
	// The rule's statement in natural language as captured during the assessment.                               
	DataRuleStatement                                                                         *string            `json:"DataRuleStatement,omitempty"`
	// The time it took to perform the assessment of this rule.                                                  
	ExecutionTime                                                                             *float64           `json:"ExecutionTime,omitempty"`
	// The rule's weight in the context of the DataQualityRuleSet as captured during the                         
	// assessment.                                                                                               
	Weight                                                                                    *float64           `json:"Weight,omitempty"`
}
