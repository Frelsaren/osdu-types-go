package masterdata

// A core taken during sidewall coring.
type SidewallCoreSample struct {
	// The measured depth at position the sidewall core was taken, typically in logger's depth.         
	Depth                                                                                      *float64 `json:"Depth,omitempty"`
	// The sidewall core sample description, if available.                                              
	Description                                                                                *string  `json:"Description,omitempty"`
	// The recovered length of the sidewall core.                                                       
	RecoveredLength                                                                            *float64 `json:"RecoveredLength,omitempty"`
	// Typically a sequential number identifying the sidewall core sample.                              
	SampleID                                                                                   *string  `json:"SampleID,omitempty"`
}
