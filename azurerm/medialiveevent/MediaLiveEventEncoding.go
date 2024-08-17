// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialiveevent


type MediaLiveEventEncoding struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_live_event#key_frame_interval MediaLiveEvent#key_frame_interval}.
	KeyFrameInterval *string `field:"optional" json:"keyFrameInterval" yaml:"keyFrameInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_live_event#preset_name MediaLiveEvent#preset_name}.
	PresetName *string `field:"optional" json:"presetName" yaml:"presetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_live_event#stretch_mode MediaLiveEvent#stretch_mode}.
	StretchMode *string `field:"optional" json:"stretchMode" yaml:"stretchMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_live_event#type MediaLiveEvent#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

