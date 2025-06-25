// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    wellPressureTestAcquisitionJob, err := UnmarshalWellPressureTestAcquisitionJob(bytes)
//    bytes, err = wellPressureTestAcquisitionJob.Marshal()

package masterdata

import "encoding/json"

func UnmarshalWellPressureTestAcquisitionJob(data []byte) (WellPressureTestAcquisitionJob, error) {
	var r WellPressureTestAcquisitionJob
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WellPressureTestAcquisitionJob) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

import "time"

// A Well Pressure Test Acquisition Job is a subset of a Well Pressure Test Acquisition
// Program - usually assigned to a single contracting company and summarizing the overall
// testing phases (multiple measurement stations, run over multiple passes from one or more
// runs).
// The Acquisition String can change between two different runs - and at each station, only
// a subset of the "in place" string can be activated - which subset can evolve from 1
// measurement to the other.
type WellPressureTestAcquisitionJob struct {
	// The access control tags associated with this entity.                                                                          
	ACL                                                                                          AccessControlList                   `json:"acl"`
	// The links to data, which constitute the inputs, from which this record instance is                                            
	// derived.                                                                                                                      
	Ancestry                                                                                     *ParentList                         `json:"ancestry,omitempty"`
	// Timestamp of the time at which initial version of this OSDU resource object was created.                                      
	// Set by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                         
	CreateTime                                                                                   *time.Time                          `json:"createTime,omitempty"`
	// The user reference, which created the first version of this resource object. Set by the                                       
	// System.                                                                                                                       
	CreateUser                                                                                   *string                             `json:"createUser,omitempty"`
	Data                                                                                         *WellPressureTestAcquisitionJobData `json:"data,omitempty"`
	// Previously called ResourceID or SRN which identifies this OSDU resource object without                                        
	// version.                                                                                                                      
	ID                                                                                           *string                             `json:"id,omitempty"`
	// The schema identification for the OSDU resource object following the pattern                                                  
	// {Namespace}:{Source}:{Type}:{VersionMajor}.{VersionMinor}.{VersionPatch}. The versioning                                      
	// scheme follows the semantic versioning, https://semver.org/.                                                                  
	Kind                                                                                         string                              `json:"kind"`
	// The entity's legal tags and compliance status. The actual contents associated with the                                        
	// legal tags is managed by the Compliance Service.                                                                              
	Legal                                                                                        LegalMetaData                       `json:"legal"`
	// The Frame of Reference meta data section linking the named properties to self-contained                                       
	// definitions.                                                                                                                  
	Meta                                                                                         []FrameOfReferenceMetaDataItem      `json:"meta,omitempty"`
	// Timestamp of the time at which this version of the OSDU resource object was created. Set                                      
	// by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                             
	ModifyTime                                                                                   *time.Time                          `json:"modifyTime,omitempty"`
	// The user reference, which created this version of this resource object. Set by the System.                                    
	ModifyUser                                                                                   *string                             `json:"modifyUser,omitempty"`
	// A generic dictionary of string keys mapping to string value. Only strings are permitted                                       
	// as keys and values.                                                                                                           
	Tags                                                                                         map[string]string                   `json:"tags,omitempty"`
	// The version number of this OSDU resource; set by the framework.                                                               
	Version                                                                                      *int64                              `json:"version,omitempty"`
}
