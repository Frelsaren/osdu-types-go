package workproduct

// The entity's legal tags and compliance status. The actual contents associated with the
// legal tags is managed by the Compliance Service.
//
// Legal meta data like legal tags, relevant other countries, legal status. This structure
// is included by the SystemProperties "legal", which is part of all OSDU records. Not
// extensible.
type LegalMetaData struct {
	// The list of legal tags, which resolve to legal properties (like country of origin, export         
	// classification code, etc.) and rules with the help of the Compliance Service.                     
	Legaltags                                                                                   []string `json:"legaltags"`
	// The list of other relevant data countries as an array of two-letter country codes, see            
	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.                                                 
	OtherRelevantDataCountries                                                                  []string `json:"otherRelevantDataCountries"`
	// The legal status. Set by the system after evaluation against the compliance rules                 
	// associated with the "legaltags" using the Compliance Service.                                     
	Status                                                                                      *string  `json:"status,omitempty"`
}
