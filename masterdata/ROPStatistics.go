package masterdata

// Individual Statistics that define the rate of penetration of an activity
type ROPStatistics struct {
	// Identifier of the probability type the value is describing        
	ProbabilityTypeID                                            string  `json:"ProbabilityTypeID"`
	// Record of the statistic estimation                                
	StatisticRecord                                              float64 `json:"StatisticRecord"`
}
