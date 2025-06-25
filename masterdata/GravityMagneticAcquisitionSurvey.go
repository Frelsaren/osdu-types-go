// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gravityMagneticAcquisitionSurvey, err := UnmarshalGravityMagneticAcquisitionSurvey(bytes)
//    bytes, err = gravityMagneticAcquisitionSurvey.Marshal()

package masterdata

import "encoding/json"

func UnmarshalGravityMagneticAcquisitionSurvey(data []byte) (GravityMagneticAcquisitionSurvey, error) {
	var r GravityMagneticAcquisitionSurvey
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GravityMagneticAcquisitionSurvey) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

import "time"

// A gravity and magnetics acquisition survey is a type of a geophysical survey acquisition
// and a business project that deploys resources to the field to record (passive) gravity
// and or magnetic geophysical measurement data.
//
// Several types of gravity and magnetic data can be acquired simultaneously from the same
// survey, and different acquisition platforms can be used to acquire the data, for example,
// land, air, marine, satellite and borehole.
//
// It may be referred to as a field survey, acquisition survey, or field program.
type GravityMagneticAcquisitionSurvey struct {
	// The access control tags associated with this entity.                                                                            
	ACL                                                                                          AccessControlList                     `json:"acl"`
	// The links to data, which constitute the inputs, from which this record instance is                                              
	// derived.                                                                                                                        
	Ancestry                                                                                     *ParentList                           `json:"ancestry,omitempty"`
	// Timestamp of the time at which initial version of this OSDU resource object was created.                                        
	// Set by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                           
	CreateTime                                                                                   *time.Time                            `json:"createTime,omitempty"`
	// The user reference, which created the first version of this resource object. Set by the                                         
	// System.                                                                                                                         
	CreateUser                                                                                   *string                               `json:"createUser,omitempty"`
	Data                                                                                         *GravityMagneticAcquisitionSurveyData `json:"data,omitempty"`
	// Previously called ResourceID or SRN which identifies this OSDU resource object without                                          
	// version.                                                                                                                        
	ID                                                                                           *string                               `json:"id,omitempty"`
	// The schema identification for the OSDU resource object following the pattern                                                    
	// {Namespace}:{Source}:{Type}:{VersionMajor}.{VersionMinor}.{VersionPatch}. The versioning                                        
	// scheme follows the semantic versioning, https://semver.org/.                                                                    
	Kind                                                                                         string                                `json:"kind"`
	// The entity's legal tags and compliance status. The actual contents associated with the                                          
	// legal tags is managed by the Compliance Service.                                                                                
	Legal                                                                                        LegalMetaData                         `json:"legal"`
	// The Frame of Reference meta data section linking the named properties to self-contained                                         
	// definitions.                                                                                                                    
	Meta                                                                                         []FrameOfReferenceMetaDataItem        `json:"meta,omitempty"`
	// Timestamp of the time at which this version of the OSDU resource object was created. Set                                        
	// by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                               
	ModifyTime                                                                                   *time.Time                            `json:"modifyTime,omitempty"`
	// The user reference, which created this version of this resource object. Set by the System.                                      
	ModifyUser                                                                                   *string                               `json:"modifyUser,omitempty"`
	// A generic dictionary of string keys mapping to string value. Only strings are permitted                                         
	// as keys and values.                                                                                                             
	Tags                                                                                         map[string]string                     `json:"tags,omitempty"`
	// The version number of this OSDU resource; set by the framework.                                                                 
	Version                                                                                      *int64                                `json:"version,omitempty"`
}
