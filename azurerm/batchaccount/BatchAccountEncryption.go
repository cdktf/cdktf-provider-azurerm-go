// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchaccount


type BatchAccountEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/batch_account#key_vault_key_id BatchAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
}

