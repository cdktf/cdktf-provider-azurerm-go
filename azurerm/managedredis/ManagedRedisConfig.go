// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedredis

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedRedisConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#location ManagedRedis#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#name ManagedRedis#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#resource_group_name ManagedRedis#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#sku_name ManagedRedis#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#customer_managed_key ManagedRedis#customer_managed_key}
	CustomerManagedKey *ManagedRedisCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// default_database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#default_database ManagedRedis#default_database}
	DefaultDatabase *ManagedRedisDefaultDatabase `field:"optional" json:"defaultDatabase" yaml:"defaultDatabase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#high_availability_enabled ManagedRedis#high_availability_enabled}.
	HighAvailabilityEnabled interface{} `field:"optional" json:"highAvailabilityEnabled" yaml:"highAvailabilityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#id ManagedRedis#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#identity ManagedRedis#identity}
	Identity *ManagedRedisIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#public_network_access ManagedRedis#public_network_access}.
	PublicNetworkAccess *string `field:"optional" json:"publicNetworkAccess" yaml:"publicNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#tags ManagedRedis#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/managed_redis#timeouts ManagedRedis#timeouts}
	Timeouts *ManagedRedisTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

