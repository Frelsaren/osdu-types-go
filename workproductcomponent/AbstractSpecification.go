package workproductcomponent

import "time"

// An array element of an analysis specifying the numeric, text or date-time value
// associated with a ParameterType.
type AbstractSpecification struct {
	// The date and time at which this specification becomes effective.                              
	EffectiveDateTime                                                                     *time.Time `json:"EffectiveDateTime,omitempty"`
	// Parameter representing a property or characteristic.                                          
	ParameterTypeID                                                                       *string    `json:"ParameterTypeID,omitempty"`
	// The date value of the parameter.                                                              
	SpecificationDate                                                                     *string    `json:"SpecificationDate,omitempty"`
	// The date and time value of the parameter.                                                     
	SpecificationDateTime                                                                 *time.Time `json:"SpecificationDateTime,omitempty"`
	// The indicator value of the parameter.                                                         
	SpecificationIndicator                                                                *bool      `json:"SpecificationIndicator,omitempty"`
	// The numeric value of the parameter.                                                           
	SpecificationQuantity                                                                 *float64   `json:"SpecificationQuantity,omitempty"`
	// The text value of the parameter.                                                              
	SpecificationText                                                                     *string    `json:"SpecificationText,omitempty"`
	// The time value of the parameter.                                                              
	SpecificationTime                                                                     *string    `json:"SpecificationTime,omitempty"`
	// The first point in time when this specification is no longer in effect.                       
	TerminationDateTime                                                                   *time.Time `json:"TerminationDateTime,omitempty"`
	// The unit for the quantity parameter, like metre (m in SI units system) for quantity           
	// Length.                                                                                       
	UnitOfMeasureID                                                                       *string    `json:"UnitOfMeasureID,omitempty"`
}
