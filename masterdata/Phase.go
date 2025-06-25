package masterdata

// Description of an individual phase that compose the WellActivityProgram. A phase
// describes the key milestones and is usually based on the major sections of the well or
// non-well-related work.
// (Such as "Acces Well", "Suspend ", "Wellhead Removal "...)
type Phase struct {
	// Reference to the different BHA runs                                                                
	BHARunIDs                                                                                    []string `json:"BHARunIDs"`
	// A reference to the object that holds the evaluation information about the drilling                 
	// program for this phase.                                                                            
	EvaluationPlanID                                                                             *string  `json:"EvaluationPlanID,omitempty"`
	// Volume of a kick that will not cause a pressure exceeding the formation fracture pressure.         
	KickTolerance                                                                                *float64 `json:"KickTolerance,omitempty"`
	// An absolute upper limit for the pressure in the annulus of an oil and gas well as                  
	// measured at the wellhead. It's a number with pressure units                                        
	Maasp                                                                                        *float64 `json:"MAASP,omitempty"`
	// Free text box to describe the objectives from the phase.                                           
	Objective                                                                                    *string  `json:"Objective,omitempty"`
	// Reference to the objects that hold the cementing design information about the drilling             
	// program for this phase                                                                             
	PlannedCementJobIDs                                                                          []string `json:"PlannedCementJobIDs,omitempty"`
	// Reference to the objects that hold the expected lithology designs information about the            
	// drilling program for this phase                                                                    
	PlannedLithologyIDs                                                                          []string `json:"PlannedLithologyIDs,omitempty"`
	// Reference to the objects that holds the information about the risks that apply to this             
	// drilling program (may be too high level for this)                                                  
	RiskIDs                                                                                      []string `json:"RiskIDs,omitempty"`
	// The order that these phases will be executed (is this needed?)                                     
	SequenceNo                                                                                   *float64 `json:"SequenceNo,omitempty"`
	// A reference to the object that holds the information about the type of phase that is               
	// described (such as Workover, Completion…)                                                          
	TypeID                                                                                       string   `json:"TypeID"`
	// A reference to the object that holds the information about the activity plan for this              
	// Well Activity program                                                                              
	WellActivityPlanID                                                                           string   `json:"WellActivityPlanID"`
	// Identifier of the description from the Well Barrier Element Test to be run                         
	WellBarrierElementTestID                                                                     *string  `json:"WellBarrierElementTestID,omitempty"`
	// A reference to the object that describes the physical structure of a borehole                      
	WellboreArchitectureID                                                                       *string  `json:"WellboreArchitectureID,omitempty"`
	// The wellbore to which this drilling program phase refers                                           
	WellboreID                                                                                   *string  `json:"WellboreID,omitempty"`
}
