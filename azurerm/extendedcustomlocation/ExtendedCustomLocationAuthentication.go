// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package extendedcustomlocation


type ExtendedCustomLocationAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/extended_custom_location#value ExtendedCustomLocation#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/extended_custom_location#type ExtendedCustomLocation#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

