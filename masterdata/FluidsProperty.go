package masterdata

// A description of fluid used in the drilling of a wellbore
type FluidsProperty struct {
	// Array of Fluids facet that describes each individual programmed property of the fluids.             
	FluidFacets                                                                               []FluidFacet `json:"FluidFacets"`
	// An open reference list of fluid properties                                                          
	FluidPropertyNameID                                                                       string       `json:"FluidPropertyNameID"`
}
