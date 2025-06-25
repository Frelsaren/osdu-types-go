package masterdata

// The sensor, meter or tool equipment used to acquire the gravity and magnetic geophysical
// measurements
type Equipment struct {
	// The manufacturer of a sensor, meter or tool used to acquire gravity or magnetics                         
	// geophysical data.                                                                                        
	Manufacturer                                                                               *string          `json:"Manufacturer,omitempty"`
	// The model name or number, typically given by the manufacturer of a sensor, meter or tool                 
	// used to acquire gravity or magnetics geophysical data.                                                   
	Model                                                                                      *string          `json:"Model,omitempty"`
	// A name given to a sensor, meter or tool used to acquire gravity or magnetics geophysical                 
	// data. This could be a common or colloquial name.                                                         
	// Where possible, for better search results, the manufacturer, model and serial number                     
	// should be completed also.                                                                                
	Name                                                                                       *string          `json:"Name,omitempty"`
	// An array of remarks about the equipment used to acquire the measurements, for example,                   
	// the orientation of equipment. This utilises the AbstractRemark fragment.                                 
	Remarks                                                                                    []AbstractRemark `json:"Remarks,omitempty"`
	// The serial number, typically given by the manufacturer, of a sensor, meter or tool used                  
	// to acquire gravity or magnetics geophysical data.                                                        
	SerialNumber                                                                               *string          `json:"SerialNumber,omitempty"`
}
