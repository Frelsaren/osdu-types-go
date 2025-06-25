package referencedata

import "time"

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Generic reference object containing the universal properties of reference data,
// especially the ones commonly thought of as Types
type ParameterTypeData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                 
	ExistenceKind                                                                               *string                `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                          
	ResourceCurationStatus                                                                      *string                `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                  
	ResourceHomeRegionID                                                                        *string                `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                               
	ResourceHostRegionIDs                                                                       []string               `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                
	ResourceLifecycleStatus                                                                     *string                `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                  
	ResourceSecurityClassification                                                              *string                `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                  
	// organization, agency, system, internal team, or individual. For informational purposes                          
	// only, the list of sources is not governed.                                                                      
	Source                                                                                      *string                `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                           
	// based on data quality. Clarifications: Since Certified is the highest classification of                         
	// suitable quality, any further change or versioning of a Certified record should be                              
	// carefully considered and justified. If a Technical Assurance value is not populated then                        
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                        
	// Technical Assurance values are not intended to be used for the identification of a single                       
	// "preferred" or "definitive" record by comparison with other records.                                            
	TechnicalAssuranceID                                                                        *string                `json:"TechnicalAssuranceID,omitempty"`
	// Name of the authority, or organisation, which governs the entity value and from which it                        
	// is sourced.                                                                                                     
	AttributionAuthority                                                                        *string                `json:"AttributionAuthority,omitempty"`
	// Name, URL, or other identifier of the publication, or repository, of the attribution                            
	// source organisation from which the entity value is sourced.                                                     
	AttributionPublication                                                                      *string                `json:"AttributionPublication,omitempty"`
	// The distinct instance of the attribution publication, by version number, sequence number,                       
	// date of publication, etc., that was used for the entity value.                                                  
	AttributionRevision                                                                         *string                `json:"AttributionRevision,omitempty"`
	// The abbreviation or mnemonic for a reference type if defined. Example: WELL and WLBR.                           
	Code                                                                                        *string                `json:"Code,omitempty"`
	// For reference values published and governed by OSDU: The date and time the record was                           
	// committed into the OSDU member GitLab reference-values repository. The sole purpose of                          
	// this date is to optimise the OSDU milestone upgrades. It allows the upgrade code to                             
	// figure out whether or not the record must be PUT into reference value storage.                                  
	CommitDate                                                                                  *time.Time             `json:"CommitDate,omitempty"`
	// The text which describes a NAME TYPE in detail.                                                                 
	Description                                                                                 *string                `json:"Description,omitempty"`
	// Native identifier from a Master Data Management System or other trusted source external                         
	// to OSDU - stored here in order to allow for multi-system connection and synchronization.                        
	// If used, the "Source" property should identify that source system.                                              
	ID                                                                                          *string                `json:"ID,omitempty"`
	// By default reference values are considered as 'active'. An absent 'InactiveIndicator'                           
	// property value means the reference value is in active use. When 'InactiveIndicator' is                          
	// set true the reverence value is no longer in use and should no longer be offered as a                           
	// choice.                                                                                                         
	InactiveIndicator                                                                           *bool                  `json:"InactiveIndicator,omitempty"`
	// The name of the entity instance.                                                                                
	Name                                                                                        *string                `json:"Name,omitempty"`
	// Alternative names, including historical, by which this entity instance is/has been known.                       
	NameAlias                                                                                   []AbstractAliasNames   `json:"NameAlias,omitempty"`
	// The unit of measure for quantity.                                                                               
	DefaultUnitOfMeasureID                                                                      *string                `json:"DefaultUnitOfMeasureID,omitempty"`
	// The number of digits in the Length Quantity that are to the right of the decimal place.                         
	ParameterTypeDecimalPlaceCount                                                              *float64               `json:"ParameterTypeDecimalPlaceCount,omitempty"`
	// The total number of digits or strings by default allowed for the Parameter Type's value.                        
	ParameterTypeDefaultLengthCount                                                             *float64               `json:"ParameterTypeDefaultLengthCount,omitempty"`
	// The total precision by default allowed for the Parameter Type's value.                                          
	ParameterTypeDefaultPrecisionCount                                                          *float64               `json:"ParameterTypeDefaultPrecisionCount,omitempty"`
	// A general purpose field to identify a data in the form of DDMMYYYY.                                             
	ParameterTypeDefaultValueDate                                                               *string                `json:"ParameterTypeDefaultValueDate,omitempty"`
	// DateTime is a date value that represents a point in time on a calendar that is expressed                        
	// in centuries.                                                                                                   
	ParameterTypeDefaultValueDateTime                                                           *time.Time             `json:"ParameterTypeDefaultValueDateTime,omitempty"`
	// Indicates whether something is applicable to to the Entity. This can be Y or N.                                 
	ParameterTypeDefaultValueIndicator                                                          *bool                  `json:"ParameterTypeDefaultValueIndicator,omitempty"`
	// This generic data element represents a parameterized numeric value.                                             
	ParameterTypeDefaultValueQuantity                                                           *float64               `json:"ParameterTypeDefaultValueQuantity,omitempty"`
	// This generic data element represents a parameterized text (long) value.                                         
	ParameterTypeDefaultValueText                                                               *string                `json:"ParameterTypeDefaultValueText,omitempty"`
	// A general purpose field to identify a time in the form of hh24:mm:ss; hh:mm:ss am/pm.                           
	ParameterTypeDefaultValueTime                                                               *string                `json:"ParameterTypeDefaultValueTime,omitempty"`
	// The most commonly used highest number that is used to constrain the values of the                               
	// Parameter Type.                                                                                                 
	ParameterTypeMaximumValueQuantity                                                           *float64               `json:"ParameterTypeMaximumValueQuantity,omitempty"`
	// The most commonly used lowest number that is used to constrain the values of the                                
	// Parameter Type.                                                                                                 
	ParameterTypeMinimumValueQuantity                                                           *float64               `json:"ParameterTypeMinimumValueQuantity,omitempty"`
	// The quantity types examples are volumetric thermal expansion,linear thermal expansion,                          
	// length                                                                                                          
	QuantityTypeID                                                                              *string                `json:"QuantityTypeID,omitempty"`
	ExtensionProperties                                                                         map[string]interface{} `json:"ExtensionProperties,omitempty"`
}
