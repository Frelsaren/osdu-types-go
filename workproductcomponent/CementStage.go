package workproductcomponent

import "time"

// A single cement stage performed within a Job.
type CementStage struct {
	// Length of the cement column                                                                                        
	AnnularLength                                                                              *float64                   `json:"AnnularLength,omitempty"`
	// Constant back pressure applied at surface while pumping the job (can be superseded by a                            
	// back pressure per pumping stage).                                                                                  
	BackPressure                                                                               *float64                   `json:"BackPressure,omitempty"`
	// Behind Cement Spacer Volume                                                                                        
	BehindCementSpacerVolume                                                                   *float64                   `json:"BehindCementSpacerVolume,omitempty"`
	// Circulating Bottomhole temperature                                                                                 
	Bhct                                                                                       *float64                   `json:"BHCT,omitempty"`
	// Static Bottomhole temperature                                                                                      
	Bhst                                                                                       *float64                   `json:"BHST,omitempty"`
	// Volume pumped to landing Bottom Dart                                                                               
	BottomDartLandVolume                                                                       *float64                   `json:"BottomDartLandVolume,omitempty"`
	// Bottom Dart Shear Pressure                                                                                         
	BottomDartShearPressure                                                                    *float64                   `json:"BottomDartShearPressure,omitempty"`
	// Number of Bottom Plugs                                                                                             
	BottomPlugNumber                                                                           *int64                     `json:"BottomPlugNumber,omitempty"`
	// Breakdown Flow Rate                                                                                                
	BreakdownFlowRate                                                                          *float64                   `json:"BreakdownFlowRate,omitempty"`
	// Breakdown Pressure                                                                                                 
	BreakdownPressure                                                                          *float64                   `json:"BreakdownPressure,omitempty"`
	// Bump Plug Pressure                                                                                                 
	BumpPlugPressure                                                                           *float64                   `json:"BumpPlugPressure,omitempty"`
	// Array of fluids worked in the stage - displaced mud, washers and spacers, cement (lead &                           
	// tail)                                                                                                              
	CementingFluid                                                                             []CementingFluid           `json:"CementingFluid,omitempty"`
	// Circulate out measured depth from top of cement plug                                                               
	CirculateOutMeasuredDepth                                                                  *float64                   `json:"CirculateOutMeasuredDepth,omitempty"`
	// Bottom Measured depth of Coiled Tubing                                                                             
	CoilTubingMeasuredDepth                                                                    *float64                   `json:"CoilTubingMeasuredDepth,omitempty"`
	// Average (or actual) displacement rate                                                                              
	DisplacementFlowrateAvg                                                                    *float64                   `json:"DisplacementFlowrateAvg,omitempty"`
	// Final displacement rate                                                                                            
	DisplacementFlowrateFinal                                                                  *float64                   `json:"DisplacementFlowrateFinal,omitempty"`
	// Maximum displacement rate                                                                                          
	DisplacementFlowrateMax                                                                    *float64                   `json:"DisplacementFlowrateMax,omitempty"`
	// Density of displacement fluid.                                                                                     
	DisplacementFluidDensity                                                                   *float64                   `json:"DisplacementFluidDensity,omitempty"`
	// Displacement fluid name.                                                                                           
	DisplacementFluidName                                                                      *string                    `json:"DisplacementFluidName,omitempty"`
	// Volume of displacement fluid.                                                                                      
	DisplacementFluidVolume                                                                    *float64                   `json:"DisplacementFluidVolume,omitempty"`
	// Final displacement pressure                                                                                        
	DisplacementPressureFinal                                                                  *float64                   `json:"DisplacementPressureFinal,omitempty"`
	// Date/time when displacing of cement started.                                                                       
	DisplaceStartDate                                                                          *time.Time                 `json:"DisplaceStartDate,omitempty"`
	// Excess Volume used                                                                                                 
	ExcessVolume                                                                               *float64                   `json:"ExcessVolume,omitempty"`
	// Excess Volume Calculation Method                                                                                   
	ExcessVolumeCalcMethodID                                                                   *string                    `json:"ExcessVolumeCalcMethodID,omitempty"`
	// Final Casing Pressure                                                                                              
	FinalCasingPressure                                                                        *float64                   `json:"FinalCasingPressure,omitempty"`
	// Final Coiled Tubing Pressure                                                                                       
	FinalCoilTubingPressure                                                                    *float64                   `json:"FinalCoilTubingPressure,omitempty"`
	// Final Pump Rate                                                                                                    
	FinalPumpRate                                                                              *float64                   `json:"FinalPumpRate,omitempty"`
	// Tubing End Pressure (not CT)                                                                                       
	FinalTubingPressure                                                                        *float64                   `json:"FinalTubingPressure,omitempty"`
	// Pre-circulation Fluid Temperature In (Pit)                                                                         
	FluidTemperatureIn                                                                         *float64                   `json:"FluidTemperatureIn,omitempty"`
	// Pre-circulation Fluid Temperature Out (Return Line)                                                                
	FluidTemperatureOut                                                                        *float64                   `json:"FluidTemperatureOut,omitempty"`
	// Pressure held to.                                                                                                  
	HeldPressure                                                                               *float64                   `json:"HeldPressure,omitempty"`
	// Hesitation Reason during operation                                                                                 
	HesitationReason                                                                           *string                    `json:"HesitationReason,omitempty"`
	// Hesitation Squeeze array                                                                                           
	HesitationSqueeze                                                                          []HesitationSqueezeHistory `json:"HesitationSqueeze,omitempty"`
	// Initial Casing Pressure                                                                                            
	InitialCasingPressure                                                                      *float64                   `json:"InitialCasingPressure,omitempty"`
	// Initial Coiled Tubing Pressure                                                                                     
	InitialCoilTubingPressure                                                                  *float64                   `json:"InitialCoilTubingPressure,omitempty"`
	// Initial Pump Rate                                                                                                  
	InitialPumpRate                                                                            *float64                   `json:"InitialPumpRate,omitempty"`
	// Tubing Start Pressure (not CT)                                                                                     
	InitialTubingPressure                                                                      *float64                   `json:"InitialTubingPressure,omitempty"`
	// Estimated volume inside casing                                                                                     
	InsideCasingVolume                                                                         *float64                   `json:"InsideCasingVolume,omitempty"`
	// Final placement measured depth of base of interval.                                                                
	IntervalBaseMeasuredDepth                                                                  *float64                   `json:"IntervalBaseMeasuredDepth,omitempty"`
	// Final placement True Vertical depth of base of interval.                                                           
	IntervalBaseTrueVerticalDepth                                                              *float64                   `json:"IntervalBaseTrueVerticalDepth,omitempty"`
	// Final placement measured depth at top of interval.                                                                 
	IntervalTopMeasuredDepth                                                                   *float64                   `json:"IntervalTopMeasuredDepth,omitempty"`
	// Final placement True Vertical depth at top of interval.                                                            
	IntervalTopTrueVerticalDepth                                                               *float64                   `json:"IntervalTopTrueVerticalDepth,omitempty"`
	// Annular flow (back flow) present after the stage was completed?                                                    
	IsAnnularFlowAfter                                                                         *bool                      `json:"IsAnnularFlowAfter,omitempty"`
	// Annular Pressure Held after job?                                                                                   
	IsAnnularPressureHeld                                                                      *bool                      `json:"IsAnnularPressureHeld,omitempty"`
	// Bottom plug used?                                                                                                  
	IsBottomPlugUsed                                                                           *bool                      `json:"IsBottomPlugUsed,omitempty"`
	// Float Held?                                                                                                        
	IsFloatHeld                                                                                *bool                      `json:"IsFloatHeld,omitempty"`
	// Hesitation when squeezing?                                                                                         
	IsHesitation                                                                               *bool                      `json:"IsHesitation,omitempty"`
	// Pill below cement plug?                                                                                            
	IsPillBelowPlug                                                                            *bool                      `json:"IsPillBelowPlug,omitempty"`
	// Is Plug Bumped?                                                                                                    
	IsPlugBumped                                                                               *bool                      `json:"IsPlugBumped,omitempty"`
	// Plug or Dart catcher?                                                                                              
	IsPlugCatcher                                                                              *bool                      `json:"IsPlugCatcher,omitempty"`
	// Is Squeeze Obtained/Successful?                                                                                    
	IsSqueezeObtained                                                                          *bool                      `json:"IsSqueezeObtained,omitempty"`
	// Tailpipe/stringer perforated?                                                                                      
	IsTailPipePerforated                                                                       *bool                      `json:"IsTailPipePerforated,omitempty"`
	// Tail pipe/stinger used for setting a plug?                                                                         
	IsTailPipeUsed                                                                             *bool                      `json:"IsTailPipeUsed,omitempty"`
	// Top plug used?                                                                                                     
	IsTopPlugUsed                                                                              *bool                      `json:"IsTopPlugUsed,omitempty"`
	// Volume of mud lost                                                                                                 
	LostVolume                                                                                 *float64                   `json:"LostVolume,omitempty"`
	// Max Static Time After Displacement                                                                                 
	MaxStaticTimeAfterDisplacement                                                             *float64                   `json:"MaxStaticTimeAfterDisplacement,omitempty"`
	// Max Static Time During Placement                                                                                   
	MaxStaticTimeDuringPlacement                                                               *float64                   `json:"MaxStaticTimeDuringPlacement,omitempty"`
	// Max Static Time Surface - cement unit and surface lines                                                            
	MaxStaticTimeSurface                                                                       *float64                   `json:"MaxStaticTimeSurface,omitempty"`
	// Min WOC Time Before Logging                                                                                        
	MinWOCTimeBeforeLogging                                                                    *float64                   `json:"MinWOCTimeBeforeLogging,omitempty"`
	// Min WOC Time Before Tagging/Testing                                                                                
	MinWOCTimeBeforeTagging                                                                    *float64                   `json:"MinWOCTimeBeforeTagging,omitempty"`
	// Mix Method                                                                                                         
	MixMethod                                                                                  *string                    `json:"MixMethod,omitempty"`
	// Date/time when mixing of cement started.                                                                           
	MixStartDate                                                                               *time.Time                 `json:"MixStartDate,omitempty"`
	// Mud Circulation Elapsed Time                                                                                       
	MudCirculationElapsedTime                                                                  *float64                   `json:"MudCirculationElapsedTime,omitempty"`
	// Mud Circulation Flowrate during the job                                                                            
	MudCirculationFlowRate                                                                     *float64                   `json:"MudCirculationFlowRate,omitempty"`
	// Mud circulation pressure                                                                                           
	MudCirculationPressure                                                                     *float64                   `json:"MudCirculationPressure,omitempty"`
	// Mud density (weight)                                                                                               
	MudDensity                                                                                 *float64                   `json:"MudDensity,omitempty"`
	// Marsh funnel viscosity measured in a conical-shaped funnel                                                         
	MudFunnelViscosity                                                                         *float64                   `json:"MudFunnelViscosity,omitempty"`
	// Gels-10Min (in hole at start of job).                                                                              
	MudGel10Min                                                                                *float64                   `json:"MudGel10Min,omitempty"`
	// Gels-10Sec (in hole at start of job).                                                                              
	MudGel10SEC                                                                                *float64                   `json:"MudGel10Sec,omitempty"`
	// Plastic viscosity (in hole at start of job).                                                                       
	MudPlasticViscosity                                                                        *float64                   `json:"MudPlasticViscosity,omitempty"`
	// Yield point (in hole at start of job).                                                                             
	MudYieldPoint                                                                              *float64                   `json:"MudYieldPoint,omitempty"`
	// Type of mud in hole                                                                                                
	OriginalMudTypeID                                                                          *string                    `json:"OriginalMudTypeID,omitempty"`
	// Estimated volume outside casing for this stage placement                                                           
	OutsideCasingVolume                                                                        *float64                   `json:"OutsideCasingVolume,omitempty"`
	// Plug Manufacturer                                                                                                  
	PlugManufacturer                                                                           *string                    `json:"PlugManufacturer,omitempty"`
	// Pressure Held time at end of stage                                                                                 
	PressureHeldTime                                                                           *float64                   `json:"PressureHeldTime,omitempty"`
	// Volume of mud circulated prior to cement job                                                                       
	PriorCirculatingVolume                                                                     *float64                   `json:"PriorCirculatingVolume,omitempty"`
	// Date/time when pumping cement ended.                                                                               
	PumpEndDate                                                                                *time.Time                 `json:"PumpEndDate,omitempty"`
	// Average Pump Pressure                                                                                              
	PumpPressureAverage                                                                        *float64                   `json:"PumpPressureAverage,omitempty"`
	// Pump Rate Maximum                                                                                                  
	PumpRateMax                                                                                *float64                   `json:"PumpRateMax,omitempty"`
	// Date/time when pumping cement started.                                                                             
	PumpStartDate                                                                              *time.Time                 `json:"PumpStartDate,omitempty"`
	// Maximum overpull applied during reciprocating                                                                      
	ReciprocationOverpull                                                                      *float64                   `json:"ReciprocationOverpull,omitempty"`
	// Maximum slackoff applied during reciprocating                                                                      
	ReciprocationSlackoff                                                                      *float64                   `json:"ReciprocationSlackoff,omitempty"`
	// Volume of mud returned                                                                                             
	ReturnsVolume                                                                              *float64                   `json:"ReturnsVolume,omitempty"`
	// Shoe Track Length                                                                                                  
	ShoeTrackLength                                                                            *float64                   `json:"ShoeTrackLength,omitempty"`
	// Shoe Track Volume                                                                                                  
	ShoeTrackVolume                                                                            *float64                   `json:"ShoeTrackVolume,omitempty"`
	// Squeeze average flowrate                                                                                           
	SqueezeFlowRateAverage                                                                     *float64                   `json:"SqueezeFlowRateAverage,omitempty"`
	// Squeeze maximum flowrate                                                                                           
	SqueezeFlowRateMax                                                                         *float64                   `json:"SqueezeFlowRateMax,omitempty"`
	// Squeeze pressure left (held) on pipe while waiting on cement.                                                      
	SqueezeHeldPressure                                                                        *float64                   `json:"SqueezeHeldPressure,omitempty"`
	// Squeeze Held Pressure Elapsed Time                                                                                 
	SqueezeHeldPressureDuration                                                                *float64                   `json:"SqueezeHeldPressureDuration,omitempty"`
	// Squeezed cement volume left inside casing                                                                          
	SqueezeInCasingVolume                                                                      *float64                   `json:"SqueezeInCasingVolume,omitempty"`
	// Squeeze objective                                                                                                  
	SqueezeObjective                                                                           *string                    `json:"SqueezeObjective,omitempty"`
	// Squeezed cement volume outside/behind                                                                              
	SqueezeOutCasingVolume                                                                     *float64                   `json:"SqueezeOutCasingVolume,omitempty"`
	// Squeeze Pressure Average                                                                                           
	SqueezePressureAverage                                                                     *float64                   `json:"SqueezePressureAverage,omitempty"`
	// Squeeze Pressure Final                                                                                             
	SqueezePressureFinal                                                                       *float64                   `json:"SqueezePressureFinal,omitempty"`
	// Squeeze Pressure Initial                                                                                           
	SqueezePressureInitial                                                                     *float64                   `json:"SqueezePressureInitial,omitempty"`
	// Squeeze Remarks                                                                                                    
	SqueezeRemarks                                                                             *string                    `json:"SqueezeRemarks,omitempty"`
	// Volume of cement that was returned to surface during squeeze operation                                             
	SqueezeReverseVolume                                                                       *float64                   `json:"SqueezeReverseVolume,omitempty"`
	// Squeeze Tool Measured Depth e.g. Packer or Retainer                                                                
	SqueezeToolMeasuredDepth                                                                   *float64                   `json:"SqueezeToolMeasuredDepth,omitempty"`
	// Squeeze Tool True Vertical Depth e.g. Packer or Retainer                                                           
	SqueezeToolTrueVerticalDepth                                                               *float64                   `json:"SqueezeToolTrueVerticalDepth,omitempty"`
	// Measured depth interval for the cement stage (Base MD - Top MD).                                                   
	StageIntervalLength                                                                        *float64                   `json:"StageIntervalLength,omitempty"`
	// Stage number.                                                                                                      
	StageNumber                                                                                float64                    `json:"StageNumber"`
	// Remarks                                                                                                            
	StageRemarks                                                                               *string                    `json:"StageRemarks,omitempty"`
	// Mid measured depth of stage tool (multi-stage cement job) e.g. circulation port.                                   
	StageToolMeasuredDepth                                                                     *float64                   `json:"StageToolMeasuredDepth,omitempty"`
	// Mid true vertical depth of stage tool (multi-stage cement job) e.g. circulation port.                              
	StageToolTrueVerticalDepth                                                                 *float64                   `json:"StageToolTrueVerticalDepth,omitempty"`
	// Stage type.                                                                                                        
	StageTypeID                                                                                string                     `json:"StageTypeID"`
	// Bottom measured depth in primary cement job (multi-stage cement job).                                              
	StringMeasuredDepth                                                                        *float64                   `json:"StringMeasuredDepth,omitempty"`
	// Tailpipe/stinger size (diameter).                                                                                  
	TailPipeDiameter                                                                           *float64                   `json:"TailPipeDiameter,omitempty"`
	// Volume pumped to landing Top Dart                                                                                  
	TopDartLandVolume                                                                          *float64                   `json:"TopDartLandVolume,omitempty"`
	// Top Dart Shear Pressure                                                                                            
	TopDartShearPressure                                                                       *float64                   `json:"TopDartShearPressure,omitempty"`
	// Total operating hours from at depth ready for cement job. Included but not limited to:                             
	// Pre-circulation, NPT, rig up/down, line-up and pumping operations.                                                 
	TotalOperatingTime                                                                         *float64                   `json:"TotalOperatingTime,omitempty"`
	// Total pumping hours from at depth ready for cement job. Included but not limited to:                               
	// Pre-circulation, washing, spacer, cement and displacement.                                                         
	TotalPumpingTime                                                                           *float64                   `json:"TotalPumpingTime,omitempty"`
	// Under Displacement Volume Actual                                                                                   
	UnderDisplacementVolumeActual                                                              *float64                   `json:"UnderDisplacementVolumeActual,omitempty"`
	// Under Displacement Volume Planned                                                                                  
	UnderDisplacementVolumePlanned                                                             *float64                   `json:"UnderDisplacementVolumePlanned,omitempty"`
	// Workstring Bottom Measured Depth during displacement for Plugs & Squeezes                                          
	WorkstringMeasuredDepth                                                                    *float64                   `json:"WorkstringMeasuredDepth,omitempty"`
	// Workstring Bottom True Vertical Depth during displacement for Plugs & Squeezes                                     
	WorkstringTrueVerticalDepth                                                                *float64                   `json:"WorkstringTrueVerticalDepth,omitempty"`
}
