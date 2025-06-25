package workproductcomponent

import "time"

// A pumping diagnostics session for this job stage.
type StimJobDiagnosticSession struct {
	// Base fluid volume entering the pumping equipment.                                                                  
	BaseFluidVol                                                                                *float64                  `json:"BaseFluidVol,omitempty"`
	// Bottomhole hydrostatic pressure.                                                                                   
	BHHydrostaticPressure                                                                       *float64                  `json:"BHHydrostaticPressure,omitempty"`
	// The measured depth of the bottom of the hole.                                                                      
	BHMeasuredDepth                                                                             *float64                  `json:"BHMeasuredDepth,omitempty"`
	// Average bottomhole treatment pressure.                                                                             
	BHTreatmentPressureAvg                                                                      *float64                  `json:"BHTreatmentPressureAvg,omitempty"`
	// Average bottomhole treatment flow rate.                                                                            
	BHTreatmentRateAvg                                                                          *float64                  `json:"BHTreatmentRateAvg,omitempty"`
	// The pressure at which gas begins to break out of an under saturated oil and form a free                            
	// gas phase in the matrix or a gas cap.                                                                              
	BubblePointPressure                                                                         *float64                  `json:"BubblePointPressure,omitempty"`
	// The size of the choke used during a flow back test.                                                                
	ChokeSize                                                                                   *float64                  `json:"ChokeSize,omitempty"`
	// A description of this pumping diagnostic session.                                                                  
	Description                                                                                 *string                   `json:"Description,omitempty"`
	// The volume change of a fluid when pressure is applied.                                                             
	FluidCompressibility                                                                        *float64                  `json:"FluidCompressibility,omitempty"`
	// The density of the fluid during this pumping diagnostic session.                                                   
	FluidDensity                                                                                *float64                  `json:"FluidDensity,omitempty"`
	// A measurement, derived from a data frac, of the efficiency of a particular fluid in                                
	// creating fracture area on a particular formation at a set of conditions.                                           
	FluidEfficiency                                                                             *float64                  `json:"FluidEfficiency,omitempty"`
	// A diagnostic test determining fluid efficiency.                                                                    
	FluidEfficiencyTest                                                                         []StimFluidEfficiencyTest `json:"FluidEfficiencyTest,omitempty"`
	// The consistency index K is the shear stress or viscosity of the fluid at one sec-1 shear                           
	// rate. An increasing K raises the effective viscosity.                                                              
	FluidKprimeFactor                                                                           *float64                  `json:"FluidKprimeFactor,omitempty"`
	// Power law component. As 'n' decreases from 1, the fluid becomes more shear thinning.                               
	// Reducing 'n' produces more non-Newtonian behavior.                                                                 
	FluidNprimeFactor                                                                           *float64                  `json:"FluidNprimeFactor,omitempty"`
	// The heat required to raise one unit mass of a substance by one degree.                                             
	FluidSpecificHeat                                                                           *float64                  `json:"FluidSpecificHeat,omitempty"`
	// In physics, thermal conductivity is the property of a material describing its ability to                           
	// conduct heat. It appears primarily in Fourier's Law for heat conduction. Thermal                                   
	// conductivity is measured in watts per kelvin-meter. Multiplied by a temperature                                    
	// difference (in kelvins) and an area (in square meters), and divided by a thickness (in                             
	// meters), the thermal conductivity predicts the rate of energy loss (in watts) through a                            
	// piece of material.                                                                                                 
	FluidThermalConductivity                                                                    *float64                  `json:"FluidThermalConductivity,omitempty"`
	// Dimensional response to temperature change is expressed by its coefficient of thermal                              
	// expansion. When the temperature of a substance changes, the energy that is stored in the                           
	// intermolecular bonds between atoms also changes. When the stored energy increases, so                              
	// does the length of the molecular bonds. As a result, solids typically expand in response                           
	// to heating and contract on cooling. The degree of expansion divided by the change in                               
	// temperature is called the material's coefficient of thermal expansion and generally                                
	// varies with temperature.                                                                                           
	FluidThermalExpansionCoefficient                                                            *float64                  `json:"FluidThermalExpansionCoefficient,omitempty"`
	// Foam quality percentage of foam for the pumping diagnostic session during the stimulation                          
	// services.                                                                                                          
	FoamQuality                                                                                 *float64                  `json:"FoamQuality,omitempty"`
	// The date and time when the fluid in the fracture is completely leaked off into the                                 
	// formation and the fracture closes on its faces.                                                                    
	FractureCloseDateTime                                                                       *time.Time                `json:"FractureCloseDateTime,omitempty"`
	// The pressure when the fracture width becomes zero.                                                                 
	FractureClosePressure                                                                       *float64                  `json:"FractureClosePressure,omitempty"`
	// The pressure loss due to fluid friction with the pipe while a fluid is being pumped.                               
	FrictionPressure                                                                            *float64                  `json:"FrictionPressure,omitempty"`
	// The measured depth of the wellbore to its injection point.                                                         
	InjectionPointMeasuredDepth                                                                 *float64                  `json:"InjectionPointMeasuredDepth,omitempty"`
	// Are the calculations corrected for temperature? A value of true (or 1) indicates that the                          
	// calculations were corrected for temperature. A value of false (or 0) or not given                                  
	// indicates otherwise.                                                                                               
	IsTemperatureCorrectionApplied                                                              *bool                     `json:"IsTemperatureCorrectionApplied,omitempty"`
	// The measured depth of the middle perforation (MPP).                                                                
	MidPerforationMeasuredDepth                                                                 *float64                  `json:"MidPerforationMeasuredDepth,omitempty"`
	// The true vertical depth of the middle perforation.                                                                 
	MidPerforationTrueVerticalDepth                                                             *float64                  `json:"MidPerforationTrueVerticalDepth,omitempty"`
	// The name of this pumping diagnostic session.                                                                       
	Name                                                                                        *string                   `json:"Name,omitempty"`
	// The number of this pumping diagnostics session.                                                                    
	Number                                                                                      *int64                    `json:"Number,omitempty"`
	// The volume of the pad divided by the (volume of the pad + the volume of the proppant                               
	// laden fluid).                                                                                                      
	PadPercent                                                                                  *float64                  `json:"PadPercent,omitempty"`
	// The pressure of the liquids in the formation pore space.                                                           
	PorePressure                                                                                *float64                  `json:"PorePressure,omitempty"`
	// The time between the shutin time and the pump on time.                                                             
	PumpDuration                                                                                *float64                  `json:"PumpDuration,omitempty"`
	// A diagnostic test involving flowing a well back after treatment.                                                   
	PumpFlowBackTest                                                                            []StimPumpFlowBackTest    `json:"PumpFlowBackTest,omitempty"`
	// The date and time pumping ended.                                                                                   
	PumpOffDateTime                                                                             *time.Time                `json:"PumpOffDateTime,omitempty"`
	// The date and time pumping began.                                                                                   
	PumpOnDateTime                                                                              *time.Time                `json:"PumpOnDateTime,omitempty"`
	// Remarks                                                                                                            
	Remarks                                                                                     *string                   `json:"Remarks,omitempty"`
	// The volume change of a reservoir material when pressure is applied.                                                
	ReservoirTotalCompressibility                                                               *float64                  `json:"ReservoirTotalCompressibility,omitempty"`
	// Initial shutin pressure.                                                                                           
	ShutInInitialPressure                                                                       *float64                  `json:"ShutInInitialPressure,omitempty"`
	// The number of the stage associated with this diagnostics session.                                                  
	StageNumber                                                                                 *int64                    `json:"StageNumber,omitempty"`
	// Static bottomhole temperature.                                                                                     
	StaticBHTemperature                                                                         *float64                  `json:"StaticBHTemperature,omitempty"`
	// An injection test involving multiple steps of injection rate and pressure, where a curve                           
	// deflection and change of slope indicates the fracture breakdown pressure. An injection                             
	// test involving multiple steps of injection rate and pressure, where a curve deflection                             
	// and change of slope indicates the fracture breakdown pressure.                                                     
	StepDownTest                                                                                []StimStepDownTest        `json:"StepDownTest,omitempty"`
	// An injection test, plotted pressure against injection rate, where a curve deflection and                           
	// change of slope indicates the fracture breakdown pressure.                                                         
	StepRateTest                                                                                []StimStepTest            `json:"StepRateTest,omitempty"`
	// Temperature of the fluid at the surface.                                                                           
	SurfaceFluidTemperature                                                                     *float64                  `json:"SurfaceFluidTemperature,omitempty"`
	// The temperature at surface during this pumping diagnostic session.                                                 
	SurfaceTemperature                                                                          *float64                  `json:"SurfaceTemperature,omitempty"`
	// The volume of fluid in the wellbore for this pumping diagnostic session.                                           
	WellboreVolume                                                                              *float64                  `json:"WellboreVolume,omitempty"`
	// The date and time at which a well ceases flowing and the valves are closed.                                        
	WellShutinDateTime                                                                          *time.Time                `json:"WellShutinDateTime,omitempty"`
}
