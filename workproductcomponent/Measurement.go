package workproductcomponent

// Unitary Measurement Serie Description
type Measurement struct {
	// Identifier of the unitary gauge active for this specific curve measurement                                                    
	ActiveGaugeID                                                                               *string                              `json:"ActiveGaugeID,omitempty"`
	// Describe the active Probe type for this current measurement station                                                           
	ActiveProbeTypeID                                                                           string                               `json:"ActiveProbeTypeID"`
	// The vertical measurement reference for this well logging acquisition activity. This                                           
	// object defines the vertical reference datum for the measured depths.                                                          
	DepthReference                                                                              *AbstractFacilityVerticalMeasurement `json:"DepthReference,omitempty"`
	// Identifier of the detailed test run at this station (such as miniDST, miniFrac,…)                                             
	DetailedTestTypeID                                                                          *string                              `json:"DetailedTestTypeID,omitempty"`
	// Boolean describing if the test serie is run at constant depth (as opposed as at different                                     
	// stations)                                                                                                                     
	IsStationary                                                                                *bool                                `json:"IsStationary,omitempty"`
	// Boolean that indicates the measurement is tubular (meaning proceeded inside the installed                                     
	// tubular) - or annular otherwise.                                                                                              
	IsTubular                                                                                   *bool                                `json:"IsTubular,omitempty"`
	// array of information to identify interval times within the complete measurement                                               
	// recording, significant in themselves - such as PreTests in Formation Pressure Test -                                          
	// BuildUp/Close-in intervals in Build Ups production pressure tests,….                                                          
	MeasurementPeriods                                                                          []MeasurementPeriod                  `json:"MeasurementPeriods,omitempty"`
	// Measured Depth of the station for Stationary Tests - mutually exclusive with                                                  
	// (PressureTestTop/bottom measuredDepth                                                                                         
	PressurePointMeasuredDepth                                                                  *float64                             `json:"PressurePointMeasuredDepth,omitempty"`
	// Measured Depth of the bottom station for Non Stationary Tests- mutually exclusive with                                        
	// (Pressure Point Measured Depth                                                                                                
	PressureTestBottomMeasuredDepth                                                             *float64                             `json:"PressureTestBottomMeasuredDepth,omitempty"`
	// Measured Depth of the top station for Non Stationary Tests - mutually exclusive with                                          
	// (Pressure Point Measured Depth                                                                                                
	PressureTestTopMeasuredDepth                                                                *float64                             `json:"PressureTestTopMeasuredDepth,omitempty"`
	// DEPRECATED - PLEASE USE "MEASUREMENT PERIOD" INSTEAD: array of information to identify                                        
	// pretests measurements within the complete measurement recording.                                                              
	PreTests                                                                                    []PreTest                            `json:"PreTests,omitempty"`
	// Sequential number of the test (or pre-test) which identifies it into the record                                               
	TestNumber                                                                                  *int64                               `json:"TestNumber,omitempty"`
	// Identifier of the reference data record indicating the reliability of the tests. This                                         
	// encapsulates the common "TestSuccess" notion - but allows for more flexibility                                                
	TestReliabilityID                                                                           *string                              `json:"TestReliabilityID,omitempty"`
	// Identifier of the Station within the acquisition job captured in the unitary serie                                            
	WellPressureTestAcquisitionStationIdentifier                                                *int64                               `json:"WellPressureTestAcquisitionStationIdentifier,omitempty"`
}
