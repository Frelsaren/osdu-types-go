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
type SeismicFieldTraceDataData struct {
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
	// Indicates if the volume is a product of the difference between 4D surveys                                                               
	Difference                                                                                  *bool                                          `json:"Difference,omitempty"`
	// The sample axis value in the vertical dimension at which Depth formatted data ends. Use                                                 
	// SeismicDomainType to know which of time or depth pairs is the actual range vs. what is                                                  
	// estimated.                                                                                                                              
	EndDepth                                                                                    *float64                                       `json:"EndDepth,omitempty"`
	// The sample axis value in the vertical dimension at which Time formatted data starts. Use                                                
	// SeismicDomainType to know which of time or depth pairs is the actual range vs. what is                                                  
	// estimated.                                                                                                                              
	EndTime                                                                                     *float64                                       `json:"EndTime,omitempty"`
	// The shotpoint that came before all others                                                                                               
	FirstShotPoint                                                                              *float64                                       `json:"FirstShotPoint,omitempty"`
	// The type of gathers in this dataset.                                                                                                    
	GatherTypeID                                                                                *string                                        `json:"GatherTypeID,omitempty"`
	// Coordinate reference system of positions in trace header, which matches what is described                                               
	// in BinGrid but is repeated here for convenience and in case bin grid is not defined.  In                                                
	// case of conflict with Bin Grid, this value applies to the coordinates as written into                                                   
	// CMPX, CMPY headers. Nevertheless, Bin Grid should be used for mapping purposes.                                                         
	HorizontalCRSID                                                                             *string                                        `json:"HorizontalCRSID,omitempty"`
	// The last shotpoint represented by the data                                                                                              
	LastShotPoint                                                                               *float64                                       `json:"LastShotPoint,omitempty"`
	// Polygon showing the coverage of live traces in the trace dataset                                                                        
	LiveTraceOutline                                                                            *AbstractSpatialLocation                       `json:"LiveTraceOutline,omitempty"`
	// Sample data format in terms of sample value precision 8bit Integer, 16bit Floating Point                                                
	// etc.                                                                                                                                    
	Precision                                                                                   *PurplePrecision                               `json:"Precision,omitempty"`
	// For most datasets, the acquisition project that generated the underlying field data.  For                                               
	// merges, probably absent (see processing project for set of acquisition projects used in                                                 
	// processing this dataset).                                                                                                               
	PrincipalAcquisitionProjectID                                                               *string                                        `json:"PrincipalAcquisitionProjectID,omitempty"`
	// Processing Parameters to simply capture process history until full provenance model can                                                 
	// be implemented.                                                                                                                         
	ProcessingParameters                                                                        []ProcessingParameters                         `json:"ProcessingParameters,omitempty"`
	// The actual maximum amplitude value found in the dataset.                                                                                
	RangeAmplitudeMax                                                                           *float64                                       `json:"RangeAmplitudeMax,omitempty"`
	// The actual minimum amplitude value found in the dataset.                                                                                
	RangeAmplitudeMin                                                                           *float64                                       `json:"RangeAmplitudeMin,omitempty"`
	// Number of samples in the vertical direction.                                                                                            
	SampleCount                                                                                 *int64                                         `json:"SampleCount,omitempty"`
	// Vertical sampling interval of data.                                                                                                     
	SampleInterval                                                                              *float64                                       `json:"SampleInterval,omitempty"`
	// 2D line name or survey name for the 2D group                                                                                            
	Seismic2DName                                                                               *string                                        `json:"Seismic2DName,omitempty"`
	// ID of the Seismic Trace Data Type                                                                                                       
	SeismicAttributeTypeID                                                                      *string                                        `json:"SeismicAttributeTypeID,omitempty"`
	// ID of the nature of the vertical axis in the trace data set, usually Depth or Time.                                                     
	SeismicDomainTypeID                                                                         *string                                        `json:"SeismicDomainTypeID,omitempty"`
	// ID of the Seismic Filtering Type                                                                                                        
	SeismicFilteringTypeID                                                                      *string                                        `json:"SeismicFilteringTypeID,omitempty"`
	// Reference to the reference-data--SeismicPhase value describing the wavelet phase as                                                     
	// estimated/processed for this SeismicTraceData instance.                                                                                 
	SeismicPhaseID                                                                              *string                                        `json:"SeismicPhaseID,omitempty"`
	// Reflection polarity of embedded wavelet.  Normal, Reverse, Plus 90, Minus 90 according to                                               
	// SEG standard.                                                                                                                           
	SeismicPolarityID                                                                           *string                                        `json:"SeismicPolarityID,omitempty"`
	// The dimensionality of trace data sets (not as acquired but as in the dataset), such as                                                  
	// 2D, 3D, 4D.                                                                                                                             
	SeismicTraceDataDimensionalityTypeID                                                        *string                                        `json:"SeismicTraceDataDimensionalityTypeID,omitempty"`
	// The observed wave mode type in this trace data set (P, Sv, etc).                                                                        
	SeismicWaveTypeID                                                                           *string                                        `json:"SeismicWaveTypeID,omitempty"`
	// Indicates how much the data has been shifted from the Vertical Datum (seismic reference                                                 
	// datum) in the domain and units of SeismicDomainType and in the sense that a positive                                                    
	// number causes a sample to move downward in physical space (lower elevation).                                                            
	ShiftApplied                                                                                *float64                                       `json:"ShiftApplied,omitempty"`
	// Defines the sorting order of the trace data as stored in the file(s).                                                                   
	SortOrderID                                                                                 *string                                        `json:"SortOrderID,omitempty"`
	// The sample axis value in the vertical dimension at which Depth formatted data starts. Use                                               
	// SeismicDomainType to know which of time or depth pairs is the actual range vs. what is                                                  
	// estimated.                                                                                                                              
	StartDepth                                                                                  *float64                                       `json:"StartDepth,omitempty"`
	// The sample axis value in the vertical dimension at which Time formatted data starts. Use                                                
	// SeismicDomainType to know which of time or depth pairs is the actual range vs. what is                                                  
	// estimated.                                                                                                                              
	StartTime                                                                                   *float64                                       `json:"StartTime,omitempty"`
	// Character metadata from headers inside file, such as the EBCDIC header of SEGD.  This is                                                
	// an array to capture each stanza separately.                                                                                             
	TextualFileHeader                                                                           []string                                       `json:"TextualFileHeader,omitempty"`
	// How many traces are in the volume                                                                                                       
	TraceCount                                                                                  *int64                                         `json:"TraceCount,omitempty"`
	// UOM for vertical trace domain values                                                                                                    
	TraceDomainUOM                                                                              *string                                        `json:"TraceDomainUOM,omitempty"`
	// Maximum trace length calculated using depth or time start and end points as appropriate                                                 
	// according to SeismicDomainType.                                                                                                         
	TraceLength                                                                                 *float64                                       `json:"TraceLength,omitempty"`
	// Datum value, the elevation of zero time/depth on the vertical axis in the domain of                                                     
	// SeismicDomainType relative to the vertical reference datum used (usually MSL). Positive                                                 
	// is upward from zero elevation to seismic datum).                                                                                        
	VerticalDatumOffset                                                                         *float64                                       `json:"VerticalDatumOffset,omitempty"`
	// Identifies a vertical reference datum type. E.g. mean sea level, ground level, mudline.                                                 
	VerticalMeasurementTypeID                                                                   *string                                        `json:"VerticalMeasurementTypeID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
