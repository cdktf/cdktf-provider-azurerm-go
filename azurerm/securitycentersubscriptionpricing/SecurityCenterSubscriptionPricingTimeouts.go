// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycentersubscriptionpricing


type SecurityCenterSubscriptionPricingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#create SecurityCenterSubscriptionPricing#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#delete SecurityCenterSubscriptionPricing#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#read SecurityCenterSubscriptionPricing#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#update SecurityCenterSubscriptionPricing#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

