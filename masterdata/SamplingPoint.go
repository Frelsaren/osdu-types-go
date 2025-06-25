package masterdata

// A free-form reference to the flow port on the Facility where this sample was acquired.
// The property is only used in conjunction with WellheadSampleAcquisition,
// SeparatorSampleAcquisition, FacilitySampleAcquisition.
// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped
// as seen below:
// SamplingPoint= [
// PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition |
// FacilitySampleAcquisition | SeparatorSampleAcquisition>[].Item.SamplingPoint
type SamplingPoint struct {
	// This is a description of the name of the  component or equipment used in capturing the           
	// sample. It can be used in storing the P & ID of the equipment or its component as defined        
	// within the Organisations Facility SOR / Repository.                                              
	SamplingPointName                                                                           *string `json:"SamplingPointName,omitempty"`
	// This is the OSDU record ID from a reference list containing the different types of               
	// sampling points.                                                                                 
	SamplingPointTypeID                                                                         *string `json:"SamplingPointTypeID,omitempty"`
}
