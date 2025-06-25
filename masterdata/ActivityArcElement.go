package masterdata

// An 'arc' establishing a link between a single pair of ActivityTemplate input and output
// parameters.
type ActivityArcElement struct {
	// The relationship to the ActivityTemplate first in a pair of ActivityTemplates in the        
	// context of a workflow.                                                                      
	PrecedingActivityTemplateID                                                             string `json:"PrecedingActivityTemplateID"`
	// The relationship to the ActivityTemplate succeeding the first ActivityTemplate in the       
	// workflow.                                                                                   
	SucceedingActivityTemplateID                                                            string `json:"SucceedingActivityTemplateID"`
}
