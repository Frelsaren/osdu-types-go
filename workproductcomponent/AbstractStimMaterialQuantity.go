package workproductcomponent

// Usage and maximum mass or flow rates for a material used in the Stage.
type AbstractStimMaterialQuantity struct {
	// The density of the material used.                                                                 
	Density                                                                                     *float64 `json:"Density,omitempty"`
	// The mass of material used.  This should be used without specifying any of the other               
	// material measures (e.g. volume, standard volume, etc.).                                           
	Mass                                                                                        *float64 `json:"Mass,omitempty"`
	// Rate at which mass of material is pumped/added to the fluid.                                      
	MassFlowRate                                                                                *float64 `json:"MassFlowRate,omitempty"`
	// Additive or Proppant when referenced to the Job Material Catalog                                  
	MaterialKindID                                                                              *string  `json:"MaterialKindID,omitempty"`
	// This is a reference to the name of the proppant or fluid additive in the StimJob Material         
	// Catalog.                                                                                          
	MaterialReferenceID                                                                         string   `json:"MaterialReferenceID"`
	// The standard volume of material used. Standard volume is the volume measured under the            
	// same conditions. This should be used without specifying any of the other material                 
	// measures (e.g., mass, volume, etc.).                                                              
	StdVolume                                                                                   *float64 `json:"StdVolume,omitempty"`
	// The volume of material used.  This should be used without specifying any of the other             
	// material measures (e.g. mass, standard volume, etc.).                                             
	Volume                                                                                      *float64 `json:"Volume,omitempty"`
	// The volume per volume measure of material used.  This should be used without specifying           
	// any of the other material measures (e.g. mass, density, standard volume, etc.).                   
	VolumeConcentration                                                                         *float64 `json:"VolumeConcentration,omitempty"`
	// Rate at which the volume of material is pumped/added to the fluid.                                
	VolumetricFlowRate                                                                          *float64 `json:"VolumetricFlowRate,omitempty"`
}
