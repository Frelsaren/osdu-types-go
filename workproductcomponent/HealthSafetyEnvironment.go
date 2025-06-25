package workproductcomponent

import "time"

// Health Safety or Environment events that occurred since the last drilling/operation
// report. Captures data related to HSE events (e.g., tests, inspections, meetings, and
// drills), test values (e.g., pressure tested to), and/or incidents (e.g., discharges,
// non-compliance notices received, etc.).
type HealthSafetyEnvironment struct {
	// Comments and remarks                                                   
	Comments                                                       *string    `json:"Comments,omitempty"`
	// Incident free duration (commonly in days).                             
	DaysIncidentFree                                               *int64     `json:"DaysIncidentFree,omitempty"`
	// Last abandonment drill.                                                
	LastAbandonDrillDate                                           *time.Time `json:"LastAbandonDrillDate,omitempty"`
	// Last blow out preventer drill.                                         
	LastBopDrillDate                                               *time.Time `json:"LastBopDrillDate,omitempty"`
	// Last blow out preventer pressure test.                                 
	LastBopPressureTestDate                                        *time.Time `json:"LastBopPressureTestDate,omitempty"`
	// Last casing pressure test date and time.                               
	LastCasingPressureTestDate                                     *time.Time `json:"LastCasingPressureTestDate,omitempty"`
	// Last diverter drill.                                                   
	LastDiverterDrillDate                                          *time.Time `json:"LastDiverterDrillDate,omitempty"`
	// Last fire or life boat drill.                                          
	LastFireBoatDrillDate                                          *time.Time `json:"LastFireBoatDrillDate,omitempty"`
	// Last rig inspection/check.                                             
	LastRigInspectionDate                                          *time.Time `json:"LastRigInspectionDate,omitempty"`
	// Last safety inspection.                                                
	LastSafetyInspectionDate                                       *time.Time `json:"LastSafetyInspectionDate,omitempty"`
	// Last safety meeting.                                                   
	LastSafetyMeetingDate                                          *time.Time `json:"LastSafetyMeetingDate,omitempty"`
	// Last trip drill.                                                       
	LastTripDrillDate                                              *time.Time `json:"LastTripDrillDate,omitempty"`
	// Next blow out preventer pressure test.                                 
	NextBopPresTestDate                                            *time.Time `json:"NextBopPresTestDate,omitempty"`
	// Inspection non-compliance notice served?                               
	NonComplianceIssued                                            *bool      `json:"NonComplianceIssued,omitempty"`
	// Blow out preventer annular preventer pressure tested to.               
	PressureAnnular                                                *float64   `json:"PressureAnnular,omitempty"`
	// Blow out preventer ram pressure tested to.                             
	PressureBOPRAM                                                 *float64   `json:"PressureBOPRam,omitempty"`
	// Choke line pressure tested to.                                         
	PressureChokeLine                                              *float64   `json:"PressureChokeLine,omitempty"`
	// Choke line manifold pressure tested to.                                
	PressureChokeMan                                               *float64   `json:"PressureChokeMan,omitempty"`
	// Blow out preventer diverter pressure tested to.                        
	PressureDiverter                                               *float64   `json:"PressureDiverter,omitempty"`
	// Kelly hose pressure tested to.                                         
	PressureKellyHose                                              *float64   `json:"PressureKellyHose,omitempty"`
	// Last casing pressure test pressure.                                    
	PressureLastCasing                                             *float64   `json:"PressureLastCasing,omitempty"`
	// Standpipe manifold pressure tested to.                                 
	PressureStandPipeManifold                                      *float64   `json:"PressureStandPipeManifold,omitempty"`
	// Governmental regulatory inspection agency inspection?                  
	RegulatoryAgencyInspection                                     *bool      `json:"RegulatoryAgencyInspection,omitempty"`
	// Number of health, safety and environment incidents reported.           
	TotalStopCards                                                 *int64     `json:"TotalStopCards,omitempty"`
	// Volume of cuttings discharged.                                         
	VolumeCuttingDischarged                                        *float64   `json:"VolumeCuttingDischarged,omitempty"`
	// Daily whole mud discarded.                                             
	VolumeFluidDischarged                                          *float64   `json:"VolumeFluidDischarged,omitempty"`
	// Oil on cuttings daily discharge.                                       
	VolumeOilCuttingDischarge                                      *float64   `json:"VolumeOilCuttingDischarge,omitempty"`
	// Volume of sanitary waste discharged.                                   
	VolumeWasteDischarged                                          *float64   `json:"VolumeWasteDischarged,omitempty"`
}
