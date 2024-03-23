// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptionpolicyremediation


type SubscriptionPolicyRemediationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/subscription_policy_remediation#create SubscriptionPolicyRemediation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/subscription_policy_remediation#delete SubscriptionPolicyRemediation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/subscription_policy_remediation#read SubscriptionPolicyRemediation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/subscription_policy_remediation#update SubscriptionPolicyRemediation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

