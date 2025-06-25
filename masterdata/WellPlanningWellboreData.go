package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type WellPlanningWellboreData struct {
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
	// The definitive description of the hole section associated with this wellbore                                                  
	HoleSectionID                                                                               *string                              `json:"HoleSectionID,omitempty"`
	// Name of Well Planning wellbore. Derived from the record identified by WellboreID.                                             
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// A reference to a set of one or more CasingDesigns associated to the planned Wellbore.                                         
	PlannedCasingDesigns                                                                        []string                             `json:"PlannedCasingDesigns,omitempty"`
	// A reference to a set of one or more HoleSections associated to the planned Wellbore.                                          
	PlannedHoleSections                                                                         []string                             `json:"PlannedHoleSections,omitempty"`
	// A reference to the descriptive object that holds the information about the planned                                            
	// lithology associated with the wellbore                                                                                        
	PlannedLithologyID                                                                          *string                              `json:"PlannedLithologyID,omitempty"`
	// A reference to the Well Activity Program that holds the planned programs for each phase                                       
	// of the Wellbore                                                                                                               
	PlannedWellActivityProgramID                                                                *string                              `json:"PlannedWellActivityProgramID,omitempty"`
	// A reference to the PPFGDataset that holds the information about the pore pressure                                             
	// associated with the definitive drilling program                                                                               
	PPFGDatasetID                                                                               *string                              `json:"PPFGDatasetID,omitempty"`
	// A reference to the objects that holds the information about the definitive version of the                                     
	// different survey programs associated with the wellbore                                                                        
	SurveyProgramIDs                                                                            []string                             `json:"SurveyProgramIDs,omitempty"`
	// The drill targets associated with this definitive Drilling Program                                                            
	TargetID                                                                                    *string                              `json:"TargetID,omitempty"`
	// The zero depth point (ZDP) definition for the all measured depths related to this                                             
	// WellPlanningWellbore.                                                                                                         
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// Identifier of the parent wellbore.                                                                                            
	WellboreID                                                                                  *string                              `json:"WellboreID,omitempty"`
	// The formation markers associated with the definitive Drilling Program                                                         
	WellboreMarkerSetID                                                                         *string                              `json:"WellboreMarkerSetID,omitempty"`
	// Identifier of the parent well side-car for the well planning domain.                                                          
	WellPlanningWellID                                                                          *string                              `json:"WellPlanningWellID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
