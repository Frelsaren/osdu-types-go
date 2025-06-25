package workproductcomponent

// A structure carrying descriptions of fault throw characteristics.
type FaultThrowDescriptions struct {
	// Specifies the fault throw type (reverse, normal etc…)                                             
	FaultThrowTypeID                                                                            *string  `json:"FaultThrowTypeID,omitempty"`
	// Time interval Maximum age the fault throw was active - if the link to horizon is defined,         
	// we can derive the link to horizons through geologic age                                           
	MaximumAge                                                                                  *float64 `json:"MaximumAge,omitempty"`
	// Time interval Minimum age the fault throw was active - if the link to horizon is defined,         
	// we can derive the link to horizons through geologic age                                           
	MinimumAge                                                                                  *float64 `json:"MinimumAge,omitempty"`
}
