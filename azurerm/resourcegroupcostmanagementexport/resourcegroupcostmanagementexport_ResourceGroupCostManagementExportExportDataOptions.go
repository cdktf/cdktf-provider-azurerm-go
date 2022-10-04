package resourcegroupcostmanagementexport


type ResourceGroupCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_cost_management_export#time_frame ResourceGroupCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/resource_group_cost_management_export#type ResourceGroupCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

