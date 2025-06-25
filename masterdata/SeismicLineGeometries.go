package masterdata

// A processing geometry comprising the 2D Interpretation Survey (often referred to as
// survey in the context of an interpretation application but not a survey in the context of
// acquisition).
type SeismicLineGeometries struct {
	// Reference to a 2D processing geometry associated with a particular seismic line used in         
	// interpretation.  Multiple datasets may refer to this geometry and support the                   
	// interpretation.                                                                                 
	SeismicLineGeometryID                                                                      *string `json:"SeismicLineGeometryID,omitempty"`
	// The distinct line name used by interpretation objects (horizons) in the Interpretation          
	// Project, which may differ from the name used in seismic line geometry.  This allows the         
	// line names in the project to be unique within the project even though the names may not         
	// be unique across all the projects that use the same line geometries.  The name used in a        
	// horizon pick is related to the appropriate geometry through this name.                          
	SeismicLineName                                                                            *string `json:"SeismicLineName,omitempty"`
}
