package provider


type AzurermProviderFeaturesCognitiveAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}.
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
}

