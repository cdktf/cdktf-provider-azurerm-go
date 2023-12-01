// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbcassandradatacenter


type CosmosdbCassandraDatacenterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cosmosdb_cassandra_datacenter#create CosmosdbCassandraDatacenter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cosmosdb_cassandra_datacenter#delete CosmosdbCassandraDatacenter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cosmosdb_cassandra_datacenter#read CosmosdbCassandraDatacenter#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/cosmosdb_cassandra_datacenter#update CosmosdbCassandraDatacenter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

