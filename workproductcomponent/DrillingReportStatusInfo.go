package workproductcomponent

import "time"

// Status information related to an Operations Report
type DrillingReportStatusInfo struct {
	// A pointer to the latest BHA Run used for this reporting period                                   
	BHARunID                                                                                  *string   `json:"BHARunID,omitempty"`
	// Diameter of the last casing joint.                                                               
	CasingDiameterLast                                                                        *float64  `json:"CasingDiameterLast,omitempty"`
	// Measured depth of the last casing joint.                                                         
	CasingMDLast                                                                              *float64  `json:"CasingMDLast,omitempty"`
	// True vertical depth of last casing joint.                                                        
	CasingTVDLast                                                                             *float64  `json:"CasingTVDLast,omitempty"`
	// Description of the hole condition.                                                               
	ConditionHole                                                                             *string   `json:"ConditionHole,omitempty"`
	// The cost per day of the mud used                                                                 
	CostDayMud                                                                                *float64  `json:"CostDayMud,omitempty"`
	// Hole nominal inside diameter.                                                                    
	DiameterHole                                                                              *float64  `json:"DiameterHole,omitempty"`
	// Measured depth to the start of the current hole diameter.                                        
	DiameterHoleStartMD                                                                       *float64  `json:"DiameterHoleStartMD,omitempty"`
	// Pilot hole nominal inside diameter.                                                              
	DiameterPilot                                                                             *float64  `json:"DiameterPilot,omitempty"`
	// The planned measured depth of the pilot hole.                                                    
	DiameterPilotPlanMD                                                                       *float64  `json:"DiameterPilotPlanMD,omitempty"`
	// The planned true vertical depth of the pilot hole.                                               
	DiameterPilotPlanTVD                                                                      *float64  `json:"DiameterPilotPlanTVD,omitempty"`
	// Distance drilled: rotating.                                                                      
	DistDrillRot                                                                              *float64  `json:"DistDrillRot,omitempty"`
	// Distance drilled: sliding.                                                                       
	DistDrillSlid                                                                             *float64  `json:"DistDrillSlid,omitempty"`
	// Distance covered while holding angle with a steerable drilling assembly.                         
	DistHold                                                                                  *float64  `json:"DistHold,omitempty"`
	// Distance reamed.                                                                                 
	DistReam                                                                                  *float64  `json:"DistReam,omitempty"`
	// Distance covered while actively steering with a steerable drilling assembly.                     
	DistSteering                                                                              *float64  `json:"DistSteering,omitempty"`
	// Distance drilled.  This should be measured along the centerline of the wellbore.                 
	DrilledDistance                                                                           *float64  `json:"DrilledDistance,omitempty"`
	// Time spent circulating from the start of the bit run.                                            
	ElapsedTimeCirc                                                                           *float64  `json:"ElapsedTimeCirc,omitempty"`
	// Drilling time.                                                                                   
	ElapsedTimeDrill                                                                          *float64  `json:"ElapsedTimeDrill,omitempty"`
	// Time spent rotary drilling.                                                                      
	ElapsedTimeDrillRot                                                                       *float64  `json:"ElapsedTimeDrillRot,omitempty"`
	// Time spent slide drilling from the start of the bit run.                                         
	ElapsedTimeDrillSlid                                                                      *float64  `json:"ElapsedTimeDrillSlid,omitempty"`
	// Time spent with no directional drilling work (commonly in hours).                                
	ElapsedTimeHold                                                                           *float64  `json:"ElapsedTimeHold,omitempty"`
	// Time the rig has been on location (commonly in days).                                            
	ElapsedTimeLOC                                                                            *float64  `json:"ElapsedTimeLoc,omitempty"`
	// Time spent reaming from the start of the bit run.                                                
	ElapsedTimeReam                                                                           *float64  `json:"ElapsedTimeReam,omitempty"`
	// Time since the bit broke ground (commonly in days).                                              
	ElapsedTimeSpud                                                                           *float64  `json:"ElapsedTimeSpud,omitempty"`
	// Time from the start of operations (commonly in days).                                            
	ElapsedTimeStart                                                                          *float64  `json:"ElapsedTimeStart,omitempty"`
	// Time spent steering the bottomhole assembly (commonly in hours).                                 
	ElapsedTimeSteering                                                                       *float64  `json:"ElapsedTimeSteering,omitempty"`
	// Name of the operator's drilling engineer.                                                        
	Engineer                                                                                  *string   `json:"Engineer,omitempty"`
	// A summary of  planned activities for the next reporting period.                                  
	Forecast24Hr                                                                              *string   `json:"Forecast24Hr,omitempty"`
	// The measured formation strength. This should be the final measurement before the end of          
	// the report period.                                                                               
	FormationStrength                                                                         *float64  `json:"FormationStrength,omitempty"`
	// The measured depth of the formation strength measurement.                                        
	FormationStrengthMD                                                                       *float64  `json:"FormationStrengthMD,omitempty"`
	// The true vertical depth of the formation strength measurement.                                   
	FormationStrengthTVD                                                                      *float64  `json:"FormationStrengthTVD,omitempty"`
	// Name of operator's wellsite geologist.                                                           
	Geologist                                                                                 *string   `json:"Geologist,omitempty"`
	// Measured depth to the kickoff point of the wellbore.                                             
	KickoffMD                                                                                 *float64  `json:"KickoffMD,omitempty"`
	// True vertical depth to the kickoff point of the wellbore.                                        
	KickoffTVD                                                                                *float64  `json:"KickoffTVD,omitempty"`
	// Maximum allowable shut-in casing pressure.                                                       
	Maasp                                                                                     *float64  `json:"MAASP,omitempty"`
	// The measured depth planned to be reached.                                                        
	MdPlanned                                                                                 *float64  `json:"MdPlanned,omitempty"`
	// Authorization for expenditure (AFE) number that this cost item applies to.                       
	NumAFE                                                                                    *string   `json:"NumAFE,omitempty"`
	// Number of contractor personnel on the rig.                                                       
	NumContract                                                                               *int64    `json:"NumContract,omitempty"`
	// Number of operator personnel on the rig.                                                         
	NumOperator                                                                               *int64    `json:"NumOperator,omitempty"`
	// Total number of personnel on board the rig.                                                      
	NumPOB                                                                                    *int64    `json:"NumPOB,omitempty"`
	// Number of service company personnel on the rig.                                                  
	NumService                                                                                *int64    `json:"NumService,omitempty"`
	// The measured plug back depth.                                                                    
	PlugTopMD                                                                                 *float64  `json:"PlugTopMD,omitempty"`
	// Kick tolerance pressure.                                                                         
	PresKickTol                                                                               *float64  `json:"PresKickTol,omitempty"`
	// Leak off test equivalent mud weight.                                                             
	PresLotEmw                                                                                *float64  `json:"PresLotEmw,omitempty"`
	// The type of pressure test that was run.                                                          
	PresTestType                                                                              *string   `json:"PresTestType,omitempty"`
	// Wellbore measured depth at the end of the report period.                                         
	ReportEndMD                                                                               *float64  `json:"ReportEndMD,omitempty"`
	// Wellbore true vertical depth at the end of the report.                                           
	ReportEndTVD                                                                              *float64  `json:"ReportEndTVD,omitempty"`
	// Name of the Rig/Work Unit used for the reporting period                                          
	Rig                                                                                       *string   `json:"Rig,omitempty"`
	// A pointer to the rig used.                                                                       
	RigID                                                                                     *string   `json:"RigID,omitempty"`
	// Average rate of penetration.                                                                     
	RopAverage                                                                                *float64  `json:"RopAverage,omitempty"`
	// Rate of penetration at the end of the reporting period.                                          
	RopCurrent                                                                                *float64  `json:"RopCurrent,omitempty"`
	// The date and time for which the well status is reported.                                         
	StatusDateTime                                                                            time.Time `json:"StatusDateTime"`
	// A summary of the activities performed and the status of the ongoing activities.                  
	Summary24Hr                                                                               *string   `json:"Summary24Hr,omitempty"`
	// Name of the rig supervisor                                                                       
	Supervisor                                                                                *string   `json:"Supervisor,omitempty"`
	// A pointer to the tubular (assembly) used in this report period.                                  
	TubularID                                                                                 *string   `json:"TubularID,omitempty"`
	// True vertical depth of a leak off test point.                                                    
	TvdLot                                                                                    *float64  `json:"TvdLot,omitempty"`
	// Type of wellbore.                                                                                
	TypeWellbore                                                                              *string   `json:"TypeWellbore,omitempty"`
	// Kick tolerance volume.                                                                           
	VolKickTol                                                                                *float64  `json:"VolKickTol,omitempty"`
}
