package workproductcomponent

// The original or 'as ingested' coordinates (Point, MultiPoint, LineString,
// MultiLineString, Polygon or MultiPolygon). The name 'AsIngestedCoordinates' was chosen to
// contrast it to 'OriginalCoordinates', which carries the uncertainty whether any
// coordinate operations took place before ingestion. In cases where the original CRS is
// different from the as-ingested CRS, the AppliedOperations can also contain the list of
// operations applied to the coordinate prior to ingestion. The data structure is similar to
// GeoJSON FeatureCollection, however in a CRS context explicitly defined within the
// AbstractAnyCrsFeatureCollection. The coordinate sequence follows GeoJSON standard, i.e.
// 'eastward/longitude', 'northward/latitude' {, 'upward/height' unless overridden by an
// explicit direction in the AsIngestedCoordinates.VerticalCoordinateReferenceSystemID}.
//
// A schema like GeoJSON FeatureCollection with a non-WGS 84 CRS context; based on
// https://geojson.org/schema/FeatureCollection.json. Attention: the coordinate order is
// fixed: Longitude/Easting/Westing/X first, followed by Latitude/Northing/Southing/Y,
// optionally height as third coordinate.
type AbstractAnyCRSFeatureCollection struct {
	Bbox                                                                                        []float64                 `json:"bbox,omitempty"`
	// The CRS reference into the CoordinateReferenceSystem catalog.                                                      
	CoordinateReferenceSystemID                                                                 *string                   `json:"CoordinateReferenceSystemID,omitempty"`
	Features                                                                                    []AnyCRSGeoJSONFeature    `json:"features"`
	// The CRS reference as persistableReference string. If populated, the                                                
	// CoordinateReferenceSystemID takes precedence.                                                                      
	PersistableReferenceCRS                                                                     string                    `json:"persistableReferenceCrs"`
	// The unit of measure for the Z-axis (only for 3-dimensional coordinates, where the CRS                              
	// does not describe the vertical unit). Note that the direction is upwards positive, i.e. Z                          
	// means height.                                                                                                      
	PersistableReferenceUnitZ                                                                   *string                   `json:"persistableReferenceUnitZ,omitempty"`
	// The VerticalCRS reference as persistableReference string. If populated, the                                        
	// VerticalCoordinateReferenceSystemID takes precedence. The property is null or empty for                            
	// 2D geometries. For 3D geometries and absent or null persistableReferenceVerticalCrs the                            
	// vertical CRS is either provided via persistableReferenceCrs's CompoundCRS or it is                                 
	// implicitly defined as EPSG:5714 MSL height.                                                                        
	PersistableReferenceVerticalCRS                                                             *string                   `json:"persistableReferenceVerticalCrs,omitempty"`
	Type                                                                                        AsIngestedCoordinatesType `json:"type"`
	// The explicit VerticalCRS reference into the CoordinateReferenceSystem catalog. This                                
	// property stays empty for 2D geometries. Absent or empty values for 3D geometries mean the                          
	// context may be provided by a CompoundCRS in 'CoordinateReferenceSystemID' or implicitly                            
	// EPSG:5714 MSL height                                                                                               
	VerticalCoordinateReferenceSystemID                                                         *string                   `json:"VerticalCoordinateReferenceSystemID,omitempty"`
	// The explicit vertical unit ID, referring to a reference-data--UnitOfMeasure record; this                           
	// is only required for features containing 3-dimensional coordinates and undefined vertical                          
	// CoordinateReferenceSystems; if a VerticalCoordinateReferenceSystemID is populated, the                             
	// VerticalUnitID is given by the VerticalCoordinateReferenceSystemID's                                               
	// data.CoordinateSystem.VerticalAxisUnitID. The VerticalUnitID definition overrides any                              
	// self-contained definition in persistableReferenceUnitZ.                                                            
	VerticalUnitID                                                                              *string                   `json:"VerticalUnitID,omitempty"`
}
