// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbAccountConfig struct {
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
	// consistency_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#consistency_policy CosmosdbAccount#consistency_policy}
	ConsistencyPolicy *CosmosdbAccountConsistencyPolicy `field:"required" json:"consistencyPolicy" yaml:"consistencyPolicy"`
	// geo_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#geo_location CosmosdbAccount#geo_location}
	GeoLocation interface{} `field:"required" json:"geoLocation" yaml:"geoLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#location CosmosdbAccount#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#name CosmosdbAccount#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#offer_type CosmosdbAccount#offer_type}.
	OfferType *string `field:"required" json:"offerType" yaml:"offerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#resource_group_name CosmosdbAccount#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#access_key_metadata_writes_enabled CosmosdbAccount#access_key_metadata_writes_enabled}.
	AccessKeyMetadataWritesEnabled interface{} `field:"optional" json:"accessKeyMetadataWritesEnabled" yaml:"accessKeyMetadataWritesEnabled"`
	// analytical_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#analytical_storage CosmosdbAccount#analytical_storage}
	AnalyticalStorage *CosmosdbAccountAnalyticalStorage `field:"optional" json:"analyticalStorage" yaml:"analyticalStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#analytical_storage_enabled CosmosdbAccount#analytical_storage_enabled}.
	AnalyticalStorageEnabled interface{} `field:"optional" json:"analyticalStorageEnabled" yaml:"analyticalStorageEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#automatic_failover_enabled CosmosdbAccount#automatic_failover_enabled}.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#backup CosmosdbAccount#backup}
	Backup *CosmosdbAccountBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#burst_capacity_enabled CosmosdbAccount#burst_capacity_enabled}.
	BurstCapacityEnabled interface{} `field:"optional" json:"burstCapacityEnabled" yaml:"burstCapacityEnabled"`
	// capabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#capabilities CosmosdbAccount#capabilities}
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#capacity CosmosdbAccount#capacity}
	Capacity *CosmosdbAccountCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#cors_rule CosmosdbAccount#cors_rule}
	CorsRule *CosmosdbAccountCorsRule `field:"optional" json:"corsRule" yaml:"corsRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#create_mode CosmosdbAccount#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#default_identity_type CosmosdbAccount#default_identity_type}.
	DefaultIdentityType *string `field:"optional" json:"defaultIdentityType" yaml:"defaultIdentityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#free_tier_enabled CosmosdbAccount#free_tier_enabled}.
	FreeTierEnabled interface{} `field:"optional" json:"freeTierEnabled" yaml:"freeTierEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#id CosmosdbAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#identity CosmosdbAccount#identity}
	Identity *CosmosdbAccountIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#ip_range_filter CosmosdbAccount#ip_range_filter}.
	IpRangeFilter *[]*string `field:"optional" json:"ipRangeFilter" yaml:"ipRangeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#is_virtual_network_filter_enabled CosmosdbAccount#is_virtual_network_filter_enabled}.
	IsVirtualNetworkFilterEnabled interface{} `field:"optional" json:"isVirtualNetworkFilterEnabled" yaml:"isVirtualNetworkFilterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#key_vault_key_id CosmosdbAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#kind CosmosdbAccount#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#local_authentication_disabled CosmosdbAccount#local_authentication_disabled}.
	LocalAuthenticationDisabled interface{} `field:"optional" json:"localAuthenticationDisabled" yaml:"localAuthenticationDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#managed_hsm_key_id CosmosdbAccount#managed_hsm_key_id}.
	ManagedHsmKeyId *string `field:"optional" json:"managedHsmKeyId" yaml:"managedHsmKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#minimal_tls_version CosmosdbAccount#minimal_tls_version}.
	MinimalTlsVersion *string `field:"optional" json:"minimalTlsVersion" yaml:"minimalTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#mongo_server_version CosmosdbAccount#mongo_server_version}.
	MongoServerVersion *string `field:"optional" json:"mongoServerVersion" yaml:"mongoServerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#multiple_write_locations_enabled CosmosdbAccount#multiple_write_locations_enabled}.
	MultipleWriteLocationsEnabled interface{} `field:"optional" json:"multipleWriteLocationsEnabled" yaml:"multipleWriteLocationsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#network_acl_bypass_for_azure_services CosmosdbAccount#network_acl_bypass_for_azure_services}.
	NetworkAclBypassForAzureServices interface{} `field:"optional" json:"networkAclBypassForAzureServices" yaml:"networkAclBypassForAzureServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#network_acl_bypass_ids CosmosdbAccount#network_acl_bypass_ids}.
	NetworkAclBypassIds *[]*string `field:"optional" json:"networkAclBypassIds" yaml:"networkAclBypassIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#partition_merge_enabled CosmosdbAccount#partition_merge_enabled}.
	PartitionMergeEnabled interface{} `field:"optional" json:"partitionMergeEnabled" yaml:"partitionMergeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#public_network_access_enabled CosmosdbAccount#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// restore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#restore CosmosdbAccount#restore}
	Restore *CosmosdbAccountRestore `field:"optional" json:"restore" yaml:"restore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#tags CosmosdbAccount#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#timeouts CosmosdbAccount#timeouts}
	Timeouts *CosmosdbAccountTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_network_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/cosmosdb_account#virtual_network_rule CosmosdbAccount#virtual_network_rule}
	VirtualNetworkRule interface{} `field:"optional" json:"virtualNetworkRule" yaml:"virtualNetworkRule"`
}

