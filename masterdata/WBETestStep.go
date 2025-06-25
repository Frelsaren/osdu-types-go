package masterdata

// Description of an individual step of the overall Well Barrier Element Test
type WBETestStep struct {
	// A list of pressures and the actual volume bled back at the given pressure.                                              
	ActualVolumesBledBack                                                                     []ActualVolumesBledBackElement   `json:"ActualVolumesBledBack,omitempty"`
	// A list of pressures and the actual volume pumped at the given pressure.                                                 
	ActualVolumesPumped                                                                       []ActualVolumesPumpedElement     `json:"ActualVolumesPumped,omitempty"`
	// Text string for noting the components tested                                                                            
	ComponentsTested                                                                          *string                          `json:"ComponentsTested,omitempty"`
	// A list of systems and their volumes.                                                                                    
	ComponentVolumesTestSystem                                                                []ComponentVolumesTestSystem     `json:"ComponentVolumesTestSystem,omitempty"`
	// The name of a pass/fail or acceptance criteria template that applies for the step                                       
	CriteriaTemplateName                                                                      *string                          `json:"CriteriaTemplateName,omitempty"`
	// The expected pressure change due to thermal effects                                                                     
	ExpectedPressureChangeDueToThermalEffects                                                 float64                          `json:"ExpectedPressureChangeDueToThermalEffects"`
	// Comments on the expected pressure change due to thermal effects. Can be used to provide                                 
	// an explanation as to why a value was not entered (e.g. expansion not expected to impact                                 
	// test results)                                                                                                           
	ExpectedPressureChangeDueToThermalEffectsComments                                         *string                          `json:"ExpectedPressureChangeDueToThermalEffectsComments,omitempty"`
	// A list of pressures and the expected volume bled back at the given pressure.                                            
	ExpectedVolumesBledBack                                                                   []ExpectedVolumesBledBackElement `json:"ExpectedVolumesBledBack,omitempty"`
	// A list of pressures and the expected volume pumped at the given pressure.                                               
	ExpectedVolumesPumped                                                                     []ExpectedVolumesPumpedElement   `json:"ExpectedVolumesPumped,omitempty"`
	// Link to the object containing the high pressure criteria.                                                               
	HighPressureCriteria                                                                      []HighPressureCriterionElement   `json:"HighPressureCriteria"`
	// Link to the object containing the low pressure criteria .                                                               
	LowPressureCriteria                                                                       []LowPressureCriterionElement    `json:"LowPressureCriteria,omitempty"`
	// For use when executing simultaneous tests using two or more different sets of criteria                                  
	ParallelWBETestStepIDs                                                                    []string                         `json:"ParallelWBETestStepIDs,omitempty"`
	// The ID of a schematic document that describes the pressure path / line-up                                               
	SchematicID                                                                               *string                          `json:"SchematicID,omitempty"`
	// A description of the individual test step.                                                                              
	StepDescription                                                                           *string                          `json:"StepDescription,omitempty"`
	// Additional notes on the test step                                                                                       
	StepNotes                                                                                 *string                          `json:"StepNotes,omitempty"`
	// An integer that describes the order in which the pressure test steps will be performed                                  
	StepNumber                                                                                int64                            `json:"StepNumber"`
	// A list of test fluids and their properties.                                                                             
	TestFluids                                                                                []TestFluid                      `json:"TestFluids,omitempty"`
	// The cumulative total volume of the systems lined up for the test                                                        
	TotalVolumeTestSystem                                                                     float64                          `json:"TotalVolumeTestSystem"`
}
