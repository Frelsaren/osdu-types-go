package workproductcomponent

import "time"

// Information regarding an individual fluid that is part of the Drilling Report.
type FluidElement struct {
	// Mud alkalinity P1 from alternate alkalinity method (volume in ml of 0.02N acid                      
	// to reach the phenolphthalein endpoint).                                                             
	AlkalinityP1                                                                               *float64    `json:"AlkalinityP1,omitempty"`
	// Mud alkalinity P2 from alternate alkalinity method (volume in ml of 0.02N acid to                   
	// titrate, the reagent mixture to the phenolphthalein endpoint).                                      
	AlkalinityP2                                                                               *float64    `json:"AlkalinityP2,omitempty"`
	// Cleaning coefficient of drilling fluid                                                              
	AnnularShapeRatio                                                                          *float64    `json:"AnnularShapeRatio,omitempty"`
	// Average specific gravity of solids.                                                                 
	ASG                                                                                        *float64    `json:"Asg,omitempty"`
	// Average size of the drill cuttings.                                                                 
	AverageCuttingSize                                                                         *float64    `json:"AverageCuttingSize,omitempty"`
	// Barite content percent.                                                                             
	BaritePC                                                                                   *float64    `json:"BaritePc,omitempty"`
	// The name of the fluid as given by the supplier                                                      
	BrandName                                                                                  *string     `json:"BrandName,omitempty"`
	// Density of water phase of NAF.                                                                      
	BrineDensity                                                                               *float64    `json:"BrineDensity,omitempty"`
	// Percent brine content.                                                                              
	BrinePC                                                                                    *float64    `json:"BrinePc,omitempty"`
	// Calcium content.                                                                                    
	Calcium                                                                                    *float64    `json:"Calcium,omitempty"`
	// Calcium chloride content.                                                                           
	CalciumChloride                                                                            *float64    `json:"CalciumChloride,omitempty"`
	// Calcium chloride percent.                                                                           
	CalciumChloridePC                                                                          *float64    `json:"CalciumChloridePc,omitempty"`
	// Carbonate content.                                                                                  
	Carbonate                                                                                  *float64    `json:"Carbonate,omitempty"`
	// Chloride content.                                                                                   
	Chloride                                                                                   *float64    `json:"Chloride,omitempty"`
	// Comments and remarks.                                                                               
	Comments                                                                                   *string     `json:"Comments,omitempty"`
	// Reference to the Company performing the analysis                                                    
	CompanyID                                                                                  *string     `json:"CompanyID,omitempty"`
	// Fluid density. This excludes the cuttings.                                                          
	Density                                                                                    *float64    `json:"Density,omitempty"`
	// Measurement of the emulsion stability and oil-wetting capability in oil-based muds.                 
	ElectStab                                                                                  *float64    `json:"ElectStab,omitempty"`
	// Equivalent circulating density where fluid reading was recorded or calculated/simulated             
	EquivalentCirculatingDensity                                                               *float64    `json:"EquivalentCirculatingDensity,omitempty"`
	// Equivalent static density where fluid reading was recorded or calculated/simulated                  
	EquivalentStaticDensity                                                                    *float64    `json:"EquivalentStaticDensity,omitempty"`
	// High temperature high pressure (HTHP) filter cake thickness.                                        
	FilterCakeHthp                                                                             *float64    `json:"FilterCakeHthp,omitempty"`
	// Filter cake thickness at low (normal) temperature and pressure.                                     
	FilterCakeLtlp                                                                             *float64    `json:"FilterCakeLtlp,omitempty"`
	// High temperature high pressure (HTHP) filtrate (volume per 30 min).                                 
	FiltrateHthp                                                                               *float64    `json:"FiltrateHthp,omitempty"`
	// API water loss (low temperature and pressure mud filtrate measurement) (volume per 30               
	// min).                                                                                               
	FiltrateLtlp                                                                               *float64    `json:"FiltrateLtlp,omitempty"`
	// High temperature high pressure (HTHP) pressure.                                                     
	FiltratePressureHthp                                                                       *float64    `json:"FiltratePressureHthp,omitempty"`
	// High temperature high pressure (HTHP) temperature.                                                  
	FiltrateTemperatureHthp                                                                    *float64    `json:"FiltrateTemperatureHthp,omitempty"`
	// Zero-second gels.                                                                                   
	Gel0SEC                                                                                    *float64    `json:"Gel0Sec,omitempty"`
	// Ten-minute gels.                                                                                    
	Gel10Min                                                                                   *float64    `json:"Gel10Min,omitempty"`
	// Ten-second gels.                                                                                    
	Gel10SEC                                                                                   *float64    `json:"Gel10Sec,omitempty"`
	// Thirty-minute gels.                                                                                 
	Gel30Min                                                                                   *float64    `json:"Gel30Min,omitempty"`
	// Total calcium hardness.                                                                             
	HardnessCA                                                                                 *float64    `json:"HardnessCa,omitempty"`
	// Iron content.                                                                                       
	Iron                                                                                       *float64    `json:"Iron,omitempty"`
	// Assumed kick density for calculation of kick tolerance where the fluid reading was                  
	// recorded.                                                                                           
	KickToleranceIntensity                                                                     *float64    `json:"KickToleranceIntensity,omitempty"`
	// Assumed kick volume for calculation of kick tolerance based on the kick intensity where             
	// the fluid reading was recorded.                                                                     
	KickToleranceVolume                                                                        *float64    `json:"KickToleranceVolume,omitempty"`
	// Lost circulation material.                                                                          
	Lcm                                                                                        *float64    `json:"Lcm,omitempty"`
	// Lime content.                                                                                       
	Lime                                                                                       *float64    `json:"Lime,omitempty"`
	// Shearing capacity of mud at low velocity                                                            
	LowShearYieldPoint                                                                         *float64    `json:"LowShearYieldPoint,omitempty"`
	// Magnesium content.                                                                                  
	Magnesium                                                                                  *float64    `json:"Magnesium,omitempty"`
	// Cation exchange capacity (CEC) of the mud sample as measured by methylene blue titration            
	// (MBT).                                                                                              
	//                                                                                                     
	// NOTE: This is temporarily set to be a GenericMeasure with no unit validation, pending               
	// addition of CEC units to the Energistics UoM spec.                                                  
	Mbt                                                                                        *float64    `json:"Mbt,omitempty"`
	// The measured depth where the fluid readings were recorded.                                          
	MeasuredDepth                                                                              *float64    `json:"MeasuredDepth,omitempty"`
	// Metal recovered from the wellbore.                                                                  
	MetalRecovered                                                                             *float64    `json:"MetalRecovered,omitempty"`
	// Methyl orange alkalinity of filtrate.                                                               
	MF                                                                                         *float64    `json:"Mf,omitempty"`
	// The class of the drilling fluid.                                                                    
	MudClass                                                                                   *string     `json:"MudClass,omitempty"`
	// The name of the Mud Engineer                                                                        
	MudEngineer                                                                                *string     `json:"MudEngineer,omitempty"`
	// Mud pH.                                                                                             
	MudPh                                                                                      *float64    `json:"MudPh,omitempty"`
	// Mud pH measurement temperature.                                                                     
	MudTempPh                                                                                  *float64    `json:"MudTempPh,omitempty"`
	// Oil on cuttings.                                                                                    
	OilCtg                                                                                     *float64    `json:"OilCtg,omitempty"`
	// Oil on dried cuttings.                                                                              
	OilCtgDry                                                                                  *float64    `json:"OilCtgDry,omitempty"`
	// Oil and grease content.                                                                             
	OilGrease                                                                                  *float64    `json:"OilGrease,omitempty"`
	// Percent oil content from retort.                                                                    
	OilPC                                                                                      *float64    `json:"OilPc,omitempty"`
	// Plastic viscosity.                                                                                  
	PlasticViscosity                                                                           *float64    `json:"PlasticViscosity,omitempty"`
	// Phenolphthalein alkalinity of whole mud.                                                            
	Pm                                                                                         *float64    `json:"Pm,omitempty"`
	// Phenolphthalein alkalinity of mud filtrate.                                                         
	PmFiltrate                                                                                 *float64    `json:"PmFiltrate,omitempty"`
	// Polymers present in the mud system.                                                                 
	Polymer                                                                                    *float64    `json:"Polymer,omitempty"`
	// Type of polymers present in the mud system.                                                         
	PolymerType                                                                                []string    `json:"PolymerType,omitempty"`
	// Potassium content.                                                                                  
	Potassium                                                                                  *float64    `json:"Potassium,omitempty"`
	// Details of the Rheometer tests performed on the fluid sample                                        
	Rheometer                                                                                  []Rheometer `json:"Rheometer,omitempty"`
	// Salt content.                                                                                       
	Salt                                                                                       *float64    `json:"Salt,omitempty"`
	// Salt percent.                                                                                       
	SaltPC                                                                                     *float64    `json:"SaltPc,omitempty"`
	// The time when fluid readings were recorded.                                                         
	SampleDateTime                                                                             *time.Time  `json:"SampleDateTime,omitempty"`
	// Sample location.                                                                                    
	SampleLocation                                                                             *string     `json:"SampleLocation,omitempty"`
	// Sand content percent.                                                                               
	SandPC                                                                                     *float64    `json:"SandPc,omitempty"`
	// Sodium chloride content.                                                                            
	SodiumChloride                                                                             *float64    `json:"SodiumChloride,omitempty"`
	// Sodium chloride percent.                                                                            
	SodiumChloridePC                                                                           *float64    `json:"SodiumChloridePc,omitempty"`
	// Solids corrected for chloride content percent.                                                      
	SolCorPC                                                                                   *float64    `json:"SolCorPc,omitempty"`
	// Percent calculated solids content.                                                                  
	SolidsCalcPC                                                                               *float64    `json:"SolidsCalcPc,omitempty"`
	// Solids high gravity content.                                                                        
	SolidsHiGrav                                                                               *float64    `json:"SolidsHiGrav,omitempty"`
	// Solids high gravity content corrected for chloride content                                          
	SolidsHiGravCorrected                                                                      *float64    `json:"SolidsHiGravCorrected,omitempty"`
	// Solids high gravity percent.                                                                        
	SolidsHiGravPC                                                                             *float64    `json:"SolidsHiGravPc,omitempty"`
	// Solids low gravity content.                                                                         
	SolidsLowGrav                                                                              *float64    `json:"SolidsLowGrav,omitempty"`
	// Solids low gravity content corrected for chloride content                                           
	SolidsLowGravCorrected                                                                     *float64    `json:"SolidsLowGravCorrected,omitempty"`
	// Low gravity solids percent.                                                                         
	SolidsLowGravPC                                                                            *float64    `json:"SolidsLowGravPc,omitempty"`
	// Solids percentage from retort.                                                                      
	SolidsPC                                                                                   *float64    `json:"SolidsPc,omitempty"`
	// Sulfide content.                                                                                    
	Sulfide                                                                                    *float64    `json:"Sulfide,omitempty"`
	// True crystallization temperature.                                                                   
	Tct                                                                                        *float64    `json:"Tct,omitempty"`
	// Flow line temperature measurement where the fluid reading was recorded.                             
	TemperatureFlowLine                                                                        *float64    `json:"TemperatureFlowLine,omitempty"`
	// Funnel viscosity temperature.                                                                       
	TemperatureViscosity                                                                       *float64    `json:"TemperatureViscosity,omitempty"`
	// Turbidity units to measure the cloudiness or haziness of a fluid.                                   
	Turbidity                                                                                  *float64    `json:"Turbidity,omitempty"`
	// The true vertical depth where the fluid readings were recorded.                                     
	Tvd                                                                                        *float64    `json:"TVD,omitempty"`
	// Description for the type of fluid.                                                                  
	Type                                                                                       *string     `json:"Type,omitempty"`
	// Funnel viscosity in seconds.                                                                        
	ViscosityFunnel                                                                            *float64    `json:"ViscosityFunnel,omitempty"`
	// Water content percent.                                                                              
	WaterPC                                                                                    *float64    `json:"WaterPc,omitempty"`
	// A factor showing the activity level of salt in oil-based mud.                                       
	WaterPhaseSalinity                                                                         *float64    `json:"WaterPhaseSalinity,omitempty"`
	// Calcium content in the whole mud sample, including oil and water phases.                            
	WholeMudCalcium                                                                            *float64    `json:"WholeMudCalcium,omitempty"`
	// Chloride content in the whole mud sample, including oil and water phases.                           
	WholeMudChloride                                                                           *float64    `json:"WholeMudChloride,omitempty"`
	// Yield point (Bingham and Herschel Bulkley models).                                                  
	YieldPoint                                                                                 *float64    `json:"YieldPoint,omitempty"`
	// Zinc oxide content.                                                                                 
	ZincOxide                                                                                  *float64    `json:"ZincOxide,omitempty"`
}
