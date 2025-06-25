package referencedata

// The Transformation authority code, corresponding to the ISO19111 ID and 'projjson' id.
type ConcatenatedTransformationAuthorityCode struct {
	// The transformation authority governing the 'Code'.          
	Authority                                              *string `json:"Authority,omitempty"`
	// The transformation code assigned by the 'Authority'.        
	Code                                                   *int64  `json:"Code,omitempty"`
}
