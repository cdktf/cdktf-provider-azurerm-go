package customprovider


type CustomProviderAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/custom_provider#endpoint CustomProvider#endpoint}.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/custom_provider#name CustomProvider#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

