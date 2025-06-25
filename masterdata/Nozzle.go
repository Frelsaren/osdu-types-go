package masterdata

// number and size of nozzles / jets in a Tubular Component
type Nozzle struct {
	// Inside Diameter of the nozzle                                      
	DiameterInside                                                float64 `json:"DiameterInside"`
	// Total length of the nozzle                                         
	Length                                                        float64 `json:"Length"`
	// Number of Nozzles installed in component                           
	NozzleCount                                                   int64   `json:"NozzleCount"`
	// Unique identifier for this instance of Nozzle                      
	NozzleID                                                      string  `json:"NozzleID"`
	// Describes the Nozzle Type (such as extended; normal; blank)        
	NozzleTypeID                                                  string  `json:"NozzleTypeID"`
	// Nozzle Orientation                                                 
	Orientation                                                   float64 `json:"Orientation"`
}
