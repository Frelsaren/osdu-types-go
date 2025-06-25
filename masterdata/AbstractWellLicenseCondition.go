package masterdata

// The Well License Condition subset is necessarily abstract, because it is not possible to
// explicitly capture every kind of condition that may be imposed on a license by an
// authority, a partner or other stakeholder.  Conditions may be simple or complex, and are
// usually triggered by date or activity.  They also may specify requirements for how and
// when they must be completed.  License conditions are increasingly more rigorous,  and
// data submissions are  considered critical elements by most regulators. Some conditions
// are one time, and others are repeating.  In general, repeating conditions are tracked as
// obligations.
//
// Conditions often have to do with regulatory compliance, environmental stewardship,
// reporting operations, providing technical or business data, paying fees or obtaining
// additional approvals under specific circumstances.
//
// Failure to comply with license conditions can result in fines or other penalties, and
// license suspension or even revocation.
type AbstractWellLicenseCondition struct {
	// Narrative descriptive remarks that accompany the license condition. Could include the              
	// definition at which the values apply (such as 101.325 kilopascals and 15 degrees Celsius).         
	ConditionDescription                                                                         *string  `json:"ConditionDescription,omitempty"`
	// Names or code values assigned to a well license condition. Can include any type of                 
	// condition with the exception of values related to stratigraphic units (explicit),                  
	// Products (explicit) or values (use numeric values).                                                
	ConditionNames                                                                               []string `json:"ConditionNames,omitempty"`
	// The type of condition applied to the well license, such as flaring rate, venting rate,             
	// production rate, H2S content limit, emissions etc.                                                 
	ConditionTypeID                                                                              *string  `json:"ConditionTypeID,omitempty"`
	// The value applied to this condition. For example, an NOX emission limitation may be set            
	// at 0.02 kg / hr and the production rate for oil set at 10 m3/day.                                  
	ConditionValue                                                                               *float64 `json:"ConditionValue,omitempty"`
	// The contact person or organization that provides support for this condition.  Normally             
	// this would be someone at the agency who imposed the condition.                                     
	ContactBusinessAssociateID                                                                   *string  `json:"ContactBusinessAssociateID,omitempty"`
	// The state that must be achieved for the condition to become effective. For example, a              
	// report may be due 60 days after operations commence (or cease).                                    
	DueCondition                                                                                 *string  `json:"DueCondition,omitempty"`
	// The date that this condition must be fulfilled.                                                    
	DueDate                                                                                      *string  `json:"DueDate,omitempty"`
	// The frequency with which this condition must be met, such as an annual review. Where               
	// conditions are cycling in this way, please use the project (for work flow) or obligations          
	// ( for payments) module to track completions.                                                       
	DueFrequency                                                                                 *string  `json:"DueFrequency,omitempty"`
	// The period within which this condition must be satisfied, usually following completion of          
	// operations.                                                                                        
	DueTerm                                                                                      *float64 `json:"DueTerm,omitempty"`
	// The date on which this row of data first came into effect from a business perspective.             
	EffectiveDate                                                                                *string  `json:"EffectiveDate,omitempty"`
	// Unique identifier of this element in the parent's list of license conditions.                      
	ElementIdentifier                                                                            *string  `json:"ElementIdentifier,omitempty"`
	// The date on which this condition is no longer in effect for the well license.                      
	ExpiryDate                                                                                   *string  `json:"ExpiryDate,omitempty"`
	// The business associate who fulfilled this condition.                                               
	FulfilledBy                                                                                  *string  `json:"FulfilledBy,omitempty"`
	// Fulfilled data is the date that this condition was entirely completed or fulfilled.                
	// This date is assigned from the perspective of the business associate who fulfilled the             
	// condition.  Some conditions are ongoing, or repeat over time and are tracked as an                 
	// obligation.                                                                                        
	FulfilledDate                                                                                *string  `json:"FulfilledDate,omitempty"`
	// A flag indicating whether this condition is currently either active / valid (true) or              
	// inactive / invalid (false).                                                                        
	IsActive                                                                                     *bool    `json:"IsActive,omitempty"`
	// A Y/N flag indicating that the holder of the license is exempt from this condition.                
	IsExempt                                                                                     *string  `json:"IsExempt,omitempty"`
	// A flag indicating that this condition has been fulfilled. This is applicable when a                
	// condition is met once.                                                                             
	IsFulfilled                                                                                  *bool    `json:"IsFulfilled,omitempty"`
	// A stratigraphic column aggregating StratigraphicColumnInterpretations and eventually               
	// StratigraphicUnitInterpretations.                                                                  
	StratigraphicColumnID                                                                        *string  `json:"StratigraphicColumnID,omitempty"`
	// The optional column rank or level of StratigraphicRoleType (see StratigraphicRoleTypeID)           
	// that is identified.                                                                                
	StratigraphicColumnRankUnitTypeID                                                            *string  `json:"StratigraphicColumnRankUnitTypeID,omitempty"`
	// The type of stratigraphy, such as chronostratigraphic, lithostratigraphic,                         
	// biostratigraphic or sequence stratigraphic.                                                        
	StratigraphicRoleTypeID                                                                      *string  `json:"StratigraphicRoleTypeID,omitempty"`
	// The name of a stratigraphic unit that is named in a license condition.  This could be a            
	// formation in which aquifers must be protected, formations into which water or other waste          
	// disposal is authorized etc.                                                                        
	StratigraphicUnitInterpretationID                                                            *string  `json:"StratigraphicUnitInterpretationID,omitempty"`
	// When a product or other substance is named as an element in a well license condition, use          
	// this reference list to select the product or products.  Production, disposal, analysis,            
	// additives and other substances are valid as necessary to describe the condition.                   
	SubstanceID                                                                                  *string  `json:"SubstanceID,omitempty"`
	// The identity of any substance that escapes or spills, particularly when unintended.                
	// Intended substances and volumes are recorded as production values.                                 
	SubstanceLossID                                                                              *string  `json:"SubstanceLossID,omitempty"`
	// The type of loss experienced.  This may include fuel spills from a vehicle, gas emissions          
	// from a wellhead or feeder line, or any other substance.  This is an important part of              
	// carbon footprint analysis.                                                                         
	SubstanceLossTypeID                                                                          *string  `json:"SubstanceLossTypeID,omitempty"`
}
