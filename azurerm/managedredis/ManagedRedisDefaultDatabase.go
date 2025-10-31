// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedredis


type ManagedRedisDefaultDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#access_keys_authentication_enabled ManagedRedis#access_keys_authentication_enabled}.
	AccessKeysAuthenticationEnabled interface{} `field:"optional" json:"accessKeysAuthenticationEnabled" yaml:"accessKeysAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#client_protocol ManagedRedis#client_protocol}.
	ClientProtocol *string `field:"optional" json:"clientProtocol" yaml:"clientProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#clustering_policy ManagedRedis#clustering_policy}.
	ClusteringPolicy *string `field:"optional" json:"clusteringPolicy" yaml:"clusteringPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#eviction_policy ManagedRedis#eviction_policy}.
	EvictionPolicy *string `field:"optional" json:"evictionPolicy" yaml:"evictionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#geo_replication_group_name ManagedRedis#geo_replication_group_name}.
	GeoReplicationGroupName *string `field:"optional" json:"geoReplicationGroupName" yaml:"geoReplicationGroupName"`
	// module block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/managed_redis#module ManagedRedis#module}
	Module interface{} `field:"optional" json:"module" yaml:"module"`
}

