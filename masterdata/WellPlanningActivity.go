package masterdata

import "time"

// Information regarding an individual activity that is part of the ActivityPlan
type WellPlanningActivity struct {
	// the catalog version of this activity                                                                                       
	ActivityCodeCatalogVersion                                                                *string                             `json:"ActivityCodeCatalogVersion,omitempty"`
	// The activity code of the activity                                                                                          
	ActivityCodeID                                                                            *string                             `json:"ActivityCodeID,omitempty"`
	// Identifier of the activity                                                                                                 
	ActivityID                                                                                string                              `json:"ActivityID"`
	// The activity level of this activity. Levels go from L1 to L6                                                               
	ActivityLevel                                                                             *string                             `json:"ActivityLevel,omitempty"`
	// Free form comments associated with this activity                                                                           
	Comment                                                                                   *string                             `json:"Comment,omitempty"`
	// The depth range over which the the activity takes place                                                                    
	DepthRange                                                                                *DepthRange                         `json:"DepthRange,omitempty"`
	// Estimated duration for the planned activity.                                                                               
	EstimatedDuration                                                                         *float64                            `json:"EstimatedDuration,omitempty"`
	// Flag used to indicates this particular activity is optional in the plan.                                                   
	IsOptional                                                                                *bool                               `json:"IsOptional,omitempty"`
	// A name given to this activity                                                                                              
	Name                                                                                      *string                             `json:"Name,omitempty"`
	// The expected duration of the non productive time of the activity                                                           
	NonProductiveTimeDuration                                                                 *float64                            `json:"NonProductiveTimeDuration,omitempty"`
	// Statistics that define the non productive time of this activity                                                            
	NonProductiveTimeStatistics                                                               []NonProductiveTimeStatisticElement `json:"NonProductiveTimeStatistics,omitempty"`
	// Reference to objects that is defined within the context of a wellbore. The WITSML 1.4.1                                    
	// standard has only one object reference, but a risk may related with multiple objects.                                      
	ObjectReferenceIDs                                                                        []string                            `json:"ObjectReferenceIDs,omitempty"`
	// The parent activity to this activity                                                                                       
	ParentID                                                                                  *string                             `json:"ParentID,omitempty"`
	// DEPRECATED: Overall duration as planned for the activity                                                                   
	PlannedDuration                                                                           *time.Time                          `json:"PlannedDuration,omitempty"`
	// Date/Time the activity is planned to end                                                                                   
	PlannedEndTime                                                                            *time.Time                          `json:"PlannedEndTime,omitempty"`
	// Date/Time the activity is planned to start                                                                                 
	PlannedStartTime                                                                          *time.Time                          `json:"PlannedStartTime,omitempty"`
	// The preceding activity in the plan                                                                                         
	PredecessorsID                                                                            *string                             `json:"PredecessorsID,omitempty"`
	// The expected productive time of the activity - if "clean time probability" distribution                                    
	// is populated then this should be the expected value of the distribution                                                    
	ProductiveTimeDuration                                                                    *float64                            `json:"ProductiveTimeDuration,omitempty"`
	// Statistics that define the expected productivity time of this activity                                                     
	ProductiveTimeStatistics                                                                  []ProductiveTimeStatisticElement    `json:"ProductiveTimeStatistics,omitempty"`
	// Statistics that define the rate of penetration of this activity                                                            
	RateOfPenetrationStatistics                                                               []ROPStatistics                     `json:"RateOfPenetrationStatistics,omitempty"`
	// A reference to the object that holds the information about the risks that apply to the                                     
	// activity                                                                                                                   
	RiskIDs                                                                                   []string                            `json:"RiskIDs,omitempty"`
}
