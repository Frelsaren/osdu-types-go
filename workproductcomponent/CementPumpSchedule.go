package workproductcomponent

import "time"

// Cement Pump Schedule Parameters
type CementPumpSchedule struct {
	// Back pressure applied during pumping stage.                                                       
	BackPressure                                                                              *float64   `json:"BackPressure,omitempty"`
	// Duration of the pumping schedule interval                                                         
	ElapsedTime                                                                               *float64   `json:"ElapsedTime,omitempty"`
	// Date/time when the pump schedule interval ended.                                                  
	EndDatetime                                                                               *time.Time `json:"EndDatetime,omitempty"`
	// The ratio of excess fluid to total fluid pumped during the step                                   
	ExcessFluidRatio                                                                          *float64   `json:"ExcessFluidRatio,omitempty"`
	// Cement slurry mix rate                                                                            
	MixRate                                                                                   *float64   `json:"MixRate,omitempty"`
	// Placement Method to land cement                                                                   
	PlacementMethod                                                                           *string    `json:"PlacementMethod,omitempty"`
	// Pumping Equipment Type                                                                            
	PumpEquipmentType                                                                         *string    `json:"PumpEquipmentType,omitempty"`
	// Pump Pressure Average                                                                             
	PumpPressureAvg                                                                           *float64   `json:"PumpPressureAvg,omitempty"`
	// Pump Pressure Maximum                                                                             
	PumpPressureMax                                                                           *float64   `json:"PumpPressureMax,omitempty"`
	// Pump Pressure Minimum                                                                             
	PumpPressureMin                                                                           *float64   `json:"PumpPressureMin,omitempty"`
	// Average Rate fluid is pumped. 0 means it is a pause.                                              
	PumpRateAvg                                                                               *float64   `json:"PumpRateAvg,omitempty"`
	// Maximum Rate fluid is pumped. 0 means it is a pause.                                              
	PumpRateMax                                                                               *float64   `json:"PumpRateMax,omitempty"`
	// Minimum Rate fluid is pumped. 0 means it is a pause.                                              
	PumpRateMin                                                                               *float64   `json:"PumpRateMin,omitempty"`
	// Pump stroke rate (SPM)                                                                            
	PumpStrokeRate                                                                            *float64   `json:"PumpStrokeRate,omitempty"`
	// Number of actual Pump strokes for the fluid to be pumped (assumed pump output known)              
	PumpStrokesActual                                                                         *int64     `json:"PumpStrokesActual,omitempty"`
	// Number of Pump strokes estimated for the fluid to be pumped (assumed pump output known)           
	PumpStrokesEstimated                                                                      *int64     `json:"PumpStrokesEstimated,omitempty"`
	// Volume pumped = eTimPump * ratePump.                                                              
	PumpVolume                                                                                *float64   `json:"PumpVolume,omitempty"`
	// Remarks                                                                                           
	Remark                                                                                    *string    `json:"Remark,omitempty"`
	// Sequence Number                                                                                   
	SequenceNumber                                                                            *int64     `json:"SequenceNumber,omitempty"`
	// Shutdown Elapsed Time                                                                             
	ShutdownElapsedTime                                                                       *float64   `json:"ShutdownElapsedTime,omitempty"`
	// Date/time when the pump schedule interval started.                                                
	StartDatetime                                                                             *time.Time `json:"StartDatetime,omitempty"`
}
