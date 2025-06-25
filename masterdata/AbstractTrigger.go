package masterdata

// A condition, which causes a re-evaluation of the decision and optionally the re-execution
// of a workflow.
type AbstractTrigger struct {
	// The condition expressed as human readable text, which triggers the re-evaluation of the                   
	// decision.                                                                                                 
	Condition                                                                                   *string          `json:"Condition,omitempty"`
	// An array of free remarks or annotations.                                                                  
	Remarks                                                                                     []AbstractRemark `json:"Remarks,omitempty"`
	// Relationships to zero or more activity templates representing workflows, which need to be                 
	// executed when the condition is met.                                                                       
	WorkflowIDs                                                                                 []string         `json:"WorkflowIDs,omitempty"`
}
