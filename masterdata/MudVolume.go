package masterdata

import "time"

// Information related to mud volumes for drilling/operations report
type MudVolume struct {
	// Comments and remarks                                                                          
	Comments                                                                               *string   `json:"Comments,omitempty"`
	// Date and time of the mud volume                                                               
	DateTime                                                                               time.Time `json:"DateTime"`
	// Total mud volume that was expected                                                            
	ExpectedMudVolume                                                                      *float64  `json:"ExpectedMudVolume,omitempty"`
	// Volume of mud built.                                                                          
	MudVolumeBuild                                                                         *float64  `json:"MudVolumeBuild,omitempty"`
	// DEPRECATED - CHANGED OBJECT TYPE.  Volume of mud built.                                       
	MudVolumeBuilt                                                                         *string   `json:"MudVolumeBuilt,omitempty"`
	// Volume of mud contained in casing annulus.                                                    
	MudVolumeCasing                                                                        *float64  `json:"MudVolumeCasing,omitempty"`
	// Volume of mud dumped.                                                                         
	MudVolumeDumped                                                                        *float64  `json:"MudVolumeDumped,omitempty"`
	// Volume of mud at the end of the report interval (including pits and hole).                    
	MudVolumeEnd                                                                           *float64  `json:"MudVolumeEnd,omitempty"`
	// Volume of mud contained in the openhole annulus.                                              
	MudVolumeHole                                                                          *float64  `json:"MudVolumeHole,omitempty"`
	// Bottom measured depth for which the mud volume reading was conducted                          
	MudVolumeMeasureDepthBase                                                              *float64  `json:"MudVolumeMeasureDepthBase,omitempty"`
	// Top measured depth for which the mud volume reading was conducted                             
	MudVolumeMeasureDepthTop                                                               *float64  `json:"MudVolumeMeasureDepthTop,omitempty"`
	// Volume of mud received from mud warehouse.                                                    
	MudVolumeReceived                                                                      *float64  `json:"MudVolumeReceived,omitempty"`
	// Volume of mud returned to mud warehouse.                                                      
	MudVolumeReturned                                                                      *float64  `json:"MudVolumeReturned,omitempty"`
	// Volume of mud contained in riser section annulus.                                             
	MudVolumeRiser                                                                         *float64  `json:"MudVolumeRiser,omitempty"`
	// Volume of mud at start of report interval (including pits and hole).                          
	MudVolumeStart                                                                         *float64  `json:"MudVolumeStart,omitempty"`
	// Volume of mud contained within active string.                                                 
	MudVolumeString                                                                        *float64  `json:"MudVolumeString,omitempty"`
	// Bottom true vertical depth interval over which the mud volume reading was conducted.          
	MudVolumeTvdBase                                                                       *float64  `json:"MudVolumeTvdBase,omitempty"`
	// Top  true vertical depth interval over which the mud volume reading was conducted.            
	MudVolumeTvdTop                                                                        *float64  `json:"MudVolumeTvdTop,omitempty"`
	// Actual total mud volume that was measured                                                     
	TotalMudVolume                                                                         *float64  `json:"TotalMudVolume,omitempty"`
}
