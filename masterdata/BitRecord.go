package masterdata

// Information on the condition of the Drill Bit as recorded during an operations report
type BitRecord struct {
	// N = new, U = used.                                                                                 
	BitClass                                                                                     *string  `json:"BitClass,omitempty"`
	// Bit cost in local currency.                                                                        
	Cost                                                                                         *float64 `json:"Cost,omitempty"`
	// The name of the local currency                                                                     
	CostCurrency                                                                                 *string  `json:"CostCurrency,omitempty"`
	// Diameter of drilled hole.                                                                          
	DiameterBit                                                                                  *float64 `json:"DiameterBit,omitempty"`
	// Minimum hole or tubing which bit will pass through (for bi-center bits).                           
	DiameterPassThrough                                                                          *float64 `json:"DiameterPassThrough,omitempty"`
	// Diameter of pilot bit (for bi-center bits).                                                        
	DiameterPilot                                                                                *float64 `json:"DiameterPilot,omitempty"`
	// Bit drive type (Motor, rotary table etc).                                                          
	DriveType                                                                                    *string  `json:"DriveType,omitempty"`
	// Condition of bit bearings (integer 0-8 or E, F, N or X).                                           
	FinalConditionBearing                                                                        *string  `json:"FinalConditionBearing,omitempty"`
	// Overall dull condition from IADC bit wear 2 character codes.                                       
	FinalConditionDull                                                                           *string  `json:"FinalConditionDull,omitempty"`
	// Condition of bit gauge in 1/16 of a inch. I = in gauge, else number of 16ths out of gauge.         
	FinalConditionGauge                                                                          *string  `json:"FinalConditionGauge,omitempty"`
	// Condition of inner tooth rows (inner 2/3 of bit) (0-8).                                            
	FinalConditionInner                                                                          *float64 `json:"FinalConditionInner,omitempty"`
	// Row and cone numbers for items which need location information (e.g. Cracked Cone, Lost            
	// Cone etc).                                                                                         
	FinalConditionLocation                                                                       *string  `json:"FinalConditionLocation,omitempty"`
	// Other comments on bit condition from IADC list (BitDullCode in Standard LISTS).                    
	FinalConditionOther                                                                          *string  `json:"FinalConditionOther,omitempty"`
	// Condition of outer tooth rows (outer 1/3 of bit) (0-8).                                            
	FinalConditionOuter                                                                          *float64 `json:"FinalConditionOuter,omitempty"`
	// Reason bit was pulled from IADC codes.                                                             
	FinalConditionReason                                                                         *string  `json:"FinalConditionReason,omitempty"`
	// IADC bit code.                                                                                     
	IADCCode                                                                                     *string  `json:"IADCCode,omitempty"`
	// Condition of bit bearings (integer 0-8 or E, F, N or X)).                                          
	InitialConditionBearing                                                                      *string  `json:"InitialConditionBearing,omitempty"`
	// Overall dull condition from IADC bit wear 2 character codes.                                       
	InitialConditionDull                                                                         *string  `json:"InitialConditionDull,omitempty"`
	// Condition of bit gauge in 1/16 of an inch. I = in gauge, else number of 16ths out of               
	// gauge.                                                                                             
	InitialConditionGauge                                                                        *string  `json:"InitialConditionGauge,omitempty"`
	// Condition of inner tooth rows (inner 2/3 of bit) (0-8).                                            
	InitialConditionInner                                                                        *float64 `json:"InitialConditionInner,omitempty"`
	// Row and cone numbers for items which need location information (e.g. Cracked Cone, Lost            
	// Cone etc).                                                                                         
	InitialConditionLocation                                                                     *string  `json:"InitialConditionLocation,omitempty"`
	// Other comments on bit condition from IADC list (BitDullCode in standard list).                     
	InitialConditionOther                                                                        *string  `json:"InitialConditionOther,omitempty"`
	// Condition of outer tooth rows (outer 1/3 of bit) (0-8).                                            
	InitialConditionOuter                                                                        *float64 `json:"InitialConditionOuter,omitempty"`
	// Reason bit was pulled from IADC codes.                                                             
	InitialConditionReason                                                                       *string  `json:"InitialConditionReason,omitempty"`
	// Manufacturer / supplier of the item.                                                               
	Manufacturer                                                                                 *string  `json:"Manufacturer,omitempty"`
	// The manufacturers code for the bit.                                                                
	MfgCode                                                                                      *string  `json:"MfgCode,omitempty"`
	// Bit number and rerun number e.g. "4.1" for the first rerun of bit 4.                               
	NumBit                                                                                       string   `json:"NumBit"`
	// Type of bit.                                                                                       
	TypeBit                                                                                      *string  `json:"TypeBit,omitempty"`
}
