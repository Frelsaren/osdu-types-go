package referencedata

// An exact or approximate condition of equality, which is a match requirement.
type Condition struct {
	// A descriptive remark about the purpose of the condition.                                      
	Description                                                                             *string  `json:"Description,omitempty"`
	// The optional, numerical tolerance to be used when comparing numbers.                          
	NumericalTolerance                                                                      *float64 `json:"NumericalTolerance,omitempty"`
	// This value is used to identify the source property that needs replacing in the source         
	// record (see data.SourceKind).                                                                 
	ReplaceProperty                                                                         *string  `json:"ReplaceProperty,omitempty"`
	// The source property name to evaluate in the scope of the connected data source.               
	SourceProperty                                                                          *string  `json:"SourceProperty,omitempty"`
	// The target kind to query for.                                                                 
	TargetKind                                                                              *string  `json:"TargetKind,omitempty"`
	// The target property path to search TargetKind instances with the value of the                 
	// SourceProperty                                                                                
	TargetProperty                                                                          *string  `json:"TargetProperty,omitempty"`
}
