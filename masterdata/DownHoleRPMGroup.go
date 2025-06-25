package masterdata

// The planned downhole rotary speed values for the BHA run.
type DownHoleRPMGroup struct {
	// The planned maximum value for the considered parameter                                           
	MaximumParameter                                             []DownHoleRPMGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                           
	MinimumParameter                                             []DownHoleRPMGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                                       
	RecommendedParameter                                         []DownHoleRPMGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
