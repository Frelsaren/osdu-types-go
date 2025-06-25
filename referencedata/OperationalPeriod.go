package referencedata

// Defines the start and end year of the satellite mission
type OperationalPeriod struct {
	// The year that the satellite mission ceased or ended. Leave blank if the mission is still       
	// active.                                                                                        
	EndYear                                                                                    *int64 `json:"EndYear,omitempty"`
	// The year that the satellite mission became operational or started                              
	StartYear                                                                                  *int64 `json:"StartYear,omitempty"`
}
