package masterdata

// A group of parameters that refer to flowrate
type FlowratePumpGroup struct {
	// The planned maximum value for the considered parameter                                            
	MaximumParameter                                             []FlowratePumpGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                            
	MinimumParameter                                             []FlowratePumpGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                                        
	RecommendedParameter                                         []FlowratePumpGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
