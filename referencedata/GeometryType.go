package referencedata

type GeometryType string

const (
	PurpleLineString      GeometryType = "LineString"
	PurpleMultiLineString GeometryType = "MultiLineString"
	PurpleMultiPoint      GeometryType = "MultiPoint"
	PurpleMultiPolygon    GeometryType = "MultiPolygon"
	PurplePoint           GeometryType = "Point"
	PurplePolygon         GeometryType = "Polygon"
)
