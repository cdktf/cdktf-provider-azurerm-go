// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycentersubscriptionpricing


type SecurityCenterSubscriptionPricingExtension struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#name SecurityCenterSubscriptionPricing#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/security_center_subscription_pricing#additional_extension_properties SecurityCenterSubscriptionPricing#additional_extension_properties}.
	AdditionalExtensionProperties *map[string]*string `field:"optional" json:"additionalExtensionProperties" yaml:"additionalExtensionProperties"`
}

