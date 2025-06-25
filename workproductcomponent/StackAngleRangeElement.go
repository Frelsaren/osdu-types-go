package workproductcomponent

// A range container carrying minimum/maximum angle values.
type StackAngleRangeElement struct {
	// The maximum angle value.         
	Maximum                    *float64 `json:"Maximum,omitempty"`
	// The minimum angle value.         
	Minimum                    *float64 `json:"Minimum,omitempty"`
}
