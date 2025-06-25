package workproductcomponent

import "time"

// This is used to store the latest summary of data quality evaluation results for each
// RuleSet and is associated by reference to a work product component or master data object.
// Only one DataQualityReview per DataQualityRuleSet
type DataQualityReview struct {
	// This ID points to a Data Quality object used to store the result from a run of a Data                                           
	// Quality Metric Evaluation engine. Captures summary information, such as which rule-set(s)                                       
	// and rule(s) have been used, when this was run and by whom. Detailed results may be                                              
	// associated with a dataset.                                                                                                      
	DataQualityID                                                                               string                                 `json:"DataQualityID"`
	// The reference to the data quality rule set run for the data quality summary. Unique Key                                         
	// for the item in the DataQualitySummarySet array.                                                                                
	DataQualityRuleSetID                                                                        string                                 `json:"DataQualityRuleSetID"`
	// The name of the data quality rule set run for the data quality summary which should stay                                        
	// consistent with the DataQualityRuleSet. Denormalized for easier discoverability of the                                          
	// summaries, but needs to be kept up to date.                                                                                     
	DataQualityRuleSetName                                                                      *string                                `json:"DataQualityRuleSetName,omitempty"`
	// The individual scores per dimension.                                                                                            
	DimensionMetrics                                                                            []DataQualitySummarySetDimensionMetric `json:"DimensionMetrics,omitempty"`
	// The reference to the assessment calculation method.                                                                             
	MethodID                                                                                    *string                                `json:"MethodID,omitempty"`
	// The date-time at which the quality assessment started.                                                                          
	StartDateTime                                                                               *time.Time                             `json:"StartDateTime,omitempty"`
	// The assessment score in % aggregated for all dimension types.                                                                   
	TotalScore                                                                                  *float64                               `json:"TotalScore,omitempty"`
}
