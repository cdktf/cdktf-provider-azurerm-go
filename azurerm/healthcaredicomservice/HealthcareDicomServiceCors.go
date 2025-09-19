// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaredicomservice


type HealthcareDicomServiceCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/healthcare_dicom_service#allow_credentials HealthcareDicomService#allow_credentials}.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/healthcare_dicom_service#allowed_headers HealthcareDicomService#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/healthcare_dicom_service#allowed_methods HealthcareDicomService#allowed_methods}.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/healthcare_dicom_service#allowed_origins HealthcareDicomService#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/healthcare_dicom_service#max_age_in_seconds HealthcareDicomService#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"optional" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

