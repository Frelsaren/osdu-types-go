package workproductcomponent

// A single, typed downtime event reference, which is 'abstracted' to
// AbstractProductionValueContext and then aggregated by ValueContexts properties.
type AbstractDowntimeContext struct {
	// The fixed type 'Downtime Event' for this AbstractDowntimeContext            
	ContextType                                                        ContextType `json:"ContextType"`
	// Reference to DowntimeEvent.                                                 
	DowntimeEventID                                                    string      `json:"DowntimeEventID"`
}
