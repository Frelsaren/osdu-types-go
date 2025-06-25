package masterdata

// Provides the overall description of the drilling fluids system.
type FluidSystem struct {
	// Description of the formulation to be used for the drilling activity                                     
	BarrelFormulation                                                                      []BarrelFormulation `json:"BarrelFormulation,omitempty"`
	// An array of planned / designed properties of the fluid for the Interval / Phase.                        
	// Properties may have a specified value or a range that should be maintained                              
	FluidProperties                                                                        []FluidsProperty    `json:"FluidProperties"`
	// The purpose the mud will play in this hole section (can be Sweep, Displacement Mud).                    
	FluidPurposeID                                                                         *string             `json:"FluidPurposeID,omitempty"`
	// Free text string of the common name or product name of the drilling mud.                                
	FluidSystemName                                                                        *string             `json:"FluidSystemName,omitempty"`
	// Type of polymers present in mud system.                                                                 
	PolymerType                                                                            *string             `json:"PolymerType,omitempty"`
}
