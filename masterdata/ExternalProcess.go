package masterdata

// A workflow configuration in the context of a scheduled job.
type ExternalProcess struct {
	// ID reference of the External Process Type                                                      
	EdsExternalProcessTypeID                                                                   string `json:"EdsExternalProcessTypeID"`
	// Reference name for the security scheme in the ConnectedSourceRegistryEntry document this       
	// external process belongs to.                                                                   
	SecuritySchemeName                                                                         string `json:"SecuritySchemeName"`
	// External Process endpoint                                                                      
	URL                                                                                        string `json:"Url"`
}
