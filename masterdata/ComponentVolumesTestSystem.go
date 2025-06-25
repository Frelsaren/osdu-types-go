package masterdata

// Description of the tested component during a singe step.
type ComponentVolumesTestSystem struct {
	// e.g. surface lines, choke/kill Lines, BOP, Casing        
	SystemComponentName                                 *string `json:"SystemComponentName,omitempty"`
	// The volume of the named system                           
	SystemComponentVolume                               float64 `json:"SystemComponentVolume"`
}
