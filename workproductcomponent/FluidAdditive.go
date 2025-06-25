package workproductcomponent

// Properties of an additive mixed into a Drilling, Stimulation or other type of Wellbore
// Fluid.
type FluidAdditive struct {
	// Code to identify additive by supplier's product number/code.                                     
	AdditiveCode                                                                                *string `json:"AdditiveCode,omitempty"`
	// Additive type or function                                                                        
	AdditiveKindID                                                                              *string `json:"AdditiveKindID,omitempty"`
	// The name of the additive.                                                                        
	AdditiveName                                                                                string  `json:"AdditiveName"`
	// The type of additive that is used, which can represent a suppliers description or type of        
	// AdditiveKind. For example, 5% HCl could be the type when AdditiveKind=acid.                      
	AdditiveType                                                                                *string `json:"AdditiveType,omitempty"`
	// The chemical abstract service number for this additive (CAS registry number).                    
	ChemicalAbstractServiceNumber                                                               *string `json:"ChemicalAbstractServiceNumber,omitempty"`
	// The additive material type.                                                                      
	MaterialTypeID                                                                              *string `json:"MaterialTypeID,omitempty"`
	// General remarks about this fluid additive.                                                       
	Remarks                                                                                     *string `json:"Remarks,omitempty"`
	// A code used to identify the supplier of the additive.                                            
	SupplierCode                                                                                *string `json:"SupplierCode,omitempty"`
	// The name of the additive supplier.                                                               
	SupplierName                                                                                *string `json:"SupplierName,omitempty"`
}
