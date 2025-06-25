package workproductcomponent

// If populated, it contains information that additional content is available containing
// grain size analysis results. The array AnalysisTypeIDs must then contain a value of
// GrainSizeAnalysis.
type GrainSizeAnalysis struct {
	// The relationship to the GrainSizeClassificationScheme reference value record, which         
	// defined the grain size interval to GrainSizeClassification mapping. It provides the         
	// detailed context to the MeanGrainSizeDescription.                                           
	GrainSizeClassificationSchemeID                                                       *string  `json:"GrainSizeClassificationSchemeID,omitempty"`
	// The type of grain size analysis carried out during this analysis.                           
	GrainSizeMethodTypeID                                                                 *string  `json:"GrainSizeMethodTypeID,omitempty"`
	// The mean grain size as determined in the analysis. It holds the relationship to a           
	// GrainSizeClassification reference value.                                                    
	MeanGrainSizeClassificationID                                                         *string  `json:"MeanGrainSizeClassificationID,omitempty"`
	// The median grain size as numeric value, as determined in this analysis.                     
	MedianGrainSize                                                                       *float64 `json:"MedianGrainSize,omitempty"`
}
