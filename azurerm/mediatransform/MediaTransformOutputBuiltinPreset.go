// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputBuiltinPreset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/media_transform#preset_name MediaTransform#preset_name}.
	PresetName *string `field:"required" json:"presetName" yaml:"presetName"`
	// preset_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/media_transform#preset_configuration MediaTransform#preset_configuration}
	PresetConfiguration *MediaTransformOutputBuiltinPresetPresetConfiguration `field:"optional" json:"presetConfiguration" yaml:"presetConfiguration"`
}

