package referencedata

// The 2-dimensional bounding box derived from the extent (Polygon or MultiPolygon) based on
// WGS 84 (EPSG:4326). The schema of this substructure is identical to the GeoJSON
// FeatureCollection https://geojson.org/schema/FeatureCollection.json. The coordinate
// sequence follows GeoJSON standard, i.e. longitude, latitude. CoordinateReferenceSystems
// with an extent crossing the anti-meridian are represented by a MultiPolygon.
//
// GeoJSON feature collection as originally published in
// https://geojson.org/schema/FeatureCollection.json. Attention: the coordinate order is
// fixed: Longitude first, followed by Latitude, optionally height above MSL (EPSG:5714) as
// third coordinate.
//
// The 2-dimensional bounding box derived from the extent (Polygon or MultiPolygon) based on
// WGS 84 (EPSG:4326). The schema of this substructure is identical to the GeoJSON
// FeatureCollection https://geojson.org/schema/FeatureCollection.json. The coordinate
// sequence follows GeoJSON standard, i.e. longitude, latitude. CoordinateTransformations
// with an extent crossing the anti-meridian are represented by a MultiPolygon.
type GeoJSONFeatureCollection struct {
	Bbox     []float64            `json:"bbox,omitempty"`
	Features []GeoJSONFeature     `json:"features"`
	Type     Wgs84CoordinatesType `json:"type"`
}
