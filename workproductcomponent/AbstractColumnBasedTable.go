package workproductcomponent

// Table describing the observations (estimations or measurements) captured from this WPC.
// It can either be restricted to "Properties Descriptions" only when we have a separated
// content storage OR used to capture the stream content itself whenever required
//
// The ColumnBasedTable is a set of columns, which have equal length (data.ColumnSize)
// included by types carrying embedded table properties. Columns have a Property Kind,
// UnitOfMeasure and Facet. There are KeyColumns (index columns) and Columns (for look-up
// values). Examples are KrPc, PVT and Facies tables.
//
// The column based table with the estimated volumes.
//
// An embedded ColumnBasedTable containing the Measurements within the period that are
// required (and only if required) for indexing
//
// An embedded ColumnBasedTable with the properties including their values associated to the
// intervals in data.Intervals[]. The association is done by array index.
//
// An embedded ColumnBasedTable with the properties including their values associated to the
// intervals in data.Markers[]. The association is done by array index.
type AbstractColumnBasedTable struct {
	// Optional relationship to a ColumnBasedTableTemplate record, which defines the KeyColumn                                                      
	// and Column definitions. Some columns defined in the template may be omitted if not                                                           
	// contained in the ColumnValues, but the ones used must be exactly identical to the                                                            
	// template's column definitions. If the ColumnBasedTableTemplateID is populated, the                                                           
	// ColumnBasedTableType is expected to be ColumnBasedTableTemplateControlled.                                                                   
	ColumnBasedTableTemplateID                                                                  *string                                             `json:"ColumnBasedTableTemplateID,omitempty"`
	// Quickly indicate the type of the column based table (KrPc, PVT, Facies, ...) and its                                                         
	// standard columns definition. It is supposed to be used when you don't use KeyColumns                                                         
	// neither Columns as attributes of this WPC.                                                                                                   
	ColumnBasedTableType                                                                        *string                                             `json:"ColumnBasedTableType,omitempty"`
	// A common column storing values of a particular property kind. Do not use this attribute                                                      
	// if you want to follow a given ColumnBasedTableType.                                                                                          
	Columns                                                                                     []AbstractReferencePropertyType                     `json:"Columns,omitempty"`
	// The count of elements in each column, i.e. the number of rows in the ColumnBasedTable.                                                       
	// All columns must have the same size, including placeholder values for the undefined cells                                                    
	// identified by ColumnValues[].UndefinedValueRows[].                                                                                           
	ColumnSize                                                                                  *int64                                              `json:"ColumnSize,omitempty"`
	// First column values are related to first key column, second column values are related to                                                     
	// the second key column, etc…                                                                                                                  
	// Column values at index KeyColumns count are related to first (non key) column, Column                                                        
	// values at index KeyColumns count + 1 are related to second (non key) column, etc...                                                          
	ColumnValues                                                                                []ProductionValuesObservationDescriptionColumnValue `json:"ColumnValues,omitempty"`
	// A column whose values are considered as keys/indices. Do not use this attribute if you                                                       
	// want to follow a given ColumnBasedTableType.                                                                                                 
	KeyColumns                                                                                  []AbstractReferencePropertyType                     `json:"KeyColumns,omitempty"`
}
