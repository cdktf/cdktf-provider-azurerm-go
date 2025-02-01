// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_backend_ssl30 ApiManagement#enable_backend_ssl30}.
	EnableBackendSsl30 interface{} `field:"optional" json:"enableBackendSsl30" yaml:"enableBackendSsl30"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_backend_tls10 ApiManagement#enable_backend_tls10}.
	EnableBackendTls10 interface{} `field:"optional" json:"enableBackendTls10" yaml:"enableBackendTls10"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_backend_tls11 ApiManagement#enable_backend_tls11}.
	EnableBackendTls11 interface{} `field:"optional" json:"enableBackendTls11" yaml:"enableBackendTls11"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_frontend_ssl30 ApiManagement#enable_frontend_ssl30}.
	EnableFrontendSsl30 interface{} `field:"optional" json:"enableFrontendSsl30" yaml:"enableFrontendSsl30"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_frontend_tls10 ApiManagement#enable_frontend_tls10}.
	EnableFrontendTls10 interface{} `field:"optional" json:"enableFrontendTls10" yaml:"enableFrontendTls10"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#enable_frontend_tls11 ApiManagement#enable_frontend_tls11}.
	EnableFrontendTls11 interface{} `field:"optional" json:"enableFrontendTls11" yaml:"enableFrontendTls11"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsEcdheEcdsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsEcdheEcdsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheEcdsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsEcdheRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsEcdheRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsEcdheRsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha256_ciphers_enabled}.
	TlsRsaWithAes128CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes128CbcSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes128_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes128_cbc_sha_ciphers_enabled}.
	TlsRsaWithAes128CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128CbcShaCiphersEnabled" yaml:"tlsRsaWithAes128CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes128_gcm_sha256_ciphers_enabled}.
	TlsRsaWithAes128GcmSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes128GcmSha256CiphersEnabled" yaml:"tlsRsaWithAes128GcmSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha256_ciphers_enabled}.
	TlsRsaWithAes256CbcSha256CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcSha256CiphersEnabled" yaml:"tlsRsaWithAes256CbcSha256CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes256_cbc_sha_ciphers_enabled ApiManagement#tls_rsa_with_aes256_cbc_sha_ciphers_enabled}.
	TlsRsaWithAes256CbcShaCiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256CbcShaCiphersEnabled" yaml:"tlsRsaWithAes256CbcShaCiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#tls_rsa_with_aes256_gcm_sha384_ciphers_enabled ApiManagement#tls_rsa_with_aes256_gcm_sha384_ciphers_enabled}.
	TlsRsaWithAes256GcmSha384CiphersEnabled interface{} `field:"optional" json:"tlsRsaWithAes256GcmSha384CiphersEnabled" yaml:"tlsRsaWithAes256GcmSha384CiphersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/api_management#triple_des_ciphers_enabled ApiManagement#triple_des_ciphers_enabled}.
	TripleDesCiphersEnabled interface{} `field:"optional" json:"tripleDesCiphersEnabled" yaml:"tripleDesCiphersEnabled"`
}

