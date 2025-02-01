// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#cache_size_in_gb HpcCache#cache_size_in_gb}.
	CacheSizeInGb *float64 `field:"required" json:"cacheSizeInGb" yaml:"cacheSizeInGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#location HpcCache#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#name HpcCache#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#resource_group_name HpcCache#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#sku_name HpcCache#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#subnet_id HpcCache#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#automatically_rotate_key_to_latest_enabled HpcCache#automatically_rotate_key_to_latest_enabled}.
	AutomaticallyRotateKeyToLatestEnabled interface{} `field:"optional" json:"automaticallyRotateKeyToLatestEnabled" yaml:"automaticallyRotateKeyToLatestEnabled"`
	// default_access_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#default_access_policy HpcCache#default_access_policy}
	DefaultAccessPolicy *HpcCacheDefaultAccessPolicy `field:"optional" json:"defaultAccessPolicy" yaml:"defaultAccessPolicy"`
	// directory_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#directory_active_directory HpcCache#directory_active_directory}
	DirectoryActiveDirectory *HpcCacheDirectoryActiveDirectory `field:"optional" json:"directoryActiveDirectory" yaml:"directoryActiveDirectory"`
	// directory_flat_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#directory_flat_file HpcCache#directory_flat_file}
	DirectoryFlatFile *HpcCacheDirectoryFlatFile `field:"optional" json:"directoryFlatFile" yaml:"directoryFlatFile"`
	// directory_ldap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#directory_ldap HpcCache#directory_ldap}
	DirectoryLdap *HpcCacheDirectoryLdap `field:"optional" json:"directoryLdap" yaml:"directoryLdap"`
	// dns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#dns HpcCache#dns}
	Dns *HpcCacheDns `field:"optional" json:"dns" yaml:"dns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#id HpcCache#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#identity HpcCache#identity}
	Identity *HpcCacheIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#key_vault_key_id HpcCache#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#mtu HpcCache#mtu}.
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#ntp_server HpcCache#ntp_server}.
	NtpServer *string `field:"optional" json:"ntpServer" yaml:"ntpServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#tags HpcCache#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/hpc_cache#timeouts HpcCache#timeouts}
	Timeouts *HpcCacheTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

