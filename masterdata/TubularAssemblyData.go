package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type TubularAssemblyData struct {
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
	// Used to describe if it belongs to a RunActivity or to a PullActivity                                                          
	ActivityTypeID                                                                              *string                              `json:"ActivityTypeID,omitempty"`
	// Used to describe the reason of Activity - such as cut/pull, pulling,…                                                         
	ActivityTypeReasonDescription                                                               *string                              `json:"ActivityTypeReasonDescription,omitempty"`
	// Type of Artificial Lift supported by the TubularAssembly where applicable E.g. could be                                       
	// "Surface Pump" / "Submersible Pump" / "Gas Lift", etc                                                                         
	ArtificialLiftTypeID                                                                        *string                              `json:"ArtificialLiftTypeID,omitempty"`
	// The measured depth at the base (bottom) of the Tubular Assembly                                                               
	AssemblyBaseMD                                                                              *float64                             `json:"AssemblyBaseMD,omitempty"`
	// True Vertical Depth of the base of the Assembly                                                                               
	AssemblyBaseReportedTVD                                                                     *float64                             `json:"AssemblyBaseReportedTVD,omitempty"`
	// The measured depth at the top of the Tubular Assembly                                                                         
	AssemblyTopMD                                                                               *float64                             `json:"AssemblyTopMD,omitempty"`
	// True Vertical Depth of the top of the Assembly                                                                                
	AssemblyTopReportedTVD                                                                      *float64                             `json:"AssemblyTopReportedTVD,omitempty"`
	// Hours Circulated before running Assembly                                                                                      
	CirculatedHours                                                                             *float64                             `json:"CirculatedHours,omitempty"`
	// Tubular Running Contractor                                                                                                    
	ContractorID                                                                                *string                              `json:"ContractorID,omitempty"`
	// The drift diameter is the inside diameter (ID) that the pipe manufacturer guarantees per                                      
	// specifications. The drift diameter is used by the well planner to determine what size                                         
	// tools or casing strings can later be run through the casing, whereas the nominal inside                                       
	// diameter is used for fluid volume calculations such as mud circulating times and cement                                       
	// slurry placement calculations. Note that the nominal inside diameter is not the same as                                       
	// the drift diameter but is always slightly larger.                                                                             
	DriftDiameter                                                                               *float64                             `json:"DriftDiameter,omitempty"`
	// Estimated Weight Below Jar                                                                                                    
	EstimatedWeightBelowJar                                                                     *float64                             `json:"EstimatedWeightBelowJar,omitempty"`
	// Density of Fluid Behind Assembly                                                                                              
	FluidBehindDensity                                                                          *float64                             `json:"FluidBehindDensity,omitempty"`
	// Type of fluid behind (outside) assembly                                                                                       
	FluidBehindTypeID                                                                           *string                              `json:"FluidBehindTypeID,omitempty"`
	// Fluid Lost Volume when running assembly in hole                                                                               
	FluidLostVolume                                                                             *float64                             `json:"FluidLostVolume,omitempty"`
	// Maximum Hole Angle                                                                                                            
	HoleAngleMax                                                                                *float64                             `json:"HoleAngleMax,omitempty"`
	// Date/time Assembly started to be run in hole                                                                                  
	InHoleDate                                                                                  *time.Time                           `json:"InHoleDate,omitempty"`
	// Weight of Assembly in Slips                                                                                                   
	InSlipsWeight                                                                               *float64                             `json:"InSlipsWeight,omitempty"`
	// Is Fluid Lost                                                                                                                 
	IsFluidLost                                                                                 *bool                                `json:"IsFluidLost,omitempty"`
	// Indicates if the Tubular Assembly is currently downhole.                                                                      
	IsInstalled                                                                                 *bool                                `json:"IsInstalled,omitempty"`
	// True if the assembly is a mixed string, else false. True where the Tubular Assembly is                                        
	// made up of joints with different Diameters, Weights, Grades, Connection, Tensile                                              
	// Strengths, Collapse Resistance or  Yield Strengths.                                                                           
	IsMixedString                                                                               *bool                                `json:"IsMixedString,omitempty"`
	// Is Parallel/Dual Assembly                                                                                                     
	IsParallel                                                                                  *bool                                `json:"IsParallel,omitempty"`
	// Date/time assembly was set/landed                                                                                             
	LandedDate                                                                                  *time.Time                           `json:"LandedDate,omitempty"`
	// Depth Adjustment for use in Landed Depth calculation. Set MD = Top MD + Assembly Length +                                     
	// Depth Adjustment. Note that Depth Adjustment can be positive or negative length.                                              
	LandedDepthAdjustment                                                                       *float64                             `json:"LandedDepthAdjustment,omitempty"`
	// Weight of Assembly when Landed                                                                                                
	LandedWeight                                                                                *float64                             `json:"LandedWeight,omitempty"`
	// This reference table describes the type of liner used in the borehole where applicable.                                       
	// For example, slotted, gravel packed or pre-perforated etc.                                                                    
	LinerTypeID                                                                                 *string                              `json:"LinerTypeID,omitempty"`
	// Maximum Outer Diameter                                                                                                        
	MaximumOuterDiameter                                                                        *float64                             `json:"MaximumOuterDiameter,omitempty"`
	// Maximum Outer Diameter Measured Depth                                                                                         
	MaximumOuterDiameterMD                                                                      *float64                             `json:"MaximumOuterDiameterMD,omitempty"`
	// This is the minimum inner diameter of the whole Tubular Assembly.                                                             
	MinimumInnerDiameter                                                                        *float64                             `json:"MinimumInnerDiameter,omitempty"`
	// The name of the Tubular Assembly.                                                                                             
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// Open Hole Size/Diameter behind Assembly                                                                                       
	OpenHoleDiameter                                                                            *float64                             `json:"OpenHoleDiameter,omitempty"`
	// Optional - Identifier of the parent assembly (in case of side-track, multi-nesting,…) -                                       
	// The Concentric Tubular model is used to identify the Assembly that an Assembly sits                                           
	// inside e.g. Surface Casing set inside Conductor, Tubing set inside Production Casing, a                                       
	// Bumper Spring set inside a Production Tubing Profile Nipple, Liner set inside Casing,                                         
	// etc. This is needed to enable a Digital Well Sketch application to understand                                                 
	// relationships between Assemblies including those in parent Wellbores.                                                         
	ParentAssemblyID                                                                            *string                              `json:"ParentAssemblyID,omitempty"`
	// Identifier of the Wellbore the Assembly is installed into or run in. Note: it may not be                                      
	// same wellbore that current assembly is installed into.                                                                        
	ParentWellboreID                                                                            *string                              `json:"ParentWellboreID,omitempty"`
	// Diameter of the Pilot Hole (drillstrings)                                                                                     
	PilotHoleSize                                                                               *float64                             `json:"PilotHoleSize,omitempty"`
	// The distance that the Tubular Assembly has penetrated below the surface of the sea floor.                                     
	SeaFloorPenetrationLength                                                                   *float64                             `json:"SeaFloorPenetrationLength,omitempty"`
	// Descriptor for Assembly, e.g. Production, Surface, Conductor, Intermediate, Drilling                                          
	StringClassID                                                                               *string                              `json:"StringClassID,omitempty"`
	// The Measured Depth that the assembly is suspended from. This  'point' is typically the                                        
	// Measured Depth of the top of the assembly e.g. Hanger though with PBRs the Suspension                                         
	// Point may not be the top.                                                                                                     
	SuspensionPointMD                                                                           *float64                             `json:"SuspensionPointMD,omitempty"`
	// Tagged Measured Depth comparison to estimated Landed Depth                                                                    
	TaggedMD                                                                                    *float64                             `json:"TaggedMD,omitempty"`
	// Nominal size (diameter) describing the whole assembly, e.g. 9.625", 12.25                                                     
	TubularAssemblyNominalSize                                                                  *float64                             `json:"TubularAssemblyNominalSize,omitempty"`
	// Sequence of the TubularAssembly (Typically BHA sequence or simply BHA #)                                                      
	TubularAssemblyNumber                                                                       *int64                               `json:"TubularAssemblyNumber,omitempty"`
	// The full record of historical and current states of the Assembly. The current active                                          
	// state is recorded in TubularAssemblyStatus.                                                                                   
	TubularAssemblyStates                                                                       []TubularAssemblyStateElement        `json:"TubularAssemblyStates,omitempty"`
	// Reflects the current status of the Assembly - as 'installed', 'pulled', 'planned',... -                                       
	// Applicable to tubing/completions as opposed to drillstrings. Historical states are                                            
	// recorded in TubularAssemblyStates.                                                                                            
	TubularAssemblyStatus                                                                       *TubularAssemblyStatusClass          `json:"TubularAssemblyStatus,omitempty"`
	// Total Length of the whole assembly.                                                                                           
	TubularAssemblyTotalLength                                                                  *float64                             `json:"TubularAssemblyTotalLength,omitempty"`
	// Type of tubular assembly.                                                                                                     
	TubularAssemblyTypeID                                                                       *string                              `json:"TubularAssemblyTypeID,omitempty"`
	// Defines whether the sequence of child tubular components runs either top to bottom, or                                        
	// bottom to top.                                                                                                                
	TubularDirectionID                                                                          *string                              `json:"TubularDirectionID,omitempty"`
	// Either a self-contained vertical reference for the depths in this TubularAssembly or a                                        
	// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the                                           
	// entity defined by VerticalReferenceEntityID e.g. the parent Well.                                                             
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
