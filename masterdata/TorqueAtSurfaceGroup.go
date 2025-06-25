package masterdata

// A group of parameters that refer to anticipated Torque at Surface
type TorqueAtSurfaceGroup struct {
	// The planned maximum value for the considered parameter                                               
	MaximumParameter                                             []TorqueAtSurfaceGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                               
	MinimumParameter                                             []TorqueAtSurfaceGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                                           
	RecommendedParameter                                         []TorqueAtSurfaceGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
