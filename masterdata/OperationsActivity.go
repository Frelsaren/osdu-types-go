package masterdata

import "time"

// Information regarding an individual activity that is part of the Operations Report
type OperationsActivity struct {
	// The activity code of the activity                                                                                                     
	ActivityCodeID                                                                               *string                                     `json:"ActivityCodeID,omitempty"`
	// Identifier of the activity.                                                                                                           
	ActivityID                                                                                   *string                                     `json:"ActivityID,omitempty"`
	// Measured depth at the base of interval over which the activity was conducted.                                                         
	ActivityMDBase                                                                               *float64                                    `json:"ActivityMDBase,omitempty"`
	// Measured depth at the top of interval over which the activity was conducted.                                                          
	ActivityMDTop                                                                                *float64                                    `json:"ActivityMDTop,omitempty"`
	// More detail on the outcome of the activity. For Example Injury, Operation Failed, Kick,                                               
	// Circulation Loss, Mud Loss                                                                                                            
	ActivityOutcomeDetailID                                                                      *string                                     `json:"ActivityOutcomeDetailID,omitempty"`
	// OK, Finish, interrupted, failed, etc.                                                                                                 
	ActivityOutcomeID                                                                            *string                                     `json:"ActivityOutcomeID,omitempty"`
	// TVD at the base of interval over which the activity was conducted.                                                                    
	ActivityTVDBase                                                                              *float64                                    `json:"ActivityTVDBase,omitempty"`
	// TVD at the top of interval over which the activity was conducted.                                                                     
	ActivityTVDTop                                                                               *float64                                    `json:"ActivityTVDTop,omitempty"`
	// Actual duration of the activity.                                                                                                      
	ActualDuration                                                                               *float64                                    `json:"ActualDuration,omitempty"`
	// Attachments associated with the activity                                                                                              
	AttachmentIDs                                                                                []string                                    `json:"AttachmentIDs,omitempty"`
	// Measured depth of the bit at the ase of interval over which the activity was conducted.                                               
	BitMDBase                                                                                    *float64                                    `json:"BitMDBase,omitempty"`
	// Measured depth of the bit at the top of interval over which the activity was conducted.                                               
	BitMDTop                                                                                     *float64                                    `json:"BitMDTop,omitempty"`
	// Comments and remarks.                                                                                                                 
	Comments                                                                                     *string                                     `json:"Comments,omitempty"`
	// Custom string to further define an activity.                                                                                          
	DetailActivity                                                                               *string                                     `json:"DetailActivity,omitempty"`
	// Date and time that activities ended.                                                                                                  
	EndDateTime                                                                                  *time.Time                                  `json:"EndDateTime,omitempty"`
	// Is the activity an Offline activity, Values are "true" (or "1") and "false" (or "0").                                                 
	IsOffline                                                                                    *bool                                       `json:"IsOffline,omitempty"`
	// Is the activity optimum? Values are "true" (or "1") and "false" (or "0").                                                             
	IsOptimum                                                                                    *bool                                       `json:"IsOptimum,omitempty"`
	// Does activity bring closer to objective?  Values are "true" (or "1") and "false" (or "0").                                            
	IsProductive                                                                                 *bool                                       `json:"IsProductive,omitempty"`
	// The item state for the data object. (Actual, Planned, Unknown and Modeled)                                                            
	ItemState                                                                                    *string                                     `json:"ItemState,omitempty"`
	// DEPRECATED: The measured depth to the activity/operation.  Use ActivityMDBase instead.                                                
	MeasuredDepth                                                                                *float64                                    `json:"MeasuredDepth,omitempty"`
	// link to an external object or document. For example Regulatory submission, Tour sheet.                                                
	ObjectReference                                                                              *string                                     `json:"ObjectReference,omitempty"`
	// Notes associated with the operation                                                                                                   
	OperationalNotes                                                                             *string                                     `json:"OperationalNotes,omitempty"`
	// Reference to the Organisation that represents the Operator                                                                            
	OperatorID                                                                                   *string                                     `json:"OperatorID,omitempty"`
	// The UID of the parent activity                                                                                                        
	ParentID                                                                                     *string                                     `json:"ParentID,omitempty"`
	// Phase refers to a large activity classification, e.g., drill surface hole.                                                            
	Phase                                                                                        *string                                     `json:"Phase,omitempty"`
	// The planned duration for the activity.                                                                                                
	PlannedDuration                                                                              *float64                                    `json:"PlannedDuration,omitempty"`
	// The planned hole depth at the start of the activity                                                                                   
	PlannedHoleDepthIn                                                                           *float64                                    `json:"PlannedHoleDepthIn,omitempty"`
	// The planned hole depth at the end of the activity                                                                                     
	PlannedHoleDepthOut                                                                          *float64                                    `json:"PlannedHoleDepthOut,omitempty"`
	// The UID of the preceding activity                                                                                                     
	PredecessorID                                                                                *string                                     `json:"PredecessorID,omitempty"`
	// Alternate proprietary activity code. For example contractor specific activity code                                                    
	ProprietaryActivityCode                                                                      []OperationsActivityProprietaryActivityCode `json:"ProprietaryActivityCode,omitempty"`
	// The Diameter of the section in which the activity took place                                                                          
	SectionDiameter                                                                              *float64                                    `json:"SectionDiameter,omitempty"`
	// Reference to the Organisation that represents the Service Provider                                                                    
	ServiceProviderID                                                                            *string                                     `json:"ServiceProviderID,omitempty"`
	// Date and time that activities started.                                                                                                
	StartDateTime                                                                                *time.Time                                  `json:"StartDateTime,omitempty"`
	// The target depth of the activity                                                                                                      
	TargetDepth                                                                                  *float64                                    `json:"TargetDepth,omitempty"`
	// A pointer to the tubular object related to this activity. Not the Hole Section that you                                               
	// are operating within.                                                                                                                 
	TubularID                                                                                    *string                                     `json:"TubularID,omitempty"`
	// DEPRECATED: True vertical depth to the activity/operation. Use ActivityTVDBase instead.                                               
	Tvd                                                                                          *float64                                    `json:"TVD,omitempty"`
	// Classifier (planned, unplanned, downtime).                                                                                            
	TypeActivityClassID                                                                          *string                                     `json:"TypeActivityClassID,omitempty"`
	// A Reference to the wellbore in which the activities take place.                                                                       
	WellboreID                                                                                   *string                                     `json:"WellboreID,omitempty"`
}
