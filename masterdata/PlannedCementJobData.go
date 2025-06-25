package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type PlannedCementJobData struct {
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
	// Set of stages for the job (usually 1 or 2).                                                                           
	CementStages                                                                                []CementStage                `json:"CementStages,omitempty"`
	// The identifier of the TubularComponent used for Cementing the wellbore                                                
	CementToolTubularComponentID                                                                string                       `json:"CementToolTubularComponentID"`
	// Identifier of cementing contractor.                                                                                   
	ContractorID                                                                                *string                      `json:"ContractorID,omitempty"`
	// The estimated Measured depth of the Top of Cement. Depth relative to Planned wellbore                                 
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                              
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedCementTopMeasuredDepth                                                             *float64                     `json:"EstimatedCementTopMeasuredDepth,omitempty"`
	// Estimated measured depth at bottom of hole. Depth relative to Planned wellbore ZDP.                                   
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                                   
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedHoleMeasuredDepth                                                                  *float64                     `json:"EstimatedHoleMeasuredDepth,omitempty"`
	// If Plug,  estimated measured depth of bottom of plug. Depth relative to Planned wellbore                              
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                              
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedPluggedBaseMeasuredDepth                                                           *float64                     `json:"EstimatedPluggedBaseMeasuredDepth,omitempty"`
	// If Plug, estimated measured depth of top of plug. Depth relative to Planned wellbore ZDP.                             
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                                   
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedPluggedTopMeasuredDepth                                                            *float64                     `json:"EstimatedPluggedTopMeasuredDepth,omitempty"`
	// Estimated measured depth of previous shoe. Depth relative to Planned wellbore ZDP.                                    
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                                   
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedPreviousShoeMeasuredDepth                                                          *float64                     `json:"EstimatedPreviousShoeMeasuredDepth,omitempty"`
	// Estimated True Vertical Depth of previous shoe. Depth relative to Planned wellbore ZDP.                               
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                                   
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedPreviousShoeTrueVerticalDepth                                                      *float64                     `json:"EstimatedPreviousShoeTrueVerticalDepth,omitempty"`
	// Estimated measured depth of cement string shoe. Depth relative to Planned wellbore ZDP.                               
	// Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                                   
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedStringSetMeasuredDepth                                                             *float64                     `json:"EstimatedStringSetMeasuredDepth,omitempty"`
	// Estimated True vertical depth of cement string shoe. Depth relative to Planned wellbore                               
	// ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds the depth                              
	// reference in data.VerticalMeasurement.                                                                                
	EstimatedStringSetTrueVerticalDepth                                                         *float64                     `json:"EstimatedStringSetTrueVerticalDepth,omitempty"`
	// Estimated duration for waiting on cement to set.                                                                      
	EstimatedWaitingOnCement                                                                    *string                      `json:"EstimatedWaitingOnCement,omitempty"`
	// Coiled Tubing Used (true=CTU used). Values are "true" (or "1") and "false" (or "0").                                  
	IsCoilTubing                                                                                *bool                        `json:"IsCoilTubing,omitempty"`
	// Offshore job? Values are "true" (or "1") and "false" (or "0").                                                        
	IsOffshoreJob                                                                               *bool                        `json:"IsOffshoreJob,omitempty"`
	// Pipe being reciprocated.  Values are "true" (or "1") and "false" (or "0").                                            
	IsReciprocating                                                                             *bool                        `json:"IsReciprocating,omitempty"`
	// Returns to seabed? Values are "true" (or "1") and "false" (or "0").                                                   
	IsReturnsToSeabed                                                                           *bool                        `json:"IsReturnsToSeabed,omitempty"`
	// Job configuration.                                                                                                    
	JobConfig                                                                                   *string                      `json:"JobConfig,omitempty"`
	// Type of cement job.                                                                                                   
	JobTypeID                                                                                   string                       `json:"JobTypeID"`
	// Human recognizable context for the cement job.                                                                        
	Name                                                                                        string                       `json:"Name"`
	// Name for the cemented string                                                                                          
	NameCementedString                                                                          string                       `json:"NameCementedString"`
	// Plug type.                                                                                                            
	PlugTypeID                                                                                  *string                      `json:"PlugTypeID,omitempty"`
	// Measured depth of squeeze. DDepth relative to Planned wellbore ZDP. Navigate via                                      
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                                   
	// data.VerticalMeasurement.                                                                                             
	SqueezeMeasuredDepth                                                                        *float64                     `json:"SqueezeMeasuredDepth,omitempty"`
	// Identifier of the tubular assembly actually installed or to be installed                                              
	TubularAssemblyID                                                                           string                       `json:"TubularAssemblyID"`
	// Type of squeeze.                                                                                                      
	TypeSqueeze                                                                                 *string                      `json:"TypeSqueeze,omitempty"`
	// Water depth if offshore. The distance from mean sea level to water bottom.                                            
	WaterDepth                                                                                  *float64                     `json:"WaterDepth,omitempty"`
	// Identifier of the Wellbore Architecture describing the geometry of the installed tubular.                             
	WellboreArchitectureID                                                                      string                       `json:"WellboreArchitectureID"`
	// Reference to the Wellbore                                                                                             
	WellboreID                                                                                  string                       `json:"WellboreID"`
	// Identifier of the TubularAssembly that describes the cement work string                                               
	WorkStringID                                                                                *string                      `json:"WorkStringID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
