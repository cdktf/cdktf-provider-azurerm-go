package subscriptioncostmanagementview


type SubscriptionCostManagementViewPivot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_view#name SubscriptionCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_view#type SubscriptionCostManagementView#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

