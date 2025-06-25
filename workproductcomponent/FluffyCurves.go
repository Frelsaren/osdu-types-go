package workproductcomponent

import "time"

// A curve is a data type that is represented by a series of digits, and are commonly
// displayed as a continuous line or a series of points referenced to the WellLog reference
// curve.   A WellLog commonly contains multiple curves.
type FluffyCurves struct {
	// DEPRECATED: Use `SamplingStop` for consistency. The curves maximum "depth" i.e., the                
	// reference value at which the curve has its last non-absent value. The curve may contain             
	// further absent values in between TopDepth and BaseDepth. Note that the SamplingDomainType           
	// may not be a depth as the property name indicates.                                                  
	BaseDepth                                                                                   *float64   `json:"BaseDepth,omitempty"`
	// Curve description is specific to that single curve mnemonic. In essence, curve                      
	// description defines the internal factors such as what the "curve" or measurement ideally            
	// is representing, how is it calculated, what are the assumptions and the "constants".                
	CurveDescription                                                                            *string    `json:"CurveDescription,omitempty"`
	// The ID of the Well Log Curve                                                                        
	CurveID                                                                                     *string    `json:"CurveID,omitempty"`
	// The Quality of the Log Curve.                                                                       
	CurveQuality                                                                                *string    `json:"CurveQuality,omitempty"`
	// The value type to be expected as curve sample values.                                               
	CurveSampleTypeID                                                                           *string    `json:"CurveSampleTypeID,omitempty"`
	// Unit of Measure for the Log Curve                                                                   
	CurveUnit                                                                                   *string    `json:"CurveUnit,omitempty"`
	// DEPRECATED: Refer to the `version` of the individual record. Log Version                            
	CurveVersion                                                                                *string    `json:"CurveVersion,omitempty"`
	// Date curve was created in the database                                                              
	DateStamp                                                                                   *time.Time `json:"DateStamp,omitempty"`
	// DEPRECATED: Use the `Curves[].IsRegularlySampled` property instead. The Coding of the               
	// depth.                                                                                              
	DepthCoding                                                                                 *string    `json:"DepthCoding,omitempty"`
	// Unit of Measure for TopDepth and BaseDepth.                                                         
	DepthUnit                                                                                   *string    `json:"DepthUnit,omitempty"`
	// DEPRECATED: Any curve can be interpreted. Whether curve can be interpolated or not                  
	Interpolate                                                                                 *bool      `json:"Interpolate,omitempty"`
	// The name of person who generated, improved or QC'd this Log Curve.  This excludes the               
	// name of the person who loaded or ingested the data.                                                 
	InterpreterName                                                                             *string    `json:"InterpreterName,omitempty"`
	// DEPRECATED: Highly subjective and difficult to ascertain what "Processed" means.                    
	// Indicates if the curve has been processed or re-processed. This could include both                  
	// wellsite downhole processing during acquisition or post-job processing from the office              
	IsProcessed                                                                                 *bool      `json:"IsProcessed,omitempty"`
	// Indicates if the curve is regularly or irregularly sampled. Not to be confused with the             
	// `data.IsRegular` attribute at Log Level which indicates regularity of the reference curve           
	// only.                                                                                               
	IsRegularlySampled                                                                          *bool      `json:"IsRegularlySampled,omitempty"`
	// The related record id of the Log Curve Business Value Type.                                         
	LogCurveBusinessValueID                                                                     *string    `json:"LogCurveBusinessValueID,omitempty"`
	// The related record id of the Log Curve Family - which is the detailed Geological Physical           
	// Quantity Measured - such as neutron porosity                                                        
	LogCurveFamilyID                                                                            *string    `json:"LogCurveFamilyID,omitempty"`
	// The related record id of the Log Curve Main Family Type - which is the Geological                   
	// Physical Quantity measured - such as porosity.                                                      
	LogCurveMainFamilyID                                                                        *string    `json:"LogCurveMainFamilyID,omitempty"`
	// DEPRECATED:  Please use the `LogCurveMainCurveFamily` and the `LogCurveFamily`                      
	// properties. This property is duplication. The related record id of the Log Curve Type -             
	// which is the standard mnemonic chosen by the company - OSDU provides an initial list                
	LogCurveTypeID                                                                              *string    `json:"LogCurveTypeID,omitempty"`
	// A short or abbreviated form of the curve name, typically provided by the logging vendor             
	// or the processing company.  Curve mnemonics have meaning to expert users.                           
	Mnemonic                                                                                    *string    `json:"Mnemonic,omitempty"`
	// Indicates that there is no measurement within the curve. This attribute is required for             
	// the Wellbore DDMS.                                                                                  
	NullValue                                                                                   *bool      `json:"NullValue,omitempty"`
	// The number of columns present in this Curve for a single reference value.  Most curves              
	// only have one column per curve.  A single curve may contain an array of columns, and                
	// these are commonly present in curves that display as images, for example Borehole Image             
	// logs or Variable Density Logs.                                                                      
	NumberOfColumns                                                                             *int64     `json:"NumberOfColumns,omitempty"`
	// The start or smallest value of the ReferenceCurveID, typically the start depth or time of           
	// the logging.                                                                                        
	SamplingStart                                                                               *float64   `json:"SamplingStart,omitempty"`
	// The stop or largest value of the ReferenceCurveID, typically the stop depth or time of              
	// the logging.                                                                                        
	SamplingStop                                                                                *float64   `json:"SamplingStop,omitempty"`
	// DEPRECATED: Use `SamplingStart` for consistency. The curves minimum "depth", i.e., the              
	// reference value at which the curve has its first non-absent value. The curve may contain            
	// further absent values in between TopDepth and BaseDepth. Note that the SamplingDomainType           
	// may not be a depth as the property name indicates.                                                  
	TopDepth                                                                                    *float64   `json:"TopDepth,omitempty"`
}
