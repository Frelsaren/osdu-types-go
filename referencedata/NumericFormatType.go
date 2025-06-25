package referencedata

// A hint how the number is expected to be presented, e.g., d or D for decimal, f or F for
// fixed point, e or E for exponential (scientific), or g or G for general (default). Not
// all languages support all codes in all cases - in principle the case means that the
// resulting case is transformed to upper case or lower case depending on the case of the
// NumericFormatType.
type NumericFormatType string

const (
	D                  NumericFormatType = "d"
	E                  NumericFormatType = "e"
	F                  NumericFormatType = "f"
	G                  NumericFormatType = "g"
	NumericFormatTypeD NumericFormatType = "D"
	NumericFormatTypeE NumericFormatType = "E"
	NumericFormatTypeF NumericFormatType = "F"
	NumericFormatTypeG NumericFormatType = "G"
)
