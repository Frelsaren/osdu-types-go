package referencedata

// Scope and extent information about the described transformation.
type FluffyPreferredUsage struct {
	// The Preferred Usage authority code, corresponding to the ISO19111 ID and 'projjson' id.                                   
	AuthorityCode                                                                             *FluffyPreferredUsageAuthorityCode `json:"AuthorityCode,omitempty"`
	// Extent or area of use information.                                                                                        
	Extent                                                                                    *FluffyPreferredExtent             `json:"Extent,omitempty"`
	// The name of the Usage.                                                                                                    
	Name                                                                                      *string                            `json:"Name,omitempty"`
	Scope                                                                                     *FluffyPreferredScope              `json:"Scope,omitempty"`
}
