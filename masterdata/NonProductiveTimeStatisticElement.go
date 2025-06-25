package masterdata

// Structure that describes a individual statistic record of the operations from an activity.
type NonProductiveTimeStatisticElement struct {
	// Identifier of the probability type the value is describing        
	ProbabilityTypeID                                            string  `json:"ProbabilityTypeID"`
	// Record of the statistic estimation                                
	StatisticRecord                                              float64 `json:"StatisticRecord"`
}
