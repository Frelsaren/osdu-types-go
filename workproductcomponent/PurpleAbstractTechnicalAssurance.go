package workproductcomponent

// Describes a record's overall suitability for general business consumption based on level
// of trust.
type PurpleAbstractTechnicalAssurance struct {
	// List of workflows and/or personas that the technical assurance value is valid for (e.g.,                           
	// This data is trusted for Seismic Processing)                                                                       
	AcceptableUsage                                                                             []FluffyAcceptableUsage   `json:"AcceptableUsage,omitempty"`
	// Any additional context to support the determination of technical assurance                                         
	Comment                                                                                     *string                   `json:"Comment,omitempty"`
	// Date when the technical assurance determination for this record has taken place                                    
	EffectiveDate                                                                               *string                   `json:"EffectiveDate,omitempty"`
	// The individuals, or roles, that reviewed and determined the technical assurance value                              
	Reviewers                                                                                   []AbstractContact         `json:"Reviewers,omitempty"`
	// Describes a record's overall suitability for general business consumption based on data                            
	// quality. Clarifications: Since Certified is the highest classification of suitable                                 
	// quality, any further change or versioning of a Certified record should be carefully                                
	// considered and justified. If a Technical Assurance value is not populated then one can                             
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                   
	// Technical Assurance values are not intended to be used for the identification of a single                          
	// "preferred" or "definitive" record by comparison with other records.                                               
	TechnicalAssuranceTypeID                                                                    string                    `json:"TechnicalAssuranceTypeID"`
	// List of workflows and/or personas that the technical assurance value is not valid for                              
	// (e.g., This data is not trusted for seismic interpretation)                                                        
	UnacceptableUsage                                                                           []FluffyUnacceptableUsage `json:"UnacceptableUsage,omitempty"`
}
