// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyX509CertificateProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/key_vault_certificate#key_usage KeyVaultCertificate#key_usage}.
	KeyUsage *[]*string `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/key_vault_certificate#subject KeyVaultCertificate#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/key_vault_certificate#validity_in_months KeyVaultCertificate#validity_in_months}.
	ValidityInMonths *float64 `field:"required" json:"validityInMonths" yaml:"validityInMonths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/key_vault_certificate#extended_key_usage KeyVaultCertificate#extended_key_usage}.
	ExtendedKeyUsage *[]*string `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// subject_alternative_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/key_vault_certificate#subject_alternative_names KeyVaultCertificate#subject_alternative_names}
	SubjectAlternativeNames *KeyVaultCertificateCertificatePolicyX509CertificatePropertiesSubjectAlternativeNames `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

