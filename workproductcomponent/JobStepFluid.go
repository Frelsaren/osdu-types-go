package workproductcomponent

// Fluid used for this job step of the stimulated interval.
type JobStepFluid struct {
	// Concentrations of additives in the fluid for this job step.                                                      
	AdditiveConcentration                                                                []AbstractStimMaterialQuantity `json:"AdditiveConcentration,omitempty"`
	// The density of the fluid.                                                                                        
	Density                                                                              *float64                       `json:"Density,omitempty"`
	// The description of the fluid.                                                                                    
	Description                                                                          *string                        `json:"Description,omitempty"`
	// Filter Cake                                                                                                      
	FilterCake                                                                           *float64                       `json:"FilterCake,omitempty"`
	// The filter size that the fluid for this job step passes through.                                                 
	FilterSize                                                                           *float64                       `json:"FilterSize,omitempty"`
	// The type of fluid filter used for this job step.                                                                 
	FilterType                                                                           *string                        `json:"FilterType,omitempty"`
	// The temperature of the fluid at surface.                                                                         
	FluidSurfaceTemperature                                                              *float64                       `json:"FluidSurfaceTemperature,omitempty"`
	// Test pressure of the fluid during this job step.                                                                 
	FluidTestPressure                                                                    *float64                       `json:"FluidTestPressure,omitempty"`
	// Test temperature of the fluid for this job step.                                                                 
	FluidTestTemperature                                                                 *float64                       `json:"FluidTestTemperature,omitempty"`
	// The fluid type.                                                                                                  
	FluidTypeID                                                                          *string                        `json:"FluidTypeID,omitempty"`
	// The shear stress measured at low shear rate after a mud has set quiescently for 10                               
	// minutes.                                                                                                         
	GelStrength10Min                                                                     *float64                       `json:"GelStrength10Min,omitempty"`
	// The shear stress measured at low shear rate after a mud has set quiescently for 10                               
	// seconds.                                                                                                         
	GelStrength10SEC                                                                     *float64                       `json:"GelStrength10Sec,omitempty"`
	// Is the fluid a kill fluid?                                                                                       
	IsKillFluid                                                                          *bool                          `json:"IsKillFluid,omitempty"`
	// The name of the fluid.                                                                                           
	Name                                                                                 *string                        `json:"Name,omitempty"`
	// The pH of the fluid.                                                                                             
	PH                                                                                   *float64                       `json:"pH,omitempty"`
	// The purpose of the fluid.                                                                                        
	Purpose                                                                              *string                        `json:"Purpose,omitempty"`
	// Remarks                                                                                                          
	Remarks                                                                              *string                        `json:"Remarks,omitempty"`
	// The specific gravity of the fluid at surface.                                                                    
	SpecificGravity                                                                      *float64                       `json:"SpecificGravity,omitempty"`
	// The name of the fluid supplier for this job step.                                                                
	Supplier                                                                             *string                        `json:"Supplier,omitempty"`
	// The supplier of the fluid.                                                                                       
	SupplierID                                                                           *string                        `json:"SupplierID,omitempty"`
	// Viscosity of the stimulation fluid for this job step.                                                            
	Viscosity                                                                            *float64                       `json:"Viscosity,omitempty"`
}
