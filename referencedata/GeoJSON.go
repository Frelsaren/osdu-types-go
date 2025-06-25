package referencedata

type GeoJSON struct {
	Bbox        []float64                `json:"bbox,omitempty"`
	Coordinates []GeoJSONPointCoordinate `json:"coordinates,omitempty"`
	Type        GeoJSONPointType         `json:"type"`
	Geometries  []GeometryElement        `json:"geometries,omitempty"`
}
