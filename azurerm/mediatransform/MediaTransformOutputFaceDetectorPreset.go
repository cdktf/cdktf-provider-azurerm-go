// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediatransform


type MediaTransformOutputFaceDetectorPreset struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/media_transform#analysis_resolution MediaTransform#analysis_resolution}.
	AnalysisResolution *string `field:"optional" json:"analysisResolution" yaml:"analysisResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/media_transform#blur_type MediaTransform#blur_type}.
	BlurType *string `field:"optional" json:"blurType" yaml:"blurType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/media_transform#experimental_options MediaTransform#experimental_options}.
	ExperimentalOptions *map[string]*string `field:"optional" json:"experimentalOptions" yaml:"experimentalOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/media_transform#face_redactor_mode MediaTransform#face_redactor_mode}.
	FaceRedactorMode *string `field:"optional" json:"faceRedactorMode" yaml:"faceRedactorMode"`
}

