// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFilterDeinterlace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/media_transform#mode MediaTransform#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/media_transform#parity MediaTransform#parity}.
	Parity *string `field:"optional" json:"parity" yaml:"parity"`
}

