package workproductcomponent

// This structure intends to define the reservoir boundary connection in an IJK grid such as
// aquifer connection. It can be reused in various Ijk Grid flow simulation related
// representations which need some boundary connections.
type AbstractIjkGridFlowSimulationBoundaryConnection struct {
	// The faces of the box of the grid which are in connection                
	Face                                                              *string  `json:"Face,omitempty"`
	// The grid which is in connection                                         
	GridID                                                            *string  `json:"GridID,omitempty"`
	// The lower included I index of the box of the grid in connection         
	LowerI                                                            *int64   `json:"LowerI,omitempty"`
	// The lower included J index of the box of the grid in connection         
	LowerJ                                                            *int64   `json:"LowerJ,omitempty"`
	// The lower included K index of the box of the grid in connection         
	LowerK                                                            *int64   `json:"LowerK,omitempty"`
	// The transmissibility multiplier of the connection                       
	TransmissibilityMultiplier                                        *float64 `json:"TransmissibilityMultiplier,omitempty"`
	// The upper included I index of the box of the grid in connection         
	UpperI                                                            *int64   `json:"UpperI,omitempty"`
	// The upper included J index of the box of the grid in connection         
	UpperJ                                                            *int64   `json:"UpperJ,omitempty"`
	// The upper included K index of the box of the grid in connection         
	UpperK                                                            *int64   `json:"UpperK,omitempty"`
}
