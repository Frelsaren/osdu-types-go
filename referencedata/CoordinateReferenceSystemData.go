package referencedata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Generic reference object containing the universal properties of reference data,
// especially the ones commonly thought of as Types
type CoordinateReferenceSystemData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                     
	ExistenceKind                                                                                *string                   `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                              
	ResourceCurationStatus                                                                       *string                   `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                      
	ResourceHomeRegionID                                                                         *string                   `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                   
	ResourceHostRegionIDs                                                                        []string                  `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                    
	ResourceLifecycleStatus                                                                      *string                   `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                      
	ResourceSecurityClassification                                                               *string                   `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                      
	// organization, agency, system, internal team, or individual. For informational purposes                              
	// only, the list of sources is not governed.                                                                          
	Source                                                                                       *string                   `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                               
	// based on data quality. Clarifications: Since Certified is the highest classification of                             
	// suitable quality, any further change or versioning of a Certified record should be                                  
	// carefully considered and justified. If a Technical Assurance value is not populated then                            
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                            
	// Technical Assurance values are not intended to be used for the identification of a single                           
	// "preferred" or "definitive" record by comparison with other records.                                                
	TechnicalAssuranceID                                                                         *string                   `json:"TechnicalAssuranceID,omitempty"`
	// Name of the authority, or organisation, which governs the entity value and from which it                            
	// is sourced.                                                                                                         
	AttributionAuthority                                                                         *string                   `json:"AttributionAuthority,omitempty"`
	// Name, URL, or other identifier of the publication, or repository, of the attribution                                
	// source organisation from which the entity value is sourced.                                                         
	AttributionPublication                                                                       *string                   `json:"AttributionPublication,omitempty"`
	// The distinct instance of the attribution publication, by version number, sequence number,                           
	// date of publication, etc., that was used for the entity value.                                                      
	AttributionRevision                                                                          *string                   `json:"AttributionRevision,omitempty"`
	// The abbreviation or mnemonic for a reference type if defined. Example: WELL and WLBR.                               
	Code                                                                                         *string                   `json:"Code,omitempty"`
	// For reference values published and governed by OSDU: The date and time the record was                               
	// committed into the OSDU member GitLab reference-values repository. The sole purpose of                              
	// this date is to optimise the OSDU milestone upgrades. It allows the upgrade code to                                 
	// figure out whether or not the record must be PUT into reference value storage.                                      
	CommitDate                                                                                   *time.Time                `json:"CommitDate,omitempty"`
	// The text which describes a NAME TYPE in detail.                                                                     
	Description                                                                                  *string                   `json:"Description,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                             
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                            
	// If used, the "Source" property should identify that source system.                                                  
	ID                                                                                           *string                   `json:"ID,omitempty"`
	// By default reference values are considered as 'active'. An absent 'InactiveIndicator'                               
	// property value means the reference value is in active use. When 'InactiveIndicator' is                              
	// set true the reverence value is no longer in use and should no longer be offered as a                               
	// choice.                                                                                                             
	InactiveIndicator                                                                            *bool                     `json:"InactiveIndicator,omitempty"`
	// The name of the entity instance.                                                                                    
	Name                                                                                         *string                   `json:"Name,omitempty"`
	// Alternative names, including historical, by which this entity instance is/has been known.                           
	NameAlias                                                                                    []AbstractAliasNames      `json:"NameAlias,omitempty"`
	// The base geographic CRS of this projected CRS. Only populated for                                                   
	// CoordinateReferenceSystemType==ProjectedCRS.                                                                        
	BaseCRS                                                                                      *BaseCRS                  `json:"BaseCRS,omitempty"`
	// The code as number as opposed to the Code defined as a string.                                                      
	CodeAsNumber                                                                                 *int64                    `json:"CodeAsNumber,omitempty"`
	// The namespace or authority name governing this CRS definition, e.g. EPSG for contents                               
	// from the EPSG Geodetic Parameter Dataset.                                                                           
	CodeSpace                                                                                    *string                   `json:"CodeSpace,omitempty"`
	// The type of coordinate reference system. This is an enumeration of concrete sub-types.                              
	CoordinateReferenceSystemType                                                                *CRSType                  `json:"CoordinateReferenceSystemType,omitempty"`
	// The coordinate system defining the dimension and individual axes used by the CRS.                                   
	CoordinateSystem                                                                             *CoordinateSystem         `json:"CoordinateSystem,omitempty"`
	// The datum of this CRS. Only populated for CoordinateReferenceSystemType in                                          
	// [GeographicCRS, VerticalCRS, EngineeringCRS].                                                                       
	Datum                                                                                        *Datum                    `json:"Datum,omitempty"`
	// The DatumEnsemble for the CRS's datum. Only populated for GeographicCRS.                                            
	DatumEnsemble                                                                                *DatumEnsemble            `json:"DatumEnsemble,omitempty"`
	// The horizontal CRS reference of a CompoundCRS. Only populated for                                                   
	// CoordinateReferenceSystemType==CompoundCRS.                                                                         
	HorizontalCRS                                                                                *HorizontalCRS            `json:"HorizontalCRS,omitempty"`
	// The InformationSource providing the CRS definition if different from AttributionAuthority.                          
	InformationSource                                                                            *string                   `json:"InformationSource,omitempty"`
	// The kind of CRS, e.g. bound, compound, derived, engineering, geocentric, geographic 2D,                             
	// geographic 3D, projected, vertical.                                                                                 
	Kind                                                                                         *string                   `json:"Kind,omitempty"`
	// Used for export and actionable instructions to a conversion/transformation engine. It is                            
	// initially based on Esri well-known text (WKT). Eventually, when Esri WKT are convertible                            
	// into ISO WKT and vice versa, the definition can be replaced by                                                      
	// https://proj.org/schemas/v0.2/projjson.schema.json.                                                                 
	PersistableReference                                                                         *string                   `json:"PersistableReference,omitempty"`
	// Scope and extent information about the described CRS.                                                               
	PreferredUsage                                                                               *PurplePreferredUsage     `json:"PreferredUsage,omitempty"`
	// The projection operation of a ProjectedCRS. Only populated for                                                      
	// CoordinateReferenceSystemType==ProjectedCRS.                                                                        
	Projection                                                                                   *Projection               `json:"Projection,omitempty"`
	// The revision date of this CRS.                                                                                      
	RevisionDate                                                                                 *time.Time                `json:"RevisionDate,omitempty"`
	// The source CRS of a BoundCRS. Only populated for CoordinateReferenceSystemType==BoundCRS.                           
	SourceCRS                                                                                    *PurpleSourceCRS          `json:"SourceCRS,omitempty"`
	// The target CRS of this bound CRS. Only populated for                                                                
	// CoordinateReferenceSystemType==BoundCRS.                                                                            
	TargetCRS                                                                                    *PurpleTargetCRS          `json:"TargetCRS,omitempty"`
	// The Transformation bound to the BaseCRS in a BoundCRS. Only populated for                                           
	// CoordinateReferenceSystemType==BoundCRS.                                                                            
	Transformation                                                                               *BoundTransformation      `json:"Transformation,omitempty"`
	// Contextual information about scope and extent/area of use.                                                          
	Usages                                                                                       []PurpleUsage             `json:"Usages,omitempty"`
	// The vertical CRS reference of a CompoundCRS. Only populated for                                                     
	// CoordinateReferenceSystemType==CompoundCRS.                                                                         
	VerticalCRS                                                                                  *VerticalCRSClass         `json:"VerticalCRS,omitempty"`
	// The 2-dimensional bounding box derived from the extent (Polygon or MultiPolygon) based on                           
	// WGS 84 (EPSG:4326). The schema of this substructure is identical to the GeoJSON                                     
	// FeatureCollection https://geojson.org/schema/FeatureCollection.json. The coordinate                                 
	// sequence follows GeoJSON standard, i.e. longitude, latitude. CoordinateReferenceSystems                             
	// with an extent crossing the anti-meridian are represented by a MultiPolygon.                                        
	Wgs84Coordinates                                                                             *GeoJSONFeatureCollection `json:"Wgs84Coordinates,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}    `json:"ExtensionProperties,omitempty"`
}
