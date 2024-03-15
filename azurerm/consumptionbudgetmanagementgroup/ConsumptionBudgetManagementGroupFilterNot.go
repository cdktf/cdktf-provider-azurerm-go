// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilterNot struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/consumption_budget_management_group#dimension ConsumptionBudgetManagementGroup#dimension}
	Dimension *ConsumptionBudgetManagementGroupFilterNotDimension `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/consumption_budget_management_group#tag ConsumptionBudgetManagementGroup#tag}
	Tag *ConsumptionBudgetManagementGroupFilterNotTag `field:"optional" json:"tag" yaml:"tag"`
}

