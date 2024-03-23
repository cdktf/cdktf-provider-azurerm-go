// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlcontainer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbSqlContainerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#account_name CosmosdbSqlContainer#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#database_name CosmosdbSqlContainer#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#name CosmosdbSqlContainer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#partition_key_path CosmosdbSqlContainer#partition_key_path}.
	PartitionKeyPath *string `field:"required" json:"partitionKeyPath" yaml:"partitionKeyPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#resource_group_name CosmosdbSqlContainer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#analytical_storage_ttl CosmosdbSqlContainer#analytical_storage_ttl}.
	AnalyticalStorageTtl *float64 `field:"optional" json:"analyticalStorageTtl" yaml:"analyticalStorageTtl"`
	// autoscale_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#autoscale_settings CosmosdbSqlContainer#autoscale_settings}
	AutoscaleSettings *CosmosdbSqlContainerAutoscaleSettings `field:"optional" json:"autoscaleSettings" yaml:"autoscaleSettings"`
	// conflict_resolution_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#conflict_resolution_policy CosmosdbSqlContainer#conflict_resolution_policy}
	ConflictResolutionPolicy *CosmosdbSqlContainerConflictResolutionPolicy `field:"optional" json:"conflictResolutionPolicy" yaml:"conflictResolutionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#default_ttl CosmosdbSqlContainer#default_ttl}.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#id CosmosdbSqlContainer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// indexing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#indexing_policy CosmosdbSqlContainer#indexing_policy}
	IndexingPolicy *CosmosdbSqlContainerIndexingPolicy `field:"optional" json:"indexingPolicy" yaml:"indexingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#partition_key_version CosmosdbSqlContainer#partition_key_version}.
	PartitionKeyVersion *float64 `field:"optional" json:"partitionKeyVersion" yaml:"partitionKeyVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#throughput CosmosdbSqlContainer#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#timeouts CosmosdbSqlContainer#timeouts}
	Timeouts *CosmosdbSqlContainerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// unique_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/cosmosdb_sql_container#unique_key CosmosdbSqlContainer#unique_key}
	UniqueKey interface{} `field:"optional" json:"uniqueKey" yaml:"uniqueKey"`
}

