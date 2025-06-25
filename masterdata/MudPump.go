package masterdata

// Pumps present on the rig
type MudPump struct {
	// Maximum power defined by the manufacturer.                   
	MaxPower                                               *float64 `json:"MaxPower,omitempty"`
	// Maximum required/delivered flowrate from/by the pump         
	MaxPumpFlowrate                                        *float64 `json:"MaxPumpFlowrate,omitempty"`
	// Maximum required/delivered pressure from/by the pump         
	MaxPumpPressure                                        *float64 `json:"MaxPumpPressure,omitempty"`
	// The name of the mud pump (model)                             
	Name                                                   *string  `json:"Name,omitempty"`
	// Number of pump of that type on the rig.                      
	Number                                                 *int64   `json:"Number,omitempty"`
	// Maximum power authorized by the drilling contractor          
	WorkingPower                                           *float64 `json:"WorkingPower,omitempty"`
}
