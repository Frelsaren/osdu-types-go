package referencedata

// Scope and extent information about the described CRS.
type PurpleUsage struct {
	// The Usage authority code, corresponding to the ISO19111 ID and 'projjson' id.                          
	AuthorityCode                                                                   *PurpleUsageAuthorityCode `json:"AuthorityCode,omitempty"`
	// Extent or area of use information.                                                                     
	Extent                                                                          *PurpleExtent             `json:"Extent,omitempty"`
	// The name of the Usage.                                                                                 
	Name                                                                            *string                   `json:"Name,omitempty"`
	// Scope declaration.                                                                                     
	Scope                                                                           *PurpleScope              `json:"Scope,omitempty"`
}
