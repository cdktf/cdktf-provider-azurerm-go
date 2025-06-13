// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilter struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/consumption_budget_management_group#dimension ConsumptionBudgetManagementGroup#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/consumption_budget_management_group#tag ConsumptionBudgetManagementGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

