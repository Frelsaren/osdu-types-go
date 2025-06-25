package workproductcomponent

// Value of the column. Generally only one of the attribute should be instantiated.
type ColumnValues struct {
	// A column of only boolean values                                                               
	BooleanColumn                                                                          []bool    `json:"BooleanColumn,omitempty"`
	// A column of only integer values                                                               
	IntegerColumn                                                                          []int64   `json:"IntegerColumn,omitempty"`
	// A column of only number values                                                                
	NumberColumn                                                                           []float64 `json:"NumberColumn,omitempty"`
	// A column of only string values                                                                
	StringColumn                                                                           []string  `json:"StringColumn,omitempty"`
	// The row indexes for which the values are flagged as undefined. The first element has          
	// index 0.                                                                                      
	UndefinedValueRows                                                                     []int64   `json:"UndefinedValueRows,omitempty"`
}
