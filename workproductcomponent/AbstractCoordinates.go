package workproductcomponent

// A geographic position on the surface of the earth.
type AbstractCoordinates struct {
	// x is Easting or Longitude.         
	X                            *float64 `json:"X,omitempty"`
	// y is Northing or Latitude.         
	Y                            *float64 `json:"Y,omitempty"`
}
