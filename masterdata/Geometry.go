package masterdata

// The Geometry of the target.
type Geometry struct {
	// Angle of dip of target reference plane with respect to horizontal                                         
	DIP                                                                                        *float64          `json:"Dip,omitempty"`
	// Direction of dip of target reference plane with respect to north azimuth reference                        
	DIPAzimuth                                                                                 *float64          `json:"DipAzimuth,omitempty"`
	// The geometry sections which define the spatial geometry ot the target                                     
	GeometrySections                                                                           []GeometrySection `json:"GeometrySections,omitempty"`
	// The measured length of a rectangular target along the shortest side                                       
	Length                                                                                     *float64          `json:"Length,omitempty"`
	// Measured length of the target along Target Ref Azimuth                                                    
	MajorAxis                                                                                  *float64          `json:"MajorAxis,omitempty"`
	// Measured length of the target perpendicular to Target Ref Azimuth, assumed to be on                       
	// target dip, strike, rotation plane.                                                                       
	MinorAxis                                                                                  *float64          `json:"MinorAxis,omitempty"`
	// Radius of arc for continuous curve segment. Center assumed to be offset from start point                  
	// normal to its beginning azimuth; positive value is to right, negative to left                             
	Radius                                                                                     *float64          `json:"Radius,omitempty"`
	// The shape of the target. Examples of this are Elliptical, Point etc                                       
	Shape                                                                                      *string           `json:"Shape,omitempty"`
	// Depth of target bottom surface below reference plane, measured normal to dip plane                        
	ThicknessDown                                                                              *float64          `json:"ThicknessDown,omitempty"`
	// Height of target top surface above reference plane, measured normal to dip plane                          
	ThicknessUp                                                                                *float64          `json:"ThicknessUp,omitempty"`
	// The measured length of a rectangular target along the longest side                                        
	Width                                                                                      *float64          `json:"Width,omitempty"`
}
