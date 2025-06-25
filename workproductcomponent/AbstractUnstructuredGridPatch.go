package workproductcomponent

// Definition of a collection of polyhedra which are not organized in any dimension. Only an
// abstract group type in order to be reused in GPGrid.
type AbstractUnstructuredGridPatch struct {
	// Indicates the uniform shape of all cells in this grid : tetrahedral, pyramidal, prism,        
	// hexahedral, polyhedral                                                                        
	CellShapeID                                                                              *string `json:"CellShapeID,omitempty"`
	// The count of faces in this grid                                                               
	FaceCount                                                                                *int64  `json:"FaceCount,omitempty"`
	// The count of nodes in this grid                                                               
	NodeCount                                                                                *int64  `json:"NodeCount,omitempty"`
}
