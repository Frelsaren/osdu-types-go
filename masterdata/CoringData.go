package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// The activity abstraction for projects and surveys (master-data).
type CoringData struct {
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
	// The relation to the ActivityTemplate carrying expected parameter definitions and default                                      
	// values.                                                                                                                       
	ActivityTemplateID                                                                          *string                              `json:"ActivityTemplateID,omitempty"`
	// General parameter value used in one instance of activity.  Includes reference to data                                         
	// objects which are inputs and outputs of the activity.                                                                         
	Parameters                                                                                  []PurpleAbstractActivityParameter    `json:"Parameters,omitempty"`
	// The relationship to a parent project acting as a parent activity.                                                             
	ParentProjectID                                                                             *string                              `json:"ParentProjectID,omitempty"`
	// The depth of the base of the core. The reference and kind of depth (e.g. driller's depth                                      
	// versus logger's depth) is described in data.VerticalMeasurement. For SidewallCores this                                       
	// is the depth of the deepest core.                                                                                             
	BottomDepth                                                                                 *float64                             `json:"BottomDepth,omitempty"`
	// The diameter of the core.                                                                                                     
	CoreDiameter                                                                                *float64                             `json:"CoreDiameter,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                                       
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                                      
	// If used, the "Source" property should identify that source system. i.e. this item is                                          
	// optional.                                                                                                                     
	CoreIdentifier                                                                              *string                              `json:"CoreIdentifier,omitempty"`
	// The vendor assigned core number.                                                                                              
	CoreNumber                                                                                  *string                              `json:"CoreNumber,omitempty"`
	// The date the core returned to surface.                                                                                        
	CoreRecoveredDate                                                                           *string                              `json:"CoreRecoveredDate,omitempty"`
	// An array contains narrative remarks pertaining to a core.                                                                     
	CoreRemarks                                                                                 []CoreRemark                         `json:"CoreRemarks,omitempty"`
	// The date of the coring operation.                                                                                             
	CoringOperationDate                                                                         *string                              `json:"CoringOperationDate,omitempty"`
	// Flag indicating whether a detailed tripping schedule is available. Available tripping                                         
	// schedules are usually a sign of controlled coring sample quality.                                                             
	HasTrippingSchedule                                                                         *bool                                `json:"HasTrippingSchedule,omitempty"`
	// Indicates if the core was oriented.                                                                                           
	IsOriented                                                                                  *bool                                `json:"IsOriented,omitempty"`
	// The name of the core. For example ACME 1 Core 1                                                                               
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// The WellLog representation carrying the preferred core-depth to logging-depth correction                                      
	// or mapping.                                                                                                                   
	PreferredDepthShiftsID                                                                      *string                              `json:"PreferredDepthShiftsID,omitempty"`
	// The kind of preservation applied to this conventional coring or sidewall core.                                                
	PreservationTypeID                                                                          *string                              `json:"PreservationTypeID,omitempty"`
	// The coring run number.                                                                                                        
	RunNumber                                                                                   *int64                               `json:"RunNumber,omitempty"`
	// The coring company that extracted the well core. For Example: ACME Limited PLC                                                
	ServiceCompanyID                                                                            *string                              `json:"ServiceCompanyID,omitempty"`
	// The depth of the top of the core. The reference and kind of depth (e.g. driller's depth                                       
	// versus logger's depth) is described in data.VerticalMeasurement. For SidewallCores this                                       
	// is the depth of the shallowest core.                                                                                          
	TopDepth                                                                                    *float64                             `json:"TopDepth,omitempty"`
	// References an entry in the VerticalMeasurements array for the Wellbore identified by                                          
	// WellboreID, or a standalone vertical reference which defines the vertical reference datum                                     
	// for all measured depths of the Coring record.                                                                                 
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// The relationship to the wellbore.                                                                                             
	WellboreID                                                                                  *string                              `json:"WellboreID,omitempty"`
	// The Conventional Coring specific sub-structure.                                                                               
	ConventionalCoring                                                                          *ConventionalCoring                  `json:"ConventionalCoring,omitempty"`
	// The Sidewall Coring specific sub-structure.                                                                                   
	SidewallCoring                                                                              *SidewallCoring                      `json:"SidewallCoring,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
