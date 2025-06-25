package workproduct

type AnyCRSGeoJSON struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates,omitempty"`
	Type        AnyCRSGeoJSONPointType         `json:"type"`
	Geometries  []GeometryElement              `json:"geometries,omitempty"`
}
