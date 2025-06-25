package workproductcomponent

// Information on the condition of the Drill Bit as recorded during an operations report
type BitRecord struct {
	// BHA Run Identifier (includes ID of TubularAssembly which parents the Bit/Under-reamer            
	// Component)                                                                                       
	BHARunID                                                                                   *string  `json:"BHARunID,omitempty"`
	// N = new, U = used.                                                                               
	BitClass                                                                                   *string  `json:"BitClass,omitempty"`
	// Any other Bit Dull comments/description, older T-B-G code or full formatted dull grade           
	// e.g.  0-2-WT-N/S-X-I-LN-TD                                                                       
	BitDullComments                                                                            *string  `json:"BitDullComments,omitempty"`
	// Bit cost in local currency.                                                                      
	Cost                                                                                       *float64 `json:"Cost,omitempty"`
	// The name of the local currency                                                                   
	CostCurrency                                                                               *string  `json:"CostCurrency,omitempty"`
	// Diameter of drilled hole.                                                                        
	DiameterBit                                                                                *float64 `json:"DiameterBit,omitempty"`
	// Minimum hole or tubing which bit will pass through (for bi-center bits).                         
	DiameterPassThrough                                                                        *float64 `json:"DiameterPassThrough,omitempty"`
	// Diameter of pilot bit (for bi-center bits).                                                      
	DiameterPilot                                                                              *float64 `json:"DiameterPilot,omitempty"`
	// Bit drive type (Motor, rotary table etc).                                                        
	DriveType                                                                                  *string  `json:"DriveType,omitempty"`
	// Pulled bit IADC Dull Grade (B) Condition of bit Bearings (integer 0-8 or E, F, N or X).          
	FinalConditionBearing                                                                      *string  `json:"FinalConditionBearing,omitempty"`
	// Pulled bit IADC Dull Grade (D) Overall Dull condition from IADC bit wear 2 character             
	// codes.                                                                                           
	FinalConditionDull                                                                         *string  `json:"FinalConditionDull,omitempty"`
	// Pulled bit IADC Dull Grade (G) Condition of bit Gauge in 1/16 of a inch. I = in gauge,           
	// else number of 16ths out of gauge.                                                               
	FinalConditionGauge                                                                        *string  `json:"FinalConditionGauge,omitempty"`
	// Pulled bit IADC Dull Grade (I) Condition of Inner tooth rows (inner 2/3 of bit) (0-8).           
	FinalConditionInner                                                                        *float64 `json:"FinalConditionInner,omitempty"`
	// Pulled bit IADC Dull Grade (L) Location Row and cone numbers where wear located (e.g.            
	// Cracked Cone, Lost Cone etc).                                                                    
	FinalConditionLocation                                                                     *string  `json:"FinalConditionLocation,omitempty"`
	// Pulled bit IADC Dull Grade (O) Other comments on bit condition.                                  
	FinalConditionOther                                                                        *string  `json:"FinalConditionOther,omitempty"`
	// Pulled bit IADC Dull Grade (O) Condition of Outer tooth rows (outer 1/3 of bit) (0-8).           
	FinalConditionOuter                                                                        *float64 `json:"FinalConditionOuter,omitempty"`
	// Pulled bit  IADC Dull Grade (R) Reason bit was pulled.                                           
	FinalConditionReason                                                                       *string  `json:"FinalConditionReason,omitempty"`
	// IADC bit code.                                                                                   
	IADCCode                                                                                   *string  `json:"IADCCode,omitempty"`
	// As run bit IADC Dull Grade (B) Condition of bit Bearings (integer 0-8 or E, F, N or X)).         
	InitialConditionBearing                                                                    *string  `json:"InitialConditionBearing,omitempty"`
	// As run bit IADC Dull Grade (D) Overall dull condition from IADC bit wear 2 character             
	// codes.                                                                                           
	InitialConditionDull                                                                       *string  `json:"InitialConditionDull,omitempty"`
	// As run bit IADC Dull Grade (G) Condition of bit Gauge in 1/16 of an inch. I = in gauge,          
	// else number of 16ths out of gauge.                                                               
	InitialConditionGauge                                                                      *string  `json:"InitialConditionGauge,omitempty"`
	// As run bit IADC Dull Grade (I) Condition of Inner tooth rows (inner 2/3 of bit) (0-8).           
	InitialConditionInner                                                                      *float64 `json:"InitialConditionInner,omitempty"`
	// As run bit IADC Dull Grade (L) Location Row and cone numbers where wear located (e.g.            
	// Cracked Cone, Lost Cone etc).                                                                    
	InitialConditionLocation                                                                   *string  `json:"InitialConditionLocation,omitempty"`
	// As run bit IADC Dull Grade (O) Other comments on bit condition from IADC list                    
	// (BitDullCode in standard list).                                                                  
	InitialConditionOther                                                                      *string  `json:"InitialConditionOther,omitempty"`
	// As run bit IADC Dull Grade (O) Condition of Outer tooth rows (outer 1/3 of bit) (0-8).           
	InitialConditionOuter                                                                      *float64 `json:"InitialConditionOuter,omitempty"`
	// As run bit IADC Dull Grade (R) Reason bit was pulled                                             
	InitialConditionReason                                                                     *string  `json:"InitialConditionReason,omitempty"`
	// Manufacturer / supplier of the item.                                                             
	Manufacturer                                                                               *string  `json:"Manufacturer,omitempty"`
	// The manufacturer's code for the bit.                                                             
	MfgCode                                                                                    *string  `json:"MfgCode,omitempty"`
	// Bit number and rerun number e.g. "4.1" for the first rerun of bit 4.                             
	NumBit                                                                                     string   `json:"NumBit"`
	// Identifier of Bit Tubular Component                                                              
	TubularComponentID                                                                         *string  `json:"TubularComponentID,omitempty"`
	// Type of bit.                                                                                     
	TypeBit                                                                                    *string  `json:"TypeBit,omitempty"`
}
