// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFilterOverlay struct {
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/media_transform#audio MediaTransform#audio}
	Audio *MediaTransformOutputCustomPresetFilterOverlayAudio `field:"optional" json:"audio" yaml:"audio"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/media_transform#video MediaTransform#video}
	Video *MediaTransformOutputCustomPresetFilterOverlayVideo `field:"optional" json:"video" yaml:"video"`
}

