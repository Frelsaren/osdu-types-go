package workproductcomponent

import "time"

// The array of marker meta data in this set.
type Markers struct {
	// Name of the feature the marker is characterizing                                                   
	FeatureName                                                                                *string    `json:"FeatureName,omitempty"`
	// Feature Type Reference Type. Possible values - Base, top, fault, salt, reef, sea floor             
	FeatureTypeID                                                                              *string    `json:"FeatureTypeID,omitempty"`
	// Associated geological age                                                                          
	GeologicalAge                                                                              *string    `json:"GeologicalAge,omitempty"`
	// The optional relationship to a HorizonInterpretation, GeobodyBoundaryInterpretation or             
	// FaultInterpretation.                                                                               
	InterpretationID                                                                           *string    `json:"InterpretationID,omitempty"`
	// Timestamp of the date and time when the when the Marker was interpreted.                           
	MarkerDate                                                                                 *time.Time `json:"MarkerDate,omitempty"`
	// A unique identifier of the marker in the list of data.Markers[], ideally a UUID. If                
	// unpopulated, the string-converted element index number is used. The first index is "0".            
	// MarkerID is also used to associate MarkerProperties via the key column                             
	// data.MarkerProperties.ColumnValues[0].StringColumn[].                                              
	MarkerID                                                                                   *string    `json:"MarkerID,omitempty"`
	// The name of the Marker interpreter (could be a person or vendor).                                  
	MarkerInterpreter                                                                          *string    `json:"MarkerInterpreter,omitempty"`
	// The depth at which the Marker was noted.                                                           
	MarkerMeasuredDepth                                                                        *float64   `json:"MarkerMeasuredDepth,omitempty"`
	// Name of the Marker                                                                                 
	MarkerName                                                                                 *string    `json:"MarkerName,omitempty"`
	// Any observation number that distinguishes a Marker observation from others with same               
	// Marker name, date.                                                                                 
	MarkerObservationNumber                                                                    *float64   `json:"MarkerObservationNumber,omitempty"`
	// The Marker's TVD converted to a Sub-Sea Vertical depth, i.e., below Mean Sea Level. Note           
	// that TVD values above MSL are negative. This is the same as true vertical depth                    
	// referenced to the vertical CRS “MSL depth”.                                                        
	MarkerSubSeaVerticalDepth                                                                  *float64   `json:"MarkerSubSeaVerticalDepth,omitempty"`
	// Marker Type Reference Type. Possible values - Biostratigraphy, Lithostratigraphy,                  
	// seismic, depth of well, sequence, flow unit                                                        
	MarkerTypeID                                                                               *string    `json:"MarkerTypeID,omitempty"`
	// The geologic reason why a portion of material (typically a rock formation) is missing              
	// from the real world material/rock being measured, compared to what was expected based on           
	// offset wells. Examples: fault, unconformity, fold. This property corresponds to marker             
	// property type reference-data--MarkerPropertyType:MissingReason.                                    
	Missing                                                                                    *string    `json:"Missing,omitempty"`
	// The distance vertically below the Marker position that marks the limit of the high                 
	// confidence range for the Marker pick.                                                              
	NegativeVerticalDelta                                                                      *float64   `json:"NegativeVerticalDelta,omitempty"`
	// The distance vertically above the Marker position that marks the limit of the high                 
	// confidence range for the Marker pick.                                                              
	PositiveVerticalDelta                                                                      *float64   `json:"PositiveVerticalDelta,omitempty"`
	// Dip angle for the Wellbore Marker.                                                                 
	SurfaceDIPAngle                                                                            *float64   `json:"SurfaceDipAngle,omitempty"`
	// Dip azimuth for the Wellbore Marker.                                                               
	SurfaceDIPAzimuth                                                                          *float64   `json:"SurfaceDipAzimuth,omitempty"`
}
