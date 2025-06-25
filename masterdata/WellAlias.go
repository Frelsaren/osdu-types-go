package masterdata

// Local name defined for the Well
type WellAlias struct {
	// Authority property          
	Authority              string  `json:"Authority"`
	// Description Property        
	Description            *string `json:"Description,omitempty"`
	// Identifier property         
	Identifier             string  `json:"Identifier"`
}
