package masterdata

import "time"

// General parameters defining the configuration of the Project.  In the case of a seismic
// acquisition project it is like receiver interval, source depth, source type.  In the case
// of a processing project, it is like replacement velocity, reference datum above mean sea
// level.
type ProjectSpecifications struct {
	// The date and time at which a ProjectSpecification becomes effective.                            
	EffectiveDateTime                                                                       *time.Time `json:"EffectiveDateTime,omitempty"`
	// Parameter type of property or characteristic.                                                   
	ParameterTypeID                                                                         *string    `json:"ParameterTypeID,omitempty"`
	// The actual date and time value of the parameter.  ISO format permits specification of           
	// time or date only.                                                                              
	ProjectSpecificationDateTime                                                            *time.Time `json:"ProjectSpecificationDateTime,omitempty"`
	// The actual indicator value of the parameter.                                                    
	ProjectSpecificationIndicator                                                           *bool      `json:"ProjectSpecificationIndicator,omitempty"`
	// The value for the specified parameter type.                                                     
	ProjectSpecificationQuantity                                                            *float64   `json:"ProjectSpecificationQuantity,omitempty"`
	// The actual text value of the parameter.                                                         
	ProjectSpecificationText                                                                *string    `json:"ProjectSpecificationText,omitempty"`
	// The date and time at which a ProjectSpecification is no longer in effect.                       
	TerminationDateTime                                                                     *time.Time `json:"TerminationDateTime,omitempty"`
	// The unit for the quantity parameter if overriding the default for this ParameterType,           
	// like metre (m in SI units system) for quantity Length.                                          
	UnitOfMeasureID                                                                         *string    `json:"UnitOfMeasureID,omitempty"`
}
