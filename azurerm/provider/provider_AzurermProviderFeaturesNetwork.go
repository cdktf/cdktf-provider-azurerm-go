package provider


type AzurermProviderFeaturesNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#relaxed_locking AzurermProvider#relaxed_locking}.
	RelaxedLocking interface{} `field:"required" json:"relaxedLocking" yaml:"relaxedLocking"`
}

