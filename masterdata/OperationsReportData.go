package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type OperationsReportData struct {
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
	// Information about a bit.                                                                                              
	BitRecord                                                                                   []BitRecord                  `json:"BitRecord,omitempty"`
	// Hole condition description.                                                                                           
	ConditionHole                                                                               *string                      `json:"ConditionHole,omitempty"`
	// Cost information captured for a defined time period during operations and/or drilling                                 
	Cost                                                                                        []Cost                       `json:"Cost,omitempty"`
	// Daily cost.                                                                                                           
	CostDay                                                                                     *float64                     `json:"CostDay,omitempty"`
	// Daily mud cost.                                                                                                       
	CostDayMud                                                                                  *float64                     `json:"CostDayMud,omitempty"`
	// DEPRECATED: Individual, company, or corporate division that work is being executed on                                 
	// behalf of. With 1.1.0 or higher, please move this value to the associated Well Activity                               
	// record (data.WellActivityID), i.e. data.Customer.                                                                     
	Customer                                                                                    *string                      `json:"Customer,omitempty"`
	// DEPRECATED: Reference to the Organisation that represents the Customer. With 1.1.0 or                                 
	// higher, please obtain the Organisation's data.Name and set it as data.Customer value in                               
	// the Well Activity record (via data.WellActivityID).                                                                   
	CustomerID                                                                                  *string                      `json:"CustomerId,omitempty"`
	// DEPRECATED: An Array of sequential operation activities descriptions performed during                                 
	// this reporting period. Potentially includes critical path and offline activities.                                     
	DrillActivity                                                                               []DrillingActivity           `json:"DrillActivity,omitempty"`
	// Date and time that the reporting period ended. A report period is commonly 24 hours.                                  
	EndDateTime                                                                                 time.Time                    `json:"EndDateTime"`
	// References to the Fluids Reports generated during this reporting period.                                              
	FluidsIDs                                                                                   []string                     `json:"FluidsIDs,omitempty"`
	// DEPRECATED: Forecast of activities for the next 24 hrs. This is a redundant property                                  
	// given the data.StatusInfo[].Forecast24Hr. Consolidate in StatusInfo.                                                  
	Forecast24Hr                                                                                *string                      `json:"Forecast24Hr,omitempty"`
	// General information about a gas reading taken during the drill report period                                          
	GasReading                                                                                  []GasReading                 `json:"GasReading,omitempty"`
	// Health Safety or Environment events that occurred since the last drilling/operation                                   
	// report. Captures data related to HSE events (e.g., tests, inspections, meetings, and                                  
	// drills), test values (e.g., pressure tested to), and/or incidents (e.g., discharges,                                  
	// non-compliance notices received, etc.).                                                                               
	Hse                                                                                         []HealthSafetyEnvironment    `json:"HSE,omitempty"`
	// Description of incidents that have occurred during the last drilling/operations report                                
	Incident                                                                                    []Incident                   `json:"Incident,omitempty"`
	// Quantity of items inventoried during drilling and/or operations                                                       
	Inventory                                                                                   []Inventory                  `json:"Inventory,omitempty"`
	// Operator personnel contact information on an operations report                                                        
	JobContact                                                                                  []AbstractContact            `json:"JobContact,omitempty"`
	// Description of the lithology for the interval.                                                                        
	Lithology                                                                                   *string                      `json:"Lithology,omitempty"`
	// Mud that has been lost during drilling / operations                                                                   
	MudLosses                                                                                   []MudLosses                  `json:"MudLosses,omitempty"`
	// Information related to mud volumes for drilling/operations report                                                     
	MudVolume                                                                                   []MudVolume                  `json:"MudVolume,omitempty"`
	// Name of Operations Report                                                                                             
	Name                                                                                        *string                      `json:"Name,omitempty"`
	// An array of sequential operation activities descriptions performed during this reporting                              
	// period. Potentially includes critical path and offline activities.                                                    
	OperationsActivity                                                                          []OperationsActivity         `json:"OperationsActivity,omitempty"`
	// Snapshot of operations personnel broken down by each company on the rig at the time of                                
	// the report.                                                                                                           
	Personnel                                                                                   []FluffyPersonnel            `json:"Personnel,omitempty"`
	// A reference to the PPFG information for this reporting period                                                         
	PorePressure                                                                                []string                     `json:"PorePressure,omitempty"`
	// Information related to pump operations on a drilling/operations report                                                
	PumpOp                                                                                      []PumpOperations             `json:"PumpOp,omitempty"`
	// Report description                                                                                                    
	ReportDescription                                                                           *string                      `json:"ReportDescription,omitempty"`
	// Sequential number assigned to report header.                                                                          
	ReportNumber                                                                                *string                      `json:"ReportNumber,omitempty"`
	// Report remarks                                                                                                        
	ReportRemarks                                                                               *string                      `json:"ReportRemarks,omitempty"`
	// Date and time that the reporting period started. A report period is commonly 24 hours.                                
	StartDateTime                                                                               time.Time                    `json:"StartDateTime"`
	// Information regarding the status of the wellbore during this reporting period                                         
	StatusInfo                                                                                  []DrillingReportStatusInfo   `json:"StatusInfo,omitempty"`
	// DEPRECATED: Company or corporate division that is responsible for executing the work.                                 
	// With 1.1.0 or higher, please move this value to the associated Well Activity record                                   
	// (data.WellActivityID), i.e. data.StewardingCompany.                                                                   
	StewardingCompany                                                                           *string                      `json:"StewardingCompany,omitempty"`
	// DEPRECATED: Reference to the Organisation that represents the StewardingCompany. With                                 
	// 1.1.0 or higher, please obtain the Organisation's data.Name and set it as                                             
	// data.StewardingCompany value in the Well Activity record.                                                             
	StewardingCompanyID                                                                         *string                      `json:"StewardingCompanyID,omitempty"`
	// DEPRECATED: Team within a company or corporate division that is responsible for executing                             
	// the work. With 1.1.0 or higher, please move this value to the associated Well Activity                                
	// record (data.WellActivityID), i.e. data.StewardingCompanyTeam.                                                        
	StewardingCompanyTeam                                                                       *string                      `json:"StewardingCompanyTeam,omitempty"`
	// DEPRECATED: Reference to the Organisation that represents the StewardingCompanyTeam. With                             
	// 1.1.0 or higher, please obtain the Organisation's data.Name and set it as                                             
	// data.StewardingCompanyTeam value in the Well Activity record.                                                         
	StewardingCompanyTeamID                                                                     *string                      `json:"StewardingCompanyTeamID,omitempty"`
	// Authorized cost for the total job                                                                                     
	TargetCost                                                                                  *float64                     `json:"TargetCost,omitempty"`
	// Planned days for the total job                                                                                        
	TargetDays                                                                                  *float64                     `json:"TargetDays,omitempty"`
	// A series of time stamped comments which comprise part of this operations report.                                      
	TimedComments                                                                               []TimedComment               `json:"TimedComments,omitempty"`
	// Cumulative cost for the job through the end of current report                                                         
	TotalCost                                                                                   *float64                     `json:"TotalCost,omitempty"`
	// Cumulative days for the job through the end of the current report                                                     
	TotalDays                                                                                   *float64                     `json:"TotalDays,omitempty"`
	// Total days of non-productive time through the end of the current report                                               
	TotalNPT                                                                                    *float64                     `json:"TotalNPT,omitempty"`
	// Meteorological readings for the defined time period on an operations report                                           
	Weather                                                                                     []Weather                    `json:"Weather,omitempty"`
	// A link to the well activity to which this report is associated.                                                       
	WellActivityID                                                                              *string                      `json:"WellActivityID,omitempty"`
	// Local name defined for the Well                                                                                       
	WellAlias                                                                                   *WellAlias                   `json:"WellAlias,omitempty"`
	// Local name defined for the wellbore                                                                                   
	WellboreAlias                                                                               []WellboreAliasElement       `json:"WellboreAlias,omitempty"`
	// A link to the wellbore that was active at the end of this report period.                                              
	WellboreID                                                                                  string                       `json:"WellboreID"`
	ExtensionProperties                                                                         map[string]interface{}       `json:"ExtensionProperties,omitempty"`
}
