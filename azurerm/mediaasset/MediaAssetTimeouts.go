// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaasset


type MediaAssetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/media_asset#create MediaAsset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/media_asset#delete MediaAsset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/media_asset#read MediaAsset#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/media_asset#update MediaAsset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

