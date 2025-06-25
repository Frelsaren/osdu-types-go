package workproductcomponent

// An array of Artefacts - each artefact has a Role, Resource tuple. An artefact is distinct
// from the file, in the sense certain valuable information is generated during loading
// process (Artefact generation process). Examples include retrieving location data,
// performing an OCR which may result in the generation of artefacts which need to be
// preserved distinctly
type AbstractGridRepresentationArtefact struct {
	// The SRN which identifies this OSDU Artefact resource.                             
	ResourceID                                                                   *string `json:"ResourceID,omitempty"`
	// The kind or schema ID of the artefact. Resolvable with the Schema Service.        
	ResourceKind                                                                 *string `json:"ResourceKind,omitempty"`
	// The record id of this artefact's role.                                            
	RoleID                                                                       *string `json:"RoleID,omitempty"`
}
