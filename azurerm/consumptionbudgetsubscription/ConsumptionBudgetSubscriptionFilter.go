// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilter struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/consumption_budget_subscription#dimension ConsumptionBudgetSubscription#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// not block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/consumption_budget_subscription#not ConsumptionBudgetSubscription#not}
	Not *ConsumptionBudgetSubscriptionFilterNot `field:"optional" json:"not" yaml:"not"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/consumption_budget_subscription#tag ConsumptionBudgetSubscription#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

