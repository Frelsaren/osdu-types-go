// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    workProduct, err := UnmarshalWorkProduct(bytes)
//    bytes, err = workProduct.Marshal()

package workproduct

import "bytes"
import "errors"
import "time"

import "encoding/json"

func UnmarshalWorkProduct(data []byte) (WorkProduct, error) {
	var r WorkProduct
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WorkProduct) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// A collection of work product components such as might be produced by a business activity
// and which is delivered to the data platform for loading.
type WorkProduct struct {
	// The access control tags associated with this entity.                                                                     
	ACL                                                                                          AccessControlList              `json:"acl"`
	// The links to data, which constitute the inputs, from which this record instance is                                       
	// derived.                                                                                                                 
	Ancestry                                                                                     *ParentList                    `json:"ancestry,omitempty"`
	// Timestamp of the time at which initial version of this OSDU resource object was created.                                 
	// Set by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                    
	CreateTime                                                                                   *time.Time                     `json:"createTime,omitempty"`
	// The user reference, which created the first version of this resource object. Set by the                                  
	// System.                                                                                                                  
	CreateUser                                                                                   *string                        `json:"createUser,omitempty"`
	Data                                                                                         *Data                          `json:"data,omitempty"`
	// Previously called ResourceID or SRN which identifies this OSDU resource object without                                   
	// version.                                                                                                                 
	ID                                                                                           *string                        `json:"id,omitempty"`
	// The schema identification for the OSDU resource object following the pattern                                             
	// {Namespace}:{Source}:{Type}:{VersionMajor}.{VersionMinor}.{VersionPatch}. The versioning                                 
	// scheme follows the semantic versioning, https://semver.org/.                                                             
	Kind                                                                                         string                         `json:"kind"`
	// The entity's legal tags and compliance status. The actual contents associated with the                                   
	// legal tags is managed by the Compliance Service.                                                                         
	Legal                                                                                        LegalMetaData                  `json:"legal"`
	// The Frame of Reference meta data section linking the named properties to self-contained                                  
	// definitions.                                                                                                             
	Meta                                                                                         []FrameOfReferenceMetaDataItem `json:"meta,omitempty"`
	// Timestamp of the time at which this version of the OSDU resource object was created. Set                                 
	// by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                        
	ModifyTime                                                                                   *time.Time                     `json:"modifyTime,omitempty"`
	// The user reference, which created this version of this resource object. Set by the System.                               
	ModifyUser                                                                                   *string                        `json:"modifyUser,omitempty"`
	// A generic dictionary of string keys mapping to string value. Only strings are permitted                                  
	// as keys and values.                                                                                                      
	Tags                                                                                         map[string]string              `json:"tags,omitempty"`
	// The version number of this OSDU resource; set by the framework.                                                          
	Version                                                                                      *int64                         `json:"version,omitempty"`
}

// The access control tags associated with this entity.
//
// The access control tags associated with this entity. This structure is included by the
// SystemProperties "acl", which is part of all OSDU records. Not extensible.
type AccessControlList struct {
	// The list of owners of this data record formatted as an email                                    
	// (core.common.model.storage.validation.ValidationDoc.EMAIL_REGEX).                               
	Owners                                                                                    []string `json:"owners"`
	// The list of viewers to which this data record is accessible/visible/discoverable                
	// formatted as an email (core.common.model.storage.validation.ValidationDoc.EMAIL_REGEX).         
	Viewers                                                                                   []string `json:"viewers"`
}

