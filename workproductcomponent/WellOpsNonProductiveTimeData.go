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
type WellOpsNonProductiveTimeData struct {
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
	// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct                                     
	// from the file, in the sense certain valuable information is generated during loading                                          
	// process (Artefact generation process). Examples include retrieving location data,                                             
	// performing an OCR which may result in the generation of artefacts which need to be                                            
	// preserved distinctly                                                                                                          
	Artefacts                                                                                   []PurpleArtefacts                    `json:"Artefacts,omitempty"`
	// The record id, which identifies this OSDU File or dataset resource.                                                           
	Datasets                                                                                    []string                             `json:"Datasets,omitempty"`
	// An array of references to content in Domain Data Management Services represented by this                                      
	// work-product-component. The references are formed as URI following                                                            
	// https://www.rfc-editor.org/rfc/rfc3986#page-16. This property is exclusively populated by                                     
	// DDMSs. If a work-product-component is represented in more than one DDMS, DDMSs are                                            
	// obliged to find the specific reference by inspecting the URI's authority values matching                                      
	// the DDMS id.                                                                                                                  
	DDMSDatasets                                                                                []string                             `json:"DDMSDatasets,omitempty"`
	// A flag that indicates if the work product component is searchable, which means covered in                                     
	// the search index.                                                                                                             
	IsDiscoverable                                                                              *bool                                `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product component is undergoing an extended load.  It                                       
	// reflects the fact that the work product component is in an early stage and may be updated                                     
	// before finalization.                                                                                                          
	IsExtendedLoad                                                                              *bool                                `json:"IsExtendedLoad,omitempty"`
	// Alternative names, including historical, by which this work-product-component is/has been                                     
	// known (it should include all the identifiers).                                                                                
	NameAliases                                                                                 []AbstractAliasNames                 `json:"NameAliases,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                                       
	// quality. Clarifications: Since Certified is the highest classification of suitable                                            
	// quality, any further change or versioning of a Certified record should be carefully                                           
	// considered and justified. If a Technical Assurance value is not populated then one can                                        
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                              
	// Technical Assurance values are not intended to be used for the identification of a single                                     
	// "preferred" or "definitive" record by comparison with other records.                                                          
	TechnicalAssurances                                                                         []PurpleAbstractTechnicalAssurance   `json:"TechnicalAssurances,omitempty"`
	// Array of Authors' names of the work product component.  Could be a person or company                                          
	// entity.                                                                                                                       
	AuthorIDs                                                                                   []string                             `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product component has been through                                        
	// (ex. well planning, exploration).                                                                                             
	BusinessActivities                                                                          []string                             `json:"BusinessActivities,omitempty"`
	// Date that a resource (work  product component here) is formed outside of OSDU before                                          
	// loading (e.g. publication date).                                                                                              
	CreationDateTime                                                                            *time.Time                           `json:"CreationDateTime,omitempty"`
	// Description.  Summary of the work product component.  Not the same as Remark which                                            
	// captures thoughts of creator about the wpc.                                                                                   
	//                                                                                                                               
	// Description, free text entry usually by Rigsite personnel, but may be amended by                                              
	// Performance Engineer (NPT Coach)                                                                                              
	Description                                                                                 *string                              `json:"Description,omitempty"`
	// List of geographic entities which provide context to the WPC.  This may include multiple                                      
	// types or multiple values of the same type.                                                                                    
	GeoContexts                                                                                 []AbstractGeoContext                 `json:"GeoContexts,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                                          
	// product component depends.  The assertion is directed only from the asserting WPC to                                          
	// ancestor objects, not children.  It should not be used to refer to files or artefacts                                         
	// within the WPC -- the association within the WPC is sufficient and Artefacts are actually                                     
	// children of the main WPC file. They should be recorded in the data.Artefacts[] array.                                         
	LineageAssertions                                                                           []LineageAssertion                   `json:"LineageAssertions,omitempty"`
	// Name                                                                                                                          
	//                                                                                                                               
	// Human recognizable context/name of the NPT Event.                                                                             
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product component                                      
	// (location of the subject matter).                                                                                             
	SpatialArea                                                                                 *AbstractSpatialLocation             `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product component                                        
	// (location of the subject matter).                                                                                             
	SpatialPoint                                                                                *AbstractSpatialLocation             `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product component to OSDU.                                                   
	SubmitterName                                                                               *string                              `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                                
	Tags                                                                                        []string                             `json:"Tags,omitempty"`
	// Name of the NPT Coach/Well Operations Performance Engineer who was responsible for                                            
	// shepherding the NPT through its lifecycle                                                                                     
	AccountableParty                                                                            *string                              `json:"AccountableParty,omitempty"`
	// Contractor Cost                                                                                                               
	ContractorCost                                                                              *float64                             `json:"ContractorCost,omitempty"`
	// Cost Group/Code (Activity Element)                                                                                            
	CostGroupID                                                                                 *string                              `json:"CostGroupID,omitempty"`
	// Downtime Event Identifier, used when NPT Equipment failures also captured as downtime                                         
	// events                                                                                                                        
	DowntimeEventID                                                                             *string                              `json:"DowntimeEventID,omitempty"`
	// End Operations Activity Natural Identifier from Daily Operations Report. Requires that                                        
	// the relationship WellOperationsReport is populated. ActivityID must match a member in the                                     
	// related WellOperationsReport data.OperationsActivity[].ActivityID.                                                            
	EndOperationsActivityID                                                                     *string                              `json:"EndOperationsActivityID,omitempty"`
	// Start Operations Activity Daily Operations Report ID  in which the NPT Event ended.                                           
	EndOperationsReportID                                                                       *string                              `json:"EndOperationsReportID,omitempty"`
	// Equipment Cost                                                                                                                
	EquipmentCost                                                                               *float64                             `json:"EquipmentCost,omitempty"`
	// Equipment Failure Hours Before Fail                                                                                           
	EquipmentFailureHoursBeforeFail                                                             *float64                             `json:"EquipmentFailureHoursBeforeFail,omitempty"`
	// Equipment Failure Location Description                                                                                        
	EquipmentFailureLocationDescription                                                         *string                              `json:"EquipmentFailureLocationDescription,omitempty"`
	// Equipment Failure Type Identifier                                                                                             
	EquipmentFailureTypeID                                                                      *string                              `json:"EquipmentFailureTypeID,omitempty"`
	// Equipment Last Inspection Date/time                                                                                           
	EquipmentLastInspectionDateTime                                                             *time.Time                           `json:"EquipmentLastInspectionDateTime,omitempty"`
	// Equipment Manufacturer                                                                                                        
	EquipmentManufacturerID                                                                     *string                              `json:"EquipmentManufacturerID,omitempty"`
	// Equipment Model Name/Number                                                                                                   
	EquipmentModel                                                                              *string                              `json:"EquipmentModel,omitempty"`
	// Equipment Part Number                                                                                                         
	EquipmentPartNumber                                                                         *string                              `json:"EquipmentPartNumber,omitempty"`
	// Equipment Serial Number                                                                                                       
	EquipmentSerialNumber                                                                       *string                              `json:"EquipmentSerialNumber,omitempty"`
	// Equipment Size / Diameter                                                                                                     
	EquipmentSize                                                                               *float64                             `json:"EquipmentSize,omitempty"`
	// Date/time the Unplanned event ended                                                                                           
	EventEndTime                                                                                *time.Time                           `json:"EventEndTime,omitempty"`
	// Date/time the Unplanned event started                                                                                         
	EventStartTime                                                                              *time.Time                           `json:"EventStartTime,omitempty"`
	// Remedial actions taken by field/rig team when NPT occurred to address the NPT.                                                
	FieldRemediation                                                                            *string                              `json:"FieldRemediation,omitempty"`
	// Calculated Gross Cost Total = (BurnRate * Gross Duration) + Contractor Cost + Equipment                                       
	// Cost + Other Cost                                                                                                             
	GrossCostTotal                                                                              *float64                             `json:"GrossCostTotal,omitempty"`
	// Gross Time Duration = NPT End Date/time - Start Date/time                                                                     
	GrossDuration                                                                               *float64                             `json:"GrossDuration,omitempty"`
	// Investigation Status History                                                                                                  
	InvestigationStatus                                                                         []NPTInvestigationStatus             `json:"InvestigationStatus,omitempty"`
	// Used to flag NPT Events which were later classified as not NPT                                                                
	IsExcludeFromNPT                                                                            *bool                                `json:"IsExcludeFromNPT,omitempty"`
	// Is Investigation Waived                                                                                                       
	IsInvestigationWaived                                                                       *bool                                `json:"IsInvestigationWaived,omitempty"`
	// Learnings arising from the NPT investigation                                                                                  
	Learnings                                                                                   *string                              `json:"Learnings,omitempty"`
	// Nested Time duration - sum of Nested NPT Net Duration occurring within this NPT (can be                                       
	// calculated)                                                                                                                   
	NestedDuration                                                                              *float64                             `json:"NestedDuration,omitempty"`
	// Calculated Net Cost Total = (Burn Rate * Net Duration) + Contractor Cost + Equipment Cost                                     
	// + Other Cost                                                                                                                  
	NetCostTotal                                                                                *float64                             `json:"NetCostTotal,omitempty"`
	// Net Duration = Gross Duration - Nested Duration - Productive Time Adjustment                                                  
	NetDuration                                                                                 *float64                             `json:"NetDuration,omitempty"`
	// NPT Actions                                                                                                                   
	NPTActions                                                                                  []NonProductiveTimeAction            `json:"NPTActions,omitempty"`
	// Minor NPT Cause Type category                                                                                                 
	NPTCauseSubTypeID                                                                           *string                              `json:"NPTCauseSubTypeID,omitempty"`
	// Major NPT Cause Type category                                                                                                 
	NPTCauseTypeID                                                                              *string                              `json:"NPTCauseTypeID,omitempty"`
	// Measured Depth at time when the NPT event ended if different from start MD referenced                                         
	// from Vertical Measure Elevation                                                                                               
	NPTEndMeasuredDepth                                                                         *float64                             `json:"NPTEndMeasuredDepth,omitempty"`
	// True Vertical Depth at time when the NPT event end if different from start TVD referenced                                     
	// from Vertical Measure Elevation                                                                                               
	NPTEndTrueVerticalDepth                                                                     *float64                             `json:"NPTEndTrueVerticalDepth,omitempty"`
	// NPT Level (nesting). 0 = Productive Time (not NPT), 1 = 1st Unplanned Activity, 2 = 2nd                                       
	// Unplanned activity occurring within 1st NPT, 3 = 3rd NPT occurring within 2nd NPT, tc.                                        
	NPTLevel                                                                                    *int64                               `json:"NPTLevel,omitempty"`
	// Measured Depth at time when the NPT event started (when relevant) referenced from                                             
	// Vertical Measure Elevation                                                                                                    
	NPTMeasuredDepth                                                                            *float64                             `json:"NPTMeasuredDepth,omitempty"`
	// True Vertical Depth at time when the NPT event started (when relevant) referenced from                                        
	// Vertical Measure Elevation                                                                                                    
	NPTTrueVerticalDepth                                                                        *float64                             `json:"NPTTrueVerticalDepth,omitempty"`
	// Ongoing operations burn rate cost (daily cost). Used to calculate Gross and Net Cost                                          
	OperationsBurnRate                                                                          *float64                             `json:"OperationsBurnRate,omitempty"`
	// Other Cost                                                                                                                    
	OtherCost                                                                                   *float64                             `json:"OtherCost,omitempty"`
	// Parent NPT Event Identifier (when nested)                                                                                     
	ParentNPTEventID                                                                            *string                              `json:"ParentNPTEventID,omitempty"`
	// Preventative Actions                                                                                                          
	PreventativeActions                                                                         *string                              `json:"PreventativeActions,omitempty"`
	// Productive Time Adjustment duration - time spent working towards Well objective(s)                                            
	// whilst  the NPT is ongoing. Productive time negates against NPT Net Time                                                      
	ProductiveTimeAdjustment                                                                    *float64                             `json:"ProductiveTimeAdjustment,omitempty"`
	// Remarks                                                                                                                       
	Remarks                                                                                     *string                              `json:"Remarks,omitempty"`
	// Responsible Company Comments                                                                                                  
	ResponsibleCompanyComments                                                                  *string                              `json:"ResponsibleCompanyComments,omitempty"`
	// Responsible Company Contact Name                                                                                              
	ResponsibleCompanyContact                                                                   *string                              `json:"ResponsibleCompanyContact,omitempty"`
	// Responsible Company Findings                                                                                                  
	ResponsibleCompanyFindings                                                                  *string                              `json:"ResponsibleCompanyFindings,omitempty"`
	// Responsible Company                                                                                                           
	ResponsibleCompanyID                                                                        *string                              `json:"ResponsibleCompanyID,omitempty"`
	// Responsible Company Resolution Date                                                                                           
	ResponsibleCompanyResolutionDate                                                            *string                              `json:"ResponsibleCompanyResolutionDate,omitempty"`
	// Identifier of the Rig/Work Unit performing the activity at the time of the NPT                                                
	RigID                                                                                       *string                              `json:"RigID,omitempty"`
	// Risk Assessment Type ID                                                                                                       
	RiskAssessmentTypeID                                                                        *string                              `json:"RiskAssessmentTypeID,omitempty"`
	// ID of specific risk associated to this NPT                                                                                    
	RiskID                                                                                      *string                              `json:"RiskID,omitempty"`
	// Root Cause Description                                                                                                        
	RootCauseDescription                                                                        *string                              `json:"RootCauseDescription,omitempty"`
	// NPT incident Reference/Identifier to other Root Cause Failure Analysis system                                                 
	RootCauseFailureReferenceID                                                                 *string                              `json:"RootCauseFailureReferenceID,omitempty"`
	// Team assigned to investigate the NPT                                                                                          
	RootCauseFailureTeamID                                                                      *string                              `json:"RootCauseFailureTeamID,omitempty"`
	// Root Cause Type ID                                                                                                            
	RootCauseTypeID                                                                             *string                              `json:"RootCauseTypeID,omitempty"`
	// NPT Safety Classification - High Potential Incident, Process Safety Incident, or both.                                        
	// NULL identifies NPT Events with no safety implications.                                                                       
	SafetyClassificationID                                                                      *string                              `json:"SafetyClassificationID,omitempty"`
	// Safety Incident Reference to another system                                                                                   
	SafetyIncidentReferenceID                                                                   *string                              `json:"SafetyIncidentReferenceID,omitempty"`
	// Severity Level defined on Time and/or Cost impact/criteria                                                                    
	SeverityLevelID                                                                             *string                              `json:"SeverityLevelID,omitempty"`
	// Start Operations Activity Natural Identifier from Daily Operations Report. Requires that                                      
	// the relationship to WellOperationsReport is populated. ActivityName must match a member                                       
	// in the related WellOperationsReport data.OperationsActivity[].ActivityID.                                                     
	StartOperationsActivityID                                                                   *string                              `json:"StartOperationsActivityID,omitempty"`
	// Start Operations Activity Daily Operations Report ID in which the NPT Event started.                                          
	StartOperationsReportID                                                                     *string                              `json:"StartOperationsReportID,omitempty"`
	// Title, free text entry usually by Rigsite personnel                                                                           
	Title                                                                                       *string                              `json:"Title,omitempty"`
	// Unplanned Event Type classification (NPT w/ Equipment Failure, NPT w/out Equipment                                            
	// Failure, Equipment Failure (no NPT)                                                                                           
	UnplannedEventTypeID                                                                        *string                              `json:"UnplannedEventTypeID,omitempty"`
	// References an entry in the VerticalMeasurements array for the Wellbore identified by                                          
	// WellboreID, or a standalone vertical reference elevation for all measured depths within                                       
	// the NPT record. If this is not populated, the VerticalMeasurement is derived from the                                         
	// Wellbore default Vertical Measure Elevation.                                                                                  
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// Identifier of the Well Activity in which the NPT occurred                                                                     
	WellActivityID                                                                              string                               `json:"WellActivityID"`
	// Parent Wellbore Identifier                                                                                                    
	WellboreID                                                                                  string                               `json:"WellboreID"`
	// Association to Wellbore Marker Set (when relevant)                                                                            
	WellboreMarkerSetID                                                                         *string                              `json:"WellboreMarkerSetID,omitempty"`
	// Well Control Event Classification level                                                                                       
	WellControlEventClassificationID                                                            *string                              `json:"WellControlEventClassificationID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
