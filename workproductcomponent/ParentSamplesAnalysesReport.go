package workproductcomponent

// A pair of properties building a reference to a report combined with a page reference
// inside the document.
type ParentSamplesAnalysesReport struct {
	// An OSDU Record representing the report from which this analysis was extracted or an OSDU        
	// Record from which this analysis was reconstructed.                                              
	ParentSamplesAnalysesReportID                                                              *string `json:"ParentSamplesAnalysesReportID,omitempty"`
	// If this dataset was extracted from a larger report, this identifies the page number or          
	// range (if in document format such as PDF) or sheet names (if in document format such as         
	// spreadsheet). For ease of parsing, ranges should be denoted with hyphens with no spaces,        
	// and divided with commas). E.g. "12-15,17" or "Sheet1,Sheet2".                                   
	ParentSamplesAnalysesReportPageNumbers                                                     *string `json:"ParentSamplesAnalysesReportPageNumbers,omitempty"`
}
