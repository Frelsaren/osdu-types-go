package workproductcomponent

// Processing Parameters to simply capture process history until full provenance model can
// be implemented.
type PurpleProcessingParameters struct {
	// Processing Parameter Type                                                                
	ProcessingParameterTypeID                                                           *string `json:"ProcessingParameterTypeID,omitempty"`
	// The quantity for the processing parameter. May include units, ordering, and other        
	// descriptions.                                                                            
	ProcessingParameterValue                                                            *string `json:"ProcessingParameterValue,omitempty"`
}
