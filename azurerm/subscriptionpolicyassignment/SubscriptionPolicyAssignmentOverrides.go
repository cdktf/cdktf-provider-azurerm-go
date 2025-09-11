// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_policy_assignment#value SubscriptionPolicyAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/subscription_policy_assignment#selectors SubscriptionPolicyAssignment#selectors}
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}

