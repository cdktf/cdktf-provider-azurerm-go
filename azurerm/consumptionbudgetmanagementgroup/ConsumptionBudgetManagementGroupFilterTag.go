// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetmanagementgroup


type ConsumptionBudgetManagementGroupFilterTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/consumption_budget_management_group#name ConsumptionBudgetManagementGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/consumption_budget_management_group#values ConsumptionBudgetManagementGroup#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/consumption_budget_management_group#operator ConsumptionBudgetManagementGroup#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

