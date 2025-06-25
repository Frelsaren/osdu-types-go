package masterdata

// Abstract class describing a key used to identify a parameter value. When multiple values
// are provided for a given parameter, provides a way to identify the parameter through its
// association with an object, a time index, an integer...
// [Without inheritance, combined specializations.] Note: floating point numbers are not
// supported as key values; the numbers have to be formatted as strings for robust equality
// operations, which are necessary for keys.
type DefaultValueKey struct {
	// Integer value from "ParameterKey" parameter, associated with this parameter. Example:          
	// {"ParameterKey": "index", "StringParameterKey: 2}.                                             
	IntegerParameterKey                                                                       *int64  `json:"IntegerParameterKey,omitempty"`
	// Relationship to an object ID, which acts as the parameter.                                     
	ObjectParameterKey                                                                        *string `json:"ObjectParameterKey,omitempty"`
	// The key name, which establishes an association between parameters.                             
	ParameterKey                                                                              *string `json:"ParameterKey,omitempty"`
	// String value from "ParameterKey" parameter, associated with this parameter. Can be used        
	// to associate with parameter values of type string or data quantity. In the later case,         
	// the string representation of the quantity value will be used. Example: {"ParameterKey":        
	// "facies", "StringParameterKey: "shale"}, {"ParameterKey":"depth",                              
	// "StringParameterKey":"1545.43m"}.                                                              
	StringParameterKey                                                                        *string `json:"StringParameterKey,omitempty"`
	// The time index acting as parameter key value.                                                  
	TimeIndexParameterKey                                                                     *string `json:"TimeIndexParameterKey,omitempty"`
}
