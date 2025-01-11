// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcaremedtechservice


type HealthcareMedtechServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/healthcare_medtech_service#create HealthcareMedtechService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/healthcare_medtech_service#delete HealthcareMedtechService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/healthcare_medtech_service#read HealthcareMedtechService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/healthcare_medtech_service#update HealthcareMedtechService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

