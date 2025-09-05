// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementview


type ResourceGroupCostManagementViewDatasetAggregation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/resource_group_cost_management_view#column_name ResourceGroupCostManagementView#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/resource_group_cost_management_view#name ResourceGroupCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

