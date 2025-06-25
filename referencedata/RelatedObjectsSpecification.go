package referencedata

// The specification to extract related objects, from which to derive the ValueExtraction.
// If this property is empty or absent, the ValueExtraction is done on the current object to
// be indexed.
type RelatedObjectsSpecification struct {
	// The RelatedConditionProperty values, which need to match in order to be accepted as               
	// de-normalized value(s). If the Policy is ExtractFirstMatch, the list is prioritized and           
	// the first match is accepted as final value. Policy ExtractAllMatches collects all                 
	// matching values as array.                                                                         
	RelatedConditionMatches                                                                     []string `json:"RelatedConditionMatches,omitempty"`
	// The property path of the target record data block, which needs subjected to the                   
	// conditional matching. The data prefix is not required.                                            
	RelatedConditionProperty                                                                    *string  `json:"RelatedConditionProperty,omitempty"`
	// The path to the property containing the ID of the target record to chase. This property           
	// is only populated if the property is extracted from a related object, which must be               
	// chased. If the property is derived from 'within' the same record, which triggered the             
	// indexing, the RelatedObjectID is left absent.                                                     
	RelatedObjectID                                                                             *string  `json:"RelatedObjectID,omitempty"`
	// The kind or schema id expected as the target object type. This property is only populated         
	// if the property is extracted from a related object, which must be chased. If the property         
	// is derived from 'within' the same record, which triggered the indexing, the                       
	// RelatedObjectKind is left absent.                                                                 
	RelatedObjectKind                                                                           *string  `json:"RelatedObjectKind,omitempty"`
	// The direction of the relationship definition seen from the object being indexed.                  
	// 'ChildToParent' assumes an outgoing relationship with the target record defined in the            
	// object being indexed. 'ParentToChildren' assumes that the related objects have a                  
	// relationship by RelatedObjectID to the id of the record being indexed.                            
	RelationshipDirection                                                                       *string  `json:"RelationshipDirection,omitempty"`
}
