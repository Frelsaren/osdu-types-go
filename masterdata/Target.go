package masterdata

// This is a geometric body describing a volume in the earth intended to be a target of one
// or more wellbores.
type Target struct {
	// A comment or description of the target                                                                           
	Comments                                                                                   *string                  `json:"Comments,omitempty"`
	// The Geometry of the target.                                                                                      
	Geometry                                                                                   *Geometry                `json:"Geometry,omitempty"`
	// The geographic location of the target                                                                            
	Location                                                                                   *AbstractSpatialLocation `json:"Location,omitempty"`
	// Human recognizable context for the target - the name must be unique within the set.                              
	Name                                                                                       *string                  `json:"Name,omitempty"`
	// The target could be Primary or Secondary                                                                         
	ObjectiveType                                                                              *string                  `json:"ObjectiveType,omitempty"`
	// Name of the parent target in this set, this one has been defined from (this represents a                         
	// pointer to the parent target. This may represents a relationship between a driller and                           
	// geological target).                                                                                              
	ParentTargetName                                                                           *string                  `json:"ParentTargetName,omitempty"`
	// Confidence factor for target                                                                                     
	TargetConfidence                                                                           *float64                 `json:"TargetConfidence,omitempty"`
	// The type of the target. Only possible values are: "Geological", "Driller" or "Unknown"                           
	Type                                                                                       *string                  `json:"Type,omitempty"`
}
