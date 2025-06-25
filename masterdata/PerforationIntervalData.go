package masterdata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type PerforationIntervalData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                               
	ExistenceKind                                                                               *string                              `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                        
	ResourceCurationStatus                                                                      *string                              `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                
	ResourceHomeRegionID                                                                        *string                              `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                             
	ResourceHostRegionIDs                                                                       []string                             `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                              
	ResourceLifecycleStatus                                                                     *string                              `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                
	ResourceSecurityClassification                                                              *string                              `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                
	// organization, agency, system, internal team, or individual. For informational purposes                                        
	// only, the list of sources is not governed.                                                                                    
	Source                                                                                      *string                              `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                         
	// based on data quality. Clarifications: Since Certified is the highest classification of                                       
	// suitable quality, any further change or versioning of a Certified record should be                                            
	// carefully considered and justified. If a Technical Assurance value is not populated then                                      
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                      
	// Technical Assurance values are not intended to be used for the identification of a single                                     
	// "preferred" or "definitive" record by comparison with other records.                                                          
	TechnicalAssuranceID                                                                        *string                              `json:"TechnicalAssuranceID,omitempty"`
	// List of geographic entities which provide context to the master data. This may include                                        
	// multiple types or multiple values of the same type.                                                                           
	GeoContexts                                                                                 []AbstractGeoContext                 `json:"GeoContexts,omitempty"`
	// Alternative names, including historical, by which this master data is/has been known (it                                      
	// should include all the identifiers).                                                                                          
	NameAliases                                                                                 []AbstractAliasNames                 `json:"NameAliases,omitempty"`
	// The spatial location information such as coordinates, CRS information (left empty when                                        
	// not appropriate).                                                                                                             
	SpatialLocation                                                                             *AbstractSpatialLocation             `json:"SpatialLocation,omitempty"`
	// Describes a record's overall suitability for general business consumption in context of                                       
	// one or more workflows/personas based on data quality and reviewer's decisions.                                                
	// Clarifications: Since Certified is the highest classification of suitable quality, any                                        
	// further change or versioning of a Certified record should be carefully considered and                                         
	// justified. If a Technical Assurance value is not populated then one can assume the data                                       
	// has not been evaluated or its quality is unknown (=Unevaluated). Technical Assurance                                          
	// values are not intended to be used for the identification of a single "preferred" or                                          
	// "definitive" record by comparison with other records.                                                                         
	TechnicalAssurances                                                                         []AbstractTechnicalAssurance         `json:"TechnicalAssurances,omitempty"`
	// DEPRECATED: (in favor of more nuanced TechnicalAssurances[] array) Describes a                                                
	// master-data record's overall suitability for general business consumption based on data                                       
	// quality. Clarifications: Since Certified is the highest classification of suitable                                            
	// quality, any further change or versioning of a Certified record should be carefully                                           
	// considered and justified. If a Technical Assurance value is not populated then one can                                        
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                              
	// Technical Assurance values are not intended to be used for the identification of a single                                     
	// "preferred" or "definitive" record by comparison with other records.                                                          
	TechnicalAssuranceTypeID                                                                    *string                              `json:"TechnicalAssuranceTypeID,omitempty"`
	// This describes the reason that caused the creation of a new version of this master data.                                      
	VersionCreationReason                                                                       *string                              `json:"VersionCreationReason,omitempty"`
	// How was the bottom hole pressure determined?  Measured, estimated, etc.                                                       
	BHPressureTypeID                                                                            *string                              `json:"BHPressureTypeID,omitempty"`
	// Measured Depth of the Casing Collar Locator used to align perforating guns on depth                                           
	CasingCollarLocatorMD                                                                       *float64                             `json:"CasingCollarLocatorMD,omitempty"`
	// Distance from CCL to Interval Top                                                                                             
	CCLTopShotDistance                                                                          *float64                             `json:"CCLTopShotDistance,omitempty"`
	// Charge Description                                                                                                            
	ChargeDescription                                                                           *string                              `json:"ChargeDescription,omitempty"`
	// Charge Make/Manufacturer                                                                                                      
	ChargeManufacturerID                                                                        *string                              `json:"ChargeManufacturerID,omitempty"`
	// Charge Shape                                                                                                                  
	ChargeShapeID                                                                               *string                              `json:"ChargeShapeID,omitempty"`
	// Charge Size                                                                                                                   
	ChargeSize                                                                                  *string                              `json:"ChargeSize,omitempty"`
	// Charge Type                                                                                                                   
	ChargeTypeID                                                                                *string                              `json:"ChargeTypeID,omitempty"`
	// Cluster reference number                                                                                                      
	ClusterRefNo                                                                                *float64                             `json:"ClusterRefNo,omitempty"`
	// Crush Damage Ratio                                                                                                            
	CrushDamageRatio                                                                            *float64                             `json:"CrushDamageRatio,omitempty"`
	// Crush Zone Diameter                                                                                                           
	CrushZoneDiameter                                                                           *float64                             `json:"CrushZoneDiameter,omitempty"`
	// Initial discharge coefficient after perforating but prior to hydraulic fracturing.  A                                         
	// coefficient used in the equation for calculation of pressure drop                                                             
	// across a perforation set.                                                                                                     
	DischargeCoefficient                                                                        *float64                             `json:"DischargeCoefficient,omitempty"`
	// Pressure during perforation                                                                                                   
	DuringPerfGaugePressure                                                                     *float64                             `json:"DuringPerfGaugePressure,omitempty"`
	// Perforation Entrance Hole Diameter                                                                                            
	EntranceHoleDiameter                                                                        *float64                             `json:"EntranceHoleDiameter,omitempty"`
	// Final surface pressure                                                                                                        
	FinalSurfPressure                                                                           *float64                             `json:"FinalSurfPressure,omitempty"`
	// Static fluid level after perforating                                                                                          
	FluidAfterMD                                                                                *float64                             `json:"FluidAfterMD,omitempty"`
	// Static fluid level before perforating                                                                                         
	FluidBeforeMD                                                                               *float64                             `json:"FluidBeforeMD,omitempty"`
	// Fluid lossrate after perforating                                                                                              
	FluidLossAfterRate                                                                          *float64                             `json:"FluidLossAfterRate,omitempty"`
	// Fluid loss rate  before perforating                                                                                           
	FluidLossBeforeRate                                                                         *float64                             `json:"FluidLossBeforeRate,omitempty"`
	// Friction Factor                                                                                                               
	FrictionFactor                                                                              *float64                             `json:"FrictionFactor,omitempty"`
	// Friction Pressure                                                                                                             
	FrictionPressure                                                                            *float64                             `json:"FrictionPressure,omitempty"`
	// Depth of gauge(s) run while perforating.                                                                                      
	GaugeMD                                                                                     *float64                             `json:"GaugeMD,omitempty"`
	// Gun Carrier Category                                                                                                          
	GunCarrierCategoryTypeID                                                                    *string                              `json:"GunCarrierCategoryTypeID,omitempty"`
	// Gun Carrier Description                                                                                                       
	GunCarrierDescription                                                                       *string                              `json:"GunCarrierDescription,omitempty"`
	// Gun Carrier Manufacturer                                                                                                      
	GunCarrierManufacturerID                                                                    *string                              `json:"GunCarrierManufacturerID,omitempty"`
	// Gun Carrier Model                                                                                                             
	GunCarrierModelID                                                                           *string                              `json:"GunCarrierModelID,omitempty"`
	// Gun Carrier Phasing (angle between perfs). Implemented via a reference list in order to                                       
	// standardize the values.                                                                                                       
	GunCarrierPhasing                                                                           *string                              `json:"GunCarrierPhasing,omitempty"`
	// Gun Carrier Type                                                                                                              
	GunCarrierTypeID                                                                            *string                              `json:"GunCarrierTypeID,omitempty"`
	// Gun diameter                                                                                                                  
	GunDiameter                                                                                 *float64                             `json:"GunDiameter,omitempty"`
	// Gun Firing Head Type                                                                                                          
	GunFiringHeadTypeID                                                                         *string                              `json:"GunFiringHeadTypeID,omitempty"`
	// Metallurgy of the gun                                                                                                         
	GunMetallurgyTypeID                                                                         *string                              `json:"GunMetallurgyTypeID,omitempty"`
	// Gun Swell Diameter                                                                                                            
	GunSwellDiameter                                                                            *float64                             `json:"GunSwellDiameter,omitempty"`
	// Initial surface pressure                                                                                                      
	InitialSurfPressure                                                                         *float64                             `json:"InitialSurfPressure,omitempty"`
	// MD at base of perfed interval                                                                                                 
	IntervalBaseMD                                                                              *float64                             `json:"IntervalBaseMD,omitempty"`
	// Interval Comments                                                                                                             
	IntervalComments                                                                            *string                              `json:"IntervalComments,omitempty"`
	// Perforation interval date/time                                                                                                
	IntervalDateTime                                                                            time.Time                            `json:"IntervalDateTime"`
	// Distance between MD top and base of interval (net distance) aka 'Top Shot to Bottom Shot'                                     
	// Interval.                                                                                                                     
	IntervalLength                                                                              *float64                             `json:"IntervalLength,omitempty"`
	// Interval Reason                                                                                                               
	IntervalReasonTypeID                                                                        *string                              `json:"IntervalReasonTypeID,omitempty"`
	// MD at top of perfed interval                                                                                                  
	IntervalTopMD                                                                               float64                              `json:"IntervalTopMD"`
	// Perforated interval type                                                                                                      
	IntervalTypeID                                                                              string                               `json:"IntervalTypeID"`
	// Additional shot Y/N (+1 to Tot Shots calc)                                                                                    
	IsAdditionalShot                                                                            *bool                                `json:"IsAdditionalShot,omitempty"`
	// Total number of gun / charge misfires                                                                                         
	Misfires                                                                                    *int64                               `json:"Misfires,omitempty"`
	// Mid point perf pressure                                                                                                       
	MPPPressure                                                                                 *float64                             `json:"MPPPressure,omitempty"`
	// The name under which this perforation interval is known.                                                                      
	Name                                                                                        *string                              `json:"Name,omitempty"`
	// Estimated perforation penetration distance                                                                                    
	PenetrationDepth                                                                            *float64                             `json:"PenetrationDepth,omitempty"`
	// SRN to Job performed to construct the Perf Intervals                                                                          
	PerforationJobID                                                                            string                               `json:"PerforationJobID"`
	// Reservoir Temperature                                                                                                         
	ReservoirTemperature                                                                        *float64                             `json:"ReservoirTemperature,omitempty"`
	// Sequence Number (guns fired sequence)                                                                                         
	SequenceNumber                                                                              int64                                `json:"SequenceNumber"`
	// Gun Shot Density (shots per foot)                                                                                             
	ShotDensity                                                                                 *float64                             `json:"ShotDensity,omitempty"`
	// Technical Result                                                                                                              
	TechnicalResult                                                                             *string                              `json:"TechnicalResult,omitempty"`
	// ID to the Zero Depth Point elevation for depths contained in the perf job incl intervals,                                     
	// depth correction used to correlate MDs to original drilling rig MD. References an entry                                       
	// in the Vertical Measurement array for the Well parented by the wellbore via WellboreID.                                       
	VerticalMeasurement                                                                         *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	// Business natural key or code of the parent Wellbore to which this record belongs                                              
	WellboreID                                                                                  string                               `json:"WellboreID"`
	// Wireline diameter                                                                                                             
	WirelineDiameterTypeID                                                                      *string                              `json:"WirelineDiameterTypeID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
