package workproductcomponent

// Describe a cell which is a part or the whole definition of an aquifer.
type SingleCellAquifer struct {
	// Area of the aquifer.                                                                              
	Area                                                                                        *float64 `json:"Area,omitempty"`
	// Depth of the aquifer                                                                              
	Depth                                                                                       *float64 `json:"Depth,omitempty"`
	// I index of the cell                                                                               
	I                                                                                           *int64   `json:"I,omitempty"`
	// Initial pressure of the aquifer.  Default property, caution if the grid property has got          
	// multi realizations.                                                                               
	InitialPressure                                                                             *float64 `json:"InitialPressure,omitempty"`
	// J index of the cell                                                                               
	J                                                                                           *int64   `json:"J,omitempty"`
	// K index of the cell                                                                               
	K                                                                                           *int64   `json:"K,omitempty"`
	// Length of the aquifer.                                                                            
	Length                                                                                      *float64 `json:"Length,omitempty"`
	// Permeability of the aquifer. Default property, caution if the grid property has got multi         
	// realizations.                                                                                     
	Permeability                                                                                *float64 `json:"Permeability,omitempty"`
	// Porosity of the aquifer.  Default property, caution if the grid property has got multi            
	// realizations.                                                                                     
	Porosity                                                                                    *float64 `json:"Porosity,omitempty"`
	// The grid where the aquifer cell is defined                                                        
	SupportingGridID                                                                            *string  `json:"SupportingGridID,omitempty"`
}