// The links to data, which constitute the inputs, from which this record instance is
// derived.
//
// A list of entity id:version references to record instances recorded in the data platform,
// from which the current record is derived and from which the legal tags must be derived.
// This structure is included by the SystemProperties "ancestry", which is part of all OSDU
// records. Not extensible.
type ParentList struct {
	// An array of none, one or many entity references of 'direct parents' in the data platform,         
	// which mark the current record as a derivative. In contrast to other relationships, the            
	// source record version is required. During record creation or update the                           
	// ancestry.parents[] relationships are used to collect the legal tags from the sources and          
	// aggregate them in the legal.legaltags[] array. As a consequence, should e.g., one or more         
	// of the legal tags of the source data expire, the access to the derivatives is also                
	// terminated. For details, see ComplianceService tutorial, 'Creating derivative Records'.           
	Parents                                                                                     []string `json:"parents,omitempty"`
}

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
type Data struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                   
	ExistenceKind                                                                               *string                  `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                            
	ResourceCurationStatus                                                                      *string                  `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                    
	ResourceHomeRegionID                                                                        *string                  `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                 
	ResourceHostRegionIDs                                                                       []string                 `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                  
	ResourceLifecycleStatus                                                                     *string                  `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                    
	ResourceSecurityClassification                                                              *string                  `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                    
	// organization, agency, system, internal team, or individual. For informational purposes                            
	// only, the list of sources is not governed.                                                                        
	Source                                                                                      *string                  `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                             
	// based on data quality. Clarifications: Since Certified is the highest classification of                           
	// suitable quality, any further change or versioning of a Certified record should be                                
	// carefully considered and justified. If a Technical Assurance value is not populated then                          
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                          
	// Technical Assurance values are not intended to be used for the identification of a single                         
	// "preferred" or "definitive" record by comparison with other records.                                              
	TechnicalAssuranceID                                                                        *string                  `json:"TechnicalAssuranceID,omitempty"`
	// Array of Annotations                                                                                              
	Annotations                                                                                 []string                 `json:"Annotations,omitempty"`
	// Array of Authors' names of the work product.  Could be a person or company entity.                                
	AuthorIDs                                                                                   []string                 `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product has been through (ex. well                            
	// planning, exploration).                                                                                           
	BusinessActivities                                                                          []string                 `json:"BusinessActivities,omitempty"`
	Components                                                                                  []string                 `json:"Components,omitempty"`
	// Date that a resource (work  product here) is formed outside of OSDU before loading (e.g.                          
	// publication date, work product delivery package assembly date).                                                   
	CreationDateTime                                                                            *time.Time               `json:"CreationDateTime,omitempty"`
	// Description of the purpose of the work product.                                                                   
	Description                                                                                 *string                  `json:"Description,omitempty"`
	// A flag that indicates if the work product is searchable, which means covered in the                               
	// search index.                                                                                                     
	IsDiscoverable                                                                              *bool                    `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product is undergoing an extended load.  It reflects                            
	// the fact that the work product is in an early stage and may be updated before                                     
	// finalization.                                                                                                     
	IsExtendedLoad                                                                              *bool                    `json:"IsExtendedLoad,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                              
	// product depends.  The assertion is directed only from the asserting WP to ancestor                                
	// objects, not children.  It should not be used to refer to files or artefacts within the                           
	// WP -- the association within the WP is sufficient and Artefacts are actually children of                          
	// the main WP file. They should be recorded in the data.Artefacts[] array.                                          
	LineageAssertions                                                                           []LineageAssertion       `json:"LineageAssertions,omitempty"`
	// Name of the instance of Work Product - could be a shipment number.                                                
	Name                                                                                        *string                  `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product (location                          
	// of the subject matter).                                                                                           
	SpatialArea                                                                                 *AbstractSpatialLocation `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product (location of                         
	// the subject matter).                                                                                              
	SpatialPoint                                                                                *AbstractSpatialLocation `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product package to OSDU.                                         
	SubmitterName                                                                               *string                  `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                    
	Tags                                                                                        []string                 `json:"Tags,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}   `json:"ExtensionProperties,omitempty"`
}

type LineageAssertion struct {
	// The object reference identifying the DIRECT, INDIRECT, REFERENCE dependency.                      
	ID                                                                                           *string `json:"ID,omitempty"`
	// Used by LineageAssertion to describe the nature of the line of descent of a work product          
	// from a prior Resource, such as DIRECT, INDIRECT, REFERENCE.  It is not for proximity              
	// (number of nodes away), it is not to cover all the relationships in a full ontology or            
	// graph, and it is not to describe the type of activity that created the asserting WP.              
	// LineageAssertion does not encompass a full provenance, process history, or activity model.        
	LineageRelationshipType                                                                      *string `json:"LineageRelationshipType,omitempty"`
}

// A polygon boundary that reflects the locale of the content of the work product (location
// of the subject matter).
//
// A geographic object which can be described by a set of points.
//
// A centroid point that reflects the locale of the content of the work product (location of
// the subject matter).
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
}

type AnyCRSGeoJSONFeature struct {
	Bbox       []float64              `json:"bbox,omitempty"`
	Geometry   *AnyCRSGeoJSON         `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	Type       FluffyType             `json:"type"`
}

type AnyCRSGeoJSON struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates,omitempty"`
	Type        AnyCRSGeoJSONPointType         `json:"type"`
	Geometries  []GeometryElement              `json:"geometries,omitempty"`
}

