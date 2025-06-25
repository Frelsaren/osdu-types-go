package referencedata

type GeoJSONFeature struct {
	Bbox       []float64              `json:"bbox,omitempty"`
	Geometry   *GeoJSON               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	Type       FeatureTypeEnum        `json:"type"`
}
