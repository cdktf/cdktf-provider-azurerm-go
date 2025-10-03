// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirservice


type HealthcareFhirServiceOciArtifact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/healthcare_fhir_service#login_server HealthcareFhirService#login_server}.
	LoginServer *string `field:"required" json:"loginServer" yaml:"loginServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/healthcare_fhir_service#digest HealthcareFhirService#digest}.
	Digest *string `field:"optional" json:"digest" yaml:"digest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/healthcare_fhir_service#image_name HealthcareFhirService#image_name}.
	ImageName *string `field:"optional" json:"imageName" yaml:"imageName"`
}

