package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type SampleData struct {
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
	// This attribute stores the array of OSDU record IDs for the parent samples used in                                     
	// creating this sample. Creation of this sample could be achieved through through                                       
	// extractions, sub sampling, derived sampling or recombination.                                                         
	ParentSampleIDs                                                                             []string                     `json:"ParentSampleIDs,omitempty"`
	// For a sample that has been recombined from separate samples, e.g. liquid sample and vapor                             
	// sample, this object records the specified: recombination conditions (pressure and                                     
	// temperature), recombination ratio, the saturation pressure and  target recombined sample                              
	// composition, whichever of these are appropriate for this recombination effort.                                        
	RecombinationSpecification                                                                  *RecombinationSpecification  `json:"RecombinationSpecification,omitempty"`
	// An array containing operational or quality comments about the sample.                                                 
	Remarks                                                                                     []AbstractRemark             `json:"Remarks,omitempty"`
	// This captures the acquisition parameters obtained during  the sample acquisition event                                
	// associated with this sample. Note that this attribute should only be used when                                        
	// associating the sample with an acquisition event from its original source and not for                                 
	// sub-sampling or derivative sources.                                                                                   
	SampleAcquisition                                                                           *SampleAcquisition           `json:"SampleAcquisition,omitempty"`
	// This is an OSDU Record ID referencing a document that contains instructions on how the                                
	// sample should be disposed.                                                                                            
	SampleDisposalInstructionID                                                                 *string                      `json:"SampleDisposalInstructionID,omitempty"`
	// Identifier from a Master Data Management System or other trusted source external to OSDU                              
	// - stored here in order to allow for multi-system connection and synchronization. If used,                             
	// the "Source" property in AbstractCommonResources schema should identify that source                                   
	// system. i.e. this item is optional.                                                                                   
	SampleIdentifier                                                                            *string                      `json:"SampleIdentifier,omitempty"`
	// This provides the name of the sample. If there are other names that need to be stored ,                               
	// leverage the Aliases available in the Abstract objects.                                                               
	SampleName                                                                                  *string                      `json:"SampleName,omitempty"`
	// This provides information about the type of the origin of the sample. It can be used to                               
	// determine if the sample was acquired from an original source location, result of                                      
	// recombination, subsampling or derived from some laboratory process.                                                   
	SampleOriginTypeID                                                                          *string                      `json:"SampleOriginTypeID,omitempty"`
	// This captures information about the preparation process executed after the sample                                     
	// acquisition event.                                                                                                    
	SamplePreparation                                                                           []SamplePreparation          `json:"SamplePreparation,omitempty"`
	// This captures information pertaining to the observed physical properties of the sample.                               
	SampleProperties                                                                            *AbstractSampleProperties    `json:"SampleProperties,omitempty"`
	// This is the OSDU record ID from the reference list of the type of rock or fluid sample                                
	// e.g. Fluid, Core, Cuttings, Core Slab, Core Plug, Core Chip, Slides.                                                  
	SampleTypeID                                                                                *string                      `json:"SampleTypeID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
