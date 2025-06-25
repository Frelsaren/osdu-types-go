package masterdata

// The property type ID and Name, which determines eventually the UnitQuantity of the
// parameter value. Used to provide a more scoped context than UnitQuantityID. If
// PropertyType is provided, UnitQuantityID is expected to be omitted.
//
// A nested object holding the relationship to a PropertyType by id (uuid) and a derived,
// human-readable name.
//
// It holds the PropertyType associated with this reference property type, further defining
// the semantics of the value. It contains a relationship to PropertyType record and its
// (de-normalized) name. String or number values can represent e.g. A date or a time by
// referring to the respective PropertyType record id.
type AbstractPropertyType struct {
	// The name of the PropertyType, de-normalized, derived from the record referenced in               
	// PropertyTypeID.                                                                                  
	Name                                                                                        *string `json:"Name,omitempty"`
	// The relationship to the PropertyType reference data item, typically containing an                
	// Energistics PWLS 3 uuid. For better traceability and usability the property name is to be        
	// populated in the Name property.                                                                  
	PropertyTypeID                                                                              *string `json:"PropertyTypeID,omitempty"`
}
