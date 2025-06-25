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
type WellLogData struct {
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
	Name                                                                                         *string                                        `json:"Name,omitempty"`
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
	// General method or circumstance of logging - MWD, completion, etc. This is free text                                                      
	// string. Use the Use Well Log Acquisition schema `LogRuns[].ConveyanceMethodID` property                                                  
	// where possible to capture the conveyance method.                                                                                         
	ActivityType                                                                                 *string                                        `json:"ActivityType,omitempty"`
	// DEPRECATED: Use the `SamplingStop` property. The `SamplingStop` is defined as the stop                                                   
	// value or largest value of the ReferenceCurveID, typically the largest value that                                                         
	// represents the depth or time of the logging. Informational Bottom Measured Depth of the                                                  
	// Well Log. Always populate SamplingStart and SamplingStop, which represents the real                                                      
	// sampling of the WellLog, including  non-depth sampling.                                                                                  
	BottomMeasuredDepth                                                                          *float64                                       `json:"BottomMeasuredDepth,omitempty"`
	// These are candidate reference curves, not real indices, which can be used to create                                                      
	// look-up/transformation tables. These candidate reference curves are generally not                                                        
	// populated, except in the cases where multiple reference curves are present, e.g. measured                                                
	// depth and time. Supported use cases can be found in WorkedExamples.                                                                      
	CandidateReferenceCurveIDs                                                                   []string                                       `json:"CandidateReferenceCurveIDs,omitempty"`
	// DEPRECATED: Use the Well Log Acquisition schema `Project.Contractors[]` array, or the                                                    
	// `LogRuns[].ContractorCompanyID` property if a specific log run is unique to another                                                      
	// contractor. The relationship to company who engaged the service company                                                                  
	// (ServiceCompanyID) to perform the logging.                                                                                               
	CompanyID                                                                                    *string                                        `json:"CompanyID,omitempty"`
	// DEPRECATED: Use Well Log Acquisition schema `LogRuns[].ConveyanceMethodID` property. The                                                 
	// conveyance method used to acquire the log data - if not an acquired log leave                                                            
	// empty/absent.                                                                                                                            
	ConveyanceMethodID                                                                           *string                                        `json:"ConveyanceMethodID,omitempty"`
	Curves                                                                                       []FluffyCurves                                 `json:"Curves,omitempty"`
	// DEPRECATED: Use Well Log Acquisition schema `WellboreFluidTypeID` property. Type of mud                                                  
	// at time of logging (oil, water based,...)                                                                                                
	DrillingFluidProperty                                                                        *string                                        `json:"DrillingFluidProperty,omitempty"`
	// Required for complex DLIS format files defined by having multiple Logical Files and/or                                                   
	// multiple Frames.  The Frame Identifier is a numerical attribute that represents the                                                      
	// interval spacing of the data within the frame.                                                                                           
	FrameIdentifier                                                                              *string                                        `json:"FrameIdentifier,omitempty"`
	// DEPRECATED: Use Well Log Acquisition schema `LogRuns[].LogPasses[].HoleTypeID` property.                                                 
	// This is now a reference value list. Description of the hole related type of logging -                                                    
	// POSSIBLE VALUE : OpenHole / CasedHole / CementedHole                                                                                     
	HoleTypeLogging                                                                              *string                                        `json:"HoleTypeLogging,omitempty"`
	// Boolean property indicating the sampling mode of the ReferenceCurveID. True means all                                                    
	// reference curve values are regularly spaced (see SamplingInterval); false means irregular                                                
	// or discrete sample spacing.                                                                                                              
	IsRegular                                                                                    *bool                                          `json:"IsRegular,omitempty"`
	// DEPRECATED: Use the Well Log Acquisition schema `LogRuns[].LogPasses[].PassTypeID`                                                       
	// property. This is now a reference value. Log Activity, used to describe the type of pass                                                 
	// such as Calibration Pass - Main Pass - Repeated Pass                                                                                     
	LogActivity                                                                                  *string                                        `json:"LogActivity,omitempty"`
	// DEPRECATED:  Use the Well Log Acquisition schema                                                                                         
	// `LogRun[].LogPasses[].LoggingDirectionID` reference value. Specifies whether curves were                                                 
	// collected downward or upward                                                                                                             
	LoggingDirection                                                                             *string                                        `json:"LoggingDirection,omitempty"`
	// DEPRECATED:  Use the Well Log Acquisition schema `LogRun[].GenericToolTypeIDs[]`                                                         
	// reference value list to capture the generic tool types or logging services for a                                                         
	// particular Log Run. Tool mnemonics can also be captured using the                                                                        
	// `LogRun[].ToolMnemonicIDs[]` reference value list. Logging Service - mainly a short                                                      
	// concatenation of the names of the tools                                                                                                  
	LoggingService                                                                               *string                                        `json:"LoggingService,omitempty"`
	// Required for complex DLIS format files defined by having multiple Logical Files and/or                                                   
	// multiple Frames.  The Logical File Identifier is a numerical attribute that represents                                                   
	// the collection of a series of data groups (e.g. logging passes within a single logging                                                   
	// run), and each Logical File may contain one or multiple Frames.                                                                          
	LogicalFileIdentifier                                                                        *string                                        `json:"LogicalFileIdentifier,omitempty"`
	// DEPRECATED: Use the `Remarks` property which utilizes the AbstractRemark fragment,                                                       
	// capturing more information about a remark or comment. Log remark provides contextual                                                     
	// information during the actual log object acquisition. Explains how the measurement in the                                                
	// wellbore is taken on a point in time or depth. Additional information may be included                                                    
	// such as bad weather, tool failure, etc. Usually a part of the log header, log remark                                                     
	// contains info specific for an acquisition run, specific for a given logging tool                                                         
	// (multiple measurements) and/or a specific interval. In essence, log remark represents the                                                
	// external factors and operational environment, directly or indirectly affecting the                                                       
	// measurement quality/uncertainty (dynamically over time/depth) - adding both noise and                                                    
	// bias to the measurements.                                                                                                                
	LogRemark                                                                                    *string                                        `json:"LogRemark,omitempty"`
	// DEPRECATED: Use the Well Log Acquisition schema `LogRun[]` array and the`LogRunID`. Log                                                  
	// Run - describe the run of the log - can be a number, but may be also a alphanumeric                                                      
	// description such as a version name                                                                                                       
	LogRun                                                                                       *string                                        `json:"LogRun,omitempty"`
	// An interval built from two nested values : StartDate and EndDate. It applies to the whole                                                
	// log services and may apply to composite logs as [start of the first run job] and [end of                                                 
	// the last run job]Log Service Date                                                                                                        
	LogServiceDateInterval                                                                       *LogServiceDateInterval                        `json:"LogServiceDateInterval,omitempty"`
	// DEPRECATED: Use the `source` of the individual record. OSDU Native Log Source - will be                                                  
	// updated for later releases - not to be used yet                                                                                          
	LogSource                                                                                    *string                                        `json:"LogSource,omitempty"`
	// DEPRECATED: Refer to the `version` of the individual record. Log Version                                                                 
	LogVersion                                                                                   *string                                        `json:"LogVersion,omitempty"`
	// The value used within curves to indicate there is no data over specific depth ranges.                                                    
	NullValue                                                                                    *string                                        `json:"NullValue,omitempty"`
	// DEPRECATED:  Use the Well Log Acquisition schema `LogRuns[].LogPasses[].PassTypeID`                                                      
	// property. This is now a reference value to a specific pass such as "Main Pass" or                                                        
	// "Repeated Pass", etc. Indicates if the Pass is the Main one (1) or a repeated one - and                                                  
	// it's level repetition                                                                                                                    
	PassNumber                                                                                   *int64                                         `json:"PassNumber,omitempty"`
	// The name of the curve that holds the primary index (reference) values.                                                                   
	ReferenceCurveID                                                                             *string                                        `json:"ReferenceCurveID,omitempty"`
	// A remark array for contextual information during the actual log object acquisition.                                                      
	// Explains how the measurement in the wellbore is taken on a point in time or depth.                                                       
	// Additional information may be included such as bad weather, tool failure, etc. Usually a                                                 
	// part of the log header, log remark contains info specific for an acquisition run,                                                        
	// specific for a given logging tool (multiple measurements) and/or a specific interval. In                                                 
	// essence, log remark represents the external factors and operational environment, directly                                                
	// or indirectly affecting the measurement quality/uncertainty (dynamically over time/depth)                                                
	// - adding both noise and bias to the measurements.                                                                                        
	Remarks                                                                                      []AbstractRemark                               `json:"Remarks,omitempty"`
	// The along wellbore reference value for the Well Log data (e.g. Measured Depth, True                                                      
	// Vertical Depth, One-way Travel Time, Calendar Time).                                                                                     
	SamplingDomainTypeID                                                                         *string                                        `json:"SamplingDomainTypeID,omitempty"`
	// For regularly sampled curves this property holds the sampling interval. For non regular                                                  
	// sampled data this property is not set. This property can be captured here for composite                                                  
	// log sets and within the Well Log Acquisition schema for raw data for each Log Run using                                                  
	// the `LogRun[].SamplingInterval`. The IsRegular flag indicates whether SamplingInterval is                                                
	// required.                                                                                                                                
	SamplingInterval                                                                             *float64                                       `json:"SamplingInterval,omitempty"`
	// The start value/first value of the ReferenceCurveID, typically the smallest value that                                                   
	// represents the depth or time of the logging. At Well Log level, this is designed to                                                      
	// represent the smallest sampling interval of any and all individual logging runs and                                                      
	// passes.                                                                                                                                  
	SamplingStart                                                                                *float64                                       `json:"SamplingStart,omitempty"`
	// The stop value/last value of the ReferenceCurveID, typically the largest value that                                                      
	// represents depth or time of the logging. At the Well Log level, this is designed to                                                      
	// represent the largest sampling interval of any and all individual logging runs and passes.                                               
	SamplingStop                                                                                 *float64                                       `json:"SamplingStop,omitempty"`
	// Populated only if the WellLog represents time-depth relationships or check shots. It is                                                  
	// expressed via the standard AbstractFacilityVerticalMeasurement. The following properties                                                 
	// are expected to be present: VerticalMeasurementPathID (typically elevation),                                                             
	// VerticalMeasurementTypeID as SeismicReferenceDatum, VerticalMeasurement holding the                                                      
	// offset to either the VerticalCRSID or the chained VerticalReferenceID in the parent                                                      
	// Wellbore.                                                                                                                                
	SeismicReferenceElevation                                                                    *AbstractFacilityVerticalMeasurement           `json:"SeismicReferenceElevation,omitempty"`
	// DEPRECATED: Use the Well Log Acquisition schema `Project.Contractors[]` array, or the                                                    
	// `LogRuns[].ContractorCompanyID` property if a specific log run is unique to another                                                      
	// contractor.                                                                                                                              
	// The relationship to a Service Company, typically the producer or logging contractor.                                                     
	ServiceCompanyID                                                                             *string                                        `json:"ServiceCompanyID,omitempty"`
	// DEPRECATED:  Use the Well Log Acquisition schema `LogRun[].GenericToolTypeIDs[]` to                                                      
	// capture the generic tool types used in a particular Log Run. Tool mnemonics can also be                                                  
	// captured using the `LogRun[].ToolMnemonicIDs[]` reference value list. Tool String                                                        
	// Description - a long concatenation of the tools used for logging services such as                                                        
	// GammaRay+NeutronPorosity                                                                                                                 
	ToolStringDescription                                                                        *string                                        `json:"ToolStringDescription,omitempty"`
	// DEPRECATED: Use the `SamplingStart` property. The `SamplingStart` is defined as the start                                                
	// value or first value of the ReferenceCurveID, typically the smallest value that                                                          
	// represents the depth or time of the logging. Informational Top Measured Depth of the Well                                                
	// Log. Always populate SamplingStart and SamplingStop, which represents the real sampling                                                  
	// of the WellLog, including  non-depth sampling.                                                                                           
	TopMeasuredDepth                                                                             *float64                                       `json:"TopMeasuredDepth,omitempty"`
	// The vertical measurement reference for the log curves, which defines the vertical                                                        
	// reference datum for the logged depths. Either VerticalMeasurement or                                                                     
	// VerticalMeasurementID are populated.                                                                                                     
	VerticalMeasurement                                                                          *AbstractFacilityVerticalMeasurement           `json:"VerticalMeasurement,omitempty"`
	// DEPRECATED: Use data.VerticalMeasurement.VerticalReferenceID instead. References an entry                                                
	// in the Vertical Measurement array for the Wellbore identified by WellboreID, which                                                       
	// defines the vertical reference datum for all curve measured depths. Either                                                               
	// VerticalMeasurementID or VerticalMeasurement are populated.                                                                              
	VerticalMeasurementID                                                                        *string                                        `json:"VerticalMeasurementID,omitempty"`
	// DEPRECATED: Use Well Log Acquisition schema `WellboreFluidTypeID` property. Type of mud                                                  
	// at time of logging (oil, water based,...)                                                                                                
	WellboreFluidTypeID                                                                          *string                                        `json:"WellboreFluidTypeID,omitempty"`
	// The Wellbore where the Well Log Work Product Component was recorded                                                                      
	WellboreID                                                                                   *string                                        `json:"WellboreID,omitempty"`
	// The Well Log Acquisition details object captures Information relevant to the well log                                                    
	// acquisition, such as the specific acquisition job, log runs and log passes that this                                                     
	// well log information derives from.                                                                                                       
	WellLogAcquisitionDetails                                                                    *WellLogAcquisitionDetails                     `json:"WellLogAcquisitionDetails,omitempty"`
	// Identifies the WellLogClass or business domain of this Well Log and associated set of                                                    
	// curves. Examples include, Petrophysical Logs, Drilling Logs and Production Logs.                                                         
	WellLogClassID                                                                               *string                                        `json:"WellLogClassID,omitempty"`
	// Well Log Type short Description such as Raw; Evaluated; Composite;....                                                                   
	WellLogTypeID                                                                                *string                                        `json:"WellLogTypeID,omitempty"`
	// Optional time reference for (calendar) time logs. The ISO date time string representing                                                  
	// zero time. Not to be confused with seismic travel time zero. The latter is defined by                                                    
	// SeismicReferenceDatum.                                                                                                                   
	ZeroTime                                                                                     *time.Time                                     `json:"ZeroTime,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
