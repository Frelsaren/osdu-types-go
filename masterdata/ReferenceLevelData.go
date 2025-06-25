package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
//
// A reference level or horizontal reference surface definition, which can be used embedded
// in other schemas.
type ReferenceLevelData struct {
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
	// The date and time at which this reference level instance becomes effective.                                           
	EffectiveDateTime                                                                           *time.Time                   `json:"EffectiveDateTime,omitempty"`
	// The height above the reference surface defined by the VerticalCoordinateReferenceSystemID                             
	// positive upwards.                                                                                                     
	Height                                                                                      *float64                     `json:"Height,omitempty"`
	// The replacement velocity value used to produce vertical static shifts in seismic data.                                
	SeismicReplacementVelocity                                                                  *float64                     `json:"SeismicReplacementVelocity,omitempty"`
	// The date and time at which a reference level instance is no longer in effect.                                         
	TerminationDateTime                                                                         *time.Time                   `json:"TerminationDateTime,omitempty"`
	// The relationship to the vertical CRS defining the absolute reference surface.                                         
	VerticalCoordinateReferenceSystemID                                                         *string                      `json:"VerticalCoordinateReferenceSystemID,omitempty"`
	// When used in context of a Wellbore, this specifies Measured Depth, True Vertical Depth,                               
	// or Elevation.                                                                                                         
	VerticalMeasurementPathID                                                                   *string                      `json:"VerticalMeasurementPathID,omitempty"`
	// When used in context of a Wellbore this specifies Driller vs Logger measurements.                                     
	VerticalMeasurementSourceID                                                                 *string                      `json:"VerticalMeasurementSourceID,omitempty"`
	// Specifies the type of vertical measurement (SRD, ES, GR, MSL,and many more).                                          
	VerticalMeasurementTypeID                                                                   *string                      `json:"VerticalMeasurementTypeID,omitempty"`
	// The positional uncertainty in the vertical direction.                                                                 
	VerticalUncertainty                                                                         *float64                     `json:"VerticalUncertainty,omitempty"`
	// When used in context of a Wellbore this specifies what directional survey or wellpath was                             
	// used to calculate the TVD.                                                                                            
	WellboreTVDTrajectoryID                                                                     *string                      `json:"WellboreTVDTrajectoryID,omitempty"`
	// A description or remarks about this reference point.                                                                  
	Description                                                                                 *string                      `json:"Description,omitempty"`
	// The name of the reference point or vertical reference plane.                                                          
	Name                                                                                        *string                      `json:"Name,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
