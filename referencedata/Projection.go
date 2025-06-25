package referencedata

// The projection operation of a ProjectedCRS. Only populated for
// CoordinateReferenceSystemType==ProjectedCRS.
type Projection struct {
	// The Projection Operation authority code, corresponding to the ISO19111 ID and 'projjson'                                  
	// id.                                                                                                                       
	AuthorityCode                                                                              *ProjectionOperationAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the projection.                                                                                               
	Name                                                                                       *string                           `json:"Name,omitempty"`
}
