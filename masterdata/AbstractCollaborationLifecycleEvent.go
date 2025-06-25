package masterdata

import "time"

// An event in the lifecycle of a CollaborationProject containing references to the
// CollaborationProjectCollections involved at the time of the event.
type AbstractCollaborationLifecycleEvent struct {
	// The date and time of the event.                                                                
	DateTime                                                                               *time.Time `json:"DateTime,omitempty"`
	// A unique identifier of this event element in a sequence of events as, e.g., in                 
	// CollaborationProject.                                                                          
	EventID                                                                                *string    `json:"EventID,omitempty"`
	// A user-given event name.                                                                       
	Name                                                                                   *string    `json:"Name,omitempty"`
	// A user defined remark associated with this event.                                              
	Remark                                                                                 *string    `json:"Remark,omitempty"`
	// The relationship to a CollaborationProjectCollection containing the System of Record           
	// resources used by the project.                                                                 
	ResourceCollectionID                                                                   *string    `json:"ResourceCollectionID,omitempty"`
	// The relationship to a CollaborationProjectCollection containing the Work in Progress           
	// resources published in the context of the project.                                             
	WIPResourceCollectionID                                                                *string    `json:"WIPResourceCollectionID,omitempty"`
}
