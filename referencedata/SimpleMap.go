package referencedata

// The external reference value is mapped to a single OSDU reference value in the target
// OSDU platform instance.
type SimpleMap struct {
	// Optional, needed if the Scope is not Global: PropertyName defines the cumulative path            
	// (dot-separated for nested structures, [] denoting arrays), to which the PropertyValue is         
	// to be assigned in the target record of kind TargetKind.                                          
	PropertyName                                                                                *string `json:"PropertyName,omitempty"`
	// Mandatory: the mapped value (reference-data relationship) in the OSDU target platform            
	// instance.                                                                                        
	ReferenceValueID                                                                            string  `json:"ReferenceValueID"`
	// Optional, needed if the Scope is not Global: TargetKind defines the record kind in which         
	// the PropertyValue is assigned to the PropertyName. The kind does not require the                 
	// specification of the full semantic version number. If specified, it denotes the first and        
	// implicitly higher versions, which are required, typically the version the PropertyValue          
	// was added.                                                                                       
	TargetKind                                                                                  *string `json:"TargetKind,omitempty"`
}
