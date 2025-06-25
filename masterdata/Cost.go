package masterdata

// Cost information captured for a defined time period during operations and/or drilling
type Cost struct {
	// Activity code classification describing higher level steps in the Well Activity                 
	ActivityPhaseID                                                                           *string  `json:"ActivityPhaseID,omitempty"`
	// AFE number that this cost item applies to.                                                      
	AFENumber                                                                                 *string  `json:"AFENumber,omitempty"`
	// Comments and remarks                                                                            
	Comments                                                                                  *string  `json:"Comments,omitempty"`
	// Cost class code.                                                                                
	CostClass                                                                                 *string  `json:"CostClass,omitempty"`
	// Cost code.                                                                                      
	CostCode                                                                                  *string  `json:"CostCode,omitempty"`
	// Currency used for Cost Amount                                                                   
	CostCurrency                                                                              string   `json:"CostCurrency"`
	// Cost group code.                                                                                
	CostGroup                                                                                 *string  `json:"CostGroup,omitempty"`
	// Name of pool/reservoir that this cost item can be accounted to.                                 
	CostPoolName                                                                              *string  `json:"CostPoolName,omitempty"`
	// Cost subcode.                                                                                   
	CostSubCode                                                                               *string  `json:"CostSubCode,omitempty"`
	// Used to indicate if Equipment or Service in scope for the Cost Item is being used. Can          
	// drive different Standby/In Use Rental Rates where applicable.                                   
	InUse                                                                                     *bool    `json:"InUse,omitempty"`
	// Invoice number for cost item; the bill is sent to the operator.                                 
	InvoiceNumber                                                                             *string  `json:"InvoiceNumber,omitempty"`
	// Is this item carried from day to day?                                                           
	IsCarryOver                                                                               *bool    `json:"IsCarryOver,omitempty"`
	// Is this an estimated cost?                                                                      
	IsEstimated                                                                               *bool    `json:"IsEstimated,omitempty"`
	// Is this item a rental?                                                                          
	IsRental                                                                                  *bool    `json:"IsRental,omitempty"`
	// Description of the cost item.                                                                   
	ItemDescription                                                                           *string  `json:"ItemDescription,omitempty"`
	// Serial number.                                                                                  
	ItemSerialNumber                                                                          *string  `json:"ItemSerialNumber,omitempty"`
	// Purchase order number provided by the operator.                                                 
	PONumber                                                                                  *string  `json:"PONumber,omitempty"`
	// High Level cost category qualifier                                                              
	ProductType                                                                               *string  `json:"ProductType,omitempty"`
	// An identification tag for the item. A serial number is a type of identification tag;            
	// however, some tags contain many pieces of information. This element only identifies the         
	// tag and does not describe the contents.                                                         
	TagName                                                                                   *string  `json:"TagName,omitempty"`
	// The field ticket number issued by the service company on location.                              
	TicketNumber                                                                              *string  `json:"TicketNumber,omitempty"`
	// Total cost for the item for this period                                                         
	TotalCostAmount                                                                           float64  `json:"TotalCostAmount"`
	// Number of cost items used that day, e.g., 1 rig dayrate, 30 joints of casing.                   
	TotalQuantityItemsUsed                                                                    *int64   `json:"TotalQuantityItemsUsed,omitempty"`
	// Cost per item, assume same currency.                                                            
	UnitCost                                                                                  *float64 `json:"UnitCost,omitempty"`
	// The kind of cost item specified (e.g., barrel, sack, gallon)                                    
	UnitKind                                                                                  *string  `json:"UnitKind,omitempty"`
	// Quantity of the unit kind (e.g., 50)                                                            
	UnitSize                                                                                  *float64 `json:"UnitSize,omitempty"`
	// Name of the vendor.                                                                             
	VendorName                                                                                *string  `json:"VendorName,omitempty"`
	// Vendor number.                                                                                  
	VendorNumber                                                                              *string  `json:"VendorNumber,omitempty"`
}
