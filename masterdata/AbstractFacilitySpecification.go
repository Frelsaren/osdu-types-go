package masterdata

import "time"

// A property, characteristic, or attribute about a facility that is not described
// explicitly elsewhere.
type AbstractFacilitySpecification struct {
	// The date and time at which the facility specification instance becomes effective.                
	EffectiveDateTime                                                                        *time.Time `json:"EffectiveDateTime,omitempty"`
	// The actual date and time value of the parameter.                                                 
	FacilitySpecificationDateTime                                                            *time.Time `json:"FacilitySpecificationDateTime,omitempty"`
	// The actual indicator value of the parameter.                                                     
	FacilitySpecificationIndicator                                                           *bool      `json:"FacilitySpecificationIndicator,omitempty"`
	// The value for the specified parameter type.                                                      
	FacilitySpecificationQuantity                                                            *float64   `json:"FacilitySpecificationQuantity,omitempty"`
	// The actual text value of the parameter.                                                          
	FacilitySpecificationText                                                                *string    `json:"FacilitySpecificationText,omitempty"`
	// Parameter type of property or characteristic.                                                    
	ParameterTypeID                                                                          *string    `json:"ParameterTypeID,omitempty"`
	// The date and time at which the facility specification instance is no longer in effect.           
	TerminationDateTime                                                                      *time.Time `json:"TerminationDateTime,omitempty"`
	// The unit for the quantity parameter, like metre (m in SI units system) for quantity              
	// Length.                                                                                          
	UnitOfMeasureID                                                                          *string    `json:"UnitOfMeasureID,omitempty"`
}
