package masterdata

// This attribute provides information about the acquisition parameters and process used in
// acquiring the target sample. Other information about the sample itself can be found in
// the Sample object.
type SampleAcquisitionDetail struct {
	// This captures the prevailing pressure and temperature recorded at the sampling point                                                                                                            
	// during the sample acquisition event. The property applies to ALL sampling acquisition                                                                                                           
	// events.                                                                                                                                                                                         
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// AcquisitionCondition.Pressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FacilitySampleAcquisition|DownholeSampleAcquisition                                                   
	// | FormationSampleAcquisition | SeparatorSampleAcquisition | WellheadSampleAcquisition                                                                                                           
	// >[].Item.AcquisitionPressure                                                                                                                                                                    
	//                                                                                                                                                                                                 
	// AcquisitionCondition.Temperature=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FacilitySampleAcquisition|DownholeSampleAcquisition                                                
	// | FormationSampleAcquisition | SeparatorSampleAcquisition | WellheadSampleAcquisition                                                                                                           
	// >[].Item.AcquisitionTemperature                                                                                                                                                                 
	AcquisitionCondition                                                                                                                                          *AbstractPTCondition                 `json:"AcquisitionCondition,omitempty"`
	// This captures the gas oil ratio (GOR) for the sample acquired during the sample                                                                                                                 
	// acquisition event. The property applies to ALL FLUID sampling acquisition events.                                                                                                               
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// AcquisitionGOR=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FacilitySampleAcquisition|DownholeSampleAcquisition                                                                  
	// | FormationSampleAcquisition | SeparatorSampleAcquisition | WellheadSampleAcquisition                                                                                                           
	// >[].Item.AcquisitionGOR                                                                                                                                                                         
	AcquisitionGOR                                                                                                                                                *float64                             `json:"AcquisitionGOR,omitempty"`
	// The depth of the base of the target interval from which the sample was acquired. The                                                                                                            
	// reference and kind of depth (e.g. driller's depth versus logger's depth) is described in                                                                                                        
	// data.VerticalMeasurement.  The property is always used except with                                                                                                                              
	// WellheadSampleAcquisition, SeparatorSampleAcquisition, FacilitySampleAcquisition.                                                                                                               
	// Note: As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                        
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// TopDepth=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<DownholeSampleAcquisition                                                                                                  
	// | FormationTestSampleAcquisition>[].Item.BaseMD                                                                                                                                                 
	BaseDepth                                                                                                                                                     *float64                             `json:"BaseDepth,omitempty"`
	// This refers to an array of OSDU record IDs for the wellbore completion objects                                                                                                                  
	// (perforated or open hole interval) that contributed to the stream used in acquiring the                                                                                                         
	// sample. It can be used for one wellbore opening or multiple opening in the case of                                                                                                              
	// commingled flow. This is typically related to acquisition events from downhole where                                                                                                            
	// samples are comingled from multiple completions. The property is only used in conjunction                                                                                                       
	// with WellheadSampleAcquisition, SeparatorSampleAcquisition, DownholeSampleAcquisition.                                                                                                          
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// WellboreOpeningIDs[0]=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<WellheadSampleAcquisition                                                                                     
	// | SeparatorSampleAcquisition>[].Item.WellboreCompletion                                                                                                                                         
	ContributingWellboreOpeningIDs                                                                                                                                []string                             `json:"ContributingWellboreOpeningIDs,omitempty"`
	// The value accounts for the application of correction procedures to the gas flow rate                                                                                                            
	// observed / measured during the sample acquisition event. The property is typically only                                                                                                         
	// used in conjunction with SeparatorSampleAcquisition.                                                                                                                                            
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// CorrectedGasRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.CorrectedGasRate                                                                
	CorrectedGasRate                                                                                                                                              *float64                             `json:"CorrectedGasRate,omitempty"`
	// The value accounts for the application of correction procedures to the oil flow rate                                                                                                            
	// observed / measured during the sample acquisition event. The property is typically only                                                                                                         
	// used in conjunction with SeparatorSampleAcquisition                                                                                                                                             
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// CorrectedOilRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.CorrectedOilRate                                                                
	CorrectedOilRate                                                                                                                                              *float64                             `json:"CorrectedOilRate,omitempty"`
	// The value accounts for the application of correction procedures to the water flow rate                                                                                                          
	// observed / measured during the sample acquisition event. The property is typically only                                                                                                         
	// used in conjunction with SeparatorSampleAcquisition.                                                                                                                                            
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// CorrectedWaterRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.CorrectedWaterRate                                                            
	CorrectedWaterRate                                                                                                                                            *float64                             `json:"CorrectedWaterRate,omitempty"`
	// This is used to capture information regarding the methodology used in correcting rates                                                                                                          
	// acquired during the sample acquisition event. The property is only used in conjunction                                                                                                          
	// with SeparatorSampleAcquisition                                                                                                                                                                 
	CorrectionRemarks                                                                                                                                             *AbstractRemark                      `json:"CorrectionRemarks,omitempty"`
	// This captures the operating conditions (prevailing pressure and temperatures) on the                                                                                                            
	// target equipment used on Topside Facilities (exclusive of wells or separators) during the                                                                                                       
	// sample acquisition event.  This attribute is provided in the event that the acquisition                                                                                                         
	// pressure and temperature recorded at the flow port or sampling point from which the                                                                                                             
	// sample is acquired is different from the operating P&T for the target facility or                                                                                                               
	// equipment. The property is only used in conjunction with FacilitySampleAcquisition                                                                                                              
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// FacilityOperatingCondition.Pressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FacilitySampleAcquisition>[].Item.FacilityPressure                                              
	//                                                                                                                                                                                                 
	// FacilityOperatingCondition.Temperature=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FacilitySampleAcquisition>[].Item.FacilityTemperature                                        
	FacilityEquipmentOperatingCondition                                                                                                                           *AbstractPTCondition                 `json:"FacilityEquipmentOperatingCondition,omitempty"`
	// This captures the operating conditions (prevailing pressure and temperatures) on the                                                                                                            
	// target formation during the sample acquisition event.  This attribute is provided in the                                                                                                        
	// event that the acquisition pressure and temperature recorded at the downhole sampling                                                                                                           
	// location is different from the Formation's P&T. The property is used in conjunction with                                                                                                        
	// all the acquisition event types except FacilitySampleAcquisition, Cuttings and                                                                                                                  
	// WellheadSampleAcquisition.                                                                                                                                                                      
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// FormationCondition.Pressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<WellheadSampleAcquisition>[].Item.FormationPressure                                                     
	//                                                                                                                                                                                                 
	// FormationCondition.Temperature=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<WellheadSampleAcquisition>[].Item.FormationTemperature                                               
	FormationCondition                                                                                                                                            *AbstractPTCondition                 `json:"FormationCondition,omitempty"`
	// This is the OSDU record ID for the predominant fluid kind obtained from the formation                                                                                                           
	// during the acquisition event. The property is only used in conjunction with                                                                                                                     
	// FormationTestSampleAcquisition                                                                                                                                                                  
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// GrossFluidKind=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FormationTestSampleAcquisition>[].Item.GrossFluidKind                                                                
	GrossFluidKindID                                                                                                                                              *string                              `json:"GrossFluidKindID,omitempty"`
	// This is the observed/ measured gas rate for this sample acquisition event. The property                                                                                                         
	// is only used in conjunction with SeparatorSampleAcquisition                                                                                                                                     
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// MeasuredGasRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.MeasuredGasRate                                                                  
	MeasuredGasRate                                                                                                                                               *float64                             `json:"MeasuredGasRate,omitempty"`
	// This is the  observed/ measured oil rate for this sample acquisition event. The property                                                                                                        
	// is only used in conjunction with SeparatorSampleAcquisition                                                                                                                                     
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// MeasuredOilRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.MeasuredOilRate                                                                  
	MeasuredOilRate                                                                                                                                               *float64                             `json:"MeasuredOilRate,omitempty"`
	// This is the observed/ measured water rate for this sample acquisition event. The property                                                                                                       
	// is only used in conjunction with SeparatorSampleAcquisition                                                                                                                                     
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// MeasuredWaterRate=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.MeasuredWaterRate                                                              
	MeasuredWaterRate                                                                                                                                             *float64                             `json:"MeasuredWaterRate,omitempty"`
	// This is the OSDU record ID for the type of mud base used  during the acquisition event or                                                                                                       
	// present in the sample required. The property is always used except with Outcrop                                                                                                                 
	MudBaseTypeID                                                                                                                                                 *string                              `json:"MudBaseTypeID,omitempty"`
	// This property is used in capturing the type of tracer used during the sample acquisition                                                                                                        
	// event.The property is always used except with Outcrop                                                                                                                                           
	MudTracerTypeID                                                                                                                                               *string                              `json:"MudTracerTypeID,omitempty"`
	// The kind of preservation applied to this sample if applied at the time of acquisition.                                                                                                          
	// The property is only used in conjunction with ConventionalCore, Sidewall Core, Cuttings,                                                                                                        
	// Outcrop, Core Plugs                                                                                                                                                                             
	PreservationTypeID                                                                                                                                            *string                              `json:"PreservationTypeID,omitempty"`
	// This refers to the different runs performed during the sample acquisition event and is                                                                                                          
	// typically identified using integers. It mostly applies to acquisition events acquired                                                                                                           
	// from the subsurface like downhole, coring, etc. The property is always used except with                                                                                                         
	// WellheadSampleAcquisition, SeparatorSampleAcquisition, FormationTestSampleAcquisition,                                                                                                          
	// FacilitySampleAcquisition                                                                                                                                                                       
	RunNumber                                                                                                                                                     *string                              `json:"RunNumber,omitempty"`
	// The name or identifier for the slot in the sample carrier where the sample was acquired.                                                                                                        
	// The property is only used in conjunction with FormationTestSampleAcquisition                                                                                                                    
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// SampleCarrierSlotName=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FormationTestSampleAcquisition>[].Item.SampleCarrierSlotName                                                  
	SampleCarrierSlotName                                                                                                                                         *string                              `json:"SampleCarrierSlotName,omitempty"`
	// The pressure used in charging the sample container. The property is only used in                                                                                                                
	// conjunction with FormationTestSampleAcquisition                                                                                                                                                 
	// Note: As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                        
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// SampleContainerCushionPressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FormationTestSampleAcquisition>[].Item.CushionPressure                                               
	SampleContainerCushionPressure                                                                                                                                *float64                             `json:"SampleContainerCushionPressure,omitempty"`
	// Actual length of recovered sample, usually a core The property is only used in                                                                                                                  
	// conjunction with ConventionalCore, Sidewall Core, Outcrop                                                                                                                                       
	SampleRecoveredLengthActual                                                                                                                                   *float64                             `json:"SampleRecoveredLengthActual,omitempty"`
	// Planned length of sample to be recovered, usually a core The property is only used in                                                                                                           
	// conjunction with ConventionalCore, Sidewall Core, Cuttings, Outcrop                                                                                                                             
	SampleRecoveredLengthPlanned                                                                                                                                  *float64                             `json:"SampleRecoveredLengthPlanned,omitempty"`
	// A free-form reference to the flow port on the Facility where this sample was acquired.                                                                                                          
	// The property is only used in conjunction with WellheadSampleAcquisition,                                                                                                                        
	// SeparatorSampleAcquisition, FacilitySampleAcquisition.                                                                                                                                          
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	// SamplingPoint= [                                                                                                                                                                                
	// PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition |                                                                                                        
	// FacilitySampleAcquisition | SeparatorSampleAcquisition>[].Item.SamplingPoint                                                                                                                    
	SamplingPoint                                                                                                                                                 *SamplingPoint                       `json:"SamplingPoint,omitempty"`
	// This captures the operating conditions (prevailing pressure and temperatures) on the                                                                                                            
	// target facility or equipment (in this case separator) during the sample acquisition                                                                                                             
	// event.  This attribute is provided in the event that the acquisition pressure and                                                                                                               
	// temperature recorded at the flow port or sampling point from which the sample is acquired                                                                                                       
	// is different from the operating P&T for the separator. The property is only used in                                                                                                             
	// conjunction with SeparatorSampleAcquisition                                                                                                                                                     
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// SeparatorOperatingCondition.Pressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.SeparatorPressure                                           
	//                                                                                                                                                                                                 
	// SeparatorOperatingCondition.Temperature=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<SeparatorSampleAcquisition>[].Item.SeparatorTemperature                                     
	SeparatorOperatingCondition                                                                                                                                   *AbstractPTCondition                 `json:"SeparatorOperatingCondition,omitempty"`
	// This is the OSDU ID for the Site where the the sample acquisition event occurred.                                                                                                               
	SiteID                                                                                                                                                        *string                              `json:"SiteID,omitempty"`
	// This references the kind of tool used in acquiring the sample. The property is always                                                                                                           
	// used except with WellheadSampleAcquisition, SeparatorSampleAcquisition,                                                                                                                         
	// FormationTestSampleAcquisition, FacilitySampleAcquisition.                                                                                                                                      
	// Note: As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                        
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// ToolKind=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<DownholeSampleAcquisition>[].Item.ToolKind.                                                                                
	ToolKind                                                                                                                                                      *string                              `json:"ToolKind,omitempty"`
	// The name of the formation tester tool section that was used in acquiring the sample. The                                                                                                        
	// property is only used in conjunction with FormationTestSampleAcquisition,                                                                                                                       
	// ConventionalCore, Sidewall Core, Cuttings                                                                                                                                                       
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// ToolSectionName=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<FormationTestSampleAcquisition>[].Item.ToolSectionName                                                              
	ToolSectionName                                                                                                                                               *string                              `json:"ToolSectionName,omitempty"`
	// The depth of the top of the target interval from which the sample was acquired. The                                                                                                             
	// reference and kind of depth (e.g. driller's depth versus logger's depth) is described in                                                                                                        
	// data.VerticalMeasurement. The property is always used except with                                                                                                                               
	// WellheadSampleAcquisition, SeparatorSampleAcquisition, FacilitySampleAcquisition.                                                                                                               
	// Note: As an example, if ingesting data formatted using PRODML, this is typically  mapped                                                                                                        
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// TopDepth=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<DownholeSampleAcquisition                                                                                                  
	// | FormationTestSampleAcquisition>[].Item.TopMD                                                                                                                                                  
	TopDepth                                                                                                                                                      *float64                             `json:"TopDepth,omitempty"`
	// Information on the list of all depths and elevations pertaining to the target wellbore                                                                                                          
	// involved in the Sample Acquisition event, like, plug back measured depth, total measured                                                                                                        
	// depth, KB elevation. The property is always used except with WellheadSampleAcquisition,                                                                                                         
	// SeparatorSampleAcquisition, FacilitySampleAcquisition                                                                                                                                           
	VerticalMeasurement                                                                                                                                           *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// This refers to the OSDU record ID of the wellbore object from which the sample was                                                                                                              
	// acquired. It typically applies in scenarios where the acquisition event only pertains to                                                                                                        
	// a single wellbore object. The property is always used except with                                                                                                                               
	// FacilitySampleAcquisition, Outcrop, SeparatorSampleAcquisition                                                                                                                                  
	// Note: If ingesting data formatted using PRODML, this is typically  mapped as seen below:                                                                                                        
	// WellboreID=                                                                                                                                                                                     
	// PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<DownholeSampleAcquisition |                                                                                                         
	// WellheadSampleAcquisition>[].Item.Wellbore                                                                                                                                                      
	WellboreID                                                                                                                                                    *string                              `json:"WellboreID,omitempty"`
	// This captures the operating conditions (prevailing pressure and temperatures) on the                                                                                                            
	// target facility or equipment (in this case wellhead) during the sample acquisition                                                                                                              
	// event.  This attribute is provided in the event that the acquisition pressure and                                                                                                               
	// temperature recorded at the flow port or sampling point from which the sample is acquired                                                                                                       
	// is different from the operating P&T at the wellhead. The property is only used in                                                                                                               
	// conjunction with WellheadSampleAcquisition.                                                                                                                                                     
	// Note:As an example, If ingesting data formatted using PRODML, this is typically  mapped                                                                                                         
	// as seen below:                                                                                                                                                                                  
	//                                                                                                                                                                                                 
	// wellheadOperatingCondition.Pressure=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<WellheadSampleAcquisition>[].Item.wellheadPressure                                              
	//                                                                                                                                                                                                 
	// wellheadOperatingCondition.Temperature=PRODML:2.1:FluidSampleAcquisitionJob.FluidSampleAcquisition<WellheadSampleAcquisition>[].Item.wellheadTemperature                                        
	WellheadOperatingCondition                                                                                                                                    *AbstractPTCondition                 `json:"WellheadOperatingCondition,omitempty"`
}
