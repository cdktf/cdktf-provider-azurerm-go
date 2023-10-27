// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyLifetimeAction struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/key_vault_certificate#action KeyVaultCertificate#action}
	Action *KeyVaultCertificateCertificatePolicyLifetimeActionAction `field:"required" json:"action" yaml:"action"`
	// trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/key_vault_certificate#trigger KeyVaultCertificate#trigger}
	Trigger *KeyVaultCertificateCertificatePolicyLifetimeActionTrigger `field:"required" json:"trigger" yaml:"trigger"`
}

