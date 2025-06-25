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
type SamplesAnalysesReportData struct {
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
	// The date and time that the document was last modified                                                                                   
	DateModified                                                                                *time.Time                                     `json:"DateModified,omitempty"`
	// The date and time that the results of the analysis containing observed measurements or                                                  
	// calculations was published                                                                                                              
	DatePublished                                                                               *time.Time                                     `json:"DatePublished,omitempty"`
	// Document language as listed in the ISO 639-3 https://en.wikipedia.org/wiki/ISO_639,                                                     
	// https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes                                                                                   
	DocumentLanguage                                                                            *string                                        `json:"DocumentLanguage,omitempty"`
	// A description text or an array of subjects covered by the document. If present this                                                     
	// information must compliment the Tag and SubTitle                                                                                        
	DocumentSubject                                                                             *string                                        `json:"DocumentSubject,omitempty"`
	// The kind of document--from a business standpoint, e.g., multi-well study, etc.                                                          
	DocumentTypeID                                                                              *string                                        `json:"DocumentTypeID,omitempty"`
	// OSDU Record IDs for the laboratories used to conduct the sample analyses contained in                                                   
	// this report. This could represent the laboratory company or the laboratory branch.                                                      
	LaboratoryIDs                                                                               []string                                       `json:"LaboratoryIDs,omitempty"`
	// List of names of laboratories used to conduct the sample analyses contained in this                                                     
	// report. This attribute is more freeform than Laboratory IDs and does not have                                                           
	// relationships to OSDU Record IDs. This attribute exists to help with a low bar of                                                       
	// ingestion but best practice is to resolve each of these to an OSDU Record ID in the                                                     
	// LaboratoryIDs attribute.                                                                                                                
	LaboratoryNames                                                                             []string                                       `json:"LaboratoryNames,omitempty"`
	// Number of pages in the document, useful in cases where if it was described in the                                                       
	// acquired manifest as opposed to a derived/calculated value                                                                              
	NumberOfPages                                                                               *int64                                         `json:"NumberOfPages,omitempty"`
	// An array containing operational or quality comments pertaining to the sample analysis                                                   
	// represented by this work product component.                                                                                             
	Remarks                                                                                     []AbstractRemark                               `json:"Remarks,omitempty"`
	// The names or identifiers of the analyzed samples in this report or document. These names                                                
	// or identifiers are often assigned by the laboratory. This list can be used to assist in                                                 
	// locating the appropriate OSDU Sample records or to find the dataset in the parent report.                                               
	ReportSampleIdentifiers                                                                     []string                                       `json:"ReportSampleIdentifiers,omitempty"`
	// A list of all sample analysis types represented by this report, whether the sample                                                      
	// analysis type relates to Rock, Fluid, or both.                                                                                          
	SampleAnalysisTypeIDs                                                                       []string                                       `json:"SampleAnalysisTypeIDs,omitempty"`
	// OSDU Record IDs for the Rock and/or Fluid Sample or Samples on which this batch or                                                      
	// batches of analysis were performed.                                                                                                     
	SampleIDs                                                                                   []string                                       `json:"SampleIDs,omitempty"`
	// List of higher level grouping terms that are often used within organisations to search                                                  
	// for analysis types in addition to the formal Family and Method properties, but which are                                                
	// variable across organisations and therefore do not need to be governed strictly by the                                                  
	// OSDU Forum. For example: SCAL, Static SCAL, Dynamic SCAL, Source Rock Analysis. We                                                      
	// suggest you use a namespace, like the operators name, like “Shell.SCAL”. BP.SCAL,                                                       
	// Equinor.SCAL, Chevron.SCAL, Exxon.SCAL.....                                                                                             
	SamplesAnalysisCategoryTagIDs                                                               []string                                       `json:"SamplesAnalysisCategoryTagIDs,omitempty"`
	// The sub-title of the document.                                                                                                          
	SubTitle                                                                                    *string                                        `json:"SubTitle,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
