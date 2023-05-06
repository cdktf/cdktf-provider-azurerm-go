package provider


type AzurermProviderFeaturesResourceGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs#prevent_deletion_if_contains_resources AzurermProvider#prevent_deletion_if_contains_resources}.
	PreventDeletionIfContainsResources interface{} `field:"optional" json:"preventDeletionIfContainsResources" yaml:"preventDeletionIfContainsResources"`
}

