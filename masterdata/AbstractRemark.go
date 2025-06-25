package masterdata

// A remark object, pairing a remark text with a source, e.g. an author, and a date, which
// is typically included in an array. The RemarkSequenceNumber acts as unique key in this
// case.
//
// This is used to capture information regarding the methodology used in correcting rates
// acquired during the sample acquisition event. The property is only used in conjunction
// with SeparatorSampleAcquisition
//
// Remarks or comments about this sample container.
type AbstractRemark struct {
	// A descriptive comment for this remark.                                                          
	Remark                                                                                     *string `json:"Remark,omitempty"`
	// The date the remark was issued.                                                                 
	RemarkDate                                                                                 *string `json:"RemarkDate,omitempty"`
	// A unique identifier for each remark record. This identifier is used to uniquely identify        
	// a member in an array of remarks.                                                                
	RemarkSequenceNumber                                                                       *int64  `json:"RemarkSequenceNumber,omitempty"`
	// The person, vendor, interpreter or other provider of information.                               
	RemarkSource                                                                               *string `json:"RemarkSource,omitempty"`
}
