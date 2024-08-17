// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementview


type ResourceGroupCostManagementViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/resource_group_cost_management_view#create ResourceGroupCostManagementView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/resource_group_cost_management_view#delete ResourceGroupCostManagementView#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/resource_group_cost_management_view#read ResourceGroupCostManagementView#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/resource_group_cost_management_view#update ResourceGroupCostManagementView#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

