package masterdata

import "time"

// Association of a rig to a particular well and well activity.
type RigAssignment struct {
	// The end time for this rig assignment to the well activity             
	EndDateTime                                                   *time.Time `json:"EndDateTime,omitempty"`
	// Remarks related to this rig assignment                                
	Remark                                                        *string    `json:"Remark,omitempty"`
	// A link to the Rig                                                     
	RigID                                                         *string    `json:"RigID,omitempty"`
	// The start time for this rig assignment to the well activity           
	StartDateTime                                                 *time.Time `json:"StartDateTime,omitempty"`
}
