package workproductcomponent

import "time"

// An interval built from two nested values : StartDate and EndDate. It applies to the whole
// log services and may apply to composite logs as [start of the first run job] and [end of
// the last run job]Log Service Date
type LogServiceDateInterval struct {
	// Date of removing tool from wellbore after the final logging run.                         
	EndDate                                                                          *time.Time `json:"EndDate,omitempty"`
	// Date of entering the wellbore with logging tools before the first logging run.           
	StartDate                                                                        *time.Time `json:"StartDate,omitempty"`
}
