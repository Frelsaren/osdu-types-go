package masterdata

// Describe a set of action to prevent for a risk or to mitigate its consequences
type MitigationElement struct {
	// Date the response must be completed                                                                      
	Deadline                                                                                  *float64          `json:"Deadline,omitempty"`
	// General description of the response                                                                      
	Description                                                                               string            `json:"Description"`
	// Name of the response                                                                                     
	Name                                                                                      string            `json:"Name"`
	// List of the staff responsible to proceed with the response. This is the 'by value'                       
	// alternative to the 'by reference' data.Mitigations[].ResponsiblesByReferenceIDs[].                       
	Responsibles                                                                              []AbstractContact `json:"Responsibles"`
	// List of the staff responsible to proceed with the response as references to UserProfile                  
	// records. This is the 'by reference' alternative to 'by value' in                                         
	// data.Mitigations[].Responsibles[].                                                                       
	ResponsiblesByReferenceIDs                                                                []string          `json:"ResponsiblesByReferenceIDs,omitempty"`
	// Describes the status of the action such as (progress, done, canceled)                                    
	Status                                                                                    string            `json:"Status"`
	// Date the description of the response has been updated.                                                   
	UpdateDate                                                                                *float64          `json:"UpdateDate,omitempty"`
}
