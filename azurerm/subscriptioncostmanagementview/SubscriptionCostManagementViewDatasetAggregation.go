package subscriptioncostmanagementview


type SubscriptionCostManagementViewDatasetAggregation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#column_name SubscriptionCostManagementView#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/subscription_cost_management_view#name SubscriptionCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

