package referencedata

// A tuple FacetType, FacetRole, both calling specific references
//
// FacetType: Enumeration of the type of additional context about the nature of a property
// type (it may include conditions, direction, qualifiers, or statistics).
//
// FacetRole: Additional context about the nature of a property type. The purpose of such
// attribute is to minimize the need to create specialized property types by mutualizing
// some well known qualifiers such as "maximum", "minimum" which apply to a lot of different
// property types.
type AbstractFacet struct {
	// Additional context about the nature of a property type. The purpose of such attribute is        
	// to minimize the need to create specialized property types by mutualizing some well known        
	// qualifiers such as "maximum", "minimum" which apply to a lot of different property types.       
	FacetRoleID                                                                                 string `json:"FacetRoleID"`
	// FacetType: An 'enumeration' of the type of additional context about the nature of a             
	// property type (it may include conditions, direction, qualifiers, or statistics).                
	FacetTypeID                                                                                 string `json:"FacetTypeID"`
}
