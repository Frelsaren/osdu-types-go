package workproductcomponent

import "time"

// ID to the Zero Depth Point Vertical Measure elevation for depths contained in the Cement
// Job, Stages and Pumping Schedule, depth correction used to correlate MDs to original
// drilling rig MD. References an entry in the Vertical Measurement array for the Well
// parented by the wellbore via WellboreID.
//
// A location along a wellbore, _usually_ associated with some aspect of the drilling of the
// wellbore, but not with any intersecting _subsurface_ natural surfaces.
//
// The well vertical measurement elevation reference for  test MD and TVD.
//
// The Vertical Measurement for the Wellbore identified which defines the vertical reference
// pressure datum for the related PPFG curves in this data set. The pressure datum is used
// to calculate the average pressure gradient in
//
// The data structure defining the zero depth point for the measured depths as well as the
// VerticalMeasurementSource (driller's depth or logger's depth), so that it becomes clear
// whether or not a depth correction is required and if so, in which direction the
// correction has to be performed (driller's depth to logger's depth or vice versa).
//
// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference which defines the vertical reference datum
// for all measured depths of the RockSample record. If this is not populated, the
// VerticalMeasurement is derived from the Coring.
//
// ID to the Zero Depth Point Vertical Measure elevation for depths contained in the
// stimulation job, job stages and other objects used to correlate MDs to original drilling
// rig MD. References an entry in the Vertical Measurement array for the Well parented by
// the wellbore via WellboreID.
//
// Either a self-contained vertical reference for the depths in this TubularAssembly or a
// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the
// entity defined by VerticalReferenceEntityID.
//
// Populated only if the WellLog represents time-depth relationships or check shots. It is
// expressed via the standard AbstractFacilityVerticalMeasurement. The following properties
// are expected to be present: VerticalMeasurementPathID (typically elevation),
// VerticalMeasurementTypeID as SeismicReferenceDatum, VerticalMeasurement holding the
// offset to either the VerticalCRSID or the chained VerticalReferenceID in the parent
// Wellbore.
//
// The vertical measurement reference for the log curves, which defines the vertical
// reference datum for the logged depths. Either VerticalMeasurement or
// VerticalMeasurementID are populated.
//
// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference elevation for all measured depths within
// the NPT record. If this is not populated, the VerticalMeasurement is derived from the
// Wellbore default Vertical Measure Elevation.
//
// The vertical measurement reference for this well logging acquisition activity. This
// object defines the vertical reference datum for the measured depths.
//
// References an entry in the Vertical Measurement array for the Wellbore identified by
// WellboreID, which defines the vertical reference datum for all marker measured depths of
// the WellboreIntervalSet Intervals array. It is strongly recommended specifying the
// VerticalMeasurement.WellboreTVDTrajectoryID  when SubSeaVerticalDepth are populated for
// the intervals.
//
// References an entry in the Vertical Measurement array for the Wellbore identified by
// WellboreID, which defines the vertical reference datum for all marker measured depths of
// the Wellbore Marker Set Markers array.
//
// References an entry in the Vertical Measurement array for the Wellbore identified by
// WellboreID, which defines the vertical reference datum for all survey station measured
// depths.
type AbstractFacilityVerticalMeasurement struct {
	// The date and time at which a vertical measurement instance becomes effective.                        
	EffectiveDateTime                                                                            *time.Time `json:"EffectiveDateTime,omitempty"`
	// The date and time at which a vertical measurement instance is no longer in effect.                   
	TerminationDateTime                                                                          *time.Time `json:"TerminationDateTime,omitempty"`
	// A vertical coordinate reference system defines the origin for height or depth values. It             
	// is expected that either VerticalCRSID or VerticalReferenceID reference is provided in a              
	// given vertical measurement array object, but not both.                                               
	VerticalCRSID                                                                                *string    `json:"VerticalCRSID,omitempty"`
	// The value of the elevation or depth. Depth is positive downwards from a vertical                     
	// reference or geodetic datum along a path, which can be vertical; elevation is positive               
	// upwards from a geodetic datum along a vertical path. Either can be negative.                         
	VerticalMeasurement                                                                          *float64   `json:"VerticalMeasurement,omitempty"`
	// Text which describes a vertical measurement in detail.                                               
	VerticalMeasurementDescription                                                               *string    `json:"VerticalMeasurementDescription,omitempty"`
	// Specifies Measured Depth, True Vertical Depth, or Elevation.                                         
	VerticalMeasurementPathID                                                                    *string    `json:"VerticalMeasurementPathID,omitempty"`
	// Specifies Driller vs Logger.                                                                         
	VerticalMeasurementSourceID                                                                  *string    `json:"VerticalMeasurementSourceID,omitempty"`
	// Specifies the type of vertical measurement (TD, Plugback, Kickoff, Drill Floor, Rotary               
	// Table...).                                                                                           
	VerticalMeasurementTypeID                                                                    *string    `json:"VerticalMeasurementTypeID,omitempty"`
	// The unit of measure for the vertical measurement. If a unit of measure and a vertical CRS            
	// are provided, the unit of measure provided is taken over the unit of measure from the CRS.           
	VerticalMeasurementUnitOfMeasureID                                                           *string    `json:"VerticalMeasurementUnitOfMeasureID,omitempty"`
	// This relationship identifies the entity (aka record) in which the VerticalReferenceID is             
	// found; It could be a different OSDU entity or a self-reference. For example, a Wellbore              
	// VerticalMeasurement may reference a member of a VerticalMeasurements[] array in its                  
	// parent Well record. Alternatively, VerticalReferenceEntityID may be populated with the ID            
	// of its own Wellbore record to make explicit that VerticalReferenceID is intended to be               
	// found in this record, not another.                                                                   
	VerticalReferenceEntityID                                                                    *string    `json:"VerticalReferenceEntityID,omitempty"`
	// The reference point from which the relative vertical measurement is made. This is only               
	// populated if the measurement has no VerticalCRSID specified. The value entered must match            
	// the VerticalMeasurementID for another vertical measurement array element in Wellbore or              
	// Well or in a related parent facility. The relationship should be  declared explicitly in             
	// VerticalReferenceEntityID. Any chain of measurements must ultimately resolve to a                    
	// Vertical CRS. It is expected that a VerticalCRSID or a VerticalReferenceID is provided in            
	// a given vertical measurement array object, but not both.                                             
	VerticalReferenceID                                                                          *string    `json:"VerticalReferenceID,omitempty"`
	// Specifies what directional survey or wellpath was used to calculate the TVD.                         
	WellboreTVDTrajectoryID                                                                      *string    `json:"WellboreTVDTrajectoryID,omitempty"`
}
