package workproductcomponent

// A set of properties describing a trajectory station property which is available for this
// instance of a WellboreTrajectory.
type AvailableTrajectoryStationPropertyElement struct {
	// The name of the curve (e.g. column in a CSV document) as originally found. If absent The        
	// name of the TrajectoryStationPropertyType is intended to be used.                               
	Name                                                                                       *string `json:"Name,omitempty"`
	// Unit of Measure for the station properties of type TrajectoryStationPropertyType.               
	StationPropertyUnitID                                                                      *string `json:"StationPropertyUnitID,omitempty"`
	// The reference to a trajectory station property type - or if interpreted as channels, the        
	// curve or channel name type, identifying e.g. MD, Inclination, Azimuth. This is a                
	// relationship to a reference-data--TrajectoryStationPropertyType record id.                      
	TrajectoryStationPropertyTypeID                                                            *string `json:"TrajectoryStationPropertyTypeID,omitempty"`
}
