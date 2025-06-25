package masterdata

import "time"

// General parameter value used in one instance of activity.
// [Without inheritance, combined specializations.]
type PurpleAbstractActivityParameter struct {
	// Parameter referencing to a top level object.                                                                  
	DataObjectParameter                                                                         *string              `json:"DataObjectParameter,omitempty"`
	// Parameter containing a double value.                                                                          
	DataQuantityParameter                                                                       *float64             `json:"DataQuantityParameter,omitempty"`
	// Identifies unit of measure for floating point value.                                                          
	DataQuantityParameterUOMID                                                                  *string              `json:"DataQuantityParameterUOMID,omitempty"`
	// When parameter is an array, used to indicate the index in the array.                                          
	Index                                                                                       *int64               `json:"Index,omitempty"`
	// Parameter containing an integer value.                                                                        
	IntegerQuantityParameter                                                                    *int64               `json:"IntegerQuantityParameter,omitempty"`
	// A nested array describing keys used to identify a parameter value. When multiple values                       
	// are provided for a given parameter, the key provides a way to identify the parameter                          
	// through its association with an object, a time index or a parameter array member via                          
	// ParameterKey value.                                                                                           
	Keys                                                                                        []PurpleParameterKey `json:"Keys,omitempty"`
	// [Added to cover lack of inheritance]                                                                          
	ParameterKindID                                                                             string               `json:"ParameterKindID"`
	// Reference data describing how the parameter was used by the activity, such as input,                          
	// output, control, constraint, agent, predecessor activity, successor activity.                                 
	ParameterRoleID                                                                             *string              `json:"ParameterRoleID,omitempty"`
	// Textual description about how this parameter was selected.                                                    
	Selection                                                                                   *string              `json:"Selection,omitempty"`
	// Parameter containing a string value.                                                                          
	StringParameter                                                                             *string              `json:"StringParameter,omitempty"`
	// Parameter containing a time index value.  It is assumed that all TimeIndexParameters                          
	// within an Activity have the same date-time format, which is then described by the                             
	// FrameOfReference mechanism.                                                                                   
	TimeIndexParameter                                                                          *time.Time           `json:"TimeIndexParameter,omitempty"`
	// Name of the parameter, used to identify it in the activity. It must have an equivalent in                     
	// the ActivityTemplate parameters.                                                                              
	Title                                                                                       string               `json:"Title"`
}
