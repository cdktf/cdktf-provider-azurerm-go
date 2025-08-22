// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbcassandratable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbCassandraTableConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#cassandra_keyspace_id CosmosdbCassandraTable#cassandra_keyspace_id}.
	CassandraKeyspaceId *string `field:"required" json:"cassandraKeyspaceId" yaml:"cassandraKeyspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#name CosmosdbCassandraTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#schema CosmosdbCassandraTable#schema}
	Schema *CosmosdbCassandraTableSchema `field:"required" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#analytical_storage_ttl CosmosdbCassandraTable#analytical_storage_ttl}.
	AnalyticalStorageTtl *float64 `field:"optional" json:"analyticalStorageTtl" yaml:"analyticalStorageTtl"`
	// autoscale_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#autoscale_settings CosmosdbCassandraTable#autoscale_settings}
	AutoscaleSettings *CosmosdbCassandraTableAutoscaleSettings `field:"optional" json:"autoscaleSettings" yaml:"autoscaleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#default_ttl CosmosdbCassandraTable#default_ttl}.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#id CosmosdbCassandraTable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#throughput CosmosdbCassandraTable#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/cosmosdb_cassandra_table#timeouts CosmosdbCassandraTable#timeouts}
	Timeouts *CosmosdbCassandraTableTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

