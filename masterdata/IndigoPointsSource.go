package masterdata

// The definition of the points used as a source for each individual critical point,
// resulting from parameters set or observed during a BHA run
type IndigoPointsSource struct {
	// Description associated with the source indicator type.                  
	Description                                                       *string  `json:"Description,omitempty"`
	// The name of the Source Indicator Type. For example Manual Input         
	Name                                                              *string  `json:"Name,omitempty"`
	// The value at that point.                                                
	Value                                                             *float64 `json:"Value,omitempty"`
	// Unit of measure of the Value value                                      
	ValueUnitID                                                       *string  `json:"ValueUnitID,omitempty"`
}
