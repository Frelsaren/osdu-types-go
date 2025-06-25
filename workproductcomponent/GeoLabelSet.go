// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    geoLabelSet, err := UnmarshalGeoLabelSet(bytes)
//    bytes, err = geoLabelSet.Marshal()

package workproductcomponent

import "encoding/json"

func UnmarshalGeoLabelSet(data []byte) (GeoLabelSet, error) {
	var r GeoLabelSet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GeoLabelSet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

import "time"

// The GeoLabelSet acts as an object for collecting consolidated geological properties or
// information related to a labelled master entity, which could be a Reservoir, Reservoir
// Segment, Field, Prospect, etc. Individual labels are name-value pairs categorised by a
// GeoLabelType which allows for specifying characteristics such as: the average porosity of
// a Reservoir as 20%, the average initial reservoir pressure for Reservoir Segment as at
// 8000 psi, or identifying the predominant depositional environment of a Field as “marine
// shallow”. The content may have been identified through various methods, including
// measurement, interpretation, estimation, data collection from reports, delivered data
// packages, data mining etc.
type GeoLabelSet struct {
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
	Data                                                                                         *GeoLabelSetData               `json:"data,omitempty"`
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
