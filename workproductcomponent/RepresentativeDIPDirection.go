package workproductcomponent

// Indicates the representative azimuth value of the fault plane
type RepresentativeDIPDirection struct {
	// The kind of north which is used as a reference for the azimuth value         
	NorthKind                                                              *string  `json:"NorthKind,omitempty"`
	// The azimuth value                                                            
	Value                                                                  *float64 `json:"Value,omitempty"`
}
