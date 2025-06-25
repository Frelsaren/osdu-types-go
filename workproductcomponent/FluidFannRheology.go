package workproductcomponent

// Fluid Fann Viscometer Rheology
type FluidFannRheology struct {
	// 100 RPM Dial Reading                                               
	Viscometer100RPMDialReading                                  *float64 `json:"Viscometer100RPMDialReading,omitempty"`
	// 200 RPM Dial Reading                                               
	Viscometer200RPMDialReading                                  *float64 `json:"Viscometer200RPMDialReading,omitempty"`
	// 300 RPM Dial Reading                                               
	Viscometer300RPMDialReading                                  *float64 `json:"Viscometer300RPMDialReading,omitempty"`
	// 30 RPM Dial Reading                                                
	Viscometer30RPMDialReading                                   *float64 `json:"Viscometer30RPMDialReading,omitempty"`
	// 3 RPM Dial Reading                                                 
	Viscometer3RPMDialReading                                    *float64 `json:"Viscometer3RPMDialReading,omitempty"`
	// 600 RPM Dial Reading                                               
	Viscometer600RPMDialReading                                  *float64 `json:"Viscometer600RPMDialReading,omitempty"`
	// 60 RPM Dial Reading                                                
	Viscometer60RPMDialReading                                   *float64 `json:"Viscometer60RPMDialReading,omitempty"`
	// 6 RPM Dial Reading                                                 
	Viscometer6RPMDialReading                                    *float64 `json:"Viscometer6RPMDialReading,omitempty"`
	// Viscometer index or sequence number. Used to retain order.         
	ViscometerIndex                                              *int64   `json:"ViscometerIndex,omitempty"`
	// Viscometer Pressure                                                
	ViscometerPressure                                           *float64 `json:"ViscometerPressure,omitempty"`
	// Viscometer Temperature                                             
	ViscometerTemperature                                        *float64 `json:"ViscometerTemperature,omitempty"`
	// Viscometer Temperature Direction                                   
	ViscometerTemperatureDirection                               *string  `json:"ViscometerTemperatureDirection,omitempty"`
}
