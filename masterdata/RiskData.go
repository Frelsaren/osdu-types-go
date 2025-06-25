package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type RiskData struct {
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
	// Describes the entity that may be affected by the risk                                                                 
	AffectedPersonnel                                                                           *string                      `json:"AffectedPersonnel,omitempty"`
	// A textual description of the cause of this risk                                                                       
	Cause                                                                                       *string                      `json:"Cause,omitempty"`
	// A textual description of the consequence of this risk occurring                                                       
	Consequence                                                                                 *string                      `json:"Consequence,omitempty"`
	// Reference to the name of the category of the loss - equivalent to the                                                 
	// "ConsequenceCategory" as defined in Bow Tie.                                                                          
	ConsequenceCategoryID                                                                       *string                      `json:"ConsequenceCategoryID,omitempty"`
	// References the consequence sub-category of the risk. Possible effects arising were a risk                             
	// event to occur.                                                                                                       
	ConsequenceSubCategoryID                                                                    *string                      `json:"ConsequenceSubCategoryID,omitempty"`
	// Description of the risk                                                                                               
	Description                                                                                 *string                      `json:"Description,omitempty"`
	// Date and time that activities started.                                                                                
	EffectiveDateTime                                                                           *string                      `json:"EffectiveDateTime,omitempty"`
	// Custom string to further extend the risk categorization                                                               
	ExtendedRiskCategory                                                                        *string                      `json:"ExtendedRiskCategory,omitempty"`
	// Initial Probability Level of the Risk.                                                                                
	// Values of 1 through 5 with 1 being the lowest.                                                                        
	InitialProbability                                                                          *float64                     `json:"InitialProbability,omitempty"`
	// Score computed as of "Pre-Mitigation" which is the combination of :                                                   
	// Initial Severity * Initial Probability                                                                                
	// Values of 1 through 25 with 1 being the lowest.                                                                       
	InitialRiskScore                                                                            *float64                     `json:"InitialRiskScore,omitempty"`
	// Initial Severity Level of the Risk.                                                                                   
	// Values of 1 through 5 with 1 being the lowest.                                                                        
	InitialSeverity                                                                             *float64                     `json:"InitialSeverity,omitempty"`
	// Describes the measure or measures that may be applied to the risk to mitigate its                                     
	// consequences                                                                                                          
	Mitigations                                                                                 []MitigationElement          `json:"Mitigations,omitempty"`
	// The common or preferred name for a risk                                                                               
	Name                                                                                        *string                      `json:"Name,omitempty"`
	// Assumes PREVENTION AND MITIGATION barriers are in place to manage the risk event                                      
	NetProbability                                                                              *float64                     `json:"NetProbability,omitempty"`
	// Assumes PREVENTION AND MITIGATION barriers are in place to manage the risk event  which                               
	// is the combination of :                                                                                               
	// Net Severity * Net Probability                                                                                        
	NetRiskScore                                                                                *float64                     `json:"NetRiskScore,omitempty"`
	// Assumes PREVENTION AND MITIGATION barriers are in place to manage the risk event                                      
	NetSeverity                                                                                 *float64                     `json:"NetSeverity,omitempty"`
	// A textual description of the steps to prevent this risk                                                               
	Preventions                                                                                 []PreventionElement          `json:"Preventions,omitempty"`
	// Date a resource is formed outside of OSDU before loading                                                              
	PublicationDate                                                                             *time.Time                   `json:"PublicationDate,omitempty"`
	// Identifier of the Related Risks collection                                                                            
	RelatedRiskSetID                                                                            *string                      `json:"RelatedRiskSetID,omitempty"`
	// Probability Level of the Risk.                                                                                        
	// Values of 1 through 5 with 1 being the lowest.                                                                        
	ResidualProbability                                                                         *float64                     `json:"ResidualProbability,omitempty"`
	// Score computed as of "Post-Mitigation" which is the combination of :                                                  
	// Combination of Residual Severity * Residual Probability                                                               
	// Values of 1 through 25 with 1 being the lowest.                                                                       
	ResidualRiskScore                                                                           *float64                     `json:"ResidualRiskScore,omitempty"`
	// Severity Level of the Risk.                                                                                           
	// Values of 1 through 5 with 1 being the lowest.                                                                        
	ResidualSeverity                                                                            *float64                     `json:"ResidualSeverity,omitempty"`
	// Array of identifiers from the risk to related objects or documents (such as BHA, Mud                                  
	// design, Activity plans...)                                                                                            
	RiskAssociatedObjectIDs                                                                     []string                     `json:"RiskAssociatedObjectIDs,omitempty"`
	// General category of the described risk such as "Reservoir", "Overburden", "Life of Well",                             
	// "Drilling", "Completion" or "Opportunity"                                                                             
	RiskCategoryID                                                                              *string                      `json:"RiskCategoryID,omitempty"`
	// Describes the "discipline" that may be affected by the risk                                                           
	RiskDisciplineID                                                                            *string                      `json:"RiskDisciplineID,omitempty"`
	// end depth of the risk interval. Depth relative to Planned wellbore ZDP. Navigate via                                  
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                                   
	// data.VerticalMeasurement.                                                                                             
	RiskEndDepth                                                                                *float64                     `json:"RiskEndDepth,omitempty"`
	// Describes the Hierarchy Level the risk is applying to (such as Well, Field, Global…)                                  
	RiskHierarchyLevelID                                                                        *string                      `json:"RiskHierarchyLevelID,omitempty"`
	// Describes the responsibles (person/role/entity) for managing the risk 'by value'. This is                             
	// an alternative to the 'by reference' property data.RiskResponsiblesByReferenceIDs[].                                  
	RiskResponsibles                                                                            []AbstractContact            `json:"RiskResponsibles,omitempty"`
	// Describes the references to responsibles (person/role/entity) for managing the risk 'by                               
	// reference' to UserProfile, an alternative to the 'by value' property                                                  
	// data.RiskResponsibles[].                                                                                              
	RiskResponsiblesByReferenceIDs                                                              []string                     `json:"RiskResponsiblesByReferenceIDs,omitempty"`
	// start depth of the risk interval. Depth relative to Planned wellbore ZDP. Navigate via                                
	// WellboreID to the side-car WellPlanningWellbore, which holds the depth reference in                                   
	// data.VerticalMeasurement.                                                                                             
	RiskStartDepth                                                                              *float64                     `json:"RiskStartDepth,omitempty"`
	// Detailed category of the described risk such as "BOP", "Casing", "Cementing", "Riserless"                             
	RiskSubCategoryID                                                                           *string                      `json:"RiskSubCategoryID,omitempty"`
	// Short description of the risk                                                                                         
	Summary                                                                                     *string                      `json:"Summary,omitempty"`
	// Date and time that activities were completed.                                                                         
	TerminationDateTime                                                                         *string                      `json:"TerminationDateTime,omitempty"`
	// The type of risk.                                                                                                     
	TypeID                                                                                      *string                      `json:"TypeID,omitempty"`
	// Identifier of the planned Wellbore                                                                                    
	WellboreID                                                                                  string                       `json:"WellboreID"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
