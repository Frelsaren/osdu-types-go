package masterdata

// Describes the start and end date range of the survey acquisition
type AcquisitionDateRange struct {
	// The end date of the acquisition survey          
	EndDate                                    *string `json:"EndDate,omitempty"`
	// The start date of the acquisition survey        
	StartDate                                  *string `json:"StartDate,omitempty"`
}
