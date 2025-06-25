package masterdata

import "time"

// Mud that has been lost during drilling / operations
type MudLosses struct {
	// Comments and remarks                                                                      
	Comments                                                                           *string   `json:"Comments,omitempty"`
	// Date and time that the mud loss occurred                                                  
	DateTime                                                                           time.Time `json:"DateTime"`
	// Mud volume lost downhole during abandonment.                                              
	MudLostVolumeAbandonDownhole                                                       *float64  `json:"MudLostVolumeAbandonDownhole,omitempty"`
	// Mud volume lost downhole behind casing.                                                   
	MudLostVolumeBehindCasingDownhole                                                  *float64  `json:"MudLostVolumeBehindCasingDownhole,omitempty"`
	// Mud volume lost downhole running casing                                                   
	MudLostVolumeCasingHole                                                            *float64  `json:"MudLostVolumeCasingHole,omitempty"`
	// Mud volume lost downhole while cementing.                                                 
	MudLostVolumeCementingDownhole                                                     *float64  `json:"MudLostVolumeCementingDownhole,omitempty"`
	// Mud volume lost downhole while circulating.                                               
	MudLostVolumeCirculatingDownhole                                                   *float64  `json:"MudLostVolumeCirculatingDownhole,omitempty"`
	// Bottom measured depth for which the mud loss reading was occurred                         
	MudLostVolumeMeasureDepthBase                                                      *float64  `json:"MudLostVolumeMeasureDepthBase,omitempty"`
	// Top measured depth for which the mud loss reading was occurred                            
	MudLostVolumeMeasureDepthTop                                                       *float64  `json:"MudLostVolumeMeasureDepthTop,omitempty"`
	// Volume of mud lost in mud cleaning equipment (at surface).                                
	MudLostVolumeMudCleanerSurface                                                     *float64  `json:"MudLostVolumeMudCleanerSurface,omitempty"`
	// Mud volume lost downhole from other location.                                             
	MudLostVolumeOtherLocationDownhole                                                 *float64  `json:"MudLostVolumeOtherLocationDownhole,omitempty"`
	// Surface volume lost other location.                                                       
	MudLostVolumeOtherLocationSurface                                                  *float64  `json:"MudLostVolumeOtherLocationSurface,omitempty"`
	// Volume of mud lost in pit room (at surface).                                              
	MudLostVolumePitsSurface                                                           *float64  `json:"MudLostVolumePitsSurface,omitempty"`
	// Volume of mud lost at shakers (at surface).                                               
	MudLostVolumeShakerSurface                                                         *float64  `json:"MudLostVolumeShakerSurface,omitempty"`
	// Volume of mud lost while tripping (at surface).                                           
	MudLostVolumeTrippingSurface                                                       *float64  `json:"MudLostVolumeTrippingSurface,omitempty"`
	// Bottom true vertical depth interval over which the mud loss reading was occurred          
	MudLostVolumeTvdBase                                                               *float64  `json:"MudLostVolumeTvdBase,omitempty"`
	// Top  true vertical depth interval over which the mud loss reading was occurred            
	MudLostVolumeTvdTop                                                                *float64  `json:"MudLostVolumeTvdTop,omitempty"`
	// Total volume of mud lost downhole.                                                        
	TotalMudLostVolumeDownhole                                                         *float64  `json:"TotalMudLostVolumeDownhole,omitempty"`
	// Total volume of mud lost at surface.                                                      
	TotalMudLostVolumeSurface                                                          *float64  `json:"TotalMudLostVolumeSurface,omitempty"`
}
