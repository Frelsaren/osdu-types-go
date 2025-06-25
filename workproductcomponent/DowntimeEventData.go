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
type DowntimeEventData struct {
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
	// The name of this downtime event.                                                                                                        
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
	// Identifies the "Where" condition of the downtime event; i.e. which component of the                                                     
	// system was affected? The downtime component values are further sorted into Surface and                                                  
	// Subsurface components                                                                                                                   
	DowntimeComponentID                                                                         *string                                        `json:"DowntimeComponentID,omitempty"`
	// Identifies the "Why" condition of the downtime event; i.e. why did it occur? These are                                                  
	// the basic reasons for, or causes of the event                                                                                           
	DowntimeReasonDetailID                                                                      *string                                        `json:"DowntimeReasonDetailID,omitempty"`
	// Identifies the "What" condition of the downtime event; i.e. what general type of activity                                               
	// or issue was involved? The downtime reason codes are grouped into Planned and Unplanned                                                 
	// categories                                                                                                                              
	DowntimeReasonID                                                                            *string                                        `json:"DowntimeReasonID,omitempty"`
	// Identifier of the record that specifies the loss caused by this event, or the type of the                                               
	// down time event is Planned, Unplanned or External.                                                                                      
	DownTimeTypeID                                                                              *string                                        `json:"DownTimeTypeID,omitempty"`
	// The identifier of the piece of equipment that caused the downtime. This is an identifier                                                
	// to an external system, not an OSDU resource.                                                                                            
	EquipmentIdentifier                                                                         *string                                        `json:"EquipmentIdentifier,omitempty"`
	// The date and time the downtime event ended                                                                                              
	EventEndTime                                                                                *string                                        `json:"EventEndTime,omitempty"`
	// The date and time the downtime event began                                                                                              
	EventStartTime                                                                              string                                         `json:"EventStartTime"`
	// Identifier of the record that specifies the event status, such as is suspended'' or                                                     
	// complete''.                                                                                                                             
	EventStatusID                                                                               *string                                        `json:"EventStatusID,omitempty"`
	// If there is a parent downtime event which has caused this event (e.g. a production                                                      
	// processing facility is down, causing downtime for connected wells), then this attribute                                                 
	// specifies the unique identifier of the entity associated with the parent event. If there                                                
	// is no parent event then this value is null                                                                                              
	ParentDowntimeEventID                                                                       *string                                        `json:"ParentDowntimeEventID,omitempty"`
	// Reporting entity which provides the reference to the master entity the downtime event                                                   
	// occurred at.                                                                                                                            
	ReportingEntityID                                                                           string                                         `json:"ReportingEntityID"`
	// The identifier of a work order which is associated with an event caused by a piece of                                                   
	// equipment. This is an identifier to an external system, not an OSDU resource.                                                           
	WorkOrderIdentifier                                                                         *string                                        `json:"WorkOrderIdentifier,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
