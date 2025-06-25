package masterdata

import "time"

// A list of alternative names for an object.  The preferred name is in a separate, scalar
// property.  It may or may not be repeated in the alias list, though a best practice is to
// include it if the list is present, but to omit the list if there are no other names.
// Note that the abstract entity is an array so the $ref to it is a simple property
// reference.
type AbstractAliasNames struct {
	// Alternative Name value of defined name type for an object.                                       
	AliasName                                                                                *string    `json:"AliasName,omitempty"`
	// A classification of alias names such as by role played or type of source, such as                
	// regulatory name, regulatory code, company code, international standard name, etc.                
	AliasNameTypeID                                                                          *string    `json:"AliasNameTypeID,omitempty"`
	// The StandardsOrganisation (reference-data) or Organisation (master-data) that provided           
	// the name (the source).                                                                           
	DefinitionOrganisationID                                                                 *string    `json:"DefinitionOrganisationID,omitempty"`
	// The date and time when an alias name becomes effective.                                          
	EffectiveDateTime                                                                        *time.Time `json:"EffectiveDateTime,omitempty"`
	// The data and time when an alias name is no longer in effect.                                     
	TerminationDateTime                                                                      *time.Time `json:"TerminationDateTime,omitempty"`
}
