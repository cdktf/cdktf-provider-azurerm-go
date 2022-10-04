package provider


type AzurermProviderFeaturesLogAnalyticsWorkspace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#permanently_delete_on_destroy AzurermProvider#permanently_delete_on_destroy}.
	PermanentlyDeleteOnDestroy interface{} `field:"optional" json:"permanentlyDeleteOnDestroy" yaml:"permanentlyDeleteOnDestroy"`
}

