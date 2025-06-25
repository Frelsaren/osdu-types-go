package referencedata

// The instructions from where to derive the value.
type ValueExtraction struct {
	// The RelatedConditionProperty values, which need to match in order to be accepted as             
	// de-normalized value(s). If the Policy is ExtractFirstMatch, the list is prioritized and         
	// the first match is accepted as final value. Policy ExtractAllMatches collects all               
	// matching values as array.                                                                       
	RelatedConditionMatches                                                                   []string `json:"RelatedConditionMatches,omitempty"`
	// The property path of the target record data block, which needs to be subjected to the           
	// conditional matching. The data prefix is not required in the path.                              
	RelatedConditionProperty                                                                  *string  `json:"RelatedConditionProperty,omitempty"`
	// The path to the property from where to extract the de-normalized value. The data prefix         
	// is not required in the path.                                                                    
	ValuePath                                                                                 string   `json:"ValuePath"`
}
