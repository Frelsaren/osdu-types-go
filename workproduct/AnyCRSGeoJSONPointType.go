package workproduct

type AnyCRSGeoJSONPointType string

const (
	AnyCRSGeometryCollection    AnyCRSGeoJSONPointType = "AnyCrsGeometryCollection"
	FluffyAnyCRSLineString      AnyCRSGeoJSONPointType = "AnyCrsLineString"
	FluffyAnyCRSMultiLineString AnyCRSGeoJSONPointType = "AnyCrsMultiLineString"
	FluffyAnyCRSMultiPoint      AnyCRSGeoJSONPointType = "AnyCrsMultiPoint"
	FluffyAnyCRSMultiPolygon    AnyCRSGeoJSONPointType = "AnyCrsMultiPolygon"
	FluffyAnyCRSPoint           AnyCRSGeoJSONPointType = "AnyCrsPoint"
	FluffyAnyCRSPolygon         AnyCRSGeoJSONPointType = "AnyCrsPolygon"
)
