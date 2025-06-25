package masterdata

// Parameters characterizing the seismic receiver configuration.
type SeismicReceiverConfiguration struct {
	// Number of receiver arrays (lines).                                                                
	CableCount                                                                                  *int64   `json:"CableCount,omitempty"`
	// Marine seismic cable towing depth below sea surface  (Positive Down).                             
	CableDepth                                                                                  *float64 `json:"CableDepth,omitempty"`
	// Total length of receiver array.                                                                   
	CableLength                                                                                 *float64 `json:"CableLength,omitempty"`
	// Horizontal distance between receiver arrays.                                                      
	CableSpacing                                                                                *float64 `json:"CableSpacing,omitempty"`
	// Number of receivers on a cable.                                                                   
	ReceiverCount                                                                               *int64   `json:"ReceiverCount,omitempty"`
	// Distance between receiver groups on the same cable.                                               
	ReceiverGroupSpacing                                                                        *float64 `json:"ReceiverGroupSpacing,omitempty"`
	// Distance between receivers on same cable.                                                         
	ReceiverSpacingInterval                                                                     *float64 `json:"ReceiverSpacingInterval,omitempty"`
	// The type of receivers, e.g. geophones, hydrophones, ocean bottom seismometers.                    
	ReceiverTypeID                                                                              *string  `json:"ReceiverTypeID,omitempty"`
	// Text remarks regarding the Seismic Receiver configuration.                                        
	Remarks                                                                                     *string  `json:"Remarks,omitempty"`
	// Name of the receiver vessel (may be the same as the source).  In the case of a VSP, this          
	// may be a platform or rig.                                                                         
	VesselName                                                                                  *string  `json:"VesselName,omitempty"`
	// The relationship to the wellbore, in which the receivers are located. Used in conjunction         
	// with VSP acquisition.                                                                             
	WellboreID                                                                                  *string  `json:"WellboreID,omitempty"`
	// Maximum depth of receivers in a wellbore. Used in conjunction with VSP acquisition.               
	WellboreReceiverMaxDepth                                                                    *float64 `json:"WellboreReceiverMaxDepth,omitempty"`
	// Minimum depth of receivers in a wellbore. Used in conjunction with VSP acquisition.               
	WellboreReceiverMinDepth                                                                    *float64 `json:"WellboreReceiverMinDepth,omitempty"`
}
