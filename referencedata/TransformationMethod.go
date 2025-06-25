package referencedata

// The Transformation method; "Concatenated" for CoordinateTransformationType ==
// ConcatenatedOperation; EPSG method code and name for CoordinateTransformationType ==
// Transformation.
type TransformationMethod struct {
	// The method authority code, corresponding to the ISO19111 ID and 'projjson' id.                           
	AuthorityCode                                                                          *MethodAuthorityCode `json:"AuthorityCode,omitempty"`
	// The Transformation method name; "Concatenated" for CoordinateTransformationType ==                       
	// ConcatenatedOperation; EPSG method code and name for CoordinateTransformationType ==                     
	// Transformation.                                                                                          
	Name                                                                                   *string              `json:"Name,omitempty"`
}
