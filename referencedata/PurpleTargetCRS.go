package referencedata

// The target CRS of this bound CRS. Only populated for
// CoordinateReferenceSystemType==BoundCRS.
type PurpleTargetCRS struct {
	// The TargetCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                              
	AuthorityCode                                                                       *PurpleTargetCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the base CRS.                                                                                      
	Name                                                                                *string                       `json:"Name,omitempty"`
	// The relationship to the TargetCRS.                                                                             
	TargetCRSID                                                                         *string                       `json:"TargetCRSID,omitempty"`
}
