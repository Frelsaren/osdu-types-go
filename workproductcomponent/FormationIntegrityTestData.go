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
type FormationIntegrityTestData struct {
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
	// Measured depth of the Wellbore bottomhole depth at time of the test relative to the                                                     
	// Vertical Measure elevation                                                                                                              
	BottomHoleMeasuredDepth                                                                     *float64                                       `json:"BottomHoleMeasuredDepth,omitempty"`
	// True vertical depth of the Wellbore bottomhole at time of the test relative to the                                                      
	// VerticalMeasure elevation                                                                                                               
	BottomHoleTrueVerticalDepth                                                                 *float64                                       `json:"BottomHoleTrueVerticalDepth,omitempty"`
	// Cement 5 Min ShutIn Point Downhole Equivalent Mud Weight                                                                                
	Cement5MinShutInPointDownholeEMW                                                            *float64                                       `json:"Cement5MinShutInPointDownholeEMW,omitempty"`
	// Cement 5 Min ShutIn Point Downhole Pressure                                                                                             
	Cement5MinShutInPointDownholePressure                                                       *float64                                       `json:"Cement5MinShutInPointDownholePressure,omitempty"`
	// Cement Inflection Point Downhole Equivalent Mud Weight                                                                                  
	CementInflectionPointDownholeEMW                                                            *float64                                       `json:"CementInflectionPointDownholeEMW,omitempty"`
	// Cement Inflection Point Downhole Pressure                                                                                               
	CementInflectionPointDownholePressure                                                       *float64                                       `json:"CementInflectionPointDownholePressure,omitempty"`
	// Cement Max Point Downhole Equivalent Mud Weight                                                                                         
	CementMaxPointDownholeEMW                                                                   *float64                                       `json:"CementMaxPointDownholeEMW,omitempty"`
	// Cement Max Point Downhole Pressure                                                                                                      
	CementMaxPointDownholePressure                                                              *float64                                       `json:"CementMaxPointDownholePressure,omitempty"`
	// Cement Unit Pumps Off Time                                                                                                              
	CementUnitPumpsOffTime                                                                      *float64                                       `json:"CementUnitPumpsOffTime,omitempty"`
	// Wellsite team remarks                                                                                                                   
	DailyOperationsTestRemarks                                                                  *string                                        `json:"DailyOperationsTestRemarks,omitempty"`
	// Pre-test required formation strength, Limit Required.                                                                                   
	ExpectedTestGradient                                                                        *float64                                       `json:"ExpectedTestGradient,omitempty"`
	// Friction Pressure.                                                                                                                      
	FrictionPressure                                                                            *float64                                       `json:"FrictionPressure,omitempty"`
	// Wellbore Azimuth of open hole. Interpolated from Definitive Survey using MD.                                                            
	HoleAzimuth                                                                                 *float64                                       `json:"HoleAzimuth,omitempty"`
	// Hole Size                                                                                                                               
	HoleDiameter                                                                                *float64                                       `json:"HoleDiameter,omitempty"`
	// Wellbore Inclination of open hole. Interpolated from Definitive Survey using MD.                                                        
	HoleInclination                                                                             *float64                                       `json:"HoleInclination,omitempty"`
	// Association to Hole Section in which FIT/LOT is performed                                                                               
	HoleSectionID                                                                               string                                         `json:"HoleSectionID"`
	// Injected volume during test. Calculated = Pumped - Returned Volume                                                                      
	InjectedVolume                                                                              *float64                                       `json:"InjectedVolume,omitempty"`
	// Interpretation Date/time                                                                                                                
	InterpretationDate                                                                          *time.Time                                     `json:"InterpretationDate,omitempty"`
	// Interpreters Surface 5 Minute Shut In Pressure                                                                                          
	InterpretedSurface5MinShutInPressure                                                        *float64                                       `json:"InterpretedSurface5minShutInPressure,omitempty"`
	// Interpreters Surface Inflection Point Pressure                                                                                          
	InterpretedSurfaceInflectionPointPressure                                                   *float64                                       `json:"InterpretedSurfaceInflectionPointPressure,omitempty"`
	// Interpreted Surface Inflection Point Pressure Maximum                                                                                   
	InterpretedSurfacePressureMaximum                                                           *float64                                       `json:"InterpretedSurfacePressureMaximum,omitempty"`
	// Interpreter's name                                                                                                                      
	Interpreter                                                                                 *string                                        `json:"Interpreter,omitempty"`
	// Interpreters remarks                                                                                                                    
	InterpretersRemarks                                                                         *string                                        `json:"InterpretersRemarks,omitempty"`
	// Test performed while drilling hole section (not shoe test)                                                                              
	IsExtendedHole                                                                              *bool                                          `json:"IsExtendedHole,omitempty"`
	// QA/QC indicator                                                                                                                         
	IsInterpreterReviewed                                                                       *bool                                          `json:"IsInterpreterReviewed,omitempty"`
	// Is Lost Circulation Material in fluid during test                                                                                       
	IsLCMUsedInTest                                                                             *bool                                          `json:"IsLCMUsedInTest,omitempty"`
	// Was the Formation Integrity test performed with a Managed Pressure Drilling applied to                                                  
	// the well                                                                                                                                
	IsMPD                                                                                       *bool                                          `json:"IsMPD,omitempty"`
	// Is there a permeable formation in the open hole segment at time of test?                                                                
	IsPermeableFormationInOH                                                                    *bool                                          `json:"IsPermeableFormationInOH,omitempty"`
	// Does Pressure While Drilling data include inflection point?                                                                             
	IsPWDInflectionPoint                                                                        *bool                                          `json:"IsPWDInflectionPoint,omitempty"`
	// Was the Formation Integrity test performed without a Riser (offshore wells only)                                                        
	IsRiserless                                                                                 *bool                                          `json:"IsRiserless,omitempty"`
	// Was a static Mud Weight measured prior to test?                                                                                         
	IsStaticMWMeasured                                                                          *bool                                          `json:"IsStaticMWMeasured,omitempty"`
	// continuous test indicator. N = Step Test.                                                                                               
	IsTestContinuous                                                                            *bool                                          `json:"IsTestContinuous,omitempty"`
	// Lithology                                                                                                                               
	LithologyTypeID                                                                             *string                                        `json:"LithologyTypeID,omitempty"`
	// Managed Pressure Drilling Delta Pressure maximum                                                                                        
	MPDDeltaPressureMax                                                                         *float64                                       `json:"MPDDeltaPressureMax,omitempty"`
	// Managed Pressure Drilling Inflection Point Pressure                                                                                     
	MPDInflectionPointPressure                                                                  *float64                                       `json:"MPDInflectionPointPressure,omitempty"`
	// Managed Pressure Drilling Prior Back Pressure Equivalent Static Density                                                                 
	MPDPriorBackPressureESD                                                                     *float64                                       `json:"MPDPriorBackPressureESD,omitempty"`
	// Type of mud base                                                                                                                        
	MudBaseType                                                                                 *string                                        `json:"MudBaseType,omitempty"`
	// Compressibility of drilling fluid                                                                                                       
	MudCompressibility                                                                          *float64                                       `json:"MudCompressibility,omitempty"`
	// Drilling fluid plastic viscosity (PV)                                                                                                   
	MudPlasticViscosity                                                                         *float64                                       `json:"MudPlasticViscosity,omitempty"`
	// Drilling fluid yield point (YP)                                                                                                         
	MudYieldPoint                                                                               *float64                                       `json:"MudYieldPoint,omitempty"`
	// Length of Open Hole below shoe.                                                                                                         
	OpenHoleLength                                                                              *float64                                       `json:"OpenHoleLength,omitempty"`
	// Operations report where the FIT/LOT test operation is described. Populated from                                                         
	// Operations Report - used to pull Rig Name via Rig associated to                                                                         
	// Report.                                                                                                                                 
	OperationsReportID                                                                          *string                                        `json:"OperationsReportID,omitempty"`
	// Total pumped volume at end of test                                                                                                      
	PumpedVolume                                                                                *float64                                       `json:"PumpedVolume,omitempty"`
	// Pump Height.                                                                                                                            
	PumpHeight                                                                                  *float64                                       `json:"PumpHeight,omitempty"`
	// Average Pump Rate.                                                                                                                      
	PumpRateAverage                                                                             *float64                                       `json:"PumpRateAverage,omitempty"`
	// Pumps off elapsed time                                                                                                                  
	PumpsOffTime                                                                                *float64                                       `json:"PumpsOffTime,omitempty"`
	// PWD 5 Minute Shut In Equivalent Mud Weight (calculated)                                                                                 
	PWD5MinShutInEMW                                                                            *float64                                       `json:"PWD5MinShutInEMW,omitempty"`
	// PWD 5 Minute Shut In Pressure (measured)                                                                                                
	PWD5MinShutInPressure                                                                       *float64                                       `json:"PWD5MinShutInPressure,omitempty"`
	// PWD 5 Minute Shut InTime                                                                                                                
	PWD5MinShutInTime                                                                           *float64                                       `json:"PWD5MinShutInTime,omitempty"`
	// PWD Adjusted Surface 5 Minute Equivalent Mud Weight                                                                                     
	PWDAdjustedSurface5MinEMW                                                                   *float64                                       `json:"PWDAdjustedSurface5MinEMW,omitempty"`
	// PWD Adjusted Surface 5 Minute Pressure                                                                                                  
	PWDAdjustedSurface5MinPressure                                                              *float64                                       `json:"PWDAdjustedSurface5MinPressure,omitempty"`
	// PWD Adjusted Surfaced Maximum Pressure                                                                                                  
	PWDAdjustedSurfacedMaxPressure                                                              *float64                                       `json:"PWDAdjustedSurfacedMaxPressure,omitempty"`
	// PWD Adjusted Surface Inflection Point Equivalent Mud Weight                                                                             
	PWDAdjustedSurfaceInflectionPointEMW                                                        *float64                                       `json:"PWDAdjustedSurfaceInflectionPointEMW,omitempty"`
	// Surface pressure adjusted for pulsed up PWD pressure                                                                                    
	PWDAdjustedSurfaceInflectionPointPressure                                                   *float64                                       `json:"PWDAdjustedSurfaceInflectionPointPressure,omitempty"`
	// PWD Adjusted Surface Maximum EMW                                                                                                        
	PWDAdjustedSurfaceMaxEMW                                                                    *float64                                       `json:"PWDAdjustedSurfaceMaxEMW,omitempty"`
	// Pressure While Drilling  Offset Pressure from Cement Unit                                                                               
	PWDCementOffsetPressure                                                                     *float64                                       `json:"PWDCementOffsetPressure,omitempty"`
	// Pressure While Drilling Inflection Point Equivalent Mud Weight (calculated).                                                            
	PWDInflectionPointEMW                                                                       *float64                                       `json:"PWDInflectionPointEMW,omitempty"`
	// Pressure While Drilling Inflection Point (measured).                                                                                    
	PWDInflectionPointPressure                                                                  *float64                                       `json:"PWDInflectionPointPressure,omitempty"`
	// Pressure While Drilling Inflection Point Time                                                                                           
	PWDInflectionPointTime                                                                      *float64                                       `json:"PWDInflectionPointTime,omitempty"`
	// Pressure While Drilling Inflection Point volume (calculated).                                                                           
	PWDInflectionPointVolume                                                                    *float64                                       `json:"PWDInflectionPointVolume,omitempty"`
	// Maximum pressure measured from Pressure While Drilling sensor after  tool download.                                                     
	PWDMaxPressure                                                                              *float64                                       `json:"PWDMaxPressure,omitempty"`
	// Pressure While Drilling Maximum Pressure Equivalent Mud Weight.                                                                         
	PWDMaxPressureEMW                                                                           *float64                                       `json:"PWDMaxPressureEMW,omitempty"`
	// Maximum pressure pulsed from Pressure While Drilling sensor                                                                             
	PWDMaxPressurePulsed                                                                        *float64                                       `json:"PWDMaxPressurePulsed,omitempty"`
	// Time of Pressure While Drilling maximum pressure                                                                                        
	PWDMaxPressureTime                                                                          *float64                                       `json:"PWDMaxPressureTime,omitempty"`
	// Volume pumped at Pressure While Drilling maximum pressure                                                                               
	PWDMaxPressureVolume                                                                        *float64                                       `json:"PWDMaxPressureVolume,omitempty"`
	// Pressure While Drilling Sensor Measured Depth                                                                                           
	PWDSensorMeasuredDepth                                                                      *float64                                       `json:"PWDSensorMeasuredDepth,omitempty"`
	// Pressure While Drilling Sensor True Vertical Depth. Calculated from Definitive Survey                                                   
	// using sensor MD.                                                                                                                        
	PWDSensorTrueVerticalDepth                                                                  *float64                                       `json:"PWDSensorTrueVerticalDepth,omitempty"`
	// Volume returned at end of test                                                                                                          
	ReturnedVolume                                                                              *float64                                       `json:"ReturnedVolume,omitempty"`
	// Diameter of the Casing/Liner shoe set in previous hole section                                                                          
	ShoeDiameter                                                                                *float64                                       `json:"ShoeDiameter,omitempty"`
	// Measured depth of the Casing/Liner Shoe relative to Vertical Measure elevation                                                          
	ShoeMeasuredDepth                                                                           *float64                                       `json:"ShoeMeasuredDepth,omitempty"`
	// True vertical depth of the Casing/Liner Shoe relative to Vertical Measure elevation                                                     
	ShoeTrueVerticalDepth                                                                       *float64                                       `json:"ShoeTrueVerticalDepth,omitempty"`
	// Smoothing Window Pressure While Drilling elapsed time                                                                                   
	SmoothingWindowPWD                                                                          *float64                                       `json:"SmoothingWindowPWD,omitempty"`
	// Time shift between surface and PWD data                                                                                                 
	SmoothingWindowPWDShiftTime                                                                 *float64                                       `json:"SmoothingWindowPWDShiftTime,omitempty"`
	// Smoothing Window Surface elapsed time                                                                                                   
	SmoothingWindowSurface                                                                      *float64                                       `json:"SmoothingWindowSurface,omitempty"`
	// Static Downhole Mud Weight                                                                                                              
	StaticDownholeMudWeight                                                                     *float64                                       `json:"StaticDownholeMudWeight,omitempty"`
	// Source of Static Downhole Mud Weight                                                                                                    
	StaticDownholeMudWeightSource                                                               *string                                        `json:"StaticDownholeMudWeightSource,omitempty"`
	// Step test time interval                                                                                                                 
	StepIntervalTime                                                                            *float64                                       `json:"StepIntervalTime,omitempty"`
	// Step test pumped volume increment                                                                                                       
	StepIntervalVolume                                                                          *float64                                       `json:"StepIntervalVolume,omitempty"`
	// Surface Pressure at Surface 5 minute shut in time                                                                                       
	Surface5MinShutInPressure                                                                   *float64                                       `json:"Surface5MinShutInPressure,omitempty"`
	// Surface Time at 5 minute shut in                                                                                                        
	Surface5MinShutInTime                                                                       *float64                                       `json:"Surface5MinShutInTime,omitempty"`
	// Surface Calculated 5 Minute Downhole Shut In Equivalent Mud Weight                                                                      
	SurfaceCalculated5MinDHShutInEMW                                                            *float64                                       `json:"SurfaceCalculated5MinDHShutInEMW,omitempty"`
	// Surface Calculated 5 Minute Downhole Shut In Pressure                                                                                   
	SurfaceCalculated5MinDHShutInPressure                                                       *float64                                       `json:"SurfaceCalculated5MinDHShutInPressure,omitempty"`
	// Surface calculated Downhole Inflection Point Equivalent Mud Weight Maximum                                                              
	SurfaceCalculatedDHEMWMax                                                                   *float64                                       `json:"SurfaceCalculatedDHEMWMax,omitempty"`
	// Surface calculated Downhole Inflection Point Pressure Maximum                                                                           
	SurfaceCalculatedDHPressureMax                                                              *float64                                       `json:"SurfaceCalculatedDHPressureMax,omitempty"`
	// Surface calculated Downhole Inflection Point Equivalent Mud Weight                                                                      
	SurfaceCalculatedInflectionPointDHEMW                                                       *float64                                       `json:"SurfaceCalculatedInflectionPointDHEMW,omitempty"`
	// Surface calculated Downhole Inflection Point Pressure                                                                                   
	SurfaceCalculatedInflectionPointDHPressure                                                  *float64                                       `json:"SurfaceCalculatedInflectionPointDHPressure,omitempty"`
	// Pressure at Surface Inflection Point                                                                                                    
	SurfaceInflectionPointPressure                                                              *float64                                       `json:"SurfaceInflectionPointPressure,omitempty"`
	// Time at Surface Inflection Point                                                                                                        
	SurfaceInflectionPointTime                                                                  *float64                                       `json:"SurfaceInflectionPointTime,omitempty"`
	// Volume at Surface Inflection Point                                                                                                      
	SurfaceInflectionPointVolume                                                                *float64                                       `json:"SurfaceInflectionPointVolume,omitempty"`
	// Surface Initial Offset Pressure                                                                                                         
	SurfaceInitialOffsetPressure                                                                *float64                                       `json:"SurfaceInitialOffsetPressure,omitempty"`
	// Mud Weight at surface at time of test                                                                                                   
	SurfaceMudWeight                                                                            *float64                                       `json:"SurfaceMudWeight,omitempty"`
	// Maximum Surface Pressure                                                                                                                
	SurfacePressureMax                                                                          *float64                                       `json:"SurfacePressureMax,omitempty"`
	// Volume pumped at Maximum Surface Pressure Pressure                                                                                      
	SurfacePressureMaximumVolume                                                                *float64                                       `json:"SurfacePressureMaximumVolume,omitempty"`
	// Time at Maximum Surface Pressure                                                                                                        
	SurfacePressureMaxTime                                                                      *float64                                       `json:"SurfacePressureMaxTime,omitempty"`
	// Surface pressure while drilling delta time                                                                                              
	SurfacePWDDeltaTime                                                                         *float64                                       `json:"SurfacePWDDeltaTime,omitempty"`
	// Fluid system stiffness                                                                                                                  
	SystemStiffness                                                                             *float64                                       `json:"SystemStiffness,omitempty"`
	// Test Data Quality Indicator. True = Representative of formation response                                                                
	TestDataQualityIndicator                                                                    *bool                                          `json:"TestDataQualityIndicator,omitempty"`
	// Date/time well integrity test was performed.                                                                                            
	TestDateTime                                                                                *time.Time                                     `json:"TestDateTime,omitempty"`
	// Test result                                                                                                                             
	TestResult                                                                                  *string                                        `json:"TestResult,omitempty"`
	// Formation Integrity test type                                                                                                           
	TestType                                                                                    *string                                        `json:"TestType,omitempty"`
	// The well vertical measurement elevation reference for  test MD and TVD.                                                                 
	VerticalMeasurement                                                                         AbstractFacilityVerticalMeasurement            `json:"VerticalMeasurement"`
	// Time v Volume v Pressure response curve(s)                                                                                              
	VolumePressureResponses                                                                     []VolumePressureResponse                       `json:"VolumePressureResponses,omitempty"`
	// Business natural key or code of the Wellbore to which this record belongs                                                               
	WellboreID                                                                                  string                                         `json:"WellboreID"`
	// Association to Marker within Wellbore                                                                                                   
	// Marker Set (MarkerID)                                                                                                                   
	WellboreMarkerID                                                                            *string                                        `json:"WellboreMarkerID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
