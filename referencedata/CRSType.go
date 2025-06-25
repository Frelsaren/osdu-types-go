package referencedata

// The type of coordinate reference system. This is an enumeration of concrete sub-types.
type CRSType string

const (
	BoundCRS       CRSType = "BoundCRS"
	CompoundCRS    CRSType = "CompoundCRS"
	DerivedCRS     CRSType = "DerivedCRS"
	EngineeringCRS CRSType = "EngineeringCRS"
	GeodeticCRS    CRSType = "GeodeticCRS"
	ProjectedCRS   CRSType = "ProjectedCRS"
	VerticalCRS    CRSType = "VerticalCRS"
)
