package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type AgreementData struct {
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
	// Unique identifier of agreement in Company contracts system of record.                                                 
	AgreementExternalID                                                                         *string                      `json:"AgreementExternalID,omitempty"`
	// Name of Company contracts system of record containing authorized version of agreement.                                
	AgreementExternalSystem                                                                     *string                      `json:"AgreementExternalSystem,omitempty"`
	// Natural unique identifier of an agreement.                                                                            
	AgreementIdentifier                                                                         *string                      `json:"AgreementIdentifier,omitempty"`
	// Familiar name of agreement.  May be a code name for highly restricted agreements.                                     
	AgreementName                                                                               *string                      `json:"AgreementName,omitempty"`
	// Reference to master agreement or other parental agreement in a hierarchy of related                                   
	// agreements.                                                                                                           
	AgreementParentID                                                                           *string                      `json:"AgreementParentID,omitempty"`
	// General purpose of agreement, such as license, purchase, trade, NDA.                                                  
	AgreementTypeID                                                                             *string                      `json:"AgreementTypeID,omitempty"`
	// A list of references to legal entities which are party to the agreement in addition to                                
	// Company.                                                                                                              
	Counterparties                                                                              []string                     `json:"Counterparties,omitempty"`
	// The Date when the agreement was put in force.                                                                         
	EffectiveDate                                                                               *time.Time                   `json:"EffectiveDate,omitempty"`
	// A list of Resources that are governed by the agreement.  Note that different terms may                                
	// apply to different resources, but that granularity is handled by the Entitlements                                     
	// Rulebase.                                                                                                             
	RestrictedResources                                                                         []RestrictedResources        `json:"RestrictedResources,omitempty"`
	// A list of obligations or allowed activities specified by the agreement that apply to                                  
	// stored resources.  These are translated into rules, which the Entitlement Rulebase                                    
	// enforces.  Each rule should reference the agreement it codifies.                                                      
	Terms                                                                                       []Terms                      `json:"Terms,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
