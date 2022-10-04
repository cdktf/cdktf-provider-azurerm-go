package provider


type AzurermProviderFeaturesTemplateDeployment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#delete_nested_items_during_deletion AzurermProvider#delete_nested_items_during_deletion}.
	DeleteNestedItemsDuringDeletion interface{} `field:"required" json:"deleteNestedItemsDuringDeletion" yaml:"deleteNestedItemsDuringDeletion"`
}

