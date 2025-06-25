package masterdata

import "time"

// Services for a well license may involve field work, office work, legal representation and
// more.  These services are specific to the license and its management.  Avoid using this
// subset for other services such as logging, coring, testing, drilling and so on.
type LicenseServiceElement struct {
	// The person or organization group who is the primary contact for this service.                       
	ContactBusinessAssociateID                                                                  *string    `json:"ContactBusinessAssociateID,omitempty"`
	// Narrative remarks about the service provided for the license.                                       
	Description                                                                                 *string    `json:"Description,omitempty"`
	// Unique identifier of this element in the parent's list of license services.                         
	ElementIdentifier                                                                           *string    `json:"ElementIdentifier,omitempty"`
	// The time of day on the end date when the service was completed.                                     
	EndTime                                                                                     *time.Time `json:"EndTime,omitempty"`
	// The business associate on whose behalf the service was provided.                                    
	RepresentedBusinessAssociateID                                                              *string    `json:"RepresentedBusinessAssociateID,omitempty"`
	// Service agents work on behalf of a licensee, and must be approved by the regulator. The             
	// service agent has legal authority, and the relationship between a regulator, a licensee             
	// is formal.  If the service agent changes, the regulator must be notified and approval for           
	// the change given. This is the date that the authority who granted the license gave                  
	// consent to change the service agent.                                                                
	ServiceAgentApprovedDate                                                                    *string    `json:"ServiceAgentApprovedDate,omitempty"`
	// The relationship between a licensee and service agent is formal.  When a service agent is           
	// discharged, regulatory notification, and sometimes approval, is required.                           
	ServiceAgentDischargeDate                                                                   *string    `json:"ServiceAgentDischargeDate,omitempty"`
	// The agent through whom a service is brokered or otherwise obtained.  Commonly used where            
	// regulations require the use of local agents, or where confidentiality is a concern.                 
	ServiceAgentID                                                                              *string    `json:"ServiceAgentID,omitempty"`
	// The business associate who provided the service, generally the supplier.                            
	ServiceByBusinessAssociateID                                                                *string    `json:"ServiceByBusinessAssociateID,omitempty"`
	// The business associate for whom this service was provided, generally the client or                  
	// customer.                                                                                           
	ServiceForBusinessAssociate                                                                 *string    `json:"ServiceForBusinessAssociate,omitempty"`
	// This attribute can be used to capture information about the quality of service.  Users of           
	// this attribute should be aware that this information may be considered private or                   
	// privileged, and in some situations data may become public, resulting in violation of                
	// certain laws.                                                                                       
	ServiceQualityTypeID                                                                        *string    `json:"ServiceQualityTypeID,omitempty"`
	// In some cases, it is necessary to order services in the sequence provided, particularly             
	// when there are relationships between the services. Use this number to organize services             
	// in the order that they were done in.  Note that some services are not necessarily                   
	// associated with a date sufficiently accurate to allow this sequencing to be reliably                
	// derived.                                                                                            
	ServiceSequenceNumber                                                                       *float64   `json:"ServiceSequenceNumber,omitempty"`
	// The service provided for the license.  Examples of services may be field related, such as           
	// surveys, or back office services such as address for service, accounting, or legal                  
	// representative.                                                                                     
	ServiceTypeID                                                                               *string    `json:"ServiceTypeID,omitempty"`
	// The time of start day when the service started.                                                     
	StartTime                                                                                   *time.Time `json:"StartTime,omitempty"`
}
