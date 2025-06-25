package referencedata

// The type of coordinate transformation. This is an enumeration of concrete sub-types.
// Transformation is a single operation between a source and a target geodetic CRS;
// ConcatenatedOperation is a chained set of Transformations.
type CTType string

const (
	ConcatenatedOperation CTType = "ConcatenatedOperation"
	Transformation        CTType = "Transformation"
)
