package apimanagement


type ApiManagementSignUp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// terms_of_service block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#terms_of_service ApiManagement#terms_of_service}
	TermsOfService *ApiManagementSignUpTermsOfService `field:"required" json:"termsOfService" yaml:"termsOfService"`
}
