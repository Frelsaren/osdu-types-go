package workproductcomponent

import "time"

// Timed comments for this job stage of the stim job.
type StimEvent struct {
	// A short description of the event.                                      
	Comment                                                        *string    `json:"Comment,omitempty"`
	// Date and time of this event.                                           
	DateTime                                                       *time.Time `json:"DateTime,omitempty"`
	// Event number.                                                          
	Number                                                         *int64     `json:"Number,omitempty"`
	// Step number. Use it to reference an existing job step entry.           
	NumberStep                                                     *int64     `json:"NumberStep,omitempty"`
}
