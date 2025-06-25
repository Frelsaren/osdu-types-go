package workproductcomponent

import "time"

// Shoetrack Cement
type ShoetrackCement struct {
	// Is Hard Cement?                               
	IsHardCement                          *bool      `json:"IsHardCement,omitempty"`
	// Comments or Remarks                           
	ShoetrackCementTestRemarks            *string    `json:"ShoetrackCementTestRemarks,omitempty"`
	// Shoetrack Drill Date                          
	ShoetrackDrillDate                    *time.Time `json:"ShoetrackDrillDate,omitempty"`
	// Shoetrack Drill Measured Depth                
	ShoetrackDrillMeasuredDepth           *float64   `json:"ShoetrackDrillMeasuredDepth,omitempty"`
	// Shoetrack Drill True Vertical Depth           
	ShoetrackDrillTrueVerticalDepth       *float64   `json:"ShoetrackDrillTrueVerticalDepth,omitempty"`
	// Shoetrack Length                              
	ShoetrackLength                       *float64   `json:"ShoetrackLength,omitempty"`
	// Shoetrack Measured Depth                      
	ShoetrackMeasuredDepth                *float64   `json:"ShoetrackMeasuredDepth,omitempty"`
	// Shoetrack True Vertical Depth                 
	ShoetrackTrueVerticalDepth            *float64   `json:"ShoetrackTrueVerticalDepth,omitempty"`
}
