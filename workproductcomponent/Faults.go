package workproductcomponent

// Array of Faults that comprise the Fault System
type Faults struct {
	// Maximum stratigraphic heave, the apparent horizontal component of the net-slip.            
	FaultHeaveNumber                                                                     *float64 `json:"FaultHeaveNumber,omitempty"`
	// ID of the Unit of Measure of the FaultHeaveNumber                                          
	FaultHeaveNumberUOM                                                                  *string  `json:"FaultHeaveNumberUOM,omitempty"`
	// The related FaultInterpretation for collaborating with Earth Modeling.                     
	FaultInterpretationID                                                                *string  `json:"FaultInterpretationID,omitempty"`
	// Net (average) Slip                                                                         
	FaultNetSlipNumber                                                                   *float64 `json:"FaultNetSlipNumber,omitempty"`
	// ID of the Unit of Measure of the FaultNetSlipNumber                                        
	FaultNetSlipNumberUOM                                                                *string  `json:"FaultNetSlipNumberUOM,omitempty"`
	// The person or team who interpreted the fault data.                                         
	Interpreter                                                                          *string  `json:"Interpreter,omitempty"`
	// Optional comment                                                                           
	Remark                                                                               *string  `json:"Remark,omitempty"`
	// RepresentationRole for this Fault element if more than one role is in use for this         
	// FaultSystem.                                                                               
	Role                                                                                 *string  `json:"Role,omitempty"`
	// Maximum linear dimension measured along strike of the slip surface                         
	SeismicFaultLength                                                                   *float64 `json:"SeismicFaultLength,omitempty"`
	// ID of the Unit of Measure of the Length of the Fault                                       
	SeismicFaultLengthUOM                                                                *string  `json:"SeismicFaultLengthUOM,omitempty"`
	// Name of an individual fault within a fault system.                                         
	SeismicFaultName                                                                     *string  `json:"SeismicFaultName,omitempty"`
	// Surface Area of the Fault Plane                                                            
	SeismicFaultSurfaceArea                                                              *float64 `json:"SeismicFaultSurfaceArea,omitempty"`
	// ID of the Unit of Measure of the Surface Area of the Fault                                 
	SeismicFaultSurfaceAreaUOM                                                           *string  `json:"SeismicFaultSurfaceAreaUOM,omitempty"`
	// Geological type of fault geometry. E.g. Thrust (thr), Reverse (rev), Normal(norm)          
	SeismicFaultTypeID                                                                   *string  `json:"SeismicFaultTypeID,omitempty"`
	// Method used to pick faults. E.g.Autotracked, Grid, Manual Picked                           
	SeismicPickingTypeID                                                                 *string  `json:"SeismicPickingTypeID,omitempty"`
	// Maximum vertical offset of faulted strata.                                                 
	StratigraphicFaultOffset                                                             *float64 `json:"StratigraphicFaultOffset,omitempty"`
	// ID of the Unit of Measure of the StratigraphicFaultOffset                                  
	StratigraphicFaultOffsetUOM                                                          *string  `json:"StratigraphicFaultOffsetUOM,omitempty"`
	// RepresentationType for this Fault element if more than one type is in use for this         
	// FaultSystem.                                                                               
	Type                                                                                 *string  `json:"Type,omitempty"`
	// Maximum vertical angle of fault                                                            
	VerticalFaultDIPAngle                                                                *float64 `json:"VerticalFaultDipAngle,omitempty"`
	// ID of the Unit of Measure of the Dip angle of the Fault                                    
	VerticalFaultDIPAngleUOM                                                             *string  `json:"VerticalFaultDipAngleUOM,omitempty"`
}
