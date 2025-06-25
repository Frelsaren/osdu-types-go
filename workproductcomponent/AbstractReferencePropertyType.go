package workproductcomponent

// The purpose of this schema is best understood in the context of a columnar dataset: the
// AbstractReferencePropertyType describes a column in a columnar dataset by declaring its
// value type (number, string), a UnitQuantity if the value type is a number, a kind if the
// string value is actually a relationship to a e.g. reference-data type.
type AbstractReferencePropertyType struct {
	// When describing a table column, e.g., in an associated dataset, this optional property                         
	// allows the association of the record column definition(s) to dataset table column(s).                          
	ColumnName                                                                                  *string               `json:"ColumnName,omitempty"`
	// Ordered array with: FacetType, FacetRole, both calling specific references                                     
	//                                                                                                                
	// FacetType: Enumerations of the type of additional context about the nature of a property                       
	// type (it may include conditions, direction, qualifiers, or statistics).                                        
	//                                                                                                                
	// FacetRole: Additional context about the nature of a property type. The purpose of such                         
	// attribute is to minimize the need to create specialized property types by mutualizing                          
	// some well known qualifiers such as "maximum", "minimum" which apply to a lot of different                      
	// property types.                                                                                                
	FacetIDs                                                                                    []AbstractFacet       `json:"FacetIDs,omitempty"`
	// It holds the PropertyType associated with this reference property type, further defining                       
	// the semantics of the value. It contains a relationship to PropertyType record and its                          
	// (de-normalized) name. String or number values can represent e.g. A date or a time by                           
	// referring to the respective PropertyType record id.                                                            
	PropertyType                                                                                *AbstractPropertyType `json:"PropertyType,omitempty"`
	// Only populated if ValueType=="string" and the values are expected to represent record                          
	// ids, e.g. to a reference-data type, then this value holds the kind (optionally without                         
	// the semantic version number).                                                                                  
	RelationshipTargetKind                                                                      *string               `json:"RelationshipTargetKind,omitempty"`
	// Only populated of the ValueType is "number". It holds the UnitOfMeasure associated with                        
	// this reference property type. It is a relationship to a UnitOfMeasure record. If the                           
	// UnitQuantityID and/or PropertyType.PropertyTypeID are populated in addition to                                 
	// UnitOfMeasureID, the referenced records must finally share the same dimension code. See                        
	// Schema Usage Guide 'Unit of Measure Foundation'.                                                               
	UnitOfMeasureID                                                                             *string               `json:"UnitOfMeasureID,omitempty"`
	// Only populated of the ValueType is "number". It holds the UnitQuantity associated with                         
	// this reference property type. It is a relationship to UnitQuantity record.                                     
	UnitQuantityID                                                                              *string               `json:"UnitQuantityID,omitempty"`
	// The number of values in a tuple, e.g. For coordinates. The default is 1.                                       
	ValueCount                                                                                  *int64                `json:"ValueCount,omitempty"`
	// The type of value to expect for this reference property, either "number" (floating point                       
	// number), "integer",  "string", or "boolean".                                                                   
	ValueType                                                                                   *string               `json:"ValueType,omitempty"`
}
