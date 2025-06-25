package masterdata

// Description of one parameter that participates in one type of activity. [Without
// inheritance, combined specializations.]
type ParameterTemplate struct {
	// If no allowed kind is given, then all kind of data types are allowed.                                          
	AllowedParameterKind                                                                        *string               `json:"AllowedParameterKind,omitempty"`
	// Textual description of additional constraint associated with the parameter. (note that it                      
	// will be better to have a formal description of the constraint)                                                 
	Constraint                                                                                  *string               `json:"Constraint,omitempty"`
	// When parameter is limited to data object of given types, describe the allowed types. Used                      
	// only when ParameterType is dataObject.  String is an OSDU kind of work product component.                      
	DataObjectContentType                                                                       []string              `json:"DataObjectContentType,omitempty"`
	// Activity Parameter value to use if one not supplied.                                                           
	DefaultValue                                                                                *DefaultValueElement  `json:"DefaultValue,omitempty"`
	// Indicates if the parameter is an input of the activity. If the parameter is a data object                      
	// and is also an output of the activity, it is strongly advised to use two parameters : one                      
	// for input and one for output. The reason is to be able to give two different versions                          
	// strings for the input and output data object which has got obviously the same UUID.                            
	IsInput                                                                                     bool                  `json:"IsInput"`
	// Indicates if the parameter is an output of the activity. If the parameter is a data                            
	// object and is also an input of the activity, it is strongly advised to use two parameters                      
	// : one for input and one for output. The reason is to be able to give two different                             
	// versions strings for the input and output data object which has got obviously the same                         
	// UUID.                                                                                                          
	IsOutput                                                                                    bool                  `json:"IsOutput"`
	// Allows to indicate that, in the same activity, this parameter template must be associated                      
	// to another parameter template identified by its title. The associated parameter value                          
	// constrains this parameter.                                                                                     
	KeyConstraints                                                                              []string              `json:"KeyConstraints,omitempty"`
	// Maximum number of parameters of this type allowed in the activity. If the maximum number                       
	// of parameters is infinite, use -1 value.                                                                       
	MaxOccurs                                                                                   int64                 `json:"MaxOccurs"`
	// Minimum number of parameter of this type required by the activity. If the minimum number                       
	// of parameters is infinite, use -1 value.                                                                       
	MinOccurs                                                                                   int64                 `json:"MinOccurs"`
	// The property type ID and Name, which determines eventually the UnitQuantity of the                             
	// parameter value. Used to provide a more scoped context than UnitQuantityID. If                                 
	// PropertyType is provided, UnitQuantityID is expected to be omitted.                                            
	PropertyType                                                                                *AbstractPropertyType `json:"PropertyType,omitempty"`
	// Name of the parameter in the activity. Key to identify parameter.                                              
	Title                                                                                       string                `json:"Title"`
	// The expected UnitQuantity for the parameter value. A more precise context can be                               
	// provided by PropertyType. If UnitQuantityID is provided, PropertyType is expected to be                        
	// omitted.                                                                                                       
	UnitQuantityID                                                                              *string               `json:"UnitQuantityID,omitempty"`
}
