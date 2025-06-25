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
type TubularComponentData struct {
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
	// The drift diameter is the inside diameter (ID) that the pipe manufacturer guarantees per                                                
	// specifications. Note that the nominal inside diameter is not the same as the drift                                                      
	// diameter but is always slightly larger. The drift diameter is used by the well planner to                                               
	// determine what size tools or casing strings can later be run through the casing, whereas                                                
	// the nominal inside diameter is used for fluid volume calculations such as mud circulating                                               
	// times and cement slurry placement calculations.                                                                                         
	DriftDiameter                                                                               *float64                                       `json:"DriftDiameter,omitempty"`
	// Nominal inner diameter of the component.                                                                                                
	InnerDiameter                                                                               *float64                                       `json:"InnerDiameter,omitempty"`
	// Unique identifier for the manufacturer of this equipment.                                                                               
	ManufacturerID                                                                              *string                                        `json:"ManufacturerID,omitempty"`
	// This is the maximum outer diameter of the component.                                                                                    
	MaximumOuterDiameter                                                                        *float64                                       `json:"MaximumOuterDiameter,omitempty"`
	// Name of the component Model as defined per the operating company                                                                        
	Model                                                                                       *string                                        `json:"Model,omitempty"`
	// The depth the packer equipment was set to seal the casing or tubing.                                                                    
	PackerSetDepthTVD                                                                           *float64                                       `json:"PackerSetDepthTVD,omitempty"`
	// Identifier of the Assembly the component is taking part on.                                                                             
	ParentAssemblyID                                                                            *string                                        `json:"ParentAssemblyID,omitempty"`
	// Identifier of the wellbore the Component is standing in.                                                                                
	ParentWellboreID                                                                            *string                                        `json:"ParentWellboreID,omitempty"`
	// Size of the Pilot Hole                                                                                                                  
	PilotHoleSize                                                                               *float64                                       `json:"PilotHoleSize,omitempty"`
	// Identifier of the Section Type.                                                                                                         
	SectionTypeID                                                                               *string                                        `json:"SectionTypeID,omitempty"`
	// Serial Number of the component as provided by the manufacturer and/or the supplier                                                      
	SerialNumber                                                                                *string                                        `json:"SerialNumber,omitempty"`
	// Depth of the tubing shoe measured from the surface.                                                                                     
	ShoeDepthTVD                                                                                *float64                                       `json:"ShoeDepthTVD,omitempty"`
	// Unique identifier for the supplier of this equipment.                                                                                   
	SupplierID                                                                                  *string                                        `json:"SupplierID,omitempty"`
	// The measured depth of the base from the specific component                                                                              
	TubularComponentBaseMD                                                                      *float64                                       `json:"TubularComponentBaseMD,omitempty"`
	// Depth of the base of the component measured from the Well-Head                                                                          
	TubularComponentBaseReportedTVD                                                             *float64                                       `json:"TubularComponentBaseReportedTVD,omitempty"`
	// Id of the Bottom Connection Type                                                                                                        
	TubularComponentBottomConnectionTypeID                                                      *string                                        `json:"TubularComponentBottomConnectionTypeID,omitempty"`
	// ID of Reference Object for type of collar used to couple the tubular with another tubing                                                
	// string.                                                                                                                                 
	TubularComponentBoxPinConfigID                                                              *string                                        `json:"TubularComponentBoxPinConfigID,omitempty"`
	// Total Length of the component                                                                                                           
	TubularComponentLength                                                                      *float64                                       `json:"TubularComponentLength,omitempty"`
	// Specifies the material type constituting the component.                                                                                 
	TubularComponentMaterialTypeID                                                              *string                                        `json:"TubularComponentMaterialTypeID,omitempty"`
	// Nominal size (diameter) of the component, e.g. 9.625", 12.25                                                                            
	TubularComponentNominalSize                                                                 *float64                                       `json:"TubularComponentNominalSize,omitempty"`
	// Nominal weight of the component.                                                                                                        
	TubularComponentNominalWeight                                                               *float64                                       `json:"TubularComponentNominalWeight,omitempty"`
	// The sequence within which the components entered the hole. That is, a sequence number of                                                
	// 1 entered first, 2 entered next, etc.                                                                                                   
	TubularComponentSequence                                                                    *int64                                         `json:"TubularComponentSequence,omitempty"`
	// Id of the Top Connection Type                                                                                                           
	TubularComponentTopConnectionTypeID                                                         *string                                        `json:"TubularComponentTopConnectionTypeID,omitempty"`
	// The measured depth of the top from the specific component                                                                               
	TubularComponentTopMD                                                                       *float64                                       `json:"TubularComponentTopMD,omitempty"`
	// Depth of the top of the component measured from the Well-Head                                                                           
	TubularComponentTopReportedTVD                                                              *float64                                       `json:"TubularComponentTopReportedTVD,omitempty"`
	// Id of tubing grade - eg. the tensile strength of the tubing material. A system of                                                       
	// classifying the material specifications for steel alloys used in the manufacture of                                                     
	// tubing.                                                                                                                                 
	TubularComponentTubingGradeID                                                               *string                                        `json:"TubularComponentTubingGradeID,omitempty"`
	// The tensile strength of the tubing material. A system of classifying the material                                                       
	// specifications for steel alloys used in the manufacture of tubing.                                                                      
	TubularComponentTubingGradeStrength                                                         *float64                                       `json:"TubularComponentTubingGradeStrength,omitempty"`
	// The axial load required to yield the pipe.                                                                                              
	TubularComponentTubingStrength                                                              *float64                                       `json:"TubularComponentTubingStrength,omitempty"`
	// Specifies the types of components that can be used in a tubular string. These are used to                                               
	// specify the type of component and multiple components are used to define a tubular string                                               
	// (Tubular).                                                                                                                              
	TubularComponentTypeID                                                                      *string                                        `json:"TubularComponentTypeID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
