package workproductcomponent

import "time"

// The explicit Vertical or Seismic Reference Datum for image file. Either
// EmbeddedVerticalReference or SharedVerticalReference must be populated if VerticalLevel
// is populated.
//
// A reference level or horizontal reference surface definition, which can be used embedded
// in other schemas.
type AbstractReferenceLevel struct {
	// The date and time at which this reference level instance becomes effective.                         
	EffectiveDateTime                                                                           *time.Time `json:"EffectiveDateTime,omitempty"`
	// The height above the reference surface defined by the VerticalCoordinateReferenceSystemID           
	// positive upwards.                                                                                   
	Height                                                                                      *float64   `json:"Height,omitempty"`
	// The replacement velocity value used to produce vertical static shifts in seismic data.              
	SeismicReplacementVelocity                                                                  *float64   `json:"SeismicReplacementVelocity,omitempty"`
	// The date and time at which a reference level instance is no longer in effect.                       
	TerminationDateTime                                                                         *time.Time `json:"TerminationDateTime,omitempty"`
	// The relationship to the vertical CRS defining the absolute reference surface.                       
	VerticalCoordinateReferenceSystemID                                                         *string    `json:"VerticalCoordinateReferenceSystemID,omitempty"`
	// When used in context of a Wellbore, this specifies Measured Depth, True Vertical Depth,             
	// or Elevation.                                                                                       
	VerticalMeasurementPathID                                                                   *string    `json:"VerticalMeasurementPathID,omitempty"`
	// When used in context of a Wellbore this specifies Driller vs Logger measurements.                   
	VerticalMeasurementSourceID                                                                 *string    `json:"VerticalMeasurementSourceID,omitempty"`
	// Specifies the type of vertical measurement (SRD, ES, GR, MSL,and many more).                        
	VerticalMeasurementTypeID                                                                   *string    `json:"VerticalMeasurementTypeID,omitempty"`
	// The positional uncertainty in the vertical direction.                                               
	VerticalUncertainty                                                                         *float64   `json:"VerticalUncertainty,omitempty"`
	// When used in context of a Wellbore this specifies what directional survey or wellpath was           
	// used to calculate the TVD.                                                                          
	WellboreTVDTrajectoryID                                                                     *string    `json:"WellboreTVDTrajectoryID,omitempty"`
}
