package referencedata

// One single configuration to derive an Search index property value and assign it to the
// index 'column' with Name.
type PurpleConfiguration struct {
	// The name of the indexed property, i.e., this is the property name used in Search.                
	Name                                                                                        *string `json:"Name,omitempty"`
	// The list of path definitions to derive the property value from.                                  
	Paths                                                                                       []Path  `json:"Paths,omitempty"`
	// Current supported policies are 'ExtractAllMatches' resulting in an array of values or            
	// 'ExtractFirstMatch' single value. The policy applies only to the Paths[].ValueExtraction.        
	Policy                                                                                      *string `json:"Policy,omitempty"`
	// The use case description this configuration satisfies.                                           
	UseCase                                                                                     *string `json:"UseCase,omitempty"`
}
