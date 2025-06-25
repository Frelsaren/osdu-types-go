package workproductcomponent

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Generic reference object containing the universal group-type properties of a Work Product
// Component for inclusion in data type specific Work Product Component objects
//
// Generic reference object containing the universal properties of a Work Product Component
// for inclusion in data type specific Work Product Component objects
type CementJobData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                                          
	ExistenceKind                                                                                *string                                        `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                                   
	ResourceCurationStatus                                                                       *string                                        `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                           
	ResourceHomeRegionID                                                                         *string                                        `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                                        
	ResourceHostRegionIDs                                                                        []string                                       `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                                         
	ResourceLifecycleStatus                                                                      *string                                        `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                           
	ResourceSecurityClassification                                                               *string                                        `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                           
	// organization, agency, system, internal team, or individual. For informational purposes                                                   
	// only, the list of sources is not governed.                                                                                               
	Source                                                                                       *string                                        `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                                    
	// based on data quality. Clarifications: Since Certified is the highest classification of                                                  
	// suitable quality, any further change or versioning of a Certified record should be                                                       
	// carefully considered and justified. If a Technical Assurance value is not populated then                                                 
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                 
	// Technical Assurance values are not intended to be used for the identification of a single                                                
	// "preferred" or "definitive" record by comparison with other records.                                                                     
	TechnicalAssuranceID                                                                         *string                                        `json:"TechnicalAssuranceID,omitempty"`
	// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct                                                
	// from the file, in the sense certain valuable information is generated during loading                                                     
	// process (Artefact generation process). Examples include retrieving location data,                                                        
	// performing an OCR which may result in the generation of artefacts which need to be                                                       
	// preserved distinctly                                                                                                                     
	Artefacts                                                                                    []AbstractGridRepresentationArtefact           `json:"Artefacts,omitempty"`
	// The record id, which identifies this OSDU File or dataset resource.                                                                      
	Datasets                                                                                     []string                                       `json:"Datasets,omitempty"`
	// An array of references to content in Domain Data Management Services represented by this                                                 
	// work-product-component. The references are formed as URI following                                                                       
	// https://www.rfc-editor.org/rfc/rfc3986#page-16. This property is exclusively populated by                                                
	// DDMSs. If a work-product-component is represented in more than one DDMS, DDMSs are                                                       
	// obliged to find the specific reference by inspecting the URI's authority values matching                                                 
	// the DDMS id.                                                                                                                             
	DDMSDatasets                                                                                 []string                                       `json:"DDMSDatasets,omitempty"`
	// A flag that indicates if the work product component is searchable, which means covered in                                                
	// the search index.                                                                                                                        
	IsDiscoverable                                                                               *bool                                          `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product component is undergoing an extended load.  It                                                  
	// reflects the fact that the work product component is in an early stage and may be updated                                                
	// before finalization.                                                                                                                     
	IsExtendedLoad                                                                               *bool                                          `json:"IsExtendedLoad,omitempty"`
	// Alternative names, including historical, by which this work-product-component is/has been                                                
	// known (it should include all the identifiers).                                                                                           
	NameAliases                                                                                  []AbstractAliasNames                           `json:"NameAliases,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                                                  
	// quality. Clarifications: Since Certified is the highest classification of suitable                                                       
	// quality, any further change or versioning of a Certified record should be carefully                                                      
	// considered and justified. If a Technical Assurance value is not populated then one can                                                   
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                         
	// Technical Assurance values are not intended to be used for the identification of a single                                                
	// "preferred" or "definitive" record by comparison with other records.                                                                     
	TechnicalAssurances                                                                          []AbstractGridRepresentationTechnicalAssurance `json:"TechnicalAssurances,omitempty"`
	// Array of Authors' names of the work product component.  Could be a person or company                                                     
	// entity.                                                                                                                                  
	AuthorIDs                                                                                    []string                                       `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product component has been through                                                   
	// (ex. well planning, exploration).                                                                                                        
	BusinessActivities                                                                           []string                                       `json:"BusinessActivities,omitempty"`
	// Date that a resource (work  product component here) is formed outside of OSDU before                                                     
	// loading (e.g. publication date).                                                                                                         
	CreationDateTime                                                                             *time.Time                                     `json:"CreationDateTime,omitempty"`
	// Description.  Summary of the work product component.  Not the same as Remark which                                                       
	// captures thoughts of creator about the wpc.                                                                                              
	Description                                                                                  *string                                        `json:"Description,omitempty"`
	// List of geographic entities which provide context to the WPC.  This may include multiple                                                 
	// types or multiple values of the same type.                                                                                               
	GeoContexts                                                                                  []AbstractGeoContext                           `json:"GeoContexts,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                                                     
	// product component depends.  The assertion is directed only from the asserting WPC to                                                     
	// ancestor objects, not children.  It should not be used to refer to files or artefacts                                                    
	// within the WPC -- the association within the WPC is sufficient and Artefacts are actually                                                
	// children of the main WPC file. They should be recorded in the data.Artefacts[] array.                                                    
	LineageAssertions                                                                            []LineageAssertion                             `json:"LineageAssertions,omitempty"`
	// Name                                                                                                                                     
	//                                                                                                                                          
	// Name of the cement job.                                                                                                                  
	Name                                                                                         string                                         `json:"Name"`
	// A polygon boundary that reflects the locale of the content of the work product component                                                 
	// (location of the subject matter).                                                                                                        
	SpatialArea                                                                                  *AbstractSpatialLocation                       `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product component                                                   
	// (location of the subject matter).                                                                                                        
	SpatialPoint                                                                                 *AbstractSpatialLocation                       `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product component to OSDU.                                                              
	SubmitterName                                                                                *string                                        `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                                           
	Tags                                                                                         []string                                       `json:"Tags,omitempty"`
	// Casing Pressure Test                                                                                                                     
	CasingPressureTest                                                                           *CasingPressureTest                            `json:"CasingPressureTest,omitempty"`
	// Cement Fluid Line Configuration                                                                                                          
	CementFluidLineConfiguration                                                                 *string                                        `json:"CementFluidLineConfiguration,omitempty"`
	// Cement Job Rating                                                                                                                        
	CementJobRating                                                                              *string                                        `json:"CementJobRating,omitempty"`
	// Cement Plug Status History                                                                                                               
	CementPlugStates                                                                             []CementPlugStatusHistory                      `json:"CementPlugStates,omitempty"`
	// Set of stages for the job (usually 1 or 2).                                                                                              
	CementStages                                                                                 []CementStage                                  `json:"CementStages,omitempty"`
	// The identifier of the TubularComponent within the Cemented String TubularAssembly used                                                   
	// for Cementing the wellbore                                                                                                               
	CementToolTubularComponentID                                                                 *string                                        `json:"CementToolTubularComponentID,omitempty"`
	// Cement Tool Type                                                                                                                         
	CementToolTypeID                                                                             *string                                        `json:"CementToolTypeID,omitempty"`
	// Identifier of cementing contractor.                                                                                                      
	ContractorID                                                                                 *string                                        `json:"ContractorID,omitempty"`
	// Contractor Job ID                                                                                                                        
	ContractorJobID                                                                              *string                                        `json:"ContractorJobID,omitempty"`
	// The estimated Measured depth of the Top of Cement (TOC) based on volume pumped.                                                          
	EstimatedCementTopMeasuredDepth                                                              *float64                                       `json:"EstimatedCementTopMeasuredDepth,omitempty"`
	// Measured depth at bottom of hole.                                                                                                        
	HoleMeasuredDepth                                                                            *float64                                       `json:"HoleMeasuredDepth,omitempty"`
	// Hole Section Assembly installed in and cement job performed in                                                                           
	HoleSectionID                                                                                *string                                        `json:"HoleSectionID,omitempty"`
	// Coiled Tubing Used in the job (true=CTU used). Values are "true" (or "1") and "false" (or                                                
	// "0").                                                                                                                                    
	IsCoilTubing                                                                                 *bool                                          `json:"IsCoilTubing,omitempty"`
	// Is the pipe pulled wet (or dry) from cementing operation depth to surface for Plugs                                                      
	IsPipePulledWet                                                                              *bool                                          `json:"IsPipePulledWet,omitempty"`
	// Plug fully drilled out                                                                                                                   
	IsPlugDrilledOut                                                                             *bool                                          `json:"IsPlugDrilledOut,omitempty"`
	// Pipe being reciprocated.  Values are "true" (or "1") and "false" (or "0").                                                               
	IsReciprocating                                                                              *bool                                          `json:"IsReciprocating,omitempty"`
	// Were fluid circulated/returned to seabed. Values are "true" (or "1") and "false" (or "0").                                               
	IsReturnsToSeabed                                                                            *bool                                          `json:"IsReturnsToSeabed,omitempty"`
	// Is pipe rotated during job.   Values are "true" (or "1") and "false" (or "0").                                                           
	IsRotating                                                                                   *bool                                          `json:"IsRotating,omitempty"`
	// Is Viscous Pilled Used                                                                                                                   
	IsViscousPillUsed                                                                            *bool                                          `json:"IsViscousPillUsed,omitempty"`
	// Job configuration/description                                                                                                            
	JobConfiguration                                                                             *string                                        `json:"JobConfiguration,omitempty"`
	// Job End Date/time                                                                                                                        
	JobEndDatetime                                                                               *time.Time                                     `json:"JobEndDatetime,omitempty"`
	// Job Start Date/time                                                                                                                      
	JobStartDatetime                                                                             *time.Time                                     `json:"JobStartDatetime,omitempty"`
	// Type of cement job                                                                                                                       
	JobTypeID                                                                                    string                                         `json:"JobTypeID"`
	// Measured depth of the landing/float collar                                                                                               
	LandingFloatCollarMeasuredDepth                                                              *float64                                       `json:"LandingFloatCollarMeasuredDepth,omitempty"`
	// Name of lead cementer                                                                                                                    
	LeadCementerName                                                                             *string                                        `json:"LeadCementerName,omitempty"`
	// Liner Top Test                                                                                                                           
	LinerTopTest                                                                                 *LinerTopTest                                  `json:"LinerTopTest,omitempty"`
	// Log Evaluation                                                                                                                           
	LogEvaluation                                                                                []LogEvaluation                                `json:"LogEvaluation,omitempty"`
	// Name of Operator Representative/Supervisor                                                                                               
	OperatorRepresentative                                                                       *string                                        `json:"OperatorRepresentative,omitempty"`
	// Operator Representative Remarks                                                                                                          
	OperatorRepresentativeRemarks                                                                *string                                        `json:"OperatorRepresentativeRemarks,omitempty"`
	// DEPRECATED: Please use PackerStingerMeasureDepth instead as this string constant string                                                  
	// property is malformed for the purpose. Packer or Stinger Measured Depth.                                                                 
	PackerStingerMeasuredDepth                                                                   *PackerStingerMeasuredDepth                    `json:"PackerStingerMeasuredDepth,omitempty"`
	// Packer or Stinger Measured Depth                                                                                                         
	PackerStingerMeasureDepth                                                                    *float64                                       `json:"PackerStingerMeasureDepth,omitempty"`
	// ROV Measured pH at Seabed                                                                                                                
	PhROVMeasured                                                                                *float64                                       `json:"PhROVMeasured,omitempty"`
	// DEPRECATED: Use the new PipePulledRateAvg instead as this constant string property is                                                    
	// malformed for the purpose.  Rate pulled from cementing operation depth to surface                                                        
	PipePulledRate                                                                               *PipePulledRate                                `json:"PipePulledRate,omitempty"`
	// Average rate pulled from cementing operation depth to surface                                                                            
	PipePulledRateAvg                                                                            *float64                                       `json:"PipePulledRateAvg,omitempty"`
	// Association to Planned Cement Job                                                                                                        
	PlannedCementJobID                                                                           *string                                        `json:"PlannedCementJobID,omitempty"`
	// If Plug, measured depth of bottom of plug.                                                                                               
	PlugBaseMeasuredDepth                                                                        *float64                                       `json:"PlugBaseMeasuredDepth,omitempty"`
	// Cement Plug Support Base Type                                                                                                            
	PlugBaseSupportTypeID                                                                        *string                                        `json:"PlugBaseSupportTypeID,omitempty"`
	// If Plug, true vertical depth of bottom of plug.                                                                                          
	PlugBaseTrueVerticalDepth                                                                    *float64                                       `json:"PlugBaseTrueVerticalDepth,omitempty"`
	// Date when Cement Plug fully drilled out (not partial drill out).                                                                         
	PlugDrillOutDate                                                                             *time.Time                                     `json:"PlugDrillOutDate,omitempty"`
	// If Plug, average inclination of wellbore across plug depth calculated from directional                                                   
	// survey                                                                                                                                   
	PlugInclinationAverage                                                                       *float64                                       `json:"PlugInclinationAverage,omitempty"`
	// If Plug, maximum inclination of wellbore across plug depth calculated from directional                                                   
	// survey                                                                                                                                   
	PlugInclinationMaximum                                                                       *float64                                       `json:"PlugInclinationMaximum,omitempty"`
	// Plug negative test pressure                                                                                                              
	PlugNegativeTestPressure                                                                     *float64                                       `json:"PlugNegativeTestPressure,omitempty"`
	// Plug Positive Test Pressure                                                                                                              
	PlugPositiveTestPressure                                                                     *float64                                       `json:"PlugPositiveTestPressure,omitempty"`
	// Plug tagged MD                                                                                                                           
	PlugTaggedMeasuredDepth                                                                      *float64                                       `json:"PlugTaggedMeasuredDepth,omitempty"`
	// Plug tagged weight                                                                                                                       
	PlugTaggedWeight                                                                             *float64                                       `json:"PlugTaggedWeight,omitempty"`
	// If Plug, measured depth of top of plug                                                                                                   
	PlugTopMeasuredDepth                                                                         *float64                                       `json:"PlugTopMeasuredDepth,omitempty"`
	// If Plug, true vertical depth of top of plug. TVDs needed for  estimating hydrostatic                                                     
	// pressure and temperature.                                                                                                                
	PlugTopTrueVerticalDepth                                                                     *float64                                       `json:"PlugTopTrueVerticalDepth,omitempty"`
	// Cement plug type.                                                                                                                        
	PlugTypeID                                                                                   *string                                        `json:"PlugTypeID,omitempty"`
	// Measured depth of the previous casing/liner shoe.                                                                                        
	PreviousShoeMeasuredDepth                                                                    *float64                                       `json:"PreviousShoeMeasuredDepth,omitempty"`
	// Estimated True Vertical Depth of previous casing/liner shoe.                                                                             
	PreviousShoeTrueVerticalDepth                                                                *float64                                       `json:"PreviousShoeTrueVerticalDepth,omitempty"`
	// Pump Through Equipment                                                                                                                   
	PumpThroughEquipmentID                                                                       *string                                        `json:"PumpThroughEquipmentID,omitempty"`
	// Distance from shoe to hole section bottom depth (Rathole)                                                                                
	RatholeLength                                                                                *float64                                       `json:"RatholeLength,omitempty"`
	// Remarks                                                                                                                                  
	Remarks                                                                                      *string                                        `json:"Remarks,omitempty"`
	// Shoe Test                                                                                                                                
	ShoeTest                                                                                     *ShoeTest                                      `json:"ShoeTest,omitempty"`
	// Shoetrack Cement                                                                                                                         
	ShoetrackCement                                                                              *ShoetrackCement                               `json:"ShoetrackCement,omitempty"`
	// Method used to perform squeeze                                                                                                           
	SqueezeMethodID                                                                              *string                                        `json:"SqueezeMethodID,omitempty"`
	// Top Measured depth of squeeze e.g. long perforated interval. Single depth when specific                                                  
	// point squeezed (top & base are same).                                                                                                    
	SqueezeTopMeasuredDepth                                                                      *float64                                       `json:"SqueezeTopMeasuredDepth,omitempty"`
	// Type of squeeze.                                                                                                                         
	SqueezeTypeID                                                                                *string                                        `json:"SqueezeTypeID,omitempty"`
	// ID of Wellbore Opening(s) through which Cement Squeezed                                                                                  
	SqueezeWellboreOpeningID                                                                     []string                                       `json:"SqueezeWellboreOpeningID,omitempty"`
	// Measured depth of the stage collar                                                                                                       
	StageCollarMeasuredDepth                                                                     *float64                                       `json:"StageCollarMeasuredDepth,omitempty"`
	// Measured depth of cemented string casing/liner shoe.                                                                                     
	StringSetMeasuredDepth                                                                       *float64                                       `json:"StringSetMeasuredDepth,omitempty"`
	// True vertical depth of cement string shoe.                                                                                               
	StringSetTrueVerticalDepth                                                                   *float64                                       `json:"StringSetTrueVerticalDepth,omitempty"`
	// TOC Interpretation                                                                                                                       
	TOCInterpretation                                                                            *TOCInterpretation                             `json:"TOCInterpretation,omitempty"`
	// ID to the Zero Depth Point Vertical Measure elevation for depths contained in the Cement                                                 
	// Job, Stages and Pumping Schedule, depth correction used to correlate MDs to original                                                     
	// drilling rig MD. References an entry in the Vertical Measurement array for the Well                                                      
	// parented by the wellbore via WellboreID.                                                                                                 
	VerticalMeasurement                                                                          AbstractFacilityVerticalMeasurement            `json:"VerticalMeasurement"`
	// Duration from cement placement and left                                                                                                  
	// undisturbed until disturbed again.                                                                                                       
	WaitOnCementDuration                                                                         *float64                                       `json:"WaitOnCementDuration,omitempty"`
	// Identifier of the parent Well Activity in which the Cement Job was performed                                                             
	WellActivityID                                                                               string                                         `json:"WellActivityID"`
	// Reference to the parent Wellbore. Cemented TubularAssembly would be installed to same                                                    
	// Wellbore.                                                                                                                                
	WellboreID                                                                                   string                                         `json:"WellboreID"`
	// Installed Casing/Liner or other Tubular Assembly on which the Cement Job is performed.                                                   
	// Not entered for Plug jobs in open hole.                                                                                                  
	WellboreTubularID                                                                            *string                                        `json:"WellboreTubularID,omitempty"`
	// Identifier of the TubularAssembly that describes the cement work string                                                                  
	WorkStringID                                                                                 *string                                        `json:"WorkStringID,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
