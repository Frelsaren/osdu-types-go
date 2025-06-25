package workproductcomponent

// The assessment score per data rule dimension type and its weight.
type DataDimensionMetric struct {
	// The reference to the DataRuleDimensionType, to which score and weight are associated.         
	DataRuleDimensionTypeID                                                                 string   `json:"DataRuleDimensionTypeID"`
	// The assessment score in % aggregated for this dimension type.                                 
	Score                                                                                   *float64 `json:"Score,omitempty"`
	// The weight factor for this score for the overall score.                                       
	Weight                                                                                  *float64 `json:"Weight,omitempty"`
}
