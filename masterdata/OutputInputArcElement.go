package masterdata

// An 'arc' establishing a link between a single pair of ActivityTemplate input and output
// parameters.
type OutputInputArcElement struct {
	// The relationship to the consuming ActivityTemplate.                        
	ConsumingActivityTemplateID                                            string `json:"ConsumingActivityTemplateID"`
	// The Title of the consuming ActivityTemplate Parameter array element.       
	ConsumingParameterTitle                                                string `json:"ConsumingParameterTitle"`
	// The relationship to the producing ActivityTemplate.                        
	ProducingActivityTemplateID                                            string `json:"ProducingActivityTemplateID"`
	// The Title of the producing ActivityTemplate Parameter array element.       
	ProducingParameterTitle                                                string `json:"ProducingParameterTitle"`
}
