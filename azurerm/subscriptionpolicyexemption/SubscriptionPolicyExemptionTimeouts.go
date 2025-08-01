// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyexemption


type SubscriptionPolicyExemptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/subscription_policy_exemption#create SubscriptionPolicyExemption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/subscription_policy_exemption#delete SubscriptionPolicyExemption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/subscription_policy_exemption#read SubscriptionPolicyExemption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/subscription_policy_exemption#update SubscriptionPolicyExemption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

