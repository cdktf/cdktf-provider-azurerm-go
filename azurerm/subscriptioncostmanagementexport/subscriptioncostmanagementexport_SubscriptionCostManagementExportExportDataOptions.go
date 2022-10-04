package subscriptioncostmanagementexport


type SubscriptionCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_export#time_frame SubscriptionCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_cost_management_export#type SubscriptionCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