type GeometryElement struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates"`
	Type        PurpleType                     `json:"type"`
}

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

type GeoJSONFeature struct {
	Bbox       []float64              `json:"bbox,omitempty"`
	Geometry   *GeoJSON               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
	Type       StickyType             `json:"type"`
}

type GeoJSON struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates,omitempty"`
	Type        GeoJSONPointType               `json:"type"`
	Geometries  []GeometryClass                `json:"geometries,omitempty"`
}

type GeometryClass struct {
	Bbox        []float64                      `json:"bbox,omitempty"`
	Coordinates []AnyCRSGeoJSONPointCoordinate `json:"coordinates"`
	Type        TentacledType                  `json:"type"`
}

// The entity's legal tags and compliance status. The actual contents associated with the
// legal tags is managed by the Compliance Service.
//
// Legal meta data like legal tags, relevant other countries, legal status. This structure
// is included by the SystemProperties "legal", which is part of all OSDU records. Not
// extensible.
type LegalMetaData struct {
	// The list of legal tags, which resolve to legal properties (like country of origin, export         
	// classification code, etc.) and rules with the help of the Compliance Service.                     
	Legaltags                                                                                   []string `json:"legaltags"`
	// The list of other relevant data countries as an array of two-letter country codes, see            
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.                                                 
	OtherRelevantDataCountries                                                                  []string `json:"otherRelevantDataCountries"`
	// The legal status. Set by the system after evaluation against the compliance rules                 
	// associated with the "legaltags" using the Compliance Service.                                     
	Status                                                                                      *string  `json:"status,omitempty"`
}

// A meta data item, which allows the association of named properties or property values to
// a Unit/Measurement/CRS/Azimuth/Time context.
type FrameOfReferenceMetaDataItem struct {
	// The kind of reference, 'Unit' for FrameOfReferenceUOM.                                                 
	//                                                                                                        
	// The kind of reference, constant 'CRS' for FrameOfReferenceCRS.                                         
	//                                                                                                        
	// The kind of reference, constant 'DateTime', for FrameOfReferenceDateTime.                              
	//                                                                                                        
	// The kind of reference, constant 'AzimuthReference', for FrameOfReferenceAzimuthReference.              
	Kind                                                                                        ReferenceKind `json:"kind"`
	// The unit symbol or name of the unit.                                                                   
	//                                                                                                        
	// The name of the CRS.                                                                                   
	//                                                                                                        
	// The name of the DateTime format and reference.                                                         
	//                                                                                                        
	// The name of the CRS or the symbol/name of the unit.                                                    
	Name                                                                                        *string       `json:"name,omitempty"`
	// The self-contained, persistable reference string uniquely identifying the Unit.                        
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying the CRS.                         
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying DateTime                         
	// reference.                                                                                             
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying AzimuthReference.                
	PersistableReference                                                                        string        `json:"persistableReference"`
	// The list of property names, to which this meta data item provides Unit context to. A full              
	// path like "StructureA.PropertyB" is required to define a unique context; "data" is                     
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides CRS context to. A full               
	// path like "StructureA.PropertyB" is required to define a unique context; "data" is                     
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides DateTime context to. A               
	// full path like "StructureA.PropertyB" is required to define a unique context; "data" is                
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides AzimuthReference                     
	// context to. A full path like "StructureA.PropertyB" is required to define a unique                     
	// context; "data" is omitted since frame-of reference normalization only applies to the                  
	// data block.                                                                                            
	PropertyNames                                                                               []string      `json:"propertyNames,omitempty"`
	// SRN to unit of measure reference.                                                                      
	UnitOfMeasureID                                                                             *string       `json:"unitOfMeasureID,omitempty"`
	// SRN to CRS reference.                                                                                  
	CoordinateReferenceSystemID                                                                 *string       `json:"coordinateReferenceSystemID,omitempty"`
}

type PurpleType string

const (
	PurpleAnyCRSLineString      PurpleType = "AnyCrsLineString"
	PurpleAnyCRSMultiLineString PurpleType = "AnyCrsMultiLineString"
	PurpleAnyCRSMultiPoint      PurpleType = "AnyCrsMultiPoint"
	PurpleAnyCRSMultiPolygon    PurpleType = "AnyCrsMultiPolygon"
	PurpleAnyCRSPoint           PurpleType = "AnyCrsPoint"
	PurpleAnyCRSPolygon         PurpleType = "AnyCrsPolygon"
)

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

