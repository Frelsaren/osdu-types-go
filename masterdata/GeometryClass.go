package masterdata

type GeometryClass struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates"`
	Type        TentacledType                  `json:"type"`
}
