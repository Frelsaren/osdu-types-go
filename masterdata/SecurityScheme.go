package masterdata

// Metadata used to retrieve or generate credentials needed for authorization.
type SecurityScheme struct {
	// Key for obtaining an access token from the secret repository                                     
	AccessTokenKeyName                                                                          *string `json:"AccessTokenKeyName,omitempty"`
	// Key for obtaining an API key from the secret repository                                          
	APIKeyKeyName                                                                               *string `json:"APIKeyKeyName,omitempty"`
	// Audience used in OAuth2 flows                                                                    
	Audience                                                                                    *string `json:"Audience,omitempty"`
	// Authorization endpoint of the identity provider used in OAuth2 flows                             
	AuthorizationURL                                                                            *string `json:"AuthorizationUrl,omitempty"`
	// Url the identity provider will send the token to in OAuth2 authorization code flows              
	CallbackURL                                                                                 *string `json:"CallbackUrl,omitempty"`
	// DEPRECATED: Superseded  by ClientIDKeyName. Client ID                                            
	ClientID                                                                                    *string `json:"ClientID,omitempty"`
	// Key for obtaining the client id for use in the OAuth2 flows from the secret repository           
	ClientIDKeyName                                                                             *string `json:"ClientIDKeyName,omitempty"`
	// Key for obtaining the client secret for use in OAuth2 flows from the secret repository           
	ClientSecretKeyName                                                                         *string `json:"ClientSecretKeyName,omitempty"`
	// OAuth2 flow, or grant type                                                                       
	FlowTypeID                                                                                  *string `json:"FlowTypeID,omitempty"`
	// Unique name given to a security scheme                                                           
	Name                                                                                        string  `json:"Name"`
	// Key for obtaining Password from the secret repository                                            
	PasswordKeyName                                                                             *string `json:"PasswordKeyName,omitempty"`
	// Key for obtaining a refresh token from the secret repository                                     
	RefreshTokenKeyName                                                                         *string `json:"RefreshTokenKeyName,omitempty"`
	// Key for obtaining the space-delimited scopes list for use in OAuth2 flows from the secret        
	// repository                                                                                       
	ScopesKeyName                                                                               *string `json:"ScopesKeyName,omitempty"`
	// Url of the secret repository containing secrets for this security scheme                         
	SecretRepoURL                                                                               *string `json:"SecretRepoUrl,omitempty"`
	// Token endpoint of the identity provider used in OAuth2 flows                                     
	TokenURL                                                                                    *string `json:"TokenUrl,omitempty"`
	// Type of security schema                                                                          
	TypeID                                                                                      string  `json:"TypeID"`
	// Key for obtaining Username from the secret repository                                            
	UsernameKeyName                                                                             *string `json:"UsernameKeyName,omitempty"`
}
