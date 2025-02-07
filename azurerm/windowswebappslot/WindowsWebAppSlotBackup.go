// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app_slot#name WindowsWebAppSlot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app_slot#schedule WindowsWebAppSlot#schedule}
	Schedule *WindowsWebAppSlotBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app_slot#storage_account_url WindowsWebAppSlot#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/windows_web_app_slot#enabled WindowsWebAppSlot#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

