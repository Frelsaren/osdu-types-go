package workproductcomponent

type GeoJSON struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates,omitempty"`
	Type        GeoJSONPointType               `json:"type"`
	Geometries  []GeometryClass                `json:"geometries,omitempty"`
}
