package referencedata

type GeometryElement struct {
	Bbox        []float64                `json:"bbox,omitempty"`
	Coordinates []GeoJSONPointCoordinate `json:"coordinates"`
	Type        GeometryType             `json:"type"`
}
