package masterdata

// A group of parameters that refer to ROP (rate of penetration)
type ROPGroup struct {
	// The planned maximum value for the considered parameter                                   
	MaximumParameter                                             []ROPGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                   
	MinimumParameter                                             []ROPGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                               
	RecommendedParameter                                         []ROPGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
