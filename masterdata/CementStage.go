package masterdata

// A single cement job.
type CementStage struct {
	// Length of the cement column                                                                              
	AnnularLength                                                                               *float64        `json:"AnnularLength,omitempty"`
	// Measured depth of base of cement. Depth relative to Planned wellbore ZDP. Navigate via                   
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                      
	// data.VerticalMeasurement.                                                                                
	CementBaseMeasuredDepth                                                                     *float64        `json:"CementBaseMeasuredDepth,omitempty"`
	// Displaced Mud, washes and spacers, cements, displacement mud.                                            
	CementingFluid                                                                              *CementingFluid `json:"CementingFluid,omitempty"`
	// Measured depth of CoilTubing (multi-stage cement job. Depth relative to Planned wellbore                 
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                 
	// reference in data.VerticalMeasurement.                                                                   
	CoilTubingMeasuredDepth                                                                     *float64        `json:"CoilTubingMeasuredDepth,omitempty"`
	// Density of displacement fluid.                                                                           
	DensDisplaceFluid                                                                           *float64        `json:"DensDisplaceFluid,omitempty"`
	// Excess volume.                                                                                           
	ExcessVolume                                                                                *float64        `json:"Excess Volume,omitempty"`
	// Average displacement rate.                                                                               
	FlowrateDisplaceAV                                                                          *float64        `json:"FlowrateDisplaceAv,omitempty"`
	// Maximum displacement rate.                                                                               
	FlowrateDisplaceMX                                                                          *float64        `json:"FlowrateDisplaceMx,omitempty"`
	// Final displacement pump rate.                                                                            
	FlowrateEnd                                                                                 *float64        `json:"FlowrateEnd,omitempty"`
	// Rate mud circulated during stage.                                                                        
	FlowrateMudCirc                                                                             *float64        `json:"FlowrateMudCirc,omitempty"`
	// Displacement fluid name.                                                                                 
	FluidDisplace                                                                               *string         `json:"FluidDisplace,omitempty"`
	// Gels-10Min (in hole at start of job).                                                                    
	Gel10Min                                                                                    *float64        `json:"Gel10Min,omitempty"`
	// Gels-10Sec (in hole at start of job).                                                                    
	Gel10SEC                                                                                    *float64        `json:"Gel10Sec,omitempty"`
	// Estimated volume inside casing.                                                                          
	InsideCasingVolume                                                                          *float64        `json:"InsideCasingVolume,omitempty"`
	// Measured depth at top of interval. Depth relative to Planned wellbore ZDP. Navigate via                  
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                      
	// data.VerticalMeasurement.                                                                                
	IntervalTopMeasuredDepth                                                                    *float64        `json:"IntervalTopMeasuredDepth,omitempty"`
	// Estimated volume outside casing for this stage placement.                                                
	OutsideCasingVolume                                                                         *float64        `json:"OutsideCasingVolume,omitempty"`
	// Plastic viscosity (in hole at start of job).                                                             
	PlasticViscoMud                                                                             *float64        `json:"PlasticViscoMud,omitempty"`
	// Mud circulation pressure.                                                                                
	PresMudCirc                                                                                 *float64        `json:"PresMudCirc,omitempty"`
	// Planned displacement pressure.                                                                           
	PressureDisplace                                                                            *float64        `json:"PressureDisplace,omitempty"`
	// Squeeze objective.                                                                                       
	SqueezeObjective                                                                            *string         `json:"SqueezeObjective,omitempty"`
	// Stage number.                                                                                            
	StageNumber                                                                                 float64         `json:"StageNumber"`
	// Stage type.                                                                                              
	StageType                                                                                   string          `json:"StageType"`
	// Measured depth of string (multi-stage cement job). Depth relative to Planned wellbore                    
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                 
	// reference in data.VerticalMeasurement.                                                                   
	StringMeasuredDepth                                                                         *float64        `json:"StringMeasuredDepth,omitempty"`
	// Measured depth of tool (multi-stage cement job. Depth relative to Planned wellbore ZDP.                  
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                      
	// reference in data.VerticalMeasurement.                                                                   
	ToolMeasuredDepth                                                                           *float64        `json:"ToolMeasuredDepth,omitempty"`
	// Type of mud in hole.                                                                                     
	TypeOriginalMud                                                                             *string         `json:"TypeOriginalMud,omitempty"`
	// Marsh funnel viscosity measured in a conical-shaped funnel, fitted with a small-bore tube                
	// on the bottom end through which mud flows under a gravity head. It is a time measurement,                
	// i.e., the time it takes for a given volume of fluid to flow through the funnel.                          
	ViscoFunnelMud                                                                              *float64        `json:"ViscoFunnelMud,omitempty"`
	// Volume of displacement fluid.                                                                            
	VolDisplaceFluid                                                                            *float64        `json:"VolDisplaceFluid,omitempty"`
	// Mud density.                                                                                             
	WtMud                                                                                       *float64        `json:"WtMud,omitempty"`
	// Yield point (in hole at start of job).                                                                   
	YieldPointMud                                                                               *float64        `json:"YieldPointMud,omitempty"`
}
