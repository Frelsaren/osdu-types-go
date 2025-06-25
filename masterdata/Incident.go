package masterdata

import "time"

// Description of incidents that have occurred during the last drilling/operations report
type Incident struct {
	// Accident description.                                                                            
	AccidentDescription                                                                       *string   `json:"AccidentDescription,omitempty"`
	// Cause description.                                                                               
	CauseDescription                                                                          *string   `json:"CauseDescription,omitempty"`
	// Comments and remarks                                                                             
	Comments                                                                                  *string   `json:"Comments,omitempty"`
	// Gross estimate of the cost incurred due to the incident.                                         
	CostLossEstimate                                                                          *float64  `json:"CostLossEstimate,omitempty"`
	// Date and time that incident occurred                                                             
	DateTime                                                                                  time.Time `json:"DateTime"`
	// Number of hours lost due to the incident.                                                        
	EstimateHoursLost                                                                         *float64  `json:"EstimateHoursLost,omitempty"`
	// Name of the person who prepared the incident report.                                             
	IncidentReporterName                                                                      *string   `json:"IncidentReporterName,omitempty"`
	// Near miss incident occurrence?                                                                   
	IsNearMiss                                                                                *bool     `json:"IsNearMiss,omitempty"`
	// Location description.                                                                            
	LocationDescription                                                                       *string   `json:"LocationDescription,omitempty"`
	// A reference to the organisation for which is the company primarily involved in managing          
	// the incident.                                                                                    
	OrganisationID                                                                            *string   `json:"OrganisationID,omitempty"`
	// Remedial action description.                                                                     
	RemedialActionDescription                                                                 *string   `json:"RemedialActionDescription,omitempty"`
	// Number of personnel killed due to the incident.                                                  
	TotalFatality                                                                             *int64    `json:"TotalFatality,omitempty"`
	// Number of personnel with major injuries.                                                         
	TotalMajorInjury                                                                          *int64    `json:"TotalMajorInjury,omitempty"`
	// Number of personnel with minor injuries.                                                         
	TotalMinorInjury                                                                          *int64    `json:"TotalMinorInjury,omitempty"`
}
