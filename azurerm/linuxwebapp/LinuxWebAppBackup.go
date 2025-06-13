// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/linux_web_app#name LinuxWebApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/linux_web_app#schedule LinuxWebApp#schedule}
	Schedule *LinuxWebAppBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/linux_web_app#storage_account_url LinuxWebApp#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/linux_web_app#enabled LinuxWebApp#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

