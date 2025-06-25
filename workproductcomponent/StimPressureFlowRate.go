package workproductcomponent

// An array of bottom hole flow rates and pressures for this stim step test.
type StimPressureFlowRate struct {
	// The bottomhole flow rate of the fluid.         
	BHFlowRate                               *float64 `json:"BHFlowRate,omitempty"`
	// The pressure of the step test.                 
	Pressure                                 *float64 `json:"Pressure,omitempty"`
}
