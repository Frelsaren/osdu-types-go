package masterdata

// A description of an individual facet composing a fluid
type FluidFacet struct {
	// Identifier to the list of potential types of Fluid Property Facet (e.g. Measured, Range        
	// (Min), Range (Max)                                                                             
	FluidPropertyFacetNameID                                                                  string  `json:"FluidPropertyFacetNameID"`
	// The measurement unit of the Fluid Property Value                                               
	FluidPropertyUnit                                                                         string  `json:"FluidPropertyUnit"`
	// The actual value of the property                                                               
	FluidPropertyValue                                                                        float64 `json:"FluidPropertyValue"`
	// The unique identifier of the fluid property (density or temperature measurement,….)            
	MeasuredPropertyID                                                                        string  `json:"MeasuredPropertyID"`
}
