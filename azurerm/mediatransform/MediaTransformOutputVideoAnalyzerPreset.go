// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputVideoAnalyzerPreset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#audio_analysis_mode MediaTransform#audio_analysis_mode}.
	AudioAnalysisMode *string `field:"optional" json:"audioAnalysisMode" yaml:"audioAnalysisMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#audio_language MediaTransform#audio_language}.
	AudioLanguage *string `field:"optional" json:"audioLanguage" yaml:"audioLanguage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#experimental_options MediaTransform#experimental_options}.
	ExperimentalOptions *map[string]*string `field:"optional" json:"experimentalOptions" yaml:"experimentalOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_transform#insights_type MediaTransform#insights_type}.
	InsightsType *string `field:"optional" json:"insightsType" yaml:"insightsType"`
}

