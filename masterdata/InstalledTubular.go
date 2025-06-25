package masterdata

// Geometrical Description of tubular assembly as installed in the wellbore
type InstalledTubular struct {
	// The measured depth at the base of the tubular                                    
	MeasuredDepthBase                                                          *float64 `json:"MeasuredDepthBase,omitempty"`
	// The  error associated with the measured depth at the base of the tubular         
	MeasuredDepthErrorBase                                                     *float64 `json:"MeasuredDepthErrorBase,omitempty"`
	// The error associated with the measured depth at the top of the tubular           
	MeasuredDepthErrorTop                                                      *float64 `json:"MeasuredDepthErrorTop,omitempty"`
	// The measured depth at the top of the tubular                                     
	MeasuredDepthTop                                                           *float64 `json:"MeasuredDepthTop,omitempty"`
	// Identifier of the tubular assembly actually installed or to be installed         
	TubularAssemblyID                                                          string   `json:"TubularAssemblyID"`
}
