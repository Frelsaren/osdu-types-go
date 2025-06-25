package workproduct

// The access control tags associated with this entity.
//
// The access control tags associated with this entity. This structure is included by the
// SystemProperties "acl", which is part of all OSDU records. Not extensible.
type AccessControlList struct {
	// The list of owners of this data record formatted as an email                                    
	// (core.common.model.storage.validation.ValidationDoc.EMAIL_REGEX).                               
	Owners                                                                                    []string `json:"owners"`
	// The list of viewers to which this data record is accessible/visible/discoverable                
	// formatted as an email (core.common.model.storage.validation.ValidationDoc.EMAIL_REGEX).         
	Viewers                                                                                   []string `json:"viewers"`
}
