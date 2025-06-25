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
type WellOperationsReportData struct {
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
	//                                                                                                                                         
	// Name of Operations Report                                                                                                               
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
	// Bit description and dull grade                                                                                                          
	BitRecord                                                                                   []BitRecord                                    `json:"BitRecord,omitempty"`
	// Hole condition description.                                                                                                             
	ConditionHole                                                                               *string                                        `json:"ConditionHole,omitempty"`
	// Cost information captured for a defined time period during operations and/or drilling                                                   
	Cost                                                                                        []Cost                                         `json:"Cost,omitempty"`
	// Daily cost.                                                                                                                             
	CostDay                                                                                     *float64                                       `json:"CostDay,omitempty"`
	// Daily mud cost.                                                                                                                         
	CostDayMud                                                                                  *float64                                       `json:"CostDayMud,omitempty"`
	// Date and time that the reporting period ended. A report period is commonly 24 hours.                                                    
	EndDateTime                                                                                 time.Time                                      `json:"EndDateTime"`
	// References to the Fluids Reports generated during this reporting period.                                                                
	FluidsIDs                                                                                   []string                                       `json:"FluidsIDs,omitempty"`
	// General information about a gas reading taken during the drill report period                                                            
	GasReading                                                                                  []GasReading                                   `json:"GasReading,omitempty"`
	// Health Safety or Environment events that occurred since the last drilling/operation                                                     
	// report. Captures data related to HSE events (e.g., tests, inspections, meetings, and                                                    
	// drills), test values (e.g., pressure tested to), and/or incidents (e.g., discharges,                                                    
	// non-compliance notices received, etc.).                                                                                                 
	Hse                                                                                         []HealthSafetyEnvironment                      `json:"HSE,omitempty"`
	// Description of incidents that have occurred during the last drilling/operations report                                                  
	Incident                                                                                    []Incident                                     `json:"Incident,omitempty"`
	// Quantity of items inventoried during drilling and/or operations                                                                         
	Inventory                                                                                   []Inventory                                    `json:"Inventory,omitempty"`
	// Operator personnel contact information on an operations report                                                                          
	JobContact                                                                                  []AbstractContactUserProfile                   `json:"JobContact,omitempty"`
	// Description of the lithology for the interval.                                                                                          
	Lithology                                                                                   *string                                        `json:"Lithology,omitempty"`
	// Mud that has been lost during drilling / operations                                                                                     
	MudLosses                                                                                   []MudLoss                                      `json:"MudLosses,omitempty"`
	// Information related to mud volumes for drilling/operations report                                                                       
	MudVolume                                                                                   []MudVolume                                    `json:"MudVolume,omitempty"`
	// An array of sequential operation activities descriptions performed during this reporting                                                
	// period. Potentially includes critical path and offline activities.                                                                      
	OperationsActivity                                                                          []OperationsActivity                           `json:"OperationsActivity,omitempty"`
	// Snapshot of operations personnel broken down by each company on the rig at the time of                                                  
	// the report.                                                                                                                             
	Personnel                                                                                   []Personnel                                    `json:"Personnel,omitempty"`
	// A reference to the PPFG information for this reporting period                                                                           
	PorePressure                                                                                []string                                       `json:"PorePressure,omitempty"`
	// Information related to pump operations on a drilling/operations report                                                                  
	PumpOp                                                                                      []PumpOperations                               `json:"PumpOp,omitempty"`
	// Report description                                                                                                                      
	ReportDescription                                                                           *string                                        `json:"ReportDescription,omitempty"`
	// Sequential number assigned to report header.                                                                                            
	ReportNumber                                                                                *string                                        `json:"ReportNumber,omitempty"`
	// Report remarks                                                                                                                          
	ReportRemarks                                                                               *string                                        `json:"ReportRemarks,omitempty"`
	// Date and time that the reporting period started. A report period is commonly 24 hours.                                                  
	StartDateTime                                                                               time.Time                                      `json:"StartDateTime"`
	// Information regarding the status of the wellbore during this reporting period                                                           
	StatusInfo                                                                                  []DrillingReportStatusInfo                     `json:"StatusInfo,omitempty"`
	// Authorized cost for the total job                                                                                                       
	TargetCost                                                                                  *float64                                       `json:"TargetCost,omitempty"`
	// Planned days for the total job                                                                                                          
	TargetDays                                                                                  *float64                                       `json:"TargetDays,omitempty"`
	// A series of time stamped comments which comprise part of this operations report.                                                        
	TimedComments                                                                               []TimedComment                                 `json:"TimedComments,omitempty"`
	// Cumulative cost for the job through the end of current report                                                                           
	TotalCost                                                                                   *float64                                       `json:"TotalCost,omitempty"`
	// Cumulative days for the job through the end of the current report                                                                       
	TotalDays                                                                                   *float64                                       `json:"TotalDays,omitempty"`
	// Total days of non-productive time through the end of the current report                                                                 
	TotalNPT                                                                                    *float64                                       `json:"TotalNPT,omitempty"`
	// Meteorological readings for the defined time period on an operations report                                                             
	Weather                                                                                     []Weather                                      `json:"Weather,omitempty"`
	// A link to the Well  Activity to which this report is associated.                                                                        
	WellActivityID                                                                              *string                                        `json:"WellActivityID,omitempty"`
	// Local name defined for the Well                                                                                                         
	WellAlias                                                                                   *WellAlias                                     `json:"WellAlias,omitempty"`
	// Local name defined for the wellbore                                                                                                     
	WellboreAlias                                                                               []WellboreAliasElement                         `json:"WellboreAlias,omitempty"`
	// A link to the wellbore that was active at the end of this reporting period.                                                             
	WellboreID                                                                                  string                                         `json:"WellboreID"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
