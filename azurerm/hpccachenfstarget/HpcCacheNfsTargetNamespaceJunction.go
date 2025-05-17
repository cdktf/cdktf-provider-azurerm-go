// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccachenfstarget


type HpcCacheNfsTargetNamespaceJunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/hpc_cache_nfs_target#namespace_path HpcCacheNfsTarget#namespace_path}.
	NamespacePath *string `field:"required" json:"namespacePath" yaml:"namespacePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/hpc_cache_nfs_target#nfs_export HpcCacheNfsTarget#nfs_export}.
	NfsExport *string `field:"required" json:"nfsExport" yaml:"nfsExport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/hpc_cache_nfs_target#access_policy_name HpcCacheNfsTarget#access_policy_name}.
	AccessPolicyName *string `field:"optional" json:"accessPolicyName" yaml:"accessPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/hpc_cache_nfs_target#target_path HpcCacheNfsTarget#target_path}.
	TargetPath *string `field:"optional" json:"targetPath" yaml:"targetPath"`
}

