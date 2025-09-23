// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountAzureFilesAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/storage_account#directory_type StorageAccount#directory_type}.
	DirectoryType *string `field:"required" json:"directoryType" yaml:"directoryType"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/storage_account#active_directory StorageAccount#active_directory}
	ActiveDirectory *StorageAccountAzureFilesAuthenticationActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/storage_account#default_share_level_permission StorageAccount#default_share_level_permission}.
	DefaultShareLevelPermission *string `field:"optional" json:"defaultShareLevelPermission" yaml:"defaultShareLevelPermission"`
}

