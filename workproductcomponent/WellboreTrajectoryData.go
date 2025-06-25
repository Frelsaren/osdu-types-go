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
type WellboreTrajectoryData struct {
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
	// The date that the survey data was acquired on the field.                                                                                
	AcquisitionDate                                                                             *time.Time                                     `json:"AcquisitionDate,omitempty"`
	// Remarks related to acquisition context which is not the same as Description which is a                                                  
	// summary of the work-product-component.                                                                                                  
	AcquisitionRemark                                                                           *string                                        `json:"AcquisitionRemark,omitempty"`
	// A flag indicating if the survey is currently active or valid within his lifecycle stage,                                                
	// not necessarily the definitive survey.                                                                                                  
	ActiveIndicator                                                                             *bool                                          `json:"ActiveIndicator,omitempty"`
	// The audit trail of operations applied to the station coordinates from the original state                                                
	// to the current state. The list may contain operations applied prior to ingestion as well                                                
	// as the operations applied to produce the Wgs84Coordinates. The text elements refer to                                                   
	// ESRI style CRS and Transformation names, which may have to be translated to EPSG standard                                               
	// names.                                                                                                                                  
	AppliedOperations                                                                           []string                                       `json:"AppliedOperations,omitempty"`
	// Date/time when the directional Survey QA/QC was performed.                                                                              
	AppliedOperationsDateTime                                                                   *time.Time                                     `json:"AppliedOperationsDateTime,omitempty"`
	// Any comments captured by the Borehole Survey specialist when performing the QA/QC work.                                                 
	AppliedOperationsRemarks                                                                    *string                                        `json:"AppliedOperationsRemarks,omitempty"`
	// Name of Borehole Survey specialist who performed the QA/QC work                                                                         
	AppliedOperationsUser                                                                       *string                                        `json:"AppliedOperationsUser,omitempty"`
	// The array of TrajectoryStationProperty definitions describing the available properties                                                  
	// for this instance of WellboreTrajectory.                                                                                                
	AvailableTrajectoryStationProperties                                                        []AvailableTrajectoryStationPropertyElement    `json:"AvailableTrajectoryStationProperties,omitempty"`
	// The North reference of the trajectory used to define the azimuth angular measurement                                                    
	// values. For example, True North, Grid North, Magnetic North.                                                                            
	AzimuthReferenceType                                                                        *string                                        `json:"AzimuthReferenceType,omitempty"`
	// Measured depth within the wellbore of the LAST surveyed station with recorded data.  If a                                               
	// stored survey has been extrapolated to a deeper depth than the last surveyed station then                                               
	// that is the extrapolated measured depth and not the survey base depth.                                                                  
	BaseDepthMeasuredDepth                                                                      float64                                        `json:"BaseDepthMeasuredDepth"`
	// Calculation Method used to calculate the Wellpath Trajectory from the directional survey                                                
	// data including TVD, X OFFSET, Y OFFSET and DOG LEG SEVERITY. Examples include Minimum                                                   
	// Curvature, Radius of Curvature, Balanced Tangential, etc.                                                                               
	CalculationMethodType                                                                       *string                                        `json:"CalculationMethodType,omitempty"`
	// The relationship to company who engaged the service company (ServiceCompanyID) to perform                                               
	// the surveying.                                                                                                                          
	CompanyID                                                                                   *string                                        `json:"CompanyID,omitempty"`
	// The date and time capturing when the last survey station was measured.                                                                  
	EndDateTime                                                                                 *time.Time                                     `json:"EndDateTime,omitempty"`
	// The measured depth to which the survey segment was extrapolated.                                                                        
	ExtrapolatedMeasuredDepth                                                                   *float64                                       `json:"ExtrapolatedMeasuredDepth,omitempty"`
	// Comment/Annotation made to WellboreTrajectory describing the projected MD Base or                                                       
	// Bottomhole, e.g., listing the Depth Reference Name, Elevation and Bottomhole MD. This can                                               
	// be used for comparison against the survey at a later date.                                                                              
	ExtrapolatedMeasuredDepthRemark                                                             *string                                        `json:"ExtrapolatedMeasuredDepthRemark,omitempty"`
	// Coordinate Reference System defining the Geodetic Datum of the station LATITUDE and                                                     
	// LONGITUDE values. If LATITUDE and LONGITUDE attributes are stored, clearly identifying                                                  
	// their Datum is required.                                                                                                                
	GeographicCRSID                                                                             *string                                        `json:"GeographicCRSID,omitempty"`
	// Geomagnetic Model Name including the applicable year used to calculate Geomagnetic field                                                
	// for a given date, coordinate and measured depth (when calculated down the wellbore).                                                    
	GeoMagneticModelID                                                                          *string                                        `json:"GeoMagneticModelID,omitempty"`
	// Identifier of the Gravity Model in use for the survey.                                                                                  
	GravityModelID                                                                              *string                                        `json:"GravityModelID,omitempty"`
	// True indicates that this trajectory is definitive for this wellbore as provided by the                                                  
	// survey contractor. False or not given indicates otherwise. There can only be one                                                        
	// trajectory per wellbore with definitive=true and it must define the geometry of the whole                                               
	// wellbore (surface to TD). The definitive trajectory may represent a composite of survey                                                 
	// data from one or more other trajectories.                                                                                               
	IsDefinitive                                                                                *bool                                          `json:"IsDefinitive,omitempty"`
	// From the survey contractors perspective providing the survey to the Operator is the                                                     
	// trajectory final (true) or intermediate/preliminary (false)? Does not mean that the                                                     
	// trajectory cannot be worked on further by the Operator.                                                                                 
	IsFinal                                                                                     *bool                                          `json:"IsFinal,omitempty"`
	// Is trajectory a result of a memory dump from a tool?                                                                                    
	IsMemory                                                                                    *bool                                          `json:"IsMemory,omitempty"`
	// Survey tool declination uncertainty.                                                                                                    
	MagneticDeclinationUncertainty                                                              *float64                                       `json:"MagneticDeclinationUncertainty,omitempty"`
	// Calculated magnetic declination used to correct a Magnetic North referenced azimuth to a                                                
	// True North azimuth. Magnetic declination angles are measured positive clockwise from True                                               
	// North to Magnetic North (or negative in the anti-clockwise direction). To convert a                                                     
	// Magnetic azimuth to a True North azimuth, the magnetic declination should be added.                                                     
	// Starting value if stations have individual values.                                                                                      
	MagneticDeclinationUsed                                                                     *float64                                       `json:"MagneticDeclinationUsed,omitempty"`
	// Calculated magnetic dip angle theoretical / reference value.                                                                            
	MagneticDIPAngleReference                                                                   *float64                                       `json:"MagneticDipAngleReference,omitempty"`
	// Survey tool dip angle uncertainty.                                                                                                      
	MagneticDIPAngleUncertainty                                                                 *float64                                       `json:"MagneticDipAngleUncertainty,omitempty"`
	// Date against Geomagnetic field calculated                                                                                               
	MagneticFieldCalculationDate                                                                *time.Time                                     `json:"MagneticFieldCalculationDate,omitempty"`
	// Coordinate Reference System defining the Projection of the station EASTING and NORTHING                                                 
	// values. If  type is "Grid North" and EASTING and NORTHING attributes are stored, clearly                                                
	// identifying their projection is required.                                                                                               
	ProjectedCRSID                                                                              *string                                        `json:"ProjectedCRSID,omitempty"`
	// Name of the Survey Company.                                                                                                             
	ServiceCompanyID                                                                            *string                                        `json:"ServiceCompanyID,omitempty"`
	// The date and time capturing when the first survey station was measured.                                                                 
	StartDateTime                                                                               *time.Time                                     `json:"StartDateTime,omitempty"`
	// The Grid Convergence angle computed at the surface position, which was used for the                                                     
	// WellboreTrajectory True North to Grid North Azimuths corrections. Only populated for                                                    
	// projected CRSs. Recommended sign convention: Gauss-Bomford.                                                                             
	SurfaceGridConvergence                                                                      *float64                                       `json:"SurfaceGridConvergence,omitempty"`
	// The Scale Factor computed for the wellbore's surface position used for calculating                                                      
	// projected Map Coordinates from survey data in WellboreTrajectory.                                                                       
	SurfaceScaleFactor                                                                          *float64                                       `json:"SurfaceScaleFactor,omitempty"`
	// Unique or Distinctive Survey Reference Number, Job Number, File Number, Identifier,                                                     
	// Label, Name, etc. as indicated on a directional survey report, file, etc.  Use this                                                     
	// attribute to allow correlation of the data in this Directional Survey back to the                                                       
	// original source document, file, etc.                                                                                                    
	SurveyReferenceIdentifier                                                                   *string                                        `json:"SurveyReferenceIdentifier,omitempty"`
	// The type of tool or equipment used to acquire this Directional Survey.  For example,                                                    
	// gyroscopic, magnetic, MWD, TOTCO, acid bottle, etc. Follow OWSG reference data and                                                      
	// support the ISCWSA survey tool definitions.                                                                                             
	SurveyToolTypeID                                                                            *string                                        `json:"SurveyToolTypeID,omitempty"`
	// The type of this directional survey.  For example a "Directional Survey" where MD,                                                      
	// Inclination and Azimuth are all measured or a "Position Log" where Inclination and                                                      
	// Azimuth are both null and only MD, TVD and X/Y Offsets are available) - or "Full Survey"                                                
	// where everything is fully filled-up, depth-inclination-azimuth.                                                                         
	SurveyType                                                                                  *string                                        `json:"SurveyType,omitempty"`
	// The version of the wellbore survey deliverable received from the service provider - as                                                  
	// given by this provider                                                                                                                  
	SurveyVersion                                                                               *string                                        `json:"SurveyVersion,omitempty"`
	// Tie-point depth measured along the wellbore from the measurement reference datum to the                                                 
	// survey station - where tie point is the place on the originating survey where the current                                               
	// survey intersect it.                                                                                                                    
	TieMeasuredDepth                                                                            *float64                                       `json:"TieMeasuredDepth,omitempty"`
	// True Vertical Depth of the TieMeasuredDepth                                                                                             
	TieTrueVerticalDepth                                                                        *float64                                       `json:"TieTrueVerticalDepth,omitempty"`
	// Measured depth in the wellbore where the directional survey starts. This should equal the                                               
	// minimum station measured depth for this directional survey, regardless of whether the                                                   
	// surveyed station represents an actual surveyed MD or not.                                                                               
	TopDepthMeasuredDepth                                                                       float64                                        `json:"TopDepthMeasuredDepth"`
	// The calculated average Tortuosity for the WellboreTrajectory.                                                                           
	Tortuosity                                                                                  *float64                                       `json:"Tortuosity,omitempty"`
	// Calculated total gravity field.                                                                                                         
	TotalGravityFieldStrength                                                                   *float64                                       `json:"TotalGravityFieldStrength,omitempty"`
	// Survey tool gravity uncertainty.                                                                                                        
	TotalGravityFieldStrengthUncertainty                                                        *float64                                       `json:"TotalGravityFieldStrengthUncertainty,omitempty"`
	// Calculated geomagnetic field theoretical/reference value.                                                                               
	TotalMagneticFieldStrength                                                                  *float64                                       `json:"TotalMagneticFieldStrength,omitempty"`
	// Survey tool magnetic uncertainty.                                                                                                       
	TotalMagneticFieldStrengthUncertainty                                                       *float64                                       `json:"TotalMagneticFieldStrengthUncertainty,omitempty"`
	// References an entry in the Vertical Measurement array for the Wellbore identified by                                                    
	// WellboreID, which defines the vertical reference datum for all survey station measured                                                  
	// depths.                                                                                                                                 
	VerticalMeasurement                                                                         AbstractFacilityVerticalMeasurement            `json:"VerticalMeasurement"`
	// Azimuth angle used for vertical section projection/computations of the survey stations.                                                 
	VerticalSectionAzimuth                                                                      *float64                                       `json:"VerticalSectionAzimuth,omitempty"`
	// Vertical Section Origin East / -West relative to the Well SHL Projected CRS 0, 0 origin.                                                
	// 0 where Vertical Section Projection originates from the wellhead.                                                                       
	VerticalSectionOriginEW                                                                     *float64                                       `json:"VerticalSectionOriginEW,omitempty"`
	// Vertical Section Origin North / -South relative to the Well SHL Projected CRS 0, 0                                                      
	// origin. 0 where Vertical Section Projection originates from the wellhead.                                                               
	VerticalSectionOriginNS                                                                     *float64                                       `json:"VerticalSectionOriginNS,omitempty"`
	// A unique name, code or number designated to the Wellbore.                                                                               
	WellboreID                                                                                  string                                         `json:"WellboreID"`
	ExtensionProperties                                                                         map[string]interface{}                         `json:"ExtensionProperties,omitempty"`
}
