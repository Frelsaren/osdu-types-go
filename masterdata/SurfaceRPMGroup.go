package masterdata

// The planned surface rotary speed values for the BHA run.
type SurfaceRPMGroup struct {
	// The planned maximum value for the considered parameter                                          
	MaximumParameter                                             []SurfaceRPMGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                          
	MinimumParameter                                             []SurfaceRPMGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                                      
	RecommendedParameter                                         []SurfaceRPMGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
