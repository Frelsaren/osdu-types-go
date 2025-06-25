package referencedata

// Scope declaration.
type PurpleScope struct {
	// The Scope authority code, corresponding to the ISO19111 ID and 'projjson' id.                          
	AuthorityCode                                                                   *PurpleScopeAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Scope.                                                                                 
	Name                                                                            *string                   `json:"Name,omitempty"`
}
