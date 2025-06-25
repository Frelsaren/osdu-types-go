package masterdata

import "time"

// This captures the acquisition parameters obtained during  the sample acquisition event
// associated with this sample. Note that this attribute should only be used when
// associating the sample with an acquisition event from its original source and not for
// sub-sampling or derivative sources.
type SampleAcquisition struct {
	// The end date and time of the acquisition event.                                                                   
	AcquisitionEndDate                                                                          *time.Time               `json:"AcquisitionEndDate,omitempty"`
	// The start date and time of the acquisition event.                                                                 
	AcquisitionStartDate                                                                        *time.Time               `json:"AcquisitionStartDate,omitempty"`
	// The company that managed the collection of the sample from its original / source                                  
	// environment.                                                                                                      
	CollectionServiceCompanyID                                                                  *string                  `json:"CollectionServiceCompanyID,omitempty"`
	// The company that handled the sample on site after the it was collected/extracted. For                             
	// example, the company that applied field preservation methods, and prepared the sample for                         
	// shipping.                                                                                                         
	HandlingServiceCompanyID                                                                    *string                  `json:"HandlingServiceCompanyID,omitempty"`
	// A remark object, pairing a remark text with a source, e.g. an author, and a date, which                           
	// is typically included in an array.                                                                                
	Remarks                                                                                     []AbstractRemark         `json:"Remarks,omitempty"`
	// This is the OSDU Record ID for the original sample container used in the acquired sample.                         
	// This is only for the initial sample acquisition. Only populate this for the original                              
	// sample.                                                                                                           
	SampleAcquisitionContainerID                                                                *string                  `json:"SampleAcquisitionContainerID,omitempty"`
	// This attribute provides information about the acquisition parameters and process used in                          
	// acquiring the target sample. Other information about the sample itself can be found in                            
	// the Sample object.                                                                                                
	SampleAcquisitionDetail                                                                     *SampleAcquisitionDetail `json:"SampleAcquisitionDetail,omitempty"`
	// A reference to the parent record, which can group this sample acquisition event record                            
	// with other sample acquisition event records, collected as part of the same job.                                   
	SampleAcquisitionJobID                                                                      *string                  `json:"SampleAcquisitionJobID,omitempty"`
	// This is a reference list of the different types of sample acquisition events used in                              
	// acquiring samples. i.e. downhole sample acquisition type, outcrop, coring, non-facility                           
	// site, etc                                                                                                         
	SampleAcquisitionTypeID                                                                     *string                  `json:"SampleAcquisitionTypeID,omitempty"`
}
