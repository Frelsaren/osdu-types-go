package workproductcomponent

import "time"

// One LabelType, e.g., GeoLabelType, associated with a value. The value type (number,
// integer, boolean, string, relationship) is implied by the label type (related via
// LabelTypeID). If the label type's ValueCount is greater than 1, then arrays of number,
// integer, boolean, string or relationships are expected to be populated.
type AbstractLabelAndValue struct {
	// The reference to a label type record, e.g., a GeoLabelType, describing the details and            
	// facets. The content of the referenced label type determines, which properties are                 
	// expected to be populated.                                                                         
	LabelTypeID                                                                               string     `json:"LabelTypeID"`
	// A Boolean label value. Only populated if the associated labelType data.ValueType is               
	// 'boolean'. The label type's ValueCount must be 1.                                                 
	ValueAsBoolean                                                                            *bool      `json:"ValueAsBoolean,omitempty"`
	// A Boolean label value. Only populated if the associated labelType data.ValueType is               
	// 'boolean'. The label type's ValueCount is greater than 1.                                         
	ValueAsBooleanArray                                                                       []bool     `json:"ValueAsBooleanArray,omitempty"`
	// A label value as date or date time. Only populated if the associated labelType                    
	// data.ValueType is 'string'. The label type's ValueCount must be 1.                                
	ValueAsDateTime                                                                           *time.Time `json:"ValueAsDateTime,omitempty"`
	// A label value as date or date time. Only populated if the associated labelType                    
	// data.ValueType is 'string'. The label type's ValueCount is greater than 1.                        
	ValueAsDateTimeArray                                                                      []string   `json:"ValueAsDateTimeArray,omitempty"`
	// An integer label value. Only populated if the associated labelType data.ValueType is              
	// 'integer'. The label type's ValueCount must be 1.                                                 
	ValueAsInteger                                                                            *int64     `json:"ValueAsInteger,omitempty"`
	// An integer label value. Only populated if the associated labelType data.ValueType is              
	// 'integer'. The label type's ValueCount is greater than 1.                                         
	ValueAsIntegerArray                                                                       []int64    `json:"ValueAsIntegerArray,omitempty"`
	// A numeric, floating point label value. Only populated if the associated labelType                 
	// data.ValueType is 'number'. The label type's ValueCount must be 1.                                
	ValueAsNumber                                                                             *float64   `json:"ValueAsNumber,omitempty"`
	// A numeric, floating point label value. Only populated if the associated labelType                 
	// data.ValueType is 'number'. The label type's ValueCount is greater than 1.                        
	ValueAsNumberArray                                                                        []float64  `json:"ValueAsNumberArray,omitempty"`
	// A label relationship to another record via its record id:version, where version is                
	// optional. The expected target type is defined in labelType.RelationshipTargetKind. Only           
	// populated if the associated labelType data.ValueType is 'string'. The label type's                
	// ValueCount must be 1.                                                                             
	ValueAsRelatedID                                                                          *string    `json:"ValueAsRelatedID,omitempty"`
	// A label relationship to another record via its record id:version, where version is                
	// optional. The expected target type is defined in labelType.RelationshipTargetKind. Only           
	// populated if the associated labelType data.ValueType is 'string'. The label type's                
	// ValueCount is greater than 1.                                                                     
	ValueAsRelatedIDs                                                                         []string   `json:"ValueAsRelatedIDs,omitempty"`
	// A textual label value. Only populated if the associated labelType data.ValueType is               
	// 'string'. Not to be used for relationships, see ValueAsRelatedID. The label type's                
	// ValueCount must be 1.                                                                             
	ValueAsString                                                                             *string    `json:"ValueAsString,omitempty"`
	// A textual label value. Only populated if the associated labelType data.ValueType is               
	// 'string'. Not to be used for relationships, see ValueAsRelatedID. The label type's                
	// ValueCount is greater than 1.                                                                     
	ValueAsStringArray                                                                        []string   `json:"ValueAsStringArray,omitempty"`
}
