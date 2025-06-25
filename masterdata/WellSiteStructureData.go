package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// The schema fragment included by facilities. A facility is a grouping of equipment that is
// located within a specific geographic boundary or site and that is used in the context of
// energy-related activities such as exploration, extraction, generation, storage,
// processing, disposal, supply, or transfer. Clarifications: (1) A facility may be surface
// or subsurface located. (2) Usually facility equipment is commonly owned or operated. (3)
// Industry definitions may vary and differ from this one. This schema fragment is included
// by Well, Wellbore, Rig, as well as Tank Batteries, Compression Stations, Storage
// Facilities, Wind Farms, Wind Turbines, Mining Facilities, etc., once these types are
// included in to the OSDU.
type WellSiteStructureData struct {
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
	// The current operator organization ID; the organization ID may also be found in the                                       
	// FacilityOperatorOrganisationID of the FacilityOperator array providing the actual dates.                                 
	CurrentOperatorID                                                                           *string                         `json:"CurrentOperatorID,omitempty"`
	// The main source of the header information.                                                                               
	DataSourceOrganisationID                                                                    *string                         `json:"DataSourceOrganisationID,omitempty"`
	// A descriptive text or remark about the Facility.                                                                         
	FacilityDescription                                                                         *string                         `json:"FacilityDescription,omitempty"`
	// A list of key facility events.                                                                                           
	FacilityEvents                                                                              []AbstractFacilityEvent         `json:"FacilityEvents,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                                  
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                                 
	// If used, the "Source" property should identify that source system.                                                       
	FacilityID                                                                                  *string                         `json:"FacilityID,omitempty"`
	// Name of the Facility.                                                                                                    
	FacilityName                                                                                *string                         `json:"FacilityName,omitempty"`
	// DEPRECATED: please use data.NameAliases. Alternative names, including historical, by                                     
	// which this facility is/has been known.                                                                                   
	FacilityNameAliases                                                                         []AbstractAliasNames            `json:"FacilityNameAliases,omitempty"`
	// The history of operator organizations of the facility.                                                                   
	FacilityOperators                                                                           []AbstractFacilityOperator      `json:"FacilityOperators,omitempty"`
	// facilitySpecification maintains the specification like slot name, wellbore drilling                                      
	// permit number, rig name etc.                                                                                             
	FacilitySpecifications                                                                      []AbstractFacilitySpecification `json:"FacilitySpecifications,omitempty"`
	// The history of life cycle states the facility has been through.                                                          
	FacilityStates                                                                              []AbstractFacilityState         `json:"FacilityStates,omitempty"`
	// The definition of a kind of capability to perform a business function or a service.                                      
	FacilityTypeID                                                                              *string                         `json:"FacilityTypeID,omitempty"`
	// A initial operator organization ID; the organization ID may also be found in the                                         
	// FacilityOperatorOrganisationID of the FacilityOperator array providing the actual dates.                                 
	InitialOperatorID                                                                           *string                         `json:"InitialOperatorID,omitempty"`
	// Identifies the Facility's general location as being onshore vs. offshore.                                                
	OperatingEnvironmentID                                                                      *string                         `json:"OperatingEnvironmentID,omitempty"`
	// The default vertical coordinate reference system used in the vertical measurements for a                                 
	// well or wellbore if absent from input vertical measurements and there is no other                                        
	// recourse for obtaining a valid CRS.                                                                                      
	DefaultVerticalCRSID                                                                        *string                         `json:"DefaultVerticalCRSID,omitempty"`
	// The default datum reference point, or zero depth point, used to determine other points                                   
	// vertically in a well.  References an entry in the VerticalMeasurements array.                                            
	DefaultVerticalMeasurementID                                                                *string                         `json:"DefaultVerticalMeasurementID,omitempty"`
	// Description on how to drive to the location                                                                              
	DrivingInstructions                                                                         *string                         `json:"DrivingInstructions,omitempty"`
	// Is the WellSiteStructure origin coordinate for the Field                                                                 
	IsFieldCentre                                                                               *bool                           `json:"IsFieldCentre,omitempty"`
	// Location Description text                                                                                                
	LocationDescription                                                                         *string                         `json:"LocationDescription,omitempty"`
	// Location Uncertainty for the WellsiteStructure location coordinates                                                      
	LocationUncertainty                                                                         *AbstractLocationUncertainty    `json:"LocationUncertainty,omitempty"`
	// Free text remarks                                                                                                        
	Remarks                                                                                     []AbstractRemark                `json:"Remarks,omitempty"`
	// Diameter or Radius of the Slot                                                                                           
	SlotSize                                                                                    *float64                        `json:"SlotSize,omitempty"`
	// The type of Wellsite Structure                                                                                           
	StructureTypeID                                                                             *string                         `json:"StructureTypeID,omitempty"`
	// List of all elevations pertaining to the well site structure like, ground level/water                                    
	// depth, default rig elevation, mud line elevation, etc.                                                                   
	VerticalMeasurements                                                                        []FluffyVerticalMeasurementID   `json:"VerticalMeasurements,omitempty"`
	// List of Well Slots and their coordinates                                                                                 
	WellSlots                                                                                   []WellSlots                     `json:"WellSlots,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}          `json:"ExtensionProperties,omitempty"`
}
