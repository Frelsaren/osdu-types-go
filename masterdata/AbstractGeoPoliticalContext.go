package masterdata

// A single, typed geo-political entity reference, which is 'abstracted' to
// AbstractGeoContext and then aggregated by GeoContexts properties.
//
// City as GeoPoliticalContext.
//
// Unique identifier for the area that is the Country. Note that if you choose, you could
// use only the City relationship and derive the Country and StateProvince (or other
// subdivision). Alternatively, you may choose to populate all three relationships
// explicitly. For example Austria, Canada, United Kingdom, USA, Venezuela.
//
// Unique identifier for the area that is the Country Note that if you choose, you could use
// only the City relationship and derive the Country and StateProvince (or other
// subdivision). Alternatively, you may choose to populate all three relationships
// explicitly. For example states, provinces or other political subdivisions of countries.
type AbstractGeoPoliticalContext struct {
	// Reference to GeoPoliticalEntity.                                                                 
	GeoPoliticalEntityID                                                                        *string `json:"GeoPoliticalEntityID,omitempty"`
	// The GeoPoliticalEntityType reference of the GeoPoliticalEntity (via GeoPoliticalEntityID)        
	// for application convenience.                                                                     
	GeoTypeID                                                                                   *string `json:"GeoTypeID,omitempty"`
}
