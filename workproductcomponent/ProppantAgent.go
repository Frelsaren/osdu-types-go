package workproductcomponent

// Defines the properties of a Proppant and it's test values.
type ProppantAgent struct {
	// Laminar flow friction coefficient.                                                                                                 
	FrictionCoefficientLaminar                                                                  *float64                                  `json:"FrictionCoefficientLaminar,omitempty"`
	// Turbulent flow friction coefficient.                                                                                               
	FrictionCoefficientTurbulent                                                                *float64                                  `json:"FrictionCoefficientTurbulent,omitempty"`
	ISO135032Properties                                                                         []StimISO135032Properties                 `json:"ISO13503_2Properties,omitempty"`
	ISO135035Point                                                                              []ISO135035PropertiesForThisProppantAgent `json:"ISO13503_5Point,omitempty"`
	// Characterizes how easily radiation passes through a material. This can be used to compute                                          
	// the concentration of proppant in a slurry using a densitometer.                                                                    
	MassAbsorptionCoefficient                                                                   *float64                                  `json:"MassAbsorptionCoefficient,omitempty"`
	// Integer high value of sieve mesh size: for 4070 sand, this value is 70. For a 70 sand                                              
	// this would be 70.                                                                                                                  
	MeshSizeHigh                                                                                *int64                                    `json:"MeshSizeHigh,omitempty"`
	// Integer low value of sieve mesh size: for 4070 sand, this value is 40.                                                             
	MeshSizeLow                                                                                 *int64                                    `json:"MeshSizeLow,omitempty"`
	// The name of the proppant.                                                                                                          
	ProppantName                                                                                *string                                   `json:"ProppantName,omitempty"`
	// Proppant type or function.                                                                                                         
	ProppantTypeID                                                                              *string                                   `json:"ProppantTypeID,omitempty"`
	// General remarks about this proppant agent.                                                                                         
	Remarks                                                                                     *string                                   `json:"Remarks,omitempty"`
	// Identifier of service company organisation providing this proppant agent for the                                                   
	// stimulation job.                                                                                                                   
	SupplierID                                                                                  *string                                   `json:"SupplierID,omitempty"`
	// The name of the proppant supplier.                                                                                                 
	SupplierName                                                                                *string                                   `json:"SupplierName,omitempty"`
	// Supplier's product number/code.                                                                                                    
	SupplierProductCode                                                                         *string                                   `json:"SupplierProductCode,omitempty"`
	// The unconfined compressive strength of the proppant.                                                                               
	UnconfinedCompressiveStrength                                                               *float64                                  `json:"UnconfinedCompressiveStrength,omitempty"`
}
