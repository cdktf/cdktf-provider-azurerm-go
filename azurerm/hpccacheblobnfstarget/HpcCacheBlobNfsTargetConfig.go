// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccacheblobnfstarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheBlobNfsTargetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#cache_name HpcCacheBlobNfsTarget#cache_name}.
	CacheName *string `field:"required" json:"cacheName" yaml:"cacheName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#name HpcCacheBlobNfsTarget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#namespace_path HpcCacheBlobNfsTarget#namespace_path}.
	NamespacePath *string `field:"required" json:"namespacePath" yaml:"namespacePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#resource_group_name HpcCacheBlobNfsTarget#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#storage_container_id HpcCacheBlobNfsTarget#storage_container_id}.
	StorageContainerId *string `field:"required" json:"storageContainerId" yaml:"storageContainerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#usage_model HpcCacheBlobNfsTarget#usage_model}.
	UsageModel *string `field:"required" json:"usageModel" yaml:"usageModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#access_policy_name HpcCacheBlobNfsTarget#access_policy_name}.
	AccessPolicyName *string `field:"optional" json:"accessPolicyName" yaml:"accessPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#id HpcCacheBlobNfsTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#timeouts HpcCacheBlobNfsTarget#timeouts}
	Timeouts *HpcCacheBlobNfsTargetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#verification_timer_in_seconds HpcCacheBlobNfsTarget#verification_timer_in_seconds}.
	VerificationTimerInSeconds *float64 `field:"optional" json:"verificationTimerInSeconds" yaml:"verificationTimerInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/hpc_cache_blob_nfs_target#write_back_timer_in_seconds HpcCacheBlobNfsTarget#write_back_timer_in_seconds}.
	WriteBackTimerInSeconds *float64 `field:"optional" json:"writeBackTimerInSeconds" yaml:"writeBackTimerInSeconds"`
}

