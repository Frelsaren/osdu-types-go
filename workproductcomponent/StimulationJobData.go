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
type StimulationJobData struct {
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
	// Date and time at which the stimulation contractor arrives on location.                                                                  
	ArrivalDateTime                                                                             *time.Time                                     `json:"ArrivalDateTime,omitempty"`
	// Bottomhole static temperature for the job (BHST)                                                                                        
	BHStaticTemperature                                                                         *float64                                       `json:"BHStaticTemperature,omitempty"`
	// Expected or calculated bottomhole treating temperature for the job.                                                                     
	BHTreatingTemperature                                                                       *float64                                       `json:"BHTreatingTemperature,omitempty"`
	// Total clean volume pumped for all job stages during this stim job.  Clean volume is                                                     
	// amount of fluid before proppant is added.                                                                                               
	CleanTotalVolume                                                                            *float64                                       `json:"CleanTotalVolume,omitempty"`
	// Identifier of service company performing the stimulation job.                                                                           
	ContractorID                                                                                *string                                        `json:"ContractorID,omitempty"`
	// Pressure recorded on fluid returning to surface.                                                                                        
	FlowBackPressure                                                                            *float64                                       `json:"FlowBackPressure,omitempty"`
	// Flow rate recorded on fluid returning to surface.                                                                                       
	FlowBackRate                                                                                *float64                                       `json:"FlowBackRate,omitempty"`
	// Volume recorded on fluid returning to surface.                                                                                          
	FlowBackVolume                                                                              *float64                                       `json:"FlowBackVolume,omitempty"`
	// Volume ratio of fluid in the fracture at the end of pumping.                                                                            
	FluidEfficiency                                                                             *float64                                       `json:"FluidEfficiency,omitempty"`
	// Maximum job fluid pumping rate encountered during treatment of all stages.                                                              
	FluidPumpingRateMax                                                                         *float64                                       `json:"FluidPumpingRateMax,omitempty"`
	// Hydraulic horsepower ordered for the stimulation job.                                                                                   
	HhpOrdered                                                                                  *float64                                       `json:"HhpOrdered,omitempty"`
	// Hydraulic horsepower actually used for the stimulation job.                                                                             
	HhpUsed                                                                                     *float64                                       `json:"HhpUsed,omitempty"`
	// Hole Section stimulation job performed in.                                                                                              
	HoleSectionID                                                                               *string                                        `json:"HoleSectionID,omitempty"`
	// Is the stimulation conveyed via coiled tubing?                                                                                          
	IsCoiledTubingConveyed                                                                      *bool                                          `json:"IsCoiledTubingConveyed,omitempty"`
	// Ending date and time of the stimulation job.                                                                                            
	JobEndDateTime                                                                              *time.Time                                     `json:"JobEndDateTime,omitempty"`
	// Name of the stimulation job.                                                                                                            
	JobName                                                                                     *string                                        `json:"JobName,omitempty"`
	// Average pressure encountered during treatment of all stages.                                                                            
	JobPressureAvg                                                                              *float64                                       `json:"JobPressureAvg,omitempty"`
	// Maximum pressure encountered during the job.                                                                                            
	JobPressureMax                                                                              *float64                                       `json:"JobPressureMax,omitempty"`
	// Start date and time of the stimulation job.                                                                                             
	JobStartDateTime                                                                            *time.Time                                     `json:"JobStartDateTime,omitempty"`
	// Total volume pumped for all stages.                                                                                                     
	JobTotalVolume                                                                              *float64                                       `json:"JobTotalVolume,omitempty"`
	// Identifier of the type of well stimulation job                                                                                          
	JobTypeID                                                                                   string                                         `json:"JobTypeID"`
	// Object that contains the Additives and Proppants on location and used in the stimulation                                                
	// job.                                                                                                                                    
	MaterialCatalog                                                                             []MaterialCatalogue                            `json:"MaterialCatalog,omitempty"`
	// Total amount of materials used for the Stimulation Job.                                                                                 
	MaterialUsed                                                                                []AbstractStimMaterialQuantity                 `json:"MaterialUsed,omitempty"`
	// Total length of stimulated interval for this stim job.                                                                                  
	OpenHoleLength                                                                              *float64                                       `json:"OpenHoleLength,omitempty"`
	// Name of Operator Representative/Supervisor                                                                                              
	OperatorRepresentative                                                                      *string                                        `json:"OperatorRepresentative,omitempty"`
	// Operator Representative Remarks                                                                                                         
	OperatorRepresentativeRemarks                                                               *string                                        `json:"OperatorRepresentativeRemarks,omitempty"`
	// The ratio of proppant placed in formation vs proppant pumped                                                                            
	PadPercent                                                                                  *float64                                       `json:"PadPercent,omitempty"`
	// Petroleum Industry Data Exchange (PIDX) UNSPSC (Segment 71) commodity code from the oil                                                 
	// and gas extraction and production enhancement services family.                                                                          
	PIDXCommodityCodeID                                                                         *string                                        `json:"PIDXCommodityCodeID,omitempty"`
	// The total mass of proppant placed in the formation for the entire stim job.                                                             
	ProppantInFormationTotal                                                                    *float64                                       `json:"ProppantInFormationTotal,omitempty"`
	// The total mass of proppant used (pumped) for the entire stim job.                                                                       
	ProppantUsedTotal                                                                           *float64                                       `json:"ProppantUsedTotal,omitempty"`
	// The total pumping time/duration for the entire stim job.                                                                                
	PumpTimeTotal                                                                               *float64                                       `json:"PumpTimeTotal,omitempty"`
	// General remarks about this Stim Job.                                                                                                    
	Remarks                                                                                     *string                                        `json:"Remarks,omitempty"`
	// Integer Number of stages treated during the stimulation job.                                                                            
	StageCount                                                                                  *int64                                         `json:"StageCount,omitempty"`
	// Name of the service company supervisor.                                                                                                 
	Supervisor                                                                                  *string                                        `json:"Supervisor,omitempty"`
	// ID to the Zero Depth Point Vertical Measure elevation for depths contained in the                                                       
	// stimulation job, job stages and other objects used to correlate MDs to original drilling                                                
	// rig MD. References an entry in the Vertical Measurement array for the Well parented by                                                  
	// the wellbore via WellboreID.                                                                                                            
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement           `json:"VerticalMeasurement,omitempty"`
	// Water management plan identifier for this stim job.                                                                                     
	WaterManagementPlanIdentifier                                                               *string                                        `json:"WaterManagementPlanIdentifier,omitempty"`
	// Source for water used during the stim job.                                                                                              
	WaterSource                                                                                 *string                                        `json:"WaterSource,omitempty"`
	// A link to the well activity to which this stimulation job is associated.                                                                
	WellActivityID                                                                              *string                                        `json:"WellActivityID,omitempty"`
	// A link to the wellbore in which the stimulation took place.                                                                             
	WellboreID                                                                                  string                                         `json:"WellboreID"`
	// Installed tubular through which the stimulation job is performed.                                                                       
	WellboreTubularID                                                                           *string                                        `json:"WellboreTubularID,omitempty"`
	// Log(s) associated to the Stim Job e.g. Composite Log(s)                                                                                 
	WellLog                                                                                     []string                                       `json:"WellLog,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
