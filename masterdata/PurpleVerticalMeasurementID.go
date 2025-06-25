package masterdata

import "time"

// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference elevation for all measured depths within
// the BHA Run record. If this is not populated, the VerticalMeasurement is derived from the
// Wellbore default Vertical Measure Elevation.
//
// A location along a wellbore, _usually_ associated with some aspect of the drilling of the
// wellbore, but not with any intersecting _subsurface_ natural surfaces.
//
// Either a self-contained vertical reference for the depths in this blowout preventer or a
// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the
// entity defined by VerticalReferenceEntityID e.g. the parent Well.
//
// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference which defines the vertical reference datum
// for all measured depths of the Coring record.
//
// The vertical measurement reference for a wellbore or borehole acquisition platform,
// defining the vertical reference datum, elevation and relevant depths.
//
// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference which defines the vertical reference datum
// for all measured depths of the HoleSection record. If this is not populated, the
// VerticalMeasurement is derived from the Wellbore.
//
// The vertical measurement reference for the interval tops and bases. Either
// VerticalMeasurement with supplementing type properties or VerticalMeasurementID (an
// external vertical reference defined in the object VerticalReferenceEntityID) are
// populated.
//
// ID to the Zero Depth Point elevation for depths contained in the perf job incl intervals,
// depth correction used to correlate MDs to original drilling rig MD. References an entry
// in the Vertical Measurement array for the Well parented by the wellbore via WellboreID.
//
// References an entry in the VerticalMeasurements array for the Wellbore identified by
// WellboreID, or a standalone vertical reference which defines the vertical reference datum
// for all measured depths of the RockSample record. If this is not populated, the
// VerticalMeasurement is derived from the Coring.
//
// Information on the list of all depths and elevations pertaining to the target wellbore
// involved in the Sample Acquisition event, like, plug back measured depth, total measured
// depth, KB elevation. The property is always used except with WellheadSampleAcquisition,
// SeparatorSampleAcquisition, FacilitySampleAcquisition
//
// The vertical measurement reference for VSP surveys, which defines the vertical reference
// datum for the measured depths.
//
// References an entry in the VerticalMeasurements array of the Rig, Well or Wellbore
// identified by VerticalReferenceEntityID or a standalone vertical reference which defines
// the vertical reference datum for all measured depths of the SurveyProgram record. For
// planned SurveyPrograms, this property may be absent; then depths are relative to Planned
// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds
// the depth reference in data.VerticalMeasurement.
//
// Either a self-contained vertical reference for the depths in this TubularAssembly or a
// reference (VerticalReferenceID) to an element in data.VerticalMeasurements[] in the
// entity defined by VerticalReferenceEntityID e.g. the parent Well.
//
// The measured elevation from a known reference datum to a permanent geodetic datum.  This
// measurement is the foundation for calculating and correlating depths from geodetic
// datums, which in turn allow depth correlation between wellbores.
//
// The vertical measurement reference for this well logging acquisition activity. This
// object defines the vertical reference datum for the measured depths.
//
// The zero depth point (ZDP) definition for the all measured depths related to this
// WellPlanningWellbore.
//
// The vertical measurement reference for this well testing acquisition activity. This
// object defines the vertical reference datum for the measured depths.
//
// The vertical measurement reference for the interval top and base. Either
// VerticalMeasurement with supplementing type properties or VerticalMeasurementID (an
// external vertical reference defined in the object VerticalReferenceEntityID) are
// populated.
type PurpleVerticalMeasurementID struct {
	// The relationship to the rig, which was used while this vertical measurement was in active            
	// use.                                                                                                 
	RigID                                                                                        *string    `json:"RigID,omitempty"`
	// The ID for a distinct vertical measurement within the Wellbore VerticalMeasurements array            
	// so that it may be referenced by other vertical measurements if necessary.                            
	VerticalMeasurementID                                                                        *string    `json:"VerticalMeasurementID,omitempty"`
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
