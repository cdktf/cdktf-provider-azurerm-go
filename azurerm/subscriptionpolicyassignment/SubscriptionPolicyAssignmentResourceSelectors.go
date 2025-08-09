// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentResourceSelectors struct {
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/subscription_policy_assignment#selectors SubscriptionPolicyAssignment#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/subscription_policy_assignment#name SubscriptionPolicyAssignment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

