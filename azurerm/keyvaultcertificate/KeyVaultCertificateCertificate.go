// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/key_vault_certificate#contents KeyVaultCertificate#contents}.
	Contents *string `field:"required" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/key_vault_certificate#password KeyVaultCertificate#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
}

