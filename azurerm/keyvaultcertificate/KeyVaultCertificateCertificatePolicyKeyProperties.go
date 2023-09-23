// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultcertificate


type KeyVaultCertificateCertificatePolicyKeyProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/key_vault_certificate#exportable KeyVaultCertificate#exportable}.
	Exportable interface{} `field:"required" json:"exportable" yaml:"exportable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/key_vault_certificate#key_type KeyVaultCertificate#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/key_vault_certificate#reuse_key KeyVaultCertificate#reuse_key}.
	ReuseKey interface{} `field:"required" json:"reuseKey" yaml:"reuseKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/key_vault_certificate#curve KeyVaultCertificate#curve}.
	Curve *string `field:"optional" json:"curve" yaml:"curve"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.74.0/docs/resources/key_vault_certificate#key_size KeyVaultCertificate#key_size}.
	KeySize *float64 `field:"optional" json:"keySize" yaml:"keySize"`
}

