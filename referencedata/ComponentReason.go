package referencedata

// Describes a distinct combination of Reason and Component IDs for a given Reason Detail
type ComponentReason struct {
	// Downtime components for which this downtime reason is valid for.            
	DowntimeComponentID                                                    *string `json:"DowntimeComponentID,omitempty"`
	// Downtime reasons for which this downtime reason detail is valid for.        
	DowntimeReasonID                                                       *string `json:"DowntimeReasonID,omitempty"`
}
