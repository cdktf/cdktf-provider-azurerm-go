// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesRecoveryService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs#purge_protected_items_from_vault_on_destroy AzurermProvider#purge_protected_items_from_vault_on_destroy}.
	PurgeProtectedItemsFromVaultOnDestroy interface{} `field:"optional" json:"purgeProtectedItemsFromVaultOnDestroy" yaml:"purgeProtectedItemsFromVaultOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs#vm_backup_stop_protection_and_retain_data_on_destroy AzurermProvider#vm_backup_stop_protection_and_retain_data_on_destroy}.
	VmBackupStopProtectionAndRetainDataOnDestroy interface{} `field:"optional" json:"vmBackupStopProtectionAndRetainDataOnDestroy" yaml:"vmBackupStopProtectionAndRetainDataOnDestroy"`
}

