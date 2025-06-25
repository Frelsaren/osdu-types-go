package referencedata

// DEPRECATED: No longer in use. A source and target system attribute specification to
// replace values.
type ReplaceAttribute struct {
	// DEPRECATED: No longer in use. An optional descriptive remark explaining the purpose of            
	// the replacement.                                                                                  
	Description                                                                                  *string `json:"Description,omitempty"`
	// DEPRECATED: No longer in use. The source system attribute delivering the value to replace.        
	SourceSystemAttribute                                                                        *string `json:"SourceSystemAttribute,omitempty"`
	// DEPRECATED: No longer in use. The target system attribute receiving the value to replace.         
	TargetSystemAttribute                                                                        *string `json:"TargetSystemAttribute,omitempty"`
}
