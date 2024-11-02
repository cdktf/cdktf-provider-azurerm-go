// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyassignment


type SubscriptionPolicyAssignmentResourceSelectorsSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/subscription_policy_assignment#kind SubscriptionPolicyAssignment#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/subscription_policy_assignment#in SubscriptionPolicyAssignment#in}.
	In *[]*string `field:"optional" json:"in" yaml:"in"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/subscription_policy_assignment#not_in SubscriptionPolicyAssignment#not_in}.
	NotIn *[]*string `field:"optional" json:"notIn" yaml:"notIn"`
}

