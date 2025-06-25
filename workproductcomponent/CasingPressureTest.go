package workproductcomponent

import "time"

// Casing Pressure Test
type CasingPressureTest struct {
	// Bled Back Volume                                               
	BledBackVolume                                         *float64   `json:"BledBackVolume,omitempty"`
	// Test Date/time                                                 
	CasingTestDate                                         *time.Time `json:"CasingTestDate,omitempty"`
	// Test Duration                                                  
	CasingTestDuration                                     *float64   `json:"CasingTestDuration,omitempty"`
	// Casing Test Elapsed Time following end of cement job           
	CasingTestElapsedTime                                  *float64   `json:"CasingTestElapsedTime,omitempty"`
	// Test Pressure                                                  
	CasingTestPressure                                     *float64   `json:"CasingTestPressure,omitempty"`
	// Casing Test Stable Flow Rate                                   
	CasingTestStableRate                                   *float64   `json:"CasingTestStableRate,omitempty"`
	// Float Depth                                                    
	FloatMeasuredDepth                                     *float64   `json:"FloatMeasuredDepth,omitempty"`
	// Fluid Density                                                  
	FluidDensity                                           *float64   `json:"FluidDensity,omitempty"`
	// Maintained Pressure % of Initial Test Pressure                 
	MaintainPressurePercent                                *float64   `json:"MaintainPressurePercent,omitempty"`
	// Packer Depth                                                   
	PackerMeasuredDepth                                    *float64   `json:"PackerMeasuredDepth,omitempty"`
	// Incremental Pressure Over Defined Time Interval                
	PressureChangePerTime                                  *float64   `json:"PressureChangePerTime,omitempty"`
	// Pumped Volume                                                  
	PumpedVolume                                           *float64   `json:"PumpedVolume,omitempty"`
	// Comments or Remarks                                            
	Remarks                                                *string    `json:"Remarks,omitempty"`
	// Test Criteria                                                  
	TestCriteria                                           *string    `json:"TestCriteria,omitempty"`
}
