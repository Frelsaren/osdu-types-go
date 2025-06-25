package referencedata

// The source CRS of a BoundCRS. Only populated for CoordinateReferenceSystemType==BoundCRS.
type PurpleSourceCRS struct {
	// The SourceCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                              
	AuthorityCode                                                                       *PurpleSourceCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Source CRS.                                                                                    
	Name                                                                                *string                       `json:"Name,omitempty"`
	// The relationship to the source CoordinateRefSystem.                                                            
	SourceCRSID                                                                         *string                       `json:"SourceCRSID,omitempty"`
}
