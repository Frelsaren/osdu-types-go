package workproductcomponent

// Allows to link the geometry of the WellboreMarkerSet to a particular index of a time
// series. This is particularly useful for intervals referring to fluid contacts where the
// topology and geometry varies against the time.
//
// Defines the link between a work-product-component (likely a representation in the RESQML
// sense) and a time series work-product-component holding the entire list of time steps in
// a series.
//
// Allows to link the geometry of the WellboreMarkerSet to a particular index of a time
// series. This is particularly useful for markers referring to fluid contacts where the
// topology and geometry varies against the time.
type AbstractTimeSeriesLink struct {
	// Index of the timestamp of the representation in the associated TimeSeries       
	TimeIndex                                                                   int64  `json:"TimeIndex"`
	// Time series the representation is associated to                                 
	TimeSeriesID                                                                string `json:"TimeSeriesID"`
}
