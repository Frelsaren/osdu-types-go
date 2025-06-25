// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    unitQuantity, err := UnmarshalUnitQuantity(bytes)
//    bytes, err = unitQuantity.Marshal()

package referencedata

import "encoding/json"

func UnmarshalUnitQuantity(data []byte) (UnitQuantity, error) {
	var r UnitQuantity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UnitQuantity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

import "time"

// A single unit quantity class representing a superclass of mutually comparable quantity
// kinds. The class may or may not represent a recognized quantity kind (i.e., it may
// logically be abstract).  Quantities of the same class will have the same quantity
// dimension.  However, quantities of the same dimension are not necessarily of the same
// class. The term "unit quantity class" is intended to have the same general  meaning as
// the term "kind of quantity" as defined by the  "International vocabulary of meteorology -
// Basic and general concepts and associated terms (VIM)" (JCGM 200:2008)
// https://www.bipm.org/utils/common/documents/jcgm/JCGM_200_2012.pdf.
type UnitQuantity struct {
	// The access control tags associated with this entity.                                                                     
	ACL                                                                                          AccessControlList              `json:"acl"`
	// The links to data, which constitute the inputs, from which this record instance is                                       
	// derived.                                                                                                                 
	Ancestry                                                                                     *ParentList                    `json:"ancestry,omitempty"`
	// Timestamp of the time at which initial version of this OSDU resource object was created.                                 
	// Set by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                    
	CreateTime                                                                                   *time.Time                     `json:"createTime,omitempty"`
	// The user reference, which created the first version of this resource object. Set by the                                  
	// System.                                                                                                                  
	CreateUser                                                                                   *string                        `json:"createUser,omitempty"`
	Data                                                                                         *UnitQuantityData              `json:"data,omitempty"`
	// Previously called ResourceID or SRN which identifies this OSDU resource object without                                   
	// version.                                                                                                                 
	ID                                                                                           *string                        `json:"id,omitempty"`
	// The schema identification for the OSDU resource object following the pattern                                             
	// {Namespace}:{Source}:{Type}:{VersionMajor}.{VersionMinor}.{VersionPatch}. The versioning                                 
	// scheme follows the semantic versioning, https://semver.org/.                                                             
	Kind                                                                                         string                         `json:"kind"`
	// The entity's legal tags and compliance status. The actual contents associated with the                                   
	// legal tags is managed by the Compliance Service.                                                                         
	Legal                                                                                        LegalMetaData                  `json:"legal"`
	// The Frame of Reference meta data section linking the named properties to self-contained                                  
	// definitions.                                                                                                             
	Meta                                                                                         []FrameOfReferenceMetaDataItem `json:"meta,omitempty"`
	// Timestamp of the time at which this version of the OSDU resource object was created. Set                                 
	// by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                        
	ModifyTime                                                                                   *time.Time                     `json:"modifyTime,omitempty"`
	// The user reference, which created this version of this resource object. Set by the System.                               
	ModifyUser                                                                                   *string                        `json:"modifyUser,omitempty"`
	// A generic dictionary of string keys mapping to string value. Only strings are permitted                                  
	// as keys and values.                                                                                                      
	Tags                                                                                         map[string]string              `json:"tags,omitempty"`
	// The version number of this OSDU resource; set by the framework.                                                          
	Version                                                                                      *int64                         `json:"version,omitempty"`
}
