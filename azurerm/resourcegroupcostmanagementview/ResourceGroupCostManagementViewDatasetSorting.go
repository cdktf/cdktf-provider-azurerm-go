package resourcegroupcostmanagementview


type ResourceGroupCostManagementViewDatasetSorting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/resource_group_cost_management_view#direction ResourceGroupCostManagementView#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/resource_group_cost_management_view#name ResourceGroupCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

