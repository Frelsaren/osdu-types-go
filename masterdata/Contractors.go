package masterdata

// References to organisations which supplied services to the Project.
type Contractors struct {
	// Name of the team, unit, crew, party, or other subdivision of the Contractor that provided        
	// services.                                                                                        
	ContractorCrew                                                                              *string `json:"ContractorCrew,omitempty"`
	// Reference to a company that provided services.                                                   
	ContractorOrganisationID                                                                    *string `json:"ContractorOrganisationID,omitempty"`
	// The identifier of a reference value for the role of a contractor providing services, such        
	// as Recording, Line Clearing, Positioning, Data Processing.                                       
	ContractorTypeID                                                                            *string `json:"ContractorTypeID,omitempty"`
}
