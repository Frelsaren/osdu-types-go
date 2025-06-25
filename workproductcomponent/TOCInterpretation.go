package workproductcomponent

// TOC Interpretation
type TOCInterpretation struct {
	// TOC Interpretation                                    
	Interpretation                                  *string  `json:"Interpretation,omitempty"`
	// Comments or Remarks of the TOC Interpretation         
	InterpretationRemarks                           *string  `json:"InterpretationRemarks,omitempty"`
	// Interpreter                                           
	Interpreter                                     *string  `json:"Interpreter,omitempty"`
	// Remedial Cement Required?                             
	IsRemedialCementRequired                        *bool    `json:"IsRemedialCementRequired,omitempty"`
	// TOC Sufficient?                                       
	IsTopOfCementSufficient                         *bool    `json:"IsTopOfCementSufficient,omitempty"`
	// Number of remedial Squeezes required                  
	NumberOfRemedial                                *int64   `json:"NumberOfRemedial,omitempty"`
	// Planned Top of Cement (TOC) Measured Depth            
	PlannedTopMeasuredDepth                         *float64 `json:"PlannedTopMeasuredDepth,omitempty"`
	// Remedial Type                                         
	RemedialType                                    *string  `json:"RemedialType,omitempty"`
	// Top of Cement Locate Method                           
	TopOfCementLocateMethodID                       *string  `json:"TopOfCementLocateMethodID,omitempty"`
	// Top of Cement (TOC) Measured Depth                    
	TopOfCementMeasuredDepth                        *float64 `json:"TopOfCementMeasuredDepth,omitempty"`
	// Top of Cement (TOC) True Vertical Depth               
	TopOfCementTrueVerticalDepth                    *float64 `json:"TopOfCementTrueVerticalDepth,omitempty"`
}
