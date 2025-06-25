package referencedata

// Scope declaration.
type PurplePreferredScope struct {
	// The Preferred Scope authority code, corresponding to the ISO19111 ID and 'projjson' id.                                   
	AuthorityCode                                                                             *PurplePreferredScopeAuthorityCode `json:"AuthorityCode,omitempty"`
	// The name of the Scope.                                                                                                    
	Name                                                                                      *string                            `json:"Name,omitempty"`
}
