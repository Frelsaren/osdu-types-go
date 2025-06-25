package referencedata

// The Projection Operation authority code, corresponding to the ISO19111 ID and 'projjson'
// id.
type ProjectionOperationAuthorityCode struct {
	// The projection operation  authority governing the 'Code'.         
	Authority                                                    *string `json:"Authority,omitempty"`
	// The projection operation code assigned by the 'Authority'.        
	Code                                                         *int64  `json:"Code,omitempty"`
}
