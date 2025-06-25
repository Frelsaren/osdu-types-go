package masterdata

// The Conventional Coring specific sub-structure.
type ConventionalCoring struct {
	// An array of core sections.                                                                     
	CoreSections                                                                        []CoreSection `json:"CoreSections,omitempty"`
	// The name identifying the type of coring procedure used to extract the core, e.g.,              
	// Conventional Core,  Drop Cores.                                                                
	CoreTypeID                                                                          *string       `json:"CoreTypeID,omitempty"`
	// The length of core recovered in this run.                                                      
	RecoveredLength                                                                     *float64      `json:"RecoveredLength,omitempty"`
}
