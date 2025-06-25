package workproductcomponent

// Definition of an activity code alias proprietary to a Company
type ProprietaryActivityCodeElement struct {
	// Authority property          
	Authority              string  `json:"Authority"`
	// Description Property        
	Description            *string `json:"Description,omitempty"`
	// Identifier property         
	Identifier             string  `json:"Identifier"`
}
