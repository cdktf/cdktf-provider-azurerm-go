// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccachenfstarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheNfsTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#cache_name HpcCacheNfsTarget#cache_name}.
	CacheName *string `field:"required" json:"cacheName" yaml:"cacheName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#name HpcCacheNfsTarget#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// namespace_junction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#namespace_junction HpcCacheNfsTarget#namespace_junction}
	NamespaceJunction interface{} `field:"required" json:"namespaceJunction" yaml:"namespaceJunction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#resource_group_name HpcCacheNfsTarget#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#target_host_name HpcCacheNfsTarget#target_host_name}.
	TargetHostName *string `field:"required" json:"targetHostName" yaml:"targetHostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#usage_model HpcCacheNfsTarget#usage_model}.
	UsageModel *string `field:"required" json:"usageModel" yaml:"usageModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#id HpcCacheNfsTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/hpc_cache_nfs_target#timeouts HpcCacheNfsTarget#timeouts}
	Timeouts *HpcCacheNfsTargetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

