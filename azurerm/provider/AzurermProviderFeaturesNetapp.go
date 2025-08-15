// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesNetapp struct {
	// When enabled, backups will be deleted when the `azurerm_netapp_backup_vault` resource is destroyed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs#delete_backups_on_backup_vault_destroy AzurermProvider#delete_backups_on_backup_vault_destroy}
	DeleteBackupsOnBackupVaultDestroy interface{} `field:"optional" json:"deleteBackupsOnBackupVaultDestroy" yaml:"deleteBackupsOnBackupVaultDestroy"`
	// When enabled, the volume will not be destroyed, safeguarding from severe data loss.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs#prevent_volume_destruction AzurermProvider#prevent_volume_destruction}
	PreventVolumeDestruction interface{} `field:"optional" json:"preventVolumeDestruction" yaml:"preventVolumeDestruction"`
}

