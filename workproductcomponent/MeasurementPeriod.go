package workproductcomponent

// Contains information about the individual, discrete events or periods during a wellbore
// pressure test
type MeasurementPeriod struct {
	// An identifying kind of the MeasurementPeriod e.g. BuildUp, FallOff, PreTest                                  
	MeasurementPeriodKind                                                                 *string                   `json:"MeasurementPeriodKind,omitempty"`
	// An embedded ColumnBasedTable containing the Measurements within the period that are                          
	// required (and only if required) for indexing                                                                 
	MeasurementPeriodMeasurements                                                         *AbstractColumnBasedTable `json:"MeasurementPeriodMeasurements,omitempty"`
	// Time index where the individual period ends into the sequence                                                
	PeriodEndTime                                                                         *float64                  `json:"PeriodEndTime,omitempty"`
	// Sequential number identifying the period within the record                                                   
	PeriodNumber                                                                          *int64                    `json:"PeriodNumber,omitempty"`
	// Time index where the individual period starts into the sequence                                              
	PeriodStartTime                                                                       *float64                  `json:"PeriodStartTime,omitempty"`
}
