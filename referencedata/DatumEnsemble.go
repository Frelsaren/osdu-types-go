package referencedata

// The DatumEnsemble for the CRS's datum. Only populated for GeographicCRS.
type DatumEnsemble struct {
	// The DatumEnsemble authority code, corresponding to the ISO19111 ID and 'projjson' id.                            
	AuthorityCode                                                                           *DatumEnsembleAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the DatumEnsemble.                                                                                   
	Name                                                                                    *string                     `json:"Name,omitempty"`
}
