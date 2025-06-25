package referencedata

// Scope and extent information about the described CRS.
type PurplePreferredUsage struct {
	// The PreferredUsage authority code, corresponding to the ISO19111 ID and 'projjson' id.                                   
	AuthorityCode                                                                            *PurplePreferredUsageAuthorityCode `json:"AuthorityCode,omitempty"`
	// Extent or area of use information.                                                                                       
	Extent                                                                                   *PurplePreferredExtent             `json:"Extent,omitempty"`
	// The name of the Usage.                                                                                                   
	Name                                                                                     *string                            `json:"Name,omitempty"`
	// Scope declaration.                                                                                                       
	Scope                                                                                    *PurplePreferredScope              `json:"Scope,omitempty"`
}
