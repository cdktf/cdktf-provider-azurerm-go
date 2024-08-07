// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpointcustomdomain


type CdnEndpointCustomDomainUserManagedHttps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/cdn_endpoint_custom_domain#key_vault_certificate_id CdnEndpointCustomDomain#key_vault_certificate_id}.
	KeyVaultCertificateId *string `field:"optional" json:"keyVaultCertificateId" yaml:"keyVaultCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/cdn_endpoint_custom_domain#key_vault_secret_id CdnEndpointCustomDomain#key_vault_secret_id}.
	KeyVaultSecretId *string `field:"optional" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/cdn_endpoint_custom_domain#tls_version CdnEndpointCustomDomain#tls_version}.
	TlsVersion *string `field:"optional" json:"tlsVersion" yaml:"tlsVersion"`
}

