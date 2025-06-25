package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type TubularComponentData struct {
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
	// DEPRECATED: This security classification is merely decorative; the security                                                   
	// classification associated to the legal.legaltags[] is evaluated by platform services                                          
	// instead. Previously:  Classifies the security level of the resource.                                                          
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
	// Axial Load Capacity of component                                                                                              
	AxialLoadCapacity                                                                           *float64                             `json:"AxialLoadCapacity,omitempty"`
	// Bottom Connection Outer Length                                                                                                
	BotConnLength                                                                               *float64                             `json:"BotConnLength,omitempty"`
	// Bottom Connection Outer Diameter                                                                                              
	BotConnOD                                                                                   *float64                             `json:"BotConnOD,omitempty"`
	// Burst Pressure                                                                                                                
	BurstPressure                                                                               *float64                             `json:"BurstPressure,omitempty"`
	// Closed End Displacement volume/length                                                                                         
	ClosedEndDisplacement                                                                       *float64                             `json:"ClosedEndDisplacement,omitempty"`
	// Collapse Pressure                                                                                                             
	CollapsePressure                                                                            *float64                             `json:"CollapsePressure,omitempty"`
	// Top and/or Bottom Connection information                                                                                      
	Connections                                                                                 []AbstractTubularComponentConnection `json:"Connections,omitempty"`
	// Density                                                                                                                       
	Density                                                                                     *float64                             `json:"Density,omitempty"`
	// The drift diameter is the inside diameter (ID) that the pipe manufacturer guarantees per                                      
	// specifications. Note that the nominal inside diameter is not the same as the drift                                            
	// diameter but is always slightly larger. The drift diameter is used by the well planner to                                     
	// determine what size tools or casing strings can later be run through the casing, whereas                                      
	// the nominal inside diameter is used for fluid volume calculations such as mud circulating                                     
	// times and cement slurry placement calculations.                                                                               
	DriftDiameter                                                                               *float64                             `json:"DriftDiameter,omitempty"`
	// Nominal inner diameter 'ID' of the component.                                                                                 
	InnerDiameter                                                                               *float64                             `json:"InnerDiameter,omitempty"`
	// Internal Reference name/description                                                                                           
	InternalReference                                                                           *string                              `json:"InternalReference,omitempty"`
	// Is Radioactive                                                                                                                
	IsRadioActive                                                                               *bool                                `json:"IsRadioActive,omitempty"`
	// Is thread lock used when making up the pipe?                                                                                  
	IsThreadLockUsed                                                                            *bool                                `json:"IsThreadLockUsed,omitempty"`
	// Average Joint Length                                                                                                          
	JointLengthAverage                                                                          *float64                             `json:"JointLengthAverage,omitempty"`
	// Linear Capacity volume/length inside component                                                                                
	LinearCapacity                                                                              *float64                             `json:"LinearCapacity,omitempty"`
	// Actual Make Up Torque                                                                                                         
	MakeUpTorqueAct                                                                             *float64                             `json:"MakeUpTorqueAct,omitempty"`
	// Maximum Make Up Torque                                                                                                        
	MakeUpTorqueMax                                                                             *float64                             `json:"MakeUpTorqueMax,omitempty"`
	// Minimum Make Up Torque                                                                                                        
	MakeUpTorqueMin                                                                             *float64                             `json:"MakeUpTorqueMin,omitempty"`
	// Optimum Make Up Torque                                                                                                        
	MakeUpTorqueOpt                                                                             *float64                             `json:"MakeUpTorqueOpt,omitempty"`
	// Unique identifier for the manufacturer of this equipment.                                                                     
	ManufacturerID                                                                              *string                              `json:"ManufacturerID,omitempty"`
	// This is the maximum hard outer diameter of the component.                                                                     
	MaximumOuterDiameter                                                                        *float64                             `json:"MaximumOuterDiameter,omitempty"`
	// Name of the component Model as defined per the operating company                                                              
	Model                                                                                       *string                              `json:"Model,omitempty"`
	// The name of this tubular component.                                                                                           
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// Description of the Size (ID) of the Nozzle used in the Tubular Component                                                      
	Nozzles                                                                                     []Nozzle                             `json:"Nozzles,omitempty"`
	// Number of Joints per pipe section                                                                                             
	NumJoints                                                                                   *int64                               `json:"NumJoints,omitempty"`
	// Outside Coupling Length                                                                                                       
	OutsideCouplingLength                                                                       *float64                             `json:"OutsideCouplingLength,omitempty"`
	// The depth the packer equipment was set to seal the casing or tubing.                                                          
	PackerSetDepthTVD                                                                           *float64                             `json:"PackerSetDepthTVD,omitempty"`
	// Identifier of the Assembly the component is part of.                                                                          
	ParentAssemblyID                                                                            *string                              `json:"ParentAssemblyID,omitempty"`
	// Identifier of the wellbore the Component is installed/run into.                                                               
	ParentWellboreID                                                                            *string                              `json:"ParentWellboreID,omitempty"`
	// Vendor part number                                                                                                            
	PartNumber                                                                                  *string                              `json:"PartNumber,omitempty"`
	// Size/diameter of the Pilot Hole when assembly is a drillstring                                                                
	PilotHoleSize                                                                               *float64                             `json:"PilotHoleSize,omitempty"`
	// Poissons Ratio                                                                                                                
	PoissonsRatio                                                                               *float64                             `json:"PoissonsRatio,omitempty"`
	// Identifier of the Section Type.                                                                                               
	SectionTypeID                                                                               *string                              `json:"SectionTypeID,omitempty"`
	// Description of the type of Sensor(s) for the Tubular Components e.g. for MWD/LWD tools                                        
	Sensors                                                                                     []Sensor                             `json:"Sensors,omitempty"`
	// Serial Number of the component as provided by the manufacturer and/or the supplier                                            
	SerialNumber                                                                                *string                              `json:"SerialNumber,omitempty"`
	// True vertical depth of the casing/tubing shoe measured from the surface.                                                      
	ShoeDepthTVD                                                                                *float64                             `json:"ShoeDepthTVD,omitempty"`
	// Unique identifier for the supplier of this equipment.                                                                         
	SupplierID                                                                                  *string                              `json:"SupplierID,omitempty"`
	// Top Connection Outer Length                                                                                                   
	TopConnLength                                                                               *float64                             `json:"TopConnLength,omitempty"`
	// Top Connection Outer Diameter                                                                                                 
	TopConnOD                                                                                   *float64                             `json:"TopConnOD,omitempty"`
	// TFA of all Nozzles                                                                                                            
	TotalFlowArea                                                                               *float64                             `json:"TotalFlowArea,omitempty"`
	// The installed measured depth of the base of the specific component                                                            
	TubularComponentBaseMD                                                                      *float64                             `json:"TubularComponentBaseMD,omitempty"`
	// True Vertical Depth of the base of the component measured from the Wellhead                                                   
	TubularComponentBaseReportedTVD                                                             *float64                             `json:"TubularComponentBaseReportedTVD,omitempty"`
	// Identifier of the Bottom Connection Type                                                                                      
	TubularComponentBottomConnectionTypeID                                                      *string                              `json:"TubularComponentBottomConnectionTypeID,omitempty"`
	// Box / Pin configuration Identifier                                                                                            
	TubularComponentBoxPinConfigID                                                              *string                              `json:"TubularComponentBoxPinConfigID,omitempty"`
	// Total Length of the component(s)                                                                                              
	TubularComponentLength                                                                      *float64                             `json:"TubularComponentLength,omitempty"`
	// Specifies the material type constituting the component.                                                                       
	TubularComponentMaterialTypeID                                                              *string                              `json:"TubularComponentMaterialTypeID,omitempty"`
	// Nominal size (outer diameter 'OD') of the component, e.g. 9.625", 12.25"                                                      
	TubularComponentNominalSize                                                                 *float64                             `json:"TubularComponentNominalSize,omitempty"`
	// Nominal size description e.g. 8-1/2" x 9-5/8"                                                                                 
	TubularComponentNominalSizeDescription                                                      *string                              `json:"TubularComponentNominalSizeDescription,omitempty"`
	// Nominal weight of the component.                                                                                              
	TubularComponentNominalWeight                                                               *float64                             `json:"TubularComponentNominalWeight,omitempty"`
	// The sequence within which the components entered the hole. That is, a sequence number of                                      
	// 1 entered first, 2 entered 2nd, etc.                                                                                          
	TubularComponentSequence                                                                    *int64                               `json:"TubularComponentSequence,omitempty"`
	// Identifier of the Top Connection Type                                                                                         
	TubularComponentTopConnectionTypeID                                                         *string                              `json:"TubularComponentTopConnectionTypeID,omitempty"`
	// The installed measured depth of the top of the specific component                                                             
	TubularComponentTopMD                                                                       *float64                             `json:"TubularComponentTopMD,omitempty"`
	// True Vertical Depth of the top of the component measured from the Wellhead                                                    
	TubularComponentTopReportedTVD                                                              *float64                             `json:"TubularComponentTopReportedTVD,omitempty"`
	// Id of tubing grade - eg. the tensile strength of the tubing material. A system of                                             
	// classifying the material specifications for steel alloys used in the manufacture of                                           
	// tubing.                                                                                                                       
	TubularComponentTubingGradeID                                                               *string                              `json:"TubularComponentTubingGradeID,omitempty"`
	// The tensile strength of the tubing material. A system of classifying the material                                             
	// specifications for steel alloys used in the manufacture of tubing.                                                            
	TubularComponentTubingGradeStrength                                                         *float64                             `json:"TubularComponentTubingGradeStrength,omitempty"`
	// The axial load required to yield the pipe.                                                                                    
	TubularComponentTubingStrength                                                              *float64                             `json:"TubularComponentTubingStrength,omitempty"`
	// Specifies the types of components that can be used in a tubular string. These are used to                                     
	// specify the type of component and multiple components are used to define a tubular string                                     
	// (Tubular).                                                                                                                    
	TubularComponentTypeID                                                                      *string                              `json:"TubularComponentTypeID,omitempty"`
	// Vendor number or other reference identifier                                                                                   
	VendorNumber                                                                                *string                              `json:"VendorNumber,omitempty"`
	// Youngs Modulus of Elasticity                                                                                                  
	YoungsModulus                                                                               *float64                             `json:"YoungsModulus,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
