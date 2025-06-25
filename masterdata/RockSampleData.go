package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type RockSampleData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                               
	ExistenceKind                                                                               *string                              `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                        
	ResourceCurationStatus                                                                      *string                              `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                
	ResourceHomeRegionID                                                                        *string                              `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                             
	ResourceHostRegionIDs                                                                       []string                             `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                              
	ResourceLifecycleStatus                                                                     *string                              `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                
	ResourceSecurityClassification                                                              *string                              `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                
	// organization, agency, system, internal team, or individual. For informational purposes                                        
	// only, the list of sources is not governed.                                                                                    
	Source                                                                                      *string                              `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                         
	// based on data quality. Clarifications: Since Certified is the highest classification of                                       
	// suitable quality, any further change or versioning of a Certified record should be                                            
	// carefully considered and justified. If a Technical Assurance value is not populated then                                      
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                      
	// Technical Assurance values are not intended to be used for the identification of a single                                     
	// "preferred" or "definitive" record by comparison with other records.                                                          
	TechnicalAssuranceID                                                                        *string                              `json:"TechnicalAssuranceID,omitempty"`
	// List of geographic entities which provide context to the master data. This may include                                        
	// multiple types or multiple values of the same type.                                                                           
	GeoContexts                                                                                 []AbstractGeoContext                 `json:"GeoContexts,omitempty"`
	// Alternative names, including historical, by which this master data is/has been known (it                                      
	// should include all the identifiers).                                                                                          
	NameAliases                                                                                 []AbstractAliasNames                 `json:"NameAliases,omitempty"`
	// The spatial location information such as coordinates, CRS information (left empty when                                        
	// not appropriate).                                                                                                             
	SpatialLocation                                                                             *AbstractSpatialLocation             `json:"SpatialLocation,omitempty"`
	// Describes a record's overall suitability for general business consumption in context of                                       
	// one or more workflows/personas based on data quality and reviewer's decisions.                                                
	// Clarifications: Since Certified is the highest classification of suitable quality, any                                        
	// further change or versioning of a Certified record should be carefully considered and                                         
	// justified. If a Technical Assurance value is not populated then one can assume the data                                       
	// has not been evaluated or its quality is unknown (=Unevaluated). Technical Assurance                                          
	// values are not intended to be used for the identification of a single "preferred" or                                          
	// "definitive" record by comparison with other records.                                                                         
	TechnicalAssurances                                                                         []AbstractTechnicalAssurance         `json:"TechnicalAssurances,omitempty"`
	// DEPRECATED: (in favor of more nuanced TechnicalAssurances[] array) Describes a                                                
	// master-data record's overall suitability for general business consumption based on data                                       
	// quality. Clarifications: Since Certified is the highest classification of suitable                                            
	// quality, any further change or versioning of a Certified record should be carefully                                           
	// considered and justified. If a Technical Assurance value is not populated then one can                                        
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                              
	// Technical Assurance values are not intended to be used for the identification of a single                                     
	// "preferred" or "definitive" record by comparison with other records.                                                          
	TechnicalAssuranceTypeID                                                                    *string                              `json:"TechnicalAssuranceTypeID,omitempty"`
	// This describes the reason that caused the creation of a new version of this master data.                                      
	VersionCreationReason                                                                       *string                              `json:"VersionCreationReason,omitempty"`
	// For cuttings: a flag indication whether the cuttings were washed and dried. More details                                      
	// in the associated RockSampleAnalysis.                                                                                         
	AreCuttingsWashedAndDried                                                                   *bool                                `json:"AreCuttingsWashedAndDried,omitempty"`
	// The depth of the base of the rock sample. For point measurements (slides, sidewall                                            
	// cores), the same value is assigned to TopDepth and BottomDepth.                                                               
	BottomDepth                                                                                 *float64                             `json:"BottomDepth,omitempty"`
	// Identifies the Coring from which this sample is created from it directly. Not populated                                       
	// for outcrops and cuttings.                                                                                                    
	CoringID                                                                                    *string                              `json:"CoringID,omitempty"`
	// The person, vendor or other provider of this information.                                                                     
	DataSource                                                                                  *string                              `json:"DataSource,omitempty"`
	// A flag to denote whether a particular core has been slabbed or not.                                                           
	IsCoreSlabbed                                                                               *bool                                `json:"IsCoreSlabbed,omitempty"`
	// For sample plugs: a flag indicating whether the plug has been cleaned. More details in                                        
	// the associated RockSampleAnalysis.                                                                                            
	IsPlugCleaned                                                                               *bool                                `json:"IsPlugCleaned,omitempty"`
	// The data vendor assigned sample ID or number.                                                                                 
	LabSampleIdentifier                                                                         *string                              `json:"LabSampleIdentifier,omitempty"`
	// The name of this RockSample.                                                                                                  
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// Identifies the RockSample from which this sample is created from it directly.                                                 
	ParentSampleID                                                                              *string                              `json:"ParentSampleID,omitempty"`
	// The kind of preservation applied to this sample.                                                                              
	PreservationTypeID                                                                          *string                              `json:"PreservationTypeID,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                                       
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                                      
	// If used, the "Source" property should identify that source system. i.e. this item is                                          
	// optional.                                                                                                                     
	RockSampleIdentifier                                                                        *string                              `json:"RockSampleIdentifier,omitempty"`
	// The date that the sample was acquired.                                                                                        
	SampleAcquiredDate                                                                          *time.Time                           `json:"SampleAcquiredDate,omitempty"`
	// The diameter of rock sample. Not applicable to cuttings                                                                       
	SampleDiameter                                                                              *float64                             `json:"SampleDiameter,omitempty"`
	// The interval between the top depth and bottom depth.                                                                          
	SampleInterval                                                                              *float64                             `json:"SampleInterval,omitempty"`
	// The length of rock sample. Not applicable to cuttings                                                                         
	SampleLength                                                                                *float64                             `json:"SampleLength,omitempty"`
	// The kind of orientation of this sample with respect to the bedding or drilling direction.                                     
	// Typical values are Horizontal, Vertical, Axial.                                                                               
	SampleOrientationID                                                                         *string                              `json:"SampleOrientationID,omitempty"`
	// Company and/or organization that owns the sample.                                                                             
	SampleOwnerID                                                                               *string                              `json:"SampleOwnerID,omitempty"`
	// An array containing operational or quality comments pertaining to a rock sample.                                              
	SampleRemarks                                                                               []SampleRemarks                      `json:"SampleRemarks,omitempty"`
	// An array containing the name of the locations where the material sample is stored. It can                                     
	// be stored in more than one location over time.                                                                                
	SampleStorageLocations                                                                      []AbstractStorageLocation            `json:"SampleStorageLocations,omitempty"`
	// Identifies a rock sample type.  E.g. Core, Cuttings, Core Slab, Core Plug, Core Chip,                                         
	// Slides. Considered mandatory.                                                                                                 
	SampleTypeID                                                                                *string                              `json:"SampleTypeID,omitempty"`
	// Weight of sample                                                                                                              
	SampleWeight                                                                                *float64                             `json:"SampleWeight,omitempty"`
	// The depth of the top of the rock sample. For point measurements (slides, sidewall cores),                                     
	// the same value is assigned to TopDepth and BottomDepth.                                                                       
	TopDepth                                                                                    *float64                             `json:"TopDepth,omitempty"`
	// References an entry in the VerticalMeasurements array for the Wellbore identified by                                          
	// WellboreID, or a standalone vertical reference which defines the vertical reference datum                                     
	// for all measured depths of the RockSample record. If this is not populated, the                                               
	// VerticalMeasurement is derived from the Coring.                                                                               
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// Unique wellbore identifier. Not required for outcrops.                                                                        
	WellboreID                                                                                  *string                              `json:"WellboreID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
