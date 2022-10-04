package apimanagement


type ApiManagementSignUpTermsOfService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#consent_required ApiManagement#consent_required}.
	ConsentRequired interface{} `field:"required" json:"consentRequired" yaml:"consentRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#text ApiManagement#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

