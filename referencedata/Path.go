package referencedata

// A single path definition to derive a property value from.
type Path struct {
	// The specification to extract related objects, from which to derive the ValueExtraction.                               
	// If this property is empty or absent, the ValueExtraction is done on the current object to                             
	// be indexed.                                                                                                           
	RelatedObjectsSpec                                                                          *RelatedObjectsSpecification `json:"RelatedObjectsSpec,omitempty"`
	// The instructions from where to derive the value.                                                                      
	ValueExtraction                                                                             ValueExtraction              `json:"ValueExtraction"`
}
