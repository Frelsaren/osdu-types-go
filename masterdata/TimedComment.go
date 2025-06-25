package masterdata

import "time"

// Information about a series of drilling activities.
type TimedComment struct {
	// A comment that has been recorded at a particular time           
	Comment                                                 *string    `json:"Comment,omitempty"`
	// The time that the comment was made                              
	CommentTime                                             *time.Time `json:"CommentTime,omitempty"`
}
