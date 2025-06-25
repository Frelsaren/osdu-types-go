package referencedata

// The Transformation bound to the BaseCRS in a BoundCRS. Only populated for
// CoordinateReferenceSystemType==BoundCRS.
type BoundTransformation struct {
	// The Transformation authority code, corresponding to the ISO19111 ID and 'projjson' id.                             
	AuthorityCode                                                                            *TransformationAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Transformation.                                                                                    
	Name                                                                                     *string                      `json:"Name,omitempty"`
	// The relationship to the bound transformation.                                                                      
	TransformationID                                                                         *string                      `json:"TransformationID,omitempty"`
}
