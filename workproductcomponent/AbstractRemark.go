package workproductcomponent

// A remark object, pairing a remark text with a source, e.g. an author, and a date, which
// is typically included in an array. The RemarkSequenceNumber acts as unique key in this
// case.
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
