package masterdata

// Violations of license conditions may have serious consequences, including penalties or
// suspension or revocation of the license.  Any condition may be violated, and the
// regulator may provide notice to the address for service that the violation must be
// remedied. These processes are complex, but this simplified subset of attributes allows
// basic information to be captured.
type AbstractWellLicenseViolation struct {
	// Unique identifier of this element in the parent's list of license violations.                     
	ElementIdentifier                                                                           *string  `json:"ElementIdentifier,omitempty"`
	// Indicates whether this violation is still in effect, or has not been fully resolved.              
	IsActive                                                                                    *bool    `json:"IsActive,omitempty"`
	// The liquid or gaseous volume or other quantity that was lost or that escaped.                     
	QuantityLost                                                                                *float64 `json:"QuantityLost,omitempty"`
	// The duration over which the volume loss was recorded, usually in hours or days.                   
	QuantityLostDuration                                                                        *float64 `json:"QuantityLostDuration,omitempty"`
	// The condition(s) in a license that have been violated.  This may be a relationship to the         
	// specific conditions. The ViolatedCondition value refers to an array element in                    
	// data.LicenseConditions[].                                                                         
	ViolatedCondition                                                                           *string  `json:"ViolatedCondition,omitempty"`
	// The consequence of the violation as assigned by the regulatory or other authority.                
	// Consequence may be financial, work to be done, reports to be submitted. In addition, the          
	// authority may suspend, terminate or otherwise revise a license by adding new conditions           
	// or obligations.  For each violation, there may be more than one consequences.                     
	ViolationConsequenceTypeIDs                                                                 []string `json:"ViolationConsequenceTypeIDs,omitempty"`
	// The date on which the violation was first on record, such as the date that the regulatory         
	// agency notified you about a problem, or the date when something such as a report was              
	// overdue. Violations may be triggered by operations or events (or failure to complete              
	// something) or may be reported by another party.                                                   
	ViolationDate                                                                               *string  `json:"ViolationDate,omitempty"`
	// A supporting narrative that describes the violation.  While in some cases the nature of           
	// the violations, others require detailed explanations.  For example, environmental                 
	// violations often require explanatory text).                                                       
	//                                                                                                   
	// In populating descriptions of the violation and the resolution, ensure that the contents          
	// are appropriate to the attribute.  Do not put resolution information in the violation             
	// description, for example.                                                                         
	ViolationDescription                                                                        *string  `json:"ViolationDescription,omitempty"`
	// The date on which a violation is determined to have been fully resolved.                          
	ViolationEndDate                                                                            *string  `json:"ViolationEndDate,omitempty"`
	// A short narrative description of the specific resolution of the violation, such as                
	// procedures modified, fence repaired, water disposal corrected etc.                                
	ViolationResolutionDescription                                                              *string  `json:"ViolationResolutionDescription,omitempty"`
	// The method by which a license violation has been resolved.  In some cases, more than one          
	// method must be used, so multiple entries may be required, with appropriate dates                  
	// associated.  Methods may be financial (fines or other payments), operational (procedures          
	// altered),  rectification (things fixed or moved) or conciliatory (remedies to citizens or         
	// other stakeholders).                                                                              
	ViolationResolutionMethodTypeIDs                                                            []string `json:"ViolationResolutionMethodTypeIDs,omitempty"`
	// The date that this violation was resolved by submitting required information, paying              
	// fines, amending procedures etc. and necessary release notice provided (usually by the             
	// regulator).  This may or may not be the same as the ViolationEndDate.                             
	ViolationResolvedDate                                                                       *string  `json:"ViolationResolvedDate,omitempty"`
	// The date on which this violation was determined to be in effect.  This may not be the             
	// same date as the violation actually first began, ViolationDate.                                   
	ViolationStartDate                                                                          *string  `json:"ViolationStartDate,omitempty"`
	// The type of violation of a license that is being recorded. Can be as simple as failure to         
	// submit necessary reports or something more difficult such as improper procedures.                 
	// Grouping these violations by class supports reporting, metrics and analysis.                      
	ViolationTypeID                                                                             *string  `json:"ViolationTypeID,omitempty"`
	// Identify the wellbore  that is in violation. The condition violated can be tracked via            
	// ViolatedCondition referring to an object in data.LicenseConditions[].                             
	WellboreID                                                                                  *string  `json:"WellboreID,omitempty"`
	// Identify the well  that is in violation. The condition violated can be tracked via                
	// ViolatedCondition referring to an object in data.LicenseConditions[].                             
	WellID                                                                                      *string  `json:"WellID,omitempty"`
}
