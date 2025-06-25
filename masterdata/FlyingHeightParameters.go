package masterdata

// Captures the average flying height and the flying height parameters such as the
// instrumentation, datum and the clearance method.
type FlyingHeightParameters struct {
	// The average flying height for this airborne platform survey.                                      
	AverageFlyingHeight                                                                         *float64 `json:"AverageFlyingHeight,omitempty"`
	// The airborne survey ground clearance method, for example, constant or drape.                      
	ClearanceMethodID                                                                           *string  `json:"ClearanceMethodID,omitempty"`
	// The topographic information utilised in this survey, for example, Shuttle Radar                   
	// Topography Mission (SRTM) v3.                                                                     
	ExternalTopographicalInformation                                                            *string  `json:"ExternalTopographicalInformation,omitempty"`
	// The instruments or equipment used to measure the specified flying height. See                     
	// AverageFlyingHeight. This can be one or several different types of equipment, for                 
	// example, laser scanner, GPS, radar altimeter.                                                     
	MeasurementInstrumentationIDs                                                               []string `json:"MeasurementInstrumentationIDs,omitempty"`
	// The relationship to the vertical CRS defining the absolute reference surface. If both the         
	// vertical and horizontal CRS are the same, then the vertical CRS should instead be                 
	// captured only in the AbstractSpatialLocation.                                                     
	VerticalCoordinateReferenceSystemID                                                         *string  `json:"VerticalCoordinateReferenceSystemID,omitempty"`
}
