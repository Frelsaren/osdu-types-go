package workproductcomponent

// Diversion details for the stimulated interval.
type Diversion struct {
	// Name of the diversion contractor.                                                  
	ContractorName                                                               *string  `json:"ContractorName,omitempty"`
	// The diversion method used for this job stage.                                      
	DiversionMethodID                                                            *string  `json:"DiversionMethodID,omitempty"`
	// Length between packer elements.                                                    
	ElementSpacing                                                               *float64 `json:"ElementSpacing,omitempty"`
	// Remarks on the diversion method.                                                   
	Remarks                                                                      *string  `json:"Remarks,omitempty"`
	// A supplier description of the diversion tool, such as its commercial name.         
	ToolDescription                                                              *string  `json:"ToolDescription,omitempty"`
}
