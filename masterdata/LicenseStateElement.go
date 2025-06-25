package masterdata

// A specific well license status value belonging to a well license status class facet,
// which is assigned by a business
// associate and valid for a time interval. It identifies the facility the license status is
// valid for.
type LicenseStateElement struct {
	// Unique identifier of this element in the parent's list of license states.                        
	ElementIdentifier                                                                           *string `json:"ElementIdentifier,omitempty"`
	// A flag that indicates whether a particular status is currently in effect.  Historical            
	// status information can be retained for legal, administrative or analysis purposes.               
	// Future status information can be projected based on the terms of a license and may also          
	// be captured in this data object.                                                                 
	IsActive                                                                                    *bool   `json:"IsActive,omitempty"`
	// The specific status that is applicable, which belongs to a status class.  They are               
	// selected from values organized in a faceted taxonomy, where the status class is the name         
	// of a facet. Each status class should represent an element in a faceted taxonomy that             
	// reflects one perspective of the state or condition of a data object.                             
	LicenseStatusTypeID                                                                         *string `json:"LicenseStatusTypeID,omitempty"`
	// Identify the business associate who assigned this status.  A status is most commonly             
	// assigned by the operator, field contractor, contact for service or the regulatory agency.        
	StatusAssignedByID                                                                          *string `json:"StatusAssignedByID,omitempty"`
	// The date on which this status was first valid.  Note that a data object can have more            
	// than one status in effect at the same time, provided that there is no conflict or                
	// confusion created by doing so.                                                                   
	StatusEffectiveDate                                                                         *string `json:"StatusEffectiveDate,omitempty"`
	// The reason why this status has been assigned to a well license.  This is particularly            
	// important if a license has been cancelled, suspended or terminated by a regulatory               
	// authority.                                                                                       
	StatusReason                                                                                *string `json:"StatusReason,omitempty"`
	// The date on which this status is no longer valid.  A status can become invalid based on          
	// events or activities, because they have been superseded by a different status or even            
	// because status information is no longer relevant.                                                
	StatusTerminationDate                                                                       *string `json:"StatusTerminationDate,omitempty"`
	// Identify the wellbore for which this status is valid, particularly useful when a status          
	// is not applicable for the entire well configuration.  In the license, each wellbore may          
	// have a different status in terms of the license. Note that these are not well statuses,          
	// but license statuses.                                                                            
	WellboreID                                                                                  *string `json:"WellboreID,omitempty"`
	// Identify the well for which this status is valid, particularly useful when a status is           
	// applicable for the entire well configuration . Note that these are not well statuses, but        
	// license statuses.                                                                                
	WellID                                                                                      *string `json:"WellID,omitempty"`
}
