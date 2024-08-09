// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetFormat struct {
	// jpg block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#jpg MediaTransform#jpg}
	Jpg *MediaTransformOutputCustomPresetFormatJpg `field:"optional" json:"jpg" yaml:"jpg"`
	// mp4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#mp4 MediaTransform#mp4}
	Mp4 *MediaTransformOutputCustomPresetFormatMp4 `field:"optional" json:"mp4" yaml:"mp4"`
	// png block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#png MediaTransform#png}
	Png *MediaTransformOutputCustomPresetFormatPng `field:"optional" json:"png" yaml:"png"`
	// transport_stream block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#transport_stream MediaTransform#transport_stream}
	TransportStream *MediaTransformOutputCustomPresetFormatTransportStream `field:"optional" json:"transportStream" yaml:"transportStream"`
}

