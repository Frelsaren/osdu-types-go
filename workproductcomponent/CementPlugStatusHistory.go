package workproductcomponent

import "time"

// Cement Plug status history
type CementPlugStatusHistory struct {
	// Plug Status Base MD                 
	PlugStatusBaseMeasuredDepth *float64   `json:"PlugStatusBaseMeasuredDepth,omitempty"`
	// Status date/time                    
	PlugStatusDatetime          *time.Time `json:"PlugStatusDatetime,omitempty"`
	// Status Remarks                      
	PlugStatusRemarks           *string    `json:"PlugStatusRemarks,omitempty"`
	// Plug Status Top MD                  
	PlugStatusTopMeasuredDepth  *float64   `json:"PlugStatusTopMeasuredDepth,omitempty"`
	// Status Type                         
	PlugStatusTypeID            *string    `json:"PlugStatusTypeID,omitempty"`
}
