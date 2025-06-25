package workproduct

// A meta data item, which allows the association of named properties or property values to
// a Unit/Measurement/CRS/Azimuth/Time context.
type FrameOfReferenceMetaDataItem struct {
	// The kind of reference, 'Unit' for FrameOfReferenceUOM.                                                 
	//                                                                                                        
	// The kind of reference, constant 'CRS' for FrameOfReferenceCRS.                                         
	//                                                                                                        
	// The kind of reference, constant 'DateTime', for FrameOfReferenceDateTime.                              
	//                                                                                                        
	// The kind of reference, constant 'AzimuthReference', for FrameOfReferenceAzimuthReference.              
	Kind                                                                                        ReferenceKind `json:"kind"`
	// The unit symbol or name of the unit.                                                                   
	//                                                                                                        
	// The name of the CRS.                                                                                   
	//                                                                                                        
	// The name of the DateTime format and reference.                                                         
	//                                                                                                        
	// The name of the CRS or the symbol/name of the unit.                                                    
	Name                                                                                        *string       `json:"name,omitempty"`
	// The self-contained, persistable reference string uniquely identifying the Unit.                        
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying the CRS.                         
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying DateTime                         
	// reference.                                                                                             
	//                                                                                                        
	// The self-contained, persistable reference string uniquely identifying AzimuthReference.                
	PersistableReference                                                                        string        `json:"persistableReference"`
	// The list of property names, to which this meta data item provides Unit context to. A full              
	// path like "StructureA.PropertyB" is required to define a unique context; "data" is                     
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides CRS context to. A full               
	// path like "StructureA.PropertyB" is required to define a unique context; "data" is                     
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides DateTime context to. A               
	// full path like "StructureA.PropertyB" is required to define a unique context; "data" is                
	// omitted since frame-of reference normalization only applies to the data block.                         
	//                                                                                                        
	// The list of property names, to which this meta data item provides AzimuthReference                     
	// context to. A full path like "StructureA.PropertyB" is required to define a unique                     
	// context; "data" is omitted since frame-of reference normalization only applies to the                  
	// data block.                                                                                            
	PropertyNames                                                                               []string      `json:"propertyNames,omitempty"`
	// SRN to unit of measure reference.                                                                      
	UnitOfMeasureID                                                                             *string       `json:"unitOfMeasureID,omitempty"`
	// SRN to CRS reference.                                                                                  
	CoordinateReferenceSystemID                                                                 *string       `json:"coordinateReferenceSystemID,omitempty"`
}
