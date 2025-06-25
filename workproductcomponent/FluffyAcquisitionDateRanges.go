package workproductcomponent

// Defines the start and end date of data acquisition with a remark
type FluffyAcquisitionDateRanges struct {
	// End date of the data acquisition          
	EndDate                              *string `json:"EndDate,omitempty"`
	// Optional free text description            
	Remark                               *string `json:"Remark,omitempty"`
	// Start date of the data acquisition        
	StartDate                            *string `json:"StartDate,omitempty"`
}
