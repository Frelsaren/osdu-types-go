package workproductcomponent

// ISO13503_2 properties for this proppant agent.
type StimISO135032Properties struct {
	// The density the material would have if no intra-granular porosity is present. (e.g.                                   
	// Boyle's Law porosimetry).                                                                                             
	AbsoluteDensity                                                                             *float64                     `json:"AbsoluteDensity,omitempty"`
	// The solubility of a proppant in 12:3 HCl:HF for 30 minutes at 150°F is an indication of                               
	// the amount of soluble materials (i.e. carbonates, feldspars, iron oxides, clays, etc)                                 
	// present in the proppant.                                                                                              
	AcidSolubility                                                                              *float64                     `json:"AcidSolubility,omitempty"`
	// Apparent density excludes extra-granular porosity by placing a known mass in a volume of                              
	// fluid and determining how much of the fluid is displaced (Archimedes).                                                
	ApparentDensity                                                                             *float64                     `json:"ApparentDensity,omitempty"`
	// Bulk density includes both the proppant and the porosity. This is measured by filling a                               
	// known volume with dry proppant and measuring the weight.                                                              
	BulkDensity                                                                                 *float64                     `json:"BulkDensity,omitempty"`
	// Percentage of undesirable agglomerated discrete proppant particles which typically occurs                             
	// more with inefficiently processed natural sand proppants as opposed to manufactured                                   
	// ceramic proppants. ISO 13503-2 and API RP19C limit the mass of clusters to less than 1%.                              
	ClustersPercent                                                                             *float64                     `json:"ClustersPercent,omitempty"`
	// Crush Test Data                                                                                                       
	CrushTestData                                                                               []ISO135032CrushTestData     `json:"CrushTestData,omitempty"`
	// Crush test classification indicating the highest stress level at which a proppant                                     
	// generated no more than 10% crushed material rounded down to the nearest 1,000 psi during                              
	// a crush test. For example, a value of 14 means 14k which is 14000 psi.                                                
	KValue                                                                                      *int64                       `json:"KValue,omitempty"`
	// A mass loss (gravimetric) test method applied to coated proppants only, which determines                              
	// the mass of resin coating applied to a natural sand or manufactured proppant by means of                              
	// thorough combustion of the flammable resin from the nonflammable proppant. Reported as a                              
	// % of original mass.                                                                                                   
	LossOnIgnition                                                                              *float64                     `json:"LossOnIgnition,omitempty"`
	// The mean diameter of particles in a sample of proppant.                                                               
	MeanParticleDiameter                                                                        *float64                     `json:"MeanParticleDiameter,omitempty"`
	// The median diameter of particles in a sample of proppant.                                                             
	MedianParticleDiameter                                                                      *float64                     `json:"MedianParticleDiameter,omitempty"`
	// Krumbein Roundness Shape Factor that is a measure of the relative sharpness of grain                                  
	// corners or of grain curvature. Krumbein and Sloss (1963) are the most widely used method                              
	// of determining shape factors.                                                                                         
	Roundness                                                                                   *float64                     `json:"Roundness,omitempty"`
	// Sieve Analysis Data                                                                                                   
	SieveAnalysisData                                                                           []ISO135032SieveAnalysisData `json:"SieveAnalysisData,omitempty"`
	// Not formally part of ISO 13503.2 properties, the specific gravity is the apparent density                             
	// of the proppant divided by the density of water.                                                                      
	SpecificGravity                                                                             *float64                     `json:"SpecificGravity,omitempty"`
	// Krumbein Sphericity Shape Factor that is a measure of how closely a proppant particle                                 
	// approaches the shape of a sphere. Krumbein and Sloss (1963) are the most widely used                                  
	// method of determining shape factors.                                                                                  
	Sphericity                                                                                  *float64                     `json:"Sphericity,omitempty"`
	// A measure of water clarity, how much the material suspended in water decreases the                                    
	// passage of light through the water. Measured according to ISO 7027 and expressed in                                   
	// Formazin Nephelometric Units (FNU).                                                                                   
	TurbidityFNU                                                                                *float64                     `json:"TurbidityFNU,omitempty"`
	// A measure of water clarity, how much the material suspended in water decreases the                                    
	// passage of light through the water.   Measured according to USEPA 180.1 and expressed in                              
	// Nephelometric Turbidity Units (NTU).                                                                                  
	TurbidityNTU                                                                                *float64                     `json:"TurbidityNTU,omitempty"`
}
