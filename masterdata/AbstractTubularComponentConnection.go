package masterdata

// Tubular Connection specific properties. Based on WITSML Tubular  Component 'Connection'
// ComplexType.
type AbstractTubularComponentConnection struct {
	// Critical Cross Section Area (for bending stiffness ratio).         
	CriticalCrossSectionArea                                     *float64 `json:"CriticalCrossSectionArea,omitempty"`
	// Inside diameter (ID) of the connection.                            
	CXNInnerDiameter                                             *float64 `json:"CXNInnerDiameter,omitempty"`
	// Outside diameter (OD) of the body of the connection.               
	CXNOuterDiameter                                             *float64 `json:"CXNOuterDiameter,omitempty"`
	// Internal length of the connection.                                 
	InsideConnectionLength                                       *float64 `json:"InsideConnectionLength,omitempty"`
	// Leak Pressure                                                      
	LeakPressure                                                 *float64 `json:"LeakPressure,omitempty"`
	// Make Up Torque                                                     
	MakeUpTorque                                                 *float64 `json:"MakeUpTorque,omitempty"`
	// External length of the connection.                                 
	OutsideConnectionLength                                      *float64 `json:"OutsideConnectionLength,omitempty"`
	// Position of connection                                             
	PositionID                                                   *string  `json:"PositionID,omitempty"`
	// Thread size/diameter.                                              
	ThreadSize                                                   *float64 `json:"ThreadSize,omitempty"`
	// Thread/Connection Type                                             
	ThreadTypeID                                                 *string  `json:"ThreadTypeID,omitempty"`
	// Tensile yield stress of the connection                             
	YieldStress                                                  *float64 `json:"YieldStress,omitempty"`
	// Torque limit for the connection                                    
	YieldTorque                                                  *float64 `json:"YieldTorque,omitempty"`
}
