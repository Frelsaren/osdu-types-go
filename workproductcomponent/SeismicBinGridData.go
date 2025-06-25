package workproductcomponent

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Generic reference object containing the universal group-type properties of a Work Product
// Component for inclusion in data type specific Work Product Component objects
//
// Generic reference object containing the universal properties of a Work Product Component
// for inclusion in data type specific Work Product Component objects
//
// The embedded bin grid definition describing the mapping from pixel coordinate system to
// DataCRS ('world coordinates'). BinGrid member properties are only populated if BinGridID
// is absent. A populated BinGridID overrides any embedded BinGrid values.
//
// The shared properties for a bin grid.
type SeismicBinGridData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                                         
	ExistenceKind                                                                               *string                                        `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                                  
	ResourceCurationStatus                                                                      *string                                        `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                          
	ResourceHomeRegionID                                                                        *string                                        `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                                       
	ResourceHostRegionIDs                                                                       []string                                       `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                                        
	ResourceLifecycleStatus                                                                     *string                                        `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                          
	ResourceSecurityClassification                                                              *string                                        `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                          
	// organization, agency, system, internal team, or individual. For informational purposes                                                  
	// only, the list of sources is not governed.                                                                                              
	Source                                                                                      *string                                        `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                                   
	// based on data quality. Clarifications: Since Certified is the highest classification of                                                 
	// suitable quality, any further change or versioning of a Certified record should be                                                      
	// carefully considered and justified. If a Technical Assurance value is not populated then                                                
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                
	// Technical Assurance values are not intended to be used for the identification of a single                                               
	// "preferred" or "definitive" record by comparison with other records.                                                                    
	TechnicalAssuranceID                                                                        *string                                        `json:"TechnicalAssuranceID,omitempty"`
	// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct                                               
	// from the file, in the sense certain valuable information is generated during loading                                                    
	// process (Artefact generation process). Examples include retrieving location data,                                                       
	// performing an OCR which may result in the generation of artefacts which need to be                                                      
	// preserved distinctly                                                                                                                    
	Artefacts                                                                                   []AbstractGridRepresentationArtefact           `json:"Artefacts,omitempty"`
	// The record id, which identifies this OSDU File or dataset resource.                                                                     
	Datasets                                                                                    []string                                       `json:"Datasets,omitempty"`
	// An array of references to content in Domain Data Management Services represented by this                                                
	// work-product-component. The references are formed as URI following                                                                      
	// https://www.rfc-editor.org/rfc/rfc3986#page-16. This property is exclusively populated by                                               
	// DDMSs. If a work-product-component is represented in more than one DDMS, DDMSs are                                                      
	// obliged to find the specific reference by inspecting the URI's authority values matching                                                
	// the DDMS id.                                                                                                                            
	DDMSDatasets                                                                                []string                                       `json:"DDMSDatasets,omitempty"`
	// A flag that indicates if the work product component is searchable, which means covered in                                               
	// the search index.                                                                                                                       
	IsDiscoverable                                                                              *bool                                          `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product component is undergoing an extended load.  It                                                 
	// reflects the fact that the work product component is in an early stage and may be updated                                               
	// before finalization.                                                                                                                    
	IsExtendedLoad                                                                              *bool                                          `json:"IsExtendedLoad,omitempty"`
	// Alternative names, including historical, by which this work-product-component is/has been                                               
	// known (it should include all the identifiers).                                                                                          
	NameAliases                                                                                 []AbstractAliasNames                           `json:"NameAliases,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                                                 
	// quality. Clarifications: Since Certified is the highest classification of suitable                                                      
	// quality, any further change or versioning of a Certified record should be carefully                                                     
	// considered and justified. If a Technical Assurance value is not populated then one can                                                  
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                        
	// Technical Assurance values are not intended to be used for the identification of a single                                               
	// "preferred" or "definitive" record by comparison with other records.                                                                    
	TechnicalAssurances                                                                         []AbstractGridRepresentationTechnicalAssurance `json:"TechnicalAssurances,omitempty"`
	// Array of Authors' names of the work product component.  Could be a person or company                                                    
	// entity.                                                                                                                                 
	AuthorIDs                                                                                   []string                                       `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product component has been through                                                  
	// (ex. well planning, exploration).                                                                                                       
	BusinessActivities                                                                          []string                                       `json:"BusinessActivities,omitempty"`
	// Date that a resource (work  product component here) is formed outside of OSDU before                                                    
	// loading (e.g. publication date).                                                                                                        
	CreationDateTime                                                                            *time.Time                                     `json:"CreationDateTime,omitempty"`
	// Description.  Summary of the work product component.  Not the same as Remark which                                                      
	// captures thoughts of creator about the wpc.                                                                                             
	Description                                                                                 *string                                        `json:"Description,omitempty"`
	// List of geographic entities which provide context to the WPC.  This may include multiple                                                
	// types or multiple values of the same type.                                                                                              
	GeoContexts                                                                                 []AbstractGeoContext                           `json:"GeoContexts,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                                                    
	// product component depends.  The assertion is directed only from the asserting WPC to                                                    
	// ancestor objects, not children.  It should not be used to refer to files or artefacts                                                   
	// within the WPC -- the association within the WPC is sufficient and Artefacts are actually                                               
	// children of the main WPC file. They should be recorded in the data.Artefacts[] array.                                                   
	LineageAssertions                                                                           []LineageAssertion                             `json:"LineageAssertions,omitempty"`
	// Name                                                                                                                                    
	Name                                                                                        *string                                        `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product component                                                
	// (location of the subject matter).                                                                                                       
	SpatialArea                                                                                 *AbstractSpatialLocation                       `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product component                                                  
	// (location of the subject matter).                                                                                                       
	SpatialPoint                                                                                *AbstractSpatialLocation                       `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product component to OSDU.                                                             
	SubmitterName                                                                               *string                                        `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                                          
	Tags                                                                                        []string                                       `json:"Tags,omitempty"`
	// DEPRECATED: Use  AbstractGeoJson.PropertiesBinGridCorners properties inside the                                                         
	// ABCDBinGridSpatialLocation. Previously:  Array of 4 corner points for bin grid in local                                                 
	// coordinates: Point A (min inline, min crossline); Point B (min inline, max crossline);                                                  
	// Point C (max inline, min crossline); Point D (max inline, max crossline).  If Point D is                                                
	// not given and BinGridDefinitionMethodTypeID=4, it must be supplied, with its spatial                                                    
	// location, before ingestion to create a parallelogram in map coordinate space.  Note                                                     
	// correspondence of inline=x, crossline=y.                                                                                                
	ABCDBinGridLocalCoordinates                                                                 []AbstractCoordinates                          `json:"ABCDBinGridLocalCoordinates,omitempty"`
	// Bin Grid ABCD points containing the projected coordinates, projected CRS and quality                                                    
	// metadata.  This attribute is required also for the P6 definition method to define the                                                   
	// projected CRS, even if the ABCD coordinates would be optional (recommended to be always                                                 
	// calculated). It is recommended to populate the GeoJSON/AnyCrsGeoJSON with properties                                                    
	// according to the AbstractGeoJson.PropertiesBinGridCorners schema fragment.                                                              
	ABCDBinGridSpatialLocation                                                                  *AbstractSpatialLocation                       `json:"ABCDBinGridSpatialLocation,omitempty"`
	// This identifies how the Bin Grid is defined:  4=ABCD four-points method was used to                                                     
	// define the grid (P6 parameters are optional and can contain derived values;                                                             
	// P6BinNodeIncrementOnIAxis and P6BinNodeIncrementOnJaxis can be used as part of four-point                                               
	// method).  Use a perspective transformation to map between map coordinates and bin                                                       
	// coordinates. Note point order.  6=P6 definition method was used to define the bin grid                                                  
	// (ABCD points are optional and can contain derived values; ABCDBinGridSpatialLocation must                                               
	// specify the projected CRS).  Use an affine transformation to map between map coordinates                                                
	// and bin coordinates.                                                                                                                    
	BinGridDefinitionMethodTypeID                                                               *string                                        `json:"BinGridDefinitionMethodTypeID,omitempty"`
	// Name of bin grid (e.g., GEOCO_GREENCYN_PHV_2012).  Probably the name as it exists in a                                                  
	// separate corporate store if OSDU is not main system.                                                                                    
	BinGridName                                                                                 *string                                        `json:"BinGridName,omitempty"`
	// Type of bin grid (Acquisition, Processing, Velocity, MagGrav, Magnetics, Gravity,                                                       
	// GeologicModel, Reprojected, etc.)                                                                                                       
	BinGridTypeID                                                                               *string                                        `json:"BinGridTypeID,omitempty"`
	// Nominal design fold as intended by the bin grid definition, expressed as the mode in                                                    
	// percentage points (60 fold = 6000%).                                                                                                    
	CoveragePercent                                                                             *float64                                       `json:"CoveragePercent,omitempty"`
	// Easting coordinate of tie point (e.g., center or A point)                                                                               
	P6BinGridOriginEasting                                                                      *float64                                       `json:"P6BinGridOriginEasting,omitempty"`
	// Inline coordinate of tie point (e.g., center or A point)                                                                                
	P6BinGridOriginI                                                                            *float64                                       `json:"P6BinGridOriginI,omitempty"`
	// Crossline coordinate of tie point (e.g., center or A point)                                                                             
	P6BinGridOriginJ                                                                            *float64                                       `json:"P6BinGridOriginJ,omitempty"`
	// Northing coordinate of tie point (e.g., center or A point)                                                                              
	P6BinGridOriginNorthing                                                                     *float64                                       `json:"P6BinGridOriginNorthing,omitempty"`
	// Increment (positive integer) for the inline coordinate. If not provided then 1 is                                                       
	// assumed.  The bin grid definition is expected to have increment 1 and the increment                                                     
	// stored with the SeismicTraceData (“inline increment”) takes precedence over the increment                                               
	// set at the BinGrid.  Alternatively the increments are allowed to be defined with the                                                    
	// BinGrid, but this should be avoided to allow for variations in sampling in trace data                                                   
	// sets.                                                                                                                                   
	P6BinNodeIncrementOnIaxis                                                                   *int64                                         `json:"P6BinNodeIncrementOnIaxis,omitempty"`
	// Increment (positive integer) for the crossline coordinate. If not provided then 1 is                                                    
	// assumed.  The bin grid definition is expected to have increment 1 and the increment                                                     
	// stored with the SeismicTraceData (“crossline increment”) takes precedence over the                                                      
	// increment set at the BinGrid. Alternatively the increments are allowed to be defined with                                               
	// the BinGrid, but this should be avoided to allow for variations in sampling in trace data                                               
	// sets.                                                                                                                                   
	P6BinNodeIncrementOnJaxis                                                                   *int64                                         `json:"P6BinNodeIncrementOnJaxis,omitempty"`
	// Distance between two inlines at the given increment apart, e.g., 30 m with                                                              
	// P6BinNodeIncrementOnIaxis=1.  Unit from projected CRS in ABCDBinGridSpatialLocation                                                     
	P6BinWidthOnIaxis                                                                           *float64                                       `json:"P6BinWidthOnIaxis,omitempty"`
	// Distance between two crosslines at the given increment apart, e.g., 25 m with                                                           
	// P6BinNodeIncrementOnJaxis=4.  Unit from projected CRS in ABCDBinGridSpatialLocation                                                     
	P6BinWidthOnJaxis                                                                           *float64                                       `json:"P6BinWidthOnJaxis,omitempty"`
	// Clockwise angle from grid north (in projCRS) in degrees from 0 to 360 of the direction of                                               
	// increasing crosslines (constant inline), i.e., of the vector from point A to B.                                                         
	P6MapGridBearingOfBinGridJaxis                                                              *float64                                       `json:"P6MapGridBearingOfBinGridJaxis,omitempty"`
	// Scale factor for Bin Grid.  If not provided then 1 is assumed. Unit is unity.                                                           
	P6ScaleFactorOfBinGrid                                                                      *float64                                       `json:"P6ScaleFactorOfBinGrid,omitempty"`
	// EPSG code: 9666 for right-handed, 1049 for left-handed.  See IOGP Guidance Note 373-07-2                                                
	// and 483-6.                                                                                                                              
	P6TransformationMethod                                                                      *int64                                         `json:"P6TransformationMethod,omitempty"`
	// Identifier (name) of the corporate database/application that stores the source bin grid                                                 
	// definitions if OSDU is not main system.                                                                                                 
	SourceBinGridAppID                                                                          *string                                        `json:"SourceBinGridAppID,omitempty"`
	// Identifier of the source bin grid as stored in a corporate database/application if OSDU                                                 
	// is not main system.                                                                                                                     
	SourceBinGridID                                                                             *int64                                         `json:"SourceBinGridID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
