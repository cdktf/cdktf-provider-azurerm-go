// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetCodec struct {
	// aac_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#aac_audio MediaTransform#aac_audio}
	AacAudio *MediaTransformOutputCustomPresetCodecAacAudio `field:"optional" json:"aacAudio" yaml:"aacAudio"`
	// copy_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#copy_audio MediaTransform#copy_audio}
	CopyAudio *MediaTransformOutputCustomPresetCodecCopyAudio `field:"optional" json:"copyAudio" yaml:"copyAudio"`
	// copy_video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#copy_video MediaTransform#copy_video}
	CopyVideo *MediaTransformOutputCustomPresetCodecCopyVideo `field:"optional" json:"copyVideo" yaml:"copyVideo"`
	// dd_audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#dd_audio MediaTransform#dd_audio}
	DdAudio *MediaTransformOutputCustomPresetCodecDdAudio `field:"optional" json:"ddAudio" yaml:"ddAudio"`
	// h264_video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#h264_video MediaTransform#h264_video}
	H264Video *MediaTransformOutputCustomPresetCodecH264Video `field:"optional" json:"h264Video" yaml:"h264Video"`
	// h265_video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#h265_video MediaTransform#h265_video}
	H265Video *MediaTransformOutputCustomPresetCodecH265Video `field:"optional" json:"h265Video" yaml:"h265Video"`
	// jpg_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#jpg_image MediaTransform#jpg_image}
	JpgImage *MediaTransformOutputCustomPresetCodecJpgImage `field:"optional" json:"jpgImage" yaml:"jpgImage"`
	// png_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/media_transform#png_image MediaTransform#png_image}
	PngImage *MediaTransformOutputCustomPresetCodecPngImage `field:"optional" json:"pngImage" yaml:"pngImage"`
}

