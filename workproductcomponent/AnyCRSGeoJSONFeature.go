package workproductcomponent

type AnyCRSGeoJSONFeature struct {
	Bbox       []float64              `json:"bbox,omitempty"`
	Geometry   *AnyCRSGeoJSON         `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	Type       FluffyType             `json:"type"`
}
