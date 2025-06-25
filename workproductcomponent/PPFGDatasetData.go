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
type PPFGDatasetData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                                          
	ExistenceKind                                                                                *string                                        `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                                   
	ResourceCurationStatus                                                                       *string                                        `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                           
	ResourceHomeRegionID                                                                         *string                                        `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                                        
	ResourceHostRegionIDs                                                                        []string                                       `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                                         
	ResourceLifecycleStatus                                                                      *string                                        `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                           
	ResourceSecurityClassification                                                               *string                                        `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                           
	// organization, agency, system, internal team, or individual. For informational purposes                                                   
	// only, the list of sources is not governed.                                                                                               
	Source                                                                                       *string                                        `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                                    
	// based on data quality. Clarifications: Since Certified is the highest classification of                                                  
	// suitable quality, any further change or versioning of a Certified record should be                                                       
	// carefully considered and justified. If a Technical Assurance value is not populated then                                                 
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                 
	// Technical Assurance values are not intended to be used for the identification of a single                                                
	// "preferred" or "definitive" record by comparison with other records.                                                                     
	TechnicalAssuranceID                                                                         *string                                        `json:"TechnicalAssuranceID,omitempty"`
	// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct                                                
	// from the file, in the sense certain valuable information is generated during loading                                                     
	// process (Artefact generation process). Examples include retrieving location data,                                                        
	// performing an OCR which may result in the generation of artefacts which need to be                                                       
	// preserved distinctly                                                                                                                     
	Artefacts                                                                                    []AbstractGridRepresentationArtefact           `json:"Artefacts,omitempty"`
	// The record id, which identifies this OSDU File or dataset resource.                                                                      
	Datasets                                                                                     []string                                       `json:"Datasets,omitempty"`
	// An array of references to content in Domain Data Management Services represented by this                                                 
	// work-product-component. The references are formed as URI following                                                                       
	// https://www.rfc-editor.org/rfc/rfc3986#page-16. This property is exclusively populated by                                                
	// DDMSs. If a work-product-component is represented in more than one DDMS, DDMSs are                                                       
	// obliged to find the specific reference by inspecting the URI's authority values matching                                                 
	// the DDMS id.                                                                                                                             
	DDMSDatasets                                                                                 []string                                       `json:"DDMSDatasets,omitempty"`
	// A flag that indicates if the work product component is searchable, which means covered in                                                
	// the search index.                                                                                                                        
	IsDiscoverable                                                                               *bool                                          `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product component is undergoing an extended load.  It                                                  
	// reflects the fact that the work product component is in an early stage and may be updated                                                
	// before finalization.                                                                                                                     
	IsExtendedLoad                                                                               *bool                                          `json:"IsExtendedLoad,omitempty"`
	// Alternative names, including historical, by which this work-product-component is/has been                                                
	// known (it should include all the identifiers).                                                                                           
	NameAliases                                                                                  []AbstractAliasNames                           `json:"NameAliases,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                                                  
	// quality. Clarifications: Since Certified is the highest classification of suitable                                                       
	// quality, any further change or versioning of a Certified record should be carefully                                                      
	// considered and justified. If a Technical Assurance value is not populated then one can                                                   
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                         
	// Technical Assurance values are not intended to be used for the identification of a single                                                
	// "preferred" or "definitive" record by comparison with other records.                                                                     
	TechnicalAssurances                                                                          []AbstractGridRepresentationTechnicalAssurance `json:"TechnicalAssurances,omitempty"`
	// Array of Authors' names of the work product component.  Could be a person or company                                                     
	// entity.                                                                                                                                  
	AuthorIDs                                                                                    []string                                       `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product component has been through                                                   
	// (ex. well planning, exploration).                                                                                                        
	BusinessActivities                                                                           []string                                       `json:"BusinessActivities,omitempty"`
	// Date that a resource (work  product component here) is formed outside of OSDU before                                                     
	// loading (e.g. publication date).                                                                                                         
	CreationDateTime                                                                             *time.Time                                     `json:"CreationDateTime,omitempty"`
	// Description.  Summary of the work product component.  Not the same as Remark which                                                       
	// captures thoughts of creator about the wpc.                                                                                              
	Description                                                                                  *string                                        `json:"Description,omitempty"`
	// List of geographic entities which provide context to the WPC.  This may include multiple                                                 
	// types or multiple values of the same type.                                                                                               
	GeoContexts                                                                                  []AbstractGeoContext                           `json:"GeoContexts,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                                                     
	// product component depends.  The assertion is directed only from the asserting WPC to                                                     
	// ancestor objects, not children.  It should not be used to refer to files or artefacts                                                    
	// within the WPC -- the association within the WPC is sufficient and Artefacts are actually                                                
	// children of the main WPC file. They should be recorded in the data.Artefacts[] array.                                                    
	LineageAssertions                                                                            []LineageAssertion                             `json:"LineageAssertions,omitempty"`
	// Name                                                                                                                                     
	Name                                                                                         *string                                        `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product component                                                 
	// (location of the subject matter).                                                                                                        
	SpatialArea                                                                                  *AbstractSpatialLocation                       `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product component                                                   
	// (location of the subject matter).                                                                                                        
	SpatialPoint                                                                                 *AbstractSpatialLocation                       `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product component to OSDU.                                                              
	SubmitterName                                                                                *string                                        `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                                           
	Tags                                                                                         []string                                       `json:"Tags,omitempty"`
	// The characters that represent absent curve values in this data set, for example  '-999',                                                 
	// 'NULL', '0', etc. Typically for legacy data                                                                                              
	AbsentValueCharacters                                                                        *string                                        `json:"AbsentValueCharacters,omitempty"`
	// Open comments from the calculation team                                                                                                  
	Comment                                                                                      *string                                        `json:"Comment,omitempty"`
	// ID that reflects the context of the PPFG  data set, for example 'Pre-Drill' or                                                           
	// 'Post-Drill'                                                                                                                             
	ContextTypeID                                                                                *string                                        `json:"ContextTypeID,omitempty"`
	// Array of curve that constitutes the whole PPFG Dataset                                                                                   
	Curves                                                                                       []PurpleCurves                                 `json:"Curves,omitempty"`
	// Free text to describe the type of gauge used for the pressure measurement                                                                
	GaugeType                                                                                    *string                                        `json:"GaugeType,omitempty"`
	// IDs of the offset Wellbores included in the context and calculations of this PPFG data set                                               
	OffsetWellboreIDs                                                                            []string                                       `json:"OffsetWellboreIDs,omitempty"`
	// ID of the PPFG curve that is the primary reference or index. Derived from the PPFG curve                                                 
	// ID                                                                                                                                       
	PrimaryReferenceCurveID                                                                      *string                                        `json:"PrimaryReferenceCurveID,omitempty"`
	// The type of the primary reference, for example 'TVDSS',  'MD' , 'TWT'                                                                    
	PrimaryReferenceType                                                                         *string                                        `json:"PrimaryReferenceType,omitempty"`
	// The date that the PPFG data set was created by the PPFG practitioner or contractor                                                       
	RecordDate                                                                                   *time.Time                                     `json:"RecordDate,omitempty"`
	// Id of the Reference WellTrajectory used for TVD's calculation                                                                            
	ReferenceWellTrajectoryID                                                                    *string                                        `json:"ReferenceWellTrajectoryID,omitempty"`
	// ID of the service Company that acquired the PPFG                                                                                         
	ServiceCompanyID                                                                             *string                                        `json:"ServiceCompanyID,omitempty"`
	// Tectonic Scenario Setting for Planning and Pore Pressure Practitioners. Built into                                                       
	// interpretive curves. Can be, for example 'Strike Slip'                                                                                   
	TectonicSetting                                                                              *string                                        `json:"TectonicSetting,omitempty"`
	// The Vertical Measurement for the Wellbore identified which defines the vertical reference                                                
	// pressure datum for the related PPFG curves in this data set. The pressure datum is used                                                  
	// to calculate the average pressure gradient in                                                                                            
	VerticalMeasurement                                                                          *AbstractFacilityVerticalMeasurement           `json:"VerticalMeasurement,omitempty"`
	// ID from the Wellbore where the PPFG Work Product Component was recorded                                                                  
	WellboreID                                                                                   *string                                        `json:"WellboreID,omitempty"`
	// ID from the Well where the PPFG Work Product Component was recorded                                                                      
	WellID                                                                                       *string                                        `json:"WellID,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
