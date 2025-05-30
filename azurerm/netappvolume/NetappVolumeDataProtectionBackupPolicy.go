// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeDataProtectionBackupPolicy struct {
	// The ID of the backup policy to associate with this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/netapp_volume#backup_policy_id NetappVolume#backup_policy_id}
	BackupPolicyId *string `field:"required" json:"backupPolicyId" yaml:"backupPolicyId"`
	// The ID of the backup vault to associate with this volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/netapp_volume#backup_vault_id NetappVolume#backup_vault_id}
	BackupVaultId *string `field:"required" json:"backupVaultId" yaml:"backupVaultId"`
	// If set to false, the backup policy will not be enabled on this volume, thus disabling scheduled backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/netapp_volume#policy_enabled NetappVolume#policy_enabled}
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
}

