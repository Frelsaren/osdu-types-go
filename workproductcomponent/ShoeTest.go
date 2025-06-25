package workproductcomponent

// Shoe Test
type ShoeTest struct {
	// Shoe Test Tool Used                                               
	IsShoeTestToolUsed                                          *bool    `json:"IsShoeTestToolUsed,omitempty"`
	// Open Hole Length                                                  
	OpenHoleLength                                              *float64 `json:"OpenHoleLength,omitempty"`
	// Comments or Remarks                                               
	ShoeTestComments                                            *string  `json:"ShoeTestComments,omitempty"`
	// Elapsed Time Before Shoe Test following end of cement job         
	ShoeTestElapsedTime                                         *float64 `json:"ShoeTestElapsedTime,omitempty"`
	// Shoe Test Equivalent Mud Weight                                   
	ShoeTestEquivalentMudWeight                                 *float64 `json:"ShoeTestEquivalentMudWeight,omitempty"`
	// Shoe Test Type                                                    
	ShoeTestTypeID                                              *string  `json:"ShoeTestTypeID,omitempty"`
}
