package referencedata

// The SourceCRS authority code, corresponding to the ISO19111 ID and 'projjson' id.
type PurpleSourceCRSAuthorityCode struct {
	// The authority governing the 'Code'.          
	Authority                               *string `json:"Authority,omitempty"`
	// The code assigned by the 'Authority'.        
	Code                                    *int64  `json:"Code,omitempty"`
}
