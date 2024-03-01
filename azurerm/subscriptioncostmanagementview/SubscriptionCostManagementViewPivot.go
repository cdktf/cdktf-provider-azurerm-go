// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptioncostmanagementview


type SubscriptionCostManagementViewPivot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/subscription_cost_management_view#name SubscriptionCostManagementView#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/subscription_cost_management_view#type SubscriptionCostManagementView#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

