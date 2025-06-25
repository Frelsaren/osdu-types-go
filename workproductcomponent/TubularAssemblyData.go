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
type TubularAssemblyData struct {
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
	// Indicates if the Assembly is activated or not                                                                                            
	ActiveIndicator                                                                              *bool                                          `json:"ActiveIndicator,omitempty"`
	// Used to describe if it belongs to a RunActivity or to a PullActivity                                                                     
	ActivityTypeID                                                                               *string                                        `json:"ActivityTypeID,omitempty"`
	// Used to describe the reason of Activity - such as cut/pull, pulling,…                                                                    
	ActivityTypeReasonDescription                                                                *string                                        `json:"ActivityTypeReasonDescription,omitempty"`
	// Type of Artificial Lift used (could be "Surface Pump" / "Submersible Pump" / "Gas Lift"….)                                               
	ArtificialLiftTypeID                                                                         *string                                        `json:"ArtificialLiftTypeID,omitempty"`
	// The measured depth of the base from the whole assembly                                                                                   
	AssemblyBaseMD                                                                               *float64                                       `json:"AssemblyBaseMD,omitempty"`
	// Depth of the base of the Assembly measured from the Well-Head                                                                            
	AssemblyBaseReportedTVD                                                                      *float64                                       `json:"AssemblyBaseReportedTVD,omitempty"`
	// The measured depth of the top from the whole assembly                                                                                    
	AssemblyTopMD                                                                                *float64                                       `json:"AssemblyTopMD,omitempty"`
	// Depth of the top of the Assembly measured from the Well-Head                                                                             
	AssemblyTopReportedTVD                                                                       *float64                                       `json:"AssemblyTopReportedTVD,omitempty"`
	// The drift diameter is the inside diameter (ID) that the pipe manufacturer guarantees per                                                 
	// specifications. Note that the nominal inside diameter is not the same as the drift                                                       
	// diameter but is always slightly larger. The drift diameter is used by the well planner to                                                
	// determine what size tools or casing strings can later be run through the casing, whereas                                                 
	// the nominal inside diameter is used for fluid volume calculations such as mud circulating                                                
	// times and cement slurry placement calculations.                                                                                          
	DriftDiameter                                                                                *float64                                       `json:"DriftDiameter,omitempty"`
	// This reference table describes the type of liner used in the borehole. For example,                                                      
	// slotted, gravel packed or pre-perforated etc.                                                                                            
	LinerTypeID                                                                                  *string                                        `json:"LinerTypeID,omitempty"`
	// This is the minimum inner diameter of the whole assembly.                                                                                
	MinimumInnerDiameter                                                                         *float64                                       `json:"MinimumInnerDiameter,omitempty"`
	// A YES or NO flag indicating the assembly is a mixed string. The length of the assembly                                                   
	// may be made up of joints with different tensile strengths, or collapse resistance and                                                    
	// yield strengths.                                                                                                                         
	MixedStringIndicator                                                                         *string                                        `json:"MixedStringIndicator,omitempty"`
	// Optional - Identifier of the parent assembly (in case of side-track, multi-nesting,…) -                                                  
	// The Concentric Tubular model is used to identify the Assembly that an Assembly sits                                                      
	// inside e.g. Surface Casing set inside Conductor, Tubing set inside Production Casing, a                                                  
	// Bumper Spring set inside a Production Tubing Profile Nipple, Liner set inside Casing,                                                    
	// etc. This is needed to enable a Digital Well Sketch application to understand                                                            
	// relationships between Assemblies and their parent Wellbores.                                                                             
	ParentAssemblyID                                                                             *string                                        `json:"ParentAssemblyID,omitempty"`
	// Identifier of the wellbore the Component is standing in.                                                                                 
	ParentWellboreID                                                                             *string                                        `json:"ParentWellboreID,omitempty"`
	// Diameter of the Pilot Hole                                                                                                               
	PilotHoleSize                                                                                *float64                                       `json:"PilotHoleSize,omitempty"`
	// the distance that the assembly has penetrated below the surface of the sea floor.                                                        
	SeaFloorPenetrationLength                                                                    *float64                                       `json:"SeaFloorPenetrationLength,omitempty"`
	// Descriptor for Assembly, e.g. Production, Surface, Conductor, Intermediate, Drilling                                                     
	StringClassID                                                                                *string                                        `json:"StringClassID,omitempty"`
	// In case of multi-nesting of assemblies, the 'point' is the Measured Depth of the top of                                                  
	// the assembly though with PBRs the Suspension Point may not be the top.                                                                   
	SuspensionPointMD                                                                            *float64                                       `json:"SuspensionPointMD,omitempty"`
	// Nominal size (diameter) describing the whole assembly, e.g. 9.625", 12.25                                                                
	TubularAssemblyNominalSize                                                                   *float64                                       `json:"TubularAssemblyNominalSize,omitempty"`
	// Sequence of the TubularAssembly (Typically BHA sequence)                                                                                 
	TubularAssemblyNumber                                                                        *int64                                         `json:"TubularAssemblyNumber,omitempty"`
	// The full record of historical and current states of the Assembly. The current active                                                     
	// state is recorded in TubularAssemblyStatus.                                                                                              
	TubularAssemblyStates                                                                        []TubularAssemblyStateElement                  `json:"TubularAssemblyStates,omitempty"`
	// Reflects the current status of the Assembly - as 'installed', 'pulled', 'planned',... -                                                  
	// Applicable to tubing/completions as opposed to drillstrings. Historical states are                                                       
	// recorded in TubularAssemblyStates.                                                                                                       
	TubularAssemblyStatus                                                                        *TubularAssemblyStatusClass                    `json:"TubularAssemblyStatus,omitempty"`
	// Total Length of the whole assembly.                                                                                                      
	TubularAssemblyTotalLength                                                                   *float64                                       `json:"TubularAssemblyTotalLength,omitempty"`
	// Type of tubular assembly.                                                                                                                
	TubularAssemblyTypeID                                                                        *string                                        `json:"TubularAssemblyTypeID,omitempty"`
	// Defines whether the sequence of child tubular components runs either top to bottom, or                                                   
	// bottom to top.                                                                                                                           
	TubularDirection                                                                             *string                                        `json:"TubularDirection,omitempty"`
	// Either a self-contained vertical reference for the depths in this TubularAssembly or a                                                   
	// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the                                                      
	// entity defined by VerticalReferenceEntityID.                                                                                             
	VerticalMeasurement                                                                          *AbstractFacilityVerticalMeasurement           `json:"VerticalMeasurement,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
