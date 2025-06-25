package masterdata

// Set of (Time / Rate / Back Pressure).
type CementPumpSchedule struct {
	// Back pressure applied during pumping stage.          
	BackPressure                                   *float64 `json:"BackPressure,omitempty"`
	// Comments and Remarks.                                
	Comments                                       *string  `json:"Comments,omitempty"`
	// Rate fluid is pumped. 0 means it is a pause.         
	PumpRate                                       *float64 `json:"PumpRate,omitempty"`
	// Volume pumped = eTimPump * ratePump.                 
	PumpVolume                                     *float64 `json:"PumpVolume,omitempty"`
}
