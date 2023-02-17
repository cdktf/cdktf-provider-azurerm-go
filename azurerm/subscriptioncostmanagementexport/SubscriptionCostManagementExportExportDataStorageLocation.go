package subscriptioncostmanagementexport


type SubscriptionCostManagementExportExportDataStorageLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_export#container_id SubscriptionCostManagementExport#container_id}.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_export#root_folder_path SubscriptionCostManagementExport#root_folder_path}.
	RootFolderPath *string `field:"required" json:"rootFolderPath" yaml:"rootFolderPath"`
}

