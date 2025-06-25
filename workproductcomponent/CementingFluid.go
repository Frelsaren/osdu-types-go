package workproductcomponent

// Original mud, washes, spacers, cement and displacement fluid properties.
type CementingFluid struct {
	// Volume of given fluid that resides in the annulus.                                                             
	AnnularFluidVolume                                                                          *float64              `json:"AnnularFluidVolume,omitempty"`
	// API Fluid Loss @ 30 Min                                                                                        
	APIFluidLoss30Min                                                                           *float64              `json:"APIFluidLoss30Min,omitempty"`
	// API fluid loss = 2 * volTestFluidLoss * SQRT(30/timeFluidLoss).                                                
	APIFluidLossRate                                                                            *float64              `json:"APIFluidLossRate,omitempty"`
	// Density of base fluid.                                                                                         
	BaseFluidDensity                                                                            *float64              `json:"BaseFluidDensity,omitempty"`
	// Drilling fluid base type                                                                                       
	BaseFluidTypeID                                                                             *string               `json:"BaseFluidTypeID,omitempty"`
	// Estimate Measured depth at bottom of fluid placement at end of Stage                                           
	BaseMeasuredDepth                                                                           *float64              `json:"BaseMeasuredDepth,omitempty"`
	// Cased Hole Excess Slurry Volume Percent                                                                        
	CasedHoleExcessPercent                                                                      *float64              `json:"CasedHoleExcessPercent,omitempty"`
	// Cement Additives list. Additives can be added in slurry but also in spacers, washes, mud.                      
	CementAdditives                                                                             []CementAdditive      `json:"CementAdditives,omitempty"`
	// Comments or Remarks                                                                                            
	CementFluidRemark                                                                           *string               `json:"CementFluidRemark,omitempty"`
	// Cement Mass Unit Type (Sacks or Tonnes)                                                                        
	CementMassUnitTypeID                                                                        *string               `json:"CementMassUnitTypeID,omitempty"`
	// Set of (Time / Rate / Back Pressure).                                                                          
	CementPumpSchedules                                                                         []CementPumpSchedule  `json:"CementPumpSchedules,omitempty"`
	// Cement Test Lab ID Number                                                                                      
	CementTestLabID                                                                             *string               `json:"CementTestLabID,omitempty"`
	// Cement Yield per Unit                                                                                          
	CementYield                                                                                 *float64              `json:"CementYield,omitempty"`
	// Compressive Strength Pressure/Temperature/Thickening                                                           
	CompressiveStrengthTest                                                                     []CompressiveStrength `json:"CompressiveStrengthTest,omitempty"`
	// Constant Gas Foam Method Density                                                                               
	ConstantGasFoamDensity                                                                      *float64              `json:"ConstantGasFoamDensity,omitempty"`
	// Fluid density                                                                                                  
	Density                                                                                     *float64              `json:"Density,omitempty"`
	// Constant gas ratio method for measuring average density.                                                       
	DensityConstantGasMethodID                                                                  *string               `json:"DensityConstantGasMethodID,omitempty"`
	// Density Method                                                                                                 
	DensityMethodID                                                                             *string               `json:"DensityMethodID,omitempty"`
	// Fluid Density Measurement Pressure                                                                             
	DensityPressure                                                                             *float64              `json:"DensityPressure,omitempty"`
	// Fluid Density Measurement Temperature                                                                          
	DensityTemperature                                                                          *float64              `json:"DensityTemperature,omitempty"`
	// Design fluid density                                                                                           
	DesignDensity                                                                               *float64              `json:"DesignDensity,omitempty"`
	// DesignVolume                                                                                                   
	DesignVolume                                                                                *float64              `json:"DesignVolume,omitempty"`
	// Dry blend cement slurry class                                                                                  
	DryBlendCementSlurryClassID                                                                 *string               `json:"DryBlendCementSlurryClassID,omitempty"`
	// Density of Dry blend.                                                                                          
	DryBlendDensity                                                                             *float64              `json:"DryBlendDensity,omitempty"`
	// Description of dry blend.                                                                                      
	DryBlendDescription                                                                         *string               `json:"DryBlendDescription,omitempty"`
	// Total mass of dry blend including additives: the volume is not constant.                                       
	DryBlendMass                                                                                *float64              `json:"DryBlendMass,omitempty"`
	// Weight of a sack of dry blend aka Sack Size or Tonnes                                                          
	DryBlendMassPerUnit                                                                         *float64              `json:"DryBlendMassPerUnit,omitempty"`
	// Name of dry blend.                                                                                             
	DryBlendName                                                                                *string               `json:"DryBlendName,omitempty"`
	// Excess Slurry Volume                                                                                           
	ExcessSlurryVolume                                                                          *float64              `json:"ExcessSlurryVolume,omitempty"`
	// Excess Slurry Volume Measurement Type                                                                          
	ExcessSlurryVolumeMeasureTypeID                                                             *string               `json:"ExcessSlurryVolumeMeasureTypeID,omitempty"`
	// True Vertical Depth at bottom of fluid placement at the end of Stage.                                          
	FluidBaseTrueVerticalDepth                                                                  *float64              `json:"FluidBaseTrueVerticalDepth,omitempty"`
	// Commercial product or trade name of fluid                                                                      
	FluidCommercialNameID                                                                       *string               `json:"FluidCommercialNameID,omitempty"`
	// Fluid description.                                                                                             
	FluidDescription                                                                            *string               `json:"FluidDescription,omitempty"`
	// Cement pumping schedule index or sequence number. Fluid Index: 1: first fluid pumped                           
	// (=original mud)                                                                                                
	// (last-1)=Tail cement, last= displacement mud                                                                   
	FluidIndex                                                                                  int64                 `json:"FluidIndex"`
	// Fluid loss dehydrating test period, used to compute the API fluid loss.                                        
	FluidLossElapsedTime                                                                        *float64              `json:"FluidLossElapsedTime,omitempty"`
	// Fluid loss pressure                                                                                            
	FluidLossPressure                                                                           *float64              `json:"FluidLossPressure,omitempty"`
	// Fluid loss temperature                                                                                         
	FluidLossTemperature                                                                        *float64              `json:"FluidLossTemperature,omitempty"`
	// Fluid loss volume                                                                                              
	FluidLossVolume                                                                             *float64              `json:"FluidLossVolume,omitempty"`
	// Name of Fluid e.g. Lead or Tail Slurry                                                                         
	FluidName                                                                                   *string               `json:"FluidName,omitempty"`
	// Fluid purpose description.                                                                                     
	FluidPurpose                                                                                *string               `json:"FluidPurpose,omitempty"`
	// Fluid Rheological Model - Newtonian/Bingham/Power Law/Herschel Bulkley.                                        
	FluidRheologicalModelID                                                                     *string               `json:"FluidRheologicalModelID,omitempty"`
	// Fluid Supplier.                                                                                                
	FluidSupplier                                                                               *string               `json:"FluidSupplier,omitempty"`
	// True Vertical Depth at top of fluid placement at the end of Stage.                                             
	FluidTopTrueVerticalDepth                                                                   *float64              `json:"FluidTopTrueVerticalDepth,omitempty"`
	// High level fluid type                                                                                          
	FluidTypeID                                                                                 *string               `json:"FluidTypeID,omitempty"`
	// Foam Injection Pressure                                                                                        
	FoamInjectionPressure                                                                       *float64              `json:"FoamInjectionPressure,omitempty"`
	// Test free fluid (mL/250ML) API 10B-2                                                                           
	FreeFluidPercent                                                                            *float64              `json:"FreeFluidPercent,omitempty"`
	// Test free fluid temperature                                                                                    
	FreeFluidTemperature                                                                        *float64              `json:"FreeFluidTemperature,omitempty"`
	// Free Fluid Test Angle e.g. 0  or 45 degrees                                                                    
	FreeFluidTestAngle                                                                          *float64              `json:"FreeFluidTestAngle,omitempty"`
	// Gas type used for foam job.                                                                                    
	GasFoamTypeID                                                                               *string               `json:"GasFoamTypeID,omitempty"`
	// Volume of gas used for foam job                                                                                
	GasFoamVolume                                                                               *float64              `json:"GasFoamVolume,omitempty"`
	// Gas foam method ratio                                                                                          
	GasMethodAverageRatio                                                                       *float64              `json:"GasMethodAverageRatio,omitempty"`
	// Gas foam method: final method ratio                                                                            
	GasMethodEndRatio                                                                           *float64              `json:"GasMethodEndRatio,omitempty"`
	// Gas foam method: initial method ratio                                                                          
	GasMethodStartRatio                                                                         *float64              `json:"GasMethodStartRatio,omitempty"`
	// Gauge Hole Size for estimating annular volume                                                                  
	GaugeHoleSize                                                                               *float64              `json:"GaugeHoleSize,omitempty"`
	// Gel reading after 10 minutes.                                                                                  
	Gel10MinReading                                                                             *float64              `json:"Gel10MinReading,omitempty"`
	// Gel strength after 10 minutes.                                                                                 
	Gel10MinStrength                                                                            *float64              `json:"Gel10MinStrength,omitempty"`
	// Gel reading after 10 seconds.                                                                                  
	Gel10SECReading                                                                             *float64              `json:"Gel10SecReading,omitempty"`
	// Gel strength after 10 seconds.                                                                                 
	Gel10SECStrength                                                                            *float64              `json:"Gel10SecStrength,omitempty"`
	// Gel reading after 1 minute.                                                                                    
	Gel1MinReading                                                                              *float64              `json:"Gel1MinReading,omitempty"`
	// Gel strength after 1 minute.                                                                                   
	Gel1MinStrength                                                                             *float64              `json:"Gel1MinStrength,omitempty"`
	// Gel reading after 30 minutes.                                                                                  
	Gel30MinReading                                                                             *float64              `json:"Gel30MinReading,omitempty"`
	// Gel strength after 30 minutes.                                                                                 
	Gel30MinStrength                                                                            *float64              `json:"Gel30MinStrength,omitempty"`
	// Foam used indicator                                                                                            
	IsFoamUsed                                                                                  *bool                 `json:"IsFoamUsed,omitempty"`
	// Consistency index (Power Law and Herschel-Bulkley).                                                            
	K                                                                                           *float64              `json:"K,omitempty"`
	// Mixability (0 - 5) - 0 is not mixable                                                                          
	MixabilityRating                                                                            *int64                `json:"MixabilityRating,omitempty"`
	// Mix Fluid Concentration                                                                                        
	MixFluidConcentration                                                                       *float64              `json:"MixFluidConcentration,omitempty"`
	// Mix Fluid Ratio                                                                                                
	MixFluidRatio                                                                               *float64              `json:"MixFluidRatio,omitempty"`
	// Mix Fluid Type                                                                                                 
	MixFluidTypeID                                                                              *string               `json:"MixFluidTypeID,omitempty"`
	// Detailed level fluid type                                                                                      
	MudTypeID                                                                                   *string               `json:"MudTypeID,omitempty"`
	// Power Law index (Power Law and Herschel-Bulkley).                                                              
	N                                                                                           *float64              `json:"N,omitempty"`
	// Open Hole Excess Slurry Volume Percent                                                                         
	OpenHoleExcessPercent                                                                       *float64              `json:"OpenHoleExcessPercent,omitempty"`
	// Volume Pumped                                                                                                  
	PumpedVolume                                                                                *float64              `json:"PumpedVolume,omitempty"`
	// Reserved (Unpumped) Volume                                                                                     
	ReservedVolume                                                                              *float64              `json:"ReservedVolume,omitempty"`
	// Fluid Fann Viscometer Rheology                                                                                 
	Rheometer                                                                                   []FluidFannRheology   `json:"Rheometer,omitempty"`
	// Measured depth interval between the top and base of the slurry placement.                                      
	SlurryHeight                                                                                *float64              `json:"SlurryHeight,omitempty"`
	// Slurry Mix Method                                                                                              
	SlurryMixMethodID                                                                           *string               `json:"SlurryMixMethodID,omitempty"`
	// Slurry Type                                                                                                    
	SlurryTypeID                                                                                *string               `json:"SlurryTypeID,omitempty"`
	// Solid Volume Per Volume Measure                                                                                
	SolidVolumeFraction                                                                         *float64              `json:"SolidVolumeFraction,omitempty"`
	// Source Water Chlorides                                                                                         
	SourceWaterChlorides                                                                        *float64              `json:"SourceWaterChlorides,omitempty"`
	// Water source description                                                                                       
	SourceWaterDescription                                                                      *string               `json:"SourceWaterDescription,omitempty"`
	// Source Water pH                                                                                                
	SourceWaterpH                                                                               *float64              `json:"SourceWaterpH,omitempty"`
	// Source Water pH Temperature                                                                                    
	SourceWaterpHTemperature                                                                    *float64              `json:"SourceWaterpHTemperature,omitempty"`
	// Source Water Temperature                                                                                       
	SourceWaterTemperature                                                                      *float64              `json:"SourceWaterTemperature,omitempty"`
	// Thickening Test - Pressure, Temperature, Consistency (Bc) v Elapsed Time                                       
	ThickeningTimeTest                                                                          []ThickeningTimeTest  `json:"ThickeningTimeTest,omitempty"`
	// Estimated Measured depth at top of fluid placement at end of Stage.                                            
	TopMeasuredDepth                                                                            *float64              `json:"TopMeasuredDepth,omitempty"`
	// Total Fluid/Slurry Volume. If CementingFluid is NOT a cement slurry, then report the                           
	// total fluid volume created of the fluid in question.                                                           
	TotalFluidVolume                                                                            *float64              `json:"TotalFluidVolume,omitempty"`
	// Total Slurry Volume                                                                                            
	TotalSlurryVolume                                                                           *float64              `json:"TotalSlurryVolume,omitempty"`
	// The elapsed time between the development of 100lbf/100sq ft gel strength and 500lbf/100                        
	// sq ft gel strength.                                                                                            
	TransitionElapsedTime                                                                       *float64              `json:"TransitionElapsedTime,omitempty"`
	// Units/Amount of Dry Blend Calculated                                                                           
	UnitsCalculated                                                                             *int64                `json:"UnitsCalculated,omitempty"`
	// Units/Amount of Dry Blend Used                                                                                 
	UnitsUsed                                                                                   *int64                `json:"UnitsUsed,omitempty"`
	// Viscosity (Newtonian) or Plastic Viscosity if Bingham. Not relevant for other rheological                      
	// models.                                                                                                        
	Viscosity                                                                                   *float64              `json:"Viscosity,omitempty"`
	// Yield point (Bingham and Herschel Bulkley models).                                                             
	YieldPoint                                                                                  *float64              `json:"YieldPoint,omitempty"`
	// Cement Slurry Yield Volume. Calculated - Sacks Used * CementYield                                              
	YieldVolume                                                                                 *float64              `json:"YieldVolume,omitempty"`
	// The elapsed time from initiation of the static portion of the test until the slurry                            
	// attains a gel strength of 100lbf/100sq ft.                                                                     
	ZeroGelElapsedTime                                                                          *float64              `json:"ZeroGelElapsedTime,omitempty"`
}
