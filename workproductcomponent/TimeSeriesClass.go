package workproductcomponent

// Allow to link the geometry of the representation to a particular index of a time series.
// This is particularly useful for IJK grids used in geomechanical or basin context where
// the topology and geometry varies against the time.
type TimeSeriesClass struct {
	// Index of the timestamp of the representation in the associated TimeSeries       
	TimeIndex                                                                   int64  `json:"TimeIndex"`
	// Time series the representation is associated to                                 
	TimeSeriesID                                                                string `json:"TimeSeriesID"`
}
