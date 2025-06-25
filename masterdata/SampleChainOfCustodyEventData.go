package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type SampleChainOfCustodyEventData struct {
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
	// The pressure and temperature conditions recorded when the current sample container is                                 
	// closed for the current chain of custody event.                                                                        
	ClosingCondition                                                                            *AbstractPTCondition         `json:"ClosingCondition,omitempty"`
	// The OSDU ID of the current container used to hold the sample at the end of the chain of                               
	// custody event.                                                                                                        
	// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:                                       
	// CurrentContainerID =                                                                                                  
	// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].CurrentContainer                                              
	CurrentContainerID                                                                          *string                      `json:"CurrentContainerID,omitempty"`
	// The location where the sample was stored at the end of the chain of custody event                                     
	CurrentStorageLocation                                                                      *AbstractStorageLocation     `json:"CurrentStorageLocation,omitempty"`
	// The custodian responsible for this change of custody event                                                            
	// Reference:                                                                                                            
	// Custodian = PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].Custodian                                         
	Custodian                                                                                   *string                      `json:"Custodian,omitempty"`
	// Date for this chain of custody event                                                                                  
	// Reference:                                                                                                            
	// CustodyDate = PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].CustodyDate                                     
	CustodyDate                                                                                 *time.Time                   `json:"CustodyDate,omitempty"`
	// The physical location or organisation where this chain of custody event occurred.                                     
	// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:                                       
	// CustodyEventLocation =                                                                                                
	// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].ContainerLocation                                             
	CustodyEventLocationID                                                                      *string                      `json:"CustodyEventLocationID,omitempty"`
	// The action for this chain of custody event. Enum. See sample action.                                                  
	// Reference:                                                                                                            
	// CustodyActionID = PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].CustodyAction                               
	CustodyEventTypeID                                                                          *string                      `json:"CustodyEventTypeID,omitempty"`
	// The initial sample properties observed in the source container at the start of this chain                             
	// of custody event.                                                                                                     
	InitialSampleProperties                                                                     map[string]interface{}       `json:"InitialSampleProperties,omitempty"`
	// The difference in sample properties observed due to losses incurred while transferring                                
	// between containers during this chain of custody event.                                                                
	LostSampleProperties                                                                        map[string]interface{}       `json:"LostSampleProperties,omitempty"`
	// The name of this 'chain of custody' event.                                                                            
	Name                                                                                        *string                      `json:"Name,omitempty"`
	// The pressure and temperature conditions recorded when the previous sample container is                                
	// opened for the current chain of custody event.                                                                        
	OpeningCondition                                                                            *AbstractPTCondition         `json:"OpeningCondition,omitempty"`
	// The OSDU record ID of the previous container used to hold the sample at the start of the                              
	// chain of custody event.                                                                                               
	// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:                                       
	// PreviousContainerID =                                                                                                 
	// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].PreviousContainer                                             
	PreviousContainerID                                                                         *string                      `json:"PreviousContainerID,omitempty"`
	// The initial physical location where this sample was stored at the start of the chain of                               
	// custody event.                                                                                                        
	PreviousStorageLocation                                                                     *AbstractStorageLocation     `json:"PreviousStorageLocation,omitempty"`
	// The remaining sample properties observed in the target container at the end of this chain                             
	// of custody event.                                                                                                     
	RemainingSampleProperties                                                                   map[string]interface{}       `json:"RemainingSampleProperties,omitempty"`
	// Pertinent information about this object stored alongside other attributes of this object.                             
	// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:                                       
	// Remarks = PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].Remark                                              
	Remarks                                                                                     []AbstractRemark             `json:"Remarks,omitempty"`
	// The OSDU Record ID for the Sample.                                                                                    
	SampleID                                                                                    *string                      `json:"SampleID,omitempty"`
	// The pressure and temperature conditions recorded during the sample transfer operation                                 
	// between containers for the current chain of custody event.                                                            
	// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:                                       
	// TransferCondition.Pressure =                                                                                          
	// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].TransferPressure                                              
	// TransferCondition.Temperature =                                                                                       
	// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].TransferTemperature                                           
	TransferCondition                                                                           *AbstractPTCondition         `json:"TransferCondition,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
