// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificatecontacts


type KeyVaultCertificateContactsContact struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/key_vault_certificate_contacts#email KeyVaultCertificateContacts#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/key_vault_certificate_contacts#name KeyVaultCertificateContacts#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/key_vault_certificate_contacts#phone KeyVaultCertificateContacts#phone}.
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
}

