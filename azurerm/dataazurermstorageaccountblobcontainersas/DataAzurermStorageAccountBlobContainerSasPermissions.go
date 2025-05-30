// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstorageaccountblobcontainersas


type DataAzurermStorageAccountBlobContainerSasPermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#add DataAzurermStorageAccountBlobContainerSas#add}.
	Add interface{} `field:"required" json:"add" yaml:"add"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#create DataAzurermStorageAccountBlobContainerSas#create}.
	Create interface{} `field:"required" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#delete DataAzurermStorageAccountBlobContainerSas#delete}.
	Delete interface{} `field:"required" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#list DataAzurermStorageAccountBlobContainerSas#list}.
	List interface{} `field:"required" json:"list" yaml:"list"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#read DataAzurermStorageAccountBlobContainerSas#read}.
	Read interface{} `field:"required" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/data-sources/storage_account_blob_container_sas#write DataAzurermStorageAccountBlobContainerSas#write}.
	Write interface{} `field:"required" json:"write" yaml:"write"`
}

