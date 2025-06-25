package workproductcomponent

// The contents positions the SeismicTraceData record in context of a time series. This is
// to be used for time lapse or 4D SeismicTraceData. This structure is optional and absent
// for SeismicTraceData not part of a time series.
type TimeLapse struct {
	// The index into the TimeSeriesID's data.UTCDateTimeValues[] or data.GeologicTimeValues[]        
	// arrays.                                                                                        
	TimeIndex                                                                                 *int64  `json:"TimeIndex,omitempty"`
	// The relationship to a TimeSeries work-product-component.                                       
	TimeSeriesID                                                                              *string `json:"TimeSeriesID,omitempty"`
}
