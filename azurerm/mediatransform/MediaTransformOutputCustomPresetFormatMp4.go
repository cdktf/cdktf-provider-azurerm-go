// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFormatMp4 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#filename_pattern MediaTransform#filename_pattern}.
	FilenamePattern *string `field:"required" json:"filenamePattern" yaml:"filenamePattern"`
	// output_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/media_transform#output_file MediaTransform#output_file}
	OutputFile interface{} `field:"optional" json:"outputFile" yaml:"outputFile"`
}

