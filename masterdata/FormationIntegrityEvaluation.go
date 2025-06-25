package masterdata

// A test performed to determine the integrity of a formation
type FormationIntegrityEvaluation struct {
	// Comments associated with the formation integrity test                                             
	Comments                                                                                    *string  `json:"Comments,omitempty"`
	// The designed maximum gas volume for a given degree of underbalance which the circulation          
	// can be performed without exceeding the weakest formation in the wellbore                          
	DesignKickTolerance                                                                         *float64 `json:"DesignKickTolerance,omitempty"`
	// The depth of the interval drilled before the integrity test. Depth relative to Planned            
	// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds           
	// the depth reference in data.VerticalMeasurement.                                                  
	DrilledIntervalBeforeTest                                                                   *float64 `json:"DrilledIntervalBeforeTest,omitempty"`
	// The total amount of pressure exerted at a true vertical depth which is denoted in the mud         
	// density.                                                                                          
	EquivalentMudWeightEstimated                                                                *float64 `json:"EquivalentMudWeightEstimated,omitempty"`
	// The name of the formation on which the integrity test has been run.                               
	FormationName                                                                               *string  `json:"FormationName,omitempty"`
	// Indicates if a leak off has occurred during the test                                              
	HasLeakOffOccurred                                                                          *bool    `json:"HasLeakOffOccurred,omitempty"`
	// The pressure observed from the LeakOff                                                            
	LeakOffPressure                                                                             *float64 `json:"LeakOffPressure,omitempty"`
	// The maximum pressure observed at the surface during the test.                                     
	MaxTestPressureAtSurface                                                                    *float64 `json:"MaxTestPressureAtSurface,omitempty"`
	// The measured depth of the Casing Shoe during this evaluation. Depth relative to Planned           
	// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds           
	// the depth reference in data.VerticalMeasurement.                                                  
	MeasuredDepthCasingShoe                                                                     *float64 `json:"MeasuredDepthCasingShoe,omitempty"`
	// The type of pressure test used in the evaluation of the formation                                 
	PressureTestType                                                                            *string  `json:"PressureTestType,omitempty"`
	// The density of the fluid measured during the evaluation test                                      
	TestFluidDensity                                                                            *float64 `json:"TestFluidDensity,omitempty"`
	// The pumping rate used during the integrity test.                                                  
	TestingPumpingRate                                                                          *float64 `json:"TestingPumpingRate,omitempty"`
	// The pressure observed at the surface during the test.                                             
	TestPressureAtSurface                                                                       *float64 `json:"TestPressureAtSurface,omitempty"`
	// The true vertical depth of the Casing Shoe during this evaluation. Depth relative to              
	// Planned wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which         
	// holds the depth reference in data.VerticalMeasurement.                                            
	TrueVerticalDepthCasingShoe                                                                 *float64 `json:"TrueVerticalDepthCasingShoe,omitempty"`
	// Amount of fluid pumped during the test                                                            
	VolumePumped                                                                                *float64 `json:"VolumePumped,omitempty"`
}
