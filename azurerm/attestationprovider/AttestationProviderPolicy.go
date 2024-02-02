// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package attestationprovider


type AttestationProviderPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/attestation_provider#data AttestationProvider#data}.
	Data *string `field:"optional" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/attestation_provider#environment_type AttestationProvider#environment_type}.
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
}

