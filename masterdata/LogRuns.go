package masterdata

// A log run identifies a single series of combinable logging tools that are entered into
// the wellbore below the reference depth (e.g. KB, the).
type LogRuns struct {
	// The relationship to a service or contractor organisation, typically the producer or                       
	// logging contractor for this particular and identifiable Log Run. Typically this is the                    
	// same as the overall well log acquisition.                                                                 
	// Only populate this property if a particular and identifiable Log Run has a different or                   
	// specifically identifiable service or contracting company to the Well Log Acquisition                      
	// project as a whole.                                                                                       
	// The overall contracting company is captured under the `Project.Contractors[]` array. This                 
	// `Project.Contractors[]` array can capture all Contractors associated with the overall                     
	// Well Log Acquisition project activity, but can not be associated to a specific Log Run.                   
	ContractorCompanyID                                                                         *string          `json:"ContractorCompanyID,omitempty"`
	// The conveyance method used to acquire the log data e.g. wireline, LWD, pipe conveyed.                     
	ConveyanceMethodID                                                                          *string          `json:"ConveyanceMethodID,omitempty"`
	// An array of well log generic tool types used in this Log Run.                                             
	GenericToolTypeIDs                                                                          []string         `json:"GenericToolTypeIDs,omitempty"`
	// The logging passes within this single logging run.                                                        
	LogPasses                                                                                   []LogPasses      `json:"LogPasses,omitempty"`
	// A unique identifier for a specific log run. This can be a number or alphanumeric.                         
	LogRunID                                                                                    *string          `json:"LogRunID,omitempty"`
	// Remarks specific to this log run.                                                                         
	Remarks                                                                                     []AbstractRemark `json:"Remarks,omitempty"`
	// The smallest value/first value of the log index for this log run.                                         
	SamplingStart                                                                               *float64         `json:"SamplingStart,omitempty"`
	// The largest value/last value of the log index for this log run.                                           
	SamplingStop                                                                                *float64         `json:"SamplingStop,omitempty"`
	// A relationship to the document record within the OSDU Platform containing the tool sensor                 
	// offsets, for example a tool string diagram or report.                                                     
	SensorOffsetDocumentID                                                                      *string          `json:"SensorOffsetDocumentID,omitempty"`
	// An array of PWLS tool mnemonics used in this Log Run.                                                     
	ToolMnemonicIDs                                                                             []string         `json:"ToolMnemonicIDs,omitempty"`
	// The type of fluid in the wellbore at time of logging                                                      
	// e.g. oil based mud, water based mud, water. This property could be overridden by the                      
	// LogPass wellbore fluid type.                                                                              
	WellboreFluidTypeID                                                                         *string          `json:"WellboreFluidTypeID,omitempty"`
}
