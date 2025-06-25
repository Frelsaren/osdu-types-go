package masterdata

// An array of many addresses displayed as table. Organisation Address table contains
// information on the address, phone numbers, primary contacts, and location of the business
// associate, allowing clients to have multiple addresses. For example, compan ies that have
// a headquarters and various satellite offices.
type Address struct {
	// One or more lines of address for a business associate.                                                                
	AddressLines                                                                                []string                     `json:"AddressLines,omitempty"`
	// The type of business associate address. For example shipping, billing, sales...                                       
	AddressTypeID                                                                               *string                      `json:"AddressTypeID,omitempty"`
	// City as GeoPoliticalContext.                                                                                          
	City                                                                                        *AbstractGeoPoliticalContext `json:"City,omitempty"`
	// Unique identifier for the area that is the Country. Note that if you choose, you could                                
	// use only the City relationship and derive the Country and StateProvince (or other                                     
	// subdivision). Alternatively, you may choose to populate all three relationships                                       
	// explicitly. For example Austria, Canada, United Kingdom, USA, Venezuela.                                              
	Country                                                                                     *AbstractGeoPoliticalContext `json:"Country,omitempty"`
	// A Y/N flag indicating whether this Organisation Address is currently either active /                                  
	// valid (Y) or inactive / invalid (N).                                                                                  
	IsActive                                                                                    *bool                        `json:"IsActive,omitempty"`
	// A flag indicating this contact information is primary or preferred for the given                                      
	// AddressTypeID.                                                                                                        
	IsPreferred                                                                                 *bool                        `json:"IsPreferred,omitempty"`
	// The type of office, such as branch, lab facility, etc.                                                                
	OfficeTypeID                                                                                *string                      `json:"OfficeTypeID,omitempty"`
	// Code number assigned by the postal service identifying a mail delivery zone.                                          
	PostalCode                                                                                  *string                      `json:"PostalCode,omitempty"`
	// The primary contact for this address.                                                                                 
	PrimaryContact                                                                              *AbstractContactUserProfile  `json:"PrimaryContact,omitempty"`
	// Unique identifier for the area that is the Country Note that if you choose, you could use                             
	// only the City relationship and derive the Country and StateProvince (or other                                         
	// subdivision). Alternatively, you may choose to populate all three relationships                                       
	// explicitly. For example states, provinces or other political subdivisions of countries.                               
	ProvinceState                                                                               *AbstractGeoPoliticalContext `json:"ProvinceState,omitempty"`
	// Narrative remarks about this address.                                                                                 
	Remark                                                                                      *string                      `json:"Remark,omitempty"`
}
