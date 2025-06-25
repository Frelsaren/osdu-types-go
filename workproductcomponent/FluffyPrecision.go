package workproductcomponent

// Sample data format in terms of sample value precision 8bit Integer, 16bit Floating Point
// etc.
type FluffyPrecision struct {
	// SRN of a reference value for binary data types, such as INT, UINT, FLOAT, IBM_FLOAT,        
	// ASCII, EBCDIC.                                                                              
	WordFormat                                                                             *string `json:"WordFormat,omitempty"`
	// Size of the word in bytes.                                                                  
	WordWidth                                                                              *int64  `json:"WordWidth,omitempty"`
}
