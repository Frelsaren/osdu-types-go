package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// This abstract entity gathers all the properties universally required to describe any kind
// of reservoir - at multiple levels of granularity (e.g. Reservoir or Reservoir Segment).
type ReservoirSegmentData struct {
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
	// Indicates the current life cycle status of the reservoir in a simple way for the ease of                              
	// searching. This same status may also be found in the LifeCycleStatuses array, providing                               
	// the actual status dates.                                                                                              
	CurrentLifeCycleStatusID                                                                    *string                      `json:"CurrentLifeCycleStatusID,omitempty"`
	// The date that CO2 storage injection first began in the Reservoir or Reservoir Segment.                                
	FirstCO2StorageInjectionDate                                                                *string                      `json:"FirstCO2StorageInjectionDate,omitempty"`
	// The date that commercial production first began in the Reservoir or Reservoir Segment.                                
	FirstProductionDate                                                                         *string                      `json:"FirstProductionDate,omitempty"`
	// Indicates whether the condition of a Reservoir or Reservoir Segment is Active, meaning                                
	// there is at least one well intentionally and currently interacting with the reservoir or                              
	// reservoir segment. The absence of a value should not be assumed to mean true (Active) or                              
	// false (Inactive).                                                                                                     
	IsActiveCondition                                                                           *bool                        `json:"IsActiveCondition,omitempty"`
	// This flag indicates that this reservoir is composed of reservoir segments, and that the                               
	// characteristics of the segments should be aggregated to best describe the characteristics                             
	// of the reservoir. When not segmented, the single reservoir segment describes the                                      
	// properties of the entire reservoir. If more than one reservoir segment is defined for                                 
	// this reservoir, this flag should be set and the characteristics for the reservoir should                              
	// be calculated as summary data over all its reservoir segments.                                                        
	IsSegmented                                                                                 *bool                        `json:"IsSegmented,omitempty"`
	// Set of attributes capturing the Life Cycle Statuses of the Reservoir, a concept which is                              
	// typically chronological.                                                                                              
	LifeCycleStatuses                                                                           []LifeCycleStatuses          `json:"LifeCycleStatuses,omitempty"`
	// Name of the reservoir or reservoir segment.                                                                           
	Name                                                                                        *string                      `json:"Name,omitempty"`
	// This is the best estimate of the original hydrocarbon total pore volume of the reservoir                              
	// segment at initial conditions.                                                                                        
	OriginalHydrocarbonPoreVolume                                                               *float64                     `json:"OriginalHydrocarbonPoreVolume,omitempty"`
	// An array of remarks that provide more context for this data record.                                                   
	Remarks                                                                                     []AbstractRemark             `json:"Remarks,omitempty"`
	// The productive area is the measured or estimated total area of the reservoir unit                                     
	// considered, usually the area within the hydrocarbon-water contact.                                                    
	ReservoirUnitArea                                                                           *float64                     `json:"ReservoirUnitArea,omitempty"`
	// The true vertical depth from the Vertical CRS to the permanent reservoir unit datum used                              
	// for pressure. This is further qualified by the associated Vertical CRS property. If no                                
	// Vertical CRS is populated, assume Mean Sea Level, and therefore "TVD Subsea".                                         
	ReservoirUnitPressureDatumTVD                                                               *float64                     `json:"ReservoirUnitPressureDatumTVD,omitempty"`
	// The average true vertical depth from the Vertical CRS to the top of the reservoir unit.                               
	// This is further qualified by the associated Vertical CRS property. If no Vertical CRS is                              
	// populated, assume Mean Sea Level, and therefore "TVD Subsea".                                                         
	ReservoirUnitTopDepthTVD                                                                    *float64                     `json:"ReservoirUnitTopDepthTVD,omitempty"`
	// The Vertical Coordinate Reference System defining the origin (i.e., zero point) for the                               
	// vertical measurements on this data record (e.g. PressureDatum and TopDepth). The most                                 
	// common Vertical CRS in this context is Mean Sea Level. Populating this attribute is                                   
	// especially important when the Vertical CRS is not Mean Sea Level but rather a local                                   
	// alternative like Caspian height or a custom Vertical CRS.                                                             
	VerticalCRSID                                                                               *string                      `json:"VerticalCRSID,omitempty"`
	// If the the vertical measurements on this data record (e.g. PressureDatum and TopDepth)                                
	// were not referenced from Mean Sea Level, this attribute can capture the difference                                    
	// between the selected Vertical CRS and Mean Sea Level.                                                                 
	VerticalDatumOffsetToMeanSeaLevel                                                           *float64                     `json:"VerticalDatumOffsetToMeanSeaLevel,omitempty"`
	// If a full-fledged earth model is built, this property takes the role of SectorID to refer                             
	// to a ReservoirCompartmentInterpretation.                                                                              
	CompartmentInterpretationID                                                                 *string                      `json:"CompartmentInterpretationID,omitempty"`
	// If CompartmentInterpretationID is populated, this identifies the member of the                                        
	// data.ReservoirCompartmentUnits[]. Such a member refers to the intersecting                                            
	// RockFluidUnitInterpretation and StratigraphicUnitInterpretation instances.                                            
	CompartmentMemberIdentifier                                                                 *string                      `json:"CompartmentMemberIdentifier,omitempty"`
	// This brief description is intended to inform the consumer of this data about the                                      
	// intent/purpose/use for which this reservoir segment exists.                                                           
	IntentPurposeUseDescription                                                                 *string                      `json:"IntentPurposeUseDescription,omitempty"`
	// This flag indicates that this reservoir segment (usually a fault block) is hydraulically                              
	// isolated (i.e., not in pressure communication) from other segments in the reservoir.                                  
	IsIsolated                                                                                  *bool                        `json:"IsIsolated,omitempty"`
	// Identifier linking to the parent reservoir entity (It could be either a segment or a                                  
	// reservoir)                                                                                                            
	ParentReservoirEntityID                                                                     *string                      `json:"ParentReservoirEntityID,omitempty"`
	// DEPRECATED: Use IntentPurposeUseDescription instead. The use of the reference value list                              
	// reference-data--ReservoirSegmentType is no longer encouraged. Previously: Identifier of                               
	// the reference-data object describing the type of the reservoir segment (note: not a                                   
	// validated field). This brief description is intended to afford the consumer of this data                              
	// the intent/purpose/use for which this reservoir segment exists.                                                       
	ReservoirSegmentTypeID                                                                      *string                      `json:"ReservoirSegmentTypeID,omitempty"`
	// Identifier of the Persisted Collection in which are referred the whole set of child                                   
	// segments constituting a comprehensive sector. A Persisted Collection is only used when no                             
	// full-fledged earth model is available. See also ReservoirCompartmentInterpretationID with                             
	// CompartmentMemberIdentifier for the earth model usage.                                                                
	SectorID                                                                                    *string                      `json:"SectorID,omitempty"`
	// The reference to an earth model structural organization, which contains the list of                                   
	// relevant fault and horizon interpretations for this ReservoirSegment.                                                 
	StructuralOrganizationInterpretationID                                                      *string                      `json:"StructuralOrganizationInterpretationID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
