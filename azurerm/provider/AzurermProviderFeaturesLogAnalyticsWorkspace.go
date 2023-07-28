package provider


type AzurermProviderFeaturesLogAnalyticsWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs#permanently_delete_on_destroy AzurermProvider#permanently_delete_on_destroy}.
	PermanentlyDeleteOnDestroy interface{} `field:"optional" json:"permanentlyDeleteOnDestroy" yaml:"permanentlyDeleteOnDestroy"`
}

