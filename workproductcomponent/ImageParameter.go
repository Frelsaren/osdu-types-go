package workproductcomponent

type ImageParameter struct {
	// The image parameter value of type "Boolean" under which the image was acquired.                  
	BooleanParameterValue                                                                      *bool    `json:"BooleanParameterValue,omitempty"`
	// The type of parameter under which the image was acquired. In the array of                        
	// data.ImageParameters[] objects, a particular ImageParameterTypeID value should appear at         
	// most once.                                                                                       
	ImageParameterTypeID                                                                       *string  `json:"ImageParameterTypeID,omitempty"`
	// The image parameter value of type "Numeric" under which the image was acquired.                  
	NumericParameterValue                                                                      *float64 `json:"NumericParameterValue,omitempty"`
	// The image parameter value of type "String" under which the image was acquired.                   
	StringParameterValue                                                                       *string  `json:"StringParameterValue,omitempty"`
}
