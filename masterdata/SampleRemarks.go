package masterdata

// An array containing narrative remarks pertaining to a rock sample.
type SampleRemarks struct {
	// A descriptive comment for this remark.                                   
	Remark                                                              *string `json:"Remark,omitempty"`
	// A unique identifier for each remark record.                              
	RemarkID                                                            *string `json:"RemarkID,omitempty"`
	// The person, vendor, interpreter or other provider of information.        
	RemarkSource                                                        *string `json:"RemarkSource,omitempty"`
}
