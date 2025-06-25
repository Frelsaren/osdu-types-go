package masterdata

// A group of parameters that refer to RPM (rotations per minute)
type RPMGroup struct {
	// The planned downhole rotary speed values for the BHA run.                  
	DownHoleRPMGroup                                            *DownHoleRPMGroup `json:"DownHoleRPMGroup,omitempty"`
	// The planned surface rotary speed values for the BHA run.                   
	SurfaceRPMGroup                                             *SurfaceRPMGroup  `json:"SurfaceRPMGroup,omitempty"`
}
