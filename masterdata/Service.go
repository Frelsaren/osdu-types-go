package masterdata

// Describes the primary services provided by a business associate. For example drilling
// contractor, logging company, seismic broker etc.
type Service struct {
	// The date on which the service came into effect.                                            
	EffectiveDate                                                                         *string `json:"EffectiveDate,omitempty"`
	// A flag indicating whether this Service is currently either active/valid (True) or          
	// inactive/invalid (False).                                                                  
	IsActive                                                                              *bool   `json:"IsActive,omitempty"`
	// Narrative remarks about this service.                                                      
	Remark                                                                                *string `json:"Remark,omitempty"`
	// A code indicating the quality of service received from this service.                       
	ServiceQualityTypeID                                                                  *string `json:"ServiceQualityTypeID,omitempty"`
	// The service that a business associate provides. For example well logging, drilling,        
	// application development.                                                                   
	ServiceTypeID                                                                         *string `json:"ServiceTypeID,omitempty"`
	// The date on which this service was no longer in effect.                                    
	TerminationDate                                                                       *string `json:"TerminationDate,omitempty"`
}
