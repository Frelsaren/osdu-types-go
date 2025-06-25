package workproductcomponent

import "time"

// Log Evaluation parameters
type LogEvaluation struct {
	// Elapsed Time Before Log from end of Cement Job                                 
	BeforeLogElapsedTime                                         *float64             `json:"BeforeLogElapsedTime,omitempty"`
	// Cement bond log quality description.                                           
	CBLLogQuality                                                *string              `json:"CBLLogQuality,omitempty"`
	// Cement Bond Log Pressure                                                       
	CBLPressure                                                  *float64             `json:"CBLPressure,omitempty"`
	// Evaluation Date/time                                                           
	EvaluationDate                                               *time.Time           `json:"EvaluationDate,omitempty"`
	// Bottom Hole Temperature Log Run?                                               
	IsBHTRun                                                     *bool                `json:"IsBHTRun,omitempty"`
	// Cement Bond Log tool run?                                                      
	IsCBLRun                                                     *bool                `json:"IsCBLRun,omitempty"`
	// Cement Found Between Shoe Collar during Evaluation Log run                     
	IsCementFoundBetweenShoeCollar                               *bool                `json:"IsCementFoundBetweenShoeCollar,omitempty"`
	// Cement Found On Tool during Evaluation Log run                                 
	IsCementFoundOnTool                                          *bool                `json:"IsCementFoundOnTool,omitempty"`
	// Cement Evaluation Tool Run?                                                    
	IsCETRun                                                     *bool                `json:"IsCETRun,omitempty"`
	// Log Evaluation Depths                                                          
	LogEvaluationDepths                                          []LogEvaluationDepth `json:"LogEvaluationDepths,omitempty"`
	// Log Evaluation Results                                                         
	LogEvaluationResults                                         *string              `json:"LogEvaluationResults,omitempty"`
	// ID to evaluation Log                                                           
	LogID                                                        *string              `json:"LogID,omitempty"`
	// Log type                                                                       
	LogType                                                      *string              `json:"LogType,omitempty"`
}
