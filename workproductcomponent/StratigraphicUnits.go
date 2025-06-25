package workproductcomponent

// Allow to link the K layers (or the "geologic k" property for example in case of K
// expansion or unstructured grid) of this grid with some stratigraphic units of a
// stratigraphic organization.
type StratigraphicUnits struct {
	// Reference to the stratigraphic column rank interpretation which this grid is derived from.          
	StratigraphicColumnRankInterpretationID                                                      string    `json:"StratigraphicColumnRankInterpretationID"`
	// For each K layer, indicate the corresponding stratigraphic unit indices within the                  
	// associated Stratigraphic Organization. A negative value means that a K layer is not                 
	// related to any stratigraphic unit (salt for example)                                                
	StratigraphicUnitsIndices                                                                    [][]int64 `json:"StratigraphicUnitsIndices"`
}
