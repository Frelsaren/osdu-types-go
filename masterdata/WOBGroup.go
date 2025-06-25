package masterdata

// A group of parameters that refer to WOB (weight on bit)
type WOBGroup struct {
	// The planned maximum value for the considered parameter                                   
	MaximumParameter                                             []WOBGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                   
	MinimumParameter                                             []WOBGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                               
	RecommendedParameter                                         []WOBGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
