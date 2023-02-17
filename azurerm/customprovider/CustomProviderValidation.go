package customprovider


type CustomProviderValidation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/custom_provider#specification CustomProvider#specification}.
	Specification *string `field:"required" json:"specification" yaml:"specification"`
}

