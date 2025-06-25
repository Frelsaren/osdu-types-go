package workproductcomponent

// Defines relationships with other objects (any kind of Resource) upon which this work
// product component depends.  The assertion is directed only from the asserting WPC to
// ancestor objects, not children.  It should not be used to refer to files or artefacts
// within the WPC -- the association within the WPC is sufficient and Artefacts are actually
// children of the main WPC file. They should be recorded in the data.Artefacts[] array.
type LineageAssertion struct {
	// The object reference identifying the DIRECT, INDIRECT, REFERENCE dependency.                    
	ID                                                                                         *string `json:"ID,omitempty"`
	// Used by LineageAssertion to describe the nature of the line of descent of a work product        
	// component from a prior Resource, such as DIRECT, INDIRECT, REFERENCE.  It is not for            
	// proximity (number of nodes away), it is not to cover all the relationships in a full            
	// ontology or graph, and it is not to describe the type of activity that created the              
	// asserting WPC.  LineageAssertion does not encompass a full provenance, process history,         
	// or activity model.                                                                              
	LineageRelationshipType                                                                    *string `json:"LineageRelationshipType,omitempty"`
}
