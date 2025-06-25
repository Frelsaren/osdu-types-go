package workproductcomponent

// The Well Log Acquisition details object captures Information relevant to the well log
// acquisition, such as the specific acquisition job, log runs and log passes that this
// well log information derives from.
type WellLogAcquisitionDetails struct {
	// A relationship to particular LogPassIDs specified within the Well Log Acquisition. There         
	// can be one or many LogPasses.                                                                    
	// Human readable reference only e.g. no reference value or pattern is available. Data              
	// quality rules can be established to assist in managing this relationship.                        
	LogPassIDs                                                                                 []string `json:"LogPassIDs,omitempty"`
	// A relationship to particular LogRunIDs specified within the Well Log Acquisition. There          
	// can be one or many LogRuns.                                                                      
	// Human readable reference only e.g. no reference value or pattern is available. Data              
	// quality rules can be established to assist in managing this relationship.                        
	LogRunIDs                                                                                  []string `json:"LogRunIDs,omitempty"`
	// A relationship to the Well Log Acquisition record relevant to this well log and set of           
	// log curves.                                                                                      
	WellLogAcquisitionID                                                                       *string  `json:"WellLogAcquisitionID,omitempty"`
}
