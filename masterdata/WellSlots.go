package masterdata

// List of slots for the WellSiteStructure. A slot is a circular opening on a platform
// wellhead deck or subsea template from which to construct a well.
type WellSlots struct {
	// Remarks                                                                                                      
	Remarks                                                                            *string                      `json:"Remarks,omitempty"`
	// Slot Identifier, may be used by a Well to associate to a Slot                                                
	SlotID                                                                             *string                      `json:"SlotID,omitempty"`
	// Slot Local East/-West Coordinate offset from WellSiteStructure Centre location                               
	SlotLocalEWOffset                                                                  *float64                     `json:"SlotLocalEWOffset,omitempty"`
	// Slot Local North/-South Coordinate offset from WellSiteStructure Centre location                             
	SlotLocalNSOffset                                                                  *float64                     `json:"SlotLocalNSOffset,omitempty"`
	// Slot Location                                                                                                
	SlotLocation                                                                       *AbstractSpatialLocation     `json:"SlotLocation,omitempty"`
	// Slot Local Uncertainty within the Wellsite Structure                                                         
	SlotLocationUncertainty                                                            *AbstractLocationUncertainty `json:"SlotLocationUncertainty,omitempty"`
	// Slot Name or Number                                                                                          
	SlotName                                                                           *string                      `json:"SlotName,omitempty"`
}
