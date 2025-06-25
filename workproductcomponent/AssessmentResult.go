package workproductcomponent

// The assessment of an executed QualityDataRule.
type AssessmentResult struct {
	// The unique identifier of an assessment result in the AssessmentResults[] array. This                           
	// identifier becomes relevant if the assessments run over arrays of objects. Typically the                       
	// 'local object identifier' of the assessed object will be used. Example: WellLog has                            
	// Curves[], then Curves[].CurveID would be used as AssessedPartID.                                               
	AssessedPartID                                                                             *string                `json:"AssessedPartID,omitempty"`
	// An extension point to add properties, which are not searchable.                                                
	ExtensionProperties                                                                        map[string]interface{} `json:"ExtensionProperties,omitempty"`
	// The assessment state for this rule.                                                                            
	QualityAssessmentStateID                                                                   *string                `json:"QualityAssessmentStateID,omitempty"`
}
