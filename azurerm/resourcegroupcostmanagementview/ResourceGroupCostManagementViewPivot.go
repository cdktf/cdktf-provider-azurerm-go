package resourcegroupcostmanagementview


type ResourceGroupCostManagementViewPivot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_cost_management_view#name ResourceGroupCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_cost_management_view#type ResourceGroupCostManagementView#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

