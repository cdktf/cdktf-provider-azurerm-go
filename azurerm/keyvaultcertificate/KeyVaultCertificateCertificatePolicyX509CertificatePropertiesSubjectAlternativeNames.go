// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/key_vault_certificate#dns_names KeyVaultCertificate#dns_names}.
	DnsNames *[]*string `field:"optional" json:"dnsNames" yaml:"dnsNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/key_vault_certificate#emails KeyVaultCertificate#emails}.
	Emails *[]*string `field:"optional" json:"emails" yaml:"emails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.86.0/docs/resources/key_vault_certificate#upns KeyVaultCertificate#upns}.
	Upns *[]*string `field:"optional" json:"upns" yaml:"upns"`
}