type FluffyType string

const (
	AnyCRSFeature FluffyType = "AnyCrsFeature"
)

type AsIngestedCoordinatesType string

const (
	AnyCRSFeatureCollection AsIngestedCoordinatesType = "AnyCrsFeatureCollection"
)

type TentacledType string

const (
	PurpleLineString      TentacledType = "LineString"
	PurpleMultiLineString TentacledType = "MultiLineString"
	PurpleMultiPoint      TentacledType = "MultiPoint"
	PurpleMultiPolygon    TentacledType = "MultiPolygon"
	PurplePoint           TentacledType = "Point"
	PurplePolygon         TentacledType = "Polygon"
)

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

type StickyType string

const (
	Feature StickyType = "Feature"
)

type Wgs84CoordinatesType string

const (
	FeatureCollection Wgs84CoordinatesType = "FeatureCollection"
)

type ReferenceKind string

const (
	AzimuthReference ReferenceKind = "AzimuthReference"
	CRS              ReferenceKind = "CRS"
	DateTime         ReferenceKind = "DateTime"
	Unit             ReferenceKind = "Unit"
)

type AnyCRSGeoJSONPointCoordinate struct {
	Double     *float64
	UnionArray []PurpleCoordinate
}

func (x *AnyCRSGeoJSONPointCoordinate) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *AnyCRSGeoJSONPointCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

type PurpleCoordinate struct {
	Double     *float64
	UnionArray []FluffyCoordinate
}

func (x *PurpleCoordinate) UnmarshalJSON(data []byte) error {
	x.UnionArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.UnionArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *PurpleCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.UnionArray != nil, x.UnionArray, false, nil, false, nil, false, nil, false)
}

type FluffyCoordinate struct {
	Double      *float64
	DoubleArray []float64
}

func (x *FluffyCoordinate) UnmarshalJSON(data []byte) error {
	x.DoubleArray = nil
	object, err := unmarshalUnion(data, nil, &x.Double, nil, nil, true, &x.DoubleArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *FluffyCoordinate) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, x.Double, nil, nil, x.DoubleArray != nil, x.DoubleArray, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
			*pi = nil
	}
	if pf != nil {
			*pf = nil
	}
	if pb != nil {
			*pb = nil
	}
	if ps != nil {
			*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
			return false, err
	}

	switch v := tok.(type) {
	case json.Number:
			if pi != nil {
					i, err := v.Int64()
					if err == nil {
							*pi = &i
							return false, nil
					}
			}
			if pf != nil {
					f, err := v.Float64()
					if err == nil {
							*pf = &f
							return false, nil
					}
					return false, errors.New("Unparsable number")
			}
			return false, errors.New("Union does not contain number")
	case float64:
			return false, errors.New("Decoder should not return float64")
	case bool:
			if pb != nil {
					*pb = &v
					return false, nil
			}
			return false, errors.New("Union does not contain bool")
	case string:
			if haveEnum {
					return false, json.Unmarshal(data, pe)
			}
			if ps != nil {
					*ps = &v
					return false, nil
			}
			return false, errors.New("Union does not contain string")
	case nil:
			if nullable {
					return false, nil
			}
			return false, errors.New("Union does not contain null")
	case json.Delim:
			if v == '{' {
					if haveObject {
							return true, json.Unmarshal(data, pc)
					}
					if haveMap {
							return false, json.Unmarshal(data, pm)
					}
					return false, errors.New("Union does not contain object")
			}
			if v == '[' {
					if haveArray {
							return false, json.Unmarshal(data, pa)
					}
					return false, errors.New("Union does not contain array")
			}
			return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")
}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
			return json.Marshal(*pi)
	}
	if pf != nil {
			return json.Marshal(*pf)
	}
	if pb != nil {
			return json.Marshal(*pb)
	}
	if ps != nil {
			return json.Marshal(*ps)
	}
	if haveArray {
			return json.Marshal(pa)
	}
	if haveObject {
			return json.Marshal(pc)
	}
	if haveMap {
			return json.Marshal(pm)
	}
	if haveEnum {
			return json.Marshal(pe)
	}
	if nullable {
			return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
