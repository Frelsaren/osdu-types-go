package masterdata

// A log pass identifies a single acquisition event or job.  A log run my contain one or
// multiple log passes.
type LogPasses struct {
	// The flowing condition of the wellbore during the log acquisition e.g. Static or Dynamic                  
	FlowingConditionID                                                                         *string          `json:"FlowingConditionID,omitempty"`
	// The type of hole logged in this pass e.g. open hole, cased hole, cemented hole.                          
	HoleTypeIDs                                                                                []string         `json:"HoleTypeIDs,omitempty"`
	// The index reference type for this log pass e.g. depth, time, counter, calendar time.                     
	// IndexTypeID corresponds to WellLog data.SamplingDomainTypeID.                                            
	IndexTypeID                                                                                *string          `json:"IndexTypeID,omitempty"`
	// A boolean property indicating the sampling mode of the index or reference curve for this                 
	// Log Pass.                                                                                                
	//                                                                                                          
	// True means all reference curve values are regularly spaced (see SamplingInterval). False                 
	// means irregular or discrete sample spacing.                                                              
	IsRegularSamplingInterval                                                                  *bool            `json:"IsRegularSamplingInterval,omitempty"`
	// The direction of logging for this log pass.                                                              
	// Specifies whether the logging information was collected in an downward or upward                         
	// direction or stationary.                                                                                 
	LoggingDirectionID                                                                         *string          `json:"LoggingDirectionID,omitempty"`
	LogPassID                                                                                  *string          `json:"LogPassID,omitempty"`
	// The type of log pass such as Calibration Pass, Main Pass, Repeated Pass.                                 
	PassTypeID                                                                                 *string          `json:"PassTypeID,omitempty"`
	// Remarks specific to this log pass.                                                                       
	Remarks                                                                                    []AbstractRemark `json:"Remarks,omitempty"`
	// A reference to the Sample Acquisition Job for samples acquired in the same logging pass.                 
	// This can be one or more Sample Acquisition Jobs.                                                         
	SampleAcquisitionJobIDs                                                                    []string         `json:"SampleAcquisitionJobIDs,omitempty"`
	// For regularly sampled curves this property holds the sampling interval for this Log                      
	// Pass.                                                                                                    
	// For a non regular sampling interval this property is not set. The                                        
	// IsRegularSamplingInterval flag indicates whether the SamplingInterval is required.                       
	SamplingInterval                                                                           *float64         `json:"SamplingInterval,omitempty"`
	// The smallest value/first value of the log index for this log pass.                                       
	SamplingStart                                                                              *float64         `json:"SamplingStart,omitempty"`
	// The largest value/last value of the log index for this log pass.                                         
	SamplingStop                                                                               *float64         `json:"SamplingStop,omitempty"`
	// The type of fluid in the wellbore at time of acquisition of the log pass e.g. oil based                  
	// mud, water based mud, water.                                                                             
	// Only populate if it is different to the wellbore fluid identified in the log run.                        
	WellboreFluidTypeID                                                                        *string          `json:"WellboreFluidTypeID,omitempty"`
}
