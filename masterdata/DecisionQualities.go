package masterdata

// The 6-component decision quality object.
type DecisionQualities struct {
	// Context about the decision quality 'appropriate frame' element.                                               
	AppropriateFrame                                                                 *AbstractDecisionQualityElement `json:"AppropriateFrame,omitempty"`
	// Context about the decision quality 'commitment to action' element.                                            
	CommitmentToAction                                                               *AbstractDecisionQualityElement `json:"CommitmentToAction,omitempty"`
	// Context about the decision quality 'creative, doable alternatives' elements.                                  
	DoableAlternatives                                                               []OneAlternative                `json:"DoableAlternatives,omitempty"`
	// Context about the decision quality 'meaningful, reliable information' element.                                
	InformationReliability                                                           *AbstractDecisionQualityElement `json:"InformationReliability,omitempty"`
	// Context about the decision quality 'logically correct reasoning' element.                                     
	ReasoningCorrectness                                                             *AbstractDecisionQualityElement `json:"ReasoningCorrectness,omitempty"`
	// Context about the decision quality 'clear values and trade-off' element.                                      
	TradeOffAnalysis                                                                 *AbstractDecisionQualityElement `json:"TradeOffAnalysis,omitempty"`
}
