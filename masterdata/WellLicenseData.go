package masterdata

// Common resources to be injected at root 'data' level for every entity, which is
// persistable in Storage. The insertion is performed by the OsduSchemaComposer script.
//
// Properties shared with all master-data schema instances.
type WellLicenseData struct {
	// Where does this data resource sit in the cradle-to-grave span of its existence?                                                
	ExistenceKind                                                                                *string                              `json:"ExistenceKind,omitempty"`
	// Describes the current Curation status.                                                                                         
	ResourceCurationStatus                                                                       *string                              `json:"ResourceCurationStatus,omitempty"`
	// The name of the home [cloud environment] region for this OSDU resource object.                                                 
	ResourceHomeRegionID                                                                         *string                              `json:"ResourceHomeRegionID,omitempty"`
	// The name of the host [cloud environment] region(s) for this OSDU resource object.                                              
	ResourceHostRegionIDs                                                                        []string                             `json:"ResourceHostRegionIDs,omitempty"`
	// Describes the current Resource Lifecycle status.                                                                               
	ResourceLifecycleStatus                                                                      *string                              `json:"ResourceLifecycleStatus,omitempty"`
	// Classifies the security level of the resource.                                                                                 
	ResourceSecurityClassification                                                               *string                              `json:"ResourceSecurityClassification,omitempty"`
	// The entity that produced the record, or from which it is received; could be an                                                 
	// organization, agency, system, internal team, or individual. For informational purposes                                         
	// only, the list of sources is not governed.                                                                                     
	Source                                                                                       *string                              `json:"Source,omitempty"`
	// DEPRECATED: Describes a record's overall suitability for general business consumption                                          
	// based on data quality. Clarifications: Since Certified is the highest classification of                                        
	// suitable quality, any further change or versioning of a Certified record should be                                             
	// carefully considered and justified. If a Technical Assurance value is not populated then                                       
	// one can assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                       
	// Technical Assurance values are not intended to be used for the identification of a single                                      
	// "preferred" or "definitive" record by comparison with other records.                                                           
	TechnicalAssuranceID                                                                         *string                              `json:"TechnicalAssuranceID,omitempty"`
	// List of geographic entities which provide context to the master data. This may include                                         
	// multiple types or multiple values of the same type.                                                                            
	GeoContexts                                                                                  []AbstractGeoContext                 `json:"GeoContexts,omitempty"`
	// Alternative names, including historical, by which this master data is/has been known (it                                       
	// should include all the identifiers).                                                                                           
	NameAliases                                                                                  []AbstractAliasNames                 `json:"NameAliases,omitempty"`
	// The spatial location information such as coordinates, CRS information (left empty when                                         
	// not appropriate).                                                                                                              
	SpatialLocation                                                                              *AbstractSpatialLocation             `json:"SpatialLocation,omitempty"`
	// Describes a record's overall suitability for general business consumption in context of                                        
	// one or more workflows/personas based on data quality and reviewer's decisions.                                                 
	// Clarifications: Since Certified is the highest classification of suitable quality, any                                         
	// further change or versioning of a Certified record should be carefully considered and                                          
	// justified. If a Technical Assurance value is not populated then one can assume the data                                        
	// has not been evaluated or its quality is unknown (=Unevaluated). Technical Assurance                                           
	// values are not intended to be used for the identification of a single "preferred" or                                           
	// "definitive" record by comparison with other records.                                                                          
	TechnicalAssurances                                                                          []AbstractTechnicalAssurance         `json:"TechnicalAssurances,omitempty"`
	// DEPRECATED: (in favor of more nuanced TechnicalAssurances[] array) Describes a                                                 
	// master-data record's overall suitability for general business consumption based on data                                        
	// quality. Clarifications: Since Certified is the highest classification of suitable                                             
	// quality, any further change or versioning of a Certified record should be carefully                                            
	// considered and justified. If a Technical Assurance value is not populated then one can                                         
	// assume the data has not been evaluated or its quality is unknown (=Unevaluated).                                               
	// Technical Assurance values are not intended to be used for the identification of a single                                      
	// "preferred" or "definitive" record by comparison with other records.                                                           
	TechnicalAssuranceTypeID                                                                     *string                              `json:"TechnicalAssuranceTypeID,omitempty"`
	// This describes the reason that caused the creation of a new version of this master data.                                       
	VersionCreationReason                                                                        *string                              `json:"VersionCreationReason,omitempty"`
	// The validated contractor for this license, as approved.  The term contractor has variable                                      
	// interpretations globally, in this instance the contractor is generally meant to be the                                         
	// contractor who will conduct initial operations on the well for drilling to final TD.                                           
	ApprovedContractorID                                                                         *string                              `json:"ApprovedContractorID,omitempty"`
	// Number for the approved drill slot on the offshore platform. Platforms have 24 to 32                                           
	// slots in which to position derricks for drilling. Each slot may require a different                                            
	// license in some jurisdictions, while others approve all operations on a platform at the                                        
	// same time. If this property is absent, the license applies to all slots.                                                       
	ApprovedDrillSlot                                                                            *string                              `json:"ApprovedDrillSlot,omitempty"`
	// The approved purpose for the well proposal in narrative form.  This often contains                                             
	// formation or depth information along with aggregated information from the Well Status and                                      
	// Classifications facets Business Intention, Play Type, Role, Product Type and Fluid                                             
	// Direction. The Well License Status subset allows this information to be decomposed into                                        
	// facets for better clarity.                                                                                                     
	ApprovedPurposeDescription                                                                   *string                              `json:"ApprovedPurposeDescription,omitempty"`
	// Sometimes specific rigs may be authorized by a license, particularly where environmental                                       
	// or access restrictions set specific requirements.  This attribute identifies the specific                                      
	// rig approved, where needed.                                                                                                    
	ApprovedRigID                                                                                *string                              `json:"ApprovedRigID,omitempty"`
	// The height of the foundation on which the derrick and engine sits.                                                             
	ApprovedRigSubstructureHeight                                                                *float64                             `json:"ApprovedRigSubstructureHeight,omitempty"`
	// The type of rig installation the license is approved for, where this is relevant. For                                          
	// example land, barge, submersible, platform, jackup, drillship, semisub, artificial                                             
	// island...                                                                                                                      
	ApprovedRigTypeID                                                                            *string                              `json:"ApprovedRigTypeID,omitempty"`
	// The approved spud date, can be the same as shown on the license application or may be                                          
	// changed as a result of the discussions involved in the approval process or changes to the                                      
	// circumstances surrounding the well plan.                                                                                       
	ApprovedSpudDate                                                                             *string                              `json:"ApprovedSpudDate,omitempty"`
	// Surveying is a specialized skill that requires certification before it can be practiced                                        
	// in most regions.  The name of the surveyor or surveying company provides evidence that a                                       
	// properly qualified individual will conduct the survey.                                                                         
	ApprovedSurveyorID                                                                           *string                              `json:"ApprovedSurveyorID,omitempty"`
	// The WellProductType(s) (e.g., fluid) that is/are approved for the well license.  In some                                       
	// regions, WellProductTypes are approved individually or as a group.  Similarly, in some                                         
	// regions, specific WellProductTypes may be excluded from the well license, particularly if                                      
	// the land right obtained does not include those WellProductTypes.                                                               
	ApprovedTargetObjectiveProductIDs                                                            []string                             `json:"ApprovedTargetObjectiveProductIDs,omitempty"`
	// A well license may be assigned to one or more wellbores at the time the license is                                             
	// granted. Other licenses may be assigned to specific well components as defined in What is                                      
	// a Well.  Where the component is known, specific relationships are needed.                                                      
	ApprovedWellboreIDs                                                                          []string                             `json:"ApprovedWellboreIDs,omitempty"`
	// The number of wells that are approved under this license.  Not all regulators define                                           
	// wells the same way.  Some regulators may identify the number of wellbores, others the                                          
	// number of well origins. Refer to PPDM "What is a Well" for more information about well                                         
	// components.                                                                                                                    
	ApprovedWellCount                                                                            *int64                               `json:"ApprovedWellCount,omitempty"`
	// A well license may be assigned to an entire planned well configuration at the time the                                         
	// license is granted. Other licenses may be assigned to specific well components.  Where                                         
	// the component is known, specific relationships are needed.                                                                     
	ApprovedWellIDs                                                                              []string                             `json:"ApprovedWellIDs,omitempty"`
	// The formation authorized for use. May not be the same as the projected formation at the                                        
	// final depth of a wellbore.                                                                                                     
	AuthorizedStratigraphicUnit                                                                  *string                              `json:"AuthorizedStratigraphicUnit,omitempty"`
	// For the formation specified in the license, the use that this formation may be used for,                                       
	// such as production, injection, disposal, storage etc.                                                                          
	AuthorizedStratigraphicUnitUseRoleID                                                         *string                              `json:"AuthorizedStratigraphicUnitUseRoleID,omitempty"`
	// Bidding round number for drilling rights that resulted in this license.                                                        
	BiddingRoundIdentifier                                                                       *string                              `json:"BiddingRoundIdentifier,omitempty"`
	// The date on which a license has expired.  Note that this may be different than the due                                         
	// date, particularly where extensions are granted or specific conditions as stated in the                                        
	// license are met or not met.                                                                                                    
	ExpiredDate                                                                                  *string                              `json:"ExpiredDate,omitempty"`
	// The date when the license has or is due to expire.  In many cases, the date when the                                           
	// license term will expire is stated in the license.  This date may therefore be a date in                                       
	// the future. There may be conditions under which the license can be extended.  Where this                                       
	// occurs, the extension should usually be captured as a new version of the license.                                              
	ExpiryDueDate                                                                                *string                              `json:"ExpiryDueDate,omitempty"`
	// Identify the business associate, most often a regulator or authorized agency, has granted                                      
	// the license.  In some cases, other authorities may grant licenses, attribute is not                                            
	// intended to be proscriptive.                                                                                                   
	GrantedByBusinessAssociateID                                                                 *string                              `json:"GrantedByBusinessAssociateID,omitempty"`
	// A Y/N flag indicating whether this license is currently either active / valid (Y) or                                           
	// inactive / invalid (N).                                                                                                        
	IsActive                                                                                     *bool                                `json:"IsActive,omitempty"`
	// When a license is delayed, set this flag to True.  Details about the application can be                                        
	// found in the application object.                                                                                               
	IsDelayed                                                                                    *bool                                `json:"IsDelayed,omitempty"`
	// When the license expires, set the expiry date to the appropriate date, and set this flag                                       
	// to True.                                                                                                                       
	IsExpired                                                                                    *bool                                `json:"IsExpired,omitempty"`
	// Details about conditions imposed on a license by an authority, a partner or other                                              
	// stakeholder.                                                                                                                   
	LicenseConditions                                                                            []AbstractWellLicenseCondition       `json:"LicenseConditions,omitempty"`
	// Date the well license was issued.  This date may not be the same as the date entered into                                      
	// the system.                                                                                                                    
	LicenseDate                                                                                  *string                              `json:"LicenseDate,omitempty"`
	// The business associate who is the contact representing the licensee for this license.                                          
	LicenseeContactBusinessAssociateID                                                           *string                              `json:"LicenseeContactBusinessAssociateID,omitempty"`
	// The licensee to whom this license has been granted.  Sometimes, the licensees may be an                                        
	// address for service, particularly for confidential operations. This is not necessarily                                         
	// the name of the operator or the owner of the well.                                                                             
	LicenseeID                                                                                   *string                              `json:"LicenseeID,omitempty"`
	// The date on which the license has been issued and comes into effect.  Normally, no                                             
	// operations that involve surface or subsurface disturbance may be started before this date.                                     
	LicenseIssuedDate                                                                            *string                              `json:"LicenseIssuedDate,omitempty"`
	// The primary term of the license as set by the agency granting the license.                                                     
	LicensePrimaryTerm                                                                           *float64                             `json:"LicensePrimaryTerm,omitempty"`
	// The date the permit or license was reissued. The regulatory approval normally expires                                          
	// after the primary term unless an extension is requested and granted.                                                           
	LicenseReissueDate                                                                           *string                              `json:"LicenseReissueDate,omitempty"`
	// The secondary term of the license, often 6 - 12 months in duration.  These are granted by                                      
	// the regulatory authority after requests for renewal are submitted by the operator or the                                       
	// address for service.                                                                                                           
	LicenseReissueTerm                                                                           *float64                             `json:"LicenseReissueTerm,omitempty"`
	// Remarks and narratives associated to the license.                                                                              
	LicenseRemarks                                                                               []AbstractRemark                     `json:"LicenseRemarks,omitempty"`
	// Services associated with the license.                                                                                          
	LicenseServices                                                                              []LicenseServiceElement              `json:"LicenseServices,omitempty"`
	// Details about LicenseStates, past and present.                                                                                 
	LicenseStates                                                                                []LicenseStateElement                `json:"LicenseStates,omitempty"`
	// The type of license granted is often based on codes and rules set out in legislation or                                        
	// regulation.  This license type attribute should indicate the exact type as specified in                                        
	// the license document.                                                                                                          
	LicenseType                                                                                  *string                              `json:"LicenseType,omitempty"`
	// Details about violations of conditions, past and present.                                                                      
	LicenseViolations                                                                            []AbstractWellLicenseViolation       `json:"LicenseViolations,omitempty"`
	// The reference to the operating environment of the well/wellbore this license applies to.                                       
	OperatingEnvironmentID                                                                       *string                              `json:"OperatingEnvironmentID,omitempty"`
	// Projected measured depth at which the drilling of the primary wellbore is planned to be                                        
	// terminated. When multiple wellbores are planned, this may be the depth of the deepest                                          
	// wellbore, depending on local rules.                                                                                            
	ProjectedMeasuredDepth                                                                       *float64                             `json:"ProjectedMeasuredDepth,omitempty"`
	// The stratigraphic unit that is proposed for final completion and production (or                                                
	// injection). Note that this may not be the formation at the planned final depth of the                                          
	// wellbore. The stratigraphic unit should be part of the ProjectedStratigraphicColumn.                                           
	ProjectedStratigraphicUnitID                                                                 *string                              `json:"ProjectedStratigraphicUnitID,omitempty"`
	// Projected true vertical depth of the well determined during the permitting of the well.                                        
	ProjectedTrueVerticalDepth                                                                   *float64                             `json:"ProjectedTrueVerticalDepth,omitempty"`
	// The section of  the regulation under which the license was granted. For Example, the type                                      
	// of permit (such as Rule 37 or Rule 38 in the State of Texas) that is given for the well.                                       
	RegulationSection                                                                            *string                              `json:"RegulationSection,omitempty"`
	// A reference table describing the overall structure of the rig with respect to the                                              
	// environment and function. For example land, barge, submersible, platform, jackup,                                              
	// drillship, semisub or artificial island.                                                                                       
	//                                                                                                                                
	// Sometimes specific types of rigs may be authorized by a license, particularly where                                            
	// environmental or access restrictions set specific requirements.                                                                
	RigTypeID                                                                                    *string                              `json:"RigTypeID,omitempty"`
	// The name of the set of formation names used for interpretations.  Different name sets may                                      
	// exist on a regional or basin basis, or may be developed by regulatory agencies, data                                           
	// vendors or operating companies.  Clarity about which name set is used is critical to                                           
	// ensuring that the name of the formation is not ambiguous.                                                                      
	StratigraphicColumnID                                                                        *string                              `json:"StratigraphicColumnID,omitempty"`
	// The measured elevation from a known reference datum to a permanent geodetic datum.  This                                       
	// measurement is the foundation for calculating and correlating depths from geodetic                                             
	// datums, which in turn allow depth correlation between wellbores.                                                               
	VerticalMeasurement                                                                          *AbstractFacilityVerticalMeasurement `json:"VerticalMeasurement,omitempty"`
	ExtensionProperties                                                                          map[string]interface{}               `json:"ExtensionProperties,omitempty"`
}
