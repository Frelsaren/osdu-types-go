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
// The purpose of this schema is best understood in the context of a columnar dataset: the
// AbstractReferencePropertyType describes a column in a columnar dataset by declaring its
// value type (number, string), a UnitQuantity if the value type is a number, a kind if the
// string value is actually a relationship to a e.g. reference-data type.
type GenericPropertyData struct {
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
	// Ordered array with: FacetType, FacetRole, both calling specific references                                                              
	//                                                                                                                                         
	// FacetType: Enumerations of the type of additional context about the nature of a property                                                
	// type (it may include conditions, direction, qualifiers, or statistics).                                                                 
	//                                                                                                                                         
	// FacetRole: Additional context about the nature of a property type. The purpose of such                                                  
	// attribute is to minimize the need to create specialized property types by mutualizing                                                   
	// some well known qualifiers such as "maximum", "minimum" which apply to a lot of different                                               
	// property types.                                                                                                                         
	FacetIDs                                                                                    []AbstractFacet                                `json:"FacetIDs,omitempty"`
	// It holds the PropertyType associated with this reference property type, further defining                                                
	// the semantics of the value. It contains a relationship to PropertyType record and its                                                   
	// (de-normalized) name. String or number values can represent e.g. A date or a time by                                                    
	// referring to the respective PropertyType record id.                                                                                     
	PropertyType                                                                                *AbstractPropertyType                          `json:"PropertyType,omitempty"`
	// Only populated if ValueType=="string" and the values are expected to represent record                                                   
	// ids, e.g. to a reference-data type, then this value holds the kind (optionally without                                                  
	// the semantic version number).                                                                                                           
	RelationshipTargetKind                                                                      *string                                        `json:"RelationshipTargetKind,omitempty"`
	// Only populated of the ValueType is "number". It holds the UnitQuantity associated with                                                  
	// this reference property type. It is a relationship to UnitQuantity record.                                                              
	UnitQuantityID                                                                              *string                                        `json:"UnitQuantityID,omitempty"`
	// The number of values in a tuple, e.g. For coordinates. The default is 1.                                                                
	ValueCount                                                                                  *int64                                         `json:"ValueCount,omitempty"`
	// The type of value to expect for this reference property, either "number" (floating point                                                
	// number), "integer",  "string", or "boolean".                                                                                            
	ValueType                                                                                   *string                                        `json:"ValueType,omitempty"`
	// The srn of the Column Based Table associated to a categorical property. Can be based on a                                               
	// defined dictionary of rock types                                                                                                        
	ClassificationTableID                                                                       *string                                        `json:"ClassificationTableID,omitempty"`
	// Indexable elements are used to link property values and geometry to a representation -                                                  
	// for instance a property could be attached to a flow grid's cell or face, or to a                                                        
	// Triangulated surface's point                                                                                                            
	IndexableElementID                                                                          *string                                        `json:"IndexableElementID,omitempty"`
	// Maximum value of the Property                                                                                                           
	MaxValue                                                                                    *float64                                       `json:"MaxValue,omitempty"`
	// Mean value of the Property                                                                                                              
	MeanValue                                                                                   *float64                                       `json:"MeanValue,omitempty"`
	// Minimum value of the Property                                                                                                           
	MinValue                                                                                    *float64                                       `json:"MinValue,omitempty"`
	// The srn of the topology the property refers to (WPC srn)                                                                                
	PropertyTopologyID                                                                          *string                                        `json:"PropertyTopologyID,omitempty"`
	// Unit of Measure of the property                                                                                                         
	PropertyUnitID                                                                              *string                                        `json:"PropertyUnitID,omitempty"`
	// Optional element indicating the realization index (metadata). Used if the property is the                                               
	// result of a multi-realization process.                                                                                                  
	RealizationIndices                                                                          []float64                                      `json:"RealizationIndices,omitempty"`
	// Standard deviation value of the Property                                                                                                
	StdDeviation                                                                                *float64                                       `json:"StdDeviation,omitempty"`
	// When using time series, allow to associate a single stamp to the property, if not present                                               
	// the property contains all time stamps of the time series.                                                                               
	TimeIndices                                                                                 *float64                                       `json:"TimeIndices,omitempty"`
	// Time series the property is associated to (srn)                                                                                         
	TimeSeriesID                                                                                *string                                        `json:"TimeSeriesID,omitempty"`
	// Always present when attached to time values, even when time values defined through time                                                 
	// series (array of date-times). If the property is attached to a time-series object,                                                      
	// denormalized array of time to ease search                                                                                               
	TimeValues                                                                                  []string                                       `json:"TimeValues,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
