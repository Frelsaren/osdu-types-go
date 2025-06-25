package referencedata

// A configuration for a UnitQuantity offering a sub-set of units and a default unit.
type FluffyConfiguration struct {
	// The default UnitOfMeasure to be used for this measurement name. If empty, the first                            
	// element of the mandatory PreferredUnitIDs array should be taken. The DefaultUnitID should                      
	// be member of the PreferredUnitIDs array.                                                                       
	DefaultUnitID                                                                               *string               `json:"DefaultUnitID,omitempty"`
	// The name of the measurement.                                                                                   
	Name                                                                                        *string               `json:"Name,omitempty"`
	// A hint how the number is expected to be presented, e.g., d or D for decimal, f or F for                        
	// fixed point, e or E for exponential (scientific), or g or G for general (default). Not                         
	// all languages support all codes in all cases - in principle the case means that the                            
	// resulting case is transformed to upper case or lower case depending on the case of the                         
	// NumericFormatType.                                                                                             
	NumericFormatType                                                                           *NumericFormatType    `json:"NumericFormatType,omitempty"`
	// The number of decimal digits for NumericFormatType f or F or e or E, or the number of                          
	// significant digits in g or G. If populated in conjunction with NumericFormatType d or D,                       
	// NumericPrecision defines the minimum number of digits. If the number has less digits than                      
	// given by NumericPrecision, it is padded with leading zeroes.                                                   
	NumericPrecision                                                                            *int64                `json:"NumericPrecision,omitempty"`
	// The preferred sub-set of units, which are meaningful for the domain, app or user. This                         
	// property is mandatory and must contain at least one element.                                                   
	PreferredUnitIDs                                                                            []string              `json:"PreferredUnitIDs"`
	// The list of property names, to which this configuration should apply. At least one of the                      
	// fields UnitQuantityID, PropertyType and or PropertyNames must be populated. Scope narrows                      
	// from UnitQuantityID, PropertyType to PropertyNames.                                                            
	PropertyNames                                                                               []string              `json:"PropertyNames,omitempty"`
	// If specified,Energistics PWLS 3 PropertyType implies a UnitQuantityID and offers a much                        
	// finer scoping. PropertyType refers to a pair PropertyTypeID (typically a UUID) and a                           
	// PropertyName. At least one of the fields UnitQuantityID, PropertyType and or                                   
	// PropertyNames must be populated. Scope narrows from UnitQuantityID, PropertyType to                            
	// PropertyNames.                                                                                                 
	PropertyType                                                                                *AbstractPropertyType `json:"PropertyType,omitempty"`
	// The relationship to the UnitQuantity for this configuration. At least one of the fields                        
	// UnitQuantityID, PropertyType and or PropertyNames must be populated. Scope narrows from                        
	// UnitQuantityID, PropertyType to PropertyNames.                                                                 
	UnitQuantityID                                                                              *string               `json:"UnitQuantityID,omitempty"`
}
