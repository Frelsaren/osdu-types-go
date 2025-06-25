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
type SeismicAcquisitionSurveyData struct {
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
	// The (non-overlapping) historical activity states and effective start and termination                                          
	// dates. The last state is replicated in the single LastActivityState for simpler queries.                                      
	ActivityStates                                                                              []AbstractActivityState              `json:"ActivityStates,omitempty"`
	// The relation to the ActivityTemplate carrying expected parameter definitions and default                                      
	// values.                                                                                                                       
	ActivityTemplateID                                                                          *string                              `json:"ActivityTemplateID,omitempty"`
	// The current or last state this activity transitioned to. It is a copy of the last element                                     
	// in ActivityStates[]. If there is only one state recorded, the ActivityStates[] can stay                                       
	// empty.                                                                                                                        
	LastActivityState                                                                           *AbstractActivityState               `json:"LastActivityState,omitempty"`
	// General parameter value used in one instance of activity.  Includes reference to data                                         
	// objects which are inputs and outputs of the activity.                                                                         
	Parameters                                                                                  []DefaultValueElement                `json:"Parameters,omitempty"`
	// The relationship to a parent project acting as a parent activity.                                                             
	ParentProjectID                                                                             *string                              `json:"ParentProjectID,omitempty"`
	// Acquisition approach used Conventional, Wide Azimuth, Multi Azimuth etc.                                                      
	AcquisitionTypeID                                                                           *string                              `json:"AcquisitionTypeID,omitempty"`
	// The calculated are covered by the survey. This value is calculated during the loading of                                      
	// the survey.                                                                                                                   
	AreaCalculated                                                                              *float64                             `json:"AreaCalculated,omitempty"`
	// The nominal area covered by the survey. This value is usually entered by the end user.                                        
	AreaNominal                                                                                 *float64                             `json:"AreaNominal,omitempty"`
	// DEPRECATED: Use ReceiverConfigurations[].CableCount. Number of receiver arrays (lines).                                       
	CableCount                                                                                  *int64                               `json:"CableCount,omitempty"`
	// DEPRECATED: Use ReceiverConfigurations[].CableLength. Total length of receiver array.                                         
	CableLength                                                                                 *float64                             `json:"CableLength,omitempty"`
	// DEPRECATED: Use ReceiverConfigurations[].CableSpacing. Horizontal distance between                                            
	// receiver arrays.                                                                                                              
	CableSpacingDistance                                                                        *float64                             `json:"CableSpacingDistance,omitempty"`
	// DEPRECATED: Use SourceConfigurations[].EnergySourceTypeID.Seismic Source type. E.g.:                                          
	// Airgun, Vibroseis, Dynamite, Watergun.                                                                                        
	EnergySourceTypeID                                                                          *string                              `json:"EnergySourceTypeID,omitempty"`
	// The number of times a point in the subsurface is sampled.  It measures of the redundancy                                      
	// of common midpoint seismic data.                                                                                              
	FoldCount                                                                                   *int64                               `json:"FoldCount,omitempty"`
	// Horizontal distance between source and last receiver.                                                                         
	MaxOffsetDistance                                                                           *float64                             `json:"MaxOffsetDistance,omitempty"`
	// Horizontal distance between source and first receiver.                                                                        
	MinOffsetDistance                                                                           *float64                             `json:"MinOffsetDistance,omitempty"`
	// Identifies the setting of acquisition (land, marine, transition zone).                                                        
	OperatingEnvironmentID                                                                      *string                              `json:"OperatingEnvironmentID,omitempty"`
	// If populated, this survey is part of a time-lapse survey sequence. It identifies the                                          
	// preceding SeismicAcquisitionSurvey. The first survey in the sequence has an empty or                                          
	// absent PrecedingTimeLapseSurveyID.                                                                                            
	PrecedingTimeLapseSurveyID                                                                  *string                              `json:"PrecedingTimeLapseSurveyID,omitempty"`
	// The seismic receiver configurations used for this acquisition project.                                                        
	ReceiverConfigurations                                                                      []SeismicReceiverConfiguration       `json:"ReceiverConfigurations,omitempty"`
	// Length of record at time of acquisition.                                                                                      
	RecordLength                                                                                *float64                             `json:"RecordLength,omitempty"`
	// Vertical sampling interval of data at time of acquisition.                                                                    
	SampleInterval                                                                              *float64                             `json:"SampleInterval,omitempty"`
	// Reference to the standard values for the general layout of the acquisition.  This is an                                       
	// hierarchical value.  The top value is like 2D, 3D, 4D, Borehole, Passive.  The second                                         
	// value is like NATS, WATS, Brick, Crosswell.  Nodes are separated by forward slash.                                            
	SeismicGeometryTypeID                                                                       *string                              `json:"SeismicGeometryTypeID,omitempty"`
	// Orientation of plane between source and receivers.                                                                            
	ShootingAzimuthAngle                                                                        *float64                             `json:"ShootingAzimuthAngle,omitempty"`
	// DEPRECATED: Use SourceConfigurations[].ShotpointSpacing.  Horizontal distance between                                         
	// shotpoint locations.                                                                                                          
	ShotpointIncrementDistance                                                                  *float64                             `json:"ShotpointIncrementDistance,omitempty"`
	// DEPRECATED: Use SourceConfigurations[].SourceArrayCount. Number of energy sources.                                            
	SourceArrayCount                                                                            *int64                               `json:"SourceArrayCount,omitempty"`
	// DEPRECATED: Use SourceConfigurations[].SourceArraySpacing. Distance between energy                                            
	// Sources.                                                                                                                      
	SourceArraySeparationDistance                                                               *float64                             `json:"SourceArraySeparationDistance,omitempty"`
	// The seismic source configurations used for this acquisition project.                                                          
	SourceConfigurations                                                                        []SeismicSourceConfiguration         `json:"SourceConfigurations,omitempty"`
	// The vertical measurement reference for VSP surveys, which defines the vertical reference                                      
	// datum for the measured depths.                                                                                                
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// DEPRECATED: use VesselNames in SourceConfigurations and ReceiverConfigurations. List of                                       
	// names of the seismic acquisition (source and streamer) vessels used (marine environment                                       
	// only).                                                                                                                        
	VesselNames                                                                                 []string                             `json:"VesselNames,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
