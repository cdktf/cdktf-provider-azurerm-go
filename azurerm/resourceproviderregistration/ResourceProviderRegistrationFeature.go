package resourceproviderregistration


type ResourceProviderRegistrationFeature struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_provider_registration#name ResourceProviderRegistration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_provider_registration#registered ResourceProviderRegistration#registered}.
	Registered interface{} `field:"required" json:"registered" yaml:"registered"`
}

