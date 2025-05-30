// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificatePolicy struct {
	// issuer_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/key_vault_certificate#issuer_parameters KeyVaultCertificate#issuer_parameters}
	IssuerParameters *KeyVaultCertificateCertificatePolicyIssuerParameters `field:"required" json:"issuerParameters" yaml:"issuerParameters"`
	// key_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/key_vault_certificate#key_properties KeyVaultCertificate#key_properties}
	KeyProperties *KeyVaultCertificateCertificatePolicyKeyProperties `field:"required" json:"keyProperties" yaml:"keyProperties"`
	// secret_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/key_vault_certificate#secret_properties KeyVaultCertificate#secret_properties}
	SecretProperties *KeyVaultCertificateCertificatePolicySecretProperties `field:"required" json:"secretProperties" yaml:"secretProperties"`
	// lifetime_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/key_vault_certificate#lifetime_action KeyVaultCertificate#lifetime_action}
	LifetimeAction interface{} `field:"optional" json:"lifetimeAction" yaml:"lifetimeAction"`
	// x509_certificate_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/key_vault_certificate#x509_certificate_properties KeyVaultCertificate#x509_certificate_properties}
	X509CertificateProperties *KeyVaultCertificateCertificatePolicyX509CertificateProperties `field:"optional" json:"x509CertificateProperties" yaml:"x509CertificateProperties"`
}

