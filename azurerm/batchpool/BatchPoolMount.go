// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolMount struct {
	// azure_blob_file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/batch_pool#azure_blob_file_system BatchPool#azure_blob_file_system}
	AzureBlobFileSystem *BatchPoolMountAzureBlobFileSystem `field:"optional" json:"azureBlobFileSystem" yaml:"azureBlobFileSystem"`
	// azure_file_share block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/batch_pool#azure_file_share BatchPool#azure_file_share}
	AzureFileShare interface{} `field:"optional" json:"azureFileShare" yaml:"azureFileShare"`
	// cifs_mount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/batch_pool#cifs_mount BatchPool#cifs_mount}
	CifsMount interface{} `field:"optional" json:"cifsMount" yaml:"cifsMount"`
	// nfs_mount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/batch_pool#nfs_mount BatchPool#nfs_mount}
	NfsMount interface{} `field:"optional" json:"nfsMount" yaml:"nfsMount"`
}

