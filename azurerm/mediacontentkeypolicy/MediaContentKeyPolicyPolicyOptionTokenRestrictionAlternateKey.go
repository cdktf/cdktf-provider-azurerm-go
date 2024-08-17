// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediacontentkeypolicy


type MediaContentKeyPolicyPolicyOptionTokenRestrictionAlternateKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_content_key_policy#rsa_token_key_exponent MediaContentKeyPolicy#rsa_token_key_exponent}.
	RsaTokenKeyExponent *string `field:"optional" json:"rsaTokenKeyExponent" yaml:"rsaTokenKeyExponent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_content_key_policy#rsa_token_key_modulus MediaContentKeyPolicy#rsa_token_key_modulus}.
	RsaTokenKeyModulus *string `field:"optional" json:"rsaTokenKeyModulus" yaml:"rsaTokenKeyModulus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_content_key_policy#symmetric_token_key MediaContentKeyPolicy#symmetric_token_key}.
	SymmetricTokenKey *string `field:"optional" json:"symmetricTokenKey" yaml:"symmetricTokenKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/media_content_key_policy#x509_token_key_raw MediaContentKeyPolicy#x509_token_key_raw}.
	X509TokenKeyRaw *string `field:"optional" json:"x509TokenKeyRaw" yaml:"x509TokenKeyRaw"`
}

