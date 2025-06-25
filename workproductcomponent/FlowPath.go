package workproductcomponent

import "time"

// Details about the downhole equipment design for this stimulation interval.
type FlowPath struct {
	// The downhole pressure at which the formation broke.                                               
	BreakDownPressure                                                                         *float64   `json:"BreakDownPressure,omitempty"`
	// The measured depth of a bridge plug.                                                              
	BridgePlugMeasuredDepth                                                                   *float64   `json:"BridgePlugMeasuredDepth,omitempty"`
	// Date that the bridge plug for this stage was removed.                                             
	BridgePlugRemovalDate                                                                     *time.Time `json:"BridgePlugRemovalDate,omitempty"`
	// The type of stim flow path.                                                                       
	FlowPathTypeID                                                                            *string    `json:"FlowPathTypeID,omitempty"`
	// The formation fracture gradient for this treatment flow path.                                     
	FractureGradient                                                                          *float64   `json:"FractureGradient,omitempty"`
	// The friction factor used to compute openhole pressure loss.                                       
	OpenHoleFrictionFactor                                                                    *float64   `json:"OpenHoleFrictionFactor,omitempty"`
	// The measured depth of a packer.                                                                   
	PackerMeasuredDepth                                                                       *float64   `json:"PackerMeasuredDepth,omitempty"`
	// The friction factor for the pipe, tubing, and/or casing.                                          
	PipeFrictionFactor                                                                        *float64   `json:"PipeFrictionFactor,omitempty"`
	// PMax prediction allows the tool assembly to be designed with expected pressures. It               
	// determines maximum allowable surface pressure and is typically calculated as a single             
	// number by which the pressure relief valves are set. This variable is the average of all           
	// the pmax pressures calculated for this flow path.                                                 
	PmaxPacPressureAvg                                                                        *float64   `json:"PmaxPacPressureAvg,omitempty"`
	// PMax prediction allows the tool assembly to be designed with expected pressures. It               
	// determines maximum allowable surface pressure and is typically calculated as a single             
	// number by which the pressure relief valves are set. This variable is the maximum of all           
	// the pmax pressures calculated for this flow path.                                                 
	PmaxPacPressureMax                                                                        *float64   `json:"PmaxPacPressureMax,omitempty"`
	// Average allowable pressure for the zone of interest with respect to the bottomhole                
	// assembly during the stimulation services.                                                         
	PmaxWeaklinkPressureAvg                                                                   *float64   `json:"PmaxWeaklinkPressureAvg,omitempty"`
	// Maximum allowable pressure for the zone of interest with respect to the bottomhole                
	// assembly during the stimulation services.                                                         
	PmaxWeaklinkPressureMax                                                                   *float64   `json:"PmaxWeaklinkPressureMax,omitempty"`
	// General remarks about the flow path for this job stage.                                           
	Remarks                                                                                   *string    `json:"Remarks,omitempty"`
	// The maximum measured depth of the tubing used for treatment of a stage.                           
	TubingBottomMeasuredDepth                                                                 *float64   `json:"TubingBottomMeasuredDepth,omitempty"`
}
