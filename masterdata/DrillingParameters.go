package masterdata

import "time"

// The bottomhole assembly drilling parameters schema, which contains statistical and
// calculated operations data for the run, related to depths, activities, temperature,
// pressure, flow rates, torque, etc.
type DrillingParameters struct {
	// The average flowrate at the bit during the run                                            
	AverageBitFlowrate                                                                *float64   `json:"AverageBitFlowrate,omitempty"`
	// Average mud pump flow rate.                                                               
	AverageFlowratePump                                                               *float64   `json:"AverageFlowratePump,omitempty"`
	// The average pump pressure used during the run                                             
	AveragePumpPressure                                                               *float64   `json:"AveragePumpPressure,omitempty"`
	// Average rate of penetration through Interval.                                             
	AverageROP                                                                        *float64   `json:"AverageROP,omitempty"`
	// Average surface turn rate (commonly in rpm) through Interval.                             
	AverageRPM                                                                        *float64   `json:"AverageRpm,omitempty"`
	// Average turn rate (commonly in rpm) downhole.                                             
	AverageRPMDownhole                                                                *float64   `json:"AverageRpmDownhole,omitempty"`
	// Average torque: downhole.                                                                 
	AverageTorqueDownhole                                                             *float64   `json:"AverageTorqueDownhole,omitempty"`
	// Average torque: off bottom.                                                               
	AverageTorqueOffBottom                                                            *float64   `json:"AverageTorqueOffBottom,omitempty"`
	// Average Torque: on bottom.                                                                
	AverageTorqueOnBottom                                                             *float64   `json:"AverageTorqueOnBottom,omitempty"`
	// Bit nozzle average velocity.                                                              
	AverageVelocityBitNozzle                                                          *float64   `json:"AverageVelocityBitNozzle,omitempty"`
	// Surface weight on bit - average through interval.                                         
	AverageWOB                                                                        *float64   `json:"AverageWOB,omitempty"`
	// Weight on bit - average downhole.                                                         
	AverageWOBDownhole                                                                *float64   `json:"AverageWOBDownhole,omitempty"`
	// Azimuth at stop measured depth.                                                           
	AzimuthBottom                                                                     *float64   `json:"AzimuthBottom,omitempty"`
	// Azimuth at start measured depth.                                                          
	AzimuthTop                                                                        *float64   `json:"AzimuthTop,omitempty"`
	// Operating time spent by bit for run.                                                      
	//                                                                                           
	// BUSINESS RULE: When reporting an actual as opposed to design, this is required.           
	BitRunTIME                                                                        *time.Time `json:"BitRunTIme,omitempty"`
	// Comments and remarks.                                                                     
	Comments                                                                          *string    `json:"Comments,omitempty"`
	// Distance drilled - rotating.                                                              
	DistanceDrilledRotating                                                           *float64   `json:"DistanceDrilledRotating,omitempty"`
	// Distance drilled - sliding                                                                
	DistanceDrilledSliding                                                            *float64   `json:"DistanceDrilledSliding,omitempty"`
	// Distance covered while holding angle with a steerable drilling assembly.                  
	DistanceHolding                                                                   *float64   `json:"DistanceHolding,omitempty"`
	// Distance reamed.                                                                          
	DistanceReamed                                                                    *float64   `json:"DistanceReamed,omitempty"`
	// Distance covered while actively steering with a steerable drilling assembly.              
	DistanceSteering                                                                  *float64   `json:"DistanceSteering,omitempty"`
	// Hookload when the string is moving down.                                                  
	HookloadDown                                                                      *float64   `json:"HookloadDown,omitempty"`
	// Hookload when the string is moving up.                                                    
	HookloadUp                                                                        *float64   `json:"HookloadUp,omitempty"`
	// Inclination at start measured depth.                                                      
	InclinationStart                                                                  *float64   `json:"InclinationStart,omitempty"`
	// Inclination at stop measured depth.                                                       
	InclinationStop                                                                   *float64   `json:"InclinationStop,omitempty"`
	// The maximum flowrate at the bit during the run                                            
	MaximumBitFlowrate                                                                *float64   `json:"MaximumBitFlowrate,omitempty"`
	// Maximum mud pump flow rate.                                                               
	MaximumFlowratePump                                                               *float64   `json:"MaximumFlowratePump,omitempty"`
	// Maximum inclination.                                                                      
	MaximumInclination                                                                *float64   `json:"MaximumInclination,omitempty"`
	// Maximum mud temperature downhole during run.                                              
	MaximumMudTemperatureDownhole                                                     *float64   `json:"MaximumMudTemperatureDownhole,omitempty"`
	// The maximum pump pressure used during the run                                             
	MaximumPumpPressure                                                               *float64   `json:"MaximumPumpPressure,omitempty"`
	// Maximum rate of penetration through Interval.                                             
	MaximumROP                                                                        *float64   `json:"MaximumROP,omitempty"`
	// Maximum turn rate (commonly in rpm).                                                      
	MaximumRPM                                                                        *float64   `json:"MaximumRpm,omitempty"`
	// Maximum torque: on bottom.                                                                
	MaximumTorqueOnBottom                                                             *float64   `json:"MaximumTorqueOnBottom,omitempty"`
	// Weight on bit - maximum.                                                                  
	MaximumWOB                                                                        *float64   `json:"MaximumWOB,omitempty"`
	// The minimum flowrate at the bit during the run                                            
	MinimumBitFlowrate                                                                *float64   `json:"MinimumBitFlowrate,omitempty"`
	// Minimum mud pump flow rate.                                                               
	MinimumFlowratePump                                                               *float64   `json:"MinimumFlowratePump,omitempty"`
	// Minimum inclination.                                                                      
	MinimumInclination                                                                *float64   `json:"MinimumInclination,omitempty"`
	// The minimum pump pressure used during the run                                             
	MinimumPumpPressure                                                               *float64   `json:"MinimumPumpPressure,omitempty"`
	// Minimum rate of penetration through Interval.                                             
	MinimumROP                                                                        *float64   `json:"MinimumROP,omitempty"`
	// Minimum turn rate (commonly in rpm).                                                      
	MinimumRPM                                                                        *float64   `json:"MinimumRpm,omitempty"`
	// Minimum torque: on bottom.                                                                
	MinimumTorqueOnBottom                                                             *float64   `json:"MinimumTorqueOnBottom,omitempty"`
	// Weight on bit - minimum.                                                                  
	MinimumWOB                                                                        *float64   `json:"MinimumWOB,omitempty"`
	// The class of the drilling fluid.                                                          
	MudClass                                                                          *string    `json:"MudClass,omitempty"`
	// Overpull = HookloadUp - HookloadRotating                                                  
	OverPull                                                                          *float64   `json:"OverPull,omitempty"`
	// Bit hydraulic.                                                                            
	PowerBit                                                                          *float64   `json:"PowerBit,omitempty"`
	// Pressure drop in bit.                                                                     
	PressureDropBit                                                                   *float64   `json:"PressureDropBit,omitempty"`
	// Hookload rotating.                                                                        
	RotatingHookload                                                                  *float64   `json:"RotatingHookload,omitempty"`
	// Measured depth at the end of the run.                                                     
	RunEndHoleMeasuredDepth                                                           *float64   `json:"RunEndHoleMeasuredDepth,omitempty"`
	// Measured depth at start of the run.                                                       
	RunStartHoleMeasuredDepth                                                         *float64   `json:"RunStartHoleMeasuredDepth,omitempty"`
	// Slackoff = HookloadRotating  - HookloadDown.                                              
	SlackOff                                                                          *float64   `json:"SlackOff,omitempty"`
	// Time spent circulating from start of bit run.                                             
	TimeCirculating                                                                   *float64   `json:"TimeCirculating,omitempty"`
	// Time spent rotary drilling from start of bit run.                                         
	TimeDrillingRotating                                                              *float64   `json:"TimeDrillingRotating,omitempty"`
	// Time spent slide drilling from start of bit run.                                          
	TimeDrillingSliding                                                               *float64   `json:"TimeDrillingSliding,omitempty"`
	// Time spent on hold from start of bit run.                                                 
	TimeHolding                                                                       *float64   `json:"TimeHolding,omitempty"`
	// Time spent reaming from start of bit run.                                                 
	TimeReaming                                                                       *float64   `json:"TimeReaming,omitempty"`
	// Time spent steering from start of bit run.                                                
	TimeSteering                                                                      *float64   `json:"TimeSteering,omitempty"`
	// Weight of the string above the jars.                                                      
	WeightAboveJar                                                                    *float64   `json:"WeightAboveJar,omitempty"`
	// Weight of the string below the jars.                                                      
	WeightBelowJar                                                                    *float64   `json:"WeightBelowJar,omitempty"`
	// Drilling fluid density.                                                                   
	WeightMud                                                                         *float64   `json:"WeightMud,omitempty"`
}
