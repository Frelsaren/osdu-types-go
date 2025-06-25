package workproductcomponent

// The embedded bin grid definition describing the mapping from pixel coordinate system to
// DataCRS ('world coordinates'). BinGrid member properties are only populated if BinGridID
// is absent. A populated BinGridID overrides any embedded BinGrid values.
//
// The shared properties for a bin grid.
type AbstractBinGrid struct {
	// DEPRECATED: Use  AbstractGeoJson.PropertiesBinGridCorners properties inside the                                   
	// ABCDBinGridSpatialLocation. Previously:  Array of 4 corner points for bin grid in local                           
	// coordinates: Point A (min inline, min crossline); Point B (min inline, max crossline);                            
	// Point C (max inline, min crossline); Point D (max inline, max crossline).  If Point D is                          
	// not given and BinGridDefinitionMethodTypeID=4, it must be supplied, with its spatial                              
	// location, before ingestion to create a parallelogram in map coordinate space.  Note                               
	// correspondence of inline=x, crossline=y.                                                                          
	ABCDBinGridLocalCoordinates                                                                 []AbstractCoordinates    `json:"ABCDBinGridLocalCoordinates,omitempty"`
	// Bin Grid ABCD points containing the projected coordinates, projected CRS and quality                              
	// metadata.  This attribute is required also for the P6 definition method to define the                             
	// projected CRS, even if the ABCD coordinates would be optional (recommended to be always                           
	// calculated). It is recommended to populate the GeoJSON/AnyCrsGeoJSON with properties                              
	// according to the AbstractGeoJson.PropertiesBinGridCorners schema fragment.                                        
	ABCDBinGridSpatialLocation                                                                  *AbstractSpatialLocation `json:"ABCDBinGridSpatialLocation,omitempty"`
	// This identifies how the Bin Grid is defined:  4=ABCD four-points method was used to                               
	// define the grid (P6 parameters are optional and can contain derived values;                                       
	// P6BinNodeIncrementOnIAxis and P6BinNodeIncrementOnJaxis can be used as part of four-point                         
	// method).  Use a perspective transformation to map between map coordinates and bin                                 
	// coordinates. Note point order.  6=P6 definition method was used to define the bin grid                            
	// (ABCD points are optional and can contain derived values; ABCDBinGridSpatialLocation must                         
	// specify the projected CRS).  Use an affine transformation to map between map coordinates                          
	// and bin coordinates.                                                                                              
	BinGridDefinitionMethodTypeID                                                               *string                  `json:"BinGridDefinitionMethodTypeID,omitempty"`
	// Name of bin grid (e.g., GEOCO_GREENCYN_PHV_2012).  Probably the name as it exists in a                            
	// separate corporate store if OSDU is not main system.                                                              
	BinGridName                                                                                 *string                  `json:"BinGridName,omitempty"`
	// Type of bin grid (Acquisition, Processing, Velocity, MagGrav, Magnetics, Gravity,                                 
	// GeologicModel, Reprojected, etc.)                                                                                 
	BinGridTypeID                                                                               *string                  `json:"BinGridTypeID,omitempty"`
	// Nominal design fold as intended by the bin grid definition, expressed as the mode in                              
	// percentage points (60 fold = 6000%).                                                                              
	CoveragePercent                                                                             *float64                 `json:"CoveragePercent,omitempty"`
	// Easting coordinate of tie point (e.g., center or A point)                                                         
	P6BinGridOriginEasting                                                                      *float64                 `json:"P6BinGridOriginEasting,omitempty"`
	// Inline coordinate of tie point (e.g., center or A point)                                                          
	P6BinGridOriginI                                                                            *float64                 `json:"P6BinGridOriginI,omitempty"`
	// Crossline coordinate of tie point (e.g., center or A point)                                                       
	P6BinGridOriginJ                                                                            *float64                 `json:"P6BinGridOriginJ,omitempty"`
	// Northing coordinate of tie point (e.g., center or A point)                                                        
	P6BinGridOriginNorthing                                                                     *float64                 `json:"P6BinGridOriginNorthing,omitempty"`
	// Increment (positive integer) for the inline coordinate. If not provided then 1 is                                 
	// assumed.  The bin grid definition is expected to have increment 1 and the increment                               
	// stored with the SeismicTraceData (“inline increment”) takes precedence over the increment                         
	// set at the BinGrid.  Alternatively the increments are allowed to be defined with the                              
	// BinGrid, but this should be avoided to allow for variations in sampling in trace data                             
	// sets.                                                                                                             
	P6BinNodeIncrementOnIaxis                                                                   *int64                   `json:"P6BinNodeIncrementOnIaxis,omitempty"`
	// Increment (positive integer) for the crossline coordinate. If not provided then 1 is                              
	// assumed.  The bin grid definition is expected to have increment 1 and the increment                               
	// stored with the SeismicTraceData (“crossline increment”) takes precedence over the                                
	// increment set at the BinGrid. Alternatively the increments are allowed to be defined with                         
	// the BinGrid, but this should be avoided to allow for variations in sampling in trace data                         
	// sets.                                                                                                             
	P6BinNodeIncrementOnJaxis                                                                   *int64                   `json:"P6BinNodeIncrementOnJaxis,omitempty"`
	// Distance between two inlines at the given increment apart, e.g., 30 m with                                        
	// P6BinNodeIncrementOnIaxis=1.  Unit from projected CRS in ABCDBinGridSpatialLocation                               
	P6BinWidthOnIaxis                                                                           *float64                 `json:"P6BinWidthOnIaxis,omitempty"`
	// Distance between two crosslines at the given increment apart, e.g., 25 m with                                     
	// P6BinNodeIncrementOnJaxis=4.  Unit from projected CRS in ABCDBinGridSpatialLocation                               
	P6BinWidthOnJaxis                                                                           *float64                 `json:"P6BinWidthOnJaxis,omitempty"`
	// Clockwise angle from grid north (in projCRS) in degrees from 0 to 360 of the direction of                         
	// increasing crosslines (constant inline), i.e., of the vector from point A to B.                                   
	P6MapGridBearingOfBinGridJaxis                                                              *float64                 `json:"P6MapGridBearingOfBinGridJaxis,omitempty"`
	// Scale factor for Bin Grid.  If not provided then 1 is assumed. Unit is unity.                                     
	P6ScaleFactorOfBinGrid                                                                      *float64                 `json:"P6ScaleFactorOfBinGrid,omitempty"`
	// EPSG code: 9666 for right-handed, 1049 for left-handed.  See IOGP Guidance Note 373-07-2                          
	// and 483-6.                                                                                                        
	P6TransformationMethod                                                                      *int64                   `json:"P6TransformationMethod,omitempty"`
	// Identifier (name) of the corporate database/application that stores the source bin grid                           
	// definitions if OSDU is not main system.                                                                           
	SourceBinGridAppID                                                                          *string                  `json:"SourceBinGridAppID,omitempty"`
	// Identifier of the source bin grid as stored in a corporate database/application if OSDU                           
	// is not main system.                                                                                               
	SourceBinGridID                                                                             *int64                   `json:"SourceBinGridID,omitempty"`
}
