package referencedata

// The horizontal CRS reference of a CompoundCRS. Only populated for
// CoordinateReferenceSystemType==CompoundCRS.
type HorizontalCRS struct {
	// The HorizontalCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                            
	AuthorityCode                                                                           *HorizontalCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The relationship to the HorizontalCrs.                                                                           
	HorizontalCRSID                                                                         *string                     `json:"HorizontalCRSID,omitempty"`
	// The name of the HorizontalCrs.                                                                                   
	Name                                                                                    *string                     `json:"Name,omitempty"`
}
