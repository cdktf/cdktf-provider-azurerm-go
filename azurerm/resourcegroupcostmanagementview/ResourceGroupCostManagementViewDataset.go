// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementview


type ResourceGroupCostManagementViewDataset struct {
	// aggregation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/resource_group_cost_management_view#aggregation ResourceGroupCostManagementView#aggregation}
	Aggregation interface{} `field:"required" json:"aggregation" yaml:"aggregation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/resource_group_cost_management_view#granularity ResourceGroupCostManagementView#granularity}.
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// grouping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/resource_group_cost_management_view#grouping ResourceGroupCostManagementView#grouping}
	Grouping interface{} `field:"optional" json:"grouping" yaml:"grouping"`
	// sorting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/resource_group_cost_management_view#sorting ResourceGroupCostManagementView#sorting}
	Sorting interface{} `field:"optional" json:"sorting" yaml:"sorting"`
}

