package masterdata

// A series of operating parameters observed during the run
type OperationParameterPlan struct {
	// A group of parameters that refer to flowrate                                                  
	FlowratePumpGroup                                                          *FlowratePumpGroup    `json:"FlowratePumpGroup,omitempty"`
	// The realization strategy utilized in this series of operation parameters                      
	RealizationStrategy                                                        *string               `json:"RealizationStrategy,omitempty"`
	// A group of parameters that refer to ROP (rate of penetration)                                 
	ROPGroup                                                                   *ROPGroup             `json:"ROPGroup,omitempty"`
	// A group of parameters that refer to RPM (rotations per minute)                                
	RPMGroup                                                                   *RPMGroup             `json:"RPMGroup,omitempty"`
	// A group of parameters that refer to anticipated Torque at Surface                             
	TorqueAtSurfaceGroup                                                       *TorqueAtSurfaceGroup `json:"TorqueAtSurfaceGroup,omitempty"`
	// A group of parameters that refer to Torque on Bottom                                          
	TorqueOnBottomGroup                                                        *TorqueOnBottomGroup  `json:"TorqueOnBottomGroup,omitempty"`
	// A group of parameters that refer to WOB (weight on bit)                                       
	WOBGroup                                                                   *WOBGroup             `json:"WOBGroup,omitempty"`
}
