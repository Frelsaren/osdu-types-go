package referencedata

// The method authority code, corresponding to the ISO19111 ID and 'projjson' id.
type MethodAuthorityCode struct {
	// The authority governing the 'Code'.                                               
	Authority                                                                    *string `json:"Authority,omitempty"`
	// EPSG Method code in case of CoordinateTransformationType == Transformation        
	Code                                                                         *int64  `json:"Code,omitempty"`
}
