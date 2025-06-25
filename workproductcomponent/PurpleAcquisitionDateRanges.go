package workproductcomponent

import "time"

// Defines the start and end date of the processing completed on the dataset
type PurpleAcquisitionDateRanges struct {
	// end date and time of images used  for processing             
	EndDate                                              *time.Time `json:"EndDate,omitempty"`
	// Start date and time of images used  for processing           
	StartDate                                            *time.Time `json:"StartDate,omitempty"`
}
