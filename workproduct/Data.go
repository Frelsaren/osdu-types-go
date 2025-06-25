package workproduct

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
type Data struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                   
	ExistenceKind                                                                               *string                  `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                            
	ResourceCurationStatus                                                                      *string                  `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                    
	ResourceHomeRegionID                                                                        *string                  `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                 
	ResourceHostRegionIDs                                                                       []string                 `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                  
	ResourceLifecycleStatus                                                                     *string                  `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                    
	ResourceSecurityClassification                                                              *string                  `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                    
	// organization, agency, system, internal team, or individual. For informational purposes                            
	// only, the list of sources is not governed.                                                                        
	Source                                                                                      *string                  `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                             
	// based on data quality. Clarifications: Since Certified is the highest classification of                           
	// suitable quality, any further change or versioning of a Certified record should be                                
	// carefully considered and justified. If a Technical Assurance value is not populated then                          
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                          
	// Technical Assurance values are not intended to be used for the identification of a single                         
	// "preferred" or "definitive" record by comparison with other records.                                              
	TechnicalAssuranceID                                                                        *string                  `json:"TechnicalAssuranceID,omitempty"`
	// Array of Annotations                                                                                              
	Annotations                                                                                 []string                 `json:"Annotations,omitempty"`
	// Array of Authors' names of the work product.  Could be a person or company entity.                                
	AuthorIDs                                                                                   []string                 `json:"AuthorIDs,omitempty"`
	// Array of business processes/workflows that the work product has been through (ex. well                            
	// planning, exploration).                                                                                           
	BusinessActivities                                                                          []string                 `json:"BusinessActivities,omitempty"`
	Components                                                                                  []string                 `json:"Components,omitempty"`
	// Date that a resource (work  product here) is formed outside of OSDU before loading (e.g.                          
	// publication date, work product delivery package assembly date).                                                   
	CreationDateTime                                                                            *time.Time               `json:"CreationDateTime,omitempty"`
	// Description of the purpose of the work product.                                                                   
	Description                                                                                 *string                  `json:"Description,omitempty"`
	// A flag that indicates if the work product is searchable, which means covered in the                               
	// search index.                                                                                                     
	IsDiscoverable                                                                              *bool                    `json:"IsDiscoverable,omitempty"`
	// A flag that indicates if the work product is undergoing an extended load.  It reflects                            
	// the fact that the work product is in an early stage and may be updated before                                     
	// finalization.                                                                                                     
	IsExtendedLoad                                                                              *bool                    `json:"IsExtendedLoad,omitempty"`
	// Defines relationships with other objects (any kind of Resource) upon which this work                              
	// product depends.  The assertion is directed only from the asserting WP to ancestor                                
	// objects, not children.  It should not be used to refer to files or artefacts within the                           
	// WP -- the association within the WP is sufficient and Artefacts are actually children of                          
	// the main WP file. They should be recorded in the data.Artefacts[] array.                                          
	LineageAssertions                                                                           []LineageAssertion       `json:"LineageAssertions,omitempty"`
	// Name of the instance of Work Product - could be a shipment number.                                                
	Name                                                                                        *string                  `json:"Name,omitempty"`
	// A polygon boundary that reflects the locale of the content of the work product (location                          
	// of the subject matter).                                                                                           
	SpatialArea                                                                                 *AbstractSpatialLocation `json:"SpatialArea,omitempty"`
	// A centroid point that reflects the locale of the content of the work product (location of                         
	// the subject matter).                                                                                              
	SpatialPoint                                                                                *AbstractSpatialLocation `json:"SpatialPoint,omitempty"`
	// Name of the person that first submitted the work product package to OSDU.                                         
	SubmitterName                                                                               *string                  `json:"SubmitterName,omitempty"`
	// Array of key words to identify the work product, especially to help in search.                                    
	Tags                                                                                        []string                 `json:"Tags,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}   `json:"ExtensionProperties,omitempty"`
}
