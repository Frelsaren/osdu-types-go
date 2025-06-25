package masterdata

import "time"

// DEPRECATED: Use data.OperationsReport instead. Information regarding an individual
// activity that is part of the Drilling Report
type DrillingActivity struct {
	// DEPRECATED: Use data.OperationsActivity.ActivityCodeID instead. The activity code of the                                        
	// activity                                                                                                                        
	ActivityCodeID                                                                              *string                                `json:"ActivityCodeID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityID instead. Identifier of the activity                                          
	ActivityID                                                                                  *string                                `json:"ActivityID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityMDBase instead. Measured depth at the                                           
	// base of interval over which the activity was conducted.                                                                         
	ActivityMDBase                                                                              *float64                               `json:"ActivityMDBase,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityMDTop instead. Measured depth at the top                                        
	// of interval over which the activity was conducted.                                                                              
	ActivityMDTop                                                                               *float64                               `json:"ActivityMDTop,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityOutcomeDetailID instead. More detail on                                         
	// the outcome of the activity. For Example Injury, Operation Failed, Kick, Circulation                                            
	// Loss, Mud Loss                                                                                                                  
	ActivityOutcomeDetailID                                                                     *string                                `json:"ActivityOutcomeDetailID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityOutcomeID instead. OK, Finish,                                                  
	// interrupted, failed, etc.                                                                                                       
	ActivityOutcomeID                                                                           *string                                `json:"ActivityOutcomeID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityTVDBase instead. TVD at the base of                                             
	// interval over which the activity was conducted.                                                                                 
	ActivityTVDBase                                                                             *float64                               `json:"ActivityTVDBase,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ActivityTVDTop instead. TVD at the top of                                               
	// interval over which the activity was conducted.                                                                                 
	ActivityTVDTop                                                                              *float64                               `json:"ActivityTVDTop,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.AttachmentIDs instead. Attachments associated                                           
	// with the activity                                                                                                               
	AttachmentIDs                                                                               []string                               `json:"AttachmentIDs,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.BitMDBase instead. Measured depth of the bit at                                         
	// the ase of interval over which the activity was conducted.                                                                      
	BitMDBase                                                                                   *float64                               `json:"BitMDBase,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.BitMDTop instead. Measured depth of the bit at                                          
	// the top of interval over which the activity was conducted.                                                                      
	BitMDTop                                                                                    *float64                               `json:"BitMDTop,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.Comments instead. Comments and remarks.                                                 
	Comments                                                                                    *string                                `json:"Comments,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.DetailActivity instead. Custom string to further                                        
	// define an activity.                                                                                                             
	DetailActivity                                                                              *string                                `json:"DetailActivity,omitempty"`
	// DEPRECATED: Transform the string value to a number and assign it to                                                             
	// data.OperationsActivity.ActualDuration (and assign the actual unit in meta[]). The                                              
	// activity duration (commonly in hours).                                                                                          
	Duration                                                                                    *time.Time                             `json:"Duration,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.EndDateTime instead. Date and time that                                                 
	// activities ended.                                                                                                               
	EndDateTime                                                                                 *time.Time                             `json:"EndDateTime,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.IsOffline instead. Is the activity an Offline                                           
	// activity, Values are "true" (or "1") and "false" (or "0").                                                                      
	IsOffline                                                                                   *bool                                  `json:"IsOffline,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.IsOptimum instead. Is the activity optimum?                                             
	// Values are "true" (or "1") and "false" (or "0").                                                                                
	IsOptimum                                                                                   *bool                                  `json:"IsOptimum,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.IsProductive instead. Does activity bring closer                                        
	// to objective?  Values are "true" (or "1") and "false" (or "0").                                                                 
	IsProductive                                                                                *bool                                  `json:"IsProductive,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ItemState instead. The item state for the data                                          
	// object. (Actual, Planned, Unknown and Modeled)                                                                                  
	ItemState                                                                                   *string                                `json:"ItemState,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.MeasuredDepth instead. The measured depth to the                                        
	// drilling activity/operation.                                                                                                    
	MeasuredDepth                                                                               *float64                               `json:"MeasuredDepth,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ObjectReference instead. link to an external                                            
	// object or document. For example Regulatory submission, Tour sheet.                                                              
	ObjectReference                                                                             *string                                `json:"ObjectReference,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.OperationalNotes instead. Notes associated with                                         
	// the operation                                                                                                                   
	OperationalNotes                                                                            *string                                `json:"OperationalNotes,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.OperatorID instead. Reference to the Organisation                                       
	// that represents the Operator                                                                                                    
	OperatorID                                                                                  *string                                `json:"OperatorID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ParentID instead. The UID of the parent activity                                        
	ParentID                                                                                    *string                                `json:"ParentID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.Phase instead. Phase refers to a large activity                                         
	// classification, e.g., drill surface hole.                                                                                       
	Phase                                                                                       *string                                `json:"Phase,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.PlannedDuration instead. The planned duration for                                       
	// the activity.                                                                                                                   
	PlannedDuration                                                                             *float64                               `json:"PlannedDuration,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.PlannedHoleDepthIn  instead. The planned hole                                           
	// depth at the start of the activity                                                                                              
	PlannedHoleDepthIn                                                                          *float64                               `json:"PlannedHoleDepthIn,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.PlannedHoleDepthOut instead. The planned hole                                           
	// depth at the end of the activity                                                                                                
	PlannedHoleDepthOut                                                                         *float64                               `json:"PlannedHoleDepthOut,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.PredecessorID instead. The UID of the preceding                                         
	// activity                                                                                                                        
	PredecessorID                                                                               *string                                `json:"PredecessorID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ProprietaryActivityCode instead. Alternate                                              
	// proprietary activity code. For example contractor specific activity code                                                        
	ProprietaryActivityCode                                                                     []DrillActivityProprietaryActivityCode `json:"ProprietaryActivityCode,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.SectionDiameter instead. The Diameter of the                                            
	// section in which the activity took place                                                                                        
	SectionDiameter                                                                             *float64                               `json:"SectionDiameter,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.ServiceProviderID instead. Reference to the                                             
	// Organisation that represents the Service Provider                                                                               
	ServiceProviderID                                                                           *string                                `json:"ServiceProviderID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.StartDateTime instead. Date and time that                                               
	// activities started.                                                                                                             
	StartDateTime                                                                               *time.Time                             `json:"StartDateTime,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.TargetDepth instead. The target depth of the                                            
	// activity                                                                                                                        
	TargetDepth                                                                                 *float64                               `json:"TargetDepth,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.TubularID instead. A pointer to the tubular                                             
	// object related to this activity. Not the Hole Section that you are operating within.                                            
	TubularID                                                                                   *string                                `json:"TubularID,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.TVD instead. True vertical depth to the drilling                                        
	// activity/operation.                                                                                                             
	Tvd                                                                                         *float64                               `json:"TVD,omitempty"`
	// DEPRECATED: Use data.OperationsActivity.TypeActivityClassID instead. Classifier (planned,                                       
	// unplanned, downtime).                                                                                                           
	TypeActivityClassID                                                                         *string                                `json:"TypeActivityClassID,omitempty"`
}
