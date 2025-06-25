package workproductcomponent

// A subdivision of a reservoir compartment generally corresponding to one geologic unit.
type ReservoirCompartmentUnit struct {
	// The stratigraphic or the geobody associated to this Reservoir compartment unit.                
	GeologicUnitID                                                                           *string  `json:"GeologicUnitID,omitempty"`
	// Identifier of the reservoir compartment unit. It should be an UUID ideally. Used to be         
	// able to easier reference a reservoir compartment from other WPC.                               
	Identifier                                                                               *string  `json:"Identifier,omitempty"`
	// The rock fluid unit which are associated to this Reservoir compartment unit.                   
	RockFluidUnitIDs                                                                         []string `json:"RockFluidUnitIDs,omitempty"`
}
