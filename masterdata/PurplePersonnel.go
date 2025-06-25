package masterdata

// List of key individuals supporting the Project.  This could be Abstracted for re-use, and
// could reference a separate Persons master data object.
type PurplePersonnel struct {
	// Reference to the company which employs Personnel.                                              
	CompanyOrganisationID                                                                     *string `json:"CompanyOrganisationID,omitempty"`
	// Name of an individual supporting the Project.                                                  
	PersonName                                                                                *string `json:"PersonName,omitempty"`
	// The identifier of a reference value for the role of an individual supporting a Project,        
	// such as Project Manager, Party Chief, Client Representative, Senior Observer.                  
	ProjectRoleID                                                                             *string `json:"ProjectRoleID,omitempty"`
}
