package masterdata

// A workflow configuration in the context of a scheduled job.
type Workflow struct {
	// Allows creation of a specific handler DAG for the source type                                               
	Handler                                                                                    *string             `json:"Handler,omitempty"`
	// Name value or object Id                                                                                     
	Parameters                                                                                 []WorkflowParameter `json:"Parameters,omitempty"`
	// Reference name for the security scheme in the ConnectedSourceRegistryEntry document this                    
	// scheduled job belongs to.                                                                                   
	SecuritySchemeName                                                                         string              `json:"SecuritySchemeName"`
	// Tag given to a workflow category, such as Fetch, Ingest, Delivery                                           
	Tag                                                                                        *string             `json:"Tag,omitempty"`
	// Data source endpoint used in workflow                                                                       
	URL                                                                                        *string             `json:"Url,omitempty"`
}
