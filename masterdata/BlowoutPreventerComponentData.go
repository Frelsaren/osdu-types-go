package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type BlowoutPreventerComponentData struct {
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
	// Measured depth of the the base of the component.                                                                      
	BaseMeasuredDepth                                                                           *float64                     `json:"BaseMeasuredDepth,omitempty"`
	// Identifier of the type of blowout preventer component.                                                                
	BlowoutPreventerComponentTypeID                                                             *string                      `json:"BlowoutPreventerComponentTypeID,omitempty"`
	// Identifier of the blowout preventer the component is a part of.                                                       
	BlowoutPreventerID                                                                          string                       `json:"BlowoutPreventerID"`
	// General comments and remarks related to the component.                                                                
	Comment                                                                                     *string                      `json:"Comment,omitempty"`
	// Description of the blowout preventer component.                                                                       
	Description                                                                                 *string                      `json:"Description,omitempty"`
	// The fluid volume required to close the component.                                                                     
	FluidVolumeToClose                                                                          *float64                     `json:"FluidVolumeToClose,omitempty"`
	// The fluid volume required to open the component.                                                                      
	FluidVolumeToOpen                                                                           *float64                     `json:"FluidVolumeToOpen,omitempty"`
	// Is ram bore variable?                                                                                                 
	IsVariable                                                                                  *bool                        `json:"IsVariable,omitempty"`
	// Is the blowout preventer component vertical?                                                                          
	IsVertical                                                                                  *bool                        `json:"IsVertical,omitempty"`
	// The last data and time the component was certified.                                                                   
	LastCertificationDateTime                                                                   *time.Time                   `json:"LastCertificationDateTime,omitempty"`
	// Total length of the component.                                                                                        
	Length                                                                                      *float64                     `json:"Length,omitempty"`
	// Unique identifier for the manufacturer of this equipment.                                                             
	ManufacturerID                                                                              *string                      `json:"ManufacturerID,omitempty"`
	// The allowable maximum outer diameter of pipe that can be safely closed around.                                        
	MaximumCloseDiameter                                                                        *float64                     `json:"MaximumCloseDiameter,omitempty"`
	// The maximum allowable hang-off weight supported by the component.                                                     
	MaximumHangOffWeight                                                                        *float64                     `json:"MaximumHangOffWeight,omitempty"`
	// The minimum outer diameter of pipe that can be safely closed around (only applicable for                              
	// blowout preventers with variable bore rams).                                                                          
	MinimumCloseDiameter                                                                        *float64                     `json:"MinimumCloseDiameter,omitempty"`
	// Manufacturer's designated model.                                                                                      
	Model                                                                                       *string                      `json:"Model,omitempty"`
	// The name of the blowout preventer component.                                                                          
	Name                                                                                        *string                      `json:"Name,omitempty"`
	// The nominal inner diameter of the blowout preventer component.                                                        
	NominalInnerDiameter                                                                        *float64                     `json:"NominalInnerDiameter,omitempty"`
	// The nominal outer diameter of the blowout preventer component.                                                        
	NominalOuterDiameter                                                                        *float64                     `json:"NominalOuterDiameter,omitempty"`
	// Maximum pressure at which the blowout preventer component is expected to operate under                                
	// normal conditions. It is typically set below the pressure rating to provide a safety                                  
	// margin.                                                                                                               
	OperatingPressureRating                                                                     *float64                     `json:"OperatingPressureRating,omitempty"`
	// Maximum pressure rating of the blowout preventer component.                                                           
	PressureRating                                                                              *float64                     `json:"PressureRating,omitempty"`
	// The sequence within which the components entered the hole. That is, a sequence number of                              
	// 1 entered first, 2 entered next, etc.                                                                                 
	SequenceNumber                                                                              *int64                       `json:"SequenceNumber,omitempty"`
	// Serial number of the component as provided by the manufacturer and/or the supplier.                                   
	SerialNumber                                                                                *string                      `json:"SerialNumber,omitempty"`
	// Measured depth of the the top of the component.                                                                       
	TopMeasuredDepth                                                                            *float64                     `json:"TopMeasuredDepth,omitempty"`
	// Identifier of the Well the component is associated with.                                                              
	WellID                                                                                      *string                      `json:"WellID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
