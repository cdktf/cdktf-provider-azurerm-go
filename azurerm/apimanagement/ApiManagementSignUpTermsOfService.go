package apimanagement


type ApiManagementSignUpTermsOfService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/api_management#consent_required ApiManagement#consent_required}.
	ConsentRequired interface{} `field:"required" json:"consentRequired" yaml:"consentRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/api_management#text ApiManagement#text}.
	Text *string `field:"optional" json:"text" yaml:"text"`
}

