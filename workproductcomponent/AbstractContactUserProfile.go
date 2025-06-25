package workproductcomponent

// A schema fragment, which can hold explicit contact information (contact by value) and a
// relationship to a UserProfile (contact by reference). Both ways of capturing contact
// information can be used simultaneously : the 'by value' may capture the state of the
// contact at the time of recording, while the 'by reference' relationship to UserProfile
// captures the current state of the contact.
//
// An object with properties that describe a specific person or other point-of-contact (like
// an email distribution list) that is relevant in this context (like a given data set or
// business project). The contact specified may be either internal or external to the
// organisation (something denoted via the Organisation object that is referenced). Note
// that some properties contain personally identifiable information, so it might not be
// appropriate to populate all properties in all scenarios.
type AbstractContactUserProfile struct {
	// Additional information about the contact                                                         
	Comment                                                                                     *string `json:"Comment,omitempty"`
	// The data governance role assigned to this contact if and only if the context has a data          
	// governance role (in context of TechnicalAssurance). The value is kept absent in all other        
	// cases.                                                                                           
	DataGovernanceRoleTypeID                                                                    *string `json:"DataGovernanceRoleTypeID,omitempty"`
	// Contact email address. Property may be left empty where it is inappropriate to provide           
	// personally identifiable information.                                                             
	EmailAddress                                                                                *string `json:"EmailAddress,omitempty"`
	// Name of the individual contact. Property may be left empty where it is inappropriate to          
	// provide personally identifiable information.                                                     
	Name                                                                                        *string `json:"Name,omitempty"`
	// Reference to the company the contact is associated with.                                         
	OrganisationID                                                                              *string `json:"OrganisationID,omitempty"`
	// Contact phone number. Property may be left empty where it is inappropriate to provide            
	// personally identifiable information.                                                             
	PhoneNumber                                                                                 *string `json:"PhoneNumber,omitempty"`
	// The identifier of a reference value for the role of the contact within the associated            
	// organisation, such as Account owner, Sales Representative, Technical Support, Project            
	// Manager, Party Chief, Client Representative, Senior Observer.                                    
	RoleTypeID                                                                                  *string `json:"RoleTypeID,omitempty"`
	// The persona in context of workflows associated with this contact, as used in                     
	// TechnicalAssurance.                                                                              
	WorkflowPersonaTypeID                                                                       *string `json:"WorkflowPersonaTypeID,omitempty"`
	// The contact information by reference to a UserProfile record.                                    
	UserProfileID                                                                               *string `json:"UserProfileID,omitempty"`
}
