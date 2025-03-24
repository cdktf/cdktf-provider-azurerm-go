// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementexport


type ResourceGroupCostManagementExportExportDataOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/resource_group_cost_management_export#time_frame ResourceGroupCostManagementExport#time_frame}.
	TimeFrame *string `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/resource_group_cost_management_export#type ResourceGroupCostManagementExport#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

