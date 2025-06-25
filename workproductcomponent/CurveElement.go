package workproductcomponent

// Describing the characteristics of the "Pressure against Time" recorded curves.
type CurveElement struct {
	// Mnemonic-level curve description is used during parsing or reading and ingesting LAS or          
	// DLIS files, to explain the type of measurement being looked at, specifically for that            
	// moment. Curve description is specific to that single (log) mnemonic and for the entire           
	// log (acquisition run) interval. In essence, curve description defines the internal               
	// factors such as what the "curve" or measurement ideally is representing, how is it               
	// calculated, what are the assumptions and the "constants".                                        
	CurveDescription                                                                            *string `json:"CurveDescription,omitempty"`
	// The ID of the Well Log Curve                                                                     
	CurveID                                                                                     *string `json:"CurveID,omitempty"`
	// The value type to be expected as curve sample values.                                            
	CurveSampleTypeID                                                                           *string `json:"CurveSampleTypeID,omitempty"`
	// Unit of Measure for the Log Curve                                                                
	CurveUnitID                                                                                 *string `json:"CurveUnitID,omitempty"`
	// The Mnemonic of the Log Curve is the value as received either from Raw Providers or from         
	// Internal Processing team                                                                         
	Mnemonic                                                                                    *string `json:"Mnemonic,omitempty"`
	// Indicates that there is no measurement within the curve                                          
	NullValue                                                                                   *bool   `json:"NullValue,omitempty"`
	// The number of columns present in this Curve for a single reference value. For simple logs        
	// this is typically 1; for image logs this holds the number of image traces or property            
	// series. Further information about the columns can be obtained via the respective log or          
	// curve APIs of the Domain Data Management Service.                                                
	NumberOfColumns                                                                             *int64  `json:"NumberOfColumns,omitempty"`
	// Unit of Measure Time Reference                                                                   
	TimeUnitID                                                                                  *string `json:"TimeUnitID,omitempty"`
}
