package referencedata

// The list of DataQualityRule items that this DataQualityRuleSet consists of.
type PurpleDataRule struct {
	// The reference to the DataQualityRule record.                                       
	DataRuleID                                                                   string   `json:"DataRuleID"`
	// Indicated if the DataQualityRule is required to pass or for information.           
	DataRulePurposeTypeID                                                        *string  `json:"DataRulePurposeTypeID,omitempty"`
	// The weight factor assigned to this rule in the current DataQualityRuleSet.         
	Weight                                                                       *float64 `json:"Weight,omitempty"`
}
