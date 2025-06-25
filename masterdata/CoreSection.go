package masterdata

// An array of core sections.
type CoreSection struct {
	// The measured depth at the base of this section, typically driller's depth.           
	BottomDepth                                                                    *float64 `json:"BottomDepth,omitempty"`
	// A unique core section identifier, typically for human consumption.                   
	CoreSectionID                                                                  *string  `json:"CoreSectionID,omitempty"`
	// The length of core section recovered.                                                
	RecoveredLength                                                                *float64 `json:"RecoveredLength,omitempty"`
	// The measured depth at the top of this section, typically in driller's depth.         
	TopDepth                                                                       *float64 `json:"TopDepth,omitempty"`
}
