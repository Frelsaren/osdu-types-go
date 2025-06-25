package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type BusinessAssociateData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                          
	ExistenceKind                                                                               *string                         `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                   
	ResourceCurationStatus                                                                      *string                         `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                           
	ResourceHomeRegionID                                                                        *string                         `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                        
	ResourceHostRegionIDs                                                                       []string                        `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                         
	ResourceLifecycleStatus                                                                     *string                         `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                           
	ResourceSecurityClassification                                                              *string                         `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                           
	// organization, agency, system, internal team, or individual. For informational purposes                                   
	// only, the list of sources is not governed.                                                                               
	Source                                                                                      *string                         `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                    
	// based on data quality. Clarifications: Since Certified is the highest classification of                                  
	// suitable quality, any further change or versioning of a Certified record should be                                       
	// carefully considered and justified. If a Technical Assurance value is not populated then                                 
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                 
	// Technical Assurance values are not intended to be used for the identification of a single                                
	// "preferred" or "definitive" record by comparison with other records.                                                     
	TechnicalAssuranceID                                                                        *string                         `json:"TechnicalAssuranceID,omitempty"`
	// List of geographic entities which provide context to the master data. This may include                                   
	// multiple types or multiple values of the same type.                                                                      
	GeoContexts                                                                                 []AbstractGeoContext            `json:"GeoContexts,omitempty"`
	// Alternative names, including historical, by which this master data is/has been known (it                                 
	// should include all the identifiers).                                                                                     
	NameAliases                                                                                 []AbstractAliasNames            `json:"NameAliases,omitempty"`
	// The spatial location information such as coordinates, CRS information (left empty when                                   
	// not appropriate).                                                                                                        
	SpatialLocation                                                                             *AbstractSpatialLocation        `json:"SpatialLocation,omitempty"`
	// Describes a record's overall suitability for general business consumption in context of                                  
	// one or more workflows/personas based on data quality and reviewer's decisions.                                           
	// Clarifications: Since Certified is the highest classification of suitable quality, any                                   
	// further change or versioning of a Certified record should be carefully considered and                                    
	// justified. If a Technical Assurance value is not populated then one can assume the data                                  
	// has not been evaluated or its quality is unknown (=Unevaluated). Technical Assurance                                     
	// values are not intended to be used for the identification of a single "preferred" or                                     
	// "definitive" record by comparison with other records.                                                                    
	TechnicalAssurances                                                                         []AbstractTechnicalAssurance    `json:"TechnicalAssurances,omitempty"`
	// DEPRECATED: (in favor of more nuanced TechnicalAssurances[] array) Describes a                                           
	// master-data record's overall suitability for general business consumption based on data                                  
	// quality. Clarifications: Since Certified is the highest classification of suitable                                       
	// quality, any further change or versioning of a Certified record should be carefully                                      
	// considered and justified. If a Technical Assurance value is not populated then one can                                   
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                         
	// Technical Assurance values are not intended to be used for the identification of a single                                
	// "preferred" or "definitive" record by comparison with other records.                                                     
	TechnicalAssuranceTypeID                                                                    *string                         `json:"TechnicalAssuranceTypeID,omitempty"`
	// This describes the reason that caused the creation of a new version of this master data.                                 
	VersionCreationReason                                                                       *string                         `json:"VersionCreationReason,omitempty"`
	// The Addresses array contains information on the address, phone numbers, primary contacts,                                
	// and location of the business associate, allowing clients to have multiple addresses. For                                 
	// example, companies that have a headquarters and various satellite offices.                                               
	Addresses                                                                                   []Address                       `json:"Addresses,omitempty"`
	// Describes the set of authorities held by a business associate to make payments, sign                                     
	// contracts etc. Considered as business context.                                                                           
	Authorities                                                                                 []OrganisationAuthority         `json:"Authorities,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                                  
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                                 
	// If used, the "Source" property should identify that source system.                                                       
	BusinessAssociateID                                                                         *string                         `json:"BusinessAssociateID,omitempty"`
	// Represents the contact information for a company. May be a phone number, fax number,                                     
	// EMail address, Web URL etc.                                                                                              
	Contacts                                                                                    []AbstractContactUserProfile    `json:"Contacts,omitempty"`
	// The current status of the Business Associate, such as Active, In Receivership, Sold,                                     
	// Merged. Main sheet                                                                                                       
	CurrentStatus                                                                               *CurrentBusinessAssociateStatus `json:"CurrentStatus,omitempty"`
	// Textual description of the nature of the organisation.                                                                   
	Description                                                                                 *string                         `json:"Description,omitempty"`
	// The date and time at which a given business associate becomes effective.                                                 
	EffectiveDate                                                                               *string                         `json:"EffectiveDate,omitempty"`
	// The name of the business associate.                                                                                      
	Name                                                                                        *string                         `json:"Name,omitempty"`
	// The relationship to an organisation.                                                                                     
	OrganisationID                                                                              *string                         `json:"OrganisationID,omitempty"`
	// Category the organisational structure fits into. Possible values - Operating Unit,                                       
	// Operating sub Unit, A Business, A Department, Government Agency, etc.                                                    
	OrganisationTypeID                                                                          *string                         `json:"OrganisationTypeID,omitempty"`
	// If populated, this is the parent business associate of the current business associate. In                                
	// case of de-mergers/splits, this relationship identified the Business Associate, which was                                
	// split up.                                                                                                                
	ParentBusinessAssociateID                                                                   *string                         `json:"ParentBusinessAssociateID,omitempty"`
	// Where the Business Associate is a person, this holds the contact information.                                            
	Person                                                                                      *AbstractContactUserProfile     `json:"Person,omitempty"`
	// The array of historical business associate status together with the time interval of                                     
	// validity and a remark.                                                                                                   
	PreviousStates                                                                              []BAStatus                      `json:"PreviousStates,omitempty"`
	// The reason why the business associated was formed.                                                                       
	PurposeDescription                                                                          *string                         `json:"PurposeDescription,omitempty"`
	// Narrative remarks about this Business Associate.                                                                         
	Remark                                                                                      *string                         `json:"Remark,omitempty"`
	// Describes the set of primary services provided by a business associate. For example                                      
	// drilling contractor, logging com pany, seismic broker etc.                                                               
	Services                                                                                    []Service                       `json:"Services,omitempty"`
	// The succeeding organisation/business associate in case of mergers and acquisitions.                                      
	SuccessorID                                                                                 *string                         `json:"SuccessorID,omitempty"`
	// The date and time at which a given business associate is no longer in effect.                                            
	TerminationDate                                                                             *string                         `json:"TerminationDate,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}          `json:"ExtensionProperties,omitempty"`
}
