package referencedata

// The datum of this CRS. Only populated for CoordinateReferenceSystemType in
// [GeographicCRS, VerticalCRS, EngineeringCRS].
type Datum struct {
	// The Datum authority code, corresponding to the ISO19111 ID and 'projjson' id.                    
	AuthorityCode                                                                   *DatumAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Datum.                                                                           
	Name                                                                            *string             `json:"Name,omitempty"`
}
