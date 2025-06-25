package referencedata

// Enumerated string indicating whether to use the normal scalar field for scaling this
// field (STANDARD), no scaling (NOSCALE), or override scalar (OVERRIDE).  Default is
// current STANDARD (such as SEG-Y rev2).
type ScalarIndicator string

const (
	Noscale  ScalarIndicator = "NOSCALE"
	Override ScalarIndicator = "OVERRIDE"
	Standard ScalarIndicator = "STANDARD"
)
