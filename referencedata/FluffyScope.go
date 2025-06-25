package referencedata

type FluffyScope struct {
	// The Scope authority code, corresponding to the ISO19111 ID and 'projjson' id.                          
	AuthorityCode                                                                   *FluffyScopeAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Scope.                                                                                 
	Name                                                                            *string                   `json:"Name,omitempty"`
}
