package masterdata

// Snapshot of operations personnel broken down by each company on the rig at the time of
// the report.
type FluffyPersonnel struct {
	// Comments and remarks                                                                               
	Comments                                                                                     *string  `json:"Comments,omitempty"`
	// Number of persons on board or on location for the selected company/organization.                   
	// HeadCount should be reported as a non-negative integer.                                            
	HeadCount                                                                                    float64  `json:"HeadCount"`
	// A Reference to the organization for which have their headcount measured during this report         
	OrganizationID                                                                               string   `json:"OrganizationID"`
	// Role the organization is playing for the service being provided                                    
	OrganizationRole                                                                             *string  `json:"OrganizationRole,omitempty"`
	// Service provided by the company.                                                                   
	ServiceType                                                                                  *string  `json:"ServiceType,omitempty"`
	// Total time worked by the company (commonly in hours).                                              
	TotalTime                                                                                    *float64 `json:"TotalTime,omitempty"`
}
