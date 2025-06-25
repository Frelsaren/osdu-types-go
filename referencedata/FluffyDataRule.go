package referencedata

// The list of QualityDataRule items that this QualityDataRuleSet consists of.
type FluffyDataRule struct {
	DataRuleID                                                                 *string `json:"DataRuleID,omitempty"`
	// Indicated if the QualityDataRule is required to pass or for information.        
	DataRulePurposeClass                                                       *string `json:"DataRulePurposeClass,omitempty"`
}
