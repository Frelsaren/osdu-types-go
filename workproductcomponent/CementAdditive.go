package workproductcomponent

// Additives/chemicals added to a Fluid during a cement job
type CementAdditive struct {
	// Additive Amount Used                                                                              
	AdditiveAmount                                                                              *float64 `json:"AdditiveAmount,omitempty"`
	// Additive Amount Unit                                                                              
	AdditiveAmountUoM                                                                           *string  `json:"AdditiveAmountUoM,omitempty"`
	// Additive Code                                                                                     
	AdditiveCode                                                                                *string  `json:"AdditiveCode,omitempty"`
	// Concentration Amount: unit type depends on typeConc.                                              
	AdditiveConcentrationAmount                                                                 *float64 `json:"AdditiveConcentrationAmount,omitempty"`
	// Additive Concentration Unit Of Measure                                                            
	AdditiveConcentrationUnitOfMeasureID                                                        *string  `json:"AdditiveConcentrationUnitOfMeasureID,omitempty"`
	// Additive density                                                                                  
	AdditiveDensity                                                                             *float64 `json:"AdditiveDensity,omitempty"`
	// Additive Formulation - Wet or Dry.                                                                
	AdditiveFormulationID                                                                       *string  `json:"AdditiveFormulationID,omitempty"`
	// Additive index or sequence number. Used to retain order.                                          
	AdditiveIndex                                                                               int64    `json:"AdditiveIndex"`
	// Volume of container                                                                               
	AdditiveLiquidVolume                                                                        *float64 `json:"AdditiveLiquidVolume,omitempty"`
	// ID of or URL to Material Safety Data Sheet                                                        
	AdditiveMSDSID                                                                              *string  `json:"AdditiveMSDSID,omitempty"`
	// Relationship to a FluidAdditiveName reference-data record.                                        
	AdditiveNameID                                                                              string   `json:"AdditiveNameID"`
	// Comments or Remarks.                                                                              
	AdditiveRemark                                                                              *string  `json:"AdditiveRemark,omitempty"`
	// The chief purpose or reason for adding a substance to a fluid used in a downhole                  
	// operation.                                                                                        
	AdditiveRoleID                                                                              *string  `json:"AdditiveRoleID,omitempty"`
	// Mass of Sack                                                                                      
	AdditiveSackMass                                                                            *float64 `json:"AdditiveSackMass,omitempty"`
	// ConcentrationType: %BWOC (%By weight of Cement), %BWOB (%By weight of blend), %BWOW (%By          
	// weight of water), %BWOBF (%By weight of base fluid)                                               
	AdditiveTypeConcentrationID                                                                 *string  `json:"AdditiveTypeConcentrationID,omitempty"`
	// The liquid, solid or gaseous substance used to manipulate the function or properties of a         
	// fluid.                                                                                            
	AdditiveTypeID                                                                              *string  `json:"AdditiveTypeID,omitempty"`
	// Vendor/Supplier of additive                                                                       
	AdditiveVendorID                                                                            *string  `json:"AdditiveVendorID,omitempty"`
}
