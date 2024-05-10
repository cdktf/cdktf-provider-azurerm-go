// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetCodecH265Video struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#complexity MediaTransform#complexity}.
	Complexity *string `field:"optional" json:"complexity" yaml:"complexity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#key_frame_interval MediaTransform#key_frame_interval}.
	KeyFrameInterval *string `field:"optional" json:"keyFrameInterval" yaml:"keyFrameInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#label MediaTransform#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// layer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#layer MediaTransform#layer}
	Layer interface{} `field:"optional" json:"layer" yaml:"layer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#scene_change_detection_enabled MediaTransform#scene_change_detection_enabled}.
	SceneChangeDetectionEnabled interface{} `field:"optional" json:"sceneChangeDetectionEnabled" yaml:"sceneChangeDetectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#stretch_mode MediaTransform#stretch_mode}.
	StretchMode *string `field:"optional" json:"stretchMode" yaml:"stretchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#sync_mode MediaTransform#sync_mode}.
	SyncMode *string `field:"optional" json:"syncMode" yaml:"syncMode"`
}

