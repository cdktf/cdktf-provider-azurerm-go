// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilter struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/consumption_budget_management_group#dimension ConsumptionBudgetManagementGroup#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/consumption_budget_management_group#not ConsumptionBudgetManagementGroup#not}
	Not *ConsumptionBudgetManagementGroupFilterNot `field:"optional" json:"not" yaml:"not"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/consumption_budget_management_group#tag ConsumptionBudgetManagementGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

