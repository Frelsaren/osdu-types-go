package masterdata

// A group of parameters that refer to Torque on Bottom
type TorqueOnBottomGroup struct {
	// The planned maximum value for the considered parameter                                              
	MaximumParameter                                             []TorqueOnBottomGroupMaximumParameter     `json:"MaximumParameter,omitempty"`
	// The planned minimum value for the considered parameter                                              
	MinimumParameter                                             []TorqueOnBottomGroupMinimumParameter     `json:"MinimumParameter,omitempty"`
	// The planned recommended value for the considered parameter                                          
	RecommendedParameter                                         []TorqueOnBottomGroupRecommendedParameter `json:"RecommendedParameter,omitempty"`
}
