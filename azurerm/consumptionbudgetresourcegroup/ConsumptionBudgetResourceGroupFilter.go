// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetresourcegroup


type ConsumptionBudgetResourceGroupFilter struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/consumption_budget_resource_group#dimension ConsumptionBudgetResourceGroup#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/consumption_budget_resource_group#tag ConsumptionBudgetResourceGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

