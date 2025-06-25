package workproductcomponent

// An injection test, plotted pressure against injection rate, where a curve deflection and
// change of slope indicates the fracture breakdown pressure.
type StimStepTest struct {
	// The pressure necessary to extend the fracture once initiated. The fracture extension                       
	// pressure may rise slightly with increasing fracture length andor height because of                         
	// friction pressure drop down the length of the fracture.                                                    
	FractureExtensionPressure                                                              *float64               `json:"FractureExtensionPressure,omitempty"`
	// A pressure and fluid flow rate data set.                                                                   
	PressureMeasurement                                                                    []StimPressureFlowRate `json:"PressureMeasurement,omitempty"`
}
