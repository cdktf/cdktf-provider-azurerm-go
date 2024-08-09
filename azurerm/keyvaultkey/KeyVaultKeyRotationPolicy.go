// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultkey


type KeyVaultKeyRotationPolicy struct {
	// automatic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/key_vault_key#automatic KeyVaultKey#automatic}
	Automatic *KeyVaultKeyRotationPolicyAutomatic `field:"optional" json:"automatic" yaml:"automatic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/key_vault_key#expire_after KeyVaultKey#expire_after}.
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/key_vault_key#notify_before_expiry KeyVaultKey#notify_before_expiry}.
	NotifyBeforeExpiry *string `field:"optional" json:"notifyBeforeExpiry" yaml:"notifyBeforeExpiry"`
}

