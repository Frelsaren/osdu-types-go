package workproductcomponent

// Volume pumped versus Pressure response curve data. Multiple curves supported, one set for
// each Pressure Data Source
type VolumePressureResponse struct {
	// Choke A Open in fixed % (percent) unit. No unit conversion supported.         
	ChokeAOpenPercent                                                       *float64 `json:"ChokeAOpenPercent,omitempty"`
	// Choke B Open in fixed % (percent) unit. No unit conversion supported.         
	ChokeBOpenPercent                                                       *float64 `json:"ChokeBOpenPercent,omitempty"`
	// Elapsed Time                                                                  
	ElapsedTime                                                             *float64 `json:"ElapsedTime,omitempty"`
	// Flow In Rate                                                                  
	FlowInRate                                                              *float64 `json:"FlowInRate,omitempty"`
	// Flow Out Rate                                                                 
	FlowOutRate                                                             *float64 `json:"FlowOutRate,omitempty"`
	// Pressure Data Source                                                          
	PressureDataSource                                                      *string  `json:"PressureDataSource,omitempty"`
	// Pumped Volume                                                                 
	PumpedVolume                                                            *float64 `json:"PumpedVolume,omitempty"`
	// Pump Rate                                                                     
	PumpRate                                                                *float64 `json:"PumpRate,omitempty"`
	// PWD Elapsed Time                                                              
	PWDElapsedTime                                                          *float64 `json:"PWDElapsedTime,omitempty"`
	// PWD Pressure                                                                  
	PWDPressure                                                             *float64 `json:"PWDPressure,omitempty"`
	// Remark                                                                        
	Remark                                                                  *string  `json:"Remark,omitempty"`
	// Sequence Number                                                               
	SequenceNumber                                                          int64    `json:"SequenceNumber"`
	// Surface Pressure                                                              
	SurfacePressure                                                         *float64 `json:"SurfacePressure,omitempty"`
	// Source for surface pressure measurement                                       
	SurfacePressureSource                                                   *string  `json:"SurfacePressureSource,omitempty"`
}
