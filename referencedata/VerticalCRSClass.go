package referencedata

// The vertical CRS reference of a CompoundCRS. Only populated for
// CoordinateReferenceSystemType==CompoundCRS.
type VerticalCRSClass struct {
	// The VerticalCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                          
	AuthorityCode                                                                         *VerticalCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the VerticalCrs.                                                                                 
	Name                                                                                  *string                   `json:"Name,omitempty"`
	// The relationship to the VerticalCrs.                                                                         
	VerticalCRSID                                                                         *string                   `json:"VerticalCRSID,omitempty"`
}
