package referencedata

// Extent or area of use information.
type PurpleExtent struct {
	// The Extent authority code, corresponding to the ISO19111 ID and 'projjson' id.                           
	AuthorityCode                                                                    *PurpleExtentAuthorityCode `json:"AuthorityCode,omitempty"`
	// Eastern longitude limit of the bounding box in degrees based on WGS 84                                   
	BoundingBoxEastBoundLongitude                                                    *float64                   `json:"BoundingBoxEastBoundLongitude,omitempty"`
	// Northern latitude limit of the bounding box in degrees based on WGS 84                                   
	BoundingBoxNorthBoundLatitude                                                    *float64                   `json:"BoundingBoxNorthBoundLatitude,omitempty"`
	// Southern latitude limit of the bounding box in degrees based on WGS 84                                   
	BoundingBoxSouthBoundLatitude                                                    *float64                   `json:"BoundingBoxSouthBoundLatitude,omitempty"`
	// Western longitude limit of the bounding box in degrees based on WGS 84                                   
	BoundingBoxWestBoundLongitude                                                    *float64                   `json:"BoundingBoxWestBoundLongitude,omitempty"`
	// The description of the Extent.                                                                           
	Description                                                                      *string                    `json:"Description,omitempty"`
	// The name of the Extent.                                                                                  
	Name                                                                             *string                    `json:"Name,omitempty"`
}
