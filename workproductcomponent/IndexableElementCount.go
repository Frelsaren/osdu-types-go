package workproductcomponent

// Defines the count of a particular indexable element in a representation
type IndexableElementCount struct {
	// The count of indexable element                
	Count                                    int64   `json:"Count"`
	// The indexable element which is counted        
	IndexableElementID                       *string `json:"IndexableElementID,omitempty"`
}
