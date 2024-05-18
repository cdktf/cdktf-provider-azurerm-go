// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputCustomPresetCodecAacAudio struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_transform#bitrate MediaTransform#bitrate}.
	Bitrate *float64 `field:"optional" json:"bitrate" yaml:"bitrate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_transform#channels MediaTransform#channels}.
	Channels *float64 `field:"optional" json:"channels" yaml:"channels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_transform#label MediaTransform#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_transform#profile MediaTransform#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/media_transform#sampling_rate MediaTransform#sampling_rate}.
	SamplingRate *float64 `field:"optional" json:"samplingRate" yaml:"samplingRate"`
}

