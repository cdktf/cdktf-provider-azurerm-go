// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcareservice


type HealthcareServiceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/healthcare_service#type HealthcareService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

