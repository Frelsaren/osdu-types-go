package masterdata

import "time"

type DataGovernanceRole struct {
	// The date and time that this data governance role in this Organisation begins to be            
	// considered active.                                                                            
	DataGovernanceRoleEffectiveDateTime                                                   *time.Time `json:"DataGovernanceRoleEffectiveDateTime,omitempty"`
	// Internal, unique identifier for the set of attributes describing and qualifying a             
	// specific data governance role.                                                                
	DataGovernanceRoleIdentifier                                                          *string    `json:"DataGovernanceRoleIdentifier,omitempty"`
	// The date and time that this data governance role in this Organisation begins to be            
	// considered inactive.                                                                          
	DataGovernanceRoleTerminationDateTime                                                 *time.Time `json:"DataGovernanceRoleTerminationDateTime,omitempty"`
	// Reference to a data governance role assigned to this individual in the associated             
	// Organisation.                                                                                 
	DataGovernanceRoleTypeID                                                              *string    `json:"DataGovernanceRoleTypeID,omitempty"`
	// Reference to the organisation to which this data governance role is associated. For           
	// example, a specific business unit.                                                            
	OrganisationID                                                                        *string    `json:"OrganisationID,omitempty"`
}
