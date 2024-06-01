// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificateissuer


type KeyVaultCertificateIssuerAdmin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/key_vault_certificate_issuer#email_address KeyVaultCertificateIssuer#email_address}.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/key_vault_certificate_issuer#first_name KeyVaultCertificateIssuer#first_name}.
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/key_vault_certificate_issuer#last_name KeyVaultCertificateIssuer#last_name}.
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.106.1/docs/resources/key_vault_certificate_issuer#phone KeyVaultCertificateIssuer#phone}.
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
}

