package masterdata

// Quantity of items inventoried during drilling and/or operations
type Inventory struct {
	// Comments and remarks                                                                            
	Comments                                                                                  *string  `json:"Comments,omitempty"`
	// DEPRECATED: Cost for the product for the report interval.                                       
	CostItem                                                                                  *string  `json:"CostItem,omitempty"`
	// Cost for the product for the report interval.                                                   
	CostOfItem                                                                                *float64 `json:"CostOfItem,omitempty"`
	// Description of the inventory item                                                               
	ItemDescription                                                                           *string  `json:"ItemDescription,omitempty"`
	// DEPRECATED: This property is incorrectly defined as a number, please use the Name               
	// property instead. Name or type of inventory item.                                               
	ItemName                                                                                  *float64 `json:"ItemName,omitempty"`
	// Name or type of inventory item.                                                                 
	Name                                                                                      *string  `json:"Name,omitempty"`
	// Price per item unit, assume same currency for all items.                                        
	PricePerUnit                                                                              *float64 `json:"PricePerUnit,omitempty"`
	// Daily quantity adjustment/correction.                                                           
	TotalQuantityAdjustment                                                                   *float64 `json:"TotalQuantityAdjustment,omitempty"`
	// Amount of the item remaining on location after all adjustments for the report interval.         
	TotalQuantityOnLocation                                                                   *float64 `json:"TotalQuantityOnLocation,omitempty"`
	// Quantity received at the site.                                                                  
	TotalQuantityReceived                                                                     *float64 `json:"TotalQuantityReceived,omitempty"`
	// Quantity returned to base from site.                                                            
	TotalQuantityReturned                                                                     *float64 `json:"TotalQuantityReturned,omitempty"`
	// Start quantity for report interval.                                                             
	TotalQuantityStart                                                                        *float64 `json:"TotalQuantityStart,omitempty"`
	// Quantity used for the report interval.                                                          
	TotalQuantityUsed                                                                         *float64 `json:"TotalQuantityUsed,omitempty"`
	// Item volume per unit.                                                                           
	VolumeItem                                                                                *float64 `json:"VolumeItem,omitempty"`
	// Item weight per unit.                                                                           
	WeightItem                                                                                *float64 `json:"WeightItem,omitempty"`
}
