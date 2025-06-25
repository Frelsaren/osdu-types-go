package workproductcomponent

// A geographic context to an entity. It can be either a reference to a GeoPoliticalEntity,
// Basin, Field, Play or Prospect.
//
// A single, typed geo-political entity reference, which is 'abstracted' to
// AbstractGeoContext and then aggregated by GeoContexts properties.
//
// A single, typed basin entity reference, which is 'abstracted' to AbstractGeoContext and
// then aggregated by GeoContexts properties.
//
// A single, typed field entity reference, which is 'abstracted' to AbstractGeoContext and
// then aggregated by GeoContexts properties.
//
// A single, typed Play entity reference, which is 'abstracted' to AbstractGeoContext and
// then aggregated by GeoContexts properties.
//
// A single, typed Prospect entity reference, which is 'abstracted' to AbstractGeoContext
// and then aggregated by GeoContexts properties.
type AbstractGeoContext struct {
	// Reference to GeoPoliticalEntity.                                                                 
	GeoPoliticalEntityID                                                                        *string `json:"GeoPoliticalEntityID,omitempty"`
	// The GeoPoliticalEntityType reference of the GeoPoliticalEntity (via GeoPoliticalEntityID)        
	// for application convenience.                                                                     
	//                                                                                                  
	// The BasinType reference of the Basin (via BasinID) for application convenience.                  
	//                                                                                                  
	// The fixed type 'Field' for this AbstractGeoFieldContext.                                         
	//                                                                                                  
	// The PlayType reference of the Play (via PlayID) for application convenience.                     
	//                                                                                                  
	// The ProspectType reference of the Prospect (via ProspectID) for application convenience.         
	GeoTypeID                                                                                   *string `json:"GeoTypeID,omitempty"`
	// Reference to Basin.                                                                              
	BasinID                                                                                     *string `json:"BasinID,omitempty"`
	// Reference to Field.                                                                              
	FieldID                                                                                     *string `json:"FieldID,omitempty"`
	// Reference to the play.                                                                           
	PlayID                                                                                      *string `json:"PlayID,omitempty"`
	// Reference to the prospect.                                                                       
	ProspectID                                                                                  *string `json:"ProspectID,omitempty"`
}
