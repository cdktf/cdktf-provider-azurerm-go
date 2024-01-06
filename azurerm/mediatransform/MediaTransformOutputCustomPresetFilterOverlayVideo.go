// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFilterOverlayVideo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#input_label MediaTransform#input_label}.
	InputLabel *string `field:"required" json:"inputLabel" yaml:"inputLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#audio_gain_level MediaTransform#audio_gain_level}.
	AudioGainLevel *float64 `field:"optional" json:"audioGainLevel" yaml:"audioGainLevel"`
	// crop_rectangle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#crop_rectangle MediaTransform#crop_rectangle}
	CropRectangle *MediaTransformOutputCustomPresetFilterOverlayVideoCropRectangle `field:"optional" json:"cropRectangle" yaml:"cropRectangle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#end MediaTransform#end}.
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#fade_in_duration MediaTransform#fade_in_duration}.
	FadeInDuration *string `field:"optional" json:"fadeInDuration" yaml:"fadeInDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#fade_out_duration MediaTransform#fade_out_duration}.
	FadeOutDuration *string `field:"optional" json:"fadeOutDuration" yaml:"fadeOutDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#opacity MediaTransform#opacity}.
	Opacity *float64 `field:"optional" json:"opacity" yaml:"opacity"`
	// position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#position MediaTransform#position}
	Position *MediaTransformOutputCustomPresetFilterOverlayVideoPosition `field:"optional" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/media_transform#start MediaTransform#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}

