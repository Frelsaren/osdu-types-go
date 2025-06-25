// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    seismic3DInterpretationSet, err := UnmarshalSeismic3DInterpretationSet(bytes)
//    bytes, err = seismic3DInterpretationSet.Marshal()

package masterdata

import "encoding/json"

func UnmarshalSeismic3DInterpretationSet(data []byte) (Seismic3DInterpretationSet, error) {
	var r Seismic3DInterpretationSet
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Seismic3DInterpretationSet) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

import "time"

// A seismic 3D interpretation survey is a collection of trace datasets with a common bin
// grid that are important for interpretation.  It does not comprise all of the datasets
// with a common bin grid, nor all of the datasets from a processing project, nor all of the
// datasets from an acquisition project, because some are not suitable for interpretation.
// An interpretation survey may include datasets from more than one acquisition or
// processing project.  Consequently, it is not an acquisition survey nor a processing
// survey.  It is not an application project, which is a collection of all the various
// objects an application and user care about for some analysis (seismic, wells, etc.).  It
// inherits properties shared by project entities because it can serve to capture the
// archiving of a master or authorized project activity. Interpretation objects (horizons)
// are hung from an interpretation survey to give context and to derive spatial location
// based on the common bin grid.  Trace datasets and seismic horizons are associated through
// LineageAssertion, although a master collection of trace datasets and horizons are
// explicitly related through a child relationship property.
type Seismic3DInterpretationSet struct {
	// The access control tags associated with this entity.                                                                      
	ACL                                                                                          AccessControlList               `json:"acl"`
	// The links to data, which constitute the inputs, from which this record instance is                                        
	// derived.                                                                                                                  
	Ancestry                                                                                     *ParentList                     `json:"ancestry,omitempty"`
	// Timestamp of the time at which initial version of this OSDU resource object was created.                                  
	// Set by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                     
	CreateTime                                                                                   *time.Time                      `json:"createTime,omitempty"`
	// The user reference, which created the first version of this resource object. Set by the                                   
	// System.                                                                                                                   
	CreateUser                                                                                   *string                         `json:"createUser,omitempty"`
	Data                                                                                         *Seismic3DInterpretationSetData `json:"data,omitempty"`
	// Previously called ResourceID or SRN which identifies this OSDU resource object without                                    
	// version.                                                                                                                  
	ID                                                                                           *string                         `json:"id,omitempty"`
	// The schema identification for the OSDU resource object following the pattern                                              
	// {Namespace}:{Source}:{Type}:{VersionMajor}.{VersionMinor}.{VersionPatch}. The versioning                                  
	// scheme follows the semantic versioning, https://semver.org/.                                                              
	Kind                                                                                         string                          `json:"kind"`
	// The entity's legal tags and compliance status. The actual contents associated with the                                    
	// legal tags is managed by the Compliance Service.                                                                          
	Legal                                                                                        LegalMetaData                   `json:"legal"`
	// The Frame of Reference meta data section linking the named properties to self-contained                                   
	// definitions.                                                                                                              
	Meta                                                                                         []FrameOfReferenceMetaDataItem  `json:"meta,omitempty"`
	// Timestamp of the time at which this version of the OSDU resource object was created. Set                                  
	// by the System. The value is a combined date-time string in ISO-8601 given in UTC.                                         
	ModifyTime                                                                                   *time.Time                      `json:"modifyTime,omitempty"`
	// The user reference, which created this version of this resource object. Set by the System.                                
	ModifyUser                                                                                   *string                         `json:"modifyUser,omitempty"`
	// A generic dictionary of string keys mapping to string value. Only strings are permitted                                   
	// as keys and values.                                                                                                       
	Tags                                                                                         map[string]string               `json:"tags,omitempty"`
	// The version number of this OSDU resource; set by the framework.                                                           
	Version                                                                                      *int64                          `json:"version,omitempty"`
}
