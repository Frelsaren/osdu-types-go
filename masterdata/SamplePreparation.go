package masterdata

import "time"

// This captures information about the preparation process executed after the sample
// acquisition event.
type SamplePreparation struct {
	// This captures other pertinent information regarding the sample preparation process.                          
	Remarks                                                                                    []AbstractRemark     `json:"Remarks,omitempty"`
	// Used to describe the pressure and temperature conditions at which the sample preparation                     
	// took place                                                                                                   
	SamplePreparationCondition                                                                 *AbstractPTCondition `json:"SamplePreparationCondition,omitempty"`
	// This represents the end date for the sample preparation process.                                             
	SamplePreparationEndDate                                                                   *time.Time           `json:"SamplePreparationEndDate,omitempty"`
	// Provide additional details on which industrial/lab method used to conduct the sample                         
	// preparation                                                                                                  
	SamplePreparationMethodID                                                                  *string              `json:"SamplePreparationMethodID,omitempty"`
	// This represents the start date for the sample preparation process.                                           
	SamplePreparationStartDate                                                                 *time.Time           `json:"SamplePreparationStartDate,omitempty"`
	// Provides extra details on any processes applied after the sample has been acquired                           
	SamplePreparationTypeID                                                                    *string              `json:"SamplePreparationTypeID,omitempty"`
}
