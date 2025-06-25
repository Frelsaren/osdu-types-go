package referencedata

// This entity is used to provide a mapping of external reference values to the current
// platform instance reference values. The scope can be global or specific to an external
// entity type. It can provide simple mappings or complex mappings, which maps the source
// value to multiple property values — well status and classification is an example for such
// complex mappings.
type Map struct {
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
