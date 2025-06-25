package workproductcomponent

// ISO13503_2 crush test data for this instance of ISO13503_2 data for the proppant agent.
type ISO135032CrushTestData struct {
	// Mass percentage of fines after being exposed to stress.         
	Fines                                                     *float64 `json:"Fines,omitempty"`
	// Stress measured at a point during a crush test.                 
	Stress                                                    *float64 `json:"Stress,omitempty"`
}
