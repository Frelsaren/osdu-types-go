package masterdata

// Parameter name value pair.
type WorkflowParameter struct {
	// Name of the parameter the handler expects         
	Name                                         *string `json:"Name,omitempty"`
	// Value of the parameter the handler expects        
	Value                                        *string `json:"Value,omitempty"`
}
