package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// A Project is a business activity that consumes financial and human resources and produces
// (digital) work products.
//
// The activity abstraction for projects and surveys (master-data).
type ConnectedSourceDataJobData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                       
	ExistenceKind                                                                               *string                      `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                
	ResourceCurationStatus                                                                      *string                      `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                        
	ResourceHomeRegionID                                                                        *string                      `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                     
	ResourceHostRegionIDs                                                                       []string                     `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                      
	ResourceLifecycleStatus                                                                     *string                      `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                        
	ResourceSecurityClassification                                                              *string                      `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                        
	// organization, agency, system, internal team, or individual. For informational purposes                                
	// only, the list of sources is not governed.                                                                            
	Source                                                                                      *string                      `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                 
	// based on data quality. Clarifications: Since Certified is the highest classification of                               
	// suitable quality, any further change or versioning of a Certified record should be                                    
	// carefully considered and justified. If a Technical Assurance value is not populated then                              
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                              
	// Technical Assurance values are not intended to be used for the identification of a single                             
	// "preferred" or "definitive" record by comparison with other records.                                                  
	TechnicalAssuranceID                                                                        *string                      `json:"TechnicalAssuranceID,omitempty"`
	// List of geographic entities which provide context to the master data. This may include                                
	// multiple types or multiple values of the same type.                                                                   
	GeoContexts                                                                                 []AbstractGeoContext         `json:"GeoContexts,omitempty"`
	// Alternative names, including historical, by which this master data is/has been known (it                              
	// should include all the identifiers).                                                                                  
	NameAliases                                                                                 []AbstractAliasNames         `json:"NameAliases,omitempty"`
	// The spatial location information such as coordinates, CRS information (left empty when                                
	// not appropriate).                                                                                                     
	SpatialLocation                                                                             *AbstractSpatialLocation     `json:"SpatialLocation,omitempty"`
	// Describes a record's overall suitability for general business consumption in context of                               
	// one or more workflows/personas based on data quality and reviewer's decisions.                                        
	// Clarifications: Since Certified is the highest classification of suitable quality, any                                
	// further change or versioning of a Certified record should be carefully considered and                                 
	// justified. If a Technical Assurance value is not populated then one can assume the data                               
	// has not been evaluated or its quality is unknown (=Unevaluated). Technical Assurance                                  
	// values are not intended to be used for the identification of a single "preferred" or                                  
	// "definitive" record by comparison with other records.                                                                 
	TechnicalAssurances                                                                         []AbstractTechnicalAssurance `json:"TechnicalAssurances,omitempty"`
	// DEPRECATED: (in favor of more nuanced TechnicalAssurances[] array) Describes a                                        
	// master-data record's overall suitability for general business consumption based on data                               
	// quality. Clarifications: Since Certified is the highest classification of suitable                                    
	// quality, any further change or versioning of a Certified record should be carefully                                   
	// considered and justified. If a Technical Assurance value is not populated then one can                                
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                      
	// Technical Assurance values are not intended to be used for the identification of a single                             
	// "preferred" or "definitive" record by comparison with other records.                                                  
	TechnicalAssuranceTypeID                                                                    *string                      `json:"TechnicalAssuranceTypeID,omitempty"`
	// This describes the reason that caused the creation of a new version of this master data.                              
	VersionCreationReason                                                                       *string                      `json:"VersionCreationReason,omitempty"`
	// References to applicable agreements in external contract database system of record.                                   
	ContractIDs                                                                                 []string                     `json:"ContractIDs,omitempty"`
	// References to organisations which supplied services to the Project.                                                   
	Contractors                                                                                 []Contractors                `json:"Contractors,omitempty"`
	// The history of expenditure approvals.                                                                                 
	FundsAuthorizations                                                                         []FundsAuthorizations        `json:"FundsAuthorizations,omitempty"`
	// The organisation which controlled the conduct of the project.                                                         
	Operator                                                                                    *string                      `json:"Operator,omitempty"`
	// List of key individuals supporting the Project.  This could be Abstracted for re-use, and                             
	// could reference a separate Persons master data object.                                                                
	Personnel                                                                                   []PurplePersonnel            `json:"Personnel,omitempty"`
	// The date and time when the Project was initiated.                                                                     
	ProjectBeginDate                                                                            *time.Time                   `json:"ProjectBeginDate,omitempty"`
	// The date and time when the Project was completed.                                                                     
	ProjectEndDate                                                                              *time.Time                   `json:"ProjectEndDate,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                               
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                              
	// If used, the "Source" property should identify that source system.                                                    
	ProjectID                                                                                   *string                      `json:"ProjectID,omitempty"`
	// The common or preferred name of a Project.                                                                            
	ProjectName                                                                                 *string                      `json:"ProjectName,omitempty"`
	// DEPRECATED: please use data.NameAliases. The history of Project names, codes, and other                               
	// business identifiers.                                                                                                 
	ProjectNames                                                                                []AbstractAliasNames         `json:"ProjectNames,omitempty"`
	// General parameters defining the configuration of the Project.  In the case of a seismic                               
	// acquisition project it is like receiver interval, source depth, source type.  In the case                             
	// of a processing project, it is like replacement velocity, reference datum above mean sea                              
	// level.                                                                                                                
	ProjectSpecifications                                                                       []ProjectSpecifications      `json:"ProjectSpecifications,omitempty"`
	// The history of life cycle states that the Project has been through..                                                  
	ProjectStates                                                                               []ProjectStates              `json:"ProjectStates,omitempty"`
	// Description of the objectives of a Project.                                                                           
	Purpose                                                                                     *string                      `json:"Purpose,omitempty"`
	// The (non-overlapping) historical activity states and effective start and termination                                  
	// dates. The last state is replicated in the single LastActivityState for simpler queries.                              
	ActivityStates                                                                              []AbstractActivityState      `json:"ActivityStates,omitempty"`
	// The relation to the ActivityTemplate carrying expected parameter definitions and default                              
	// values.                                                                                                               
	ActivityTemplateID                                                                          *string                      `json:"ActivityTemplateID,omitempty"`
	// The current or last state this activity transitioned to. It is a copy of the last element                             
	// in ActivityStates[]. If there is only one state recorded, the ActivityStates[] can stay                               
	// empty.                                                                                                                
	LastActivityState                                                                           *AbstractActivityState       `json:"LastActivityState,omitempty"`
	// General parameter value used in one instance of activity.  Includes reference to data                                 
	// objects which are inputs and outputs of the activity.                                                                 
	Parameters                                                                                  []DefaultValueElement        `json:"Parameters,omitempty"`
	// The relationship to a parent project acting as a parent activity.                                                     
	ParentProjectID                                                                             *string                      `json:"ParentProjectID,omitempty"`
	// Indicates if a scheduled job is active (will be executed) or not (won't be executed)                                  
	ActiveIndicator                                                                             bool                         `json:"ActiveIndicator"`
	// ID of the external partition containing the desired data                                                              
	ConnectedSourceDataPartitionID                                                              string                       `json:"ConnectedSourceDataPartitionID"`
	// ID reference of the parent Connected Source Registry Entry                                                            
	ConnectedSourceRegistryEntryID                                                              string                       `json:"ConnectedSourceRegistryEntryID"`
	// Source system schema authority.                                                                                       
	ConnectedSourceSchemaAuthority                                                              *string                      `json:"ConnectedSourceSchemaAuthority,omitempty"`
	// The maximum create/update time for data records (UTC).                                                                
	CreateTimeMax                                                                               *time.Time                   `json:"CreateTimeMax,omitempty"`
	// The wait time, in seconds, for the eds_Ingest DAG run upon completion of the Manifest                                 
	// Ingestion (osdu_ingest) DAG run, used to retrieve details from the XCom summary.                                      
	EdsIngestWaitTime                                                                           *float64                     `json:"EdsIngestWaitTime,omitempty"`
	// A list of external processes configuration to be executed by EDS                                                      
	ExternalProcesses                                                                           []ExternalProcess            `json:"ExternalProcesses,omitempty"`
	// A temporary solution before these references are stored in a related, external record for                             
	// scalability: The data type/schema/kind of data being retrieved form the external source.                              
	// The failed records list which is retrieved from the external source.                                                  
	FailedRecords                                                                               []string                     `json:"FailedRecords,omitempty"`
	// End time of the records to be fetched from the source system.                                                         
	FetchEndDateTime                                                                            *time.Time                   `json:"FetchEndDateTime,omitempty"`
	// The data type/schema/kind of data being retrieved form the external source. The returned                              
	// value should validate against the corresponding registered schema in the OSDU schema                                  
	// service.                                                                                                              
	FetchKind                                                                                   string                       `json:"FetchKind"`
	// Start time of the records to be fetched from the source system.                                                       
	FetchStartDateTime                                                                          *time.Time                   `json:"FetchStartDateTime,omitempty"`
	// Filter applied to the data fetch request, using data members in the FetchKind schema.                                 
	Filter                                                                                      string                       `json:"Filter"`
	// Indicates if the dummy parent data mapping should be considered or not, by default it                                 
	// will be false (and false if absent), if true ParentDataMappingDummyMasterIDs should be                                
	// set.                                                                                                                  
	IsDummyParentMappingEnabled                                                                 *bool                        `json:"IsDummyParentMappingEnabled,omitempty"`
	// The last successful run date of the job (UTC)                                                                         
	LastSuccessfulRunDateUTC                                                                    *time.Time                   `json:"LastSuccessfulRunDateUTC,omitempty"`
	// The maximum number of records to be processed in this job.                                                            
	LimitRecords                                                                                *int64                       `json:"LimitRecords,omitempty"`
	// Descriptive label given to a scheduled job.                                                                           
	Name                                                                                        string                       `json:"Name"`
	// List of access control lists (ACLs) to be injected into the manifests of external data                                
	// before ingestion                                                                                                      
	OnIngestionACL                                                                              *AccessControlList           `json:"OnIngestionAcl,omitempty"`
	// Consumer partition the incoming data will be placed in.                                                               
	OnIngestionDataPartitionID                                                                  string                       `json:"OnIngestionDataPartitionID"`
	// List of legal tags to be injected into the manifests of external data before ingestion                                
	OnIngestionLegalTags                                                                        *LegalMetaData               `json:"OnIngestionLegalTags,omitempty"`
	// Consumer schema authority for the incoming data will be placed in.                                                    
	OnIngestionSchemaAuthority                                                                  *string                      `json:"OnIngestionSchemaAuthority,omitempty"`
	// A fixed dummy master ID is set in advance and will be used if the master record is                                    
	// unavailable in the consumer system. Multiple master ids can be defined as array members;                              
	// ids are defined without data-partition example: 'master-data--Well:DummyWell'.                                        
	ParentDataMappingDummyMasterIDs                                                             []string                     `json:"ParentDataMappingDummyMasterIDs,omitempty"`
	// Schedule this job should run on, in CRON format                                                                       
	ScheduleUTC                                                                                 string                       `json:"ScheduleUTC"`
	// TriggerNaturalizationDAG (default false) triggers, if true, a naturalization DAG                                      
	// (directed acyclic graph) that will add the data file (SEG-Y, LAS, etc.) to the target                                 
	// OSDU Platform from the source system and convert the WPC's child dataset from "external"                              
	// to "internal".                                                                                                        
	TriggerNaturalizationDAG                                                                    *bool                        `json:"TriggerNaturalizationDAG,omitempty"`
	// DEPRECATED: Superseded by the contents of appropriate parameters in an ActivityTemplate                               
	// instance identified by data.ActivityTemplateID. In earlier versions: List of workflows                                
	// and their configuration used in this scheduled job                                                                    
	Workflows                                                                                   []Workflow                   `json:"Workflows,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
