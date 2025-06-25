package workproductcomponent

// A range container carrying minimum/maximum angle values.
type LengthRange struct {
	// The maximum length value.         
	Maximum                     *float64 `json:"Maximum,omitempty"`
	// The minimum length value.         
	Minimum                     *float64 `json:"Minimum,omitempty"`
}
