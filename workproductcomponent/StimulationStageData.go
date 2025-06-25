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
type StimulationStageData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                                         
	ExistenceKind                                                                               *string                                        `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                                  
	ResourceCurationStatus                                                                      *string                                        `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                          
	ResourceHomeRegionID                                                                        *string                                        `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                                       
	ResourceHostRegionIDs                                                                       []string                                       `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                                        
	ResourceLifecycleStatus                                                                     *string                                        `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                          
	ResourceSecurityClassification                                                              *string                                        `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                          
	// organization, agency, system, internal team, or individual. For informational purposes                                                  
	// only, the list of sources is not governed.                                                                                              
	Source                                                                                      *string                                        `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                                   
	// based on data quality. Clarifications: Since Certified is the highest classification of                                                 
	// suitable quality, any further change or versioning of a Certified record should be                                                      
	// carefully considered and justified. If a Technical Assurance value is not populated then                                                
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                
	// Technical Assurance values are not intended to be used for the identification of a single                                               
	// "preferred" or "definitive" record by comparison with other records.                                                                    
	TechnicalAssuranceID                                                                        *string                                        `json:"TechnicalAssuranceID,omitempty"`
	// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct                                               
	// from the file, in the sense certain valuable information is generated during loading                                                    
	// process (Artefact generation process). Examples include retrieving location data,                                                       
	// performing an OCR which may result in the generation of artefacts which need to be                                                      
	// preserved distinctly                                                                                                                    
	Artefacts                                                                                   []AbstractGridRepresentationArtefact           `json:"Artefacts,omitempty"`
	// The record id, which identifies this OSDU File or dataset resource.                                                                     
	Datasets                                                                                    []string                                       `json:"Datasets,omitempty"`
	// An array of references to content in Domain Data Management Services represented by this                                                
	// work-product-component. The references are formed as URI following                                                                      
	// https://www.rfc-editor.org/rfc/rfc3986#page-16. This property is exclusively populated by                                               
	// DDMSs. If a work-product-component is represented in more than one DDMS, DDMSs are                                                      
	// obliged to find the specific reference by inspecting the URI's authority values matching                                                
	// the DDMS id.                                                                                                                            
	DDMSDatasets                                                                                []string                                       `json:"DDMSDatasets,omitempty"`
	// A flag that indicates if the work product component is searchable, which means covered in                                               
	// the search index.                                                                                                                       
	IsDiscoverable                                                                              *bool                                          `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product component is undergoing an extended load.  It                                                 
	// reflects the fact that the work product component is in an early stage and may be updated                                               
	// before finalization.                                                                                                                    
	IsExtendedLoad                                                                              *bool                                          `json:"IsExtendedLoad,omitempty"`
	// Alternative names, including historical, by which this work-product-component is/has been                                               
	// known (it should include all the identifiers).                                                                                          
	NameAliases                                                                                 []AbstractAliasNames                           `json:"NameAliases,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                                                 
	// quality. Clarifications: Since Certified is the highest classification of suitable                                                      
	// quality, any further change or versioning of a Certified record should be carefully                                                     
	// considered and justified. If a Technical Assurance value is not populated then one can                                                  
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                                        
	// Technical Assurance values are not intended to be used for the identification of a single                                               
	// "preferred" or "definitive" record by comparison with other records.                                                                    
	TechnicalAssurances                                                                         []AbstractGridRepresentationTechnicalAssurance `json:"TechnicalAssurances,omitempty"`
	// Array of Authors' names of the work product component.  Could be a person or company                                                    
	// entity.                                                                                                                                 
	AuthorIDs                                                                                   []string                                       `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product component has been through                                                  
	// (ex. well planning, exploration).                                                                                                       
	BusinessActivities                                                                          []string                                       `json:"BusinessActivities,omitempty"`
	// Date that a resource (work  product component here) is formed outside of OSDU before                                                    
	// loading (e.g. publication date).                                                                                                        
	CreationDateTime                                                                            *time.Time                                     `json:"CreationDateTime,omitempty"`
	// Description.  Summary of the work product component.  Not the same as Remark which                                                      
	// captures thoughts of creator about the wpc.                                                                                             
	Description                                                                                 *string                                        `json:"Description,omitempty"`
	// List of geographic entities which provide context to the WPC.  This may include multiple                                                
	// types or multiple values of the same type.                                                                                              
	GeoContexts                                                                                 []AbstractGeoContext                           `json:"GeoContexts,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                                                    
	// product component depends.  The assertion is directed only from the asserting WPC to                                                    
	// ancestor objects, not children.  It should not be used to refer to files or artefacts                                                   
	// within the WPC -- the association within the WPC is sufficient and Artefacts are actually                                               
	// children of the main WPC file. They should be recorded in the data.Artefacts[] array.                                                   
	LineageAssertions                                                                           []LineageAssertion                             `json:"LineageAssertions,omitempty"`
	// Name                                                                                                                                    
	Name                                                                                        *string                                        `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product component                                                
	// (location of the subject matter).                                                                                                       
	SpatialArea                                                                                 *AbstractSpatialLocation                       `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product component                                                  
	// (location of the subject matter).                                                                                                       
	SpatialPoint                                                                                *AbstractSpatialLocation                       `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product component to OSDU.                                                             
	SubmitterName                                                                               *string                                        `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                                          
	Tags                                                                                        []string                                       `json:"Tags,omitempty"`
	// The average annulus pressure for all job steps for the stage treatment.                                                                 
	AnnulusPressureAvg                                                                          *float64                                       `json:"AnnulusPressureAvg,omitempty"`
	// The highest annulus pressure for all job steps while treating the stage.                                                                
	AnnulusPressureMax                                                                          *float64                                       `json:"AnnulusPressureMax,omitempty"`
	// Average base fluid return flow rate of all Steps for this Stage.                                                                        
	BaseFluidReturnRateAvg                                                                      *float64                                       `json:"BaseFluidReturnRateAvg,omitempty"`
	// Average bottomhole flow rate for this Stage.                                                                                            
	BHPumpRateAvg                                                                               *float64                                       `json:"BHPumpRateAvg,omitempty"`
	// The average measured or calculated bottom hole temperature whilst pumping with well fluid                                               
	// injection or circulation of the wellbore at the point of interest. Point of interest is                                                 
	// generally the injection point or region of interest for the test or treatment.                                                          
	BHTreatingTemperatureAvg                                                                    *float64                                       `json:"BHTreatingTemperatureAvg,omitempty"`
	// The volume pumped for the body portion of the stage treatment.                                                                          
	BodyVolume                                                                                  *float64                                       `json:"BodyVolume,omitempty"`
	// The pressure at which the formation fractures and accepts injected fluid.                                                               
	BreakDownPressure                                                                           *float64                                       `json:"BreakDownPressure,omitempty"`
	// The average casing pressure of any step for the stage treatment.                                                                        
	CasingPressureAvg                                                                           *float64                                       `json:"CasingPressureAvg,omitempty"`
	// The highest casing pressure of any step while treating the stage.                                                                       
	CasingPressureMax                                                                           *float64                                       `json:"CasingPressureMax,omitempty"`
	// Maximum casing fluid pumping rate of any step while treating the stage.                                                                 
	CasingPumpRateMax                                                                           *float64                                       `json:"CasingPumpRateMax,omitempty"`
	// Delta time recorded for the closure of the fracture to occur during the stage treatment.                                                
	ClosureDuration                                                                             *float64                                       `json:"ClosureDuration,omitempty"`
	// An analysis parameter used in hydraulic fracture design to indicate the pressure at which                                               
	// the fracture effectively closes without proppant in place.                                                                              
	ClosurePressure                                                                             *float64                                       `json:"ClosurePressure,omitempty"`
	// Average conductivity of a fracture created during the treatment supported by proppant                                                   
	// during the stimulation services. Hydraulic conductivity, symbolically represented as K,                                                 
	// is a property of vascular plants, soil or rock, that describes the ease with which water                                                
	// can move through pore spaces or fractures. It depends on the intrinsic permeability of                                                  
	// the material and on the degree of saturation. Saturated hydraulic conductivity, Ksat,                                                   
	// describes water movement through saturated media.                                                                                       
	ConductivityAvg                                                                             *float64                                       `json:"ConductivityAvg,omitempty"`
	// The average static temperature of the wellbore injection point(s) or formation at                                                       
	// equilibrium (steady state) with no fluid or tool movement, allowing for equilibrium                                                     
	// conditions at the wellbore injection point; (DHST: down hole static temperature.                                                        
	DHStaticTemperatureAvg                                                                      *float64                                       `json:"DHStaticTemperatureAvg,omitempty"`
	// Diversion details for the stimulated interval.                                                                                          
	Diversion                                                                                   *Diversion                                     `json:"Diversion,omitempty"`
	// Details about the downhole equipment design for this stimulation interval.                                                              
	FlowPath                                                                                    *FlowPath                                      `json:"FlowPath,omitempty"`
	// Average fluid pumping rate for all job Steps while treating the Stage.                                                                  
	FluidPumpRateAvg                                                                            *float64                                       `json:"FluidPumpRateAvg,omitempty"`
	// Maximum fluid pumping rate of any Step while treating the Stage.                                                                        
	FluidPumpRateMax                                                                            *float64                                       `json:"FluidPumpRateMax,omitempty"`
	// Volume pumped during flush portion of stage treatment.                                                                                  
	FlushVolume                                                                                 *float64                                       `json:"FlushVolume,omitempty"`
	// Measured depth of the bottom of the formation.                                                                                          
	FormationBottomMeasuredDepth                                                                *float64                                       `json:"FormationBottomMeasuredDepth,omitempty"`
	// True vertical depth of the bottom of the formation.                                                                                     
	FormationBottomTrueVerticalDepth                                                            *float64                                       `json:"FormationBottomTrueVerticalDepth,omitempty"`
	// The length of formation broken (fracced) per day.                                                                                       
	FormationBreakLengthPerDay                                                                  *float64                                       `json:"FormationBreakLengthPerDay,omitempty"`
	// The name of the formation being stimulated in this job stage.                                                                           
	FormationName                                                                               *string                                        `json:"FormationName,omitempty"`
	// Measured depth of the top of the formation.                                                                                             
	FormationTopMeasuredDepth                                                                   *float64                                       `json:"FormationTopMeasuredDepth,omitempty"`
	// True vertical depth of the top of the formation.                                                                                        
	FormationTopTrueVerticalDepth                                                               *float64                                       `json:"FormationTopTrueVerticalDepth,omitempty"`
	// The formation fracture gradient for the stage after treatment.                                                                          
	FractureGradientFinal                                                                       *float64                                       `json:"FractureGradientFinal,omitempty"`
	// The formation fracture gradient for stage before treatment.                                                                             
	FractureGradientInitial                                                                     *float64                                       `json:"FractureGradientInitial,omitempty"`
	// The height of the fracture  created after treating the stage.                                                                           
	FractureHeight                                                                              *float64                                       `json:"FractureHeight,omitempty"`
	// The length of the fracture created after treating the stage.                                                                            
	FractureLength                                                                              *float64                                       `json:"FractureLength,omitempty"`
	// Average fracture width created during the treatment of the stage.                                                                       
	FractureWidthAvg                                                                            *float64                                       `json:"FractureWidthAvg,omitempty"`
	// Friction pressure loss encountered by fluid transported to stage interval.                                                              
	FrictionPressure                                                                            *float64                                       `json:"FrictionPressure,omitempty"`
	// Average hydraulic horse power used for the stage.                                                                                       
	HHPAvg                                                                                      *float64                                       `json:"HHPAvg,omitempty"`
	// Maximum hydraulic horse power used for the stage.                                                                                       
	HHPMax                                                                                      *float64                                       `json:"HHPMax,omitempty"`
	// Carbon dioxide hydraulic horsepower ordered for the stage.                                                                              
	HHPOrderedCO2                                                                               *float64                                       `json:"HHPOrderedCO2,omitempty"`
	// Fluid hydraulic horsepower ordered for the stage.                                                                                       
	HHPOrderedFluid                                                                             *float64                                       `json:"HHPOrderedFluid,omitempty"`
	// Carbon dioxide hydraulic horsepower actually used for the stage.                                                                        
	HHPUsedCO2                                                                                  *float64                                       `json:"HHPUsedCO2,omitempty"`
	// Fluid hydraulic horsepower actually used for the stage.                                                                                 
	HHPUsedFluid                                                                                *float64                                       `json:"HHPUsedFluid,omitempty"`
	// The Hole Section stimulated during this job Stage.                                                                                      
	HoleSectionID                                                                               *string                                        `json:"HoleSectionID,omitempty"`
	// Measured depth at base of the stimulated interval.                                                                                      
	IntervalBaseMeasuredDepth                                                                   *float64                                       `json:"IntervalBaseMeasuredDepth,omitempty"`
	// True Vertical depth at base of interval.                                                                                                
	IntervalBaseTrueVerticalDepth                                                               *float64                                       `json:"IntervalBaseTrueVerticalDepth,omitempty"`
	// Measured depth at the top of the stimulated interval.                                                                                   
	IntervalTopMeasuredDepth                                                                    *float64                                       `json:"IntervalTopMeasuredDepth,omitempty"`
	// True Vertical depth at top of interval.                                                                                                 
	IntervalTopTrueVerticalDepth                                                                *float64                                       `json:"IntervalTopTrueVerticalDepth,omitempty"`
	// Did screen out occur? True (true or 1) indicates that screen out occurred. False (false                                                 
	// or 0) or not given indicates otherwise.                                                                                                 
	IsScreenedOut                                                                               *bool                                          `json:"IsScreenedOut,omitempty"`
	// Timed comments for this job stage of the stim job.                                                                                      
	JobEvent                                                                                    []StimEvent                                    `json:"JobEvent,omitempty"`
	// Job Step array                                                                                                                          
	JobStep                                                                                     []JobStep                                      `json:"JobStep,omitempty"`
	// Lithostratigraphic (Formation Top name) Identifier                                                                                      
	LithostratigraphicID                                                                        *string                                        `json:"LithostratigraphicID,omitempty"`
	// Material used during this job stage. For example, proppant or gel (additive).                                                           
	MaterialUsage                                                                               []AbstractStimMaterialQuantity                 `json:"MaterialUsage,omitempty"`
	// Usage and maximum mass or volume flow rates for a material for this job stage.                                                          
	MaterialUsageRateMax                                                                        []AbstractStimMaterialQuantity                 `json:"MaterialUsageRateMax,omitempty"`
	// The difference between the pressure which holds a fracture closed (minimal principal                                                    
	// stress) and that pressure which is necessary to open the fracture.                                                                      
	NetPressure                                                                                 *float64                                       `json:"NetPressure,omitempty"`
	// The diameter of the open hole.                                                                                                          
	OpenHoleDiameter                                                                            *float64                                       `json:"OpenHoleDiameter,omitempty"`
	// A name for the open hole. To be used for open hole completions.                                                                         
	OpenHoleName                                                                                *string                                        `json:"OpenHoleName,omitempty"`
	// The percentage of volume pumped used for the pad.                                                                                       
	PadPercent                                                                                  *float64                                       `json:"PadPercent,omitempty"`
	// Volume pumped for pad portion of stage treatment.                                                                                       
	PadVolume                                                                                   *float64                                       `json:"PadVolume,omitempty"`
	// Total number of perforation balls used while treating the stage.                                                                        
	PerfBallCount                                                                               *int64                                         `json:"PerfBallCount,omitempty"`
	// The size (diameter) of the perforation balls used while treating the Stage                                                              
	PerfBallSize                                                                                *float64                                       `json:"PerfBallSize,omitempty"`
	// Set of Perforation Intervals stimulated in this stage.                                                                                  
	PerforationSet                                                                              []string                                       `json:"PerforationSet,omitempty"`
	// Perforation Set Discharge Coefficient                                                                                                   
	PerforationSetDischargeCoefficient                                                          *float64                                       `json:"PerforationSetDischargeCoefficient,omitempty"`
	// Perforation Set Friction Factor                                                                                                         
	PerforationSetFrictionFactor                                                                *float64                                       `json:"PerforationSetFrictionFactor,omitempty"`
	// Perforation Set Friction Pressure                                                                                                       
	PerforationSetFrictionPressure                                                              *float64                                       `json:"PerforationSetFrictionPressure,omitempty"`
	// The average proppant concentration at the bottom of the interval.                                                                       
	ProppantBHConcAvg                                                                           *float64                                       `json:"ProppantBHConcAvg,omitempty"`
	// The maximum proppant concentration at the bottom of the stage interval for all job steps.                                               
	ProppantBHConcMax                                                                           *float64                                       `json:"ProppantBHConcMax,omitempty"`
	// The planned, total proppant mass for this job stage.                                                                                    
	ProppantDesignMass                                                                          *float64                                       `json:"ProppantDesignMass,omitempty"`
	// The proppant height.                                                                                                                    
	ProppantHeight                                                                              *float64                                       `json:"ProppantHeight,omitempty"`
	// The weight of proppant placed in the formation.                                                                                         
	ProppantInFormationMass                                                                     *float64                                       `json:"ProppantInFormationMass,omitempty"`
	// The cumulative total amount of proppant in the formation including the current stage.                                                   
	ProppantInFormationTotal                                                                    *float64                                       `json:"ProppantInFormationTotal,omitempty"`
	// The proppant concentration at the perforations.                                                                                         
	ProppantPerforationConc                                                                     *float64                                       `json:"ProppantPerforationConc,omitempty"`
	// Total proppant mass used as a percent of the design mass.                                                                               
	ProppantPumpedPercent                                                                       *float64                                       `json:"ProppantPumpedPercent,omitempty"`
	// The average proppant concentration on the surface.                                                                                      
	ProppantSurfaceConcAvg                                                                      *float64                                       `json:"ProppantSurfaceConcAvg,omitempty"`
	// The maximum proppant concentration on the surface.                                                                                      
	ProppantSurfaceConcMax                                                                      *float64                                       `json:"ProppantSurfaceConcMax,omitempty"`
	// The weight of proppant left in the wellbore after pumping has stopped.                                                                  
	ProppantWellboreMass                                                                        *float64                                       `json:"ProppantWellboreMass,omitempty"`
	// A pumping diagnostics session for this job stage.                                                                                       
	PumpDiagnosticSession                                                                       []StimJobDiagnosticSession                     `json:"PumpDiagnosticSession,omitempty"`
	// The total pumping time for the treatment of the stage.                                                                                  
	PumpTimeTotal                                                                               *float64                                       `json:"PumpTimeTotal,omitempty"`
	// General remarks about this Stage of the Stim Job.                                                                                       
	Remarks                                                                                     *string                                        `json:"Remarks,omitempty"`
	// Identifier to the Reservoir stimulated in the Job Stage                                                                                 
	ReservoirID                                                                                 *string                                        `json:"ReservoirID,omitempty"`
	// Identifier to the Reservoir Segment stimulated in the Job Stage                                                                         
	ReservoirSegmentID                                                                          *string                                        `json:"ReservoirSegmentID,omitempty"`
	// The screen out pressure.                                                                                                                
	ScreenOutPressure                                                                           *float64                                       `json:"ScreenOutPressure,omitempty"`
	// Shut in pressure data for this job stage.                                                                                               
	ShutInPressure                                                                              []StimShutInPressure                           `json:"ShutInPressure,omitempty"`
	// The initial shut-in pressure.                                                                                                           
	ShutinPressureInitial                                                                       *float64                                       `json:"ShutinPressureInitial,omitempty"`
	// The average slurry return rate of all steps for the stage treatment.                                                                    
	SlurryReturnRateAvg                                                                         *float64                                       `json:"SlurryReturnRateAvg,omitempty"`
	// The total slurry volume pumped for all steps while treating the stage.                                                                  
	SlurryVolumeTotal                                                                           *float64                                       `json:"SlurryVolumeTotal,omitempty"`
	// Ending date and time for the stage treatment.                                                                                           
	StageEndDateTime                                                                            *time.Time                                     `json:"StageEndDateTime,omitempty"`
	// Length (along hole measured depth) of the stimulated interval.                                                                          
	StageIntervalLength                                                                         *float64                                       `json:"StageIntervalLength,omitempty"`
	// The number associated with the stage.                                                                                                   
	StageNumber                                                                                 *int64                                         `json:"StageNumber,omitempty"`
	// Starting date and time for the stage treatment.                                                                                         
	StageStartDateTime                                                                          *time.Time                                     `json:"StageStartDateTime,omitempty"`
	// The SRN of the parent Stimulation Job in which the Stage was performed.                                                                 
	StimulationJobID                                                                            *string                                        `json:"StimulationJobID,omitempty"`
	// The average pressure for treating the stage across all steps.                                                                           
	SurfacePressureAvg                                                                          *float64                                       `json:"SurfacePressureAvg,omitempty"`
	// Maximum surface pressure during treatment of the stage.                                                                                 
	SurfacePressureMax                                                                          *float64                                       `json:"SurfacePressureMax,omitempty"`
	// Text describing the technology used while pumping the stage.                                                                            
	TechnologyDescription                                                                       *string                                        `json:"TechnologyDescription,omitempty"`
	// Maximum tubing fluid pumping rate of any step while treating the stage.                                                                 
	TubingFluidRateMax                                                                          *float64                                       `json:"TubingFluidRateMax,omitempty"`
	// The average tubing pressure for the stage treatment.                                                                                    
	TubingPressureAvg                                                                           *float64                                       `json:"TubingPressureAvg,omitempty"`
	// The highest tubing pressure of any step while treating the stage.                                                                       
	TubingPressureMax                                                                           *float64                                       `json:"TubingPressureMax,omitempty"`
	// Water source description for fluid pumped during stage.                                                                                 
	WaterSource                                                                                 *string                                        `json:"WaterSource,omitempty"`
	// A link to the Well Activity to which this stimulation job was performed.                                                                
	WellActivityID                                                                              *string                                        `json:"WellActivityID,omitempty"`
	// A link to the Wellbore in which the stimulation took place.                                                                             
	WellboreID                                                                                  string                                         `json:"WellboreID"`
	// ID to the Wellbore Marker within the Wellbore Marker Set that is stimulated in the Job                                                  
	// Stage                                                                                                                                   
	WellboreMarkerID                                                                            *string                                        `json:"WellboreMarkerID,omitempty"`
	// ID to the Wellbore Marker Set containing the Marker (Formation Top) that is stimulated in                                               
	// the Job Stage                                                                                                                           
	WellboreMarketSetID                                                                         *string                                        `json:"WellboreMarketSetID,omitempty"`
	// Log(s) associated to the Stimulation Stage                                                                                              
	WellLog                                                                                     []string                                       `json:"WellLog,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
