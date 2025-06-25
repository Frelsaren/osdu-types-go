package referencedata

// The coordinate system defining the dimension and individual axes used by the CRS.
type CoordinateSystem struct {
	// The CoordinateSystem authority code, corresponding to the ISO19111 ID and 'projjson' id.                               
	AuthorityCode                                                                              *CoordinateSystemAuthorityCode `json:"AuthorityCode,omitempty"`
	// The horizontal unit for 2-dimensional or 3-dimensional coordinate systems. This property                               
	// is not populated for 1-dimensional coordinate systems (Vertical CRSs).                                                 
	HorizontalAxisUnitID                                                                       *string                        `json:"HorizontalAxisUnitID,omitempty"`
	// The name of the Coordinate System.                                                                                     
	Name                                                                                       *string                        `json:"Name,omitempty"`
	// The vertical unit for 1-dimensional or 3-dimensional coordinate systems.                                               
	VerticalAxisUnitID                                                                         *string                        `json:"VerticalAxisUnitID,omitempty"`
}
