package masterdata

import "time"

// This provides information pertaining to the certification process conducted on a fluid
// sample container object.
type ContainerCertification struct {
	// This is the date of the last inspection performed on the sample container.                           
	// Reference:                                                                                           
	// InspectionDate=PRODML:2.1:FluidSampleContainer.Model.LastInspectionDate                              
	InspectionDate                                                                               *time.Time `json:"InspectionDate,omitempty"`
	// Indicator to determine if the sample container can be transported.                                   
	IsTransportable                                                                              *bool      `json:"IsTransportable,omitempty"`
	// This is the proposed date for the next inspection to be performed on the sample container.           
	NextInspectionDate                                                                           *time.Time `json:"NextInspectionDate,omitempty"`
	// This is the OSDU object identifier for the file or document containing detailed                      
	// information on the certification process performed ascertaining its viability for safe               
	// transportation of the sample.                                                                        
	TransportCertificateDocumentID                                                               *string    `json:"TransportCertificateDocumentID,omitempty"`
	// This is the period or amount of time for which the last inspection or certification                  
	// process is valid.                                                                                    
	ValidityPeriod                                                                               *float64   `json:"ValidityPeriod,omitempty"`
}
