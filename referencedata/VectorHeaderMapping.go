package referencedata

// Array of objects which define the meaning and format of a tabular structure used in a
// binary file as a header.  The initial use case is the trace headers of a SEG-Y file.
// Note that some of this information may be repeated in the SEG-Y EBCDIC header.
type VectorHeaderMapping struct {
	// Relationship to a reference value for a name of a property header such as INLINE, CDPX.                 
	KeyName                                                                                   *string          `json:"KeyName,omitempty"`
	// Beginning byte position of header value, 1 indexed.                                                     
	Position                                                                                  *int64           `json:"Position,omitempty"`
	// Enumerated string indicating whether to use the normal scalar field for scaling this                    
	// field (STANDARD), no scaling (NOSCALE), or override scalar (OVERRIDE).  Default is                      
	// current STANDARD (such as SEG-Y rev2).                                                                  
	ScalarIndicator                                                                           *ScalarIndicator `json:"ScalarIndicator,omitempty"`
	// Scalar value (as defined by standard) when a value present in the header needs to be                    
	// overwritten for this value.                                                                             
	ScalarOverride                                                                            *float64         `json:"ScalarOverride,omitempty"`
	// Relationship to units of measure reference if header standard is not followed.                          
	UoM                                                                                       *string          `json:"UoM,omitempty"`
	// Relationship to a reference value for binary data types, such as INT, UINT, FLOAT,                      
	// IBM_FLOAT, ASCII, EBCDIC.                                                                               
	WordFormat                                                                                *string          `json:"WordFormat,omitempty"`
	// Size of the word in bytes.                                                                              
	WordWidth                                                                                 *int64           `json:"WordWidth,omitempty"`
}
