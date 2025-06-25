package masterdata

// Definition of an alias
type WellboreAliasElement struct {
	// Authority property          
	Authority              string  `json:"Authority"`
	// Description Property        
	Description            *string `json:"Description,omitempty"`
	// Identifier property         
	Identifier             string  `json:"Identifier"`
}
