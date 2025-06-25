package workproductcomponent

// Array of curve that constitutes the whole PPFG Dataset
type PurpleCurves struct {
	// IDs of the type and level of processing that has been applied to the curve. An array of           
	// curve processing operations that have been applied, for example 'Smoothed', 'Calibrated',         
	// etc                                                                                               
	CurveDataProcessingTypeIDs                                                                  []string `json:"CurveDataProcessingTypeIDs,omitempty"`
	// ID of the PPFG Curve Family of the PPFG quantity measured, for example 'Pore Pressure             
	// from Corrected Drilling Exponent'. An individual curve that belongs to a Main Family              
	CurveFamilyID                                                                               *string  `json:"CurveFamilyID,omitempty"`
	// ID of the mnemonic of the Curve Family which is the value as received either from                 
	// external providers or from internal processing team, for example 'PP DxC'                         
	CurveFamilyMnemonicID                                                                       *string  `json:"CurveFamilyMnemonicID,omitempty"`
	// The ID of the PPFG Curve                                                                          
	CurveID                                                                                     *string  `json:"CurveID,omitempty"`
	// ID of the lithological unit represented by the curve                                              
	CurveLithologyID                                                                            *string  `json:"CurveLithologyID,omitempty"`
	// ID of the Main Family Type of the PPFG quantity measured, for example 'Pore Pressure'.            
	// Primarily used for high level curve classification                                                
	CurveMainFamilyID                                                                           *string  `json:"CurveMainFamilyID,omitempty"`
	// The original or as supplied PPFG curve name. Intended to hold historical or legacy                
	// information                                                                                       
	CurveName                                                                                   *string  `json:"CurveName,omitempty"`
	// ID of the PPFG Curve probability, for example 'Most Likely Case' and 'P50'                        
	CurveProbabilityID                                                                          *string  `json:"CurveProbabilityID,omitempty"`
	// ID of the empirical calibrated model used for pressure calculations from a petrophysical          
	// curve (sonic or resistivity logs), for example 'Eaton' and  'Bowers',...                          
	CurveTransformModelTypeID                                                                   *string  `json:"CurveTransformModelTypeID,omitempty"`
	// Unit of Measure of the Physical Quantity Measured by the curve. An ID to relevant unit of         
	// measure reference                                                                                 
	CurveUOM                                                                                    *string  `json:"CurveUOM,omitempty"`
}
