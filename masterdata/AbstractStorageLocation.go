package masterdata

// A record about the storage location of an item, e.g. a rock or fluid sample, seismic
// tape, where (facility), by whom (organisation), when (dates) and how (description).
//
// The location where the sample was stored at the end of the chain of custody event
//
// The initial physical location where this sample was stored at the start of the chain of
// custody event.
type AbstractStorageLocation struct {
	// The date the item arrived at the storage location.                                              
	EffectiveDateTime                                                                          *string `json:"EffectiveDateTime,omitempty"`
	// The item (sample, tape)  identifier, for example a barcode, which identifies the item in        
	// the StorageFacility.                                                                            
	SampleIdentifier                                                                           *string `json:"SampleIdentifier,omitempty"`
	// Identifies the warehouse in which the item is stored.                                           
	StorageFacilityID                                                                          *string `json:"StorageFacilityID,omitempty"`
	// The name of the location where the item is stored. It can be stored in more than one            
	// location over time.                                                                             
	StorageLocationDescription                                                                 *string `json:"StorageLocationDescription,omitempty"`
	// Identifies the organisation with which the item is stored.                                      
	StorageOrganisationID                                                                      *string `json:"StorageOrganisationID,omitempty"`
	// The date and time at which the item is  no longer stored in the given location.  If the         
	// item is still in this storage, the 'TerminationDateTime' is left absent.                        
	TerminationDateTime                                                                        *string `json:"TerminationDateTime,omitempty"`
}
