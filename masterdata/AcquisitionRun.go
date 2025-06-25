package masterdata

import "time"

// A Run is defined between two consecutive "extractions" from the ground" from acquisition
// string - meaning a single run has necessarily the same probes and gauges installed. A run
// is made of multiple (at least 2) passes.
type AcquisitionRun struct {
	// Identifiers of he conveyance method used to acquire the pressure test data - if not an                                        
	// acquired log leave empty/absent.                                                                                              
	// Mainly a unique ID - but can be multiple in case of stuck pipe,…                                                              
	ConveyanceMethodIDs                                                                         []string                             `json:"ConveyanceMethodIDs,omitempty"`
	// The vertical measurement reference for this well testing acquisition activity. This                                           
	// object defines the vertical reference datum for the measured depths.                                                          
	DepthReferenceSystem                                                                        *AbstractFacilityVerticalMeasurement `json:"DepthReferenceSystem,omitempty"`
	// Type of gauge used for the test                                                                                               
	InstalledGaugeTypeIDs                                                                       []string                             `json:"InstalledGaugeTypeIDs,omitempty"`
	// Type of probe used for the test                                                                                               
	InstalledProbeTypeIDs                                                                       []string                             `json:"InstalledProbeTypeIDs"`
	// Main Category of the Presssure Test - could be Formation Test, Transient Test,                                                
	// Interference Transient Tests,…                                                                                                
	PressureTestCategoryID                                                                      []string                             `json:"PressureTestCategoryID,omitempty"`
	// Array of unitary Acquisition Stations - which is defined as a depth constant, stop, of                                        
	// the acquisition string, within a hole - where one or many tests can be tried out                                              
	PressureTestsAcquisitionStations                                                            []AcquisitionStation                 `json:"PressureTestsAcquisitionStations"`
	// Identifier of the tubular assembly this run went through                                                                      
	RunAssemblyID                                                                               *string                              `json:"RunAssemblyID,omitempty"`
	// Date and Time of the end of this specific run                                                                                 
	RunEndDate                                                                                  *time.Time                           `json:"RunEndDate,omitempty"`
	// Identifier of this specific run within the Job. (Can be a Sequential Number, a GUID,… but                                     
	// must be unique)                                                                                                               
	RunIdentifier                                                                               *int64                               `json:"RunIdentifier,omitempty"`
	// Alphanumeric Name of the Run as captured in the acquisition report                                                            
	RunName                                                                                     *string                              `json:"RunName,omitempty"`
	// Date and Time of the start of this specific run                                                                               
	RunStartDate                                                                                *time.Time                           `json:"RunStartDate,omitempty"`
	// Identifier of the Tool Name (Branded Model Name) in the associated reference data list                                        
	ToolNameID                                                                                  string                               `json:"ToolNameID"`
	// Tool String Description - a long concatenation of the tools used for testing services                                         
	// such as MDT                                                                                                                   
	ToolStringDescription                                                                       *string                              `json:"ToolStringDescription,omitempty"`
	// The type of fluid in the wellbore at time of logging                                                                          
	// e.g. oil based mud, water based mud, water.                                                                                   
	WellboreFluidTypeID                                                                         *string                              `json:"WellboreFluidTypeID,omitempty"`
}
