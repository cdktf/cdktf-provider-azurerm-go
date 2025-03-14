// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountlocaluser


type StorageAccountLocalUserPermissionScope struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/storage_account_local_user#permissions StorageAccountLocalUser#permissions}
	Permissions *StorageAccountLocalUserPermissionScopePermissions `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/storage_account_local_user#resource_name StorageAccountLocalUser#resource_name}.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/storage_account_local_user#service StorageAccountLocalUser#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

