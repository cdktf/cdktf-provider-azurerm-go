// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package consumptionbudgetsubscription


type ConsumptionBudgetSubscriptionFilter struct {
	// dimension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/consumption_budget_subscription#dimension ConsumptionBudgetSubscription#dimension}
	Dimension interface{} `field:"optional" json:"dimension" yaml:"dimension"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/consumption_budget_subscription#tag ConsumptionBudgetSubscription#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

