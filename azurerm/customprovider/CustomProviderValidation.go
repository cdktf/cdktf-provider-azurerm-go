package customprovider


type CustomProviderValidation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/custom_provider#specification CustomProvider#specification}.
	Specification *string `field:"required" json:"specification" yaml:"specification"`
}

