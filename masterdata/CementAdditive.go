package masterdata

// Additives used during a cement job
type CementAdditive struct {
	// Additive Amount.                                                                                  
	AdditiveAmount                                                                              float64  `json:"AdditiveAmount"`
	// Additive density.                                                                                 
	AdditiveDensity                                                                             *float64 `json:"AdditiveDensity,omitempty"`
	// Wet or Dry.                                                                                       
	AdditiveFormulation                                                                         *string  `json:"AdditiveFormulation,omitempty"`
	// Additive name.                                                                                    
	AdditiveName                                                                                string   `json:"AdditiveName"`
	// The chief purpose or reason for adding a substance to a fluid used in a downhole                  
	// operation.                                                                                        
	AdditiveRoleID                                                                              *string  `json:"AdditiveRoleID,omitempty"`
	// The liquid, solid or gaseous substance used to manipulate the function or properties of a         
	// fluid.                                                                                            
	AdditiveTypeID                                                                              *string  `json:"AdditiveTypeID,omitempty"`
	// Vendor/Supplier of additive                                                                       
	AdditiveVendorID                                                                            *string  `json:"AdditiveVendorID,omitempty"`
	// Concentration Amount: unit type depends of typeConc.                                              
	Concentration                                                                               float64  `json:"Concentration"`
	// ConcentrationType: %BWOC (%By weight of Cement), %BWOB (%By weight of blend), %BWOW (%By          
	// weight of water), %BWOBF (%By weight of base fluid)                                               
	TypeConc                                                                                    string   `json:"TypeConc"`
	// Concentration in terms of volume per sack.                                                        
	VolSack                                                                                     float64  `json:"VolSack"`
	// Concentration in terms of weight per sack.                                                        
	WtSack                                                                                      float64  `json:"WtSack"`
}
