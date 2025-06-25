package masterdata

// An array contains narrative remarks pertaining to a core.
type CoreRemark struct {
	// A descriptive comment for this remark.                                   
	Remark                                                              *string `json:"Remark,omitempty"`
	// A unique identifier for each remark record.                              
	RemarkID                                                            *string `json:"RemarkID,omitempty"`
	// The person, vendor, interpreter or other provider of information.        
	RemarkSource                                                        *string `json:"RemarkSource,omitempty"`
}
