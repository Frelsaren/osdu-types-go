package workproductcomponent

// A set of properties describing a marker property which is available for this instance of
// a WellboreMarkerSet.
type MarkerProperty struct {
	// The reference to a marker property type - or if interpreted as CSV columns, the                
	// 'well-known column type. It is a relationship to a reference-data--MarkerPropertyType          
	// record id.                                                                                     
	MarkerPropertyTypeID                                                                      *string `json:"MarkerPropertyTypeID,omitempty"`
	// Unit of Measure for the marker properties of type MarkerPropertyType.                          
	MarkerPropertyUnitID                                                                      *string `json:"MarkerPropertyUnitID,omitempty"`
	// The name of the marker property (e.g. column in a CSV document) as originally found. If        
	// absent The name of the MarkerPropertyType is intended to be used.                              
	Name                                                                                      *string `json:"Name,omitempty"`
}
