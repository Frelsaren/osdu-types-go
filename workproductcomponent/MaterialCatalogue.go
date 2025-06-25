package workproductcomponent

// Object that defines the additives and proppants on location and used in the stimulation
// job.
type MaterialCatalogue struct {
	// Define the additives on location and used in the stimulation job.                
	Additives                                                           []FluidAdditive `json:"Additives,omitempty"`
	// Define the proppants on location and used in the stimulation job.                
	ProppantAgents                                                      []ProppantAgent `json:"ProppantAgents,omitempty"`
}
