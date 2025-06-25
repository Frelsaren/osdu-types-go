package workproductcomponent

// Indicate if a grid has been topologically expanded in a particular dimension (K
// expansion, J expansion, I expansion)
type ExpansionInDirection string

const (
	I ExpansionInDirection = "I"
	J ExpansionInDirection = "J"
	K ExpansionInDirection = "K"
)
