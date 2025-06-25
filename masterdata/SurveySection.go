package masterdata

// Description of the Survey Section Component
type SurveySection struct {
	// Comments and remarks.                                                                            
	Comments                                                                                   *string  `json:"Comments,omitempty"`
	// Hole measured depth at which the survey run will end/ended. Depth relative to Planned            
	// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds          
	// the depth reference in data.VerticalMeasurement.                                                 
	EndMeasuredDepth                                                                           *float64 `json:"EndMeasuredDepth,omitempty"`
	// Error model used to calculate ellipses of uncertainty.                                           
	ErrorModel                                                                                 *string  `json:"ErrorModel,omitempty"`
	// Maximum allowable depth frequency for survey stations for this survey run.                       
	FrequencyMX                                                                                *float64 `json:"FrequencyMx,omitempty"`
	// Normally true, higher index trajectory takes precedence over overlapping section of              
	// previous trajectory.  Values are "true" (or "1") and "false" (or "0").                           
	IsOverwrite                                                                                *bool    `json:"IsOverwrite,omitempty"`
	// Name of survey program section.                                                                  
	Name                                                                                       *string  `json:"Name,omitempty"`
	// Order in which program sections will be executed / were executed.                                
	Sequence                                                                                   float64  `json:"Sequence"`
	// Hole measured depth at which the survey run will begin/began. Depth relative to Planned          
	// wellbore ZDP. Navigate via WellboreID to the side-car WellPlanningWellbore, which holds          
	// the depth reference in data.VerticalMeasurement.                                                 
	StartMeasuredDepth                                                                         *float64 `json:"StartMeasuredDepth,omitempty"`
	// Company who will run / has run survey tool.                                                      
	SurveyCompanyID                                                                            *string  `json:"SurveyCompanyID,omitempty"`
	// The item state for the data object.                                                              
	SurveySectionExistenceKind                                                                 string   `json:"SurveySectionExistenceKind"`
	// Name of survey tool, as defined by the manufacturer, to be used / used for this section.         
	SurveyToolName                                                                             *string  `json:"SurveyToolName,omitempty"`
	// Standardized Type of tool used.                                                                  
	SurveyToolTypeID                                                                           *string  `json:"SurveyToolTypeID,omitempty"`
	// Parent Wellbore Identifier.                                                                      
	WellboreID                                                                                 *string  `json:"WellboreID,omitempty"`
	// ID of the WellboreTrajectory that the Survey Program is associated to, e.g., for                 
	// compiling a Definitive Survey from Parts.                                                        
	WellboreTrajectoryID                                                                       *string  `json:"WellboreTrajectoryID,omitempty"`
}
