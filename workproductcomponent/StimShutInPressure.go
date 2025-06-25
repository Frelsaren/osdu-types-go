package workproductcomponent

// Shut in pressure data for this job stage.
type StimShutInPressure struct {
	// The shut-in pressure.                                                              
	Pressure                                                                     *float64 `json:"Pressure,omitempty"`
	// The cumulative time span after shut-in at which the pressure was measured.         
	TimeAfterShutin                                                              *float64 `json:"TimeAfterShutin,omitempty"`
}
