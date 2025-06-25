package workproductcomponent

// A Routine Core Analysis Measurement.
type RCAMeasurement struct {
	// The conditions under which this analysis has been carried out                                                    
	Conditions                                                                                  *Conditions             `json:"Conditions,omitempty"`
	// Other measurements made in the context of this analysis                                                          
	OtherMeasurements                                                                           []AbstractSpecification `json:"OtherMeasurements,omitempty"`
	// The measured sample permeability.                                                                                
	Permeability                                                                                *float64                `json:"Permeability,omitempty"`
	// The measurement type used to obtain the Permeability values.                                                     
	PermeabilityMeasurementTypeID                                                               *string                 `json:"PermeabilityMeasurementTypeID,omitempty"`
	// The measured sample porosity.                                                                                    
	Porosity                                                                                    *float64                `json:"Porosity,omitempty"`
	// The measurement type used to obtain the Porosity values.                                                         
	PorosityMeasurementTypeID                                                                   *string                 `json:"PorosityMeasurementTypeID,omitempty"`
	// An array containing operational or quality comments pertaining to a rock sample analysis.                        
	Remarks                                                                                     []RCAMeasurementRemark  `json:"Remarks,omitempty"`
}
