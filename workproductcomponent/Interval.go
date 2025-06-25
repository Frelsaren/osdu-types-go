package workproductcomponent

// An interval given either by relationships to top/base markers or standalone top/base
// depths. To avoid confusion about entry and exit depths the naming convention uses Start
// as the entry point of the well path into the geologic unit and Stop as the exit point.
// For unfolded geologic units and vertical wellbores Start is Top and Stop is Base.
type Interval struct {
	// An array of StratigraphicUnitInterpretation, GeobodyInterpretation or                             
	// RockFluidUnitInterpretation record Ids associated to this interval.                               
	GeologicUnitInterpretationIDs                                                               []string `json:"GeologicUnitInterpretationIDs,omitempty"`
	// The unique identifier of the interval array member in the data.Intervals[] array. Ideally         
	// a UUID. IntervalID is also used to associate IntervalProperties via the key column                
	// data.IntervalProperties.ColumnValues[0].StringColumn[].                                           
	IntervalID                                                                                  *string  `json:"IntervalID,omitempty"`
	// The optional relationship to a HorizonInterpretation, GeobodyBoundaryInterpretation or            
	// FaultInterpretation. If the interval start (typically the top) is associated with a               
	// marker, this is considered a denormalization of the data.Markers[].InterpretationID for           
	// the data.Markers[].MarkerID equals to StartMarkerID.                                              
	StartBoundaryInterpretationID                                                               *string  `json:"StartBoundaryInterpretationID,omitempty"`
	// Name of the interval start (typically the top); when associated with a marker in a                
	// WellboreMarkerSet then this name is a denormalization of data.Markers[].MarkerName where          
	// the data.Markers[].MarkerID equals to StartMarkerID.                                              
	StartIntervalName                                                                           *string  `json:"StartIntervalName,omitempty"`
	// Individual markers are not globally identifiable. TopMarkerID is the unique id (MarkerID)         
	// of the top interval marker (typically the interval base) in the data.Markers[] array              
	// where the data.Markers[].MarkerID equals to StartMarkerID.                                        
	StartMarkerID                                                                               *string  `json:"StartMarkerID,omitempty"`
	// Optional reference to the WellboreMarkerSet containing the interval start (typically the          
	// top), with MarkerID equals StartMarkerID.                                                         
	StartMarkerSetID                                                                            *string  `json:"StartMarkerSetID,omitempty"`
	// The minimal MeasuredDepth of the interval. In the most common case this is the top.  If           
	// this value is associated with a marker then this value is a denormalization of                    
	// data.Markers[].MarkerMeasuredDepth where the data.Markers[].MarkerID equals to                    
	// StartMarkerID.                                                                                    
	StartMeasuredDepth                                                                          *float64 `json:"StartMeasuredDepth,omitempty"`
	// True vertical depth sub-sea of the start of the interval. This is the same as true                
	// vertical depth referenced to the vertical CRS "MSL depth". If the start of the interval           
	// is associated with a marker then this value is a denormalization of                               
	// data.Markers[].MarkerSubSeaVerticalDepth where the data.Markers[].MarkerID equals to              
	// StartMarkerID.                                                                                    
	StartSubSeaVerticalDepth                                                                    *float64 `json:"StartSubSeaVerticalDepth,omitempty"`
	// The optional relationship to a HorizonInterpretation, GeobodyBoundaryInterpretation or            
	// FaultInterpretation. If the interval stop (typically the base) is associated with a               
	// marker, this is considered a denormalization of the data.Markers[].InterpretationID where         
	// the data.Markers[].MarkerID equals to StopMarkerID.                                               
	StopBoundaryInterpretationID                                                                *string  `json:"StopBoundaryInterpretationID,omitempty"`
	// Name of the interval stop (typically the base); when associated with a marker in a                
	// WellboreMarkerSet then this name is a denormalization of data.Markers[].MarkerName where          
	// the data.Markers[].MarkerID equals to StopMarkerID.                                               
	StopIntervalName                                                                            *string  `json:"StopIntervalName,omitempty"`
	// Individual markers are not globally identifiable. StopMarkerID is the unique id                   
	// (MarkerID) of the interval stop (typically the interval base) in the data.Markers[] array         
	// where the data.Markers[].MarkerID equals to StopMarkerID.                                         
	StopMarkerID                                                                                *string  `json:"StopMarkerID,omitempty"`
	// Optional reference to the WellboreMarkerSet containing the top with MarkerID equals               
	// StopMarkerID.                                                                                     
	StopMarkerSetID                                                                             *string  `json:"StopMarkerSetID,omitempty"`
	// The maximum MeasuredDepth of the interval (typically the base). If the interval stop is           
	// associated with a marker then this value is a denormalization of                                  
	// data.Markers[].MarkerMeasuredDepth where the data.Markers[].MarkerID equals to                    
	// StopMarkerID.                                                                                     
	StopMeasuredDepth                                                                           *float64 `json:"StopMeasuredDepth,omitempty"`
	// True vertical depth sub-sea of the interval stop (typically the base). This is the same           
	// as true vertical depth referenced to the vertical CRS "MSL depth". If the interval stop           
	// is associated with a marker then this value is a denormalization of                               
	// data.Markers[].MarkerSubSeaVerticalDepth where the data.Markers[].MarkerID equals to              
	// StopMarkerID.                                                                                     
	StopSubSeaVerticalDepth                                                                     *float64 `json:"StopSubSeaVerticalDepth,omitempty"`
}
