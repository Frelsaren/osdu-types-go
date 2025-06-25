package masterdata

// Displaced Mud, washes and spacers, cements, displacement mud.
type CementingFluid struct {
	// Volume of given fluid that resides in the annulus.                                                      
	AnnularFluidVolume                                                                     *float64            `json:"AnnularFluidVolume,omitempty"`
	// Density of base fluid.                                                                                  
	BaseFluidDensity                                                                       *float64            `json:"BaseFluidDensity,omitempty"`
	// Type of base fluid: Fresh water, Sea water, Brine, Brackish water.                                      
	BaseFluidType                                                                          *string             `json:"BaseFluidType,omitempty"`
	// Additives can be added in slurry but also in spacers, washes, mud.                                      
	CementAdditives                                                                        []CementAdditive    `json:"CementAdditives,omitempty"`
	// Set of (Time / Rate / Back Pressure).                                                                   
	CementPumpSchedule                                                                     *CementPumpSchedule `json:"CementPumpSchedule,omitempty"`
	// Volume of cement.                                                                                       
	CementVolume                                                                           *float64            `json:"CementVolume,omitempty"`
	// Slurry class.                                                                                           
	ClassSlurryDryBlend                                                                    *string             `json:"ClassSlurryDryBlend,omitempty"`
	// Slurry density at pressure.                                                                             
	DensAtPres                                                                             *float64            `json:"DensAtPres,omitempty"`
	// Fluid density.                                                                                          
	Density                                                                                *float64            `json:"Density,omitempty"`
	// Density of Dry blend.                                                                                   
	DryBlendDensity                                                                        *float64            `json:"DryBlendDensity,omitempty"`
	// Description of dry blend.                                                                               
	DryBlendDescription                                                                    *string             `json:"DryBlendDescription,omitempty"`
	// Mass of dry blend: the blend is made of different solid additives: the volume is not                    
	// constant.                                                                                               
	DryBlendMass                                                                           *float64            `json:"DryBlendMass,omitempty"`
	// Weight of a sack of dry blend.                                                                          
	DryBlendMassSack                                                                       *float64            `json:"DryBlendMassSack,omitempty"`
	// Name of dry blend.                                                                                      
	DryBlendName                                                                           *string             `json:"DryBlendName,omitempty"`
	// Excess Percent.                                                                                         
	ExcessPercent                                                                          *float64            `json:"ExcessPercent,omitempty"`
	// Measured depth at bottom of slurry placement.                                                           
	FluidBaseMeasuredDepth                                                                 *float64            `json:"FluidBaseMeasuredDepth,omitempty"`
	// True Vertical Depth at bottom of fluid placement at the end of Stage.                                   
	FluidBaseTrueVerticalDepth                                                             *float64            `json:"FluidBaseTrueVerticalDepth,omitempty"`
	// Fluid description.                                                                                      
	FluidDescription                                                                       *string             `json:"FluidDescription,omitempty"`
	// Fluid Index: 1: first fluid pumped (=original mud),                                                     
	// (last-1)=Tail cement, last= displacement mud                                                            
	FluidIndex                                                                             *float64            `json:"FluidIndex,omitempty"`
	// Newtonian/Bingham/Power Law/Herschel Bulkley.                                                           
	FluidRheologicalModelID                                                                *string             `json:"FluidRheologicalModelID,omitempty"`
	// Estimated Measured depth at top of slurry placement.                                                    
	FluidTopMeasuredDepth                                                                  *float64            `json:"FluidTopMeasuredDepth,omitempty"`
	// True Vertical Depth at top of fluid placement at the end of Stage.                                      
	FluidTopTrueVerticalDepth                                                              *float64            `json:"FluidTopTrueVerticalDepth,omitempty"`
	// Fluid/Slurry Volume.                                                                                    
	FluidVolume                                                                            *float64            `json:"FluidVolume,omitempty"`
	// Gas type used for foam job.                                                                             
	GasFoamType                                                                            *string             `json:"GasFoamType,omitempty"`
	// Volume of gas used for foam job.                                                                        
	GasFoamVol                                                                             *float64            `json:"GasFoamVol,omitempty"`
	// Gel reading after 10 minutes.                                                                           
	Gel10MinReading                                                                        *float64            `json:"Gel10MinReading,omitempty"`
	// Gel strength after 10 minutes.                                                                          
	Gel10MinStrength                                                                       *float64            `json:"Gel10MinStrength,omitempty"`
	// Gel reading after 10 seconds.                                                                           
	Gel10SECReading                                                                        *float64            `json:"Gel10SecReading,omitempty"`
	// Gel strength after 10 seconds.                                                                          
	Gel10SECStrength                                                                       *float64            `json:"Gel10SecStrength,omitempty"`
	// Gel reading after 1 minute.                                                                             
	Gel1MinReading                                                                         *float64            `json:"Gel1MinReading,omitempty"`
	// Gel strength after 1 minute.                                                                            
	Gel1MinStrength                                                                        *float64            `json:"Gel1MinStrength,omitempty"`
	// Foam used indicator.  Values are "true" (or "1") and "false" (or "0").                                  
	IsFoamUsed                                                                             *bool               `json:"IsFoamUsed,omitempty"`
	// Consistency index (Power Law and HB).                                                                   
	K                                                                                      *float64            `json:"K,omitempty"`
	// Power Law index (Power Law and HB).                                                                     
	N                                                                                      *float64            `json:"N,omitempty"`
	// Other Volume.                                                                                           
	OtherVolume                                                                            *float64            `json:"OtherVolume,omitempty"`
	// Volume Pumped.                                                                                          
	PumpedVolume                                                                           *float64            `json:"PumpedVolume,omitempty"`
	// Purpose description.                                                                                    
	Purpose                                                                                *string             `json:"Purpose,omitempty"`
	// Mix Water Ratio.                                                                                        
	RatioMixWater                                                                          *float64            `json:"RatioMixWater,omitempty"`
	// Fluid type: Mud, Wash, Spacer, Slurry.                                                                  
	TypeFluid                                                                              *string             `json:"TypeFluid,omitempty"`
	// Viscosity (Newtonian) or Plastic Viscosity if Bingham.                                                  
	Viscosity                                                                              *float64            `json:"Viscosity,omitempty"`
	// Water volume.                                                                                           
	WaterVolume                                                                            *float64            `json:"WaterVolume,omitempty"`
	// Yield point (Bingham and Herschel Bulkley models).                                                      
	YieldPoint                                                                             *float64            `json:"YieldPoint,omitempty"`
	// Slurry Yield.                                                                                           
	YieldVolume                                                                            *float64            `json:"YieldVolume,omitempty"`
}
