// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccacheblobnfstarget


type HpcCacheBlobNfsTargetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hpc_cache_blob_nfs_target#create HpcCacheBlobNfsTarget#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hpc_cache_blob_nfs_target#delete HpcCacheBlobNfsTarget#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hpc_cache_blob_nfs_target#read HpcCacheBlobNfsTarget#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/hpc_cache_blob_nfs_target#update HpcCacheBlobNfsTarget#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

