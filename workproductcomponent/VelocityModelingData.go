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
type VelocityModelingData struct {
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
	// Type of anisotropy used in the velocity model                                                                                            
	AnisotropyType                                                                               *string                                        `json:"AnisotropyType,omitempty"`
	// The average distance between grid nodes or mesh vertices in each direction (i,j,k).  Note                                                
	// that vertical case is equivalent to sample interval.                                                                                     
	AverageNodeSpacings                                                                          []float64                                      `json:"AverageNodeSpacings,omitempty"`
	// the Bin Grid of the coordinates are specified in seismic bin inline/crossline. Only                                                      
	// required when it is a Grid `PropertyFieldRepresentationType`.                                                                            
	BinGridID                                                                                    *string                                        `json:"BinGridID,omitempty"`
	// The type of numerical value(s) stored in each grid cell, such as Float or Double.                                                        
	CellValueTypes                                                                               []string                                       `json:"CellValueTypes,omitempty"`
	// Reference to history in source system, for example Jobpro jobset id, dataset id, workflow                                                
	// id                                                                                                                                       
	DataSourceReferenceKeys                                                                      []int64                                        `json:"DataSourceReferenceKeys,omitempty"`
	// System providing data source history, eg. Jobpro, etc.                                                                                   
	DataSourceSystem                                                                             *string                                        `json:"DataSourceSystem,omitempty"`
	// The number of grid nodes in each direction (i,j,k)                                                                                       
	DimensionNodeCounts                                                                          []int64                                        `json:"DimensionNodeCounts,omitempty"`
	// Is this model defined along a line, on a surface, for a volume, for a time series?                                                       
	DimensionType                                                                                *string                                        `json:"DimensionType,omitempty"`
	// Given a discretisation of a property field (e.g. a  mesh), the value of a property may                                                   
	// refer to a vertex, the center of a cell, or the region covered by a cell.  When vertical                                                 
	// interpolation is constant, this also includes an indication of Z Grid Registration, which                                                
	// whether the sample value pertains to the top, center, of bottom of grid.                                                                 
	DiscretisationSchemeType                                                                     *string                                        `json:"DiscretisationSchemeType,omitempty"`
	// Boolean to show that datum reference is not a constant.  Any description or horizon                                                      
	// information must be described in model file(s).                                                                                          
	FloatingDatumIndicator                                                                       *bool                                          `json:"FloatingDatumIndicator,omitempty"`
	// The CRS for surface coordinates used in fault locations if not specified in File.                                                        
	HorizontalCRSID                                                                              *string                                        `json:"HorizontalCRSID,omitempty"`
	// For discretely sampled models, the mathematical form of interpolation between nodes, such                                                
	// as linear in space, bicubic spline, linear in time, trilinear, horizon-based.                                                            
	InterpolationMethodID                                                                        *string                                        `json:"InterpolationMethodID,omitempty"`
	// The purpose or intended use of the velocity model, such as Stacking|Depth Migration|Time                                                 
	// Migration|Time-depth.                                                                                                                    
	ObjectiveType                                                                                *string                                        `json:"ObjectiveType,omitempty"`
	// Is the velocity field represented as a grid, mesh or profile.                                                                            
	PropertyFieldRepresentationType                                                              *string                                        `json:"PropertyFieldRepresentationType,omitempty"`
	// List of properties represented, eg. Vp, Vs, ....  Length ValuesPerNodeOrCell.                                                            
	PropertyNameTypes                                                                            []string                                       `json:"PropertyNameTypes,omitempty"`
	// Units of measure for each property type in Cell Values.  Array of length                                                                 
	// ValuesPerNodeOrCell.                                                                                                                     
	PropertyUOMIDs                                                                               []string                                       `json:"PropertyUOMIDs,omitempty"`
	// Total depth or time covered by velocity model.  In units of SeismicDomainUoM.                                                            
	RecordLength                                                                                 *float64                                       `json:"RecordLength,omitempty"`
	// Comments about the velocity model reflecting the thinking of the modeler.  Distinguished                                                 
	// from Description which is a general explanation of the model.                                                                            
	Remark                                                                                       *string                                        `json:"Remark,omitempty"`
	// Value used to produce vertical static shifts in data                                                                                     
	ReplacementVelocity                                                                          *float64                                       `json:"ReplacementVelocity,omitempty"`
	// The relationship to the Seismic2DInterpretationSet relevant to this velocity modeling                                                    
	// activity.                                                                                                                                
	Seismic2DInterpretationSetID                                                                 *string                                        `json:"Seismic2DInterpretationSetID,omitempty"`
	// The relationship to the Seismic3DInterpretationSet relevant to this velocity modeling                                                    
	// activity.                                                                                                                                
	Seismic3DInterpretationSetID                                                                 *string                                        `json:"Seismic3DInterpretationSetID,omitempty"`
	// The relationship to the SeismicAcquisitionSurvey relevant to this velocity modeling                                                      
	// activity.                                                                                                                                
	SeismicAcquisitionSurveyID                                                                   *string                                        `json:"SeismicAcquisitionSurveyID,omitempty"`
	// Vertical domain of velocities.  E.g. Time, Depth.                                                                                        
	SeismicDomainTypeID                                                                          *string                                        `json:"SeismicDomainTypeID,omitempty"`
	// Unit of measurement for vertical domain                                                                                                  
	SeismicDomainUOM                                                                             *string                                        `json:"SeismicDomainUOM,omitempty"`
	// The list of seismic line geometries holding the trace to 'world' coordinate mappings for                                                 
	// the 2D seismic lines. Only populated if the Velocity Modeling is based on 2D                                                             
	// interpretations. The `PropertyFieldRepresentationType` should be either a Mesh or Profile.                                               
	SeismicLineGeometries                                                                        []string                                       `json:"SeismicLineGeometries,omitempty"`
	// The relationship to the SeismicProcessingProject relevant to this velocity modeling                                                      
	// activity.                                                                                                                                
	SeismicProcessingProjectID                                                                   *string                                        `json:"SeismicProcessingProjectID,omitempty"`
	// Total number of vertices in the model.                                                                                                   
	TotalNodeCount                                                                               *float64                                       `json:"TotalNodeCount,omitempty"`
	// The number of numerical values stored at each node or cell                                                                               
	ValuesPerNodeOrCell                                                                          *int64                                         `json:"ValuesPerNodeOrCell,omitempty"`
	// Type of algorithm used to derive velocities such as Stacking NMO, Tomography, etc.                                                       
	VelocityAnalysisMethodID                                                                     *string                                        `json:"VelocityAnalysisMethodID,omitempty"`
	// Direction associated with the velocity.  Orientation of velocity specification such as                                                   
	// vertical, dip and azimuth.                                                                                                               
	VelocityDirectionType                                                                        *string                                        `json:"VelocityDirectionType,omitempty"`
	// Name of the Velocity Type describing the statistic represented, such as                                                                  
	// RMS|Average|Interval|Instantaneous|Stacking|Migration.                                                                                   
	VelocityType                                                                                 *string                                        `json:"VelocityType,omitempty"`
	// Datum value, the elevation of zero time/depth on the vertical axis in the domain of                                                      
	// seismicdomaintype relative to the vertical reference datum used (usually MSL). Positive                                                  
	// is upward from zero elevation to seismic datum).                                                                                         
	VerticalDatumOffset                                                                          *float64                                       `json:"VerticalDatumOffset,omitempty"`
	// Identifies a vertical reference datum type. E.g. mean sea level, ground level, mudline.                                                  
	VerticalMeasurementTypeID                                                                    *string                                        `json:"VerticalMeasurementTypeID,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
