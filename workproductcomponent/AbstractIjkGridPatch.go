package workproductcomponent

// Definition of a collection of hexahedra which are organized and indexable by means of 3
// (almost orthogonal) directions and called I, J and K. Only an abstract group type in
// order to be reused in GPGrid.
//
// The group of elements which all column grid layers (ijk or unstructured) contain
type AbstractIjkGridPatch struct {
	// Indicate if a grid has been topologically expanded in a particular dimension (K                                 
	// expansion, J expansion, I expansion)                                                                            
	ExpansionInDirection                                                                         *ExpansionInDirection `json:"ExpansionInDirection,omitempty"`
	// Indicate if at least two adjacent nodes in K Direction are collocated. Also known as                            
	// pinched node. Usually occur in erosional context.                                                               
	HasCollocatedNodeInKDirection                                                                *bool                 `json:"HasCollocatedNodeInKDirection,omitempty"`
	// Indicate that it exists at least one gap in the whole K direction of the grid. A gap is                         
	// really a hole, it is not a layer of dead cells.                                                                 
	HasKGaps                                                                                     *bool                 `json:"HasKGaps,omitempty"`
	// Indicate that it exists at least one gap in the lateral direction of the grid. A gap is                         
	// really a hole, it is not a slice of dead cells.                                                                 
	HasLateralGaps                                                                               *bool                 `json:"HasLateralGaps,omitempty"`
	// Indicate that the nodes of the grid are given by means of a parameter along the pillar.                         
	// Otherwise nodes of the grid are explicitly given by means of an XYZ triplet.                                    
	HasParametricGeometry                                                                        *bool                 `json:"HasParametricGeometry,omitempty"`
	// Indicate that the grid contains some split nodes i.e some node which are not on a pillar.                       
	HasSplitNode                                                                                 *bool                 `json:"HasSplitNode,omitempty"`
	// Indicate that some of the pillars of the grid are truncated (Fault contact in Y shape for                       
	// example)                                                                                                        
	HasTruncations                                                                               *bool                 `json:"HasTruncations,omitempty"`
	// Indicate the K direction of the grid : up, down or not monotonic                                                
	KDirectionID                                                                                 *string               `json:"KDirectionID,omitempty"`
	// Count of cells in the K-direction (aka third and/or slowest and/or vertical direction) in                       
	// the grid. Must be positive.                                                                                     
	Nk                                                                                           *int64                `json:"Nk,omitempty"`
	// Indicate the most complex pillar shape of a grid : vertical, straight, curved                                   
	PillarShapeID                                                                                *string               `json:"PillarShapeID,omitempty"`
	// TRUE if the grid is periodic in J, i.e., has the topology of a complete 360 degree circle.                      
	IsRadial                                                                                     *bool                 `json:"IsRadial,omitempty"`
	// Indicates that the IJK grid is right handed, as determined by the triple product of                             
	// tangent vectors in the I, J, and K directions.                                                                  
	IsRightHanded                                                                                *bool                 `json:"IsRightHanded,omitempty"`
	// Count of cells in the I-direction (aka first and/or fastest direction) in the grid. Must                        
	// be positive.                                                                                                    
	NI                                                                                           *int64                `json:"Ni,omitempty"`
	// Count of cells in the J-direction (aka second direction) in the grid. Must be positive.                         
	Nj                                                                                           *int64                `json:"Nj,omitempty"`
}
