package workproductcomponent

// The name and version of the software being executed in the context of this activity
type PurpleSoftware struct {
	// The name of the software, application or plug-in used while performing this activity.           
	SoftwareName                                                                               *string `json:"SoftwareName,omitempty"`
	// The version of the software, application or plug-in used while performing this activity.        
	Version                                                                                    *string `json:"Version,omitempty"`
}
