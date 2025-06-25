package masterdata

// Description of the formulation of the barrel that will be part of the drilling mud
type BarrelFormulation struct {
	// The unit of measure of the planned product concentration                                          
	ConcentrationUnitOfMeasure                                                                  *string  `json:"ConcentrationUnitOfMeasure,omitempty"`
	// The actual value of the planned product concentration                                             
	ConcentrationValue                                                                          *float64 `json:"ConcentrationValue,omitempty"`
	// The code associated with the product                                                              
	ProductCode                                                                                 *string  `json:"ProductCode,omitempty"`
	// The function that the component plays in the formulation. For example Base Oil, Fresh             
	// Water, Chemical Additive                                                                          
	ProductFunction                                                                             *string  `json:"ProductFunction,omitempty"`
	// Name of the Product                                                                               
	ProductName                                                                                 string   `json:"ProductName"`
	// The packaging that the component comes in. For example Sacks, Bags. For Liquids this is           
	// usually N/A                                                                                       
	ProductPackage                                                                              *string  `json:"ProductPackage,omitempty"`
	// The specific gravity value of the component.                                                      
	ProductSg                                                                                   *float64 `json:"ProductSg,omitempty"`
	// The unit size of the component. For example if this product came in 25 lb bags then the           
	// Unit would be lbs, the UnitSize would be 25 and the Product Package would be "Bag"                
	ProductUnitOfMeasure                                                                        *string  `json:"ProductUnitOfMeasure,omitempty"`
	// The unit size of the component. For example if this product came in 25 lb bags then the           
	// Unit would be lbs, the UnitSize would be 25 and the Product Package would be "Bag"                
	ProductUnitSize                                                                             *string  `json:"ProductUnitSize,omitempty"`
	// The actual number of units of the component. The unit size of the component. For example          
	// if you wanted to use 15 x 25 lb bags then the Unit would be lbs, the UnitSize would be 25         
	// and the Product Package would be "Bag" and the Quantity would be 15                               
	Quantity                                                                                    *float64 `json:"Quantity,omitempty"`
}
