package masterdata

// A fluid used in the drilling of a wellbore's section
type FluidsInterval struct {
	// Comments and remarks.                                                                                
	Comments                                                                                    *string     `json:"Comments,omitempty"`
	// Provides the overall description of the drilling fluids system.                                      
	FluidsSystem                                                                                FluidSystem `json:"FluidsSystem"`
	// A fixed list of reference values describing the high level type of the drilling fluid                
	FluidTypeID                                                                                 *string     `json:"FluidTypeID,omitempty"`
	// Description of the Hole Section for this Fluids Program                                              
	HoleSectionID                                                                               *string     `json:"HoleSectionID,omitempty"`
	// The identifier of the tubular to be installed.                                                       
	InstalledTubularAssemblyID                                                                  *string     `json:"InstalledTubularAssemblyID,omitempty"`
	// The bottom measured depth of the interval in which the fluid will be used (in many cases,            
	// this Measured Depth will be the same than the one provided by                                        
	// WellboreArchitecture.HoleSection content). Depth relative to Planned wellbore ZDP.                   
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                  
	// reference in data.VerticalMeasurement.                                                               
	IntervalBaseMeasuredDepth                                                                   float64     `json:"IntervalBaseMeasuredDepth"`
	// The name of an interval in which the fluid will be used                                              
	IntervalName                                                                                string      `json:"IntervalName"`
	// The top measured depth of the interval in which the fluid will be used (in many cases,               
	// this Measured Depth will be the same than the one provided by                                        
	// WellboreArchitecture.HoleSection content). Depth relative to Planned wellbore ZDP.                   
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                  
	// reference in data.VerticalMeasurement.                                                               
	IntervalTopMeasuredDepth                                                                    float64     `json:"IntervalTopMeasuredDepth"`
	// A reference number allowing traceability back to the analysis of the fluid in a lab                  
	LabReferenceNumber                                                                          *float64    `json:"LabReferenceNumber,omitempty"`
	// Funnel viscosity in seconds.                                                                         
	ViscosityFunnel                                                                             *float64    `json:"ViscosityFunnel,omitempty"`
}
