package workproduct

// The normalized coordinates (Point, MultiPoint, LineString, MultiLineString, Polygon or
// MultiPolygon) based on WGS 84 (EPSG:4326 for 2-dimensional coordinates, EPSG:4326 +
// EPSG:5714 (MSL) for 3-dimensional coordinates). This derived coordinate representation is
// intended for global discoverability only. The schema of this substructure is identical to
// the GeoJSON FeatureCollection https://geojson.org/schema/FeatureCollection.json. The
// coordinate sequence follows GeoJSON standard, i.e. longitude, latitude {, height}
//
// GeoJSON feature collection as originally published in
// https://geojson.org/schema/FeatureCollection.json. Attention: the coordinate order is
// fixed: Longitude first, followed by Latitude, optionally height above MSL (EPSG:5714) as
// third coordinate.
type GeoJSONFeatureCollection struct {
	Bbox     []float64            `json:"bbox,omitempty"`
	Features []GeoJSONFeature     `json:"features"`
	Type     Wgs84CoordinatesType `json:"type"`
}
