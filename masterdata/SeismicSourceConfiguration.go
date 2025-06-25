package masterdata

// Parameters characterizing the seismic source configuration.
type SeismicSourceConfiguration struct {
	// Seismic Source type. E.g.: Airgun, Vibroseis, Dynamite,Watergun.                                 
	EnergySourceTypeID                                                                         *string  `json:"EnergySourceTypeID,omitempty"`
	// Text remarks regarding the Seismic source configuration.                                         
	Remarks                                                                                    *string  `json:"Remarks,omitempty"`
	// Horizontal distance between shotpoint locations.                                                 
	ShotpointSpacing                                                                           *float64 `json:"ShotpointSpacing,omitempty"`
	// Number of energy sources.                                                                        
	SourceArrayCount                                                                           *int64   `json:"SourceArrayCount,omitempty"`
	// Depth of the energy source.                                                                      
	SourceArrayDepth                                                                           *float64 `json:"SourceArrayDepth,omitempty"`
	// Maximum depth of receivers in a wellbore. Used in conjunction with VSP acquisition.              
	SourceArrayMaxDepth                                                                        *float64 `json:"SourceArrayMaxDepth,omitempty"`
	// Minimum depth of Sources in a wellbore. Used in conjunction with VSP acquisition.                
	SourceArrayMinDepth                                                                        *float64 `json:"SourceArrayMinDepth,omitempty"`
	// Distance between energy sources.                                                                 
	SourceArraySpacing                                                                         *float64 `json:"SourceArraySpacing,omitempty"`
	// Maximum frequency of the vibroseis source.                                                       
	SourceArraySweepFreqMax                                                                    *float64 `json:"SourceArraySweepFreqMax,omitempty"`
	// Minimum frequency of the vibroseis source.                                                       
	SourceArraySweepFreqMin                                                                    *float64 `json:"SourceArraySweepFreqMin,omitempty"`
	// Length of the vibroseis source sweep.                                                            
	SourceArraySweepLength                                                                     *float64 `json:"SourceArraySweepLength,omitempty"`
	// Volume of the energy source.                                                                     
	SourceArrayVolume                                                                          *float64 `json:"SourceArrayVolume,omitempty"`
	// The relationship to the wellbore, in which the source or sources are located.                    
	SourceWellboreID                                                                           *string  `json:"SourceWellboreID,omitempty"`
	// Name of the source vessel (may be the same as the receiver).  In the case of a VSP, this         
	// may be a platform or rig.                                                                        
	VesselName                                                                                 *string  `json:"VesselName,omitempty"`
}
