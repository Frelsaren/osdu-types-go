package workproductcomponent

import "time"

// Details about this flowback test step.
type StepElement struct {
	// Bottomhole flow rate for the specific step.                                                        
	BHFlowRate                                                                                 *float64   `json:"BHFlowRate,omitempty"`
	// Time stamp of the pressure measurement.                                                            
	DateTime                                                                                   *time.Time `json:"DateTime,omitempty"`
	// Calculated entry friction loss accounting for perforation and near wellbore restrictions           
	// for the specific step.                                                                             
	EntryFriction                                                                              *float64   `json:"EntryFriction,omitempty"`
	// Flowback rate.                                                                                     
	FlowbackRate                                                                               *float64   `json:"FlowbackRate,omitempty"`
	// Cumulative volume of flowback since the start of the test.                                         
	FlowbackVolume                                                                             *float64   `json:"FlowbackVolume,omitempty"`
	// Calculated near-wellbore friction loss for the specific step.                                      
	NearWellboreFriction                                                                       *float64   `json:"NearWellboreFriction,omitempty"`
	// Calculated perforation friction loss for the specific step.                                        
	PerfFrictionLoss                                                                           *float64   `json:"PerfFrictionLoss,omitempty"`
	// Calculated pipe friction loss for the specific step.                                               
	PipeFrictionLoss                                                                           *float64   `json:"PipeFrictionLoss,omitempty"`
	// The number of the step. Identifies the step within the step down test.                             
	StepNumber                                                                                 *int64     `json:"StepNumber,omitempty"`
	// Surface flow rate entering the well for the specific step.                                         
	SurfaceFlowRate                                                                            *float64   `json:"SurfaceFlowRate,omitempty"`
	// Surface pressure measured for the specific step.                                                   
	SurfacePressure                                                                            *float64   `json:"SurfacePressure,omitempty"`
}
