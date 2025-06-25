package masterdata

// A test performed on a drillstem
type DSTRun struct {
	// Comments made regarding the Drillstem Test                                                        
	Comments                                                                                    *string  `json:"Comments,omitempty"`
	// Perforation Base Depth for the DST to run through. Depth relative to Planned wellbore             
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth          
	// reference in data.VerticalMeasurement.                                                            
	DepthPerforationBase                                                                        *float64 `json:"DepthPerforationBase,omitempty"`
	// Perforation Top Depth for the DST to run through. Depth relative to Planned wellbore ZDP.         
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth               
	// reference in data.VerticalMeasurement.                                                            
	DepthPerforationTop                                                                         *float64 `json:"DepthPerforationTop,omitempty"`
	// The name of the formation in which the test was performed.                                        
	FormationName                                                                               *string  `json:"FormationName,omitempty"`
	// Description of the Hole Section in which the evaluation will be performed                         
	HoleSectionID                                                                               *string  `json:"HoleSectionID,omitempty"`
	// Free text describing the type of the hole the DST is running through (such as "Cased",            
	// "OpenHole",…)                                                                                     
	HoleType                                                                                    *string  `json:"HoleType,omitempty"`
	// The maximum start depth for the drillstem test. Depth relative to Planned wellbore ZDP.           
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth               
	// reference in data.VerticalMeasurement.                                                            
	MaximumDepthTestStart                                                                       *float64 `json:"MaximumDepthTestStart,omitempty"`
	// The minimum start depth for the drillstem test. Depth relative to Planned wellbore ZDP.           
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth               
	// reference in data.VerticalMeasurement.                                                            
	MinimumDepthTestStart                                                                       *float64 `json:"MinimumDepthTestStart,omitempty"`
	// Depth of the Packer while running the DST                                                         
	PackerDepth                                                                                 *float64 `json:"PackerDepth,omitempty"`
	// Sequential number of the perforation that the DST is running through                              
	PerforationNumber                                                                           *float64 `json:"PerforationNumber,omitempty"`
	// Sequential number of the DST                                                                      
	Sequence                                                                                    *float64 `json:"Sequence,omitempty"`
	// Planned test length                                                                               
	TestIntervalLength                                                                          *float64 `json:"TestIntervalLength,omitempty"`
	// Planned test duration                                                                             
	TestTimeDuration                                                                            *float64 `json:"TestTimeDuration,omitempty"`
}
