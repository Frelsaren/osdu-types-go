package referencedata

// The source CRS of the Transformation.
type FluffySourceCRS struct {
	// The source CRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                              
	AuthorityCode                                                                        *FluffySourceCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the source CRS.                                                                                     
	Name                                                                                 *string                       `json:"Name,omitempty"`
	// The relationship to the source CoordinateReferenceSystem.                                                       
	SourceCRSID                                                                          *string                       `json:"SourceCRSID,omitempty"`
}
