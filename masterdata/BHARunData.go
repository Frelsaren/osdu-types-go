package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// A Project is a business activity that consumes financial and human resources and produces
// (digital) work products.
type BHARunData struct {
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
	// References to applicable agreements in external contract database system of record.                                           
	ContractIDs                                                                                 []string                             `json:"ContractIDs,omitempty"`
	// References to organisations which supplied services to the Project.                                                           
	Contractors                                                                                 []Contractors                        `json:"Contractors,omitempty"`
	// The history of expenditure approvals.                                                                                         
	FundsAuthorizations                                                                         []FundsAuthorizations                `json:"FundsAuthorizations,omitempty"`
	// The organisation which controlled the conduct of the project.                                                                 
	Operator                                                                                    *string                              `json:"Operator,omitempty"`
	// List of key individuals supporting the Project.  This could be Abstracted for re-use, and                                     
	// could reference a separate Persons master data object.                                                                        
	Personnel                                                                                   []PurplePersonnel                    `json:"Personnel,omitempty"`
	// The date and time when the Project was initiated.                                                                             
	ProjectBeginDate                                                                            *time.Time                           `json:"ProjectBeginDate,omitempty"`
	// The date and time when the Project was completed.                                                                             
	ProjectEndDate                                                                              *time.Time                           `json:"ProjectEndDate,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                                       
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                                      
	// If used, the "Source" property should identify that source system.                                                            
	ProjectID                                                                                   *string                              `json:"ProjectID,omitempty"`
	// The common or preferred name of a Project.                                                                                    
	ProjectName                                                                                 *string                              `json:"ProjectName,omitempty"`
	// DEPRECATED: please use data.NameAliases. The history of Project names, codes, and other                                       
	// business identifiers.                                                                                                         
	ProjectNames                                                                                []AbstractAliasNames                 `json:"ProjectNames,omitempty"`
	// General parameters defining the configuration of the Project.  In the case of a seismic                                       
	// acquisition project it is like receiver interval, source depth, source type.  In the case                                     
	// of a processing project, it is like replacement velocity, reference datum above mean sea                                      
	// level.                                                                                                                        
	ProjectSpecifications                                                                       []ProjectSpecifications              `json:"ProjectSpecifications,omitempty"`
	// The history of life cycle states that the Project has been through..                                                          
	ProjectStates                                                                               []ProjectStates                      `json:"ProjectStates,omitempty"`
	// Description of the objectives of a Project.                                                                                   
	Purpose                                                                                     *string                              `json:"Purpose,omitempty"`
	// Date and time that activities started (BHA started to be made up, then Run In Hole)                                           
	ActivityStartDateTime                                                                       *time.Time                           `json:"ActivityStartDateTime,omitempty"`
	// Date and time that activities stopped (BHA was POOH).                                                                         
	ActivityStopDateTime                                                                        *time.Time                           `json:"ActivityStopDateTime,omitempty"`
	// Actual Performance of the BHA                                                                                                 
	ActualPerformanceDescription                                                                *string                              `json:"ActualPerformanceDescription,omitempty"`
	// Bit run number.                                                                                                               
	BitRunNumber                                                                                *string                              `json:"BitRunNumber,omitempty"`
	// Dogleg severity over the depth range of the BHA run.                                                                          
	Dogleg                                                                                      *float64                             `json:"Dogleg,omitempty"`
	// The parameters that were actually used during the BHA run                                                                     
	DrillingParams                                                                              []DrillingParameters                 `json:"DrillingParams,omitempty"`
	// Start on bottom - date and time.                                                                                              
	DrillingStartDateTime                                                                       *time.Time                           `json:"DrillingStartDateTime,omitempty"`
	// Start off bottom - date and time.                                                                                             
	DrillingStopDateTime                                                                        *time.Time                           `json:"DrillingStopDateTime,omitempty"`
	// Diameter of the hole drilled by the BHA. Note that Hole Size determine by the diameter of                                     
	// the Drill Bit and/or a Hole Opener/Under-reamer                                                                               
	HoleDiameter                                                                                *float64                             `json:"HoleDiameter,omitempty"`
	// This represents a foreign key to the Hole Section in which this BHA Run was performed.                                        
	HoleSectionID                                                                               *string                              `json:"HoleSectionID,omitempty"`
	// Part or all of the BHA is left in the hole                                                                                    
	IsLeftInHole                                                                                *bool                                `json:"IsLeftInHole,omitempty"`
	// The BHA component(s) contained a Nuclear Source.                                                                              
	IsNuclearSource                                                                             *bool                                `json:"IsNuclearSource,omitempty"`
	// Dogleg severity - Maximum.                                                                                                    
	MaximumDogleg                                                                               *float64                             `json:"MaximumDogleg,omitempty"`
	// Measured depth at run start. Depth relative to Planned wellbore ZDP. Navigate via                                             
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                                           
	// data.VerticalMeasurement.                                                                                                     
	MeasuredDepthRunStart                                                                       *float64                             `json:"MeasuredDepthRunStart,omitempty"`
	// Measured depth at run stop. Depth relative to Planned wellbore ZDP. Navigate via                                              
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                                           
	// data.VerticalMeasurement.                                                                                                     
	MeasuredDepthRunStop                                                                        *float64                             `json:"MeasuredDepthRunStop,omitempty"`
	// Human recognizable context for the BHA run.                                                                                   
	Name                                                                                        string                               `json:"Name"`
	// Objective of bottom hole assembly.                                                                                            
	ObjectiveBha                                                                                *string                              `json:"ObjectiveBha,omitempty"`
	// Planned/anticipated Dogleg severity over the depth range of the BHA run.                                                      
	PlannedDogleg                                                                               *float64                             `json:"PlannedDogleg,omitempty"`
	// Description of planned or expected Performance of the BHA                                                                     
	PredictedPerformanceDescription                                                             *string                              `json:"PredictedPerformanceDescription,omitempty"`
	// Free text allowing any comment associated to the run                                                                          
	RunComment                                                                                  *string                              `json:"RunComment,omitempty"`
	// Identifiers of the associated run parameter plans.                                                                            
	RunParameterPlans                                                                           []RunParameterPlan                   `json:"RunParameterPlans,omitempty"`
	// Bottom hole assembly status.                                                                                                  
	StatusBhaID                                                                                 *string                              `json:"StatusBhaID,omitempty"`
	// The BHA (drilling string) run number.                                                                                         
	StringRunNumber                                                                             *float64                             `json:"StringRunNumber,omitempty"`
	// Reason for trip.                                                                                                              
	TripReasonID                                                                                *string                              `json:"TripReasonID,omitempty"`
	// Duration of time to run BHA from surface to bottom of the hole                                                                
	TripTimeIn                                                                                  *float64                             `json:"TripTimeIn,omitempty"`
	// Duration of time to pull BHA from bottom of the hole to surface                                                               
	TripTimeOut                                                                                 *float64                             `json:"TripTimeOut,omitempty"`
	// This represents a foreign key to the tubular (assembly) that was utilized in this run.                                        
	TubularID                                                                                   string                               `json:"TubularID"`
	// References an entry in the VerticalMeasurements array for the Wellbore identified by                                          
	// WellboreID, or a standalone vertical reference elevation for all measured depths within                                       
	// the BHA Run record. If this is not populated, the VerticalMeasurement is derived from the                                     
	// Wellbore default Vertical Measure Elevation.                                                                                  
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// Unique identifier for the wellbore.  This uniquely represents                                                                 
	// the wellbore referenced by the (possibly non-unique) nameWellbore.                                                            
	WellboreID                                                                                  *string                              `json:"WellboreID,omitempty"`
	// Desired weight to be placed above the Jar - to optimize its efficiency                                                        
	WtAboveJar                                                                                  *float64                             `json:"WtAboveJar,omitempty"`
	// Desired weight to be placed blow the Jar - to optimize its efficiency                                                         
	WtBelowJar                                                                                  *float64                             `json:"WtBelowJar,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
