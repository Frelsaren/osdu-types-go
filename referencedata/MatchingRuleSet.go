package referencedata

// One set of  conditions, which have to be matched to determine sufficient equality of an
// object.
type MatchingRuleSet struct {
	// The list of conditions, which all have to be met in order to declare a match.                           
	Conditions                                                                              []Condition        `json:"Conditions,omitempty"`
	// A text explaining the purpose or goal of this MatchingRuleSet.                                          
	Description                                                                             *string            `json:"Description,omitempty"`
	// DEPRECATED: Combined with Conditions, no longer in use. The specifications for system                   
	// attribute replacements.                                                                                 
	ReplaceAttributes                                                                       []ReplaceAttribute `json:"ReplaceAttributes,omitempty"`
	// The name of the rule (expected to be unique in the array of MatchingRuleSets[]).                        
	RuleName                                                                                *string            `json:"RuleName,omitempty"`
}
