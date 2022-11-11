package provider


type AzurermProviderFeaturesManagedDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#expand_without_downtime AzurermProvider#expand_without_downtime}.
	ExpandWithoutDowntime interface{} `field:"optional" json:"expandWithoutDowntime" yaml:"expandWithoutDowntime"`
}

