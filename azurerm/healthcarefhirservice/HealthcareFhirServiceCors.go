// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirservice


type HealthcareFhirServiceCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_fhir_service#allowed_headers HealthcareFhirService#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_fhir_service#allowed_methods HealthcareFhirService#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_fhir_service#allowed_origins HealthcareFhirService#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_fhir_service#credentials_allowed HealthcareFhirService#credentials_allowed}.
	CredentialsAllowed interface{} `field:"optional" json:"credentialsAllowed" yaml:"credentialsAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/healthcare_fhir_service#max_age_in_seconds HealthcareFhirService#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"optional" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}

