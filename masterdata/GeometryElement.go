package masterdata

type GeometryElement struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates"`
	Type        PurpleType                     `json:"type"`
}
