package masterdata

import "time"

// The spatial location information such as coordinates, CRS information (left empty when
// not appropriate).
//
// A geographic object which can be described by a set of points.
//
// 2D coordinates that defines the start of the section.
//
// The geographic location of the target
//
// Slot Location
//
// The bottom hole location of the wellbore denoted by a specified geographic horizontal
// coordinate reference system (Horizontal CRS), such as WGS84, NAD27, or ED50. If both
// GeographicBottomHoleLocation and ProjectedBottomHoleLocation properties are populated on
// this wellbore, they must identify the same point, just in different CRSs.
//
// The bottom hole location of the wellbore denoted by a projected horizontal coordinate
// reference system (Horizontal CRS), such a UTM zone. 'Projected' in this property does not
// mean 'planned' or 'projected-to-bit'. If both GeographicBottomHoleLocation and
// ProjectedBottomHoleLocation properties are populated on this wellbore, they must identify
// the same point, just in different CRSs.
type AbstractSpatialLocation struct {
	// The audit trail of operations applied to the coordinates from the original state to the                                   
	// current state. The list may contain operations applied prior to ingestion as well as the                                  
	// operations applied to produce the Wgs84Coordinates. The text elements refer to ESRI style                                 
	// CRS and Transformation names, which may have to be translated to EPSG standard names.                                     
	AppliedOperations                                                                           []string                         `json:"AppliedOperations,omitempty"`
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
	AsIngestedCoordinates                                                                       *AbstractAnyCRSFeatureCollection `json:"AsIngestedCoordinates,omitempty"`
	// The date of the Quality Check.                                                                                            
	CoordinateQualityCheckDateTime                                                              *time.Time                       `json:"CoordinateQualityCheckDateTime,omitempty"`
	// The user who performed the Quality Check.                                                                                 
	CoordinateQualityCheckPerformedBy                                                           *string                          `json:"CoordinateQualityCheckPerformedBy,omitempty"`
	// Freetext remarks on Quality Check.                                                                                        
	CoordinateQualityCheckRemarks                                                               []string                         `json:"CoordinateQualityCheckRemarks,omitempty"`
	// A qualitative description of the quality of a spatial location, e.g. unverifiable, not                                    
	// verified, basic validation.                                                                                               
	QualitativeSpatialAccuracyTypeID                                                            *string                          `json:"QualitativeSpatialAccuracyTypeID,omitempty"`
	// An approximate quantitative assessment of the quality of a location (accurate to > 500 m                                  
	// (i.e. not very accurate)), to < 1 m, etc.                                                                                 
	QuantitativeAccuracyBandID                                                                  *string                          `json:"QuantitativeAccuracyBandID,omitempty"`
	// Indicates the expected look of the SpatialParameterType, e.g. Point, MultiPoint,                                          
	// LineString, MultiLineString, Polygon, MultiPolygon. The value constrains the type of                                      
	// geometries in the GeoJSON Wgs84Coordinates and AsIngestedCoordinates.                                                     
	SpatialGeometryTypeID                                                                       *string                          `json:"SpatialGeometryTypeID,omitempty"`
	// Date when coordinates were measured or retrieved.                                                                         
	SpatialLocationCoordinatesDate                                                              *time.Time                       `json:"SpatialLocationCoordinatesDate,omitempty"`
	// A type of spatial representation of an object, often general (e.g. an Outline, which                                      
	// could be applied to Field, Reservoir, Facility, etc.) or sometimes specific (e.g. Onshore                                 
	// Outline, State Offshore Outline, Federal Offshore Outline, 3 spatial representations that                                 
	// may be used by Countries).                                                                                                
	SpatialParameterTypeID                                                                      *string                          `json:"SpatialParameterTypeID,omitempty"`
	// The normalized coordinates (Point, MultiPoint, LineString, MultiLineString, Polygon or                                    
	// MultiPolygon) based on WGS 84 (EPSG:4326 for 2-dimensional coordinates, EPSG:4326 +                                       
	// EPSG:5714 (MSL) for 3-dimensional coordinates). This derived coordinate representation is                                 
	// intended for global discoverability only. The schema of this substructure is identical to                                 
	// the GeoJSON FeatureCollection https://geojson.org/schema/FeatureCollection.json. The                                      
	// coordinate sequence follows GeoJSON standard, i.e. longitude, latitude {, height}                                         
	Wgs84Coordinates                                                                            *GeoJSONFeatureCollection        `json:"Wgs84Coordinates,omitempty"`
}
