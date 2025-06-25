package referencedata

type ConcatenatedTransformation struct {
	// The Transformation authority code, corresponding to the ISO19111 ID and 'projjson' id.                                         
	AuthorityCode                                                                            *ConcatenatedTransformationAuthorityCode `json:"AuthorityCode,omitempty"`
	// The Transformation name as part of this concatenated operation list.                                                           
	Name                                                                                     *string                                  `json:"Name,omitempty"`
	// The relationship to the single Transformation item in the list of concatenated                                                 
	// transformations.                                                                                                               
	TransformationID                                                                         *string                                  `json:"TransformationID,omitempty"`
}
