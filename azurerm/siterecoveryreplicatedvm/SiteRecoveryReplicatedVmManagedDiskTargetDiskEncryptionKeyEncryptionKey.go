// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/site_recovery_replicated_vm#key_url SiteRecoveryReplicatedVm#key_url}.
	KeyUrl *string `field:"optional" json:"keyUrl" yaml:"keyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/site_recovery_replicated_vm#vault_id SiteRecoveryReplicatedVm#vault_id}.
	VaultId *string `field:"optional" json:"vaultId" yaml:"vaultId"`
}

