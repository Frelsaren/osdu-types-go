package referencedata

// The target CRS of this Transformation.
type FluffyTargetCRS struct {
	// The target CRS authority code, corresponding to the ISO19111 ID and 'projjson' id.                              
	AuthorityCode                                                                        *FluffyTargetCRSAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the target CRS.                                                                                     
	Name                                                                                 *string                       `json:"Name,omitempty"`
	// The relationship to the target CRS.                                                                             
	TargetCRSID                                                                          *string                       `json:"TargetCRSID,omitempty"`
}
