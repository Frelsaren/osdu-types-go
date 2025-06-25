package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type BlowoutPreventerData struct {
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
	// Type of accumulator/description.                                                                                              
	AccumulatorDescription                                                                      *string                              `json:"AccumulatorDescription,omitempty"`
	// Accumulator fluid capacity.                                                                                                   
	AccumulatorFluidCapacity                                                                    *float64                             `json:"AccumulatorFluidCapacity,omitempty"`
	// Accumulator operating pressure rating.                                                                                        
	AccumulatorOperatingPressureRating                                                          *float64                             `json:"AccumulatorOperatingPressureRating,omitempty"`
	// Accumulator pre-charge pressure.                                                                                              
	AccumulatorPreChargePressure                                                                *float64                             `json:"AccumulatorPreChargePressure,omitempty"`
	// Accumulator pre-charge volume.                                                                                                
	AccumulatorPreChargeVolume                                                                  *float64                             `json:"AccumulatorPreChargeVolume,omitempty"`
	// Identifier of the type of connection to the blowout preventer.                                                                
	BlowoutPreventerConnectionTypeID                                                            *string                              `json:"BlowoutPreventerConnectionTypeID,omitempty"`
	// Inner diameter of the booster line.                                                                                           
	BoosterLineInnerDiameter                                                                    *float64                             `json:"BoosterLineInnerDiameter,omitempty"`
	// Length of the booster line along the riser.                                                                                   
	BoosterLineLength                                                                           *float64                             `json:"BoosterLineLength,omitempty"`
	// Outer diameter of the booster line.                                                                                           
	BoosterLineOuterDiameter                                                                    *float64                             `json:"BoosterLineOuterDiameter,omitempty"`
	// Inner diameter of the choke line.                                                                                             
	ChokeLineInnerDiameter                                                                      *float64                             `json:"ChokeLineInnerDiameter,omitempty"`
	// Length of the choke line along the riser.                                                                                     
	ChokeLineLength                                                                             *float64                             `json:"ChokeLineLength,omitempty"`
	// Outer diameter of the choke line.                                                                                             
	ChokeLineOuterDiameter                                                                      *float64                             `json:"ChokeLineOuterDiameter,omitempty"`
	// Choke manifold pressure.                                                                                                      
	ChokeManifoldPressure                                                                       *float64                             `json:"ChokeManifoldPressure,omitempty"`
	// Identifier of the type of choke manifold.                                                                                     
	ChokeManifoldTypeID                                                                         *string                              `json:"ChokeManifoldTypeID,omitempty"`
	// The class designation of the blowout preventer.                                                                               
	Class                                                                                       *string                              `json:"Class,omitempty"`
	// Comments and remarks related to the blowout preventer.                                                                        
	Comment                                                                                     *string                              `json:"Comment,omitempty"`
	// Size of the connection to the blowout preventer.                                                                              
	ConnectionSize                                                                              *float64                             `json:"ConnectionSize,omitempty"`
	// Identifier of the type of control manifold.                                                                                   
	ControlManifoldTypeID                                                                       *string                              `json:"ControlManifoldTypeID,omitempty"`
	// A description of the blowout preventer.                                                                                       
	Description                                                                                 *string                              `json:"Description,omitempty"`
	// Description of the diverter.                                                                                                  
	DiverterDescription                                                                         *string                              `json:"DiverterDescription,omitempty"`
	// Diameter of the diverter.                                                                                                     
	DiverterDiameter                                                                            *float64                             `json:"DiverterDiameter,omitempty"`
	// Working pressure rating of the diverter.                                                                                      
	DiverterWorkingPressureRating                                                               *float64                             `json:"DiverterWorkingPressureRating,omitempty"`
	// The height of the blowout preventer                                                                                           
	Height                                                                                      *float64                             `json:"Height,omitempty"`
	// The pressure required to operate the blowout preventer's hydraulic system, which includes                                     
	// opening and closing the preventer. It is typically lower than the maximum operating                                           
	// pressure.                                                                                                                     
	HydraulicOperatingPressure                                                                  *float64                             `json:"HydraulicOperatingPressure,omitempty"`
	// An identification tag for the blowout preventer. A serial number is a type of                                                 
	// identification tag; however, some tags contain many pieces of information.This element                                        
	// only identifies the tag and does not describe the contents.                                                                   
	IdentificationTag                                                                           *string                              `json:"IdentificationTag,omitempty"`
	// Date and time the blowout preventer was installed.                                                                            
	InstallationDateTime                                                                        *time.Time                           `json:"InstallationDateTime,omitempty"`
	// Is this a rotating blowout preventer?                                                                                         
	IsRotating                                                                                  *bool                                `json:"IsRotating,omitempty"`
	// Inner diameter of the kill line.                                                                                              
	KillLineInnerDiameter                                                                       *float64                             `json:"KillLineInnerDiameter,omitempty"`
	// Length of kill line line along the riser.                                                                                     
	KillLineLength                                                                              *float64                             `json:"KillLineLength,omitempty"`
	// Outer diameter of the kill line.                                                                                              
	KillLineOuterDiameter                                                                       *float64                             `json:"KillLineOuterDiameter,omitempty"`
	// The last data and time the blowout preventer was certified.                                                                   
	LastCertificationDateTime                                                                   *time.Time                           `json:"LastCertificationDateTime,omitempty"`
	// Unique identifier for the manufacturer of this equipment.                                                                     
	ManufacturerID                                                                              *string                              `json:"ManufacturerID,omitempty"`
	// Manufacturer's designated model.                                                                                              
	Model                                                                                       *string                              `json:"Model,omitempty"`
	// The name of the blowout preventer.                                                                                            
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// The nominal inner diameter of the blowout preventer.                                                                          
	NominalInnerDiameter                                                                        *float64                             `json:"NominalInnerDiameter,omitempty"`
	// The nominal outer diameter of the blowout preventer.                                                                          
	NominalOuterDiameter                                                                        *float64                             `json:"NominalOuterDiameter,omitempty"`
	// Maximum pressure at which the blowout preventer is expected to operate under normal                                           
	// conditions. It is typically set below the pressure rating to provide a safety margin.                                         
	OperatingPressureRating                                                                     *float64                             `json:"OperatingPressureRating,omitempty"`
	// Maximum pressure rating of the blowout preventer.                                                                             
	PressureRating                                                                              *float64                             `json:"PressureRating,omitempty"`
	// Date and time the blowout preventer was removed.                                                                              
	RemovalDateTime                                                                             *time.Time                           `json:"RemovalDateTime,omitempty"`
	// The ID of the Rig associated with the blowout preventer.                                                                      
	RigUtilizationID                                                                            *string                              `json:"RigUtilizationID,omitempty"`
	// Indicates the service standard the blowout preventer must be able to operate under.                                           
	ServiceStandard                                                                             *string                              `json:"ServiceStandard,omitempty"`
	// The measured depth where the blowout preventer was/will be set.                                                               
	SetMeasuredDepth                                                                            *float64                             `json:"SetMeasuredDepth,omitempty"`
	// Inner diameter of the surface line.                                                                                           
	SurfaceLineInnerDiameter                                                                    *float64                             `json:"SurfaceLineInnerDiameter,omitempty"`
	// Length of the surface line.                                                                                                   
	SurfaceLineLength                                                                           *float64                             `json:"SurfaceLineLength,omitempty"`
	// Outer diameter of the surface line.                                                                                           
	SurfaceLineOuterDiameter                                                                    *float64                             `json:"SurfaceLineOuterDiameter,omitempty"`
	// The vertical space required above the blowout preventer stack to allow for safe and                                           
	// efficient operation, maintenance,  and components.                                                                            
	TopClearance                                                                                *float64                             `json:"TopClearance,omitempty"`
	// Either a self-contained vertical reference for the depths in this blowout preventer or a                                      
	// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the                                           
	// entity defined by VerticalReferenceEntityID e.g. the parent Well.                                                             
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// The weight of the blowout preventer.                                                                                          
	Weight                                                                                      *float64                             `json:"Weight,omitempty"`
	// The ID of the well associated with the blowout preventer.                                                                     
	WellID                                                                                      string                               `json:"WellID"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
