package referencedata

// The base geographic CRS of this projected CRS. Only populated for
// CoordinateReferenceSystemType==ProjectedCRS.
type BaseCRS struct {
	// The BaseCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                      
	AuthorityCode                                                                     *BaseCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The relationship to the BaseCRS.                                                                     
	BaseCRSID                                                                         *string               `json:"BaseCRSID,omitempty"`
	// The name of the base CRS.                                                                            
	Name                                                                              *string               `json:"Name,omitempty"`
}
