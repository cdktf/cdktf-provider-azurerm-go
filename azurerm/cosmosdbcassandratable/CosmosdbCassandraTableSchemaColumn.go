// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbcassandratable


type CosmosdbCassandraTableSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_cassandra_table#name CosmosdbCassandraTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_cassandra_table#type CosmosdbCassandraTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

