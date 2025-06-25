package masterdata

// The prevailing operating conditions (Pressure and Temperature)  for the recombination
// operation.
//
// The pair of absolute pressure and temperature values describing the condition for a
// particular volume measurement or estimation. The unit of measure context is defined via
// the meta[] block in the record. Search responses will return pressure in Pa (Pascal) and
// temperature in K (Kelvin).
//
// The pressure and temperature reference values for the gas stream.
//
// The pressure and temperature reference values for the oil stream.
//
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
//
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
//
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
//
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
//
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
//
// Used to describe the pressure and temperature conditions at which the sample preparation
// took place
//
// The pressure and temperature conditions recorded when the current sample container is
// closed for the current chain of custody event.
//
// The pressure and temperature conditions recorded when the previous sample container is
// opened for the current chain of custody event.
//
// The pressure and temperature conditions recorded during the sample transfer operation
// between containers for the current chain of custody event.
// Eg. if ingesting from PRODML Sample object, then the mapping can be seen below:
// TransferCondition.Pressure =
// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].TransferPressure
// TransferCondition.Temperature =
// PRODML:2.1:FluidSample.FluidSampleChainOfCustodyEvent[].TransferTemperature
//
// This provides the recommended operating conditions (Pressure and Temperature) rating for
// the sample container.
type AbstractPTCondition struct {
	// The recorded absolute pressure condition. The unit of measure context is defined via              
	// meta[] in the Storage record while the Search responses return the value in base SI unit          
	// Pa (Pascal).                                                                                      
	Pressure                                                                                     float64 `json:"Pressure"`
	// The recorded temperature condition. The unit of measure context is defined via meta[] in          
	// the Storage record while the Search responses return the value in base SI unit K (Kelvin).        
	Temperature                                                                                  float64 `json:"Temperature"`
}
