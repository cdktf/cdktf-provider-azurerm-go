// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimage


type SharedImageIdentifier struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/shared_image#offer SharedImage#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/shared_image#publisher SharedImage#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/shared_image#sku SharedImage#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}

