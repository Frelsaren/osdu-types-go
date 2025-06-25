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
type ConnectedSourceRegistryEntryData struct {
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
	// References to applicable agreements governing the use of the data source                                              
	AgreementIDs                                                                                []string                     `json:"AgreementIDs,omitempty"`
	// Connectivity information for Airflow endpoints to get more information of Manifest                                    
	// Ingestion.                                                                                                            
	AirflowStableAPIURL                                                                         *string                      `json:"AirflowStableAPIUrl,omitempty"`
	// A placeholder to keep the data provider Dataset URL.                                                                  
	DatasetURL                                                                                  *string                      `json:"DatasetURL,omitempty"`
	// Additional information/description about the data source                                                              
	Description                                                                                 *string                      `json:"Description,omitempty"`
	// Flag that determines whether the external source has a full OSDU implementation (true) or                             
	// a wrapper facade over proprietary APIs (false)                                                                        
	FullOSDUImplementationIndicator                                                             *bool                        `json:"FullOSDUImplementationIndicator,omitempty"`
	// Descriptive label given to the data source. This could be the name of an organisation                                 
	// and/or the name of a specific database or system.                                                                     
	Name                                                                                        string                       `json:"Name"`
	// DEPRECATED: Please use reference-data--ExternalReferenceValueMapping reference catalog                                
	// items instead. Temporary property awaiting a external reference-value mapping framework.                              
	// A generic dictionary of string reference-data as keys mapping to reference-value as                                   
	// string value. Only predefined reference-data and its values are permitted.                                            
	ReferenceValueMappings                                                                      map[string]interface{}       `json:"ReferenceValueMappings,omitempty"`
	// A placeholder to store the data provider Search service URL.                                                          
	SearchURL                                                                                   *string                      `json:"SearchURL,omitempty"`
	// List of security schemes available for use in authorizing against OSDU-compliant APIs of                              
	// a connected data source.                                                                                              
	SecuritySchemes                                                                             []SecurityScheme             `json:"SecuritySchemes"`
	// List of SMTP server schemes available for use in mailing the detailed EDS's report.                                   
	SMTPSchemes                                                                                 []SMTPScheme                 `json:"SmtpSchemes,omitempty"`
	// Identifier of the organisation that the registered source belongs to.                                                 
	SourceOrganisationID                                                                        *string                      `json:"SourceOrganisationID,omitempty"`
	// A placeholder to keep the data provider Storage service URL.                                                          
	StorageURL                                                                                  *string                      `json:"StorageURL,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
