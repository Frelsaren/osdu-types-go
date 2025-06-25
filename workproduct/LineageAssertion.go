package workproduct

type LineageAssertion struct {
	// The object reference identifying the DIRECT, INDIRECT, REFERENCE dependency.                      
	ID                                                                                           *string `json:"ID,omitempty"`
	// Used by LineageAssertion to describe the nature of the line of descent of a work product          
	// from a prior Resource, such as DIRECT, INDIRECT, REFERENCE.  It is not for proximity              
	// (number of nodes away), it is not to cover all the relationships in a full ontology or            
	// graph, and it is not to describe the type of activity that created the asserting WP.              
	// LineageAssertion does not encompass a full provenance, process history, or activity model.        
	LineageRelationshipType                                                                      *string `json:"LineageRelationshipType,omitempty"`
}
