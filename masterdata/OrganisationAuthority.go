package masterdata

// Describes the authority held by a business associate to make payments, sign contracts
// etc. Considered in a business context. Application or database authorities are held in
// ENTITLEMENTS.
type OrganisationAuthority struct {
	// The unique identifier of the BusinessAssociate who authorized this level of authority.           
	AuthorisedByID                                                                              *string `json:"AuthorisedByID,omitempty"`
	// Unique identifier for the row that describes the authority that a business associate has         
	// over business objects.                                                                           
	AuthorityID                                                                                 *string `json:"AuthorityID,omitempty"`
	// The type of authority given to a business associate, often an employee of a company.             
	// Authority may be extended for purchase authorizations, to sign contracts or to enter into        
	// negotiations etc.                                                                                
	AuthorityTypeID                                                                             *string `json:"AuthorityTypeID,omitempty"`
	// Date on which this Authority came into effect.                                                   
	EffectiveDate                                                                               *string `json:"EffectiveDate,omitempty"`
	// A flag indicating whether this data is currently either active / valid (True) or inactive        
	// / invalid (False).                                                                               
	IsActive                                                                                    *bool   `json:"IsActive,omitempty"`
	// Narrative remarks about this BusinessAssociate Authority.                                        
	Remark                                                                                      *string `json:"Remark,omitempty"`
	// The unique identifier of the BusinessAssociate who the authorized person represents for          
	// this authority. May be a subsidiary, a company etc.                                              
	RepresentedByID                                                                             *string `json:"RepresentedByID,omitempty"`
	// Date on which this Authority was no longer in effect.                                            
	TerminationDate                                                                             *string `json:"TerminationDate,omitempty"`
}
