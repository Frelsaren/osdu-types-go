package referencedata

type GeoJSONPointType string

const (
	FluffyLineString      GeoJSONPointType = "LineString"
	FluffyMultiLineString GeoJSONPointType = "MultiLineString"
	FluffyMultiPoint      GeoJSONPointType = "MultiPoint"
	FluffyMultiPolygon    GeoJSONPointType = "MultiPolygon"
	FluffyPoint           GeoJSONPointType = "Point"
	FluffyPolygon         GeoJSONPointType = "Polygon"
	GeometryCollection    GeoJSONPointType = "GeometryCollection"
)
