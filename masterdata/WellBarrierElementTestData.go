package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// A Project is a business activity that consumes financial and human resources and produces
// (digital) work products.
type WellBarrierElementTestData struct {
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
	// The ID of the Blowout Preventer associated with the test.                                                             
	BlowOutPreventerID                                                                          *string                      `json:"BlowOutPreventerID,omitempty"`
	// The ID of the Hole Section in which the test was performed.                                                           
	HoleSectionID                                                                               *string                      `json:"HoleSectionID,omitempty"`
	// Textual description of the operational procedure and decision points.                                                 
	OperationalProcedureDescription                                                             *string                      `json:"OperationalProcedureDescription,omitempty"`
	// The ID of the Rig associated with the test.                                                                           
	RigUtilizationID                                                                            *string                      `json:"RigUtilizationID,omitempty"`
	// Date/time well integrity test was performed.                                                                          
	TestDateTime                                                                                *time.Time                   `json:"TestDateTime,omitempty"`
	// Description or summary text.                                                                                          
	TestDescription                                                                             *string                      `json:"TestDescription,omitempty"`
	// The guidance associated with the test.                                                                                
	TestGuidance                                                                                *string                      `json:"TestGuidance,omitempty"`
	// The load case associated with the test.                                                                               
	TestLoadCase                                                                                *string                      `json:"TestLoadCase,omitempty"`
	// One or more test steps, their parameters and their criteria.                                                          
	TestSteps                                                                                   []WBETestStep                `json:"TestSteps"`
	// The Well Barrier Element Test Sub-Type - typically the barrier component being tested.                                
	// For example, a test of a PackerHanger. Free text or Open list (perhaps local).                                        
	TestSubTypeID                                                                               *string                      `json:"TestSubTypeID,omitempty"`
	// References an Open reference list for the Well Barrier Element Test Type.                                             
	// Different test types will have different templated pass/fail or acceptance criteria. The                              
	// list is Reference data has governance open to allow for addition of different test types                              
	// (e.g. BSEE BOP test and A. N. Other BOP Test).                                                                        
	TestTypeID                                                                                  string                       `json:"TestTypeID"`
	// The ID of the Tubular Assembly associated with the test.                                                              
	TubularAssemblyID                                                                           *string                      `json:"TubularAssemblyID,omitempty"`
	// The name of the Well Barrier Element Test Plan.                                                                       
	WBETestPlanName                                                                             *string                      `json:"WBETestPlanName,omitempty"`
	// The ID of the Well Activity associated with the test.                                                                 
	WellActivityID                                                                              *string                      `json:"WellActivityID,omitempty"`
	// The ID of the Well Activity Program that holds the Phase description.                                                 
	WellActivityProgramID                                                                       *string                      `json:"WellActivityProgramID,omitempty"`
	// Number that identifies the Phase within its WellActivityProgram.                                                      
	WellActivityProgramPhaseSequenceNumber                                                      *float64                     `json:"WellActivityProgramPhaseSequenceNumber,omitempty"`
	// The ID of the Wellbore Architecture applicable at the time of the test.                                               
	WellboreArchitectureID                                                                      *string                      `json:"WellboreArchitectureID,omitempty"`
	// The ID of the Wellbore on which the test was or will be performed. This ID can be used to                             
	// relate to either the Wellbore or WellPlanningWellbore object, depending on the scenario.                              
	WellboreID                                                                                  string                       `json:"WellboreID"`
	// The ID of the Well Operations Report where the barrier element test operation is                                      
	// described.                                                                                                            
	WellOperationsReportID                                                                      *string                      `json:"WellOperationsReportID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
