package workproductcomponent

// ISO13503_2 sieve analysis data for this instance of ISO13503_2 data for the proppant
// agent.
type ISO135032SieveAnalysisData struct {
	// The percentage of mass retained in the sieve.                                                    
	PercentRetained                                                                            *float64 `json:"PercentRetained,omitempty"`
	// ASTM US Standard Mesh opening size used in the sieve analysis test. To indicate Pan, use         
	// 0.                                                                                               
	SieveNumber                                                                                *int64   `json:"SieveNumber,omitempty"`
}
