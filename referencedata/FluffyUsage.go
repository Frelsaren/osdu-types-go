package referencedata

// Scope and extent information about the described transformation.
type FluffyUsage struct {
	// The Usage authority code, corresponding to the ISO19111 ID and 'projjson' id.                          
	AuthorityCode                                                                   *FluffyUsageAuthorityCode `json:"AuthorityCode,omitempty"`
	// Extent or area of use information.                                                                     
	Extent                                                                          *FluffyExtent             `json:"Extent,omitempty"`
	// The name of the Usage.                                                                                 
	Name                                                                            *string                   `json:"Name,omitempty"`
	Scope                                                                           *FluffyScope              `json:"Scope,omitempty"`
}
