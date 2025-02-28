// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcareservice


type HealthcareServiceCorsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/healthcare_service#allow_credentials HealthcareService#allow_credentials}.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/healthcare_service#allowed_headers HealthcareService#allowed_headers}.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/healthcare_service#allowed_methods HealthcareService#allowed_methods}.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/healthcare_service#allowed_origins HealthcareService#allowed_origins}.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/healthcare_service#max_age_in_seconds HealthcareService#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"optional" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

