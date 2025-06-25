package masterdata

// Context about the decision quality 'appropriate frame' element.
//
// The context for a decision quality element according to
// https://www.decisionprofessionals.com, i.e., one of AppropriateFrame, DoableAlternatives,
// InformationReliability, TradeOffAnalysis, ReasoningCorrectness, CommitmentToAction.
//
// Context about the decision quality 'commitment to action' element.
//
// Context about the decision quality 'meaningful, reliable information' element.
//
// Context about the decision quality 'logically correct reasoning' element.
//
// Context about the decision quality 'clear values and trade-off' element.
type AbstractDecisionQualityElement struct {
	// The assessed decision quality (Sufficient, within acceptable range and residual risk,                     
	// insufficient)                                                                                             
	AssessmentID                                                                                *string          `json:"AssessmentID,omitempty"`
	// A simple text providing the necessary evidence for this aspect of decision quality.                       
	EvidenceAsText                                                                              *string          `json:"EvidenceAsText,omitempty"`
	// The related PersistedCollection, which keeps the data context for this aspect of decision                 
	// quality.                                                                                                  
	EvidenceDataCollectionID                                                                    *string          `json:"EvidenceDataCollectionID,omitempty"`
	// The related Document holding the evidence for this aspect of decision quality.                            
	EvidenceDocumentID                                                                          *string          `json:"EvidenceDocumentID,omitempty"`
	// The array of remarks.                                                                                     
	Remarks                                                                                     []AbstractRemark `json:"Remarks,omitempty"`
}
