// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFilter struct {
	// crop_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#crop_rectangle MediaTransform#crop_rectangle}
	CropRectangle *MediaTransformOutputCustomPresetFilterCropRectangle `field:"optional" json:"cropRectangle" yaml:"cropRectangle"`
	// deinterlace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#deinterlace MediaTransform#deinterlace}
	Deinterlace *MediaTransformOutputCustomPresetFilterDeinterlace `field:"optional" json:"deinterlace" yaml:"deinterlace"`
	// fade_in block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#fade_in MediaTransform#fade_in}
	FadeIn *MediaTransformOutputCustomPresetFilterFadeIn `field:"optional" json:"fadeIn" yaml:"fadeIn"`
	// fade_out block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#fade_out MediaTransform#fade_out}
	FadeOut *MediaTransformOutputCustomPresetFilterFadeOut `field:"optional" json:"fadeOut" yaml:"fadeOut"`
	// overlay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#overlay MediaTransform#overlay}
	Overlay interface{} `field:"optional" json:"overlay" yaml:"overlay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/media_transform#rotation MediaTransform#rotation}.
	Rotation *string `field:"optional" json:"rotation" yaml:"rotation"`
}

