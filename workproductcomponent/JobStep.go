package workproductcomponent

import "time"

// Summary of the pumping schedule, also known as 'pumping stages', sub-stages or event
// level details.
// This is where you define pressures, rates, volumes, amount of proppant for the
// Displacement, Flush, Pad, Ramp, etc steps in the pumping operation for this job stage.
type JobStep struct {
	// Balls recovered during execution of the step.                                                           
	BallsRecovered                                                              *int64                         `json:"BallsRecovered,omitempty"`
	// Balls used during execution of the step.                                                                
	BallsUsed                                                                   *int64                         `json:"BallsUsed,omitempty"`
	// Base fluid volume recorded after equipment set to bypass.                                               
	BaseFluidBypassVolume                                                       *float64                       `json:"BaseFluidBypassVolume,omitempty"`
	// Base quality percentage of foam.                                                                        
	BaseFluidQualityAvg                                                         *float64                       `json:"BaseFluidQualityAvg,omitempty"`
	// Base fluid volume entering the equipment.                                                               
	BaseFluidVolumeEquip                                                        *float64                       `json:"BaseFluidVolumeEquip,omitempty"`
	// Bottomhole material usage rate average.                                                                 
	BHMaterialUsedRateAvg                                                       []AbstractStimMaterialQuantity `json:"BHMaterialUsedRateAvg,omitempty"`
	// Bottomhole material usage rate end.                                                                     
	BHMaterialUsedRateEnd                                                       []AbstractStimMaterialQuantity `json:"BHMaterialUsedRateEnd,omitempty"`
	// Bottomhole material usage rate start.                                                                   
	BHMaterialUsedRateStart                                                     []AbstractStimMaterialQuantity `json:"BHMaterialUsedRateStart,omitempty"`
	// Average bottomhole pressure.                                                                            
	BHPressureAvg                                                               *float64                       `json:"BHPressureAvg,omitempty"`
	// Final bottomhole pressure.                                                                              
	BHPressureEnd                                                               *float64                       `json:"BHPressureEnd,omitempty"`
	// Starting bottomhole pressure.                                                                           
	BHPressureStart                                                             *float64                       `json:"BHPressureStart,omitempty"`
	// Base quality carbon dioxide percent of foam.                                                            
	CO2BaseFluidQualityAvg                                                      *float64                       `json:"CO2BaseFluidQualityAvg,omitempty"`
	// Final CO2 pump rate in volume per time at the surface.                                                  
	CO2SurfaceRateEnd                                                           *float64                       `json:"CO2SurfaceRateEnd,omitempty"`
	// A short description of the job step.                                                                    
	Description                                                                 *string                        `json:"Description,omitempty"`
	// Ending dirty fluid flow rate.                                                                           
	DirtyMaterialRateEnd                                                        *float64                       `json:"DirtyMaterialRateEnd,omitempty"`
	// Starting dirty fluid flow rate.                                                                         
	DirtyMaterialRateStart                                                      *float64                       `json:"DirtyMaterialRateStart,omitempty"`
	// Fluid used for this job step of the stimulated interval.                                                
	Fluid                                                                       *JobStepFluid                  `json:"Fluid,omitempty"`
	// The step volume of the base fluid.                                                                      
	FluidBaseVolume                                                             *float64                       `json:"FluidBaseVolume,omitempty"`
	// Fluid volume circulated.                                                                                
	FluidCirculatedVolume                                                       *float64                       `json:"FluidCirculatedVolume,omitempty"`
	// Fluid volume pumped.                                                                                    
	FluidPumpedVolume                                                           *float64                       `json:"FluidPumpedVolume,omitempty"`
	// Fluid Report Identifier of the fluid used for this Job Step                                             
	FluidReportID                                                               *string                        `json:"FluidReportID,omitempty"`
	// Fluid volume returned.                                                                                  
	FluidReturnedVolume                                                         *float64                       `json:"FluidReturnedVolume,omitempty"`
	// Fluid volume squeezed.                                                                                  
	FluidSqueezedVolume                                                         *float64                       `json:"FluidSqueezedVolume,omitempty"`
	// Average fluid temperature.                                                                              
	FluidTemperatureAvg                                                         *float64                       `json:"FluidTemperatureAvg,omitempty"`
	// Fluid volume washed.                                                                                    
	FluidWashedVolume                                                           *float64                       `json:"FluidWashedVolume,omitempty"`
	// The fracture gradient when the step ends.                                                               
	FractureGradientFinal                                                       *float64                       `json:"FractureGradientFinal,omitempty"`
	// The fracture gradient before starting the step.                                                         
	FractureGradientInitial                                                     *float64                       `json:"FractureGradientInitial,omitempty"`
	// Numeric value used to scale a calculated rheological friction.                                          
	FrictionFactor                                                              *float64                       `json:"FrictionFactor,omitempty"`
	// Average hydraulic horsepower used during the step.                                                      
	HhpAvg                                                                      *float64                       `json:"HhpAvg,omitempty"`
	// Maximum hydraulic horsepower used during the step.                                                      
	HhpMax                                                                      *float64                       `json:"HhpMax,omitempty"`
	// Internal gas phase percentage of the foam.                                                              
	InternalPhaseFractionAvg                                                    *float64                       `json:"InternalPhaseFractionAvg,omitempty"`
	// The type of  job step.                                                                                  
	JobStepTypeID                                                               *string                        `json:"JobStepTypeID,omitempty"`
	// Material used during the step                                                                           
	MaterialUsed                                                                []AbstractStimMaterialQuantity `json:"MaterialUsed,omitempty"`
	// Average material used per minute entering the flow stream.                                              
	MaterialUsedRateAvg                                                         []AbstractStimMaterialQuantity `json:"MaterialUsedRateAvg,omitempty"`
	// Ending quantity of material used per minute entering the flow stream.                                   
	MaterialUsedRateEnd                                                         []AbstractStimMaterialQuantity `json:"MaterialUsedRateEnd,omitempty"`
	// Maximum rate of material used per minute entering the flow stream.                                      
	MaterialUsedRateMax                                                         []AbstractStimMaterialQuantity `json:"MaterialUsedRateMax,omitempty"`
	// Starting quantity of material used per minute entering the flow stream.                                 
	MaterialUsedRateStart                                                       []AbstractStimMaterialQuantity `json:"MaterialUsedRateStart,omitempty"`
	// Nitrogen base quality percentage of foam.                                                               
	N2BaseFluidQualityAvg                                                       *float64                       `json:"N2BaseFluidQualityAvg,omitempty"`
	// Final nitrogen pump rate at the surface.                                                                
	N2SurfaceRateEnd                                                            *float64                       `json:"N2SurfaceRateEnd,omitempty"`
	// The friction pressure loss contribution from pipe.                                                      
	PipeFrictionPressureLoss                                                    *float64                       `json:"PipeFrictionPressureLoss,omitempty"`
	// Average proppant concentration at the wellhead.                                                         
	ProppantConcAvg                                                             *float64                       `json:"ProppantConcAvg,omitempty"`
	// The average proppant concentration at bottomhole.                                                       
	ProppantConcBHAvg                                                           *float64                       `json:"ProppantConcBHAvg,omitempty"`
	// The final proppant concentration at bottomhole.                                                         
	ProppantConcBHEnd                                                           *float64                       `json:"ProppantConcBHEnd,omitempty"`
	// Maximum proppant concentration at bottomhole during the stimulation step.                               
	ProppantConcBHMax                                                           *float64                       `json:"ProppantConcBHMax,omitempty"`
	// The initial proppant concentration at bottomhole.                                                       
	ProppantConcBHStart                                                         *float64                       `json:"ProppantConcBHStart,omitempty"`
	// The average proppant concentration at the surface.                                                      
	ProppantConcSurfaceAvg                                                      *float64                       `json:"ProppantConcSurfaceAvg,omitempty"`
	// The final proppant concentration at the surface.                                                        
	ProppantConcSurfaceEnd                                                      *float64                       `json:"ProppantConcSurfaceEnd,omitempty"`
	// Maximum proppant concentration at the wellhead.                                                         
	ProppantConcSurfaceMax                                                      *float64                       `json:"ProppantConcSurfaceMax,omitempty"`
	// The initial proppant concentration at the surface.                                                      
	ProppantConcSurfaceStart                                                    *float64                       `json:"ProppantConcSurfaceStart,omitempty"`
	// Average proppant concentration exiting the equipment.                                                   
	ProppantSlurryConcAvg                                                       *float64                       `json:"ProppantSlurryConcAvg,omitempty"`
	// Maximum proppant concentration exiting the equipment.                                                   
	ProppantSlurryConcMax                                                       *float64                       `json:"ProppantSlurryConcMax,omitempty"`
	// Total pumping time for the step.                                                                        
	PumpTime                                                                    *float64                       `json:"PumpTime,omitempty"`
	// General remarks about this job step.                                                                    
	Remarks                                                                     *string                        `json:"Remarks,omitempty"`
	// The volume of the slurry (dirty) for this job step.                                                     
	SlurryFluidVolume                                                           *float64                       `json:"SlurryFluidVolume,omitempty"`
	// Average slurry return rate.                                                                             
	SlurryReturnRateAvg                                                         *float64                       `json:"SlurryReturnRateAvg,omitempty"`
	// Date and time the step ended.                                                                           
	StepEndDateTime                                                             *time.Time                     `json:"StepEndDateTime,omitempty"`
	// A human readable name for this job step.                                                                
	StepName                                                                    *string                        `json:"StepName,omitempty"`
	// Step number.                                                                                            
	StepNumber                                                                  *int64                         `json:"StepNumber,omitempty"`
	// Date and time the step started.                                                                         
	StepStartDateTime                                                           *time.Time                     `json:"StepStartDateTime,omitempty"`
	// Average surface pressure.                                                                               
	SurfacePressureAvg                                                          *float64                       `json:"SurfacePressureAvg,omitempty"`
	// Final surface pressure.                                                                                 
	SurfacePressureEnd                                                          *float64                       `json:"SurfacePressureEnd,omitempty"`
	// Maximum pumping pressure on surface.                                                                    
	SurfacePressureMax                                                          *float64                       `json:"SurfacePressureMax,omitempty"`
	// Starting surface pressure.                                                                              
	SurfacePressureStart                                                        *float64                       `json:"SurfacePressureStart,omitempty"`
	// Average flow rate at the wellhead.                                                                      
	WellheadFlowRateAvg                                                         *float64                       `json:"WellheadFlowRateAvg,omitempty"`
	// Maximum flow rate at the wellhead.                                                                      
	WellheadFlowRateMax                                                         *float64                       `json:"WellheadFlowRateMax,omitempty"`
	// Slurry volume entering the well.                                                                        
	WellheadSlurryVolume                                                        *float64                       `json:"WellheadSlurryVolume,omitempty"`
}
